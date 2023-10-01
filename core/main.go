package main

import rl "github.com/gen2brain/raylib-go/raylib"

type ObjectWhoMoove struct {
	Position rl.Vector2
		
							   // de tel objet c'est ça
							   // objectif, en affiché un simple
}

type gameEngine struct {
	width int32
	heigh int32
	title string
	//maxBarril int
	score int
	dead bool
	//mario ObjectWhoMoove
	//barril ObjectWhoMoove
	musicMenu rl.Music
	musicIsPaused bool

	textureCharacter rl.Texture2D
	textureMap rl.Texture2D

	objSrc rl.Rectangle
	objDest rl.Rectangle
	srcMap  []string
	plateformSpriteSrc rl.Rectangle
	plateformSpriteDest rl.Rectangle

	wallSpriteSrc rl.Rectangle
	littleSpriteSrc rl.Rectangle
	tex rl.Texture2D

	playerSrc rl.Rectangle
	playerDest rl.Rectangle
	playerVector rl.Vector2
	playerSpeed float32
	playerMoving bool
	playerDir int
	playerUp,playerDown,playerRight,playerLeft bool
	cam2d rl.Camera2D
	isRunning bool
	playerFrame int
	FrameCount int
	gargantuaDest rl.Rectangle
	gargantuaTex rl.Texture2D
	gargantuaSrc rl.Rectangle
	gargantuaSpeed int
	framecountGargantua int

	//mapFile  string
	//playerCanJump bool
	//timePlayed float32


	tileDest rl.Rectangle
	tileSrc rl.Rectangle
	tileMap []int
	tileMapLink string
	//srcMap []string
	gravity int
	playerCanJump bool
	mapW, mapH int
}


func main() {

	var game gameEngine // défini game a la struct window
	game.Init(1366, 768," Donkey Kong Mario", true, false, 0) // initialise la fenetre : donc attribue les valeurs automatiquement ?
	// qu'est game.mario ?????????
	rl.InitAudioDevice()
	game.initGame()

}

