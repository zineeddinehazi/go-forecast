package pkg

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func titleCase(s string) string {
	return strings.Title(strings.ToLower(s))
}

func GetUserInput() string {
	// Replace fmt.Scanln to get the whole string with white spaces
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter the city name to get the weather forecast.\n(for example: Algiers, Paris, New York): ")
	cityBytes, err := reader.ReadBytes('\n') // Read full line
	if err != nil {
		log.Fatal(err)
	}
	city := strings.TrimSpace(string(cityBytes)) // Convert read line into string
	titleCaseCity := titleCase(city)
	return titleCaseCity
}
