package rosengo

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
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
		if obj.ShaderEnabled {
			cx, cy := ebiten.CursorPosition()
			w, h := screen.Size()
			log.Println(cx, cy, w, h)
			obj.SetShaderUniforms(map[string]interface{}{
				"Time": float32(s.ticks) / 60,
				"ScreenSize": []float32{
					float32(w),
					float32(h),
				},
				"Cursor": []float32{
					float32(cx),
					float32(cy),
				},
			})
		}
		if obj.IsActive {
			obj.Draw(screen)
		}
	}
}
