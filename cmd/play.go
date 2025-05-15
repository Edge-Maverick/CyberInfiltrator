package cmd

import (
        "fmt"
        "hacksim/game"
        "hacksim/ui"

        tea "github.com/charmbracelet/bubbletea"
        "github.com/spf13/cobra"
)

var scenarioFlag string
var debugModeFlag bool

// playCmd represents the play command
var playCmd = &cobra.Command{
        Use:   "play",
        Short: "Start a new hacking simulation",
        Long: `Start a new hacking simulation game session.
You can select different scenarios to play through.
Available scenarios: network-breach, data-heist, system-takeover.`,
        Run: func(cmd *cobra.Command, args []string) {
                // Initialize game state
                gameState := game.NewGameState(scenarioFlag)
                
                if debugModeFlag {
                        // Run in debug mode
                        fmt.Println("Running HackSim in debug mode")
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
                        
                        fmt.Println("\nExample Commands:")
                        
                        // Run a few sample commands automatically to demonstrate functionality
                        sampleCommands := []string{"help", "scan", "connect 192.168.1.1", "ls", "cat welcome.txt"}
                        
                        for _, cmd := range sampleCommands {
                                fmt.Printf("\n$ %s\n", cmd)
                                output := gameState.ProcessCommand(cmd)
                                fmt.Println(output)
                        }
                        
                        fmt.Println("\nIn standard mode, the game provides an interactive terminal interface.")
                        
                } else {
                        // Start with splash screen in TUI mode
                        splashModel := ui.NewSplashModel(gameState)
                        p := tea.NewProgram(splashModel, tea.WithAltScreen())
                        if _, err := p.Run(); err != nil {
                                cmd.PrintErr("Error running game: ", err)
                        }
                }
        },
}

func init() {
        rootCmd.AddCommand(playCmd)
        
        // Add flags specific to the play command
        playCmd.Flags().StringVarP(&scenarioFlag, "scenario", "s", "network-breach", "Select a game scenario to play")
        playCmd.Flags().BoolVarP(&debugModeFlag, "debug", "d", false, "Run in debug mode (simpler output)")
}
