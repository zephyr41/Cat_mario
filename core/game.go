package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

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
	p.playerCanJump = false
	p.playerIsJumping = false
	p.gargantuaTex = rl.LoadTexture("../assets/gargantua.png")
	p.gargantuaSrc = rl.NewRectangle(0, 0, 200, 200)
	p.gargantuaDest = rl.NewRectangle(0, 0, 700, 700)
	p.gargantuaSpeed = 1

	// source du joueur
	p.playerSrc = rl.NewRectangle(1, 195, 32, 32)                               // selectionne un bout d'image dans la sheet sprite
	p.playerDest = rl.NewRectangle(0, 0, 32, 32)                                // met une zone ou afficher ce bout d'image
	p.playerVector = rl.NewVector2((p.playerDest.Width), (p.playerDest.Height)) // permet de lui donner une position
	// p.tileSrc = rl.NewRectangle(1550, 110, 113, 185)
	// p.tileDest = rl.NewRectangle(0, 0, 153, 83)
	// initialistion du saut du joueur :
	p.playerIsJumping = false
	p.playerSpeed = 1.45
	p.gravity = 1.0
	p.jumpForce = -20.0

	p.cam2d = rl.NewCamera2D(rl.NewVector2(float32(p.width/2), float32(500)),
		rl.NewVector2(float32(p.playerDest.X-p.playerDest.Width/2), float32(p.playerDest.Y-p.playerDest.Height/4)), 0.0, 4.0)

	// rl.InitAudioDevice()
	//p.musicMenu = rl.LoadMusicStream("../audio/peace.wav")
	// p.musicIsPaused = false
	//rl.PlayMusicStream(p.musicMenu)
	//p.loadMap()
	p.hitboxWidth = p.playerDest.Width / 4
	p.hitboxHeight = p.playerDest.Height / 4
	p.hitboxX = p.playerDest.X + p.playerDest.Width/4  // Décalage horizontal pour centrer
	p.hitboxY = p.playerDest.Y + p.playerDest.Height/4 // Décalage vertical pour centrer

	p.adjustedHitbox = rl.NewRectangle(p.hitboxX, p.hitboxY, p.hitboxWidth, p.hitboxHeight)
	for p.isRunning {
		//rl.UpdateMusicStream(p.musicMenu)
		p.input()
		p.update()
		p.render()

	}
	p.quit()
	rl.SetExitKey(0) // définit les boutons pour être ouvert fermé ?

}

func (p *gameEngine) loadMap() {
	file, err := os.ReadFile(p.tileMapLink)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	const mapPath = "..assets/cat-mario.tmx"

	fmt.Println("Map width:", p.mapFileWidth)
	fmt.Println("Map height:", p.mapFileHeight)
	fmt.Println("Map Link", p.tileMapLink)
	fmt.Println("Map File", p.mapFile)

	remNewLines := strings.Replace(string(file), "\n", " ", -1)
	sliced := strings.Split(remNewLines, " ")

	p.mapW = -1
	p.mapH = -1

	for i := 0; i < len(sliced); i++ {

		s, _ := strconv.ParseInt(sliced[i], 10, 64)
		fmt.Println("slice", i, sliced[i], "s", s)
		m := int(s)
		if p.mapW == -1 {
			p.mapW = m

		} else if p.mapH == -1 {
			p.mapH = m
		} else if i < p.mapW*p.mapH+2 {
			p.tileMap = append(p.tileMap, m)
		} else {
			p.srcMap = append(p.srcMap, sliced[i])
		}
		if len(p.tileMap) > p.mapW*p.mapH {
			p.tileMap = p.tileMap[:len(p.tileMap)-1]
		}
		p.mapW = 5
		p.mapH = 5

	}
}

