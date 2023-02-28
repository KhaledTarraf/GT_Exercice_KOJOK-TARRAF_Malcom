package main

import (
	"fmt"
)

func main() {
	var allumettes int
	var joueur int
	var tour int

	// Demande le nombre d'allumettes au joueur
	fmt.Println("Combien d'allumettes souhaitez-vous utiliser (minimum 4) ?")
	fmt.Scan(&allumettes)

	// Vérifie que le nombre d'allumettes est supérieur ou égal à 4
	for allumettes < 4 {
		fmt.Println("Le nombre d'allumettes doit être supérieur ou égal à 4.")
		fmt.Scan(&allumettes)
	}

	// Affiche le nombre d'allumettes initial
	fmt.Println("Il y a ", allumettes, " allumettes")

	// Initialise le joueur et le tour
	joueur = 1
	tour = 1

	// Boucle jusqu'à ce qu'il ne reste plus d'allumettes
	for allumettes > 0 {
		// Affiche le tour actuel
		fmt.Println("Tour ", tour)

		// Demande au joueur combien d'allumettes il souhaite retirer
		fmt.Println("Joueur ", joueur, ", combien d'allumettes voulez-vous retirer (1, 2 ou 3) ?")
		var choix int
		fmt.Scan(&choix)

		// Vérifie que le choix est valide
		for choix < 1 || choix > 3 || choix > allumettes {
			fmt.Println("Choix invalide.")
			fmt.Println("Joueur ", joueur, ", combien d'allumettes voulez-vous retirer (1, 2 ou 3) ?")
			fmt.Scan(&choix)
		}

		// Retire le nombre d'allumettes choisi
		allumettes -= choix

		// Affiche le nombre d'allumettes restantes
		fmt.Println("Il reste ", allumettes, " allumettes")

		// Change de joueur
		if joueur == 1 {
			joueur = 2
		} else {
			joueur = 1
		}

		// Incrémente le tour
		tour++
	}

	// Affiche le joueur gagnant
	fmt.Println("Joueur ", joueur, " a perdu !")
}
