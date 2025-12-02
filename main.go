// Projet: ascii-art (Zone01) - représentation ASCII d'une chaîne
// Usage: go run . "Hello\nThere"

package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	bannerFile   = "standard.txt" // nom du fichier banner à côté de main.go
	charHeight   = 8              // hauteur d'un caractère dans le banner
	firstRune    = 32             // ' '
	lastRune     = 126            // '~'
	numberOfRune = lastRune - firstRune + 1
)

func main() {
	// On attend exactement 1 argument (la chaîne à afficher).
	if len(os.Args) != 2 {
		// Sujet: aucun message d'erreur demandé, juste rien faire.
		return
	}

	input := os.Args[1]

	// Cas particulier: chaîne vide => aucune sortie (voir sujet)
	if input == "" {
		return
	}

	// Interpréter les séquences "\n" comme des vraies nouvelles lignes.
	normalized := strings.ReplaceAll(input, `\n`, "\n")

	// Charger le banner en mémoire.
	banner, err := loadBanner(bannerFile)
	if err != nil {
		// Le sujet ne définit pas l'erreur, mais panic est acceptable ici.
		fmt.Fprintln(os.Stderr, "error loading banner:", err)
		return
	}

	// Découper en lignes logiques.
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

	// Le format attendu est: pour chaque caractère
	// 8 lignes de dessin + 1 ligne vide (séparateur), soit 9 lignes par caractère.
	expectedLines := numberOfRune * (charHeight + 1)
	if len(rawLines) < expectedLines {
		// Certains banners peuvent avoir une ligne vide finale en plus,
		// on ne panique pas, mais on évite l'out of range.
	}

	banner := make(map[rune][]string)

	for i := 0; i < numberOfRune; i++ {
		start := i * (charHeight + 1)
		end := start + charHeight
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
		// - si ce n'est PAS la dernière -> on affiche juste une ligne vide
		//   (correspond à le cas "Hello\n\nThere": ligne vide au milieu).
		// - si c'est la dernière -> on ignore (cas "Hello\n" ou "\n").
		if line == "" {
			if idx != len(lines)-1 {
				fmt.Println()
			}
			continue
		}

		// Pour chaque "ligne verticale" du caractère (0 à 7)
		for row := 0; row < charHeight; row++ {
			var b strings.Builder

			for _, r := range line {
				// Le sujet garantit des caractères ASCII 32..126 + \n.
				// On ignore les caractères hors plage (option simple).
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

		// Ligne vide entre deux blocs (comme dans les exemples du sujet)
		if idx != len(lines)-1 {
			fmt.Println()
		}
	}
}