func (w *gameEngine) input() { // récupère les inputs de la map

	if rl.IsKeyDown(rl.KeyUp) && !w.playerIsJumping { // key left
		w.playerCanJump = false
		w.playerIsJumping = true
		if w.playerIsJumping {
			w.playerDest.Y -= w.jumpForce
		}
		w.adjustedHitbox.Y -= w.playerSpeed
		w.playerUp = true // dit qu'il va en haut
		w.playerDir = 17  // permet de set quel frame on veut dans la grille de sprite
		// pareil pour tous

	}
	if rl.IsKeyDown(rl.KeyDown) { // key left
		w.playerDest.Y += w.playerSpeed

		w.playerMoving = true
		w.playerDown = true
		w.playerDir = 18
		w.adjustedHitbox.Y += w.playerSpeed

	}
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
		w.adjustedHitbox.Y += w.playerSpeed

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

	if !rl.CheckCollisionRecs(p.adjustedHitbox, p.plateformSpriteDest) && !p.playerCanJump {
		p.playerDest.Y += 1

	}
	// if !p.playerIsJumping && p.playerDest.Y <= p.plateformSpriteDest.Y{
	// 	for p.hitboxY < p.plateformSpriteDest.Y{
	// 		p.playerDest.Y -= 10
	// 	}
	if p.playerMoving {
		if p.playerDown {
			p.playerDest.Y += p.playerSpeed
		}
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
	//adjustedPlayerDest := rl.NewRectangle(p.playerDest.X-p.playerDest.Width/2, p.playerDest.Y-p.playerDest.Height/2, p.playerDest.Width, p.playerDest.Height)

	if rl.CheckCollisionRecs(p.adjustedHitbox, p.plateformSpriteDest) {
		p.playerDest.Y -= 0
		if p.playerIsJumping {
			p.playerDest.Y -= 0
			fmt.Println("le personnage peut sauter")
		} else {
			fmt.Println("le personnage ne peux pas sauter")
		}

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
	// for i := 0; i < len(g.tileMap); i++ { // sert à mapper la carte :
	// 	if g.tileMap[i] != 0 {
	// 		g.tileDest.X = g.tileDest.Width * float32(i%g.mapW)
	// 		g.tileDest.Y = g.tileDest.Height * float32(i%g.mapH)
	// 		g.objDest.X = g.tileDest.X
	// 		g.objDest.Y = g.tileDest.Y
	// 		if g.srcMap[i] == "p" {
	// 			g.tex = g.textureMap
	// 			g.objSrc = g.plateformSpriteSrc
	// 			g.objDest = g.plateformSpriteDest
	// 		}
	// 		g.tileSrc.X = g.tileSrc.Width * float32((g.tileMap[i]-1)%int(g.textureMap.Width/int32(g.tileSrc.Width)))
	// 		g.tileSrc.Y = g.tileSrc.Height * float32((g.tileMap[i]-1)%int(g.textureMap.Width/int32(g.tileSrc.Width)))

	// 	}

	// }
	g.adjustedPlayerDest = rl.NewRectangle(g.playerDest.X-g.playerDest.Width/4, g.playerDest.Y-g.playerDest.Height/4, g.playerDest.Width, g.playerDest.Height)
	g.adjustedHitbox.X = g.adjustedPlayerDest.X + g.playerDest.Width - 46
	g.adjustedHitbox.Y = g.adjustedPlayerDest.Y + g.playerDest.Height - 39

	rl.DrawRectangleLines(int32(g.adjustedHitbox.X), int32(g.adjustedHitbox.Y), int32(g.hitboxX+g.hitboxWidth), int32(g.hitboxY+g.hitboxHeight), rl.Red)
	rl.DrawRectangleLines(int32(g.playerDest.X), int32(g.playerDest.Y), int32(g.playerDest.X+g.playerDest.Width), int32(g.playerDest.Y+g.playerDest.Height), rl.Blue)

	rl.DrawTexturePro(g.textureMap, g.plateformSpriteSrc, g.plateformSpriteDest, rl.NewVector2(0, 0), 0, rl.Red)
	//rl.DrawTexturePro(g.gargantuaTex, g.plateformSpriteSrc, g.plateformSpriteDest, rl.NewVector2(0, 0), 0, rl.Red)
	//rl.DrawRectangleV(rl.NewVector2(g.plateformSpriteDest.X,g.plateformSpriteDest.Y ), rl.NewVector2(g.plateformSpriteDest.Width,g.plateformSpriteDest.Height ),rl.Beige)
	rl.DrawTexturePro(g.gargantuaTex, g.gargantuaSrc, g.gargantuaDest, rl.NewVector2(10, 10), 0, rl.White)
	rl.DrawTexturePro(g.textureCharacter, g.playerSrc, g.playerDest, g.playerVector, 0, rl.Red) // drawTextureMario
	//rl.DrawRectangleV(rl.NewVector2(0,0), rl.NewVector2(40,40), rl.Blue)

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
