package main

import (
	"fmt"
)

func main() {
	// init perso
	c1 := initCharacter("Jeffrey Dahmer", "Homme", "Mage", "Humain", 1, 100, 40, [10]string{"épee", "ppp", "zz", "zz", "Potion"}) //Ne pas oublier de remplir l'inventaire

	// Affiche les infos du perso avec displayInfo
	displayInfo(c1)

	//Affiche l'inventaire
	accessInventory(c1)

	// utilise Potion de vie
	takePot(&c1)

	// Menu Home
	Interface(&c1)

}

type character struct {
	name       string
	sexe       string
	classe     string
	race       string
	niv        int
	pv_max     int
	pv_act     int
	inventaire [10]string
}

func initCharacter(nom, sexe, classe, race string, niveau, pv_max, pv_act int, inventaire [10]string) character {
	return character{
		name:       nom,
		sexe:       sexe,
		classe:     classe,
		race:       race,
		niv:        niveau,
		pv_max:     pv_max,
		pv_act:     pv_act,
		inventaire: inventaire,
	}
}

func displayInfo(c character) {
	fmt.Println("「 ✦ 𝐍𝐚𝐦𝐞 ✦ 」", c.name)
	fmt.Println("♀/♂", c.sexe)
	fmt.Println("Classe:", c.classe)
	fmt.Println("Race:", c.race)
	fmt.Println("Niveau:", c.niv)
	fmt.Printf("PV: %d / %d\n", c.pv_act, c.pv_max)
}

func accessInventory(c character) {
	fmt.Printf("--Inventaire--: \n")
	for i := 0; i < len(c.inventaire); i++ {
		fmt.Printf("%d. %s\n", i+1, c.inventaire[i])
	}
}

func takePot(c *character) {
	for i := 0; i < len(c.inventaire); i++ {
		if c.inventaire[i] == "Potion" {
			// Consomme la potion
			c.inventaire[i] = ""

			// Soigne le joueur
			c.pv_act += 50
			if c.pv_act > c.pv_max {
				c.pv_act = c.pv_max
			}

			fmt.Printf("%s utilise une Potion ! PV : %d / %d\n", c.name, c.pv_act, c.pv_max)
			return // on s’arrête après avoir utilisé 1 potion
		}
	}
	// Si pas trouvé
	fmt.Println("Vous n'avez pas de Potion dans l'inventaire !")
}

func Interface(c *character) {
	for {
		fmt.Println(" \n ❀  MENU PRINCIPAL ❀ ")
		fmt.Println("1. Information du Personnage")
		fmt.Println("2. Inventaire")
		fmt.Println("3. QUITTER")

		var choice int
		fmt.Println("Où voulez-vous allez")
		fmt.Scanln(&choice)

		var new_choice int

		switch choice {
		case 1:
			for {
				fmt.Println("\n --Information du Personnage--")
				displayInfo(*c)
				fmt.Println(" 0. Retour")
				fmt.Println(" Votre choix")
				fmt.Scanln(&new_choice)
				if new_choice == 0 {
					break
				}
			}
		case 2:
			for {
				fmt.Println("\n ✧ Inventaire ✧ ")
				accessInventory(*c)
				fmt.Println(" 0. Retour")
				fmt.Println(" Votre choix")
				fmt.Scanln(&new_choice)
				if new_choice == 0 {
					break
				}
			}
		case 3:
			fmt.Println(" 👋 Au revoir 👋")
			return // quitte le programme
		default:
			fmt.Println(" Choix invalide, Veuillez réessayer")
		}
	}
}
