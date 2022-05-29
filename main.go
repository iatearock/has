package main

import (
	"embed"
	"fmt"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/iatearock/dango"
	"github.com/jakecoffman/cp"
	"golang.org/x/image/font"
)

//go:embed assets/*
var dataFS embed.FS

const (
	screenWidth  = 640
	screenHeight = 480
)

var (
	vfs          *dango.FS
	sceneManager *SceneManager
	font24       font.Face
	game         *Game // global game state
	space        *cp.Space
)

type Game struct {
	cam          *Camera
	audioManager *AudioManager
	b            *Boat
	island       *Island
	start        *Island
	end          *Island
	windVel      cp.Vector // radian
	wake         *Wake
	boatOnIsland map[*Boat]*Island
}

func NewGame() *Game {
	g := &Game{b: NewBoat(20, 200), windVel: cp.Vector{X: 1, Y: 0},
		island: NewIsland(300, 300, 50),
		start:  NewIsland(20, 200, 50), end: NewIsland(600, 200, 50),
		wake:         NewWake(),
		boatOnIsland: map[*Boat]*Island{},
	}
	g.cam = &Camera{}
	g.cam.Reset()
	g.cam.ViewPort = [2]float64{screenWidth, screenHeight}
	g.cam.Position = [2]float64{20, 200}
	// var err error
	// g.audioManager, err = NewAudioPlayer(audio.NewContext(44100)) // sample rate
	g.audioManager = NewAudioManager()
	// if err != nil {
	// 	log.Printf("Cannot start audio player, %s\n", err)
	// }
	return g
}

func (g *Game) Update() error {
	return sceneManager.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	sceneManager.Draw(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS:%0.2f\nTPS:%0.2f",
		ebiten.CurrentFPS(), ebiten.CurrentTPS()))
}

func (g *Game) Layout(outsideW, outsideH int) (screenW, screenH int) {
	return screenWidth, screenHeight
}
func main() {
	log.SetFlags(log.Lshortfile | log.Ltime)
	rand.Seed(0)
	log.Println("HAS loading data...")
	vfs = dango.NewFS(dataFS)
	loadFonts()
	space = cp.NewSpace()
	space.SetDamping(0.8)
	InitSpace(space)

	game = NewGame()

	log.Println("HAS starting...")
	sceneManager = NewSceneManager(screenWidth, screenHeight)
	sceneManager.GoTo(&SceneTitle{sceneManager})

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("HAS")
	ebiten.SetWindowResizingMode(2)

	if err := ebiten.RunGame(game); err != nil {
		log.Println(err)
	}
	log.Println("HAS shutdonw")
}

func loadFonts() {
	var err error
	font24, err = vfs.GetFontFace(
		"assets/font/Roboto/Roboto-Regular.ttf", 24, 72)
	if err != nil {
		log.Println(err)
	}
}
