package cmd

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
)

// helpCmd represents the help command
var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Show detailed game help",
	Long:  `Display detailed help about game mechanics, commands, and available scenarios.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Display ASCII art title
		title := figure.NewFigure("HackSim Help", "", true)
		fmt.Println(title.String())

		fmt.Println("\n=== GAME SCENARIOS ===")
		fmt.Println("network-breach: Break into a corporate network and navigate through their security systems.")
		fmt.Println("data-heist: Extract valuable data from a secured server while avoiding detection.")
		fmt.Println("system-takeover: Gain full control of a critical infrastructure system.")

		fmt.Println("\n=== GAME CONTROLS ===")
		fmt.Println("Navigation: Arrow keys")
		fmt.Println("Select: Enter")
		fmt.Println("Back/Cancel: Esc")
		fmt.Println("Quit: Ctrl+C")
		
		fmt.Println("\n=== GAME MECHANICS ===")
		fmt.Println("1. Navigate through terminal interfaces and complete objectives")
		fmt.Println("2. Solve puzzles by entering correct commands or finding security weaknesses")
		fmt.Println("3. Avoid detection by security systems")
		fmt.Println("4. Complete mission objectives to progress")
		
		fmt.Println("\n=== TIPS ===")
		fmt.Println("- Take time to explore the environment")
		fmt.Println("- Watch for clues in error messages")
		fmt.Println("- Sometimes doing things quietly is better than rushing")
	},
}

func init() {
	rootCmd.AddCommand(helpCmd)
}
