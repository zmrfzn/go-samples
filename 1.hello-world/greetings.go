package main

import (
	"fmt"
	"runtime"
	"time"
)

// Greeting generates a time-based greeting
func generateGreeting() string {
	hour := time.Now().Hour()
	var timeOfDay string

	switch {
	case hour < 12:
		timeOfDay = "morning"
	case hour < 17:
		timeOfDay = "afternoon"
	default:
		timeOfDay = "evening"
	}

	return fmt.Sprintf("Good %s, Gopher!", timeOfDay)
}

// Generate Go-related facts
func getGoFacts() []string {
	return []string{
		"🚀 Go was developed by Google in 2007",
		"🐹 Go's mascot is a gopher",
		"⚡ Go has built-in concurrency support",
		"🔧 Go compiles to native machine code",
		"📦 Go has a powerful standard library",
	}
}

// Print facts with numbering
func printFacts(facts []string) {
	fmt.Println("\nFun Facts About Go:")
	for i, fact := range facts {
		fmt.Printf("%d. %s\n", i+1, fact)
	}
}

// Banner prints a welcome banner
func printBanner() {
	fmt.Println(`
    *************************
    *   Modular Hello World *
    *************************
    `)
}

// Collect and print execution statistics
func printExecutionStats() {
	fmt.Println("\nProgram Execution Stats:")
	fmt.Printf("Executed at: %s\n", time.Now().Format("15:04:05"))
	fmt.Printf("Go Version: %s\n", runtime.Version())
	fmt.Printf("Operating System: %s\n", runtime.GOOS)
	fmt.Printf("Architecture: %s\n", runtime.GOARCH)
}

func main() {
	printBanner()
	printExecutionStats()
}
