package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (p *gameEngine) Init(width int32, height int32, title string, isRunning bool, dead bool, score int) { // initialise les propriété de la fenetre
	// Pour selectionner notre joueur :
	// va déssiner un rectangle dans la feuille de sprite pour pouvoir dessiner le chat : il aura : x & y est position de ou il est sur la feuille
	// width et heigt seront la largeur de la fenetre
	// src est la source (la feuille) donc ça position dans le sprite
	// Dest est la ou on l'envoie dans notre (notre fênetre de jeu)
	p.width = width
	p.heigh = height
	p.title = title
	p.isRunning = isRunning
	p.dead = dead
	p.score = score

}

func (p *gameEngine) initGame() { // Initialise le jeu, en créant la fenêtre ,
	rl.InitWindow(p.width, p.heigh, p.title)
	rl.SetTargetFPS(60) // définit les fps a x ici 60
	p.isRunning = true
	p.textureCharacter = rl.LoadTexture("../assets/Tile.png")
	p.grassSprite = rl.LoadTexture("../assets/TX_Tileset_Ground.png")
	p.propsSprite = rl.LoadTexture("../assets/TX_Village_Props.png")

	p.plateformSpriteSrc = rl.NewRectangle(251, 1583, 1000, 394)

	p.plateformSpriteDest = rl.NewRectangle(0, 30, 153, 49)
	p.tex = rl.LoadTexture("assets/TX_Tileset_Ground.png")
	p.textureMap = rl.LoadTexture("../assets/Mossy_TileSet.png")

	p.gargantuaTex = rl.LoadTexture("../assets/gargantua.png")
	p.gargantuaSrc = rl.NewRectangle(0, 0, 200, 200)
	p.gargantuaDest = rl.NewRectangle(170, 465, 100, 100)
	p.gargantuaSpeed = 2

	// source du joueur
	p.playerSrc = rl.NewRectangle(1, 195, 32, 32)                               // selectionne un bout d'image dans la sheet sprite
	p.playerDest = rl.NewRectangle(40, 600, 32, 32)                             // met une zone ou afficher ce bout d'image
	p.playerVector = rl.NewVector2((p.playerDest.Width), (p.playerDest.Height)) // permet de lui donner une position

	p.tileDest = rl.NewRectangle(0, 0, 32, 32)
	p.tileSrc = rl.NewRectangle(0, 0, 32, 32)
	// initialistion du saut du joueur :
	p.playerCanJump = false
	p.playerIsJumping = false
	p.playerSpeed = 5
	p.gravity = 4.15
	p.tileDest = rl.NewRectangle(0, 0, 32, 32)
	p.tileSrc = rl.NewRectangle(0, 0, 32, 32)
	p.cam2d = rl.NewCamera2D(rl.NewVector2(float32(p.width/2), float32(500)),
		rl.NewVector2(float32(p.playerDest.X-p.playerDest.Width/4), float32(p.playerDest.Y-p.playerDest.Height/4)), 0.0, 2.0)

	p.hitboxWidth = p.playerDest.Width / 4
	p.hitboxHeight = p.playerDest.Height / 4
	p.hitboxX = p.playerDest.X + p.playerDest.Width/4  // Décalage horizontal pour centrer la hitbox
	p.hitboxY = p.playerDest.Y + p.playerDest.Height/4 // Décalage vertical pour centrer la hitbox
	p.adjustedHitbox = rl.NewRectangle(p.hitboxX, p.hitboxY, p.hitboxWidth, p.hitboxHeight)
	p.mapPath = "../assets/one.map" // lien de la map
	p.life = 3                      // lien des vies
	p.loadMap()                     // charge la map
	p.display()                     // affiche le mnu

	p.quit()
	rl.SetExitKey(0) // définit Echap comme exit Key

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
	rl.BeginMode2D(p.cam2d) // initialise la caméra
	for p.isRunning {       // tant que le jeu est en cours

		switch currentGameState {
		case MenuDisplay: // set les keys pour attrivuer au valeur
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

		switch currentGameState { // si on est dans le menu afficher :
		case MenuDisplay:
			// Menu
			rl.DrawText("PLAY - Appuyez sur ENTER pour jouer :", 100, 150, 35, rl.White)

			rl.DrawText("OPTIONS - Appuyez sur O pour accéder aux options :", 100, 300, 35, rl.White)

			rl.DrawText("QUIT - Appuyez sur ESCAPE pour quitter :", 100, 450, 35, rl.White)

			rl.ClearBackground(rl.DarkBlue)

		case Game:
			// JEUX
			rl.ClearBackground(rl.Black)
			rl.DrawText("JEU EN COURS - Appuyez sur ESCAPE pour revenir au menu :", 10, 10, 13, rl.White)

			for p.isRunning {
				p.input()
				p.update()
				p.render()

			}

		case Options:
			// OPTION // pour les menus

			rl.ClearBackground(rl.White)

			rl.DrawText("Setings Glogbal :", 580, 1, 35, rl.White)

			if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(15, 90, 50, 50)) {
				rl.DrawRectangle(15, 90, 50, 50, rl.White)

			} else {
				rl.DrawRectangle(15, 90, 50, 50, rl.LightGray)
			}
			if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(70, 90, 50, 50)) {
				rl.DrawRectangle(70, 90, 50, 50, rl.Yellow)

				// }
			} else {
				rl.DrawRectangle(70, 90, 50, 50, rl.Yellow)
			}

			rl.DrawText("FPS-TOUCHES", 580, 85, 35, rl.White)

			// QUiTTEZ //
			rl.DrawText("OPTIONS - Appuyez sur ESCAPE pour revenir au menu :", 300, 45, 35, rl.Brown)
		}
		rl.EndMode2D()
		rl.EndDrawing()
	}
}

