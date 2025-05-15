package ui

import (
        "fmt"
        "strings"
        "time"

        "hacksim/game"

        "github.com/charmbracelet/bubbles/textinput"
        "github.com/charmbracelet/bubbletea"
        "github.com/charmbracelet/lipgloss"
)

// DashboardModel represents the main game dashboard
type DashboardModel struct {
        gameState  *game.State
        terminal   TerminalModel
        statusBar  StatusBarModel
        sidePanel  SidePanelModel
        activeView int // 0 = terminal, 1 = side panel
        width      int
        height     int
}

// TerminalModel represents the terminal view
type TerminalModel struct {
        output     []string
        maxLines   int
        input      textinput.Model
        executing  bool
        spinner    string
        lastCmd    string
}

// StatusBarModel represents the status bar
type StatusBarModel struct {
        missionName  string
        progress     float64
        securityLevel string
        alert        string
        alertTimer   int
}

// SidePanelModel represents the side panel
type SidePanelModel struct {
        menuItems   []string
        selectedIdx int
}

// NewDashboardModel creates a new dashboard model
func NewDashboardModel(gameState *game.State) DashboardModel {
        // Initialize terminal
        input := textinput.New()
        input.Placeholder = "Type commands here..."
        input.Focus()
        input.CharLimit = 60
        input.Width = 60
        
        terminal := TerminalModel{
                output:   []string{"Welcome to HackSim Terminal v1.0", "Type 'help' for available commands."},
                maxLines: 15,
                input:    input,
        }
        
        // Initialize status bar
        statusBar := StatusBarModel{
                missionName:   gameState.CurrentScenario.Name,
                progress:      0.0,
                securityLevel: "Low",
        }
        
        // Initialize side panel menu
        sidePanel := SidePanelModel{
                menuItems: []string{
                        "Mission Objectives",
                        "Network Map",
                        "Available Tools",
                        "System Status",
                        "Quit Game",
                },
                selectedIdx: 0,
        }
        
        return DashboardModel{
                gameState:  gameState,
                terminal:   terminal,
                statusBar:  statusBar,
                sidePanel:  sidePanel,
                activeView: 0,
                width:      TerminalWidth,
                height:     TerminalHeight,
        }
}

// Init initializes the dashboard model
func (m DashboardModel) Init() tea.Cmd {
        // Initialize terminal prompt
        return textinput.Blink
}

// Update handles UI updates
func (m DashboardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
        var cmds []tea.Cmd
        
        switch msg := msg.(type) {
        case tea.KeyMsg:
                switch msg.String() {
                case "ctrl+c", "q":
                        return m, tea.Quit
                        
                case "tab":
                        // Toggle between terminal and side panel
                        m.activeView = (m.activeView + 1) % 2
                        
                case "enter":
                        if m.activeView == 0 && !m.terminal.executing {
                                // Process terminal command
                                cmd := m.terminal.input.Value()
                                if cmd != "" {
                                        m.terminal.lastCmd = cmd
                                        m.terminal.executing = true
                                        m.terminal.input.SetValue("")
                                        
                                        return m, tea.Batch(
                                                textinput.Blink,
                                                m.executeCommand(cmd),
                                        )
                                }
                        } else if m.activeView == 1 {
                                // Process side panel selection
                                return m, m.handleMenuSelection()
                        }
                        
                case "up", "down":
                        if m.activeView == 1 {
                                // Navigate side panel menu
                                if msg.String() == "up" {
                                        m.sidePanel.selectedIdx--
                                        if m.sidePanel.selectedIdx < 0 {
                                                m.sidePanel.selectedIdx = len(m.sidePanel.menuItems) - 1
                                        }
                                } else {
                                        m.sidePanel.selectedIdx++
                                        if m.sidePanel.selectedIdx >= len(m.sidePanel.menuItems) {
                                                m.sidePanel.selectedIdx = 0
                                        }
                                }
                        }
                }
                
        case game.CommandOutputMsg:
                // Add command output to terminal
                m.terminal.output = append(m.terminal.output, "$ "+m.terminal.lastCmd)
                outputLines := strings.Split(string(msg), "\n")
                m.terminal.output = append(m.terminal.output, outputLines...)
                
                // Trim output to maxLines
                if len(m.terminal.output) > m.terminal.maxLines {
                        m.terminal.output = m.terminal.output[len(m.terminal.output)-m.terminal.maxLines:]
                }
                
                m.terminal.executing = false
                
                // Update game state based on the command
                if m.gameState.CheckObjectiveCompletion(m.terminal.lastCmd) {
                        m.statusBar.progress = m.gameState.Progress
                        m.showAlert("Objective completed!", 3)
                }
                
                return m, textinput.Blink
                
        case alertTimeoutMsg:
                m.statusBar.alert = ""
                m.statusBar.alertTimer = 0
                
        case tea.WindowSizeMsg:
                m.width = msg.Width
                m.height = msg.Height
        }
        
        // Update terminal input
        if m.activeView == 0 {
                var cmd tea.Cmd
                m.terminal.input, cmd = m.terminal.input.Update(msg)
                cmds = append(cmds, cmd)
        }
        
        return m, tea.Batch(cmds...)
}

// executeCommand processes a terminal command
func (m DashboardModel) executeCommand(cmd string) tea.Cmd {
        return func() tea.Msg {
                // Process the command in the game state
                output := m.gameState.ProcessCommand(cmd)
                
                // Simulate processing time
                time.Sleep(500 * time.Millisecond)
                
                return game.CommandOutputMsg(output)
        }
}

