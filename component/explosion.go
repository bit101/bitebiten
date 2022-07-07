package component

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Particle struct {
	x, y, vx, vy float64
	maxAge       int
}

func NewParticle() *Particle {
	return &Particle{}
}

func (p *Particle) Start(x, y, maxSpeed float64, maxAge int) {
	angle := rand.Float64() * math.Pi * 2
	speed := rand.Float64() * maxSpeed
	p.x = x
	p.y = y
	p.vx = math.Cos(angle) * speed
	p.vy = math.Sin(angle) * speed
	p.maxAge = rand.Int() % maxAge
}

func (p *Particle) Update() {
	p.x += p.vx
	p.y += p.vy
}

func (p *Particle) Draw(screen *ebiten.Image) {
}

type Explosion struct {
	Actor
	particles          []*Particle
	count, age, maxAge int
	image              *ebiten.Image
	active             bool
	maxSpeed           float64
	color              color.RGBA
	size               float64
}

func NewExplosion(count int) *Explosion {
	e := &Explosion{
		maxAge:   60,
		maxSpeed: 3.0,
		count:    count,
		color:    color.RGBA{255, 255, 255, 255},
		size:     2,
	}
	for i := 0; i < e.count; i++ {
		e.particles = append(e.particles, NewParticle())
	}

	e.image = ebiten.NewImage(1, 1)
	e.image.Fill(e.color)
	return e
}

func (e *Explosion) SetColor(r, g, b uint8) {
	e.color.R = r
	e.color.G = g
	e.color.B = b
	e.image.Fill(e.color)
}

func (e *Explosion) SetSize(size float64) {
	e.size = size
}

func (e *Explosion) SetMaxSpeed(maxSpeed float64) {
	e.maxSpeed = maxSpeed
}

func (e *Explosion) SetMaxAge(maxAge int) {
	e.maxAge = maxAge
}

func (e *Explosion) SetCount(count int) {
	e.count = count
	e.particles = []*Particle{}
	for i := 0; i < e.count; i++ {
		e.particles = append(e.particles, NewParticle())
	}
}

func (e *Explosion) Start() {
	for _, p := range e.particles {
		p.Start(e.X, e.Y, e.maxSpeed, e.maxAge)
	}
	e.active = true
	e.age = 0
}

func (e *Explosion) Update() {
	if e.active {
		for _, p := range e.particles {
			p.Update()
		}
		e.age++
		if e.age >= e.maxAge {
			e.active = false
		}
	}
}

func (e *Explosion) Draw(screen *ebiten.Image) {
	if e.active {
		for _, p := range e.particles {
			alpha := 1.0 - float64(e.age)/float64(p.maxAge)
			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Scale(e.size, e.size)
			opts.GeoM.Translate(p.x-e.size/2, p.y-e.size/2)
			opts.ColorM.Scale(1, 1, 1, alpha)
			screen.DrawImage(e.image, opts)
		}
	}
}
