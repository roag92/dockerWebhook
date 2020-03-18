package server

const StateSuccess = "success"
const StateFailure = "failure"
const StateError = "error"

type Response struct {
	State string `json:"state"`
	Description string `json:"description"`
	Context string `json:"context"`
	TargetUrl string `json:"target_url"`
}
