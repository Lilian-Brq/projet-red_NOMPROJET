package main

import (
	"fmt"
)

type character struct {
	name       string
	sexe       string
	classe     string
	race       string
	niv        int
	pv_max     int
	pv_act     int
	inventaire []string
}

func initCharacter(nom, sexe, classe, race string, niveau, pv_max, pv_act int, inventaire []string) character {
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

func main() {
	// init perso
	c1 := initCharacter("Jeffrey Dahmer", "humain", 1, 100, 100, 3)

	// affiche info du perso
	fmt.Println("Nom:", c1.name)
	fmt.Println("Nom:", c1.classe)
	fmt.Println("Nom:", c1.sexe)
	fmt.Println("Nom:", c1.race)
	fmt.Println("Nom:", c1.pv_max)
	fmt.Println("Nom:", c1.pv_act)
	fmt.Println("Nom:", c1.inventaire)

}
