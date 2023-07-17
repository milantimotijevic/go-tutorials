package main

import (
	"encoding/xml"
	"fmt"
)

var print = fmt.Println

type Creature struct {
	Specie string   `xml:"specie"`
	Age    int      `xml:"age"`
	Sounds []string `xml:"sounds"`
	// in my example it worked even without a reference, but who knows...
	Abilities []*Ability `xml:"abilities>ability"` // could chain these as much as I want apparently
	Alignment string     `xml:"alignment,attr"`    // alignment will be set as attr on Creature
}

type Ability struct {
	Name           string `xml:"name"`
	SpellSlotLevel int    `xml:"spellSlotLevel"`
}

func main() {
	print("-- Go XML --")

	telekinesis := Ability{"Telekinesis", 4}
	telepathy := Ability{"Telepathy", 4}
	mindFlayer := Creature{
		Specie:    "Mind Flayer",
		Age:       75,
		Sounds:    []string{"Khhsshhhhh", "Skreeeeeee", "Tzzzruummgghhbbh"},
		Abilities: []*Ability{&telekinesis, &telepathy},
		Alignment: "Lawful Evil",
	}
	// can also call .Marshal(mindFlayer) but the result won't be formatted
	mindflayerXml, _ := xml.MarshalIndent(mindFlayer, " ", "  ")
	print(string(mindflayerXml))

	print(xml.Header + string(mindflayerXml)) // add generic xml header

	// unmarshal xml to struct
	var creature Creature
	xml.Unmarshal(mindflayerXml, &creature) // needs a reference

	print(creature)

	mindflayerXml2, _ := xml.MarshalIndent(&creature, " ", "  ")
	print(string(mindflayerXml2))
}
