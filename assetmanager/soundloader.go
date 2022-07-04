package assetmanager

import (
	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

var sounds map[string]*audio.Player
var audioContext *audio.Context
var soundOn bool

func LoadSound(name string, bs []byte) {
	if sounds == nil {
		sounds = make(map[string]*audio.Player)
	}
	if audioContext == nil {
		sampleRate := 44100
		audioContext = audio.NewContext(sampleRate)
	}

	d, err := wav.Decode(audioContext, bytes.NewReader(bs))
	if err != nil {
		log.Fatal("Could not create audio stream")
	}
	player, err := audioContext.NewPlayer(d)
	if err != nil {
		log.Fatal("Could not create audio stream")
	}

	sounds[name] = player
}

func PlaySound(name string) {
	if !soundOn {
		return
	}
	snd, ok := sounds[name]
	if !ok {
		log.Fatalf("Could not find sound: %s", name)
	}
	snd.Rewind()
	snd.Play()
}

func SetSoundOn(b bool) {
	soundOn = b
}

func IsSoundOn() bool {
	return soundOn
}

func ToggleSound() {
	SetSoundOn(!soundOn)
}
