package rosengo

type GameState int

const (
	Starting GameState = iota
	Ending
	OnMenu
	InGame
)
