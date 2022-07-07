package main

import (
	"os"

	"github.com/bit101/bitebiten/game"
	"github.com/bit101/bitebiten/world"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type GameState struct {
	x int
	y int
}

func NewGameState() *GameState {
	return &GameState{}
}

func (s *GameState) Init(gameRef *game.Game) {
	s.y = world.GetHeightInt() / 2
}

func (s *GameState) Update() error {
	s.x++
	if s.x > world.GetWidthInt() {
		s.x = 0
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		os.Exit(0)
	}
	return nil
}

func (s *GameState) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "Hello world", s.x, s.y)
}

func main() {
	firstState := NewGameState()
	game := game.NewGame(800, 600, "my game", firstState)
	ebiten.RunGame(game)
}
