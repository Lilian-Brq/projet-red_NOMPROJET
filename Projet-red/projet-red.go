package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// init perso
	c1 := initCharacter("Jeffrey Dahmer", "Homme", "Mage", "Humain", 1, 100, 10, [10]string{"épee", "ppp", "zz", "zz", "Potion", "Potion de poison"}, [5]string{"Coup de poing,"}) //Ne pas oublier de remplir l'inventaire
	// creation du perso
	characterCreation(&c1)
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
	skill      [5]string
}

func initCharacter(nom, sexe, classe, race string, niveau, pv_max, pv_act int, inventaire [10]string, skill [5]string) character {
	return character{
		name:       nom,
		sexe:       sexe,
		classe:     classe,
		race:       race,
		niv:        niveau,
		pv_max:     pv_max,
		pv_act:     pv_act,
		inventaire: inventaire,
		skill:      skill,
	}
}

// Affiche les infos du perso avec displayInfo
func displayInfo(c character) {
	fmt.Println("「 ✦ 𝐍𝐚𝐦𝐞 ✦ 」", c.name)
	fmt.Println("♀/♂", c.sexe)
	fmt.Println("Classe:", c.classe)
	fmt.Println("Race:", c.race)
	fmt.Println("Niveau:", c.niv)
	fmt.Println("skill:", c.skill)
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
				fmt.Println("\n 777. Utiliser Livre de sort")
				fmt.Println("\n 999. Utiliser Potion ")
				fmt.Println("\n 111. Utiliser Potion de poison ")
				fmt.Println(" 1. Marchand ")
				fmt.Println(" 0. Retour")
				fmt.Println(" Votre choix")
				fmt.Scan(&new_choice)
				switch new_choice {
				case 999:
					takePot(c)
				case 111:
					poisonPot(c)
				case 777:
					fmt.Println(" Votre choix parmis : 1. Boule de feu 2.wigadium 3.xxxxxx 4.xxxxxx")
					fmt.Scan(&new_choice)
					switch new_choice {
					case 1:
						for i := 0; i < len(c.inventaire); i++ {
							if c.inventaire[i] == "Livre de sort : Boule de feu" {
								spellBook(c, "Boule de feu")
								removeInventory(c, "Livre de sort : Boule de feu")

							}
						}
					case 2:
						for i := 0; i < len(c.inventaire); i++ {
							if c.inventaire[i] == "Livre de sort : " {
								spellBook(c, "")
								removeInventory(c, "Livre de sort : ")
								break
							}
						}
					}

				case 1:
					fmt.Println("\n Marchand \n ")
					fmt.Println(" 1. épée (gratuit)")
					fmt.Println(" 2. Pommes")
					fmt.Println(" 3. Cuir de sanglier")
					fmt.Println(" 4. Plume de corbeau")
					fmt.Println(" 5. Fourure de loup")
					fmt.Println(" 6. Peau de Troll")
					fmt.Println(" 7. Potion")
					fmt.Println(" 8. Potion de poison")
					fmt.Println(" 9. Livre de sort : Boule de feu")
					fmt.Println("\n 777. Retour Inventaire")
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
					case 8:
						addInventory(c, "Potion de poison")
					case 9:
						addInventory(c, "Livre de sort : boule de feu")
					case 777:
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
				fmt.Println("7. Potion")
				fmt.Println("8. Potion de poison")
				fmt.Println("9. Livre de sort : Boule de feu")
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
				case 7:
					addInventory(c, "Potion")
				case 8:
					addInventory(c, "Potion de poison")
				case 9:
					addInventory(c, "Livre de sort : Boule de feu")
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
		fmt.Printf("\n               ⸜(｡˃ ᵕ ˂ )⸝♡ \n \n %s est récusité avec 50pv \n", c.name)
		c.pv_act = 50
	}
}

func poisonPot(c *character) {
	for i := 0; i < len(c.inventaire); i++ {
		if c.inventaire[i] == "Potion de poison" {
			// Consomme la potion
			removeInventory(c, "Potion de poison")
			// boucle timer 3sec
			for i := 0; i < 3; i++ {
				time.Sleep(1 * time.Second)
				c.pv_act -= 10
				fmt.Printf("Dégâts de poison ! PV : %d / %d\n", c.pv_act, c.pv_max)
				//si pv =0 le perso meurt
				if c.pv_act == 0 {
					Wasted(c)
					break
				}
			}

		}
	}
}

func spellBook(c *character, sort string) {
	for i := 0; i < len(c.skill); i++ {
		if c.skill[i] == sort {
			fmt.Println("⚠️ Compétences déjà apprise.")
			return
		}
		if c.skill[i] == "" { // première case vide
			c.skill[i] = sort
			fmt.Printf("✅ %s Compétence acquise !\n", sort)
			return
		}
	}
	fmt.Println("⚠️ skill plein, impossible d’ajouter d'autres compétences.")
}

func characterCreation(c *character) {
	var name string
	fmt.Println("Entrez votre Nom : ")
	fmt.Scan(&name)
	name = strings.ToLower(name)
	name = strings.Title(name)
	c.name = name
	fmt.Println("personnage créé avec le nom :", c.name)
	for {
		var classe int
		fmt.Printf("Choisissez votre classe de personnage :\n")
		fmt.Printf("1. Humain \n 2. Elfe \n 3. Nain \n")
		fmt.Println("Votre choix ?")
		fmt.Scan(&classe)
		switch classe {
		case 1:
			c.classe = "Humain"
			c.pv_max = 100
			c.pv_act = 50
			fmt.Printf("Vous etes un Humain avec %d / %d PV \n", c.pv_act, c.pv_max)
		case 2:
			c.classe = "Elfe"
			c.pv_max = 80
			c.pv_act = 60
			fmt.Printf("Vous etes un Elfe avec %d / %d PV \n", c.pv_act, c.pv_max)
		case 3:
			c.classe = "Nain"
			c.pv_max = 120
			c.pv_act = 60
			fmt.Printf("Vous etes un Nain avec %d / %d PV \n", c.pv_act, c.pv_max)
		}
		break
	}
}
