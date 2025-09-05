package main

import (
	"main/bowling"
	"os"
)

func main() {
	bowling.ScoreRolls(os.Args[1:])
}