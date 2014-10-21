package treewatch_test

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/lucas-clemente/treewatch"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Treewatch", func() {
	dir, err := filepath.Abs("testdata")
	if err != nil {
		panic("could not find testdata/")
	}

	filename := dir + "/foo"

	AfterEach(func() {
		os.Remove(filename)
	})

	It("does nothing when nothing happens", func() {
		watcher, err := treewatch.NewTreeWatcher(dir)
		Expect(err).To(BeNil())
		watcher.Stop()
		_, ok := <-watcher.Changes()
		Expect(ok).To(BeFalse())
	})

	It("listens for changes", func() {
		watcher, err := treewatch.NewTreeWatcher(dir)
		Expect(err).To(BeNil())
		Expect(ioutil.WriteFile(filename, []byte("foo"), 0644)).To(BeNil())
		Expect(<-watcher.Changes()).To(Equal(filename))
		watcher.Stop()
		_, ok := <-watcher.Changes()
		Expect(ok).To(BeFalse())
	})
})
