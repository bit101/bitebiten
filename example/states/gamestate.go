package states

import (
	_ "embed"
	_ "image/png"
	"os"

	"github.com/bit101/bitebiten/assetmanager"
	"github.com/bit101/bitebiten/component"
	"github.com/bit101/bitebiten/game"
	"github.com/bit101/bitebiten/world"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

//go:embed ship.png
var shipImg []byte

//go:embed out.png
var anim []byte

type GameState struct {
	x         int
	y         int
	ship      *component.ImageView
	anim      *component.AnimView
	explosion *component.Explosion
	game      *game.Game
}

func NewGameState() *GameState {
	return &GameState{}
}

func (s *GameState) Init(gameRef *game.Game) {
	s.game = gameRef
	// init text position
	s.y = 20

	// init ship
	assetmanager.LoadImage("ship", shipImg)
	s.ship = component.NewImageView("ship")
	s.ship.SetPos(world.GetWidth()/2, world.GetHeight()/2)
	s.ship.SetRegistration(16, 16)

	// init anim
	assetmanager.LoadImage("anim", anim)
	s.anim = component.NewAnimView("anim", 40, 40, 25, 5)
	s.anim.SetPos(100, 100)
	s.anim.SetRegistration(20, 20)

	// init explosion
	s.explosion = component.NewExplosion(100)
}

func (s *GameState) Update() error {
	s.game.SetPointer(false)
	// update text
	s.x++
	if s.x > world.GetWidthInt() {
		s.x = 0
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		os.Exit(0)
	}

	// is pointer over ship? set cursor to pointer
	cx, cy := ebiten.CursorPosition()
	if cx > world.GetWidthInt()/2-15 && cx < world.GetWidthInt()/2+15 &&
		cy > world.GetHeightInt()/2-15 && cy < world.GetHeightInt()/2+15 {
		s.game.SetPointer(true)
	}

	// update ship
	s.ship.Rotation += 0.05

	// update anim
	s.anim.Update()
	s.anim.Y++
	if s.anim.Y > world.GetHeight() {
		s.anim.Y = 0
	}

	// update explosion
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		s.explosion.SetPos(float64(x), float64(y))
		s.explosion.Start()
	}
	s.explosion.Update()

	s.game.UpdateCursor()

	return nil
}

func (s *GameState) Draw(screen *ebiten.Image) {
	// draw text
	ebitenutil.DebugPrintAt(screen, "Hello world", s.x, s.y)

	// draw ship
	s.ship.Draw(screen)

	// draw anim
	s.anim.Draw(screen)

	// draw explosion
	s.explosion.Draw(screen)
}
