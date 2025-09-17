package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	// Création du personnage
	c1 := initCharacter("", "", "", "", 0, 0, 0, []string{"Potion de Soin"}, [5]string{}, 100)
	characterCreation(&c1)

	// Menu principal
	Interface(&c1)
}

// =======================
// Structures
// =======================
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
	initiative   int
	exp          int
	expMax       int
	mana         int
	mana_max     int
}

type equipment struct {
	head  [1]string
	chest [1]string
	legs  [1]string
}

type monster struct {
	name       string
	niv        int
	pv_max     int
	pv_act     int
	attack     int
	initiative int
	expReward  int
}

type monsterEntry struct {
	name      string
	weight    int
	fightFunc func(*character)
}

// =======================
// Initialisation perso
// =======================
func initCharacter(nom, sexe, classe, race string, niveau, pv_max, pv_act int, inventaire []string, skill [5]string, monnaie int) character {
	return character{
		name:         nom,
		sexe:         sexe,
		race:         race,
		niv:          niveau,
		pv_max:       pv_max,
		pv_act:       pv_act,
		inventaire:   inventaire,
		skill:        skill,
		money:        monnaie,
		maxInventory: 10,
		initiative:   10, // valeur de base
		exp:          0,  // XP actuelle
		expMax:       20, // XP pour lvl up
		mana:         10, // Mana actuel
		mana_max:     10, // Mana max
	}
}

// =======================
// Création personnage
// =======================
func characterCreation(c *character) {
	var name string

	fmt.Println("\nBienvenue dans 🌱🐾")

	fmt.Println(" _______  _______  _______  ______     _______  __    _  ______     _______  ___      _______  _     _  _______ ")
	fmt.Println("|       ||       ||       ||      |   |   _   ||  |  | ||      |   |       ||   |    |   _   || | _ | ||       |")
	fmt.Println("|  _____||    ___||    ___||  _    |  |  |_|  ||   |_| ||  _    |  |       ||   |    |  |_|  || || || ||  _____|")
	fmt.Println("| |_____ |   |___ |   |___ | | |   |  |       ||       || | |   |  |       ||   |    |       ||       || |_____ ")
	fmt.Println("|_____  ||    ___||    ___|| |_|   |  |       ||  _    || |_|   |  |      _||   |___ |       ||       ||_____  |")
	fmt.Println(" _____| ||   |___ |   |___ |       |  |   _   || | |   ||       |  |     |_ |       ||   _   ||   _   | _____| |")
	fmt.Println("|_______||_______||_______||______|   |__| |__||_|  |__||______|   |_______||_______||__| |__||__| |__||_______|\n")

	fmt.Println("Entrez votre Nom : ")
	fmt.Scan(&name)
	name = strings.ToLower(name)
	name = strings.Title(name)
	c.name = name

	var race int
	for {
		fmt.Println("\n𓆩𓆪 Choisissez votre race 𓆩𓆪")
		fmt.Println("1. Humain (100 PV)")
		fmt.Println("2. Elfe (80 PV)")
		fmt.Println("3. Nain (120 PV)")
		fmt.Scan(&race)
		switch race {
		case 1:
			c.race = "Humain"
			c.pv_max = 100
			c.pv_act = 50
			c.niv = 1
			c.skill[0] = "Coup de Poing"
			fmt.Printf("----Vous êtes un %s avec %d / %d PV (Niveau %d)----\n", c.race, c.pv_act, c.pv_max, c.niv)
			return
		case 2:
			c.race = "Elfe"
			c.pv_max = 80
			c.pv_act = 40
			c.niv = 1
			c.skill[0] = "Coup de Poing"
			fmt.Printf("----Vous êtes un %s avec %d / %d PV (Niveau %d)----\n", c.race, c.pv_act, c.pv_max, c.niv)
			return
		case 3:
			c.race = "Nain"
			c.pv_max = 120
			c.pv_act = 60
			c.niv = 1
			c.skill[0] = "Coup de Poing"
			fmt.Printf("----Vous êtes un %s avec %d / %d PV (Niveau %d)----\n", c.race, c.pv_act, c.pv_max, c.niv)
			return
		default:
			fmt.Println("Choix invalide, recommencez.")
			continue
		}
	}
}

