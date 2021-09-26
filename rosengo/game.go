package rosengo

import (
	"fmt"
	"github.com/MrWormHole/rosengo/rosengo/manager"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 480
	ScreenHeight = 640
)

type Game struct {
	transparent   bool
	mute          bool
	activeScene   *Scene
	allScenes     []*Scene
	audioManager  manager.AudioManager
	spriteManager manager.SpriteManager
}

func NewGame() (*Game, error) {
	audioManager, err := manager.NewAudioManager(48000)
	if err != nil {
		return nil, fmt.Errorf("rosengo.NewGame: %v", err)
	}
	err = audioManager.Load("assets/sounds")
	if err != nil {
		return nil, fmt.Errorf("rosengo.NewGame: %v", err)
	}
	// Note: audio players should be closed before they are played????

	spriteManager, err := manager.NewSpriteManager()
	if err != nil {
		return nil, fmt.Errorf("rosengo.NewGame: %v", err)
	}
	err = spriteManager.Load("assets/images")
	if err != nil {
		return nil, fmt.Errorf("rosengo.NewGame: %v", err)
	}

	return &Game{
		transparent:   false,
		mute:          false,
		activeScene:   nil, // TODO: decide on this
		allScenes:     nil, // TODO: decide on this
		audioManager:  audioManager,
		spriteManager: spriteManager,
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
