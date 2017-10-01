package microdungeon

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMicroDungeon(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MicroDungeon Tests")
}
