package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	fichierBanniere = "standard.txt"
	hauteurLigne    = 8
	premiereRune    = 32  // ' '
	derniereRune    = 126 // '~'
	nombreRunes     = derniereRune - premiereRune + 1
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	texte := os.Args[1]
	if texte == "" {
		return
	}

	texte = strings.ReplaceAll(texte, `\n`, "\n")

	banniere, err := chargerBanniere(fichierBanniere)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error loading banner:", err)
		return
	}

	lignes := strings.Split(texte, "\n")
	afficherAscii(lignes, banniere)
}

func chargerBanniere(chemin string) (map[rune][]string, error) {
	donnees, err := os.ReadFile(chemin)
	if err != nil {
		return nil, err
	}

	lignes := strings.Split(string(donnees), "\n")
	banniere := make(map[rune][]string)

	// Dans standard.txt : bloc de 9 lignes par rune
	// ligne 0 vide, lignes 1..8 = dessin
	for i := 0; i < nombreRunes; i++ {
		debut := i*(hauteurLigne+1) + 1
		fin := debut + hauteurLigne
		if fin > len(lignes) {
			break
		}
		r := rune(premiereRune + i)
		banniere[r] = lignes[debut:fin]
	}
	return banniere, nil
}

func afficherAscii(lignes []string, banniere map[rune][]string) {
	for i, ligne := range lignes {
		if ligne == "" {
			if i != len(lignes)-1 {
				fmt.Println()
			}
			continue
		}

		for rang := 0; rang < hauteurLigne; rang++ {
			var buf strings.Builder
			for _, r := range ligne {
				if r < premiereRune || r > derniereRune {
					continue
				}
				bloc, ok := banniere[r]
				if !ok || len(bloc) <= rang {
					continue
				}
				buf.WriteString(bloc[rang])
			}
			fmt.Println(buf.String())
		}
	}
}
