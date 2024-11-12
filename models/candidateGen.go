package models

import (
	"math"
	"math/rand"
	"time"
)

type CandidateGen struct {
	Name         string  `json:"name"`
	X            float64 `json:"x"`
	Y            float64 `json:"y"`
	Radius       float64 `json:"radius"`
	Distribution string  `json:"distribution"` // "uniform" or "normal"
	Power        float64 `json:"power"`
	Color        string  `json:"color"`
}

// Method to generate a Candidate with uniform distribution
func (cg *CandidateGen) GenerateCandidate() Candidate {
	// Create a new random generator with a unique seed
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	var candidateX, candidateY float64

	// Check if distribution is set to "uniform"
	if cg.Distribution == "uniform" {
		// Generate a random angle between 0 and 2π
		angle := rng.Float64() * 2 * math.Pi
		// Generate a random radius (scaled to the outer boundary of the circle)
		r := cg.Radius * math.Sqrt(rng.Float64())

		// Convert polar coordinates (r, angle) to Cartesian coordinates (x, y)
		candidateX = cg.X + r*math.Cos(angle)
		candidateY = cg.Y + r*math.Sin(angle)
	}

	// Create a new Candidate based on the generated position and properties of CandidateGen
	return Candidate{
		Name:  cg.Name,
		X:     candidateX,
		Y:     candidateY,
		Power: cg.Power,
	}
}
