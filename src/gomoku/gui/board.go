package gui

import (
	"gomoku/arena"
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

type Board struct {
	Background *Texture
	Deck       *Texture
}

func NewBoard() *Board {
	board := &Board{
		Background: GetTextureFromImage("data/img/bg.jpg"),
		Deck:       GetTextureFromImage("data/img/board.png"),
	}
	// Display background to the right scale
	var ratio float64
	var finalW int32
	var finalH int32

	if board.Deck.size.W > DisplayMode.W {
		ratio = float64(DisplayMode.W) / float64(board.Deck.size.W)
		finalW = int32(float64(board.Deck.size.W) * ratio)
		finalH = int32(float64(board.Deck.size.H) * ratio)
	}

	if board.Deck.size.H > DisplayMode.H {
		ratio = float64(DisplayMode.H) / float64(board.Deck.size.H)
		finalW = int32(float64(board.Deck.size.W) * ratio)
		finalH = int32(float64(board.Deck.size.H) * ratio)
	}

	board.Background.pos = sdl.Rect{X: 0, Y: 0, W: DisplayMode.W, H: DisplayMode.H}
	board.Deck.pos = sdl.Rect{X: DisplayMode.W/2 - finalW/2, Y: 0, W: finalW, H: finalH}
	return board
}

func (b *Board) PlayScene() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			Running = false
		case *sdl.KeyUpEvent:
			if t.Keysym.Sym == sdl.K_ESCAPE {
				CurrScene = SceneMap["MenuMain"]
			}
		}
	}
	Renderer.Clear()

	Renderer.Copy(b.Background.texture, &b.Background.size, &b.Background.pos)
	Renderer.Copy(b.Deck.texture, &b.Deck.size, &b.Deck.pos)

	// Display content of board in top of background
	demoArena := arena.Gomoku.Arena
	for idx, val := range demoArena.Goban {
		log.Printf("Goban data [%d]: %d\n", idx, val)
	}
	Renderer.Present()

}
