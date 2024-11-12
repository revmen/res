package models

type Experiment struct {
	Name          string         `json:"name"`
	Notes         string         `json:"notes"`
	VoterGens     []VoterGen     `json:"voterGens"`
	CandidateGens []CandidateGen `json:"candidateGens"`
	Frames        []Frame        `json:"frames"`
}

func NewExperiment() *Experiment {
	return &Experiment{
		Name: "New Experiment",
		// Notes:         "",  Start with empty notes
		VoterGens:     []VoterGen{},
		CandidateGens: []CandidateGen{},
		Frames:        []Frame{},
	}
}

func (e *Experiment) Get() *Experiment {
	return e
}

func (e *Experiment) Save(exp *Experiment) {
	// Overwrite the current experiment with the new one
	e.Name = exp.Name
	e.Notes = exp.Notes
	e.VoterGens = exp.VoterGens
	e.CandidateGens = exp.CandidateGens
	e.Frames = exp.Frames
}