func (g *gameEngine) loadMap() { // permet de load map
	f, err := os.ReadFile(g.mapPath) // ouvre le fichier

	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(`\r?\n`)
	remNewLines := re.ReplaceAllString(string(f), " ")

	sliced := strings.Split(remNewLines, " ")
	//fmt.Println("sliced:", sliced)
	g.mapW = -1
	g.mapH = -1
	for i := 0; i < len(sliced); i++ {

		s, _ := strconv.ParseInt(sliced[i], 10, 64)
		//fmt.Println("slice", i, sliced[i], "s", s)
		m := int(s)
		if g.mapW == -1 {
			g.mapW = m

		} else if g.mapH == -1 {
			g.mapH = m
		} else if i < g.mapW*g.mapH+2 {
			g.tileMap = append(g.tileMap, m)
		} else {
			g.srcMap = append(g.srcMap, sliced[i])
		}

	}
	if len(g.tileMap) > g.mapW*g.mapH {
		g.tileMap = g.tileMap[:len(g.tileMap)-1]
	}
}
func (w *gameEngine) input() { // récupère les inputs de la map

	if rl.IsKeyDown(rl.KeyUp) {
		w.playerUp = true // dit qu'il va en haut
		w.playerDir = 17  // permet de set quel frame on veut dans la grille de sprite
	}

	if rl.IsKeyDown(rl.KeyLeft) { // key left
		w.playerDest.X -= w.playerSpeed
		w.playerMoving = true
		w.playerDir = 5
		w.playerLeft = true
		w.adjustedHitbox.X -= w.playerSpeed // permet d'ajuster la hibox

	}

	if rl.IsKeyDown(rl.KeyRight) { // key left
		w.playerDest.X += w.playerSpeed
		w.playerMoving = true
		w.playerRight = true
		w.playerDir = 6
		w.adjustedHitbox.X += w.playerSpeed // ajuste la hitbox

	}

}

func (p *gameEngine) update() { // va définir les mouvements du personnage en les mettans en jours
	p.isRunning = !rl.WindowShouldClose() // tant que la fentre run
	p.gargantuaSrc.X = 0
	p.playerSrc.X = 7
	if p.playerDest.Y >= 1000 { // permet de faire respawn le joueur a son emplacement
		p.playerDest.Y = 600
		p.playerDest.X = 60
	}
	if !rl.CheckCollisionRecs(p.adjustedHitbox, p.plateformSpriteDest) && !p.playerCanJump { // sert à généré la gravité tant que le joueur n'est pas en collision
		p.playerDest.Y += 5
		p.playerMoving = true
		p.playerCanJump = false
		p.playerUp = false
	}

	// si le joueur n'est pas en collision et qu'il saute, alors il monte de 160 pixel
	if !rl.CheckCollisionRecs(p.playerDest, p.tileDest) && p.playerCanJump && p.playerUp {
		p.playerMoving = true
		p.jumpHmax = int(p.tileDest.Y) - 160
		p.jumpHmax -= 5
		p.playerDest.Y -= 5
		if p.playerDest.Y <= float32(p.jumpHmax) {
			p.playerCanJump = false

			p.playerUp = false
		}
	}
	// permet de faire bouger le joueur tout en mettant a jours la feuille de sprite pour les animations

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
	if p.FrameCount%12 == 1 {
		p.gargantuaSrc.X += 200 // permet d'animer le troue noir
		p.gargantuaSrc.Y += 200
	}
	if p.playerFrame > 3 { // permet de dire que tout les x temps l'image du chat de base devient x
		p.playerFrame = 0
	}

	p.playerSrc.X = p.playerSrc.Width * float32(p.playerFrame) // change la source du sprite pour animer le chat en largeur
	p.playerSrc.Y = p.playerSrc.Width * float32(p.playerDir)   // change la source du sprite pour animer le chat en hauteur

	rl.UpdateMusicStream(p.musicMenu) // lance une musique qui ne marche pas :(
	if p.musicIsPaused {
		rl.PauseMusicStream(p.musicMenu)
	} else {
		rl.ResumeMusicStream(p.musicMenu)
	}

	p.cam2d.Target = rl.NewVector2(float32(p.playerDest.X-p.playerDest.Width/2), float32(p.playerDest.Y-p.playerDest.Height/4)) // met a jour la caméra du chat
	p.playerMoving = false                                                                                                      // dit qu'il ne bouge plus si aucune touche n'est préssé , (func input)

}

