package rosengo

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	Name() string
	GameObjects() []*GameObject
	GameState() GameState
	Update() error
	Draw(screen *ebiten.Image)
}

type scene struct {
	name    string
	ticks   int
	objects []*GameObject
	state   GameState
}

func NewScene(name string, objects []*GameObject, state GameState) Scene {
	return &scene{
		name:    name,
		objects: objects,
		state:   state,
	}
}

func (s *scene) Name() string {
	return s.name
}

func (s *scene) GameObjects() []*GameObject {
	return s.objects
}

func (s *scene) GameState() GameState {
	return s.state
}

func (s *scene) Update() error {
	s.ticks++
	return nil
}

func (s *scene) Draw(screen *ebiten.Image) {
	for _, obj := range s.objects {
		if obj.IsActive {
			obj.Draw(screen)
		}
	}
}