// =======================
// Affichage infos
// =======================
func displayInfo(c character) {
	fmt.Println("╔══════════════════════════════════╗")
	fmt.Printf("║ 「 ✦ %s ✦ 」\n", c.name)
	fmt.Println("╟──────────────────────────────────╢")
	fmt.Printf("║ Race    : %s\n", c.race)
	fmt.Printf("║ Niveau  : %d | Exp: %d / %d\n", c.niv, c.exp, c.expMax)
	fmt.Printf("║ PV      : %d / %d\n", c.pv_act, c.pv_max)
	fmt.Printf("║ Mana    : %d / %d\n", c.mana, c.mana_max)
	fmt.Printf("║ Skill   : %v\n", c.skill)
	fmt.Printf("║ Rubis   : %d\n", c.money)
	fmt.Println("╚══════════════════════════════════╝")
}

// =======================
// Inventaire
// =======================
func accessInventory(c character) {
	for i := 0; i < len(c.inventaire); i++ {
		if c.inventaire[i] != "" {
			fmt.Printf("%d. %s\n", i+1, c.inventaire[i])
		}
	}
}

func addInventory(c *character, item string) {
	if len(c.inventaire) >= c.maxInventory {
		fmt.Println("⚠️ Inventaire plein, impossible d'ajouter l'objet.\n")
		return
	}
	c.inventaire = append(c.inventaire, item)
	fmt.Printf("✅ %s ajouté à l'inventaire.\n", item)
}

func removeInventory(c *character, item string) {
	for i := 0; i < len(c.inventaire); i++ {
		if c.inventaire[i] == item {
			c.inventaire[i] = ""
			fmt.Printf("❌ %s retiré de l'inventaire.\n", item)
			return
		}
	}
	fmt.Printf("⚠️ %s n'est pas dans l'inventaire.\n", item)
}

// =======================
// Potions
// =======================
func takePot(c *character) {
	if containsInventory(c.inventaire, "Potion de Soin") {
		removeInventory(c, "Potion de Soin")
		c.pv_act += 50
		if c.pv_act > c.pv_max {
			c.pv_act = c.pv_max
		}
		fmt.Printf(" ♡ %s utilise une Potion de Soin ♡ ! PV : %d/%d \n", c.name, c.pv_act, c.pv_max)
	} else {
		fmt.Println("Vous n'avez pas de Potion de Soin.")
	}
}

func poisonPot(c *character, g *monster) {
	if g == nil {
		fmt.Println("⚠️ Impossible d'utiliser une potion de poison sans ennemi.\n")
		return
	}

	for i := 0; i < len(c.inventaire); i++ {
		if c.inventaire[i] == "Potion de poison" {
			removeInventory(c, "Potion de poison")
			fmt.Printf("Vous utilisez Potion de poison")
			for i := 0; i < 3; i++ {
				time.Sleep(1 * time.Second)
				g.pv_act -= 10
				fmt.Printf(" 🧪 Dégâts de poison 🧪 ! PV : %d / %d\n", g.pv_act, g.pv_max)
				if g.pv_act <= 0 {
					Wasted(c, g)
					break
				}
				fmt.Printf("%s à infligé 30 dégâts à %s ", c.name, g.name)
				fmt.Printf("%s à %d / %d PV", g.name, g.pv_act, g.pv_max)
			}
			return
		}
	}
}

func manaPot(c *character) {
	if containsInventory(c.inventaire, "Potion de mana") {
		removeInventory(c, "Potion de mana")
		c.mana += 10
		if c.mana > c.mana_max {
			c.mana = c.mana_max
		}
		fmt.Printf(" 🔮 %s boit une Potion de mana 🔮 ! Mana : %d/%d\n", c.name, c.mana, c.mana_max)
	} else {
		fmt.Println("Vous n'avez pas de Potion de mana.\n")
	}
}

