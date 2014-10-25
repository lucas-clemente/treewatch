# Treewatch

[![GoDoc](https://godoc.org/github.com/lucas-clemente/treewatch?status.svg)](https://godoc.org/github.com/lucas-clemente/treewatch)
[![Build Status](https://travis-ci.org/lucas-clemente/treewatch.svg?branch=master)](https://travis-ci.org/lucas-clemente/treewatch)

Watch a file system tree for changes.

```go
watcher, err := NewTreeWatcher("/path/to/dir")
c := watcher.Changes()

filename := <-c
// filename changed (e.g. created, modified, removed)

watcher.Stop()
```

Currently supports:

- Darwin (OS X)
- Linux (requires `libinotifytools`)

Todo:

- Detailed events
