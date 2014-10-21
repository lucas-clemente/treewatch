// +build cgo

package treewatch

/*
#cgo LDFLAGS: -framework CoreFoundation -framework CoreServices

#include <CoreFoundation/CoreFoundation.h>
#include <CoreServices/CoreServices.h>

void treeWatcherCallbackC(ConstFSEventStreamRef streamRef, void *clientCallBackInfo, size_t numEvents, void *eventPaths, const FSEventStreamEventFlags eventFlags[], const FSEventStreamEventId eventIds[]);
*/
import "C"

import (
	"runtime"
	"unsafe"
)

type treeWatcherDarwin struct {
	changes chan struct{}
	runloop C.CFRunLoopRef
}

func newTreeWatcherImpl(path string) (TreeWatcher, error) {
	pathRef := C.CFStringCreateWithCString(nil, C.CString(path), C.kCFStringEncodingASCII)
	cPathsArray := unsafe.Pointer(pathRef)
	pathsRef := C.CFArrayCreate(nil, &cPathsArray, 1, nil)

	tw := &treeWatcherDarwin{
		changes: make(chan struct{}),
	}

	setupDone := make(chan struct{})

	// This thread (!) acts as CF runloop
	// The run loop is terminated from Stop()
	go func() {
		runtime.LockOSThread()
		stream := C.FSEventStreamCreate(
			nil,
			C.FSEventStreamCallback(unsafe.Pointer(C.treeWatcherCallbackC)),
			&C.FSEventStreamContext{version: 0, info: unsafe.Pointer(tw)},
			pathsRef,
			C.FSEventStreamEventId(uint64(0xFFFFFFFFFFFFFFFF)),
			C.CFTimeInterval(0.1),
			C.kFSEventStreamCreateFlagFileEvents|C.kFSEventStreamCreateFlagNoDefer,
		)
		tw.runloop = C.CFRunLoopGetCurrent()
		C.FSEventStreamScheduleWithRunLoop(stream, tw.runloop, C.kCFRunLoopDefaultMode)
		C.FSEventStreamStart(stream)
		setupDone <- struct{}{}
		C.CFRunLoopRun()
		C.FSEventStreamStop(stream)
		C.FSEventStreamUnscheduleFromRunLoop(stream, tw.runloop, C.kCFRunLoopDefaultMode)
		C.FSEventStreamInvalidate(stream)
		C.FSEventStreamRelease(stream)
		close(tw.changes)
	}()

	<-setupDone

	return tw, nil
}

//export treeWatcherCallback
func treeWatcherCallback(treeWatcherC unsafe.Pointer) {
	tw := (*treeWatcherDarwin)(treeWatcherC)
	tw.changes <- struct{}{}
}

func (tw *treeWatcherDarwin) Changes() <-chan struct{} {
	return tw.changes
}

func (tw *treeWatcherDarwin) Stop() {
	C.CFRunLoopStop(tw.runloop)
}
