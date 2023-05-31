package interactionmongodb_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestInteractionmongodb(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Interactionmongodb Suite")
}
