package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

//var running = true // permet de savoir si la fenetre est ouverte ou fermer

func (p *gameEngine) Init(width int32, height int32, title string, isRunning bool, dead bool, score int) { // initialise les propriété de la fenetre
	p.width = width
	p.heigh = height
	p.title = title
	p.isRunning = isRunning
	p.dead = dead
	p.score = score
	// mario dest
	// mario src

}

func (p *gameEngine) initGame() { // Initialise le jeu, en créant la fenêtre ,

	rl.InitWindow(p.width, p.heigh, p.title) // définit la taille de la fenetre
	p.isRunning = true
	defer rl.CloseWindow()
	rl.SetExitKey(0)    // définit les boutons pour être ouvert fermé ?
	rl.SetTargetFPS(60) // définit les fps a x
	p.texture = rl.LoadTexture("../assets/allsprites.png")
	for p.isRunning {
		p.input()
		p.update()
		p.render()
	}
	quit()
}

func (w *gameEngine) input() { // récupère les inputs de la map

}

func (w *gameEngine) update() { // va définir les mouvements du personnage
	w.isRunning = !rl.WindowShouldClose()

}

func (g *gameEngine) render() { // permet le rendu de la fenetre c'est à dire les dessins
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	// faire une condition pour dire tant que le joueur n'est pas mort :
	//rl.DrawTexture(g.TxSprites, g.frameRec)
	// sourceTest := rl.Rectangle{}
	// destRecTest := rl.Rectangle{}
	// originTest := rl.Vector2{}

	rl.DrawTexturePro(g.texture, rl.NewRectangle(0, 0, 600,800), rl.NewRectangle(0, 0,600, 800), rl.NewVector2(0,0), 0, rl.White)
	//rl.DrawTexture(g.texture,0,0,rl.White)
	rl.EndDrawing()

	// rl.DrawRectangle(650, 100, 100, 38, rl.Blue)

	// rl.DrawRectangle(650, 235, 100, 38, rl.Yellow)
	// rl.DrawRectangle(650, 500, 100, 38, rl.Red)

	// if rl.IsMouseButtonPressed(1) {
	// 	rl.DrawRectangle(650, 500, 100, 38, rl.Purple)
	// }

	//rl.DrawTexturePro(g.TxClouds, rl.NewRectangle(-g.CloudRec.X, 0, float32(g.TxClouds.Width), float32(g.TxClouds.Height)),
	//DrawTexturePro(texture, sourceRec, destRec, origin)
}

// func DrawTexturePro(texture rl.Texture2D, sourceRec, destRec rl.Rectangle, origin rl.Vector2) { // permet de faire un carrée pour afficher un sprite :
// 	texture = rl.LoadTexture("../assets/Allsprites.png")
// 	sourceRec = rl.NewRectangle(300,300, 100, 100)
// 	destRec = rl.Rectangle{300, 300, 100, 100}
// 	origin = rl.Vector2{0, 0}
// }

func quit() {
	rl.CloseWindow()
}

func drawScene() {
}