// =======================
// Mort & résurrection
// =======================
func Wasted(c *character, g *monster) {
	if c.pv_act <= 0 {
		fmt.Printf(" %s est mort (╥﹏╥)\n", c.name)
		fmt.Printf("\n               ⸜(｡˃ ᵕ ˂ )⸝♡ \n \n %s a réssuscité avec 50pv \n", c.name)
		c.pv_act = 50
	}
	if g.pv_act <= 0 {
		fmt.Printf(" %s est mort (╥﹏╥) \n", g.name)
	}
}

// =======================
// Livres de sorts
// =======================
func spellBook(c *character, sort string) {
	for i := 0; i < len(c.skill); i++ {
		if c.skill[i] == sort {
			fmt.Println("⚠️ Sort déjà appris.")
			return
		}
		if c.skill[i] == "" {
			c.skill[i] = sort
			fmt.Printf("✅ Sort appris : %s\n", sort)
			return
		}
	}
	fmt.Println("⚠️ Plus de place pour apprendre un sort.")
}

// =======================
// Combat (Missions)
// =======================

// ------- Slime -------
func initSlime() monster {
	return monster{
		name:       "Slime",
		niv:        3,
		pv_max:     40,
		pv_act:     40,
		attack:     4,
		initiative: 6,
		expReward:  10,
	}
}

func slimePattern(c *character, g *monster, turn int) {
	dmg := g.attack
	if turn%4 == 0 {
		dmg *= 2
	}
	c.pv_act -= dmg
	fmt.Printf(" ⚔️ %s inflige %d dégâts à %s ⚔️ ! (%d/%d PV)\n", g.name, dmg, c.name, c.pv_act, c.pv_max)
	Wasted(c, g)
}

func SlimeFight(c *character) {
	g := initSlime()
	turn := 1
	playerTurn := c.initiative >= g.initiative

	for c.pv_act > 0 && g.pv_act > 0 {
		fmt.Printf("\n--- Tour %d ---\n", turn)
		if playerTurn {
			characterTurn(c, &g)
			if g.pv_act <= 0 {
				fmt.Println(" 🏆 Victoire 🏆 ! Le slime est vaincu.")
				c.exp += g.expReward
				fmt.Printf(" 🖤 Vous gagnez %d XP 🖤 !\n", g.expReward)
				levelUp(c)
				break
			}
			slimePattern(c, &g, turn)
			if c.pv_act <= 0 {
				fmt.Println(" ಥ‿ಥ Défaite ! Vous avez été vaincu. ಥ‿ಥ ")
				break
			}
		} else {
			slimePattern(c, &g, turn)
			if c.pv_act <= 0 {
				fmt.Println(" ಥ‿ಥ Défaite ! Vous avez été vaincu. ಥ‿ಥ ")
				break
			}
			characterTurn(c, &g)
			if g.pv_act <= 0 {
				fmt.Println(" 🏆 Victoire 🏆 ! Le slime est vaincu.")
				c.exp += g.expReward
				fmt.Printf(" 🖤 Vous gagnez %d XP 🖤 !\n", g.expReward)
				levelUp(c)
				break
			}
		}
		turn++
	}
}

// ------ Goblin ------
func initGoblin() monster {
	return monster{
		name:       "Gobelin d'entraînement",
		niv:        9,
		pv_max:     100,
		pv_act:     100,
		attack:     8,
		initiative: 8,
		expReward:  30,
	}
}

func goblinPattern(c *character, g *monster, turn int) {
	dmg := g.attack
	if turn%3 == 0 {
		dmg *= 2
	}
	c.pv_act -= dmg
	fmt.Printf(" ⚔️ %s inflige %d dégâts à %s ⚔️ ! (%d/%d PV)\n", g.name, dmg, c.name, c.pv_act, c.pv_max)
	Wasted(c, g)
}

