# Projet Red - ETHAN SOLENE LILIAN


# 🌱🐾 Seed & Claws

**Univers : Stardew Valley**

Projet réalisé dans le cadre du **Projet RED**.

---

## 📖 Présentation

Seed & Claws est un mini RPG en ligne de commande développé en **Golang**.

Le joueur crée son personnage, explore, combat, fabrique des équipements et intéragit avec des PNJ marchands, forgerons, alchimistes et plus encore.

Le jeu reprend les mécaniques classiques d'un RPG : 
- Création et personnalisation d'un personnage
- Inventaire et gestion des ressources
- Système économique (achat, vente, craft)
- Combat au tour par tour contre des créatures
- Progression via l'expérience et les niveaux

---

## 🎭 Création du personnage

- Choix du **nom** (normalisation automatique : Majuscule + minuscules)
- Choix de la **race** :
    - Humain -> 100 PV max
    - Elfe -> 80 PV max
    - Nain -> 120 PV max

---

## 🎮 Début du jeu

- Niveau : 1
- Argent : 100 pièces d'or
- PV : 50% PV max
- Inventaire : 10 emplacements (extensibles avec des sacs)
- Skill : Coup de poing
- Equipement : Aucun

---

## 🧑‍🤝‍🧑 PNJ

- **Marchand** : vend/achète des items, des sacs et des potions
- **Forgeron** : craft des armes et des équipements
- **Enchanteur** : enchante des équipements
- **Alchimiste** : fabrique des potions (soins, poison, mana)
- **Professeur** : enseigne des compétences
- **Couturière** : fabrique des vêtements, des équipements en tissus ou en cuir

---

## 📦 Inventaire & économie

- Limite de base : 10 items
- Extensible avec des sacs (+15 max)
- Ressources à collecter/acheter : fourrure de loup, plume de corbeau, cuir de sanglier, peau de troll, etc

### Recettes du forgeron :
- **Chapeau de l'aventurier** : +10 PV max
- **Tunique de l'aventurier** : + 25 PV max
- **Bottes de l'aventurier** : + 15 PV max

---

## ⚔️ Combat

- Tour par tour : joueur <-> monstre
- Menu de combat : 
    - Attaquer (main nue) : 5 dégâts
    - Différents types d'armes (épée, hache, bâton magique, et d'autres encore)

### Sorts :
- Pleins de sorts différents à utiliser dans le combat :
    - Plus ou moins rares
    - Plus ou moins puissants
    - De l'attaque jusqu'au soins

👉 Consommation de mana pour utiliser des sorts.

---

## 🏆 Progression

- Gain d'XP après chaque combat
- Passage de niveau du personnage -> augmentation PV, stats et mana
- Passage de niveau des comptétences
- Excédent d'XP conservé

---

## 📙 Répartion des taches
- Tache 1 (Création personnage) : Lilian Solène
- Tache 2 (Initialisation du personnage) : Lilian Solène
- Tache 3  (Affichage des informations du perso) : Lilian Solène
- Tache 4 (Accès à l’inventaire) : Lilian Solène
- Tache 5 (Potion de vie) : Lilian Solène
- Tache 5 (Création du menu) : Lilian Solène
- Tache 7 (Marchand) : Lilian 
- Tache 8 (Wasted ) : Solène
- Tache 9 (Potion de poisson) : Lilian Solène
- Tache 10 : (Livre de sort) : Lilian Solène
- Tache 11 (Amélioration personnage) : Lilian
- Tache 12(Limite d’inventaire) : Lilian
- Tache 13 (Monnaie) : Solène
- Tache 14 (Monnaie) : Solène
- Tache 15(Forgeron) : Ethan
- Tache 16 (Equipment initialisation) : Ethan
- Tache 17(Equipment) : Solène
- Tache 18 (sac à dos) : Lilian
- Tache 19 (Monstres) : Solène Ethan
- Tache 20  (Ia) : Ethan
- Tache 21  (Interface combat) : Lilian
- Tache 22 (Combat tour à tour) : Ethan
- Taches bonus : Ethan
