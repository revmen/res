package models

// A frame is a one collection of generated candidates and voters

type Frame struct {
	Candidates []Candidate `json:"candidates"`
	Voters     []Voter     `json:"voters"`
}

func (f *Frame) AddCandidate(c Candidate) {
	f.Candidates = append(f.Candidates, c)
}
