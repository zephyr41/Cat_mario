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
	maxBarril int
	score int
	dead bool
	mario ObjectWhoMoove
	//barril ObjectWhoMoove
	texture rl.Texture2D
	playerSrc rl.Rectangle
	playerDest rl.Rectangle
	playerVector rl.Vector2
	playerSpeed float32
	
	isRunning bool



}


func main() {

	var game gameEngine // défini game a la struct window
	game.Init(1366, 768," Donkey Kong Mario", true, false, 0) // initialise la fenetre : donc attribue les valeurs automatiquement ?
	// qu'est game.mario ?????????

	game.initGame()

}

