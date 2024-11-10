package models

type Experiment struct {
	Name          string         `json:"name"`
	Notes         string         `json:"notes"`
	VoterGens     []VoterGen     `json:"voterGens"`
	CandidateGens []CandidateGen `json:"candidateGens"`
	Frames        []Frame        `json:"frames"`
}
