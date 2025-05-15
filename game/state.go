package game

import (
        "fmt"
        "strings"
        "time"
)

// State represents the overall game state
type State struct {
        CurrentScenario Scenario
        Network         Network
        FileSystem      FileSystem
        Progress        float64
        SecurityLevel   int // 0-10, higher means more secure
        DetectionChance float64
        ToolsUnlocked   []string
        CommandHistory  []string
        StartTime       time.Time
}

// CommandOutputMsg is a message containing command output
type CommandOutputMsg string

// NewGameState creates a new game state with the specified scenario
func NewGameState(scenarioName string) *State {
        // Initialize scenario
        var scenario Scenario
        switch scenarioName {
        case "data-heist":
                scenario = DataHeistScenario()
        case "system-takeover":
                scenario = SystemTakeoverScenario()
        default:
                scenario = NetworkBreachScenario() // Default scenario
        }
        
        // Initialize network based on scenario
        network := NewNetwork(scenario.NetworkNodes)
        
        // Initialize file system
        fileSystem := NewFileSystem(scenario.FileSystems)
        
        // Create game state
        return &State{
                CurrentScenario: scenario,
                Network:         network,
                FileSystem:      fileSystem,
                Progress:        0.0,
                SecurityLevel:   scenario.InitialSecurityLevel,
                DetectionChance: 0.1,
                ToolsUnlocked:   []string{"scan", "connect", "ls", "cat", "help"},
                CommandHistory:  []string{},
                StartTime:       time.Now(),
        }
}

// ProcessCommand processes a command and returns the output
func (s *State) ProcessCommand(cmd string) string {
        // Add command to history
        s.CommandHistory = append(s.CommandHistory, cmd)
        
        // Parse command and arguments
        parts := strings.Fields(cmd)
        if len(parts) == 0 {
                return "Error: Empty command"
        }
        
        command := parts[0]
        args := parts[1:]
        
        // Check if the tool is unlocked
        if !s.isToolUnlocked(command) && command != "help" {
                return "Error: Command '" + command + "' not found or access denied"
        }
        
        // Execute the command
        switch command {
        case "help":
                return s.helpCommand(args)
        case "scan":
                return s.scanCommand(args)
        case "connect":
                return s.connectCommand(args)
        case "ls":
                return s.lsCommand(args)
        case "cat":
                return s.catCommand(args)
        case "crack":
                return s.crackCommand(args)
        case "exploit":
                return s.exploitCommand(args)
        case "download":
                return s.downloadCommand(args)
        case "upload":
                return s.uploadCommand(args)
        case "status":
                return s.statusCommand(args)
        case "exit":
                return "Disconnected from current session."
        default:
                return "Error: Command '" + command + "' not recognized"
        }
}

// helpCommand displays help for available commands
func (s *State) helpCommand(args []string) string {
        if len(args) > 0 {
                // Help for specific command
                command := args[0]
                switch command {
                case "scan":
                        return "scan - Scan for network nodes\nUsage: scan [ip]"
                case "connect":
                        return "connect - Connect to a network node\nUsage: connect <ip> [port]"
                case "ls":
                        return "ls - List files in current directory\nUsage: ls [directory]"
                case "cat":
                        return "cat - View file contents\nUsage: cat <filename>"
                case "crack":
                        return "crack - Attempt to crack password or encryption\nUsage: crack <target>"
                case "exploit":
                        return "exploit - Exploit a vulnerability\nUsage: exploit <name> [target]"
                case "download":
                        return "download - Download a file\nUsage: download <filename>"
                case "upload":
                        return "upload - Upload a file\nUsage: upload <file> <destination>"
                case "status":
                        return "status - Display current system status\nUsage: status"
                case "exit":
                        return "exit - Disconnect from current session\nUsage: exit"
                default:
                        return "No help available for '" + command + "'"
                }
        }
        
        // General help
        var availableCommands []string
        for _, tool := range s.ToolsUnlocked {
                availableCommands = append(availableCommands, tool)
        }
        
        return "Available commands:\n" + strings.Join(availableCommands, ", ") +
                "\n\nUse 'help <command>' for more information on a specific command."
}

