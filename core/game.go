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
	p.display()
	 // définit la taille de la fenetre
	p.isRunning = true
	defer rl.CloseWindow()
	rl.SetExitKey(0)    // définit les boutons pour être ouvert fermé ?
	rl.SetTargetFPS(60) // définit les fps a x
	p.texture = rl.LoadTexture("../assets/allspritess.png")
	// rl.DrawTexturePro(g.texture, rl.NewRectangle(0, 200, 200,70), rl.NewRectangle(0, 0,200, 70), rl.NewVector2(0,0), 0, rl.White) // drawTextureMario
	// Pour selectionner notre joueur (faire une structure plus tard ) : 
	// va déssiner un rectangle dans la feuille de sprite pour pouvoir dessiner mario : il aura : x & y la position de ou il part
	// width et heigt seront la largeur de la fenetre
	// src est la source (la feuille) donc ça position
	// Dest est la ou on l'envoie (notre fênetre de jeu)
	p.playerDest = rl.NewRectangle(0,200,200,70)
	p.playerSrc = rl.NewRectangle(0,200,200,70)
	p.playerVector = rl.NewVector2(0,0)

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

//_________________________________________________________________Menu_______________________________________________________________//
// iota est un identificateur prédéclaré représentant le numéro ordinal entier non typé de la spécification 
// const actuelle dans une déclaration const (généralement entre parenthèses). Il est indexé à zéro.
const (
    MenuDisplay = iota
    Game
    Options
)
var currentGameState = MenuDisplay

func (p *gameEngine) display(){
	rl.InitWindow(p.width, p.heigh, p.title)
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
//_________________________________________________________________Menu_______________________________________________________________//


func (g *gameEngine) render() { // permet le rendu de la fenetre c'est à dire les dessins
	// rl.BeginDrawing()
	// rl.ClearBackground(rl.Black)

	// // faire une condition pour dire tant que le joueur n'est pas mort :
	// //rl.DrawTexture(g.TxSprites, g.frameRec)
	// // sourceTest := rl.Rectangle{}
	// // destRecTest := rl.Rectangle{}
	// // originTest := rl.Vector2{}

	// rl.DrawTexturePro(g.texture, rl.NewRectangle(0, 200, 200,70), rl.NewRectangle(0, 0,200, 70), rl.NewVector2(0,0), 0, rl.White) // drawTextureMario
	rl.DrawTexturePro(g.texture, g.playerSrc, g.playerDest, g.playerVector, 0, rl.White) // drawTextureMario
	//rl.DrawTexture(g.texture,0,0,rl.White)
	rl.EndDrawing()

}

func (p *gameEngine) quit() {
	// rl.UnloadTexture(p.texture)
	// rl.CloseWindow()
}

