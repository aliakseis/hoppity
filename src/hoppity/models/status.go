package models

// Represents JSON status object returned by Rest Service
// swagger:response status
type Status struct {
	Status   		string		`json:"status"`
	StatusMessage 	string		`json:"status_message"`
}
