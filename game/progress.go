package game

import (
	"fmt"
	"time"
)

// AchievementType defines the type of achievement
type AchievementType string

// Define achievement types
const (
	AchievementConnected   AchievementType = "connected"
	AchievementCommandRun  AchievementType = "command"
	AchievementDataFound   AchievementType = "data"
	AchievementObjectivesCompleted AchievementType = "objectives"
)

// Achievement represents a player achievement
type Achievement struct {
	Type        AchievementType
	Name        string
	Description string
	Timestamp   time.Time
	Value       string // IP address, command name, etc.
}

// GameProgression tracks player progress and achievements
type GameProgression struct {
	StartTime      time.Time
	LastActionTime time.Time
	Achievements   []Achievement
	NodesVisited   map[string]bool
	CommandsRun    map[string]int
	FilesAccessed  map[string]bool
}

// NewGameProgression creates a new game progression tracker
func NewGameProgression() *GameProgression {
	now := time.Now()
	return &GameProgression{
		StartTime:      now,
		LastActionTime: now,
		Achievements:   []Achievement{},
		NodesVisited:   make(map[string]bool),
		CommandsRun:    make(map[string]int),
		FilesAccessed:  make(map[string]bool),
	}
}

// AddAchievement adds a new achievement
func (p *GameProgression) AddAchievement(achievementType AchievementType, name, description, value string) {
	p.Achievements = append(p.Achievements, Achievement{
		Type:        achievementType,
		Name:        name,
		Description: description,
		Timestamp:   time.Now(),
		Value:       value,
	})
}

// RecordNodeVisit records that a node has been visited
func (p *GameProgression) RecordNodeVisit(ip string) bool {
	firstVisit := !p.NodesVisited[ip]
	p.NodesVisited[ip] = true
	p.LastActionTime = time.Now()
	return firstVisit
}

// RecordCommandRun records that a command has been run
func (p *GameProgression) RecordCommandRun(command string) {
	p.CommandsRun[command]++
	p.LastActionTime = time.Now()
}

// RecordFileAccess records that a file has been accessed
func (p *GameProgression) RecordFileAccess(filename string) bool {
	firstAccess := !p.FilesAccessed[filename]
	p.FilesAccessed[filename] = true
	p.LastActionTime = time.Now()
	return firstAccess
}

// GetStats returns statistics about the current game session
func (p *GameProgression) GetStats() string {
	elapsedTime := time.Since(p.StartTime).Round(time.Second)
	
	uniqueCommands := len(p.CommandsRun)
	totalCommands := 0
	for _, count := range p.CommandsRun {
		totalCommands += count
	}
	
	stats := fmt.Sprintf("Session Duration: %s\n", elapsedTime)
	stats += fmt.Sprintf("Nodes Visited: %d\n", len(p.NodesVisited))
	stats += fmt.Sprintf("Files Accessed: %d\n", len(p.FilesAccessed))
	stats += fmt.Sprintf("Unique Commands: %d\n", uniqueCommands)
	stats += fmt.Sprintf("Total Commands: %d\n", totalCommands)
	stats += fmt.Sprintf("Achievements: %d\n", len(p.Achievements))
	
	return stats
}

// ListAchievements returns a list of achievements
func (p *GameProgression) ListAchievements() []string {
	var result []string
	
	if len(p.Achievements) == 0 {
		return []string{"No achievements yet"}
	}
	
	for _, a := range p.Achievements {
		timeStr := a.Timestamp.Format("15:04:05")
		result = append(result, fmt.Sprintf("[%s] %s - %s", timeStr, a.Name, a.Description))
	}
	
	return result
}
