package rosengo

import (
	"flag"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 480
	ScreenHeight = 640
)

var (
	cpuProfile = flag.String("cpuprofile", "", "write cpu profile to file")
)

type Game struct {
	transparent bool
	mute        bool
	scene       *Scene
}

func NewGame() (*Game, error) {
	return &Game{
		transparent: false,
		mute:        false,
		scene:       nil,
	}, nil
}

func (g *Game) SetTransparent(value bool) {
	g.transparent = value
}

func (g *Game) SetMute(value bool) {
	g.mute = value
}

// Layout is implented to obtain frame for Ebiten interface
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

// Update is implemented to update for Ebiten interface
func (g *Game) Update() error {
	return nil
}

// Draw is implemented to draw for Ebiten interface
func (g *Game) Draw(screen *ebiten.Image) {
	return
}
