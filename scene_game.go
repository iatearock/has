package main

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/jakecoffman/cp"
	"iatearock.com/has/ui"
)

type SceneGame struct {
	sceneManager *SceneManager
	direction    *ui.Direction
	anchor       *ui.Anchor
	inputHandler *InputHandler
	eventQueue   *ControlEvent

	timeStart  time.Time
	timeSince  time.Duration
	waveShader *ebiten.Shader

	startMessage *ui.Message
	endMessage   *ui.Message

	wave    *audio.Player // audio
	seagull *audio.Player

	sceneEnd bool
}

func NewSceneGame(sm *SceneManager) *SceneGame {
	s := &SceneGame{sceneManager: sm}
	dirImg, _ := vfs.GetImage("assets/images/direction.png")
	endPos := game.island[len(game.island)-1].body.Position()
	s.direction = ui.NewDirection(dirImg, 600, 40, endPos)
	anchorImg, _ := vfs.GetImage("assets/images/anchor.png")
	s.anchor = ui.NewAnchor(anchorImg, 570, 40)

	// audio
	s.wave, _ = game.audioManager.NewInfiniteLoop("wave", "assets/audio/632517__thedutchmancreative__waves.ogg")
	s.seagull, _ = game.audioManager.NewPlayer("seagull", "assets/audio/510917__lydmakeren__seagulls-short.ogg")
	s.wave.Play()
	s.seagull.SetVolume(0.2)
	s.seagull.Play()

	// user inpu
	s.eventQueue = NewDefualtControlEvent(game.b)
	s.inputHandler = NewDefaultInputHandler(s.eventQueue.commands)

	s.timeStart = time.Now()
	shaderFile, _ := vfs.ReadFile("assets/shaders/wave.kage")
	s.waveShader, _ = ebiten.NewShader(shaderFile)

	s.startMessage = ui.NewMessage("The commodore is after you.\nFollow the compass to the safe house.",
		cp.Vector{X: -250, Y: 30})
	s.endMessage = ui.NewMessage("Welcom home! Mission Completed.\n Press q to quit game.",
		cp.Vector{X: 450, Y: 200})
	s.endMessage.SetDisplay(false)

	s.sceneEnd = false
	return s
}

func (s *SceneGame) Update() error {
	if !s.sceneEnd {
		s.inputHandler.Update()
		game.b.Update()
		space.Step(1.0 / 60.0)

		game.wake.Update()

		// audio
		if !s.seagull.IsPlaying() {
			// if seagull not playing, then randomly start playing
			if rand.Float64() < 0.0011111 {
				// approx play once in 15 seconds, assuming 60fps
				s.seagull.Rewind()
				s.seagull.Play()
			}
		}

		s.timeSince = time.Since(s.timeStart)
		s.sceneEnd = s.IsFinished()
	}

	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		s.sceneManager.GoTo(NewSceneTitle(s.sceneManager))
		game.audioManager.CloseAll()
	}

	//

	// update camera and HUD
	bp := game.b.body.Position()
	s.direction.SetPosition(bp) // HUD arrow pointer
	s.anchor.SetDisplay(game.b.isAnchor)
	game.cam.Position = [2]float64{bp.X, bp.Y}
	return nil
}

func (s *SceneGame) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{128, 170, 187, 255})
	waveShader(screen, s.waveShader,
		float32(s.timeSince.Milliseconds()),
		float32(game.cam.Position[0]), float32(game.cam.Position[1]))

	game.cam.Update()
	for _, is := range game.island {
		is.Draw(screen)
	}
	game.wake.Draw(screen)
	game.b.Draw(screen)

	// HUD
	s.direction.Draw(screen)
	s.anchor.Draw(screen)

	s.startMessage.Draw(screen, font24, game.cam.matrix)
	if s.sceneEnd {
		s.endMessage.Draw(screen, font24, game.cam.matrix)
	}

	// text.Draw(screen, "Press q to return to Title Screen", font24, 50, screenHeight/8*7, color.White)
	ebitenutil.DebugPrintAt(screen, "Press q to return to Title Screen", 10, 40)
	ebitenutil.DebugPrintAt(screen, "Press d to anchor/dock", 10, 55)

	s.DrawInfo(screen)
}

func (s *SceneGame) DrawInfo(screen *ebiten.Image) {
	if game.b.inStore {
		text.Draw(screen, "Press D to dock/undock", font24, 300, 40, color.White)
	}
}

func (s *SceneGame) IsFinished() bool {
	// boat at finish island
	if island, ok := game.boatOnIsland[game.b]; ok {
		if island.end {
			// boat at finishing island
			if game.b.inStore && game.b.isAnchor {
				return true
			}
		}
	}
	return false
}

func waveShader(image *ebiten.Image, s *ebiten.Shader, t float32, cx, cy float32) {
	sop := &ebiten.DrawRectShaderOptions{}
	sop.Uniforms = map[string]interface{}{
		"Pi":     float32(3.14159265359),
		"Time":   t, // milliseconds,
		"Camera": []float32{cx, cy},
		// "Cursor": []float32{float32(mx), float32(my)},
	}
	image.DrawRectShader(screenWidth, screenHeight, s, sop)
}
