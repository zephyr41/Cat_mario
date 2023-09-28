package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var running = true // permet de savoir si la fenetre est ouverte ou fermer
func input()       {}
func update() {
	running = !rl.WindowShouldClose()
}

func (p *window) Init(width int32, height int32, title string) { // initialise les propriété de la fenetre 
	p.width = width
	p.heigh = height
	p.title = title
}

func (p *window) initGame() { // Initialise le jeu, en créant la fenêtre , 
	rl.InitWindow(p.width, p.heigh, p.title) // définit la taille de la fenetre
	defer rl.CloseWindow()                   // définit les boutons pour être ouvert fermé ?
	rl.SetTargetFPS(60)                      // définit les fps a x
	for running {
		input()
		update()
		render()
	}
	quit()
}

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)
	drawScene()
	rl.EndDrawing()

	rl.DrawRectangle(650,100,100,38,rl.Blue)

	rl.DrawRectangle(650,235,100,38,rl.Yellow)

	rl.DrawRectangle(650,500,100,38,rl.Red)
	if rl.IsMouseButtonPressed(1){
		rl.DrawRectangle(650,500,100,38,rl.Purple)
	}
}

func quit() {
	rl.CloseWindow()
}

func drawScene() {

}