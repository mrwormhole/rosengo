package rosengo

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type GameObject struct {
	Img                 *ebiten.Image
	Opts                *ebiten.DrawImageOptions
	Shader              *ebiten.Shader
	ShaderOpts          *ebiten.DrawRectShaderOptions
	IsActive            bool
	ShaderEnabled       bool
	X, Y, Width, Height int
}

func NewGameObject(img *ebiten.Image, x, y int) (*GameObject, error) {
	opts := &ebiten.DrawImageOptions{
		GeoM:          ebiten.GeoM{},
		ColorM:        ebiten.ColorM{},
		CompositeMode: 0,
		Filter:        0,
	}
	opts.GeoM.Translate(float64(x), float64(y))
	w, h := img.Size()

	return &GameObject{
		Img:        img,
		Opts:       opts,
		ShaderOpts: nil,
		IsActive:   true,
		X:          x,
		Y:          y,
		Width:      w,
		Height:     h,
	}, nil
}

func (g *GameObject) SetShader(shader *ebiten.Shader, images [4]*ebiten.Image) {
	g.Shader = shader

	g.ShaderOpts = &ebiten.DrawRectShaderOptions{
		Images:        images,
		Uniforms:      map[string]interface{}{},
		GeoM:          g.Opts.GeoM,
		CompositeMode: g.Opts.CompositeMode,
	}
	g.ShaderEnabled = true
}

func (g *GameObject) SetShaderUniforms(uniforms map[string]interface{}) {
	g.ShaderOpts.Uniforms = uniforms
}

func (g *GameObject) Draw(screen *ebiten.Image) {
	if g.ShaderEnabled {
		screen.DrawRectShader(g.Width, g.Height, g.Shader, g.ShaderOpts)
	} else {
		screen.DrawImage(g.Img, g.Opts)
	}
}
