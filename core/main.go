package main



type window struct {
	width int32
	heigh int32
	isClose bool
	title string


}

func main() {

	var game window // défini game a la struct window
	game.Init(1366, 768, true," Donkey Kong Mario") // initialise la fenetre : donc attribue les valeurs automatiquement ?
	var menu window
	menu.Init(1366, 768, true, "Donkey Kong Mario")
	Play(game) // envoie la fenêtre initialisé ?

}


