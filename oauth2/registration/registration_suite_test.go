package registration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRegistration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Registration Suite")
}
