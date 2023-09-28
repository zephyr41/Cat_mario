package main

import rl "github.com/gen2brain/raylib-go/raylib"


func input(){}
func update(){}
func render(){
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black )
}

func Play(screenWidth int32, screenHeight int32) {
	rl.InitWindow(screenWidth, screenHeight, "Donkey Kong Mario") // définit la taille de la fenetre
	defer rl.CloseWindow() // définit les boutons pour être ouvert fermé ?
	rl.SetTargetFPS(60) // définit les fps a x
	for !rl.WindowShouldClose() { // boucle tant que c'est pas fermé
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)
		rl.EndDrawing()
	}
}