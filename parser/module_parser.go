package parser

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ParseGoMod() string {
	// read module name (for imports)
	mod, err := os.Open("go.mod")

	if err != nil {
		log.Fatal("Unable to read go.mod file") // refactor this to not fail
	}

	scanner := bufio.NewScanner(mod)
	var modName string

	scanner.Scan()
	fullModLine := scanner.Text()
	modSplit := strings.Fields(fullModLine)
	modName = modSplit[1]

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return modName
}
