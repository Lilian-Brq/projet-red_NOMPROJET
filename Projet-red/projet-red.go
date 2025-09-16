package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// init perso
	c1 := initCharacter("", "", "", "", 0, 0, 0, [10]string{"Potion"}, [5]string{""}, 100)
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
	money      int
	equipment  equipment
}

type equipment struct {
	head  [1]string
	chest [1]string
	legs  [1]string
}

type monster struct {
	name   string
	niv    int
	pv_max int
	pv_act int
	attack int
}

// Character's init
func initCharacter(nom, sexe, classe, race string, niveau, pv_max, pv_act int, inventaire [10]string, skill [5]string, monnaie int) character {
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
		money:      monnaie,
	}
}

// Show the informations of the character with displayInfo
func displayInfo(c character) {
	fmt.Println("「 ✦ 𝐍𝐚𝐦𝐞 ✦ 」", c.name)
	fmt.Println("♀/♂", c.sexe)
	fmt.Println("Classe:", c.classe)
	fmt.Println("Race:", c.race)
	fmt.Println("Niveau:", c.niv)
	fmt.Println("skill:", c.skill)
	fmt.Printf("PV: %d / %d\n", c.pv_act, c.pv_max)
	fmt.Printf("rubis: %d\n", c.money)
}

// Show the Inventory
func accessInventory(c character) {
	for i := 0; i < len(c.inventaire); i++ {
		if c.inventaire[i] != "" {
			fmt.Printf("%d. %s\n", i+1, c.inventaire[i])
		}
	}
}

// Add an item in the inventory
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

// Remove an item of the inventory
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

// Use a potion
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

