package theme

// Color Roles in a Theme
const (
	Foreground = iota
	Background
	Base
	AlternateBase
	Text
	Selection
	Highlight
)

// Role is used for type-safety
type Role int
