package rosengo

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type GameObject struct {
	Img                 *ebiten.Image
	Opts                *ebiten.DrawImageOptions
	IsActive            bool
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
		Img:      img,
		Opts:     opts,
		IsActive: true,
		X:        x,
		Y:        y,
		Width:    w,
		Height:   h,
	}, nil
}
