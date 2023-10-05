package main

import (
	"encoding/xml"
	"fmt"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

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

}

func (p *gameEngine) initGame() { // Initialise le jeu, en créant la fenêtre ,
	rl.InitWindow(p.width, p.heigh, p.title)
	rl.SetTargetFPS(60) // définit les fps a x
	p.isRunning = true
	p.tex = rl.LoadTexture("../assets/Mossy_TileSet.png")
	p.textureCharacter = rl.LoadTexture("../assets/Tile.png")
	p.plateformSpriteSrc = rl.NewRectangle(251, 1583, 1000, 394)

	p.plateformSpriteDest = rl.NewRectangle(0, 30, 153, 49)

	//p.objDest = rl.NewRectangle(0, 0, 306, 166)
	p.textureMap = rl.LoadTexture("../assets/Mossy_TileSet.png")

	p.gargantuaTex = rl.LoadTexture("../assets/gargantua.png")
	p.gargantuaSrc = rl.NewRectangle(0, 0, 200, 200)
	p.gargantuaDest = rl.NewRectangle(0, 0, 700, 700)
	p.gargantuaSpeed = 1

	// source du joueur
	p.playerSrc = rl.NewRectangle(1, 195, 32, 32)                               // selectionne un bout d'image dans la sheet sprite
	p.playerDest = rl.NewRectangle(100, 0, 32, 32)                              // met une zone ou afficher ce bout d'image
	p.playerVector = rl.NewVector2((p.playerDest.Width), (p.playerDest.Height)) // permet de lui donner une position
	// p.tileSrc = rl.NewRectangle(1550, 110, 113, 185)
	// p.tileDest = rl.NewRectangle(0, 0, 153, 83)
	// initialistion du saut du joueur :
	p.playerCanJump = false
	p.playerIsJumping = false
	p.playerSpeed = 1.45
	p.gravity = 4.15

	p.cam2d = rl.NewCamera2D(rl.NewVector2(float32(p.width/2), float32(500)),
		rl.NewVector2(float32(p.playerDest.X-p.playerDest.Width/2), float32(p.playerDest.Y-p.playerDest.Height/4)), 0.0, 4.0)

	// rl.InitAudioDevice()
	//p.musicMenu = rl.LoadMusicStream("../audio/peace.wav")
	// p.musicIsPaused = false
	//rl.PlayMusicStream(p.musicMenu)

	p.hitboxWidth = p.playerDest.Width / 4
	p.hitboxHeight = p.playerDest.Height / 4
	p.hitboxX = p.playerDest.X + p.playerDest.Width/4  // Décalage horizontal pour centrer
	p.hitboxY = p.playerDest.Y + p.playerDest.Height/4 // Décalage vertical pour centrer
	p.adjustedHitbox = rl.NewRectangle(p.hitboxX, p.hitboxY, p.hitboxWidth, p.hitboxHeight)
	p.mapPath = "../assets/cat-mario.tmx"
	p.loadMap()
	for p.isRunning {
		//rl.UpdateMusicStream(p.musicMenu)
		p.input()
		p.update()
		p.render()

	}
	p.quit()
	rl.SetExitKey(0) // définit les boutons pour être ouvert fermé ?

}

func (g *gameEngine) loadMap() {
	file, err := os.Open("../assets/cat-mario.tmx") // va chercher le fichier
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier:", err)
		return
	}
	defer file.Close() // Assurez-vous de fermer le fichier à la fin

	// Créer un décodeur XML 
	decoder := xml.NewDecoder(file)

	g.myGroup.UnmarshalXML(decoder, xml.StartElement{}) // interprete fichier décoder ainsi que la première balise qu'il trouve dans le fihcier xml

	g.myGroup.DecodeGroup(&g.mapObject) // on lui donne la map à décoder
	fmt.Println("Contenu de la carte décoder :")
    fmt.Printf("%+v\n", g.mapObject)
}
func (w *gameEngine) input() { // récupère les inputs de la map

	if rl.IsKeyDown(rl.KeyUp) { // key left
		//w.adjustedHitbox.Y -= w.playerSpeed
		w.playerUp = true // dit qu'il va en haut
		w.playerDir = 17  // permet de set quel frame on veut dans la grille de sprite
		// pareil pour tous

	}
	// if rl.IsKeyDown(rl.KeyDown) { // key left
	// 	w.playerDest.Y += w.playerSpeed

	// 	w.playerMoving = true
	// 	w.playerDown = true
	// 	w.playerDir = 18
	// 	w.adjustedHitbox.Y += w.playerSpeed

	// }
	if rl.IsKeyDown(rl.KeyLeft) { // key left
		w.playerDest.X -= w.playerSpeed
		w.playerMoving = true
		w.playerDir = 5
		w.playerLeft = true
		w.adjustedHitbox.X -= w.playerSpeed

	}

	if rl.IsKeyDown(rl.KeyRight) { // key left
		w.playerDest.X += w.playerSpeed
		w.playerMoving = true
		w.playerRight = true
		w.playerDir = 6
		w.adjustedHitbox.X += w.playerSpeed

	}

	if rl.IsKeyPressed(rl.KeyM) { // key left
		w.musicIsPaused = !w.musicIsPaused

		w.playerUp = true
	}

}

