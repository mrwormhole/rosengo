package rosengo

import "github.com/hajimehoshi/ebiten/v2"

type Scene interface {
	Name() string
	GameObjects() []GameObject
	GameState() GameState
	Update() error
	Draw(screen *ebiten.Image) error
}

type scene struct {
	name    string
	objects []GameObject
	state   GameState
}

func NewScene(name string, objects []GameObject, state GameState) Scene {
	return &scene{
		name:    name,
		objects: objects,
		state:   state,
	}
}

func (s *scene) Name() string {
	return s.name
}

func (s *scene) GameObjects() []GameObject {
	return s.objects
}

func (s *scene) GameState() GameState {
	return s.state
}

func (s *scene) Update() error {
	return nil
}

func (s *scene) Draw(screen *ebiten.Image) error {
	return nil
}
