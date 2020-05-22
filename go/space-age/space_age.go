package space

// Planet represent a planet in the solar system.
type Planet string

type orbitalPeriod float64

const (
	mercury orbitalPeriod = 0.2408467
	venus   orbitalPeriod = 0.61519726
	earth   orbitalPeriod = 1.0
	mars    orbitalPeriod = 1.8808158
	jupiter orbitalPeriod = 11.862615
	saturn  orbitalPeriod = 29.447498
	uranus  orbitalPeriod = 84.016846
	neptune orbitalPeriod = 164.79132

	earthYearSeconds = 365.25 * 24 * 60 * 60
)

// Age converts seconds into the age in years on the specified planet.
func Age(s float64, p Planet) float64 {
	var o orbitalPeriod
	switch p {
	case "Mercury":
		o = mercury
	case "Venus":
		o = venus
	case "Mars":
		o = mars
	case "Jupiter":
		o = jupiter
	case "Saturn":
		o = saturn
	case "Uranus":
		o = uranus
	case "Neptune":
		o = neptune
	default:
		o = earth
	}
	planetSeconds := earthYearSeconds * float64(o)
	return s / planetSeconds
}
