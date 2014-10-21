package treewatch

// A TreeWatcher watches a file system tree and sends notifications
// when anything happens.
type TreeWatcher interface {
	// A struct{} is sent every time something in the tree changes.
	Changes() <-chan struct{}

	// Stop the tree watcher
	Stop()
}

// NewTreeWatcher creates and starts a new treewatcher for a given directory.
func NewTreeWatcher(path string) (TreeWatcher, error) {
	return newTreeWatcherImpl(path)
}
