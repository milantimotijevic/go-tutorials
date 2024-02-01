package main

import (
	"fmt"
)

func extractMapKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

func extractMapValues[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))

	for _, v := range m {
		values = append(values, v)
	}

	return values
}

func main() {
	fmt.Println("-- Go Generic Maps --")

	var companions map[string]string = map[string]string{
		"Gale":        "Wizard",
		"Astarion":    "Rogue",
		"Shadowheart": "Cleric",
	}

	fmt.Println(extractMapKeys(companions))
	fmt.Println(extractMapValues(companions))
}
