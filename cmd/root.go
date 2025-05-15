package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hacksim",
	Short: "A terminal-based hacking simulation game",
	Long: `
  _    _          _____ _  _______ _____ __  __ 
 | |  | |   /\   / ____| |/ / ____|_   _|  \/  |
 | |__| |  /  \ | |    | ' / (___   | | | \  / |
 |  __  | / /\ \| |    |  < \___ \  | | | |\/| |
 | |  | |/ ____ \ |____| . \____) |_| |_| |  | |
 |_|  |_/_/    \_\_____|_|\_\_____/|_____|_|  |_|
                                                
HackSim - An immersive terminal-based hacking simulation game.
Navigate through virtual systems, execute hacks, and complete missions.
Use 'hacksim help' for more information about available commands.
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Initialize global flags
	rootCmd.PersistentFlags().StringP("theme", "t", "green", "Set color theme (green, blue, red)")
}
