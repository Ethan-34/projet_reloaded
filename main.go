package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
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

result := Process(string(data))
err = os.WriteFile(os.Args[2], []byte(result), 0644)
if err != nil {
	fmt.Fprintln(os.Stderr, "Erreur :", err)
	os.Exit(1)
}
}