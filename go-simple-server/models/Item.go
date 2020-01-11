package models
// Item : items in stash
type Item struct {
	Verified   bool
	W          int
	H          int
	Icon       string
	League     string
	ID         string
	Name       string
	TypeLine   string
	Identified string
	Ilvl       string
	Properties []Property
	// //ExplicitiMods
	// //flavourText
	// //frameType
	StackSize    int
	MaxStackSize int
	ArtFileName  string
	//extended
	X           int
	Y           int
	InventoryID string
}