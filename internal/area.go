package example

type figures int

// define and export the constants for the figures
const (
	Square figures = iota
	Circle
	Triangle
)

func Area(f figures) (func(float64) float64, bool) {
	switch f {
	case Square:
		return func(a float64) float64 {
			return a * a
		}, true
	case Circle:
		return func(r float64) float64 {
			return 3.14 * r * r
		}, true
	case Triangle:
		return func(a float64) float64 {
			return (a * a * 1.732) / 4
		}, true
	default:
		return nil, false
	}
}
