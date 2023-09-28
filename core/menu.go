package main

import rl "github.com/gen2brain/raylib-go/raylib"

func Menu (w gameEngine){
	rl.InitWindow(1366, 768, "Donkey Kong Mario") // définit la taille de la fenetre
	defer rl.CloseWindow() // définit les boutons pour être ouvert fermé ?
	rl.SetTargetFPS(60) // définit les fps a x

	for !rl.WindowShouldClose() { // boucle tant que c'est pas fermé
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

		rl.EndDrawing()
	}
	
}
type Bouton struct {  // On défini un nouveau type appelé Bouton
	Texte     string   
	Rectangle rl.Rectangle  // Un rectangle pour définir la position et la taille du Bouton
	Couleur   rl.Color   // La couleur du Bouton
}