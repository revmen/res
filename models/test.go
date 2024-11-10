package models

import "res/methods"

type Test struct {
	Name   string         `json:"name"`
	Method methods.Method `json:"method"`
}
