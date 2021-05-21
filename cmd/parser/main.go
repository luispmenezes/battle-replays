package main

import (
	BattleReplays "github.com/luispmenezes/BattleReplays/pkg"
	"log"
	"os"
)

func main() {
	f, err := os.Open("210511-182739.clientreplay")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = BattleReplays.NewParser(f)
	if err != nil {
		log.Fatal(err)
	}
}
