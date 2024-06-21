package players

type Availability string

const AVAILABILITY_FREE_AGENT Availability = "Free Agent"

type Player struct{}

func (p *Player) Availability() Availability {
	return AVAILABILITY_FREE_AGENT
}
