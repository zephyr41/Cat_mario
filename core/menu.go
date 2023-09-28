package main

import rl "github.com/gen2brain/raylib-go/raylib"

func Menu (){
	rl.InitWindow(1366, 768, "Donkey Kong Mario") // définit la taille de la fenetre
	defer rl.CloseWindow() // définit les boutons pour être ouvert fermé ?
	rl.SetTargetFPS(60) // définit les fps a x

	for !rl.WindowShouldClose() { // boucle tant que c'est pas fermé
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

		rl.EndDrawing()
	}
	type Bouton struct {  // On défini un nouveau type appelé Bouton
		Texte     string   
		Rectangle rl.Rectangle  // Un rectangle pour définir la position et la taille du Bouton
		Couleur   rl.Color   // La couleur du Bouton
	}
	
	var jouerBouton := Button{
		Texte:     "JOUER",
		Rectangle: rl.NewRectangle(300, 200, 200, 50),
		Couleur:   rl.Blue,
	}
	var optionsBouton := Button{
		Texte:     "OPTIONS",
		Rectangle: rl.NewRectangle(300, 300, 200, 50),
		Couleur:   rl.Green,
	}
	
	quitterBouton := Button{
		Texte:     "QUITTER",
		Rectangle: rl.NewRectangle(300, 400, 200, 50),
		Couleur:   rl.Red,
	}
	}


	