// alertTimeoutMsg is a message for alert timeouts
type alertTimeoutMsg struct{}

// showAlert displays an alert in the status bar for a specified duration
func (m *DashboardModel) showAlert(alert string, seconds int) tea.Cmd {
        m.statusBar.alert = alert
        m.statusBar.alertTimer = seconds
        
        return tea.Tick(time.Duration(seconds)*time.Second, func(time.Time) tea.Msg {
                return alertTimeoutMsg{}
        })
}

// handleMenuSelection processes a menu selection in the side panel
func (m DashboardModel) handleMenuSelection() tea.Cmd {
        switch m.sidePanel.selectedIdx {
        case 0: // Mission Objectives
                objectives := m.gameState.GetObjectives()
                output := "MISSION OBJECTIVES:\n" + strings.Join(objectives, "\n")
                return func() tea.Msg {
                        return game.CommandOutputMsg(output)
                }
                
        case 1: // Network Map
                networkMap := m.gameState.GetNetworkMap()
                return func() tea.Msg {
                        return game.CommandOutputMsg("NETWORK MAP:\n" + networkMap)
                }
                
        case 2: // Available Tools
                tools := m.gameState.GetAvailableTools()
                output := "AVAILABLE TOOLS:\n" + strings.Join(tools, "\n")
                return func() tea.Msg {
                        return game.CommandOutputMsg(output)
                }
                
        case 3: // System Status
                status := m.gameState.GetSystemStatus()
                return func() tea.Msg {
                        return game.CommandOutputMsg("SYSTEM STATUS:\n" + status)
                }
                
        case 4: // Quit Game
                return tea.Quit
        }
        
        return nil
}

// View renders the dashboard
func (m DashboardModel) View() string {
        // Calculate dimensions
        mainWidth := m.width * 3 / 4
        sideWidth := m.width - mainWidth - 3
        termHeight := m.height - 4 // Reserve space for status bar
        
        // Render status bar
        statusBarView := m.renderStatusBar(m.width)
        
        // Render terminal
        terminalView := m.renderTerminal(mainWidth, termHeight)
        
        // Render side panel
        sidePanelView := m.renderSidePanel(sideWidth, termHeight)
        
        // Join terminal and side panel horizontally
        mainContent := lipgloss.JoinHorizontal(
                lipgloss.Top,
                terminalView,
                sidePanelView,
        )
        
        // Join status bar and main content vertically
        return lipgloss.JoinVertical(
                lipgloss.Left,
                statusBarView,
                mainContent,
        )
}

// renderStatusBar creates the status bar view
func (m DashboardModel) renderStatusBar(width int) string {
        statusText := fmt.Sprintf(" MISSION: %s | PROGRESS: %.0f%% | SECURITY: %s ", 
                m.statusBar.missionName, 
                m.statusBar.progress*100,
                m.statusBar.securityLevel,
        )
        
        // If there's an alert, display it
        if m.statusBar.alert != "" {
                alertStyle := lipgloss.NewStyle().
                        Foreground(lipgloss.Color(ColorTermBlack)).
                        Background(lipgloss.Color(ColorAlertRed)).
                        Bold(true).
                        PaddingLeft(1).
                        PaddingRight(1)
                        
                alertText := alertStyle.Render(" " + m.statusBar.alert + " ")
                
                // Add padding to center the alert
                padding := width - lipgloss.Width(statusText) - lipgloss.Width(alertText)
                if padding > 0 {
                        statusText += strings.Repeat(" ", padding) + alertText
                }
        }
        
        return RenderStatusBar(statusText, width)
}

// renderTerminal creates the terminal view
func (m DashboardModel) renderTerminal(width, height int) string {
        // Join output lines
        outputText := strings.Join(m.terminal.output, "\n")
        
        // Create terminal view with output and input
        var content string
        if m.terminal.executing {
                spinner := NewSpinner()
                content = outputText + "\n$ " + m.terminal.lastCmd + " " + spinner.View()
        } else {
                promptStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(ColorMatrixGreen))
                content = outputText + "\n" + promptStyle.Render("$ ") + m.terminal.input.View()
        }
        
        // Apply focus style if terminal is active
        terminalStyle := Styles.Terminal
        if m.activeView == 0 {
                terminalStyle = terminalStyle.BorderForeground(lipgloss.Color(ColorMatrixGreen))
        }
        
        // Set width and height
        terminalStyle = terminalStyle.Width(width - 2).Height(height - 2)
        
        return terminalStyle.Render(content)
}

// renderSidePanel creates the side panel view
func (m DashboardModel) renderSidePanel(width, height int) string {
        // Render menu items
        menuView := RenderList(m.sidePanel.menuItems, m.sidePanel.selectedIdx)
        
        // Add title
        title := Styles.Subtitle.Render("SYSTEM MENU")
        content := lipgloss.JoinVertical(
                lipgloss.Left,
                title,
                "",
                menuView,
        )
        
        // Apply focus style if side panel is active
        panelStyle := Styles.Panel
        if m.activeView == 1 {
                panelStyle = panelStyle.BorderForeground(lipgloss.Color(ColorMatrixGreen))
        }
        
        // Set width and height
        panelStyle = panelStyle.Width(width - 2).Height(height - 2)
        
        return panelStyle.Render(content)
}
