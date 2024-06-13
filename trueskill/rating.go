package trueskill

import (
	"fmt"

	"goskill/gaussian"
)

// Rating is a player with a certain skill (mu, sigma).
type Rating struct {
	gaussian.Gaussian
}

// NewRating returns a player from the provided mu (mean) and sigma
// (standard deviation).
func NewRating(mu, sigma float64) Rating {
	return Rating{
		Gaussian: gaussian.NewFromMeanAndStdDev(mu, sigma),
	}
}

// Mu returns the player mean.
func (r Rating) Mu() float64 {
	return r.Mean()
}

// Sigma returns the player standard deviation.
func (r Rating) Sigma() float64 {
	return r.StdDev()
}

func (r Rating) String() string {
	return fmt.Sprintf("Player(mu=%.3f sigma=%.3f)", r.Mu(), r.Sigma())
}