// Game Display
func Interface(c *character) {
	for {
		fmt.Println(" \n ❀  MENU PRINCIPAL ❀ \n ")
		fmt.Println("1. Information du Personnage")
		fmt.Println("2. Inventaire")
		fmt.Println("3. Marchand ")
		fmt.Println("4. Forgeron")
		fmt.Println("5. QUITTER")

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
					for {
						fmt.Println("\n Marchand")
						fmt.Println(" 1. Acheter")
						fmt.Println(" 2. Vendre")
						fmt.Println(" 0. Retour Menu")

						fmt.Scan(&new_choice)

						switch new_choice {
						case 1:
							for {
								fmt.Println("\n Marchand \n ")
								fmt.Println(" 1. épée (gratuit)")
								fmt.Println(" 2. Pommes : 1 Rubis")
								fmt.Println(" 3. Cuir de sanglier : 3 Rubis")
								fmt.Println(" 4. Plume de corbeau : 1 Rubis")
								fmt.Println(" 5. Fourure de loup : 4 Rubis")
								fmt.Println(" 6. Peau de Troll : 7 Rubis")
								fmt.Println(" 7. Potion : 3 Rubis")
								fmt.Println(" 8. Potion de poison : 6 Rubis")
								fmt.Println(" 9. Livre de sort -> Boule de feu : 25 Rubis")
								fmt.Println("\n 777. Retour Inventaire")
								fmt.Println("\n 0. Retour Menu")
								fmt.Println("\n Votre choix ?")

								fmt.Scan(&new_choice)

								switch new_choice {
								case 1:
									purchase(c, "épée")
								case 2:
									purchase(c, "Pomme")
								case 3:
									purchase(c, "Cuir de Sanglier")
								case 4:
									purchase(c, "Plume de Corbeau")
								case 5:
									purchase(c, "Fourure de loup")
								case 6:
									purchase(c, "Peau de Troll")
								case 7:
									purchase(c, "Potion")
								case 8:
									purchase(c, "Potion de poison")
								case 9:
									purchase(c, "Livre de sort : boule de feu")
								case 777:
									accessInventory(*c)
								default:
									fmt.Println("\n Choix Invalide, Veuillez réessayer")
								}
								if new_choice == 0 {
									break
								}
							}
						case 2:
							for {
								fmt.Println("\n Marchand \n ")
								fmt.Println(" 1. épée : 0 rubis")
								fmt.Println(" 2. Pommes : 1 Rubis")
								fmt.Println(" 3. Cuir de sanglier : 2 Rubis")
								fmt.Println(" 4. Plume de corbeau : 1 Rubis")
								fmt.Println(" 5. Fourure de loup : 3 Rubis")
								fmt.Println(" 6. Peau de Troll : 5 Rubis")
								fmt.Println(" 7. Potion : 2 Rubis")
								fmt.Println(" 8. Potion de poison : 4 Rubis")
								fmt.Println(" 9. Livre de sort -> Boule de feu : 18 Rubis")
								fmt.Println("\n 777. Retour Inventaire")
								fmt.Println("\n 0. Retour Menu")
								fmt.Println("\n Votre choix ?")

								fmt.Scan(&new_choice)

								switch new_choice {
								case 1:
									sell(c, "épée")
								case 2:
									sell(c, "Pomme")
								case 3:
									sell(c, "Cuir de Sanglier")
								case 4:
									sell(c, "Plume de Corbeau")
								case 5:
									sell(c, "Fourure de loup")
								case 6:
									sell(c, "Peau de Troll")
								case 7:
									sell(c, "Potion")
								case 8:
									sell(c, "Potion de poison")
								case 9:
									sell(c, "Livre de sort : boule de feu")
								case 777:
									accessInventory(*c)
								default:
									fmt.Println("\n Choix Invalide, Veuillez réessayer")
								}
								if new_choice == 0 {
									break
								}
							}
						}
						if new_choice == 0 {
							break
						}
					}
				}
			}
		case 3:
			for {
				fmt.Println("\n Marchand")
				fmt.Println(" 1. Acheter")
				fmt.Println(" 2. Vendre")
				fmt.Println(" 0. Retour Menu")

				fmt.Scan(&new_choice)

				switch new_choice {
				case 1:
					for {
						fmt.Println("\n Marchand \n ")
						fmt.Println(" 1. épée (gratuit)")
						fmt.Println(" 2. Pommes : 1 Rubis")
						fmt.Println(" 3. Cuir de sanglier : 3 Rubis")
						fmt.Println(" 4. Plume de corbeau : 1 Rubis")
						fmt.Println(" 5. Fourure de loup : 4 Rubis")
						fmt.Println(" 6. Peau de Troll : 7 Rubis")
						fmt.Println(" 7. Potion : 3 Rubis")
						fmt.Println(" 8. Potion de poison : 6 Rubis")
						fmt.Println(" 9. Livre de sort -> Boule de feu : 25 Rubis")
						fmt.Println("\n 777. Retour Inventaire")
						fmt.Println("\n 0. Retour Menu")
						fmt.Println("\n Votre choix ?")

						fmt.Scan(&new_choice)

						switch new_choice {
						case 1:
							purchase(c, "épée")
						case 2:
							purchase(c, "Pomme")
						case 3:
							purchase(c, "Cuir de Sanglier")
						case 4:
							purchase(c, "Plume de Corbeau")
						case 5:
							purchase(c, "Fourure de loup")
						case 6:
							purchase(c, "Peau de Troll")
						case 7:
							purchase(c, "Potion")
						case 8:
							purchase(c, "Potion de poison")
						case 9:
							purchase(c, "Livre de sort : boule de feu")
						case 777:
							accessInventory(*c)
						default:
							fmt.Println("\n Choix Invalide, Veuillez réessayer")
						}
						if new_choice == 0 {
							break
						}
					}
				case 2:
					for {
						fmt.Println("\n Marchand \n ")
						fmt.Println(" 1. épée : 0 rubis")
						fmt.Println(" 2. Pommes : 1 Rubis")
						fmt.Println(" 3. Cuir de sanglier : 2 Rubis")
						fmt.Println(" 4. Plume de corbeau : 1 Rubis")
						fmt.Println(" 5. Fourure de loup : 3 Rubis")
						fmt.Println(" 6. Peau de Troll : 5 Rubis")
						fmt.Println(" 7. Potion : 2 Rubis")
						fmt.Println(" 8. Potion de poison : 4 Rubis")
						fmt.Println(" 9. Livre de sort -> Boule de feu : 18 Rubis")
						fmt.Println("\n 777. Retour Inventaire")
						fmt.Println("\n 0. Retour Menu")
						fmt.Println("\n Votre choix ?")

						fmt.Scan(&new_choice)

						switch new_choice {
						case 1:
							sell(c, "épée")
						case 2:
							sell(c, "Pomme")
						case 3:
							sell(c, "Cuir de Sanglier")
						case 4:
							sell(c, "Plume de Corbeau")
						case 5:
							sell(c, "Fourure de loup")
						case 6:
							sell(c, "Peau de Troll")
						case 7:
							sell(c, "Potion")
						case 8:
							sell(c, "Potion de poison")
						case 9:
							sell(c, "Livre de sort : boule de feu")
						case 777:
							accessInventory(*c)
						default:
							fmt.Println("\n Choix Invalide, Veuillez réessayer")
						}
						if new_choice == 0 {
							break
						}
					}
				}
				if new_choice == 0 {
					break
				}
			}
		case 4:
			for {
				fmt.Println("\n Forgeron ")
				fmt.Println("Tout objet forgé coûte 5 rubis.")
				fmt.Println(" 1. Chapeau de l'aventurier")
				fmt.Println(" 2. Tunique de l'aventurier")
				fmt.Println(" 3. Bottes de l'aventurier")
				fmt.Println(" 0. Retour")
				fmt.Println(" Votre choix ?")
				fmt.Scan(&new_choice)

				switch new_choice {
				case 1:
					CraftForgeron(c, "Chapeau de l'aventurier")
				case 2:
					CraftForgeron(c, "Tunique de l'aventurier")
				case 3:
					CraftForgeron(c, "Bottes de l'aventurier")
				default:
					fmt.Println(" Choix Invalide, Veuillez réessayer")
				}
				if new_choice == 0 {
					break
				}
			}

		case 5:
			fmt.Println(" 👋 Au revoir 👋")
			return // quitte le programme
		default:
			fmt.Println(" Choix invalide, Veuillez réessayer")
		}
	}
}

