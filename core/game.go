package main

import rl "github.com/gen2brain/raylib-go/raylib"
<<<<<<< HEAD

=======
const (
	screenWidth = 1366
	screenHeight = 768
)
func input(){}
func updaete(){}
func render(){
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black )
}
>>>>>>> 25faa2b3a68820e538e79879881922bdb3717aee
func Play() {
	rl.InitWindow(1366, 768, "Donkey Kong Mario") // définit la taille de la fenetre
	defer rl.CloseWindow() // définit les boutons pour être ouvert fermé ?
	rl.SetTargetFPS(60) // définit les fps a x

	for !rl.WindowShouldClose() { // boucle tant que c'est pas fermé
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

		rl.EndDrawing()
	}
	//firstx	
}