package main

import (
	"bytes"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
)

type AudioManager struct {
	ctx            *audio.Context
	players        map[string]*audio.Player
	bytesPerSample int64
}

func NewAudioManager() *AudioManager {
	return &AudioManager{ctx: audio.NewContext(44100),
		players:        map[string]*audio.Player{},
		bytesPerSample: 4,
	}
}

func (m *AudioManager) NewPlayer(name, path string) (*audio.Player, error) {

	byteSlice, err := vfs.ReadFile(path)
	if err != nil {
		log.Println("Error reading audio")
		return nil, err
	}
	s, err := vorbis.DecodeWithSampleRate(m.ctx.SampleRate(), bytes.NewReader(byteSlice))
	if err != nil {
		log.Println("Error decoding audio")
		return nil, err
	}
	p, err := m.ctx.NewPlayer(s)
	if err != nil {
		return nil, err
	}
	m.players[name] = p
	return p, nil
}

func (m *AudioManager) NewInfiniteLoop(name, path string) (*audio.Player, error) {
	byteSlice, err := vfs.ReadFile(path)
	if err != nil {
		log.Println("Error reading audio")
		return nil, err
	}
	s, err := vorbis.DecodeWithSampleRate(m.ctx.SampleRate(), bytes.NewReader(byteSlice))
	if err != nil {
		log.Println("Error decoding audio")
		return nil, err
	}
	w := time.Second * time.Duration(s.Length()) / time.Duration(m.ctx.SampleRate()) / 4
	w = w - time.Duration(time.Millisecond*500)
	stream := audio.NewInfiniteLoop(s, int64(w))
	p, err := m.ctx.NewPlayer(stream) // TODO try newplayerfrombytes
	if err != nil {
		return nil, err
	}
	m.players[name] = p
	return p, nil

}

func (m *AudioManager) CloseAll() {
	for _, p := range m.players {
		p.Close()
	}
}
