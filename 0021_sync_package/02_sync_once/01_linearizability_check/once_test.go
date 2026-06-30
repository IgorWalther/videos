package linearizability

import (
	"fmt"
	"os"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/anishathalye/porcupine"
	"github.com/stretchr/testify/require"
)

const (
	reservedValue = -27
	expectedValue = 27
)

var modelOnce = porcupine.Model{
	Init: func() any { return reservedValue },
	Step: func(state any, input, output any) (bool, any) {
		if state == reservedValue {
			return true, output
		}

		return output.(int64) == expectedValue, output
	},
	Equal: func(a, b any) bool { return a == b },
	DescribeOperation: func(input, output any) string {
		return fmt.Sprintf("Once() -> '%d'", output.(int64))
	},
}

func TestBrokenOnce_RealViolation(t *testing.T) {
	o := sync.Once{}

	events := make([]porcupine.Operation, 0, 4)
	eventsMx := new(sync.Mutex)

	wg := new(sync.WaitGroup)
	holder := atomic.Int64{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			if i == 0 {
				time.Sleep(10 * time.Millisecond)
			}

			start := time.Now().UnixNano()

			o.Do(func() {
				time.Sleep(3000 * time.Millisecond)
				holder.Store(expectedValue)
			})

			end := time.Now().UnixNano()
			result := holder.Load()

			func() {
				eventsMx.Lock()
				defer eventsMx.Unlock()

				events = append(events, porcupine.Operation{
					ClientId: i,
					Call:     start,
					Return:   end,
					Output:   result,
				})
			}()
		}()
	}

	wg.Wait()

	res, info := porcupine.CheckOperationsVerbose(modelOnce, events, 0)

	f, err := os.Create("./info.html")
	require.NoError(t, err)

	err = porcupine.Visualize(modelOnce, info, f)
	require.NoError(t, err)

	require.Equal(t, porcupine.Ok, res)
}
