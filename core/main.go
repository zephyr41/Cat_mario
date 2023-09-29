package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type gameEngine struct {
	width      int32
	height     int32
	title      string
	isRunning  bool
	dead       bool
	score      int
	texture    rl.Texture2D
	musicMenu  rl.Music
	// Add other properties as needed
}

func (p *gameEngine) Init(width, height int32, title string, isRunning, dead bool, score int) {
	p.width = width
	p.height = height
	p.title = title
	p.isRunning = isRunning
	p.dead = dead
	p.score = score
}

func (p *gameEngine) initGame() {
	rl.InitWindow(p.width, p.height, p.title)
	rl.InitAudioDevice()
	rl.SetMasterVolume(1)

	p.isRunning = true
	p.texture = rl.LoadTexture("../assets/assets.png")

	// Load and play the music
	p.musicMenu = rl.LoadMusicStream("../audio/music/menu.mp3")
	rl.PlayMusicStream(p.musicMenu)

	rl.SetExitKey(0)
	rl.SetTargetFPS(60)
}

func (p *gameEngine) update() {
	rl.UpdateMusicStream(p.musicMenu)
}

func (p *gameEngine) render() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)
	rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)
	rl.EndDrawing()
}

func (p *gameEngine) quit() {
	rl.UnloadTexture(p.texture)
	rl.UnloadMusicStream(p.musicMenu)
	rl.CloseAudioDevice()
	rl.CloseWindow()
}

func main() {
	var engine gameEngine

	engine.Init(800, 600, "My Game", false, false, 0)
	engine.initGame()

	for engine.isRunning {
		engine.update()
		engine.render()
	}

	engine.quit()
}
