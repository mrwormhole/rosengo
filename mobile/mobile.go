package mobile

import (
	"github.com/MrWormHole/rosengo/rosengo"
	"github.com/hajimehoshi/ebiten/v2/mobile"
)

func init() {
	rosengo, err := rosengo.NewGame()
	if err != nil {
		panic(err)
	}
	mobile.SetGame(rosengo)
}

// Dummy is a dummy exported function.
//
// gomobile doesn't compile a package that doesn't include any exported function.
// Dummy forces gomobile to compile this package.
func Dummy() {}
