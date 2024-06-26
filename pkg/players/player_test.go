package players_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/Arugulamoon/pwhl/pkg/players"
)

var _ = Describe("Player", func() {
	var jenner *players.Player

	Describe("availability", func() {
		Context("when unsigned", func() {
			It("is free agent", func() {
				Expect(jenner.Availability()).
					To(Equal(players.AVAILABILITY_FREE_AGENT))
			})
		})

		Context("when signed", func() {
			It("is not free agent", Pending, func() {})
		})
	})
})
