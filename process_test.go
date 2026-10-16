package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestProcess(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"1E (hex) files were added", "30 files were added"},
		{"ff (hex)", "255"},
		{"It has been 10 (bin) years", "It has been 2 years"},
		{"11 (bin) and 101 (bin)", "3 and 5"},

		{"go (up)", "GO"},
		{"GO (low)", "go"},
		{"bridge (cap)", "Bridge"},
		{"hELLO (cap)", "Hello"},
		{"this is so exciting (up, 2)", "this is SO EXCITING"},
		{"I SHOULD STOP SHOUTING (low, 3)", "I should stop shouting"},
		{"hello world (up, 5)", "HELLO WORLD"},
		{"one (up) two (cap)", "ONE Two"},

		{"Hello , world !", "Hello, world!"},
		{"Hello,world", "Hello, world"},
		{"I was thinking ... You were right", "I was thinking... You were right"},
		{"Wait !? ok", "Wait!? ok"},

		{"' hi '", "'hi'"},
		{"He said ' hello there ' loudly", "He said 'hello there' loudly"},
		{"I don't know", "I don't know"},

		{"a apple", "an apple"},
		{"a hour", "an hour"},
		{"A Egg", "An Egg"},
		{"a banana", "a banana"},

		{"1E (hex) files and 10 (bin) cats were ready (up)", "30 files and 2 cats were READY"},
		{"the answer is ff (hex) , right ?", "the answer is 255, right?"},
		{"Wait ' a apple ' !", "Wait 'an apple'!"},

		{"", ""},
		{"Nothing to change here", "Nothing to change here"},
		{"hello    world", "hello world"},
	}

	for _, c := range cases {
		got := Process(c.in)
		if got != c.want {
			t.Errorf("Process(%q)\n  obtenu  : %q\n  attendu : %q", c.in, got, c.want)
		}
	}
}

func TestHexEtBinaire(t *testing.T) {
	for n := 0; n <= 1000; n++ {
		want := strconv.Itoa(n)

		in := fmt.Sprintf("%x (hex)", n)
		if got := Process(in); got != want {
			t.Errorf("Process(%q) = %q, attendu %q", in, got, want)
		}

		in = fmt.Sprintf("%b (bin)", n)
		if got := Process(in); got != want {
			t.Errorf("Process(%q) = %q, attendu %q", in, got, want)
		}
	}
}

func TestCompteur(t *testing.T) {
	words := []string{"alpha", "bravo", "charlie", "delta", "echo"}

	for k := 1; k <= 7; k++ {
		in := strings.Join(words, " ") + fmt.Sprintf(" (up, %d)", k)
		var expected []string
		for i, w := range words {
			if i >= len(words)-k {
				w = strings.ToUpper(w)
			}
			expected = append(expected, w)
		}
		want := strings.Join(expected, " ")

		if got := Process(in); got != want {
			t.Errorf("k=%d : obtenu %q, attendu %q", k, got, want)
		}
	}
}

func TestPonctuation(t *testing.T) {
	for _, p := range []string{".", ",", "!", "?", ":", ";", "...", "!?"} {
		want := "word" + p + " next"
		in := "word " + p + " next"
		if got := Process(in); got != want {
			t.Errorf("Process(%q) = %q, attendu %q", in, got, want)
		}
	}
}

func TestArticles(t *testing.T) {
	for _, letter := range "abcdefghijklmnopqrstuvwxyz" {
		in := "a " + string(letter) + "xx"
		want := in
		if strings.ContainsRune("aeiouh", letter) {
			want = "an " + string(letter) + "xx"
		}
		if got := Process(in); got != want {
			t.Errorf("Process(%q) = %q, attendu %q", in, got, want)
		}
	}
}