package interactionmongodb_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestInteractionmongodb(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Interactionmongodb Suite")
}
