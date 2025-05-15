package game

import (
        "strings"
        "testing"
)

func TestNewGameState(t *testing.T) {
        tests := []struct {
                name     string
                scenario string
                wantName string
        }{
                {"Default scenario", "network-breach", "Network Breach"},
                {"Data heist scenario", "data-heist", "Data Heist"},
                {"System takeover scenario", "system-takeover", "System Takeover"},
                {"Invalid scenario defaults to network breach", "invalid", "Network Breach"},
        }

        for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                        state := NewGameState(tt.scenario)
                        if state.CurrentScenario.Name != tt.wantName {
                                t.Errorf("NewGameState(%s) got scenario name = %s, want %s", 
                                        tt.scenario, state.CurrentScenario.Name, tt.wantName)
                        }
                })
        }
}

func TestProcessCommand(t *testing.T) {
        state := NewGameState("network-breach")
        
        // Unlock the status command for testing
        state.unlockTool("status")
        
        tests := []struct {
                command string
                wantContains string
        }{
                {"help", "Available commands"},
                {"status", "Mission: Network Breach"},
                {"scan", "Network scan results"},
                {"connect 192.168.1.1", "Connected to 192.168.1.1"},
                {"ls", "Directory listing"},
                {"invalid-command", "Command 'invalid-command' not found or access denied"},
                {"", "Empty command"},
        }
        
        for _, tt := range tests {
                t.Run(tt.command, func(t *testing.T) {
                        result := state.ProcessCommand(tt.command)
                        if !strings.Contains(result, tt.wantContains) {
                                t.Errorf("ProcessCommand(%s) = %s, want to contain %s", 
                                        tt.command, result, tt.wantContains)
                        }
                })
        }
}

func TestObjectiveCompletion(t *testing.T) {
        state := NewGameState("network-breach")
        
        // Store initial progress
        initialProgress := state.Progress
        
        // Set up the test by marking a command objective
        found := false
        for i, obj := range state.CurrentScenario.Objectives {
                if obj.Type == "command" && obj.Target == "scan" {
                        found = true
                        // Make sure the objective is not already completed
                        state.CurrentScenario.Objectives[i].Completed = false
                        break
                }
        }
        
        if !found {
                t.Fatal("Test setup failed: could not find scan objective")
        }
        
        // Manually check objective completion for the scan command
        changed := state.CheckObjectiveCompletion("scan")
        
        // Progress should increase and return true
        if !changed {
                t.Error("CheckObjectiveCompletion should return true for scan command")
        }
        
        if state.Progress <= initialProgress {
                t.Errorf("Progress should increase after completing 'scan' objective, got %.2f, initial was %.2f", 
                        state.Progress, initialProgress)
        }
        
        // Check if the objective was marked as completed
        scanObjectiveCompleted := false
        for _, obj := range state.CurrentScenario.Objectives {
                if obj.Type == "command" && obj.Target == "scan" && obj.Completed {
                        scanObjectiveCompleted = true
                        break
                }
        }
        
        if !scanObjectiveCompleted {
                t.Error("Scan objective should be marked as completed")
        }
}