//_________________________________________________________________Menu_______________________________________________________________//

func (g *gameEngine) render() { // permet le rendu de la fenetre c'est à dire les dessins
	rl.BeginDrawing()                                // commence le decssin
	rl.ClearBackground(rl.Black)                     // met du noir en fond
	rl.BeginMode2D(g.cam2d)                          // active la caméra
	rl.DrawText("CAT MARIO", -45, 500, 35, rl.White) // tuto pour le joueur
	rl.DrawText("Press <-  -> for moove ", -60, 550, 1, rl.White)
	rl.DrawText("Press up to jump ", 90, 550, 1, rl.White)
	g.drawScene()   // déssine la scène
	rl.EndMode2D()  // enlève la caméra si l'on quitte
	rl.EndDrawing() // fin du desisn

}
func (g *gameEngine) drawScene() {
	// ajuste la hitbox pour centrer sur le chat est géré la collision
	g.adjustedPlayerDest = rl.NewRectangle(g.playerDest.X-g.playerDest.Width/4, g.playerDest.Y-g.playerDest.Height/4, g.playerDest.Width, g.playerDest.Height)
	g.adjustedHitbox.X = g.adjustedPlayerDest.X + g.playerDest.Width - 46
	g.adjustedHitbox.Y = g.adjustedPlayerDest.Y + g.playerDest.Height - 49

	// fais une boucle qui parourt, et pour chaque nombre initialise la source et la destion,

	for i := 0; i < len(g.tileMap); i++ {
		if g.tileMap[i] != 0 {
			g.tileDest.X = g.tileDest.Width * float32(i%g.mapW)
			g.tileDest.Y = g.tileDest.Height * float32(i/g.mapW)

			if g.srcMap[i] == "g" {
				g.tex = g.grassSprite
				// pareil ligne du bas mais pour de la terre
			}
			if g.srcMap[i] == "p" {
				g.tex = g.propsSprite
				// si sur la one.map il est affiché p alors la texture devient les props
			}

			// change la source de l'image pour afficher la case x du sprite en question
			g.tileSrc.X = g.tileSrc.Width * float32((g.tileMap[i]-1)%int(g.tex.Width/int32(g.tileSrc.Width)))
			g.tileSrc.Y = g.tileSrc.Height * float32((g.tileMap[i]-1)/int(g.tex.Width/int32(g.tileSrc.Width)))
			rl.DrawTexturePro(g.tex, g.tileSrc, g.tileDest, rl.NewVector2(g.tileDest.Width, g.tileDest.Height), 0, rl.White)
			if rl.CheckCollisionRecs(g.playerDest, g.tileDest) {
				g.playerMoving = true
				g.playerIsJumping = false // le joueur n'est pas entrain de sauter
				g.playerCanJump = true

				g.playerDest.Y -= 1 // permet de gérér la collision
				g.playerUp = false
			}

		}

	}

	rl.DrawTexturePro(g.gargantuaTex, g.gargantuaSrc, g.gargantuaDest, rl.NewVector2(0, 0), 0, rl.White) // dessigne gargantua, s/O interstellar
	rl.DrawTexturePro(g.textureCharacter, g.playerSrc, g.playerDest, g.playerVector, 0, rl.White)        // dessine le joueur
}

func (p *gameEngine) quit() {
	rl.UnloadTexture(p.textureCharacter) // unload la texture
	rl.UnloadTexture(p.textureMap)       //
	rl.UnloadTexture(p.gargantuaTex)     // unload les textures
	rl.UnloadMusicStream(p.musicMenu)
	rl.UnloadTexture(p.propsSprite)
	rl.UnloadTexture(p.grassSprite)
	rl.CloseAudioDevice() //
	rl.CloseWindow()
}
