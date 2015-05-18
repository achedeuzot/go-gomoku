package gui

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type Scene interface {
	PlayScene()
}

type MenuMain struct {
	background uint8
}

var CurrScene *Scene
var SceneMap map[string]*Scene = make(map[string]*Scene)

func initScenes() {
	var s Scene

	menu_main := &MenuMain{
		background: 0,
	}
	s = menu_main
	SceneMap["menu_main"] = &s
	CurrScene = SceneMap["menu_main"]
}

func (s *MenuMain) PlayScene() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			fmt.Printf("[%d ms] Quit Event\ttype:%d\n",
				t.Timestamp, t.Type)
			Running = false
		}
	}
}
