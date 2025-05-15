package ui

import (
        "bufio"
        "fmt"
        "hacksim/game"
        "os"
        "time"

        "github.com/charmbracelet/lipgloss"
)

// SimpleTUI provides a simpler TUI for environments where Bubbletea doesn't work well
type SimpleTUI struct {
        gameState *game.State
        running   bool
        scanner   *bufio.Scanner
}

// NewSimpleTUI creates a new SimpleTUI instance
func NewSimpleTUI(gameState *game.State) *SimpleTUI {
        return &SimpleTUI{
                gameState: gameState,
                running:   true,
                scanner:   bufio.NewScanner(os.Stdin),
        }
}

// Run starts the simple TUI
func (tui *SimpleTUI) Run() {
        // Display splash screen
        tui.displaySplashScreen()
        
        // Display game header
        tui.displayHeader()
        
        // Main game loop
        for tui.running {
                tui.displayPrompt()
                if tui.scanner.Scan() {
                        cmd := tui.scanner.Text()
                        if cmd == "exit" || cmd == "quit" {
                                tui.running = false
                                continue
                        }
                        
                        // Process command and display output
                        output := tui.gameState.ProcessCommand(cmd)
                        fmt.Println(output)
                        
                        // Check objectives
                        if tui.gameState.CheckObjectiveCompletion(cmd) {
                                completedStyle := lipgloss.NewStyle().
                                        Foreground(lipgloss.Color(ColorMatrixGreen)).
                                        Bold(true)
                                fmt.Println(completedStyle.Render("✓ Objective completed!"))
                                
                                // Display progress
                                fmt.Printf("Progress: %.0f%%\n", tui.gameState.Progress*100)
                                
                                // Check if all objectives are completed
                                if tui.gameState.Progress >= 1.0 {
                                        tui.displayMissionComplete()
                                }
                        }
                }
        }
        
        // Display exit message
        fmt.Println("Exiting HackSim. Thanks for playing!")
}

// displaySplashScreen shows the game splash screen
func (tui *SimpleTUI) displaySplashScreen() {
        // ASCII Art logo
        logo := `
  _    _          _____ _  _______ _____ __  __ 
 | |  | |   /\   / ____| |/ / ____|_   _|  \/  |
 | |__| |  /  \ | |    | ' / (___   | | | \  / |
 |  __  | / /\ \| |    |  < \___ \  | | | |\/| |
 | |  | |/ ____ \ |____| . \____) |_| |_| |  | |
 |_|  |_/_/    \_\_____|_|\_\_____/|_____|_|  |_|
`
        
        logoStyle := lipgloss.NewStyle().
                Foreground(lipgloss.Color(ColorMatrixGreen)).
                Bold(true)
        
        fmt.Println(logoStyle.Render(logo))
        
        // Subtitle
        subtitleStyle := lipgloss.NewStyle().
                Foreground(lipgloss.Color(ColorSoftGreen))
        fmt.Println(subtitleStyle.Render("A Terminal Hacking Simulator"))
        fmt.Println()
        
        // Loading animation
        fmt.Print("Initializing cyber infiltration protocols")
        for i := 0; i < 5; i++ {
                time.Sleep(300 * time.Millisecond)
                fmt.Print(".")
        }
        fmt.Println(" DONE")
        
        time.Sleep(1 * time.Second)
        fmt.Print("\033[H\033[2J") // Clear screen
}

// displayHeader shows the game header
func (tui *SimpleTUI) displayHeader() {
        // Display scenario information
        titleStyle := lipgloss.NewStyle().
                Foreground(lipgloss.Color(ColorMatrixGreen)).
                Bold(true)
        
        infoStyle := lipgloss.NewStyle().
                Foreground(lipgloss.Color(ColorSoftGreen))
        
        fmt.Println(titleStyle.Render("MISSION: " + tui.gameState.CurrentScenario.Name))
        fmt.Println(infoStyle.Render(tui.gameState.CurrentScenario.Description))
        fmt.Println()
        
        // Display objectives
        fmt.Println(titleStyle.Render("OBJECTIVES:"))
        for i, obj := range tui.gameState.CurrentScenario.Objectives {
                prefix := "[ ]"
                if obj.Completed {
                        prefix = "[✓]"
                }
                fmt.Printf("%s %d. %s\n", prefix, i+1, obj.Description)
        }
        fmt.Println()
}

// displayPrompt shows the command prompt
func (tui *SimpleTUI) displayPrompt() {
        promptStyle := lipgloss.NewStyle().
                Foreground(lipgloss.Color(ColorMatrixGreen)).
                Bold(true)
        
        // Show connected node in prompt if applicable
        prompt := "hacksim$ "
        if tui.gameState.Network.CurrentNode != "" {
                prompt = fmt.Sprintf("%s@%s$ ", "user", tui.gameState.Network.CurrentNode)
        }
        
        fmt.Print(promptStyle.Render(prompt))
}

// displayMissionComplete shows mission completion screen
func (tui *SimpleTUI) displayMissionComplete() {
        successStyle := lipgloss.NewStyle().
                Foreground(lipgloss.Color(ColorMatrixGreen)).
                Bold(true)
        
        fmt.Println("\n" + successStyle.Render("MISSION COMPLETE!"))
        fmt.Println(successStyle.Render("All objectives have been accomplished."))
        fmt.Println()
        
        // Display stats
        fmt.Printf("Security Level: %d/10\n", tui.gameState.SecurityLevel)
        fmt.Printf("Commands Used: %d\n", tui.gameState.CommandsIssued)
        
        // Ask to continue or exit
        fmt.Println("\nType 'exit' to quit or any other command to continue exploring.")
}