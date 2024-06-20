package main

import "mygame/pkg/game"

func main() {
	game := game.NewGame("dictionary.txt")
	game.Start()
}