// Death Animation
func Wasted(c *character) {
	if c.pv_act == 0 {
		fmt.Printf(" %s est mort (;;)\n", c.name)
		fmt.Printf("\n               ⸜(｡˃ ᵕ ˂ )⸝♡ \n \n %s est récusité avec 50pv \n", c.name)
		c.pv_act = 50
	}
}

// Poison Potion
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

// Learning Skills
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

var prices = map[string]int{
	"Potion":                       3,
	"Potion de poison":             6,
	"Livre de sort : Boule de feu": 25,
	"Fourrure de Loup":             4,
	"Peau de Troll":                7,
	"Cuir de Sanglier":             3,
	"Plume de Corbeau":             1,
	"Pommes":                       1,
}

var selling = map[string]int{
	"Potion":                       2,
	"Potion de poison":             4,
	"Livre de sort : Boule de feu": 18,
	"Fourrure de Loup":             3,
	"Peau de Troll":                5,
	"Cuir de Sanglier":             2,
	"Plume de Corbeau":             1,
	"Pommes":                       1,
}

// Return if the inventory is full
func inventoryFull(inv [10]string) bool {
	for i := 0; i < len(inv); i++ {
		if inv[i] == "" {
			return false
		}
	}
	return true
}

// Function for purchasing items
func purchase(c *character, item string) {
	price, ok := prices[item]
	if !ok {
		// item gratuit
		price = 0
	}

	if inventoryFull(c.inventaire) {
		fmt.Println("⚠️ Inventaire plein, impossible d’acheter.")
		return
	}

	if c.money < price {
		fmt.Printf("Rubis insuffisant : %d requis, il te manque %d.\n", price, price-c.money)
		return
	}

	// Débite l’argent et ajoute l’objet
	removeMoney(c, price)
	addInventory(c, item)
	if price > 0 {
		fmt.Printf(" -%d Rubis | Rubis restant : %d\n", price, c.money)
	}
}

