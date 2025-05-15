package main

import (
        "fmt"
        "hacksim/game"
)

func main() {
        fmt.Println("=== HACKSIM - TERMINAL HACKING SIMULATOR ===")
        fmt.Println("============================================")
        fmt.Println("This demo shows two different scenarios in the game.\n")
        
        // First scenario demo: Network Breach
        demoScenario("network-breach")
        
        fmt.Println("\n\n=== SECOND SCENARIO DEMO ===\n")
        
        // Second scenario demo: Data Heist
        demoScenario("data-heist")
}

// demoScenario runs a demo of the specified scenario
func demoScenario(scenarioName string) {
        // Create a new game state with the specified scenario
        gameState := game.NewGameState(scenarioName)

        fmt.Println("HackSim - Terminal Hacking Simulator")
        fmt.Println("=====================================")
        fmt.Println("Scenario:", gameState.CurrentScenario.Name)
        fmt.Println("Description:", gameState.CurrentScenario.Description)

        fmt.Println("\nOBJECTIVES:")
        for i, obj := range gameState.CurrentScenario.Objectives {
                fmt.Printf("%d. %s\n", i+1, obj.Description)
        }

        fmt.Println("\nNETWORK NODES:")
        for _, node := range gameState.CurrentScenario.NetworkNodes {
                fmt.Printf("- %s (%s): %s\n", node.Name, node.IP, node.Type)
        }

        fmt.Println("\nFILE SYSTEM:")
        fmt.Println(gameState.FileSystem.ListFiles("/"))

        fmt.Println("\nExample Commands and Outputs:")
        fmt.Println("-----------------------------")

        // Sample commands based on scenario
        var sampleCommands []string
        
        if scenarioName == "network-breach" {
                sampleCommands = []string{
                        "help",
                        "scan",
                        "connect 192.168.1.1",
                        "ls",
                        "cat welcome.txt",
                        "scan 192.168.1.1",
                        "connect 192.168.1.10",
                        "crack admin",
                }
        } else if scenarioName == "data-heist" {
                sampleCommands = []string{
                        "help",
                        "scan",
                        "connect 10.0.1.20",
                        "ls /mnt",
                        "cat security_policy.pdf",
                        "exploit web 10.0.1.20",
                }
        } else {
                sampleCommands = []string{
                        "help",
                        "scan",
                        "status",
                }
        }

        for _, cmd := range sampleCommands {
                fmt.Printf("\n$ %s\n", cmd)
                output := gameState.ProcessCommand(cmd)
                fmt.Println(output)
        }

        fmt.Println("\nGame Stats:")
        fmt.Println("-----------")
        fmt.Printf("Progress: %.0f%%\n", gameState.Progress*100)
        fmt.Printf("Security Level: %d/10\n", gameState.SecurityLevel)
        fmt.Printf("Available Tools: %s\n", gameState.ToolsUnlocked)

        fmt.Println("\nIn the full game, players can interact with a terminal UI to execute commands and complete objectives.")
}