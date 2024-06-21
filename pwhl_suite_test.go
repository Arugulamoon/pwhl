package pwhl_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPwhl(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pwhl Suite")
}
