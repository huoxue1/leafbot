package main

import (
	"github.com/3343780376/leafBot"
	"os"
)

func init() {
	leafBot.UseEchoHandle("commit")
}

func main() {
	dir, _ := os.Getwd()
	leafBot.LoadConfig(dir + "\\example\\config.json")

	leafBot.InitBots()

}
