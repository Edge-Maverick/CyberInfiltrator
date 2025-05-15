package ui

import (
        "strings"
        "time"

        "hacksim/game"

        "github.com/charmbracelet/bubbles/timer"
        "github.com/charmbracelet/bubbletea"
        "github.com/common-nighthawk/go-figure"
)

// SplashModel represents the splash screen model
type SplashModel struct {
        gameState     *game.State
        timer         timer.Model
        animationDone bool
        currentText   string
        textAnimator  func() (string, bool)
}

// NewSplashModel creates a new splash screen model
func NewSplashModel(gameState *game.State) SplashModel {
        t := timer.NewWithInterval(3*time.Second, 100*time.Millisecond)
        
        flavorText := "Initializing cyber infiltration protocols..."
        textAnimator := AnimateText(flavorText, 50*time.Millisecond)
        
        return SplashModel{
                gameState:    gameState,
                timer:        t,
                textAnimator: textAnimator,
        }
}

// Init initializes the splash model
func (m SplashModel) Init() tea.Cmd {
        return tea.Batch(
                m.timer.Init(),
                m.animateText(),
        )
}

// animateText returns a command that animates the text
func (m SplashModel) animateText() tea.Cmd {
        return tea.Tick(50*time.Millisecond, func(time.Time) tea.Msg {
                return textAnimationTickMsg{}
        })
}

// textAnimationTickMsg is a message for text animation ticks
type textAnimationTickMsg struct{}

// Update handles UI updates
func (m SplashModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
        switch msg := msg.(type) {
        case tea.KeyMsg:
                switch msg.String() {
                case "ctrl+c", "q":
                        return m, tea.Quit
                
                // Skip splash on any key
                case "enter", " ":
                        return NewDashboardModel(m.gameState), nil
                }
                
        case timer.TickMsg, timer.TimeoutMsg:
                var cmd tea.Cmd
                m.timer, cmd = m.timer.Update(msg)
                
                // When timer finishes, transition to dashboard
                if _, ok := msg.(timer.TimeoutMsg); ok {
                        return NewDashboardModel(m.gameState), nil
                }
                
                return m, cmd
                
        case textAnimationTickMsg:
                if !m.animationDone {
                        text, done := m.textAnimator()
                        m.currentText = text
                        m.animationDone = done
                        
                        if !done {
                                return m, m.animateText()
                        }
                }
        }
        
        return m, nil
}

// View renders the splash screen
func (m SplashModel) View() string {
        // Create ASCII art logo
        logo := figure.NewColorFigure("HACKSIM", "", "green", true)
        logoStr := logo.String()
        
        // Create subtitle
        subtitle := Styles.Info.Render("A Terminal Hacking Simulator")
        
        // Animation text with blinking cursor
        animText := m.currentText
        if !m.animationDone {
                animText += "█"
        }
        
        animTextRendered := Styles.Normal.Render(animText)
        
        // Instruction text
        instruction := Styles.Subtitle.Render("Press ENTER to start or wait...")
        
        // Calculate spaces for vertical centering
        verticalPadding := 5
        
        return "\n" + strings.Repeat("\n", verticalPadding) +
                CenterText(logoStr, TerminalWidth) + "\n" +
                CenterText(subtitle, TerminalWidth) + "\n\n" +
                CenterText(animTextRendered, TerminalWidth) + "\n\n" +
                CenterText(instruction, TerminalWidth) + "\n"
}
