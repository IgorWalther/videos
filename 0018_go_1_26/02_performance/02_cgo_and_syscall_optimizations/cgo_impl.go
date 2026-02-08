package perfbench

/*
static inline void cgobench_nop(void) {}

// Go function that C will call back into.
extern void cgobench_callback(void);

static inline void cgobench_call_callback(void) {
    cgobench_callback();
}
*/
import "C"

//export cgobench_callback
func cgobench_callback() {}

func cgoNop() {
	C.cgobench_nop()
}

func cgoCallWithCallback() {
	C.cgobench_call_callback()
}
