package main

import (
	"qtcore"
)

func main() {
	argv := []string{"./coreapp", "-v", "-x"}
	qtcore.NewQCoreApplication(len(argv), argv, 0)
}
