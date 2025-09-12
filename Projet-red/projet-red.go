package main

import (
	"fmt"
)

func main() {
	// init perso
	c1 := initCharacter("Jeffrey Dahmer", "Homme", "Mage", "Humain", 1, 100, 40, [10]string{"épee", "ppp", "zz", "zz"}) //Ne pas oublier de remplir l'inventaire

	// Affiche les infos du perso avec displayInfo
	displayInfo(c1)

	//Affiche l'inventaire
	accessInventory(c1)

	// utilise Potion de vie
	takePot(&c1)

	// supprime item dans l'inventaire

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
	fmt.Println("Nom:", c.name)
	fmt.Println("Sexe:", c.sexe)
	fmt.Println("Classe:", c.classe)
	fmt.Println("Race:", c.race)
	fmt.Println("Niveau:", c.niv)
	fmt.Printf("PV: %d / %d\n", c.pv_act, c.pv_max)
}

func accessInventory(c character) {
	for i := 0; i < len(c.inventaire); i++ {
		fmt.Printf("--Inventaire--:\n %d. %s\n", i+1, c.inventaire[i])
	}
}

// focntion qui permet de retirer un item dans l'inventaire
func removeItem(inventory [10]string, item string) []string {
	newInventory := [10]string{}
	remove := false
	for i := 0; i < len(inventory); i++ {
		if !remove && inventory[i] == item {
			remove = true
			continue
		}
		newInventory = append(newInventory, inventory[i])
	}
	return newInventory
}
func takePot(c *character) {
	vide := false
	for i := 0; i < len(c.inventaire); i++ {
		if c.inventaire[i] == "Potion" {
			vide := true

		}
	}

}
