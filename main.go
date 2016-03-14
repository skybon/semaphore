package semaphore

// Semaphore is a channel-based semaphore unit.
type Semaphore chan struct{}

// Acquire decreases semaphore's free access count.
func (s Semaphore) Acquire() {
	<-s
}

// Release increases semaphore's free access count.
func (s Semaphore) Release() {
	s <- struct{}{}
}

// Exec can be used for safely executing a callback, securing the access with the semaphore.
func (s Semaphore) Exec(cb func()) {
	s.Acquire()
	defer s.Release()

	cb()
}

// LimitCount returns number of concurrent actions possible under the semaphore.
func (s Semaphore) LimitCount() int {
	return cap(s)
}

// MakeSemaphore creates a new Semaphore.
func MakeSemaphore(allowance int) Semaphore {
	var sem = make(Semaphore, allowance)
	for i := 0; i < allowance; i++ {
		sem <- struct{}{}
	}

	return sem
}
