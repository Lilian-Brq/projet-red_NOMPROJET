package main

import (
	"fmt"
)

func main() {
	// init perso
	c1 := initCharacter("Jeffrey Dahmer", "Homme", "Mage", "Humain", 1, 100, 0, [10]string{"épee", "ppp", "zz", "zz", "Potion"}) //Ne pas oublier de remplir l'inventaire
	Wasted(&c1)
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

// Affiche les infos du perso avec displayInfo
func displayInfo(c character) {
	fmt.Println("「 ✦ 𝐍𝐚𝐦𝐞 ✦ 」", c.name)
	fmt.Println("♀/♂", c.sexe)
	fmt.Println("Classe:", c.classe)
	fmt.Println("Race:", c.race)
	fmt.Println("Niveau:", c.niv)
	fmt.Printf("PV: %d / %d\n", c.pv_act, c.pv_max)
}

//Affiche l'inventaire

func accessInventory(c character) {
	for i := 0; i < len(c.inventaire); i++ {
		if c.inventaire[i] != "" {
			fmt.Printf("%d. %s\n", i+1, c.inventaire[i])
		}
	}
}

// Ajoute un item à l'inventaire
func addInventory(c *character, item string) {
	for i := 0; i < len(c.inventaire); i++ {
		if c.inventaire[i] == "" { // première case vide
			c.inventaire[i] = item
			fmt.Printf("✅ %s ajouté à l'inventaire !\n", item)
			return
		}
	}
	fmt.Println("⚠️ Inventaire plein, impossible d’ajouter l’objet.")
}

// Retirer un item de l’inventaire
func removeInventory(c *character, item string) {
	for i := 0; i < len(c.inventaire); i++ {
		if c.inventaire[i] == item {
			c.inventaire[i] = ""
			fmt.Printf("❌ %s retiré de l'inventaire.\n", item)
			return
		}
	}
	fmt.Printf("⚠️ %s n’est pas dans l’inventaire.\n", item)
}

// utilise Potion de vie
func takePot(c *character) {
	for i := 0; i < len(c.inventaire); i++ {
		if c.inventaire[i] == "Potion" {
			// Consomme la potion
			removeInventory(c, "Potion")

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
		fmt.Println(" \n ❀  MENU PRINCIPAL ❀ \n ")
		fmt.Println("1. Information du Personnage")
		fmt.Println("2. Inventaire")
		fmt.Println("3. Marchand ")
		fmt.Println("4. QUITTER")

		var choice int
		fmt.Println("Votre choix")
		fmt.Scan(&choice)

		var new_choice int

		switch choice {
		case 1:
			for {
				fmt.Println("\n --Information du Personnage--\n ")
				displayInfo(*c)
				fmt.Println(" 0. Retour")
				fmt.Println(" Votre choix")
				fmt.Scan(&new_choice)
				if new_choice == 0 {
					break
				}
			}
		case 2:
			for {
				fmt.Println("\n ✧ Inventaire ✧ \n ")
				accessInventory(*c)

				fmt.Println("\n 9. Utiliser Potion ")
				fmt.Println(" 1. Marchand ")
				fmt.Println(" 0. Retour")
				fmt.Println(" Votre choix")
				fmt.Scan(&new_choice)
				switch new_choice {
				case 9:
					takePot(c)
				case 1:
					fmt.Println("\n Marchand \n ")
					fmt.Println(" 1. épée (gratuit)")
					fmt.Println(" 2. Pommes")
					fmt.Println(" 3. Cuir de sanglier")
					fmt.Println(" 4. Plume de corbeau")
					fmt.Println(" 5. Fourure de loup")
					fmt.Println(" 6. Peau de Troll")
					fmt.Println("\n 9. Retour Inventaire")
					fmt.Println(" 0. Retour Menu")
					fmt.Println(" Votre choix ?")
					fmt.Scan(&new_choice)

					switch new_choice {
					case 1:
						addInventory(c, "épée")
					case 2:
						addInventory(c, "Pomme")
					case 3:
						addInventory(c, "Cuir de Sanglier")
					case 4:
						addInventory(c, "Plume de Corbeau")
					case 5:
						addInventory(c, "Fourure de loup")
					case 6:
						addInventory(c, "Peau de Troll")
					case 7:
						addInventory(c, "Potion")
					case 9:
						accessInventory(*c)
					default:
						fmt.Println(" Choix Invalide, Veuillez réessayer")
					}

				}
				if new_choice == 0 {
					break
				}
			}
		case 3:
			for {
				fmt.Println("\n Marchand ")
				fmt.Println(" 1. épée (gratuit)")
				fmt.Println("2. Pommes")
				fmt.Println(" 3. Cuir de sanglier")
				fmt.Println(" 4. Plume de corbeau")
				fmt.Println(" 5. Fourure de loup")
				fmt.Println("6. Peau de Troll")
				fmt.Println(" 0. Retour")
				fmt.Println(" Votre choix ?")
				fmt.Scan(&new_choice)

				switch new_choice {
				case 1:
					addInventory(c, "épée")
				case 2:
					addInventory(c, "Pomme")
				case 3:
					addInventory(c, "Cuir de Sanglier")
				case 4:
					addInventory(c, "Plume de Corbeau")
				case 5:
					addInventory(c, "Fourure de loup")
				case 6:
					addInventory(c, "Peau de Troll")
				default:
					fmt.Println(" Choix Invalide, Veuillez réessayer")
				}
				if new_choice == 0 {
					break
				}
			}

		case 4:
			fmt.Println(" 👋 Au revoir 👋")
			return // quitte le programme
		default:
			fmt.Println(" Choix invalide, Veuillez réessayer")
		}
	}
}

func Wasted(c *character) {
	if c.pv_act == 0 {
		fmt.Printf(" %s est mort (;;)\n", c.name)
		fmt.Printf("\n               ⸜(｡˃ ᵕ ˂ )⸝♡ \n \n %s est récusité avec 50pv", c.name)
		c.pv_act = 50
	}
}
