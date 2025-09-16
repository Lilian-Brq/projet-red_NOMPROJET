package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// init perso
	c1 := initCharacter("", "", "", "", 0, 0, 0, []string{"Potion"}, [5]string{""}, 100)
	// creation du perso
	characterCreation(&c1)
	// Menu Home
	Interface(&c1)

}

type character struct {
	name         string
	sexe         string
	classe       string
	race         string
	niv          int
	pv_max       int
	pv_act       int
	inventaire   []string
	skill        [5]string
	money        int
	equipment    equipment
	maxInventory int
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
func initCharacter(nom, sexe, classe, race string, niveau, pv_max, pv_act int, inventaire []string, skill [5]string, monnaie int) character {
	return character{
		name:         nom,
		sexe:         sexe,
		classe:       classe,
		race:         race,
		niv:          niveau,
		pv_max:       pv_max,
		pv_act:       pv_act,
		inventaire:   inventaire,
		skill:        skill,
		money:        monnaie,
		maxInventory: 10,
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
	if len(c.inventaire) >= c.maxInventory {
		fmt.Println("⚠️ Inventaire plein, impossible d’ajouter l’objet.")
		return
	}
	c.inventaire = append(c.inventaire, item)
	fmt.Printf("✅ %s ajouté à l'inventaire !\n", item)
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
			fmt.Printf("Vous utilisez Potion de soin")
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
				fmt.Println("\n 222. Equipe un vêtement ")
				fmt.Println("\n 666. Agrandir sac ")
				fmt.Println("\n 777. Utiliser Livre de sort")
				fmt.Println("\n 999. Utiliser Potion ")
				fmt.Println(" 1. Marchand ")
				fmt.Println(" 0. Retour")
				fmt.Println(" Votre choix")
				fmt.Scan(&new_choice)
				switch new_choice {
				case 222:
					fmt.Println("Selon votre inventaire : Vous avez : 1. Chapeau de paille ; 2. Salopette; 3. Bottes de pluie")
					fmt.Scan(&new_choice)
					switch new_choice {
					case 1:
						for i := 0; i < len(c.inventaire); i++ {
							if c.inventaire[i] == "Chapeau de paille" {
								equipItem(c, "Bottes de pluie")
								removeInventory(c, "Chapeau de paille")

							}
						}
					case 2:
						for i := 0; i < len(c.inventaire); i++ {
							if c.inventaire[i] == "Salopette" {
								equipItem(c, "Bottes de pluie")
								removeInventory(c, "Salopette")

							}
						}
					case 3:
						for i := 0; i < len(c.inventaire); i++ {
							if c.inventaire[i] == "Bottes de pluie" {
								equipItem(c, "Bottes de pluie")
								removeInventory(c, "Bottes de pluie")

							}
						}
					default:
						fmt.Println("\n Choix Invalide, Veuillez réessayer")
					}

				case 999:
					takePot(c)
				case 666:
					fmt.Println("Vous avez : 1. Petit sac; 2. Moyen sac; 3. Grand sac")
					fmt.Scan(&new_choice)
					switch new_choice {
					case 1:
						for i := 0; i < len(c.inventaire); i++ {
							if c.inventaire[i] == "Petit sac" {
								upgradeInventorySlot(c, "Petit sac")
								removeInventory(c, "Petit sac")

							}
						}
					case 2:
						for i := 0; i < len(c.inventaire); i++ {
							if c.inventaire[i] == "Moyen sac" {
								upgradeInventorySlot(c, "Moyen sac")
								removeInventory(c, "Moyen sac")

							}
						}
					case 3:
						for i := 0; i < len(c.inventaire); i++ {
							if c.inventaire[i] == "Grand sac" {
								upgradeInventorySlot(c, "Grand sac")
								removeInventory(c, "Grand sac")

							}
						}
					default:
						fmt.Println("\n Choix Invalide, Veuillez réessayer")
					}
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
					default:
						fmt.Println("\n Choix Invalide, Veuillez réessayer")
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
								fmt.Println(" 10. Petit sac (+5) : 10 rubis")
								fmt.Println(" 11. Moyen sac (+10) : 20 rubis")
								fmt.Println(" 12. Grand sac (+15) : 30 rubis")
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
								case 10:
									if len(c.inventaire) >= c.maxInventory {
										fmt.Println("⚠️ Inventaire plein, impossible d’ajouter l’objet.")
										break
									}
									purchase(c, "Petit sac")
								case 11:
									if len(c.inventaire) >= c.maxInventory {
										fmt.Println("⚠️ Inventaire plein, impossible d’ajouter l’objet.")
										break
									}
									purchase(c, "Moyen sac")
								case 12:
									if len(c.inventaire) >= c.maxInventory {
										fmt.Println("⚠️ Inventaire plein, impossible d’ajouter l’objet.")
										break
									}
									purchase(c, "Grand sac")
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
				if new_choice == 0 {
					break
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
						fmt.Println(" 10. Petit sac (+5) : 10 rubis")
						fmt.Println(" 11. Moyen sac (+10) : 20 rubis")
						fmt.Println(" 12. Grand sac (+15) : 30 rubis")
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
						case 10:
							if len(c.inventaire) >= c.maxInventory {
								fmt.Println("⚠️ Inventaire plein, impossible d’ajouter l’objet.")
								break
							}
							purchase(c, "Petit sac")
						case 11:
							if len(c.inventaire) >= c.maxInventory {
								fmt.Println("⚠️ Inventaire plein, impossible d’ajouter l’objet.")
								break
							}
							purchase(c, "Moyen sac")
						case 12:
							if len(c.inventaire) >= c.maxInventory {
								fmt.Println("⚠️ Inventaire plein, impossible d’ajouter l’objet.")
								break
							}
							purchase(c, "Grand sac")
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
func Wasted(c *character, g *monster) {
	if c.pv_act <= 0 {
		fmt.Printf(" %s est mort (;;)\n", c.name)
		fmt.Printf("\n               ⸜(｡˃ ᵕ ˂ )⸝♡ \n \n %s est récusité avec 50pv \n", c.name)
		c.pv_act = 50
	}
	if g.pv_act <= 0 {
		fmt.Printf(" %s est mort (;;)\n", g.name)
	}
}

// Poison Potion
func poisonPot(c *character, g *monster) {
	if g == nil {
		fmt.Println("⚠️ Impossible d’utiliser une potion de poison sans ennemi.")
		return
	}

	for i := 0; i < len(c.inventaire); i++ {
		if c.inventaire[i] == "Potion de poison" {
			// Consomme la potion
			removeInventory(c, "Potion de poison")
			// boucle timer 3sec
			fmt.Printf("Vous utilisez Potion de poison")
			for i := 0; i < 3; i++ {
				time.Sleep(1 * time.Second)
				g.pv_act -= 10
				fmt.Printf("Dégâts de poison ! PV : %d / %d\n", g.pv_act, g.pv_max)
				//si pv = 0 le perso meurt
				if g.pv_act == 0 {
					Wasted(c, g)
					break
				}
				fmt.Printf("%s à infligé 30 dégâts à %s ", c.name, g.name)
				fmt.Printf("%s à %d / %d PV", g.name, g.pv_act, g.pv_max)
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
	"Pomme":                        1,
	"Petit Sac":                    5,
	"Moyen Sac":                    10,
	"Grand Sac":                    15,
}

var selling = map[string]int{
	"Potion":                       2,
	"Potion de poison":             4,
	"Livre de sort : Boule de feu": 18,
	"Fourrure de Loup":             3,
	"Peau de Troll":                5,
	"Cuir de Sanglier":             2,
	"Plume de Corbeau":             1,
	"Pomme":                        1,
}

// Function for purchasing items
func purchase(c *character, item string) {
	price, ok := prices[item]
	if !ok {
		// item gratuit
		price = 0
	}

	if len(c.inventaire) >= c.maxInventory {
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
			c.niv = 1
			c.skill[0] = "Coup de Poing"
			fmt.Printf("Vous êtes un %s avec %d / %d PV (Niveau %d)\n", c.race, c.pv_act, c.pv_max, c.niv)
			return
		case 2:
			c.race = "Elfe"
			c.pv_max = 80
			c.pv_act = 40
			c.niv = 1
			c.skill[0] = "Coup de Poing"
			fmt.Printf("Vous êtes un %s avec %d / %d PV (Niveau %d)\n", c.race, c.pv_act, c.pv_max, c.niv)
			return
		case 3:
			c.race = "Nain"
			c.pv_max = 120
			c.pv_act = 60
			c.niv = 1
			c.skill[0] = "Coup de Poing"
			fmt.Printf("Vous êtes un %s avec %d / %d PV (Niveau %d)\n", c.race, c.pv_act, c.pv_max, c.niv)
			return
		default:
			fmt.Println("\n Choix Invalide, Veuillez réessayer")
			continue
		}
	}
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
func containsInventory(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func equipItem(c *character, item string) {
	switch item {
	case "Chapeau de paille":
		if c.equipment.head[0] != "" { //Si un casque est déjà équipé il le remet dans l'inventaire
			addInventory(c, c.equipment.head[0])
			fmt.Printf("Vous avez enlever votre %s \n", c.equipment.head[0])
		}
		//equipe le casque
		c.equipment.head[0] = item
		removeInventory(c, item)
		c.pv_max += 10
		fmt.Println("✅ Vous avez équipé votre chapeau de paille (+10 PV)")

	case "salopette":
		if c.equipment.chest[0] != "" { //Si la tenue est déjà équipée il le remet dans l'inventaire
			addInventory(c, c.equipment.chest[0])
			fmt.Printf("Vous avez enlever votre %s \n", c.equipment.chest[0])
		}
		//equipe la salopette
		c.equipment.chest[0] = item
		removeInventory(c, item)
		c.pv_max += 25
		fmt.Println("✅ Vous avez équipé votre salopette (+25 PV)")

	case "Bottes de pluie":
		if c.equipment.legs[0] != "" { //Si les bottes sont déjà équipées il les remet dans l'inventaire
			addInventory(c, c.equipment.legs[0])
			fmt.Printf("Vous avez enlever votre %s \n", c.equipment.legs[0])
		}
		//equipe les bottes
		c.equipment.legs[0] = item
		removeInventory(c, item)
		c.pv_max += 15
		fmt.Println("✅ Vous avez équipé vos bottes (+15 PV)")

	default:
		fmt.Println("⚠️ Cet objet ne peut pas être équipé.")
	}
}

// Goblin
func initGoblin() monster {
	return monster{name: "Gobelin d'entraînement", niv: 1, pv_max: 40, pv_act: 40, attack: 5}
}

// Inventory ugrapding with bag
func upgradeInventorySlot(c *character, size string) {
	switch size {
	case "Petit sac":
		c.maxInventory += 5
		fmt.Println("Vous avez utilisé un Petit sac ( +5 places ) ")
	case "Moyen sac":
		c.maxInventory += 10
		fmt.Println("Vous avez utilisé un Moyen sac ( +10 places ) ")
	case "Grand sac":
		c.maxInventory += 15
		fmt.Println("Vous avez utilisé un Grand sac ( +15 places ) ")
	default:
		fmt.Println(" Type de sac inconnu.")
	}
}

// Goblin attack pattern
func goblinPattern(c *character, g *monster, turn int) {
	dmg := g.attack
	if turn%3 == 0 {
		dmg *= 2
	}
	c.pv_act -= dmg
	fmt.Printf("%s inflige à %s %d dégâts ! (%d/%d PV)\n", g.name, c.name, dmg, c.pv_act, c.pv_max)
	Wasted(c, g)
}

// Training
func trainingFight(c *character) {
	g := initGoblin()
	turn := 1
	for c.pv_act > 0 && g.pv_act > 0 {
		fmt.Printf("\n--- Tour %d ---\n", turn)
		characterTurn(c, &g)
		if g.pv_act <= 0 {
			fmt.Println("Victoire ! Le goblin est vaincu.")
			break
		}
		goblinPattern(c, &g, turn)
		if c.pv_act <= 0 {
			fmt.Println("Défaite ! Vous avez été vaincu.")
			break
		}
		turn++
	}
}

func characterTurn(c *character, g *monster) {
	for {
		fmt.Println(" \n ☠  MENU COMBAT ☠ \n ")
		fmt.Println("1. Menu")
		fmt.Println("2. Attaquer")
		fmt.Println("3. Inventaire ")
		fmt.Println("4. QUITTER")

		var choice int
		fmt.Println("Votre choix")
		fmt.Scan(&choice)
		var new_choice int

		switch choice {
		case 1:
			//menu
		case 2:
			var attack int
			fmt.Println("Quelle attaque souhaitez vous utiliser ?")
			for i, s := range c.skill {
				if s != "" {
					fmt.Printf("%d. %s\n", i+1, s)
				}
			}
			fmt.Scan(&attack)

			switch attack {
			case 1:
				fmt.Printf("%s", c.skill[0])
				g.pv_act = g.pv_act - 8
			case 2:
				fmt.Printf("%s", c.skill[1])
				g.pv_act = g.pv_act - 18
			case 3:
				fmt.Printf("%s", c.skill[2])
				g.pv_act = g.pv_act - 0
			case 4:
				fmt.Printf("%s", c.skill[3])
				g.pv_act = g.pv_act - 0
			case 5:
				fmt.Printf("%s", c.skill[4])
				g.pv_act = g.pv_act - 0
			}
		case 3:
			for {
				fmt.Println("\n ✧ Inventaire ✧ \n ")
				accessInventory(*c)
				fmt.Println("\n 999. Utiliser Potion ")
				fmt.Println("\n 111. Utiliser Potion de poison ")
				fmt.Println(" 0. Retour")
				fmt.Println(" Votre choix")
				fmt.Scan(&new_choice)
				switch new_choice {
				case 999:
					takePot(c)
				case 111:
					poisonPot(c, g)
				}
				if new_choice == 0 {
					break
				}
			}
		}
	}
}
