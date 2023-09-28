package main



type window struct {
	width int32
	heigh int32
	isClose bool


}

func main() {

	var game window // défini game a la struct window
	game.Init(1366, 768, true) // initialise la fenetre : donc attribue les valeurs automatiquement ?
	var menu window
	menu.Init(1366, 768, true)
	Play(game) // envoie la fenêtre initialisé ?

}