func GoblinFight(c *character) {
	g := initGoblin()
	turn := 1
	playerTurn := c.initiative >= g.initiative

	for c.pv_act > 0 && g.pv_act > 0 {
		fmt.Printf("\n--- Tour %d ---\n", turn)
		if playerTurn {
			characterTurn(c, &g)
			if g.pv_act <= 0 {
				fmt.Println(" 🏆 Victoire 🏆 ! Le gobelin est vaincu.")
				c.exp += g.expReward
				fmt.Printf(" 🖤 Vous gagnez %d XP 🖤 !\n", g.expReward)
				levelUp(c)
				break
			}
			goblinPattern(c, &g, turn)
			if c.pv_act <= 0 {
				fmt.Println(" ಥ‿ಥ Défaite ! Vous avez été vaincu. ಥ‿ಥ")
				break
			}
		} else {
			goblinPattern(c, &g, turn)
			if c.pv_act <= 0 {
				fmt.Println(" ಥ‿ಥ Défaite ! Vous avez été vaincu. ಥ‿ಥ ")
				break
			}
			characterTurn(c, &g)
			if g.pv_act <= 0 {
				fmt.Println(" 🏆 Victoire 🏆 ! Le gobelin est vaincu.")
				c.exp += g.expReward
				fmt.Printf(" 🖤 Vous gagnez %d XP 🖤 !\n", g.expReward)
				levelUp(c)
				break
			}
		}
		turn++
	}
}

// ------- Loup -------
func initLoup() monster {
	return monster{
		name:       "Slime",
		niv:        13,
		pv_max:     150,
		pv_act:     150,
		attack:     14,
		initiative: 6,
		expReward:  10,
	}
}

func loupPattern(c *character, g *monster, turn int) {
	dmg := g.attack
	if turn%8 == 0 {
		dmg *= 2
	}
	c.pv_act -= dmg
	fmt.Printf(" ⚔️ %s inflige %d dégâts à %s ⚔️ ! (%d/%d PV)\n", g.name, dmg, c.name, c.pv_act, c.pv_max)
	Wasted(c, g)
}

func LoupFight(c *character) {
	g := initLoup()
	turn := 1
	playerTurn := c.initiative >= g.initiative

	for c.pv_act > 0 && g.pv_act > 0 {
		fmt.Printf("\n--- Tour %d ---\n", turn)
		if playerTurn {
			characterTurn(c, &g)
			if g.pv_act <= 0 {
				fmt.Println(" 🏆 Victoire 🏆 ! Le loup est vaincu.")
				c.exp += g.expReward
				fmt.Printf(" 🖤 Vous gagnez %d XP 🖤 !\n", g.expReward)
				levelUp(c)
				break
			}
			loupPattern(c, &g, turn)
			if c.pv_act <= 0 {
				fmt.Println(" ಥ‿ಥ Défaite ! Vous avez été vaincu. ಥ‿ಥ")
				break
			}
		} else {
			loupPattern(c, &g, turn)
			if c.pv_act <= 0 {
				fmt.Println(" ಥ‿ಥ Défaite ! Vous avez été vaincu. ಥ‿ಥ ")
				break
			}
			characterTurn(c, &g)
			if g.pv_act <= 0 {
				fmt.Println(" 🏆 Victoire 🏆 ! Le loup est vaincu.")
				c.exp += g.expReward
				fmt.Printf(" 🖤 Vous gagnez %d XP 🖤 !\n", g.expReward)
				levelUp(c)
				break
			}
		}
		turn++
	}
}

// ------- Troll -------
func initTroll() monster {
	return monster{
		name:       "Slime",
		niv:        25,
		pv_max:     250,
		pv_act:     250,
		attack:     20,
		initiative: 10,
		expReward:  100,
	}
}

