package models

type Candidate struct {
	Name  string  `json:"name"`
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
	Power float64 `json:"power"`
}
