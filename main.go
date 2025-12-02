// MADE BY AER
// Projet: ascii-art (Zone01) - représentation ASCII d'une chaîne
// Usage: go run . "Hello\nThere"
// Contexte: lit le banner "standard.txt" (hauteur 8) et affiche l'ASCII art.

package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	bannerFile   = "standard.txt" // nom du fichier banner à côté de main.go
	charHeight   = 8              // HAUTEUR EXACTE: 8 lignes par caractère
	firstRune    = 32             // ' '
	lastRune     = 126            // '~'
	numberOfRune = lastRune - firstRune + 1
)

func main() {
	// On attend exactement 1 argument (la chaîne à afficher).
	if len(os.Args) != 2 {
		return
	}

	input := os.Args[1]

	// Chaîne vide => aucune sortie
	if input == "" {
		return
	}

	// Interpréter les "\n" littéraux comme vrais sauts de ligne.
	normalized := strings.ReplaceAll(input, `\n`, "\n")

	// Charger le banner.
	banner, err := loadBanner(bannerFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error loading banner:", err)
		return
	}

	lines := strings.Split(normalized, "\n")
	printAscii(lines, banner)
}

// loadBanner lit le fichier de banner et construit une map rune -> []string (8 lignes).
func loadBanner(path string) (map[rune][]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	rawLines := strings.Split(string(data), "\n")
	banner := make(map[rune][]string)

	// Chaque rune: 9 lignes dans le fichier
	//  - ligne 0 : vide
	//  - lignes 1..8 : dessin (8 lignes)
	for i := 0; i < numberOfRune; i++ {
		start := i*(charHeight+1) + 1 // on saute la première ligne vide
		end := start + charHeight     // on prend exactement 8 lignes de dessin

		if end > len(rawLines) {
			break
		}

		charLines := rawLines[start:end]
		r := rune(firstRune + i)
		banner[r] = charLines
	}

	return banner, nil
}

// printAscii prend les lignes logiques (séparées par \n) et affiche l'ASCII art.
func printAscii(lines []string, banner map[rune][]string) {
	for idx, line := range lines {
		// Ligne vide:
		// - si ce n'est PAS la dernière -> une seule ligne vide
		// - si c'est la dernière -> on ignore
		if line == "" {
			if idx != len(lines)-1 {
				fmt.Println()
			}
			continue
		}

		// Pour chaque des 8 lignes de hauteur
		for row := 0; row < charHeight; row++ {
			var b strings.Builder

			for _, r := range line {
				if r < firstRune || r > lastRune {
					continue
				}
				charLines, ok := banner[r]
				if !ok || len(charLines) <= row {
					continue
				}
				b.WriteString(charLines[row])
			}

			fmt.Println(b.String())
		}
	}
}