func trollPattern(c *character, g *monster, turn int) {
	dmg := g.attack
	if turn%10 == 0 {
		dmg *= 2
	}
	c.pv_act -= dmg
	fmt.Printf(" ⚔️ %s inflige %d dégâts à %s ⚔️ ! (%d/%d PV)\n", g.name, dmg, c.name, c.pv_act, c.pv_max)
	Wasted(c, g)
}

func TrollFight(c *character) {
	g := initTroll()
	turn := 1
	playerTurn := c.initiative >= g.initiative

	for c.pv_act > 0 && g.pv_act > 0 {
		fmt.Printf("\n--- Tour %d ---\n", turn)
		if playerTurn {
			characterTurn(c, &g)
			if g.pv_act <= 0 {
				fmt.Println(" 🏆 Victoire 🏆 ! Le troll est vaincu.")
				c.exp += g.expReward
				fmt.Printf(" 🖤 Vous gagnez %d XP 🖤 !\n", g.expReward)
				levelUp(c)
				break
			}
			trollPattern(c, &g, turn)
			if c.pv_act <= 0 {
				fmt.Println(" ಥ‿ಥ Défaite ಥ‿ಥ ! Vous avez été vaincu.")
				break
			}
		} else {
			trollPattern(c, &g, turn)
			if c.pv_act <= 0 {
				fmt.Println(" ಥ‿ಥ Défaite ಥ‿ಥ ! Vous avez été vaincu.")
				break
			}
			characterTurn(c, &g)
			if g.pv_act <= 0 {
				fmt.Println(" 🏆 Victoire 🏆 ! Le troll est vaincu.")
				c.exp += g.expReward
				fmt.Printf(" 🖤 Vous gagnez %d XP 🖤 !\n", g.expReward)
				levelUp(c)
				break
			}
		}
		turn++
	}
}

// ----- Combat -------
func characterTurn(c *character, g *monster) {
	for {
		fmt.Println("╔════════════════════════════╗")
		fmt.Println("║      ☠  MENU COMBAT ☠      ║")
		fmt.Println("╟────────────────────────────╢")
		fmt.Println("║ 1. Attaquer                ║")
		fmt.Println("║ 2. Inventaire              ║")
		fmt.Println("║ 3. Fuir                    ║")
		fmt.Println("╚════════════════════════════╝")

		var choice int
		fmt.Println("Votre choix")
		fmt.Scan(&choice)
		var new_choice int

		switch choice {
		case 1:
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
				fmt.Printf("%s\n", c.skill[0])
				g.pv_act = g.pv_act - 8
				return
			case 2:
				if c.skill[1] != "" {
					fmt.Printf("%s\n", c.skill[1])
					g.pv_act = g.pv_act - 18
					return
				}
			case 3:
				if c.skill[2] != "" {
					fmt.Printf("%s\n", c.skill[2])
					g.pv_act = g.pv_act - 25
					return
				}
			case 4:
				if c.skill[3] != "" {
					fmt.Printf("%s\n", c.skill[3])
					g.pv_act = g.pv_act - 30
					return
				}
			case 5:
				if c.skill[4] != "" {
					fmt.Printf("%s\n", c.skill[4])
					g.pv_act = g.pv_act - 35
					return
				}
			default:
				fmt.Println("Choix invalide")
			}
		case 2:
			for {
				fmt.Println("\n ✧ Inventaire ✧ \n ")
				accessInventory(*c)
				fmt.Println("\n 999. Utiliser Potion de Soin")
				fmt.Println("\n 111. Utiliser Potion de poison")
				fmt.Println("\n 222. Utiliser Potion de mana")
				fmt.Println(" 0. Retour")
				fmt.Println(" Votre choix")
				fmt.Scan(&new_choice)
				switch new_choice {
				case 999:
					takePot(c)
					return
				case 111:
					poisonPot(c, g)
					return
				case 222:
					manaPot(c)
					return
				}
				if new_choice == 0 {
					break
				}
			}
		case 3:
			fmt.Println("Fuite du combat !")
			return
		}
	}
}

