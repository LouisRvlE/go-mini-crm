package main

import (
	"flag"
)

func main() {
	addFlag := flag.Bool("add", false, "Ajouter un contact directement")
	nameFlag := flag.String("name", "", "Nom du contact à ajouter")
	emailFlag := flag.String("email", "", "Email du contact à ajouter")
	flag.Parse()

	store := NewMemoryStore()

	if *addFlag {
		handleDirectAdd(store, *nameFlag, *emailFlag)
	}

	runInteractiveMode(store)
}
