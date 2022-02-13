package rosengo

import (
	"fmt"
	"github.com/MrWormHole/rosengo/rosengo/manager"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	"image/color"
)

const (
	ScreenWidth  = 900
	ScreenHeight = 1600
)

var (
	OutsideWidth  int
	OutsideHeight int
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
	// IMAGE TEST
	/*dummyImage, err := spriteManager.GetImage("milky-way640x480")
	if err != nil {
		return nil, fmt.Errorf("rosengo.NewGame: %v", err)
	}
	w, h := dummyImage.Size()
	fmt.Println("Image size:", w, h)*/

	r := 500
	offset := 5
	test := ebiten.NewImage(r, r)
	test.Fill(color.White)
	sub := test.SubImage(image.Rect(offset, offset, r-offset, r-offset)).(*ebiten.Image)
	sub.Fill(color.Black)
	ebitenutil.DebugPrint(test, " test")

	test2 := ebiten.NewImage(411, 411)
	test2.Fill(color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 255,
	})

	// IMAGE TEST
	// GAMEOBJECT TEST
	dummyGameObject, err := NewGameObject(test2, ScreenWidth/2-411/2, ScreenHeight/2-411/2)
	if err != nil {
		return nil, fmt.Errorf("rosengo.NewGame: %v", err)
	}
	// GAMEOBJECT TEST

	// SCENE TEST
	s := NewScene("introduction", []*GameObject{dummyGameObject}, Starting)
	scenes := []Scene{s}
	// SCENE TEST
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
	OutsideWidth = outsideWidth
	OutsideHeight = outsideHeight

	return ScreenWidth, ScreenHeight
}

// Update is implemented to update for Ebiten interface
func (g *Game) Update() error {
	fmt.Println("SUPERNOVA: ", OutsideWidth, OutsideHeight)
	touchIDs := ebiten.TouchIDs()
	for _, id := range touchIDs {
		touchX, touchY := ebiten.TouchPosition(id)
		fmt.Println("TOUCH ID: ", id, " touchX:", touchX, " touchY:", touchY)
	}

	return g.activeScene.Update()
}

// Draw is implemented to draw for Ebiten interface
func (g *Game) Draw(screen *ebiten.Image) {
	g.activeScene.Draw(screen)
}
