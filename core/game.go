package main

import (

	rl "github.com/gen2brain/raylib-go/raylib"
)

//var running = true // permet de savoir si la fenetre est ouverte ou fermer

func (p *gameEngine) Init(width int32, height int32, title string, isRunning bool, dead bool, score int) { // initialise les propriété de la fenetre
	// Pour selectionner notre joueur (faire une structure plus tard ? ) :
	// va déssiner un rectangle dans la feuille de sprite pour pouvoir dessiner mario : il aura : x & y la position de ou il part
	// width et heigt seront la largeur de la fenetre
	// src est la source (la feuille) donc ça position
	// Dest est la ou on l'envoie (notre fênetre de jeu)
	p.width = width
	p.heigh = height
	p.title = title
	p.isRunning = isRunning
	p.dead = dead
	p.score = score
	// mario src

}

func (p *gameEngine) initGame() { // Initialise le jeu, en créant la fenêtre ,
	rl.InitWindow(p.width, p.heigh, p.title)
	// définit la taille de la fenetre
	p.isRunning = true
	p.texture = rl.LoadTexture("../assets/assets.png")
	p.playerSrc = rl.NewRectangle(2, 192, 25, 27) // selectionne un bout d'image dans la sheet sprite
	p.playerDest = rl.NewRectangle(0, 32, 64, 64) // met une zone ou afficher ce bout d'image
	p.playerVector = rl.NewVector2(-(p.playerDest.Width), -(p.playerDest.Height)) // permet de lui donner une position
	p.playerSpeed = 0.05

	p.display()
	rl.SetExitKey(0)    // définit les boutons pour être ouvert fermé ?
	rl.SetTargetFPS(60) // définit les fps a x

}
// _________________________________________________________________Menu_______________________________________________________________//
// iota est un identificateur prédéclaré représentant le numéro ordinal entier non typé de la spécification
// const actuelle dans une déclaration const (généralement entre parenthèses). Il est indexé à zéro.
const (
	MenuDisplay = iota
	Game
	Options
)

var currentGameState = MenuDisplay

func (p *gameEngine) display() {

	// defer rl.CloseWindow()

	// rl.SetTargetFPS(60)

	for p.isRunning {
		switch currentGameState {
		case MenuDisplay:
			if rl.IsKeyReleased(rl.KeyEnter) {
				currentGameState = Game
			} else if rl.IsKeyReleased(rl.KeyO) {
				currentGameState = Options
			} else if rl.IsKeyReleased(rl.KeyEscape) {
				p.quit()
			}
		case Game:
			if rl.IsKeyReleased(rl.KeyEscape) {
				currentGameState = MenuDisplay

			}
		case Options:
			if rl.IsKeyReleased(rl.KeyEscape) {
				currentGameState = MenuDisplay
			}
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		switch currentGameState {
		case MenuDisplay:
			// Menue
			rl.DrawText("PLAY - Appuyez sur ENTER pour jouer :", 100, 150, 35, rl.White)

			rl.DrawText("OPTIONS - Appuyez sur O pour accéder aux options :", 100, 300, 35, rl.White)

			rl.DrawText("QUIT - Appuyez sur ESCAPE pour quitter :", 100, 450, 35, rl.White)

			rl.ClearBackground(rl.DarkBlue)

		case Game:
			// JEUX
			rl.ClearBackground(rl.Black)
			rl.DrawText("JEU EN COURS - Appuyez sur ESCAPE pour revenir au menu :", 100, 150, 35, rl.White)
			p.input()
			p.update()
			p.render()

		case Options:
			// OPTION //

			rl.ClearBackground(rl.White)

			rl.DrawText("Setings Glogbal :", 580, 1, 35, rl.White)

			if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(15, 90, 50, 50)) {
				rl.DrawRectangle(15, 90, 50, 50, rl.White)
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				}
			} else {
				rl.DrawRectangle(15, 90, 50, 50, rl.LightGray)
			}
			if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(70, 90, 50, 50)) {
				rl.DrawRectangle(70, 90, 50, 50, rl.Yellow)
				if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
				}
			} else {
				rl.DrawRectangle(70, 90, 50, 50, rl.LightGray)
			}

			rl.DrawText("FPS-TOUCHES", 580, 85, 35, rl.White)

			// QUiTTEZ //
			rl.DrawText("OPTIONS - Appuyez sur ESCAPE pour revenir au menu :", 300, 45, 35, rl.White)
		}
		rl.ClearBackground(rl.DarkBlue)

		rl.EndDrawing()
	}
}


func (w *gameEngine) input() { // récupère les inputs de la map

	if rl.IsKeyDown(rl.KeyUp) { // key left
		w.playerDest.Y -= w.playerSpeed
	} 
	if rl.IsKeyDown(rl.KeyDown) { // key left
		w.playerDest.Y += w.playerSpeed
	} 
	if rl.IsKeyDown(rl.KeyLeft) { // key left
		w.playerDest.X -= w.playerSpeed
	} 
	if  rl.IsKeyDown(rl.KeyRight) { // key left
		w.playerDest.X += w.playerSpeed
	}
}

func (w *gameEngine) update() { // va définir les mouvements du personnage
	w.isRunning = !rl.WindowShouldClose()

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

	
	rl.DrawTexturePro(g.texture, g.playerSrc, g.playerDest, g.playerVector, 0, rl.White) // drawTextureMario
	//rl.DrawTexture(g.texture,0,0,rl.White)

}

func (p *gameEngine) quit() {
	rl.UnloadTexture(p.texture)
	rl.CloseWindow()
}
