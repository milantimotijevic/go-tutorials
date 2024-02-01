package main

import (
	"flag"
	"fmt"
	"os"
)

var print = fmt.Println

func main() {
	print("-- Sub-line Commands Example --")
	// think of them as commands that can be tailed by their own flags

	respecCmd := flag.NewFlagSet("respec", flag.ExitOnError)
	respecArea := respecCmd.String("area", "", "Character area to respec")

	vendorCmd := flag.NewFlagSet("vendor", flag.ExitOnError)
	direction := vendorCmd.String("direction", "purchase", "Vendor operation direction")
	barter := vendorCmd.Bool("barter", false, "Enable bartering feature")

	if len(os.Args) < 2 {
		print("You must either select 'respec' or 'vendor'")
		os.Exit(1)
	}

	// we will read these as args
	switch os.Args[1] {
	case "respec":
		respecCmd.Parse(os.Args[2:])
		print("You have selected the respec service.")
		print("You are respeccing:", *respecArea)
	case "vendor":
		vendorCmd.Parse(os.Args[2:])
		print("You have selected the vendor service.")
		fmt.Printf("You wish to %v with bartering enabled: %v", *direction, *barter)
	default:
		print("No (valid) args passed")
		os.Exit(1)
	}
}
