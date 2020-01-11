package models
// FullResponse : main chunk of data
type FullResponse struct {
	NextChangeID string `json:"next_change_id"`
	Stashes      []Stash
}