package decoders_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDecoders(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Decoders Suite")
}
