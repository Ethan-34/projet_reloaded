package main 

import ( 
	"fmt"
	"strings"
	"strconv"
	"unicode"
	"os"
	"regexp"
)

if len(os.Args) != 3 {
fmt.Fprintln(os.Stderr, "usage: go run . <in> <out>")
os.Exit(1)
}
input, output := os.Args[1], os.Args[2]

func main() {

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}
	a := strings.Fields(string(data))
	for _, word := range a {
		fmt.Println(word)
	}
}

_, err := os.ReadFile("nope.txt")
fmt.Println(err)
if errors.Is(err, fs.ErrNotExist) {
}