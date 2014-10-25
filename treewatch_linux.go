// +build cgo

package treewatch

/*
#cgo LDFLAGS: -linotifytools

#include <stdlib.h>
#include <inotifytools/inotifytools.h>
#include <inotifytools/inotify.h>

*/
import "C"
import (
	"fmt"
	"unsafe"
)

type treeWatcherLinux struct {
	changes chan string
	stopped bool
}

func newTreeWatcherImpl(path string) (TreeWatcher, error) {
	tw := &treeWatcherLinux{
		changes: make(chan string),
		stopped: false,
	}

	errC := C.inotifytools_initialize()
	if errC != 1 {
		return nil, makeInotifyErr("inotifytools_initialize")
	}

	errC = C.inotifytools_watch_recursively(C.CString(path), C.IN_CREATE|C.IN_DELETE|C.IN_MODIFY|C.IN_MOVED_FROM|C.IN_MOVED_TO)
	if errC != 1 {
		return nil, makeInotifyErr("inotifytools_watch_recursively")
	}

	format := C.CString("%w%f")
	go func() {
		for !tw.stopped {
			event := C.inotifytools_next_event(1)
			if event == nil {
				continue
			}
			const l = 4096
			var fname = (*C.char)(C.malloc(l))
			writtenLen := C.inotifytools_snprintf(fname, l, event, format)
			tw.changes <- C.GoStringN(fname, writtenLen+1)
			C.free(unsafe.Pointer(fname))
		}
		close(tw.changes)
	}()

	return tw, nil
}

func makeInotifyErr(f string) error {
	return fmt.Errorf("%s() failed with error %d", f, C.inotifytools_error())
}

func (tw *treeWatcherLinux) Changes() <-chan string {
	return tw.changes
}

func (tw *treeWatcherLinux) Stop() {
	tw.stopped = true
	// Read all remaining
	go func() {
		for _ = range <-tw.changes {
		}
	}()
}
