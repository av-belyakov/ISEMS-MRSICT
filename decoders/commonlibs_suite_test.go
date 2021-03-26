package decoders_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCommonlibs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Commonlibs Suite")
}
