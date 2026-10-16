package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: go run . <in> <out>")
		os.Exit(1)
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			fmt.Fprintln(os.Stderr, "Erreur : fichier introuvable :", os.Args[1])
		} else {
			fmt.Fprintln(os.Stderr, "Erreur :", err)
		}
		os.Exit(1)
	}

	for _, word := range strings.Fields(string(data)) {
		fmt.Println(word)
	}
}