func (p *gameEngine) update() { // va définir les mouvements du personnage
	p.isRunning = !rl.WindowShouldClose()
	p.gargantuaSrc.X = 0
	p.playerSrc.X = 7

	if !rl.CheckCollisionRecs(p.adjustedHitbox, p.plateformSpriteDest) && !p.playerCanJump { // sert à généré la gravité
		p.playerDest.Y += 4
		p.playerMoving = true
	}
	if rl.CheckCollisionRecs(p.adjustedHitbox, p.plateformSpriteDest) && p.playerUp {
		p.playerMoving = true
		p.playerIsJumping = true
		p.playerCanJump = false
		p.playerDest.Y -= 1
		// le joueur doit passer de 38 a
	}
	if !rl.CheckCollisionRecs(p.adjustedHitbox, p.plateformSpriteDest) && p.playerIsJumping {
		p.playerMoving = true
		p.playerDest.Y -= p.gravity * 2
		if p.playerDest.Y <= -20 {
			p.playerCanJump = false
			p.playerIsJumping = false

		}
	}

	if p.playerMoving {
		if p.playerLeft {
			p.playerDest.X -= p.playerSpeed
		}
		if p.playerRight {
			p.playerDest.X += p.playerSpeed
		}
		if p.FrameCount%8 == 1 {
			p.playerFrame++
		}

	}
	p.FrameCount++
	if p.playerFrame > 3 {
		p.playerFrame = 0
	}
	if p.framecountGargantua%12 == 1 {
		p.gargantuaSrc.X += 200
		p.gargantuaSrc.Y += 200
	}

	p.playerSrc.X = p.playerSrc.Width * float32(p.playerFrame)
	p.playerSrc.Y = p.playerSrc.Width * float32(p.playerDir)

	rl.UpdateMusicStream(p.musicMenu)
	if p.musicIsPaused {
		rl.PauseMusicStream(p.musicMenu)
	} else {
		rl.ResumeMusicStream(p.musicMenu)
	}

	p.cam2d.Target = rl.NewVector2(float32(p.playerDest.X-p.playerDest.Width/2), float32(p.playerDest.Y-p.playerDest.Height/4))
	p.playerMoving = false
	p.playerUp, p.playerDown, p.playerRight, p.playerLeft = false, false, false, false
}

//_________________________________________________________________Menu_______________________________________________________________//

func (g *gameEngine) render() { // permet le rendu de la fenetre c'est à dire les dessins
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)
	rl.BeginMode2D(g.cam2d)
	rl.DrawText("CAT MARIO", 0, 0, 35, rl.White)
	// // faire une condition pour dire tant que le joueur n'est pas mort :
	//rl.DrawTexture(g.TxSprites, g.frameRec)
	// // sourceTest := rl.Rectangle{}
	// // destRecTest := rl.Rectangle{}
	// // originTest := rl.Vector2{}

	//rl.DrawTexture(g.texture,0,0,rl.White)
	g.drawScene()
	rl.EndMode2D()
	rl.EndDrawing()

}
func (g *gameEngine) drawScene() {

	g.adjustedPlayerDest = rl.NewRectangle(g.playerDest.X-g.playerDest.Width/4, g.playerDest.Y-g.playerDest.Height/4, g.playerDest.Width, g.playerDest.Height)
	g.adjustedHitbox.X = g.adjustedPlayerDest.X + g.playerDest.Width - 46
	g.adjustedHitbox.Y = g.adjustedPlayerDest.Y + g.playerDest.Height - 39

	rl.DrawRectangleLines(int32(g.adjustedHitbox.X), int32(g.adjustedHitbox.Y), int32(g.hitboxX+g.hitboxWidth), int32(g.hitboxY+g.hitboxHeight), rl.White)

	rl.DrawTexturePro(g.textureMap, g.plateformSpriteSrc, g.plateformSpriteDest, rl.NewVector2(0, 0), 0, rl.White)

	rl.DrawTexturePro(g.textureMap, g.plateformSpriteSrc, g.plateformSpriteDest, rl.NewVector2(0, 0), 0, rl.White)
	rl.DrawTexturePro(g.gargantuaTex, g.gargantuaSrc, g.gargantuaDest, rl.NewVector2(10, 10), 0, rl.White)
	rl.DrawTexturePro(g.textureCharacter, g.playerSrc, g.playerDest, g.playerVector, 0, rl.White) // drawTextureMario
}