// scanCommand scans for network nodes
func (s *State) scanCommand(args []string) string {
        targetIP := ""
        if len(args) > 0 {
                targetIP = args[0]
        }
        
        return s.Network.Scan(targetIP)
}

// connectCommand connects to a network node
func (s *State) connectCommand(args []string) string {
        if len(args) < 1 {
                return "Error: Missing IP address\nUsage: connect <ip> [port]"
        }
        
        ip := args[0]
        port := "22" // Default port
        if len(args) > 1 {
                port = args[1]
        }
        
        result, success := s.Network.Connect(ip, port)
        
        // If successfully connected to a target node, check for objective completion
        if success {
                s.checkNetworkObjectives(ip)
        }
        
        return result
}

// lsCommand lists files in the current directory
func (s *State) lsCommand(args []string) string {
        path := "."
        if len(args) > 0 {
                path = args[0]
        }
        
        return s.FileSystem.ListFiles(path)
}

// catCommand displays file contents
func (s *State) catCommand(args []string) string {
        if len(args) < 1 {
                return "Error: Missing filename\nUsage: cat <filename>"
        }
        
        filename := args[0]
        content, err := s.FileSystem.ReadFile(filename)
        
        if err != nil {
                return "Error: " + err.Error()
        }
        
        // Check for data-related objectives when reading files
        s.checkDataObjectives(filename, content)
        
        return content
}

// crackCommand attempts to crack a password or encryption
func (s *State) crackCommand(args []string) string {
        if len(args) < 1 {
                return "Error: Missing target\nUsage: crack <target>"
        }
        
        target := args[0]
        
        // Simulate cracking process with some difficulty based on security level
        // For this simulation, always succeed
        success := true
        
        if success {
                s.unlockTool("exploit")
                return "Successfully cracked " + target + "!\nNew tool unlocked: exploit"
        }
        
        return "Failed to crack " + target + ". Try a different approach."
}

// exploitCommand exploits a vulnerability
func (s *State) exploitCommand(args []string) string {
        if len(args) < 1 {
                return "Error: Missing exploit name\nUsage: exploit <name> [target]"
        }
        
        exploitName := args[0]
        target := ""
        if len(args) > 1 {
                target = args[1]
        }
        
        // Check if this exploit succeeds (simplified)
        success := true // For simulation, always succeed
        
        if success {
                s.Progress += 0.2 // Exploits generally advance the mission
                s.unlockTool("download")
                s.unlockTool("upload")
                
                return "Successfully exploited " + exploitName + " on " + target + "!\n" +
                        "Gained elevated access. New tools unlocked: download, upload"
        }
        
        return "Failed to exploit " + exploitName + " on " + target + "."
}

// downloadCommand simulates downloading a file
func (s *State) downloadCommand(args []string) string {
        if len(args) < 1 {
                return "Error: Missing filename\nUsage: download <filename>"
        }
        
        filename := args[0]
        
        // Check if file exists
        content, err := s.FileSystem.ReadFile(filename)
        if err != nil {
                return "Error: " + err.Error()
        }
        
        // Check for data-related objectives when downloading files
        s.checkDataObjectives(filename, content)
        
        return "Successfully downloaded " + filename
}

// uploadCommand simulates uploading a file
func (s *State) uploadCommand(args []string) string {
        if len(args) < 2 {
                return "Error: Missing parameters\nUsage: upload <file> <destination>"
        }
        
        file := args[0]
        destination := args[1]
        
        // Simulate upload success (simplified)
        return "Successfully uploaded " + file + " to " + destination
}

// statusCommand displays the current system status
func (s *State) statusCommand(args []string) string {
        elapsedTime := time.Since(s.StartTime).Round(time.Second)
        
        status := "=== SYSTEM STATUS ===\n"
        status += "Mission: " + s.CurrentScenario.Name + "\n"
        status += "Progress: " + fmt.Sprintf("%.0f%%", s.Progress*100) + "\n"
        status += "Security Level: " + fmt.Sprintf("%d/10", s.SecurityLevel) + "\n"
        status += "Connected Node: " + s.Network.CurrentNode + "\n"
        status += "Session Time: " + elapsedTime.String() + "\n"
        status += "Available Tools: " + strings.Join(s.ToolsUnlocked, ", ")
        
        return status
}

