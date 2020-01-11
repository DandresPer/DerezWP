package models
// Stash : stashes in FullResponse
type Stash struct {
	ID                string
	Public            bool
	AccountName       string
	LastCharacterName string
	Stash             string
	StashType         string
	League            string
	Items             []Item
}