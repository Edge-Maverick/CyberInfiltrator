package main

import (
        "hacksim/game"
        "log"
        "net/http"
        "sync"

        "github.com/gin-gonic/gin"
        "github.com/gorilla/websocket"
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

// RunServer initializes and runs the web server
func RunServer() {
        r := gin.Default()
        
        // Serve static files
        r.Static("/static", "./web/static")
        
        // Load HTML templates
        r.LoadHTMLGlob("web/templates/*")
        
        // Game home page
        r.GET("/", func(c *gin.Context) {
                c.HTML(http.StatusOK, "index.html", gin.H{
                        "title": "HackSim - Terminal Hacking Simulator",
                })
        })
        
        // Create a new game session
        r.POST("/api/session", func(c *gin.Context) {
                var req struct {
                        Scenario string `json:"scenario" binding:"required"`
                }
                
                if err := c.ShouldBindJSON(&req); err != nil {
                        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
                        return
                }
                
                // Generate a session ID (in a real app, use a more secure method)
                sessionID := "session-" + req.Scenario + "-" + c.ClientIP()
                
                // Create a new game session
                gameState := sessionManager.CreateSession(sessionID, req.Scenario)
                
                c.JSON(http.StatusOK, gin.H{
                        "session_id": sessionID,
                        "scenario": gameState.CurrentScenario.Name,
                        "description": gameState.CurrentScenario.Description,
                })
        })
        
        // WebSocket endpoint for real-time game communication
        r.GET("/ws/:session_id", func(c *gin.Context) {
                sessionID := c.Param("session_id")
                
                // Get the game session
                gameState, exists := sessionManager.GetSession(sessionID)
                if !exists {
                        c.JSON(http.StatusNotFound, gin.H{"error": "Session not found"})
                        return
                }
                
                // Upgrade HTTP connection to WebSocket
                conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
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
        })
        
        // Start the server on port 5000 (Replit's standard exposed port)
        log.Println("Starting HackSim Web Server on port 5000...")
        if err := r.Run(":5000"); err != nil {
                log.Fatal("Failed to start server:", err)
        }
}