func exploration(c *character) {
	rand.Seed(time.Now().UnixNano())

	monsters := []monsterEntry{
		{name: "Slime", weight: 60, fightFunc: SlimeFight},
		{name: "Gobelin", weight: 29, fightFunc: GoblinFight},
		{name: "Loup", weight: 9, fightFunc: LoupFight},
		{name: "Troll", weight: 2, fightFunc: TrollFight},
	}

	totalWeight := 0
	for _, m := range monsters {
		totalWeight += m.weight
	}

	roll := rand.Intn(totalWeight)
	sum := 0

	for _, m := range monsters {
		sum += m.weight
		if roll < sum {
			fmt.Printf("\n🌿 Vous explorez les environs...\n⚔️ Un %s apparaît !\n", m.name)
			m.fightFunc(c)
			return
		}
	}
}

// =======================
// Level up
// =======================
func levelUp(c *character) {
	if c.exp >= c.expMax {
		c.niv++
		c.exp -= c.expMax
		c.expMax += 20
		c.pv_max += 10
		c.pv_act = c.pv_max
		c.mana_max += 5
		c.mana = c.mana_max
		c.initiative += 2
		fmt.Printf("🎉 %s passe niveau %d ! PV max +10, Mana max +5, Initiative +2\n", c.name, c.niv)
	}
}

