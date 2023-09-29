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
	// mario src

}

func (p *gameEngine) initGame() { // Initialise le jeu, en créant la fenêtre ,

	rl.InitWindow(p.width, p.heigh, p.title) // définit la taille de la fenetre
	p.isRunning = true
	defer rl.CloseWindow()
	rl.SetExitKey(0)    // définit les boutons pour être ouvert fermé ?
	rl.SetTargetFPS(60) // définit les fps a x
	p.playerSprite = rl.LoadTexture("../assets/allsprites.png")
	for p.isRunning {
		p.input()
		p.update()
		p.render()
	}
	p.quit()
}

func (w *gameEngine) input() { // récupère les inputs de la map

}

func (w *gameEngine) update() { // va définir les mouvements du personnage
	w.isRunning = !rl.WindowShouldClose()

}


const (
    MenuDisplay = iota
    Game
    Options
)
var currentGameState = MenuDisplay

func Display(){
		// rl.InitWindow(800, 600, "Mon Jeu")
		// defer rl.CloseWindow()
	
		// rl.SetTargetFPS(60)
	
		for !rl.WindowShouldClose() {
			switch currentGameState {
			case MenuDisplay:
				if rl.IsKeyPressed(rl.KeyEnter) {
					currentGameState = Game
				} else if rl.IsKeyPressed(rl.KeyO) {
					currentGameState = Options
				} else if rl.IsKeyPressed(rl.KeyEscape) {
					rl.CloseWindow()
				}
			case Game:
				if rl.IsKeyPressed(rl.KeyEscape) {
					currentGameState = MenuDisplay
				}
			case Options:
				if rl.IsKeyPressed(rl.KeyEscape) {
					currentGameState = MenuDisplay
				}
			}
	
			rl.BeginDrawing()
			rl.ClearBackground(rl.RayWhite)
	
			switch currentGameState {
			case MenuDisplay:
				// Menue
				rl.DrawText("PLAY - Appuyez sur ENTER pour jouer", 10, 10, 20, rl.DarkGray)
				rl.DrawText("OPTIONS - Appuyez sur O pour accéder aux options", 10, 30, 20, rl.DarkGray)
				rl.DrawText("QUIT - Appuyez sur ESCAPE pour quitter", 10, 50, 20, rl.DarkGray)
			case Game:
				// JEUX
				rl.DrawText("JEU EN COURS - Appuyez sur ESCAPE pour revenir au menu", 10, 10, 20, rl.White)
				rl.DrawRectangle(1,1,10000,10000,rl.Black)

			case Options:
				rl.DrawRectangle(1,1,10000,10000,rl.Blue)
				// OPTION
				
				//QUiTTEZ
				rl.DrawText("OPTIONS - Appuyez sur ESCAPE pour revenir au menu", 10, 10, 20, rl.White)
			}
	
			rl.EndDrawing()
		}
	}
//_________________________________________________________________Display___________________________________________________________//


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

}

func (p *gameEngine) quit() {
	rl.UnloadTexture(p.playerSprite)
	rl.CloseWindow()
}

func drawScene() {
}
