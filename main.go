package main

import (
	"flag"
	"fmt"
	"github.com/MrWormHole/rosengo/rosengo"
	"github.com/hajimehoshi/ebiten/v2"
	"os"
	"runtime/pprof"
	"runtime/trace"
)

var (
	memProfile  = flag.String("memprofile", "", "write memory profile to file")
	traceOut    = flag.String("trace", "", "write trace to file")
	transparent = flag.Bool("transparent", false, "background transparency")
	mute        = flag.Bool("mute", false, "mute")
)

func main() {
	flag.Parse()

	if *traceOut != "" {
		f, err := os.Create(*traceOut)
		if err != nil {
			panic(err)
		}
		defer func() {
			if err := f.Close(); err != nil {
				panic(err)
			}
		}()

		err = trace.Start(f)
		if err != nil {
			panic(fmt.Sprintf("could not write trace: %s", err))
		}
		defer trace.Stop()
	}

	game, err := rosengo.NewGame()
	if err != nil {
		panic(err)
	}

	if *transparent {
		ebiten.SetScreenTransparent(true)
		ebiten.SetWindowDecorated(false)
		game.SetTransparent(true)
	}

	if *mute {
		game.SetMute(true)
	}

	const scale = 2
	ebiten.SetWindowSize(rosengo.ScreenWidth*scale, rosengo.ScreenHeight*scale)
	ebiten.SetWindowResizable(true)
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}

	if *memProfile != "" {
		f, err := os.Create(*memProfile)
		if err != nil {
			panic(err)
		}
		defer func() {
			if err := f.Close(); err != nil {
				panic(err)
			}
		}()

		if err := pprof.WriteHeapProfile(f); err != nil {
			panic(fmt.Sprintf("could not write memory profile: %s", err))
		}
	}
}