// =======================
// Utilitaires
// =======================
func containsInventory(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

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

// =======================
// Menu principal
// =======================
func Interface(c *character) {
	for {
		fmt.Println("\n ❀ MENU PRINCIPAL ❀ ")
		fmt.Println("1. Informations du personnage")
		fmt.Println("2. Inventaire")
		fmt.Println("3. Marchand")
		fmt.Println("4. Forgeron")
		fmt.Println("5. Exploration")
		fmt.Println("6. Qui sont-ils ?")
		fmt.Println("7. Quitter")

		var new_choice int
		fmt.Scan(&new_choice)
		switch new_choice {
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
				fmt.Println("\n 30. Equipe un vêtement ")
				fmt.Println(" 40. Agrandir sac ")
				fmt.Println(" 50. Utiliser Livre de sort")
				fmt.Println(" 60. Utiliser Potion de Soin")
				fmt.Println(" 70. Utiliser Potion de Mana")
				fmt.Println(" \n 80. Marchand ")
				fmt.Println(" 0. Retour")
				fmt.Println(" Votre choix")
				fmt.Scan(&new_choice)
				switch new_choice {
				case 30:
					fmt.Println("Selon votre inventaire : Vous avez : 1. Chapeau de paille ; 2. Salopette; 3. Bottes de pluie")
					fmt.Scan(&new_choice)
					switch new_choice {
					case 1:
						for i := 0; i < len(c.inventaire); i++ {
							if c.inventaire[i] == "Chapeau de paille" {
								equipItem(c, "Chapeau de paille")
								removeInventory(c, "Chapeau de paille")

							}
						}
					case 2:
						for i := 0; i < len(c.inventaire); i++ {
							if c.inventaire[i] == "Salopette" {
								equipItem(c, "Salopette")
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
				case 60:
					takePot(c)
				case 40:
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
					}
				case 50:
					fmt.Println(" Votre choix parmis :\n 1. Boule de feu\n 2.wigadium\n 3.xxxxxx\n 4.xxxxxx")
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
				case 70:
					manaPot(c)
				case 80:
					merchantMenu(c)
				}
				if new_choice == 0 {
					break
				}
			}
		case 3:
			merchantMenu(c)
		case 4:
			for {
				fmt.Println("\n Forgeron ")
				fmt.Println("⚠️ Tout objet forgé coûte 5 rubis. ⚠️")
				fmt.Println(" 1. Chapeau de paille (nécessite 1 Plume de Corbeau et 1 Paille)")
				fmt.Println(" 2. Salopette (nécessite 2 Fourrure de Loup et 1 Peau de Troll)")
				fmt.Println(" 3. Bottes de pluie (nécessite 1 Fourrure de Loup et 1 Cuir de Sanglier)")
				fmt.Println(" 0. Retour")
				fmt.Println(" Votre choix ?")
				fmt.Scan(&new_choice)

				switch new_choice {
				case 1:
					CraftForgeron(c, "Chapeau de paille")
				case 2:
					CraftForgeron(c, "Salopette")
				case 3:
					CraftForgeron(c, "Bottes de pluie")
				default:
					fmt.Println("⚠️ Choix Invalide, Veuillez réessayer")
				}
				if new_choice == 0 {
					break
				}
			}
		case 5:
			exploration(c)
		case 6:
			fmt.Println("🎶 Réponse Mission 6 : ABBA (Mamma Mia, Gimme! Gimme! Gimme!, etc.)")
			fmt.Println("📖 Réponse Mission 6 : Steven Spielberg (Ready Player One)")
		case 7:
			fmt.Println("👋 Fin du jeu, merci d'avoir joué !")
			return
		}
	}
}

// =======================
// Achat & Vente
// =======================
var prices = map[string]int{
	"Potion de Soin":               3,
	"Potion de poison":             6,
	"Potion de mana":               4,
	"Livre de sort : Boule de feu": 25,
	"Fourrure de Loup":             4,
	"Peau de Troll":                7,
	"Cuir de Sanglier":             3,
	"Plume de Corbeau":             1,
	"Paille":                       1,
	"Pomme":                        1,
	"Petit Sac":                    10,
	"Moyen Sac":                    20,
	"Grand Sac":                    30,
}

var selling = map[string]int{
	"Potion de Soin":               2,
	"Potion de poison":             4,
	"Potion de mana":               3,
	"Livre de sort : Boule de feu": 18,
	"Fourrure de Loup":             3,
	"Peau de Troll":                5,
	"Cuir de Sanglier":             2,
	"Plume de Corbeau":             1,
	"Pomme":                        1,
}

func removeMoney(c *character, money int) {
	c.money -= money
	fmt.Printf("❌ %d rubis ont été retirés de votre bourse.\n", money)
}

func addMoney(c *character, money int) {
	c.money += money
	fmt.Printf("✅ %d rubis ont été ajoutés à votre bourse.\n", money)
}

func purchase(c *character, item string) {
	price, ok := prices[item]
	if !ok {
		// item gratuit
		price = 0
	}

	if len(c.inventaire) >= c.maxInventory {
		fmt.Println("⚠️ Inventaire plein, impossible d'acheter.")
		return
	}

	if c.money < price {
		fmt.Printf("Rubis insuffisant : %d requis, il te manque %d.\n", price, price-c.money)
		return
	}

	// Débite l'argent et ajoute l'objet
	removeMoney(c, price)
	addInventory(c, item)
	if price > 0 {
		fmt.Printf(" -%d Rubis | Rubis restant : %d\n", price, c.money)
	}
}

func sell(c *character, item string) {
	price := selling[item]
	if containsInventory(c.inventaire, item) {
		removeInventory(c, item)
		addMoney(c, price)
	} else {
		fmt.Println("⚠️ Vous n'avez pas l'objet sur vous.")
		return
	}
}

func merchantMenu(c *character) {
	for {
		fmt.Println("\n Marchand")
		fmt.Println(" 1. Acheter")
		fmt.Println(" 2. Vendre")
		fmt.Println(" 0. Retour Menu")

		var new_choice int
		fmt.Scan(&new_choice)

		switch new_choice {
		case 1:
			for {
				fmt.Println("\n Marchand \n ")
				fmt.Println(" 1. épée (gratuit)")
				fmt.Println(" 2. Pomme : 1 Rubis")
				fmt.Println(" 3. Paille : 1 Rubis")
				fmt.Println(" 4. Cuir de Sanglier : 3 Rubis")
				fmt.Println(" 5. Plume de Corbeau : 1 Rubis")
				fmt.Println(" 6. Fourure de Loup : 4 Rubis")
				fmt.Println(" 7. Peau de Troll : 7 Rubis")
				fmt.Println(" 8. Potion de Soin: 3 Rubis")
				fmt.Println(" 9. Potion de Poison : 6 Rubis")
				fmt.Println(" 10. Livre de sort -> Boule de feu : 25 Rubis")
				fmt.Println(" 11. Petit Sac (+5) : 10 rubis")
				fmt.Println(" 12. Sac (+10) : 20 rubis")
				fmt.Println(" 13. Grand Sac (+15) : 30 rubis")
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
					purchase(c, "Paille")
				case 4:
					purchase(c, "Cuir de Sanglier")
				case 5:
					purchase(c, "Plume de Corbeau")
				case 6:
					purchase(c, "Fourure de Loup")
				case 7:
					purchase(c, "Peau de Troll")
				case 8:
					purchase(c, "Potion de Soin")
				case 9:
					purchase(c, "Potion de Poison")
				case 10:
					purchase(c, "Livre de sort : Boule de feu")
				case 11:
					purchase(c, "Petit Sac")
				case 12:
					purchase(c, "Sac")
				case 13:
					purchase(c, "Grand Sac")
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
				fmt.Println(" 2. Pomme : 1 Rubis")
				fmt.Println(" 3. Cuir de Sanglier : 2 Rubis")
				fmt.Println(" 4. Plume de Corbeau : 1 Rubis")
				fmt.Println(" 5. Fourure de Loup : 3 Rubis")
				fmt.Println(" 6. Peau de Troll : 5 Rubis")
				fmt.Println(" 7. Potion de Soin : 2 Rubis")
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
					sell(c, "Fourure de Loup")
				case 6:
					sell(c, "Peau de Troll")
				case 7:
					sell(c, "Potion de Soin")
				case 8:
					sell(c, "Potion de Poison")
				case 9:
					sell(c, "Livre de sort : Boule de feu")
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

// =======================
// Forgeron
// =======================
func CraftForgeron(c *character, items string) {
	if items == "Chapeau de paille" {
		if containsInventory(c.inventaire, "Plume de Corbeau") && containsInventory(c.inventaire, "Paille") {
			if c.money >= 5 {
				removeInventory(c, "Plume de Corbeau")
				removeInventory(c, "Paille")
				removeMoney(c, 5)
				addInventory(c, "Chapeau de paille")
				fmt.Println("✅ Le Chapeau de paille a été fabriqué !")
			} else {
				fmt.Println("⚠️ Vous n'avez pas assez de rubis ! ")
			}
		} else {
			fmt.Println("⚠️ Vous n'avez pas les objets qu'il faut.")
		}
	}
	if items == "Salopette" {
		if containsInventory(c.inventaire, "Fourrure de Loup") && containsInventory(c.inventaire, "Peau de Troll") {
			if c.money >= 5 {
				removeInventory(c, "Fourrure de Loup")
				removeInventory(c, "Fourrure de Loup")
				removeInventory(c, "Peau de Troll")
				removeMoney(c, 5)
				addInventory(c, "Salopette")
				fmt.Println("✅ La Salopette a été fabriquée !")
			} else {
				fmt.Println("⚠️ Vous n'avez pas assez de rubis ! ")
			}
		} else {
			fmt.Println("⚠️ Vous n'avez pas les objets qu'il faut.")
		}
	}
	if items == "Bottes de pluie" {
		if containsInventory(c.inventaire, "Fourrure de Loup") && containsInventory(c.inventaire, "Cuir de Sanglier") {
			if c.money >= 5 {
				removeInventory(c, "Fourrure de Loup")
				removeInventory(c, "Cuir de Sanglier")
				removeMoney(c, 5)
				addInventory(c, "Bottes de pluie")
				fmt.Println("✅ Les bottes de pluie ont été fabriquées !")
			} else {
				fmt.Println("⚠️ Vous n'avez pas assez de rubis ! ")
			}
		} else {
			fmt.Println("⚠️ Vous n'avez pas les objets qu'il faut.")
		}
	}
}
