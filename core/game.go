package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"

)

//var running = true // permet de savoir si la fenetre est ouverte ou fermer
//var groundTexture rl.Texture2D


func (p *gameEngine) Init(width int32, height int32, title string, isRunning bool, dead bool, score int, ) { // initialise les propriété de la fenetre
	p.width = width
	p.heigh = height
	p.title = title
	p.isRunning = isRunning
	p.dead = dead
	p.score = score
}

func (p *gameEngine) initGame() { // Initialise le jeu, en créant la fenêtre ,
	rl.InitWindow(p.width, p.heigh, p.title) // définit la taille de la fenetre
	p.isRunning = true
	defer rl.CloseWindow()
	rl.SetExitKey(0)    // définit les boutons pour être ouvert fermé ?
	rl.SetTargetFPS(60) // définit les fps a x
	p.load()
	for p.isRunning {
		input()
		update(p)
		p.render()
	}
	quit()
}

func (p *gameEngine) load(){ // méthode pour load tous les assets 
	p.txSprites = rl.LoadTexture("assets/Allsprites.png") // va chercher les textures s/o le sprite sheet
}


func input() {


}

func update(w *gameEngine) { // va définir les mouvements du personnage
	w.isRunning = !rl.WindowShouldClose()
	
}

func (g *gameEngine) render() { // permet le rendu de la fenetre c'est à dire les dessins
rl.BeginDrawing()
rl.ClearBackground(rl.Black)
rl.LoadTexture("assets/Allsprites.png")
 // faire une condition pour dire tant que le joueur n'est pas mort : 
	//rl.DrawTexture(g.TxSprites, g.frameRec)
	rl.EndDrawing()
	
	rl.DrawRectangle(650,100,100,38,rl.Blue)

	rl.DrawRectangle(650,235,100,38,rl.Yellow)

	rl.DrawRectangle(650,500,100,38,rl.Red)

//le debut de la mort des bouttons

	if rl.IsMouseButtonPressed(0){
		rl.DrawRectangle(650,100,150,80,rl.Purple)
	}
	if rl.IsMouseButtonPressed(0){
		rl.DrawRectangle(650,235,150,80,rl.Purple)
	}
	if rl.IsMouseButtonPressed(0){
		rl.DrawRectangle(650,500,150,80,rl.Purple)
	}
}

func DrawTexturePro(texture rl.Texture2D, sourceRec, destRec rl.Rectangle, origin rl.Vector2, rotation float32, tint rl.Color) { // permet de faire un carrée pour afficher un sprite : 

}
func quit() {
	rl.CloseWindow()
}

