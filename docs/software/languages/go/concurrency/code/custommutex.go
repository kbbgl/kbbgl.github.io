package main

import "sync"

type ReadWriteMutex struct {
	// count number of reader goroutines currently in critical section
	readersCounter int

	// mutex for synchronizing readers' access
	readersLock sync.Mutex

	// mutex for blocking any writers' access
	globalLock sync.Mutex
}

func (rw *ReadWriteMutex) ReadLock() {
	
	// synchronizes access so that only one
	// goroutine is allowed at any time
	rw.readersLock.Lock()

	rw.readersCounter++

	if rw.readersCounter == 1 {
		rw.globalLock.Lock()
	}
	rw.readersLock.Unlock()
}

func (rw *ReadWriteMutex) WriteLock() {
	rw.globalLock.Lock()
}

func (rw *ReadWriteMutex) ReadUnlock() {
	rw.readersLock.Lock()
	rw.readersCounter--
	if rw.readersCounter == {
		rw.globalLock.Unlock()
	}
	rw.readersLock.Unlock()
}

func (rw *ReadWriteMutex) WriteUnlock() {
	rw.globalLock.Unlock()
}