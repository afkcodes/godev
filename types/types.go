package main

import "fmt"

var (
	floatVar32 float32 = 10.2
	floatVar64 float64 = 10.2
	name       string  = "FOO"
	intVar32   int32   = 1
	intVar64   int64   = 23232
	intVar     int     = -1
	uintVar    uint    = 1
	uintVar32  uint32  = 1
	uintVar64  uint64  = 10
	uintVar8   uint8   = 0x2
	byteVar    byte    = 0x4
	runeVar    rune    = 'a'
)

type Player struct {
	name        string
	health      int
	attackPower float64
}

// func getHealth(player Player) int {
// 	return player.health
// }

func (player Player) getHealth() int {
	return player.health
}

type Weapon string

var weapon Weapon = "AXE"

func getWeapon(Weapon Weapon) Weapon {
	return weapon
}

func main() {
	player := Player{
		name:        "Ashish",
		health:      100,
		attackPower: 90.2,
	}

	fmt.Printf("health %d \n", player.getHealth())

	// map

	// users := map[string]int{} // empty map
	// users := make(map[string]int) // empty map

	users := map[string]int{
		"ashish": 28,
	}

	users["monchan"] = 30
	users["baz"] = 30

	age := users["ashish"]
	bazAge, ok := users["baz"]

	if !ok {
		fmt.Println("baz does not exist in users")
	} else {
		fmt.Println("baz exists in map", bazAge)
	}

	fmt.Printf("age %v  %d\n", age, bazAge)

	for k := range users {
		fmt.Println(k)
	}

	delete(users, "baz")

	// Slices = are dynamic size
	numbers := []int{1, 2, 3, 4, 5}
	otherNumbers := make([]int, 0) // empty slice
	fmt.Println(numbers, otherNumbers)

	numbers = append(numbers, 6, 7, 8)
	fmt.Println(numbers, otherNumbers)

	numArr := [2]int{1, 2}
	// numArr[0] = 1
	// numArr[1] = 2
	fmt.Println(numArr)

}
