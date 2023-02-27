package main

import (
	"fmt"
)

func main() {
	var nbAllumettes int
	var nbAllumettesRestantes int
	var nbAllumettesPrises int
	var joueur int
	var joueurGagnant int

	fmt.Println("Combien d'allumettes voulez-vous prendre ? (minimum 4)")
	fmt.Scanln(&nbAllumettes)
	nbAllumettesRestantes = nbAllumettes
	joueur = 1

	for nbAllumettesRestantes > 0 {
		fmt.Println("Il reste", nbAllumettesRestantes, "allumettes")
		fmt.Println("Joueur", joueur, "combien d'allumettes voulez-vous prendre ?")
		fmt.Scanln(&nbAllumettesPrises)
		for nbAllumettesPrises < 1 || nbAllumettesPrises > 3 {
			fmt.Println("Vous devez prendre entre 1 et 3 allumettes")
			fmt.Scanln(&nbAllumettesPrises)
		}
		nbAllumettesRestantes -= nbAllumettesPrises
		if joueur == 1 {
			joueur = 2
		} else {
			joueur = 1
		}
	}
	if joueur == 1 {
		joueurGagnant = 2
	} else {
		joueurGagnant = 1
	}
	fmt.Println("Le joueur", joueurGagnant, "a gagné !")
}
