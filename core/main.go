package main

import (
	"image"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/lafriks/go-tiled"
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
	test1         [4]float32
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
	wallSpriteSrc       rl.Rectangle
	littleSpriteSrc     rl.Rectangle
	tex                 rl.Texture2D
	plateformSpriteSrc  rl.Rectangle
	plateformSpriteDest rl.Rectangle
	objSrc              rl.Rectangle
	objDest             rl.Rectangle
	gargantuaDest       rl.Rectangle
	gargantuaTex        rl.Texture2D
	gargantuaSrc        rl.Rectangle
	gargantuaSpeed      int
	framecountGargantua int
	playerJumpVelocity  float32
	testjump            bool

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

	// playerCanJump bool
	cam2d rl.Camera2D

	hitboxHeight       float32
	hitboxWidth        float32
	hitboxX            float32
	hitboxY            float32
	adjustedHitbox     rl.Rectangle
	adjustedPlayerDest rl.Rectangle

	// tileDest                    rl.Rectangle
	// tileSrc                     rl.Rectangle
	// tileMap                     []int
	mapPath string
	// srcMap                      []string
	// mapFileWidth, mapFileHeight int
	// img                         image.NRGBA
	mapObject tiled.Map
	myGroup   tiled.Group
	img       *image.NRGBA
}

func main() {

	var game gameEngine                                        // défini game a la struct window
	game.Init(1366, 768, " Donkey Kong Mario", true, false, 0) // initialise la fenetre : donc attribue les valeurs automatiquement ?
	// qu'est game.mario ?????????
	rl.InitAudioDevice()
	game.initGame()

}
