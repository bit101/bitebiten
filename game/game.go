package game

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type State interface {
	Update() error
	Draw(screen *ebiten.Image)
	Init(game *Game)
}

type Game struct {
	width, height int
	currentState  State
	isPointer     bool
}

func NewGame(w, h int, title string, firstState State) *Game {
	ebiten.SetWindowSize(w, h)
	ebiten.SetWindowTitle(title)

	game := &Game{width: w, height: h}
	game.isPointer = false
	game.SetState(firstState)
	err := ebiten.RunGame(game)
	if err != nil {
		log.Fatal(err)
	}
	return game
}

func (g *Game) Update() error {
	return g.currentState.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.currentState.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.width, g.height
}

func (g *Game) SetState(newState State) {
	g.currentState = newState
	newState.Init(g)
}

func (g *Game) SetPointer(b bool) {
	g.isPointer = b
}

func (g *Game) UpdateCursor() {
	if g.isPointer {
		ebiten.SetCursorShape(ebiten.CursorShapePointer)
	} else {
		ebiten.SetCursorShape(ebiten.CursorShapeCrosshair)
	}
}
