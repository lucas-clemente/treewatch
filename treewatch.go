package treewatch

// A TreeWatcher watches a file system tree and sends notifications
// when anything happens.
type TreeWatcher interface {
	// The path of the directory the change happened in is sent here.
	Changes() <-chan string

	// Stop the tree watcher
	Stop()
}

// NewTreeWatcher creates and starts a new treewatcher for a given directory.
func NewTreeWatcher(path string) (TreeWatcher, error) {
	return newTreeWatcherImpl(path)
}
