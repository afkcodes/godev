package main

import "fmt"

// weapon type
// axe
// sword
// wooden stick
// knife

type WeaponType int

const (
	Axe WeaponType = iota
	Sword
	WoodenStick
	Knife
)

func (w WeaponType) String() string {
	switch w {
	case Sword:
		return "SWORD"
	case Axe:
		return "AXE"
	default:
		return ""
	}
}

func getDamage(weaponType WeaponType) int {
	switch weaponType {
	case Axe:
		return 100
	case Sword:
		return 90
	case Knife:
		return 40
	case WoodenStick:
		return 10
	default:
		panic("weapon does not exist")
	}
}

func main() {
	fmt.Printf("damage of the weapon (%d) (%d): \n:", Axe, getDamage(Axe))
	fmt.Printf("damage of the weapon (%d) (%d): \n:", Sword, getDamage(Sword))
	fmt.Printf("damage of the weapon (%d) (%d): \n:", WoodenStick, getDamage(WoodenStick))
	fmt.Printf("damage of the weapon (%d) (%d): \n:", Knife, getDamage(Knife))

	fmt.Printf("damage of the weapon (%s) (%d): \n:", Axe, getDamage(Axe))
	fmt.Printf("damage of the weapon (%s) (%d): \n:", Sword, getDamage(Sword))
}
