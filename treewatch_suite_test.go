package treewatch_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTreewatch(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Treewatch Suite")
}