// isToolUnlocked checks if a tool is available to use
func (s *State) isToolUnlocked(tool string) bool {
        for _, t := range s.ToolsUnlocked {
                if t == tool {
                        return true
                }
        }
        return false
}

// unlockTool adds a tool to the unlocked tools list if it's not already there
func (s *State) unlockTool(tool string) {
        if !s.isToolUnlocked(tool) {
                s.ToolsUnlocked = append(s.ToolsUnlocked, tool)
        }
}

// checkNetworkObjectives checks if connecting to this IP fulfills any objectives
func (s *State) checkNetworkObjectives(ip string) {
        for i, objective := range s.CurrentScenario.Objectives {
                if !objective.Completed && objective.Type == "connect" && objective.Target == ip {
                        s.CurrentScenario.Objectives[i].Completed = true
                        s.updateProgress()
                }
        }
}

// checkDataObjectives checks if accessing this file fulfills any objectives
func (s *State) checkDataObjectives(filename, content string) {
        for i, objective := range s.CurrentScenario.Objectives {
                if !objective.Completed && objective.Type == "data" && objective.Target == filename {
                        s.CurrentScenario.Objectives[i].Completed = true
                        s.updateProgress()
                }
        }
}

// CheckObjectiveCompletion checks if a command completes an objective
func (s *State) CheckObjectiveCompletion(cmd string) bool {
        beforeProgress := s.Progress
        
        // Check command-based objectives
        parts := strings.Fields(cmd)
        if len(parts) > 0 {
                command := parts[0]
                
                for i, objective := range s.CurrentScenario.Objectives {
                        if !objective.Completed && objective.Type == "command" && objective.Target == command {
                                s.CurrentScenario.Objectives[i].Completed = true
                                s.updateProgress()
                        }
                }
        }
        
        // Return true if progress changed
        return s.Progress > beforeProgress
}

// updateProgress recalculates the overall mission progress
func (s *State) updateProgress() {
        completed := 0
        for _, obj := range s.CurrentScenario.Objectives {
                if obj.Completed {
                        completed++
                }
        }
        
        s.Progress = float64(completed) / float64(len(s.CurrentScenario.Objectives))
        
        // If progress reached certain thresholds, adjust security level
        if s.Progress >= 0.5 && s.SecurityLevel < 5 {
                s.SecurityLevel = 5
        } else if s.Progress >= 0.75 && s.SecurityLevel < 8 {
                s.SecurityLevel = 8
        }
}

// GetObjectives returns a list of all objectives with their status
func (s *State) GetObjectives() []string {
        var result []string
        
        for _, obj := range s.CurrentScenario.Objectives {
                status := "[ ]"
                if obj.Completed {
                        status = "[✓]"
                }
                
                result = append(result, status+" "+obj.Description)
        }
        
        return result
}

// GetNetworkMap returns a text representation of the network map
func (s *State) GetNetworkMap() string {
        return s.Network.GetNetworkMap()
}

// GetAvailableTools returns a list of available tools with descriptions
func (s *State) GetAvailableTools() []string {
        toolDescriptions := map[string]string{
                "scan":     "Network scanner - Discover nodes on the network",
                "connect":  "Connection utility - Connect to remote systems",
                "ls":       "List files and directories",
                "cat":      "View file contents",
                "crack":    "Password cracking tool",
                "exploit":  "Vulnerability exploitation framework",
                "download": "Download files from remote systems",
                "upload":   "Upload files to remote systems",
                "status":   "Display current system status",
                "help":     "Show help information",
        }
        
        var result []string
        for _, tool := range s.ToolsUnlocked {
                if desc, ok := toolDescriptions[tool]; ok {
                        result = append(result, tool+" - "+desc)
                } else {
                        result = append(result, tool)
                }
        }
        
        return result
}

// GetSystemStatus returns a detailed system status
func (s *State) GetSystemStatus() string {
        status := fmt.Sprintf("Current Node: %s\n", s.Network.CurrentNode)
        status += fmt.Sprintf("Security Level: %d/10\n", s.SecurityLevel)
        status += fmt.Sprintf("Detection Risk: %.1f%%\n", s.DetectionChance*100)
        status += fmt.Sprintf("Mission Progress: %.0f%%\n", s.Progress*100)
        
        return status
}
