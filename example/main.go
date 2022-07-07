package main

import (
	"github.com/bit101/bitebiten/example/states"
	"github.com/bit101/bitebiten/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	game := game.NewGame(800, 600, "my game", states.NewGameState())
	ebiten.RunGame(game)
}