func (p *gameEngine) quit() {
	rl.UnloadTexture(p.textureCharacter)
	rl.UnloadTexture(p.textureMap)
	rl.UnloadTexture(p.gargantuaTex)
	rl.UnloadMusicStream(p.musicMenu)
	rl.CloseAudioDevice()
	rl.CloseWindow()
}

// _________________________________________________________________Menu_______________________________________________________________//
// iota est un identificateur prédéclaré représentant le numéro ordinal entier non typé de la spécification
// const actuelle dans une déclaration const (généralement entre parenthèses). Il est indexé à zéro.
// const (
// 	MenuDisplay = iota
// 	Game
// 	Options
// )

// var currentGameState = MenuDisplay

// func (p *gameEngine) display() {
// 	rl.UpdateMusicStream(p.musicMenu)
// 	// defer rl.CloseWindow()

// 	// rl.SetTargetFPS(60)
// 	rl.BeginMode2D(p.cam2d)
// 	for p.isRunning {

// 		switch currentGameState {
// 		case MenuDisplay:
// 			if rl.IsKeyReleased(rl.KeyEnter) {
// 				currentGameState = Game
// 			} else if rl.IsKeyReleased(rl.KeyO) {
// 				currentGameState = Options
// 			} else if rl.IsKeyReleased(rl.KeyEscape) {
// 				p.quit()
// 			}
// 		case Game:
// 			if rl.IsKeyReleased(rl.KeyEscape) {
// 				currentGameState = MenuDisplay

// 			}
// 		case Options:
// 			if rl.IsKeyReleased(rl.KeyEscape) {
// 				currentGameState = MenuDisplay
// 			}
// 		}

// 		rl.BeginDrawing()
// 		rl.ClearBackground(rl.RayWhite)

// 		switch currentGameState {
// 		case MenuDisplay:
// 			// Menue
// 			rl.DrawText("PLAY - Appuyez sur ENTER pour jouer :", 100, 150, 35, rl.White)

// 			rl.DrawText("OPTIONS - Appuyez sur O pour accéder aux options :", 100, 300, 35, rl.White)

// 			rl.DrawText("QUIT - Appuyez sur ESCAPE pour quitter :", 100, 450, 35, rl.White)

// 			rl.ClearBackground(rl.DarkBlue)

// 		case Game:
// 			// JEUX
// 			rl.ClearBackground(rl.Black)
// 			rl.DrawText("JEU EN COURS - Appuyez sur ESCAPE pour revenir au menu :", 10, 10, 13, rl.White)

// 			p.input()
// 			p.update()
// 			p.render()

// 		case Options:
// 			// OPTION //

// 			rl.ClearBackground(rl.White)

// 			rl.DrawText("Setings Glogbal :", 580, 1, 35, rl.White)

// 			//ESSAIS DE BOUTTONS//
// 			if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(15, 90, 50, 50)) {
// 				rl.DrawRectangle(15, 90, 50, 50, rl.White)
// 				// if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
// 				// }
// 			} else {
// 				rl.DrawRectangle(15, 90, 50, 50, rl.LightGray)
// 			}
// 			if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(70, 90, 50, 50)) {
// 				rl.DrawRectangle(70, 90, 50, 50, rl.Yellow)
// 				// if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
// 				// }
// 			} else {
// 				rl.DrawRectangle(70, 90, 50, 50, rl.Yellow)
// 			}
// 			//ESSAIS DE BOUTTONS//
// 			// buttons := []struct {
// 			// 	Bounds rl.Rectangle
// 			// 	Text   string
// 			// }{
// 			// 	{rl.NewRectangle(screenWidth/20, screenHeight/20, 150, 40), "Back"},
// 			// 	{rl.NewRectangle(screenWidth-(150+screenWidth/20), screenHeight-(40+screenHeight/20), 150, 40), "Quit"},
// 			// }

// 			// for _, button := range buttons {
// 			// 	color := rl.Yellow
// 			// 	if rl.CheckCollisionPointRec(rl.GetMousePosition(), button.Bounds) {
// 			// 		color = rl.DarkGray
// 			// 		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
// 			// 			switch button.Text {
// 			// 			case "Back":
// 			// 				currentScreen = 1
// 			// 			case "Quit":
// 			// 				rl.UnloadMusicStream(bgMusic)
// 			// 				rl.CloseAudioDevice()
// 			// 				rl.CloseWindow()
// 			// 				return
// 			// 			}

// 			rl.DrawText("FPS-TOUCHES", 580, 85, 35, rl.White)

// 			// QUiTTEZ //
// 			rl.DrawText("OPTIONS - Appuyez sur ESCAPE pour revenir au menu :", 300, 45, 35, rl.Brown)
// 		}
// 		rl.EndMode2D()
// 		rl.EndDrawing()
// 	}
// }
//test
