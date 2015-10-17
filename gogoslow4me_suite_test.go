package gogoslow4me_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGogoslow4me(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gogoslow4me Suite")
}
