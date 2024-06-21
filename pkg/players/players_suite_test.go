package players_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPlayers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Players Suite")
}
