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
var simpleModeFlag bool

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
                        // Run in debug mode (non-interactive)
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
                        var sampleCommands []string
                        if scenarioFlag == "data-heist" {
                                sampleCommands = []string{"help", "scan", "connect 10.0.1.20", "ls", "cat mission.txt"}
                        } else if scenarioFlag == "system-takeover" {
                                sampleCommands = []string{"help", "scan", "connect 172.16.1.100", "ls", "cat readme.txt"}
                        } else {
                                sampleCommands = []string{"help", "scan", "connect 192.168.1.1", "ls", "cat welcome.txt"}
                        }
                        
                        for _, cmd := range sampleCommands {
                                fmt.Printf("\n$ %s\n", cmd)
                                output := gameState.ProcessCommand(cmd)
                                fmt.Println(output)
                        }
                        
                        fmt.Println("\nIn standard mode, the game provides an interactive terminal interface.")
                        
                } else if simpleModeFlag {
                        // Run in simple TUI mode (works reliably across environments)
                        simpleTUI := ui.NewSimpleTUI(gameState)
                        simpleTUI.Run()
                } else {
                        // Start with full Bubbletea TUI mode
                        fmt.Println("Starting HackSim in full TUI mode...")
                        fmt.Println("If you encounter display issues, try running with --simple flag")
                        fmt.Println("Press Ctrl+C to exit if needed")
                        
                        // Start with splash screen in full TUI mode
                        splashModel := ui.NewSplashModel(gameState)
                        p := tea.NewProgram(splashModel, tea.WithAltScreen())
                        if _, err := p.Run(); err != nil {
                                fmt.Println("Error running TUI:", err)
                                fmt.Println("Falling back to simple mode...")
                                
                                // Fall back to simple mode if Bubbletea fails
                                simpleTUI := ui.NewSimpleTUI(gameState)
                                simpleTUI.Run()
                        }
                }
        },
}

func init() {
        rootCmd.AddCommand(playCmd)
        
        // Add flags specific to the play command
        playCmd.Flags().StringVarP(&scenarioFlag, "scenario", "s", "network-breach", "Select a game scenario to play")
        playCmd.Flags().BoolVarP(&debugModeFlag, "debug", "d", false, "Run in debug mode (non-interactive demo)")
        playCmd.Flags().BoolVarP(&simpleModeFlag, "simple", "", false, "Run in simple mode (reliable terminal interface)")
}
