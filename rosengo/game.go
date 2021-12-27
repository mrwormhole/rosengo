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
	activeScene   Scene
	allScenes     []Scene
	audioManager  *manager.AudioManager
	spriteManager *manager.SpriteManager
}

func NewGame() (*Game, error) {
	sampleRate := 48000
	audioManager, err := manager.NewAudioManager(sampleRate)
	if err != nil {
		return nil, fmt.Errorf("rosengo.NewGame: %v", err)
	}
	err = audioManager.LoadAll("sounds")
	if err != nil {
		return nil, fmt.Errorf("rosengo.NewGame: %v", err)
	}
	// Note: audio players should be closed before they are played????

	spriteManager, err := manager.NewSpriteManager()
	if err != nil {
		return nil, fmt.Errorf("rosengo.NewGame: %v", err)
	}
	err = spriteManager.LoadAll("images")
	if err != nil {
		return nil, fmt.Errorf("rosengo.NewGame: %v", err)
	}

	// --------------------------------- TESTING AREA ---------------------------------
	dummyImage, err := spriteManager.GetImage("0")
	if err != nil {
		return nil, fmt.Errorf("rosengo.NewGame: %v", err)
	}
	w, h := dummyImage.Size()
	dummyGameObject, err := NewGameObject(dummyImage, ScreenWidth/2-w/2, ScreenHeight/2-h/2)
	if err != nil {
		return nil, fmt.Errorf("rosengo.NewGame: %v", err)
	}
	s := NewScene("introduction", []*GameObject{dummyGameObject}, Starting)
	scenes := []Scene{s}
	// --------------------------------- TESTING AREA ---------------------------------

	return &Game{
		transparent:   false,
		mute:          false,
		activeScene:   scenes[0],
		allScenes:     scenes,
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

// Layout is implemented to obtain frame for Ebiten interface
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

// Update is implemented to update for Ebiten interface
func (g *Game) Update() error {
	return nil
}

// Draw is implemented to draw for Ebiten interface
func (g *Game) Draw(screen *ebiten.Image) {
	g.activeScene.Draw(screen)
}
