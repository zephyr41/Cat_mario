package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type gameEngine struct {
	/* afficher l'explication
	   ##########################################################
	   #          TOUT EN RAPPORT AVEC LA MAP                   #
	   #     Largeur, music, la partie est elle en cours...     #
	   ##########################################################
	*/
	width         int32
	heigh         int32
	title         string
	score         int
	dead          bool
	musicMenu     rl.Music
	musicIsPaused bool
	isRunning     bool // savoir si la partie tourne
	FrameCount    int

	/* afficher l'explication
	   ##########################################################
	   #          TOUT EN RAPPORT AVEC LES TEXTURES             #
	   # Source des sprites, leurs destinations.... (affichage) #
	   ##########################################################
	*/
	textureCharacter    rl.Texture2D
	textureMap          rl.Texture2D
	plateformSpriteSrc  rl.Rectangle
	plateformSpriteDest rl.Rectangle
	gargantuaDest       rl.Rectangle
	gargantuaTex        rl.Texture2D
	gargantuaSrc        rl.Rectangle
	gargantuaSpeed      int

	/* afficher l'explication
	   ##########################################################
	   #          TOUT EN RAPPORT AVEC LE JOUEUR                #
	   #Source  joueur, gravité, mouvement, affichage joueur... #
	   ##########################################################
	*/
	playerSrc    rl.Rectangle
	playerDest   rl.Rectangle
	playerVector rl.Vector2

	playerMoving                                  bool
	playerDir                                     int
	playerUp, playerDown, playerRight, playerLeft bool

	playerSpeed     float32
	playerIsJumping bool
	playerCanJump   bool
	gravity         float32
	jumpHmax        int

	playerFrame int

	cam2d rl.Camera2D

	hitboxHeight       float32
	hitboxWidth        float32
	hitboxX            float32
	hitboxY            float32
	adjustedHitbox     rl.Rectangle
	adjustedPlayerDest rl.Rectangle

	mapPath string

	grassSprite rl.Texture2D
	propsSprite rl.Texture2D
	tex         rl.Texture2D
	life        int
	tileDest    rl.Rectangle
	tileSrc     rl.Rectangle
	tileMap     []int
	srcMap      []string
	mapW, mapH  int
}

func main() {

	var game gameEngine                                        // défini game a la struct window
	game.Init(1366, 768, " Donkey Kong Mario", true, false, 0) // initialise la fenetre : donc attribue les valeurs automatiquement ?

	rl.InitAudioDevice()
	game.initGame()

}
