package main



type window struct {
	width int32
	heigh int32
	title string
}

func main() {

	var game window // défini game a la struct window
	game.Init(1366, 768," Donkey Kong Mario") // initialise la fenetre : donc attribue les valeurs automatiquement ?
	game.initGame()
	
}


