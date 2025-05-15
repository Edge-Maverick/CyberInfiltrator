package cmd

import (
        "encoding/json"
        "fmt"
        "hacksim/game"
        "log"
        "net/http"
        "sync"
        "time"

        "github.com/gorilla/websocket"
        "github.com/spf13/cobra"
)

// SessionManager handles game sessions
type SessionManager struct {
        sessions    map[string]*game.State
        connections map[string]*websocket.Conn
        mu          sync.Mutex
}

// NewSessionManager creates a new session manager
func NewSessionManager() *SessionManager {
        return &SessionManager{
                sessions:    make(map[string]*game.State),
                connections: make(map[string]*websocket.Conn),
        }
}

// CreateSession creates a new game session
func (sm *SessionManager) CreateSession(sessionID, scenario string) *game.State {
        sm.mu.Lock()
        defer sm.mu.Unlock()
        
        gameState := game.NewGameState(scenario)
        sm.sessions[sessionID] = gameState
        return gameState
}

// GetSession gets a game session by ID
func (sm *SessionManager) GetSession(sessionID string) (*game.State, bool) {
        sm.mu.Lock()
        defer sm.mu.Unlock()
        
        session, exists := sm.sessions[sessionID]
        return session, exists
}

// SaveConnection saves a WebSocket connection for a session
func (sm *SessionManager) SaveConnection(sessionID string, conn *websocket.Conn) {
        sm.mu.Lock()
        defer sm.mu.Unlock()
        
        sm.connections[sessionID] = conn
}

// GetConnection gets a WebSocket connection for a session
func (sm *SessionManager) GetConnection(sessionID string) (*websocket.Conn, bool) {
        sm.mu.Lock()
        defer sm.mu.Unlock()
        
        conn, exists := sm.connections[sessionID]
        return conn, exists
}

// CloseSession closes a session and its connection
func (sm *SessionManager) CloseSession(sessionID string) {
        sm.mu.Lock()
        defer sm.mu.Unlock()
        
        if conn, exists := sm.connections[sessionID]; exists {
                conn.Close()
                delete(sm.connections, sessionID)
        }
        
        delete(sm.sessions, sessionID)
}

// WebSocket upgrader
var upgrader = websocket.Upgrader{
        ReadBufferSize:  1024,
        WriteBufferSize: 1024,
        CheckOrigin: func(r *http.Request) bool {
                return true // Allow connections from any origin
        },
}

// Global session manager
var sessionManager = NewSessionManager()

// handleWebSocket handles WebSocket connections for terminal interactivity
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
        // Extract session ID from query parameter
        sessionID := r.URL.Query().Get("session")
        scenario := r.URL.Query().Get("scenario")
        
        // Set default values if not provided
        if sessionID == "" {
                sessionID = "session-default"
        }
        
        if scenario == "" {
                scenario = "network-breach"
        }
        
        // Create a new session if it doesn't exist
        gameState, exists := sessionManager.GetSession(sessionID)
        if !exists {
                gameState = sessionManager.CreateSession(sessionID, scenario)
        }
        
        // Upgrade HTTP connection to WebSocket
        conn, err := upgrader.Upgrade(w, r, nil)
        if err != nil {
                log.Println("Failed to upgrade to WebSocket:", err)
                return
        }
        
        // Save the WebSocket connection
        sessionManager.SaveConnection(sessionID, conn)
        
        // Send initial game information
        initialMsg := map[string]interface{}{
                "type": "info",
                "scenario": gameState.CurrentScenario.Name,
                "description": gameState.CurrentScenario.Description,
                "objectives": gameState.CurrentScenario.Objectives,
        }
        conn.WriteJSON(initialMsg)
        
        // WebSocket message handling loop
        for {
                // Read message from client
                messageType, message, err := conn.ReadMessage()
                if err != nil {
                        log.Println("WebSocket read error:", err)
                        sessionManager.CloseSession(sessionID)
                        break
                }
                
                if messageType == websocket.TextMessage {
                        // Process the command
                        cmd := string(message)
                        output := gameState.ProcessCommand(cmd)
                        
                        // Check if an objective was completed
                        objectiveCompleted := gameState.CheckObjectiveCompletion(cmd)
                        
                        // Send the response back to the client
                        response := map[string]interface{}{
                                "type": "command_output",
                                "command": cmd,
                                "output": output,
                                "objective_completed": objectiveCompleted,
                                "progress": gameState.Progress,
                        }
                        
                        if err := conn.WriteJSON(response); err != nil {
                                log.Println("WebSocket write error:", err)
                                break
                        }
                        
                        // Check if game is complete
                        if gameState.Progress >= 1.0 {
                                completeMsg := map[string]interface{}{
                                        "type": "game_complete",
                                        "message": "Congratulations! All objectives completed.",
                                }
                                conn.WriteJSON(completeMsg)
                        }
                }
        }
}

// serveHome serves the home page
func serveHome(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path != "/" {
                http.Error(w, "Not found", http.StatusNotFound)
                return
        }
        if r.Method != http.MethodGet {
                http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
                return
        }
        http.ServeFile(w, r, "web/templates/index.html")
}

// handleSessionCreation handles the creation of new game sessions
func handleSessionCreation(w http.ResponseWriter, r *http.Request) {
        // Only allow POST method
        if r.Method != http.MethodPost {
                http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
                return
        }

        // Decode the request body
        var req struct {
                Scenario string `json:"scenario"`
        }

        // Read the body
        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
                http.Error(w, "Invalid request body", http.StatusBadRequest)
                return
        }

        // Generate a session ID
        sessionID := "session-" + req.Scenario + "-" + time.Now().Format("150405")

        // Create a new game state
        gameState := sessionManager.CreateSession(sessionID, req.Scenario)

        // Prepare the response
        response := map[string]interface{}{
                "session_id":  sessionID,
                "scenario":    gameState.CurrentScenario.Name,
                "description": gameState.CurrentScenario.Description,
        }

        // Write the response
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(response)
}

// startWebServer starts the HTTP server
func startWebServer(port string) {
        // Create routes
        http.HandleFunc("/", serveHome)
        http.HandleFunc("/ws", handleWebSocket)
        http.HandleFunc("/api/session", handleSessionCreation)
        http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

        // Start server
        addr := fmt.Sprintf("0.0.0.0:%s", port)
        log.Printf("Starting HackSim Web Server on %s", addr)
        log.Fatal(http.ListenAndServe(addr, nil))
}

// webCmd represents the web command
var webCmd = &cobra.Command{
        Use:   "web",
        Short: "Start HackSim as a web application",
        Long: `Start HackSim as a web application that can be accessed via a browser.
This mode allows playing the game through a web interface instead of the terminal.`,
        Run: func(cmd *cobra.Command, args []string) {
                fmt.Println("Starting HackSim Web Server...")
                
                // Get port from flag or use default
                port, _ := cmd.Flags().GetString("port")
                if port == "" {
                        port = "5000"
                }
                
                // Start web server
                startWebServer(port)
        },
}

func init() {
        rootCmd.AddCommand(webCmd)
        
        // Add flags
        webCmd.Flags().StringP("port", "p", "5000", "Port for the web server")
}