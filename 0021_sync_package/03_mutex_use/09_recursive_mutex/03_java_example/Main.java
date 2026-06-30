import java.util.concurrent.locks.ReentrantLock;

public class Main {
    public static void main(String[] args) {
        Reentrant r = new Reentrant();
        r.outer();

        System.out.println("All done!");
    }
}

class Reentrant {
    private final ReentrantLock lock = new ReentrantLock();

    public void outer() {
        lock.lock();

        try {
            inner();
        } finally {
            lock.unlock();
        }
    }

    public void inner() {
        lock.lock();

        try {
            // action
        } finally {
            lock.unlock();
        }
    }
}
