package treewatch

/*

#include <CoreServices/CoreServices.h>

// To make gcc happy
void treeWatcherCallback(void*);

// Bridge to Go code
void treeWatcherCallbackC(ConstFSEventStreamRef streamRef, void *clientCallBackInfo, size_t numEvents, void *eventPaths, const FSEventStreamEventFlags eventFlags[], const FSEventStreamEventId eventIds[]) {
  treeWatcherCallback(clientCallBackInfo);
}

*/
import "C"
