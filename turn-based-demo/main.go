package main

import (
	"fmt"
	"math/rand"
	"time"
)

func selectPlayerAbility() Ability {
	for {
		fmt.Println("Choose an option:")

		for index, ability := range player.abilities {
			fmt.Printf("%v. %v\n", index, ability.name)
		}

		var abilityId int
		fmt.Scan(&abilityId)

		if abilityId < 0 || abilityId > len(player.abilities) {
			fmt.Println("Invalid ability choice")
			continue
		}

		ability := player.abilities[abilityId]
		fmt.Printf("You selected %v\n", ability.name)

		return ability
	}
}

/*
	Passing by reference is "cheaper"
	With objects, it's the only way to mutate the passed object
*/
func applyAbility(caster *Character, target *Character, ability Ability) {
	abilityValue := rand.Intn(ability.maxValue-ability.minValue+1) + ability.minValue
	castMessage := ability.getCastMessage(caster.name, target.name, abilityValue)

	fmt.Println(castMessage)
	target.health = target.health - abilityValue

	if target.health < 0 {
		target.health = 0
	}
}

func main() {
	fmt.Println("Let the battle commence!")
	announceStats()

	for {
		rand.Seed(time.Now().UnixNano())

		selectedAbility := selectPlayerAbility()

		applyAbility(&player, &enemy, selectedAbility)
		announceStats()
		if enemy.health <= 0 {
			fmt.Printf("%v dies!\n", enemy.name)
			break
		}

		applyAbility(&enemy, &player, enemy.abilities[0])
		announceStats()
		if player.health <= 0 {
			fmt.Printf("%v dies!\n", player.name)
			break
		}
	}
}