// To create the character at the beginning of the game
func characterCreation(c *character) {
	var name string
	fmt.Println("Bienvenue dans 🌱🐾 Seed & Claws")
	fmt.Println("Entrez votre Nom : ")
	fmt.Scan(&name)
	name = strings.ToLower(name)
	name = strings.Title(name)
	c.name = name
	fmt.Println("personnage créé avec le nom :", c.name)
	for {
		var race int
		fmt.Printf("Choisissez votre classe de personnage :\n")
		fmt.Printf(" 1. Humain \n 2. Elfe \n 3. Nain \n")
		fmt.Println("Votre choix ?")
		fmt.Scan(&race)
		switch race {
		case 1:
			c.race = "Humain"
			c.pv_max = 100
			c.pv_act = 50
			fmt.Printf("Vous etes un Humain avec %d / %d PV \n", c.pv_act, c.pv_max)
		case 2:
			c.race = "Elfe"
			c.pv_max = 80
			c.pv_act = 40
			fmt.Printf("Vous etes un Elfe avec %d / %d PV \n", c.pv_act, c.pv_max)
		case 3:
			c.race = "Nain"
			c.pv_max = 120
			c.pv_act = 60
			fmt.Printf("Vous etes un Nain avec %d / %d PV \n", c.pv_act, c.pv_max)
		default:
			fmt.Println("\n Choix Invalide, Veuillez réessayer")
		}
		break
	}
	c.niv = 1
	fmt.Printf("Vous etes de niveau 1\n")
	c.skill[0] = "Coup de Poing,"
}

// Remove money of the inventory
func removeMoney(c *character, money int) {
	c.money -= money
	fmt.Printf("❌ %d rubis ont été retirés de votre bourse.\n", money)
}

// Add money in the inventory
func addMoney(c *character, money int) {
	c.money += money
	fmt.Printf("✅ %d rubis ont été ajoutés à votre bourse.\n", money)
}

// Function for selling items
func sell(c *character, item string) {
	price := selling[item]
	if containsInventory(c.inventaire, item) {
		removeInventory(c, item)
		addMoney(c, price)
	} else {
		fmt.Println("Vous n'avez pas l'objet sur vous.")
		return
	}
}

// Blacksmith craft
func CraftForgeron(c *character, items string) {
	if items == "Chapeau de l'aventurier" {
		if containsInventory(c.inventaire, "Plume de Corbeau") && containsInventory(c.inventaire, "Cuir de Sanglier") {
			if c.money >= 5 {
				removeInventory(c, "Plume de Corbeau")
				removeInventory(c, "Cuir de Sanglier")
				removeMoney(c, 5)
				addInventory(c, "Chapeau de l'aventurier")
			} else {
				fmt.Println("Vous n'avez pas assez de rubis ! ")
			}
		} else {
			fmt.Println("Vous n'avez pas les objets qu'il faut.")
		}
	}
	if items == "Tunique de l'aventurier" {
		if containsInventory(c.inventaire, "Fourrure de Loup") && containsInventory(c.inventaire, "Peau de Troll") {
			if c.money >= 5 {
				removeInventory(c, "Fourrure de Loup")
				removeInventory(c, "Fourrure de Loup")
				removeInventory(c, "Peau de Troll")
				removeMoney(c, 5)
				addInventory(c, "Tunique de l'aventurier")
			} else {
				fmt.Println("Vous n'avez pas assez de rubis ! ")
			}
		} else {
			fmt.Println("Vous n'avez pas les objets qu'il faut.")
		}
	}
	if items == "Bottes de l'aventurier" {
		if containsInventory(c.inventaire, "Fourrure de Loup") && containsInventory(c.inventaire, "Cuir de Sanglier") {
			if c.money >= 5 {
				removeInventory(c, "Fourrure de Loup")
				removeInventory(c, "Cuir de Sanglier")
				removeMoney(c, 5)
				addInventory(c, "Bottes de l'aventurier")
			} else {
				fmt.Println("Vous n'avez pas assez de rubis ! ")
			}
		} else {
			fmt.Println("Vous n'avez pas les objets qu'il faut.")
		}
	}
}

// If an item is in the Inventory
func containsInventory(slice [10]string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func initGoblin(nom string, niveau, pv_max, pv_act int, attack int) monster {
	return monster{
		name:   "Goblin",
		niv:    5,
		pv_max: 40,
		pv_act: 40,
		attack: 5,
	}
}
