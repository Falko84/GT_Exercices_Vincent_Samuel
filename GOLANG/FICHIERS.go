package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var choix int
	fmt.Println("Que voulez-vous faire ?")
	fmt.Println("1. Récupérer tout le texte contenu dans un .txt")
	fmt.Println("2. Ajouter du texte dans ce fichier depuis le go.")
	fmt.Println("3. Supprimer tout le contenu du fichier")
	fmt.Println("4. Remplacer le contenu par du texte donné par l'utilisateur")
	fmt.Scanln(&choix)
	switch choix {
	case 1:
		// Récupérer tout le texte contenu dans un .txt
		fmt.Println("Récupérer tout le texte contenu dans un .txt")
		// Ouverture du fichier
		fichier, err := os.Open("C:\\Users\\julie\\Desktop\\test.txt")
		if err != nil {
			fmt.Println(err)
		}
		// Lecture du fichier
		scanner := bufio.NewScanner(fichier)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		// Fermeture du fichier
		fichier.Close()
	case 2:
		// Ajouter du texte dans ce fichier depuis le go.
		fmt.Println("Ajouter du texte dans ce fichier depuis le go.")
		// Ouverture du fichier
		fichier, err := os.OpenFile("C:\\Users\\julie\\Desktop\\test.txt", os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err)
		}
		// Ajout du texte
		texte := "Ceci est un test."
		if _, err = fichier.WriteString(texte); err != nil {
			fmt.Println(err)
		}
		// Fermeture du fichier
		fichier.Close()
	case 3:
		// Supprimer tout le contenu du fichier
		fmt.Println("Supprimer tout le contenu du fichier")
		// Ouverture du fichier
		fichier, err := os.OpenFile("C:\\Users\\julie\\Desktop\\test.txt", os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Println(err)
		}
		// Fermeture du fichier
		fichier.Close()
	case 4:
		// Remplacer le contenu par du texte donné par l'utilisateur
		fmt.Println("Remplacer le contenu par du texte donné par l'utilisateur")
		// Ouverture du fichier
		fichier, err := os.OpenFile("C:\\Users\\julie\\Desktop\\test.txt", os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Println(err)
		}
		// Remplacement du contenu
		fmt.Println("Entrez le texte que vous voulez mettre dans le fichier :")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		texte := scanner.Text()
		if _, err = fichier.WriteString(texte); err != nil {
			fmt.Println(err)
		}
		// Fermeture du fichier
		fichier.Close()
	}
}
