package models

type VoterGen struct {
	Name         string  `json:"name"`
	X            float64 `json:"x"`
	Y            float64 `json:"y"`
	Radius       float64 `json:"radius"`
	Distribution string  `json:"distribution"` // "uniform" or "weighted"
	Strategic    float64 `json:"strategic"`
	Color        string  `json:"color"`
}
