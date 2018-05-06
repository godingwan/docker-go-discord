package main

import (
	"flag"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

// variables used for command line parameters
var (
	Email    string
	Password string
)

func init() {
	flag.StringVar(&Email, "e", "", "Account Email")
	flag.StringVar(&Password, "p", "", "Account Password")
	flag.Parse()

	if Email == "" || Password == "" {
		flag.Usage()
		os.Exit(1)
	}
}

func main() {
	// Create a new Discord session using the provided login information.
	dg, err := discordgo.New(Email, Password)
	if err != nil {
		log.Println("Error created Discord session,", err)
		return
	}

	// Print out your token.
	log.Printf("Your authentication token is: \n\n%s\n", dg.Token)
}
