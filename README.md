# HackSim: Terminal-Based Hacking Simulation Game

![HackSim Logo](generated-icon.png)

HackSim is an immersive terminal-based hacking simulation game built with Go and modern web technologies. It allows players to complete hacking missions through an authentic terminal experience accessible via desktop and mobile browsers.

## Table of Contents

- [Features](#features)
- [Game Architecture](#game-architecture)
- [Requirements](#requirements)
- [Installation](#installation)
- [Running the Game](#running-the-game)
- [Project Structure](#project-structure)
- [Implementation Guide](#implementation-guide)
  - [Backend Setup](#backend-setup)
  - [Frontend Setup](#frontend-setup)
  - [WebSocket Integration](#websocket-integration)
  - [Deployment](#deployment)
- [Extending the Game](#extending-the-game)
- [Mobile Support](#mobile-support)
- [Troubleshooting](#troubleshooting)
- [Contributing](#contributing)
- [License](#license)

## Features

- Terminal-based interface with authentic command-line experience
- Web-based accessibility (no installation required for players)
- Multiple hacking scenarios (Network Breach, Data Heist, System Takeover)
- Realistic commands including scan, connect, ls, cat, and crack
- Mission objectives and progress tracking
- Mobile-friendly interface with command suggestions
- Real-time interaction via WebSockets
- Collapsible mission objectives panel
- Progress tracking with visual indicators

## Game Architecture

HackSim uses a client-server architecture:

- **Backend**: Go server with WebSocket support for real-time communication
- **Frontend**: JavaScript terminal emulator (xterm.js) with custom styling
- **Game Engine**: Core game logic implemented in Go with simulated file systems and networks
- **Web Interface**: Responsive design supporting both desktop and mobile browsers

## Requirements

- Go 1.16 or higher
- Web browser with WebSocket support
- Internet connection (for multiplayer features)

## Installation

To install and run the game locally:

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/hacksim.git
   cd hacksim
   ```

2. Install Go dependencies:
   ```
   go mod tidy
   ```

3. Build the project:
   ```
   go build -o hacksim
   ```

## Running the Game

### Terminal Mode

Run the game in terminal mode:
```
./hacksim play
```

For a simpler interface:
```
./hacksim play --simple
```

### Web Mode

Start the web server:
```
./hacksim web
```

Then open a browser and navigate to:
```
http://localhost:5000
```

## Project Structure

```
├── assets/ - ASCII art and other static assets
├── cmd/ - Command-line interface components
│   ├── help.go - Help command implementation
│   ├── play.go - Game play command implementation
│   ├── root.go - Root command implementation
│   └── web.go - Web server command implementation
├── demos/ - Demo scripts and scenarios
├── game/ - Core game engine
│   ├── filesystem.go - Virtual filesystem implementation
│   ├── network.go - Virtual network implementation
│   ├── progress.go - Progress tracking
│   ├── scenarios.go - Game scenarios definition
│   └── state.go - Game state management
├── ui/ - Terminal UI components
│   ├── common.go - Shared UI elements
│   ├── dashboard.go - Main dashboard UI
│   └── splash.go - Splash screen
├── web/ - Web interface
│   ├── static/ - Static web assets (CSS, JS)
│   ├── templates/ - HTML templates
│   ├── main.go - Web entry point
│   └── server.go - Web server implementation
├── go.mod - Go module definition
├── go.sum - Go module checksums
└── main.go - Main application entry point
```

## Implementation Guide

### Backend Setup

1. **Create Go Module**:
   ```
   go mod init hacksim
   ```

2. **Install Required Dependencies**:
   ```
   go get github.com/spf13/cobra
   go get github.com/charmbracelet/bubbletea
   go get github.com/gin-gonic/gin
   go get github.com/gorilla/websocket
   ```

3. **Implement Game Engine**:
   - Create virtual filesystem (`game/filesystem.go`)
   - Implement network simulation (`game/network.go`)
   - Define game scenarios (`game/scenarios.go`)
   - Set up state management (`game/state.go`)

4. **Create Web Server**:
   - Set up Gin router in `web/server.go`
   - Implement WebSocket endpoint for real-time communication
   - Create session management for multiplayer support

### Frontend Setup

1. **HTML Structure**:
   - Create main interface in `web/templates/index.html`
   - Include terminal container and mission info elements
   - Add mobile-friendly controls

2. **CSS Styling**:
   - Define terminal styling in `web/static/css/terminal.css`
   - Set up responsive design for mobile devices
   - Implement matrix-inspired theme with green-on-black color scheme

3. **JavaScript Implementation**:
   - Initialize xterm.js terminal in `web/static/js/terminal.js`
   - Set up WebSocket connection for real-time updates
   - Handle user input and display command results
   - Implement mission objective tracking

### Real-World Terminal Experience

To implement the authentic terminal experience:

1. **Install xterm.js**:
   ```html
   <!-- Add to your HTML file -->
   <script src="https://cdn.jsdelivr.net/npm/xterm@5.1.0/lib/xterm.min.js"></script>
   <script src="https://cdn.jsdelivr.net/npm/xterm-addon-fit@0.7.0/lib/xterm-addon-fit.min.js"></script>
   <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/xterm@5.1.0/css/xterm.min.css">
   ```

2. **Initialize the Terminal**:
   ```javascript
   // Terminal instance
   let terminal;
   let fitAddon;
   
   // Initialize the terminal
   function initTerminal() {
       // Determine appropriate font size based on device width
       const isMobile = window.innerWidth < 768;
       const fontSize = isMobile ? 12 : 14;
       
       terminal = new Terminal({
           cursorBlink: true,
           theme: {
               background: '#000000',
               foreground: '#33ff33',
               cursor: '#33ff33',
               cursorAccent: '#000000',
               selection: 'rgba(51, 255, 51, 0.3)',
               black: '#000000',
               red: '#ff3333',
               green: '#33ff33',
               yellow: '#ffff33',
               blue: '#3333ff',
               magenta: '#ff33ff',
               cyan: '#33ffff',
               white: '#f0f0f0',
               brightBlack: '#333333',
               brightRed: '#ff6666',
               brightGreen: '#66ff66',
               brightYellow: '#ffff66',
               brightBlue: '#6666ff',
               brightMagenta: '#ff66ff',
               brightCyan: '#66ffff',
               brightWhite: '#ffffff',
               selectionForeground: '#0C0C0C'
           },
           fontFamily: 'Courier New, monospace',
           fontSize: fontSize,
           scrollback: 1000,
           allowTransparency: true
       });
       
       // Create and attach the fit addon
       fitAddon = new FitAddon.FitAddon();
       terminal.loadAddon(fitAddon);
       
       // Open the terminal
       terminal.open(document.getElementById('terminal-container'));
       fitAddon.fit();
   }
   ```

3. **Create Professional Command Prompt**:
   ```javascript
   // Show command prompt with username, hostname and path
   function showPrompt() {
       const username = 'hacker';
       const hostname = 'hacksim';
       const currentPath = '~';
       terminal.write('\r\n\x1b[1;32m' + username + '@' + hostname + '\x1b[0m:\x1b[1;34m' + currentPath + '\x1b[0m\x1b[1;32m$ \x1b[0m');
   }
   ```

4. **Format Command Output**:
   ```javascript
   // Format command output with colors and proper alignment
   function formatOutput(output, command) {
       // Handle specific commands with special formatting
       if (command === 'ls') {
           // Special handling for directory listings
           let lines = output.split('\n');
           
           if (lines.length > 1) {
               // Write header line
               terminal.write(lines[0] + '\r\n');
               
               // Process the actual directory listing with proper spacing
               for (let i = 1; i < lines.length; i++) {
                   if (lines[i].trim()) {
                       let parts = lines[i].trim().split(/\s+/);
                       if (parts.length >= 2) {
                           // Format directory entries with colors
                           if (parts[0].startsWith('d')) {
                               // Directory - blue color
                               terminal.write(parts[0] + '  \x1b[1;34m' + parts[1] + '\x1b[0m\r\n');
                           } else {
                               // File - normal color
                               terminal.write(parts[0] + '  ' + parts[1] + '\r\n');
                           }
                       } else {
                           terminal.write(lines[i] + '\r\n');
                       }
                   }
               }
           } else {
               terminal.write(output + '\r\n');
           }
       } else {
           // Add colors to specific keywords
           let formattedOutput = output
               .replace(/error:/gi, '\x1b[1;31merror:\x1b[0m')
               .replace(/success/gi, '\x1b[1;32msuccess\x1b[0m')
               .replace(/warning:/gi, '\x1b[1;33mwarning:\x1b[0m');
           
           // Display formatted output
           terminal.write(formattedOutput + '\r\n');
       }
   }
   ```

5. **Add Command Delay for Realism**:
   ```javascript
   // Process command with realistic delay
   function processCommand(command) {
       if (command && socket && socket.readyState === WebSocket.OPEN) {
           // Don't echo command here - server will echo it back with proper formatting
           socket.send(command);
           
           // Add subtle "thinking" delay for realism
           setTimeout(() => {
               // Response will be handled by the WebSocket message event
           }, 150);
       } else {
           terminal.write('\r\nConnection error: Terminal offline\r\n');
           showPrompt();
       }
   }
   ```

6. **Handle Terminal Input**:
   ```javascript
   // Command buffer
   let commandBuffer = '';
   
   // Handle key input
   terminal.onKey(({ key, domEvent }) => {
       const printable = !domEvent.altKey && !domEvent.ctrlKey && !domEvent.metaKey;
       
       // Handle Enter key
       if (domEvent.keyCode === 13) {
           terminal.write('\r\n');
           
           // Process command
           if (commandBuffer.trim()) {
               processCommand(commandBuffer.trim());
               commandBuffer = '';
           } else {
               showPrompt();
           }
       }
       // Handle Backspace
       else if (domEvent.keyCode === 8) {
           if (commandBuffer.length > 0) {
               commandBuffer = commandBuffer.substr(0, commandBuffer.length - 1);
               terminal.write('\b \b');
           }
       }
       // Handle printable characters
       else if (printable) {
           commandBuffer += key;
           terminal.write(key);
       }
   });
   ```

### WebSocket Integration

1. **Server-Side WebSocket**:
   - Implement WebSocket upgrader in `web/server.go`
   - Create message handling loop for client commands
   - Send game state updates to clients

2. **Client-Side WebSocket**:
   - Establish connection in `web/static/js/terminal.js`
   - Send user commands to server
   - Process and display server responses
   - Handle connection errors and reconnection

### Deployment

1. **Local Deployment**:
   - Build executable: `go build -o hacksim`
   - Run web server: `./hacksim web`
   - Access at http://localhost:5000

2. **Replit Deployment**:
   - Create a new Repl with Go template
   - Upload or clone the HackSim code to your Repl
   - Configure the `.replit` file to run the web server:
     ```
     run = "go run main.go web"
     ```
   - Create the following workflows for different modes:
     ```
     [HackSim Game]
     run = "go mod tidy && go run main.go"
     
     [run-hacksim]
     run = "go run main.go && go run main.go play --simple"
     
     [HackSim Web]
     run = "go run main.go web"
     ```
   - The Replit web view will automatically open when the server starts
   - For deployment, use the "Deploy" button in Replit to make it publicly accessible

3. **Server Deployment**:
   - Choose a hosting provider with Go support
   - Set up firewall to allow port 5000 (or configure for standard HTTP/HTTPS)
   - Configure reverse proxy (Nginx/Apache) for better security

4. **Docker Deployment**:
   ```
   # Create a Dockerfile
   FROM golang:1.16-alpine
   WORKDIR /app
   COPY . .
   RUN go mod download
   RUN go build -o hacksim
   EXPOSE 5000
   CMD ["./hacksim", "web"]
   
   # Build and run
   docker build -t hacksim .
   docker run -p 5000:5000 hacksim
   ```

## Extending the Game

### Adding New Scenarios

1. Create a new scenario definition in `game/scenarios.go`:
   ```go
   var NewScenario = &Scenario{
       Name:        "scenario-name",
       Description: "Scenario description",
       Objectives:  []Objective{
           {Description: "First objective", Complete: false},
           {Description: "Second objective", Complete: false},
       },
       // Define virtual filesystem and network
   }
   ```

2. Register the scenario in `game/scenarios.go`:
   ```go
   func GetScenarios() map[string]*Scenario {
       return map[string]*Scenario{
           "network-breach":   NetworkBreachScenario,
           "data-heist":       DataHeistScenario,
           "system-takeover":  SystemTakeoverScenario,
           "scenario-name":    NewScenario,
       }
   }
   ```

3. Add the scenario to the web interface in `web/templates/index.html`:
   ```html
   <div class="scenario-card" data-scenario="scenario-name">
       <h3>Scenario Title</h3>
       <p>Scenario description</p>
       <button class="start-btn">Start Mission</button>
   </div>
   ```

### Adding New Commands

1. Implement command logic in `game/state.go`:
   ```go
   func (s *State) ProcessCommand(cmd string) string {
       // ... existing command processing ...
       
       case "newcommand":
           return s.handleNewCommand(args)
           
       // ... more commands ...
   }
   
   func (s *State) handleNewCommand(args []string) string {
       // Implement command logic
       return "Command output"
   }
   ```

2. Add command to help text in `cmd/help.go`:
   ```go
   var commandHelp = map[string]string{
       // ... existing commands ...
       "newcommand": "Description of the new command",
   }
   ```

## Mobile Support

HackSim includes special features for mobile users:

1. **Responsive Design**:
   - Fluid layout that adapts to screen size
   - Touch-friendly interface elements
   - Optimized font sizes for readability

2. **Mobile Command Input**:
   - Dedicated text input field for commands
   - Send button for command execution
   - Virtual keyboard support

3. **Command Suggestions**:
   - Quick-access buttons for common commands
   - Tap to insert commands into input field
   - Context-sensitive suggestions

4. **Collapsible Panels**:
   - Objectives panel can be toggled to save space
   - Simplified mission counter (completed/total)
   - Optimized vertical space usage

### Mobile Implementation Guide

To implement the mobile-friendly features:

1. **Add Mobile Input Container**:
   ```html
   <div class="mobile-input-section">
       <div class="command-suggestions" id="command-suggestions">
           <button class="command-btn" data-command="help">help</button>
           <button class="command-btn" data-command="ls">ls</button>
           <button class="command-btn" data-command="scan">scan</button>
           <!-- Add more command buttons as needed -->
       </div>
       <div class="mobile-input-container">
           <input type="text" id="mobile-input" placeholder="Enter command...">
           <button id="send-command-btn">Send</button>
       </div>
   </div>
   ```

2. **Add CSS for Mobile Elements**:
   ```css
   .mobile-input-section {
       display: flex;
       flex-direction: column;
       width: 100%;
       margin: 0;
       padding: 0;
   }

   .command-suggestions {
       display: flex;
       flex-wrap: wrap;
       gap: 5px;
       padding: 8px;
       background-color: var(--header-bg);
       border-top: 1px solid var(--matrix-green);
   }

   .command-btn {
       background-color: #000;
       color: var(--matrix-green);
       border: 1px solid var(--matrix-green);
       padding: 5px 10px;
       cursor: pointer;
       font-size: 14px;
   }

   .mobile-input-container {
       display: flex;
       width: 100%;
       background-color: var(--header-bg);
   }

   #mobile-input {
       flex: 1;
       padding: 12px 10px;
       font-family: 'Courier New', monospace;
       font-size: 16px;
       background-color: #000;
       color: var(--matrix-green);
       border: none;
   }

   #send-command-btn {
       padding: 12px 20px;
       background-color: #000;
       color: var(--matrix-green);
       border: none;
       border-left: 1px solid var(--matrix-green);
   }
   ```

3. **Add JavaScript for Mobile Command Handling**:
   ```javascript
   // Handle mobile input
   const mobileInput = document.getElementById('mobile-input');
   const sendCommandBtn = document.getElementById('send-command-btn');
   const commandBtns = document.querySelectorAll('.command-btn');
   
   function sendMobileCommand() {
       const command = mobileInput.value.trim();
       if (command && socket && socket.readyState === WebSocket.OPEN) {
           socket.send(command);
           mobileInput.value = '';
       } else if (command) {
           terminal.write('\r\nNot connected to server. Please try again.\r\n');
           showPrompt();
       }
   }
   
   mobileInput.addEventListener('keydown', (e) => {
       if (e.key === 'Enter') {
           sendMobileCommand();
       }
   });
   
   sendCommandBtn.addEventListener('click', sendMobileCommand);
   
   // Handle command suggestion buttons
   commandBtns.forEach(btn => {
       btn.addEventListener('click', () => {
           const command = btn.getAttribute('data-command');
           mobileInput.value = command;
           mobileInput.focus();
       });
   });
   ```

4. **Implement Collapsible Objectives Panel**:
   ```html
   <div class="mobile-mission-controls">
       <button id="toggle-objectives-btn" class="control-btn">Objectives</button>
       <div class="objective-counter">
           <span id="completed-count">0</span>/<span id="total-count">0</span>
       </div>
   </div>
   <div class="objectives-panel" id="objectives-panel" style="display: none;">
       <div class="objectives-header">
           <h4>Mission Objectives</h4>
           <button id="close-objectives-btn" class="close-btn">×</button>
       </div>
       <ul id="objectives-list"></ul>
   </div>
   ```

5. **JavaScript for Objectives Toggle**:
   ```javascript
   const toggleBtn = document.getElementById('toggle-objectives-btn');
   const closeBtn = document.getElementById('close-objectives-btn');
   const objectivesPanel = document.getElementById('objectives-panel');
   
   toggleBtn.addEventListener('click', () => {
       objectivesPanel.style.display = 'block';
   });
   
   closeBtn.addEventListener('click', () => {
       objectivesPanel.style.display = 'none';
   });
   
   // Update counter when objectives are completed
   function markObjectiveCompleted(index) {
       const objective = objectivesList.querySelector(`li[data-id="${index}"]`);
       if (objective) {
           objective.classList.add('completed');
           
           // Update completed objectives counter
           const completedCount = document.getElementById('completed-count');
           const currentCompleted = parseInt(completedCount.textContent);
           completedCount.textContent = currentCompleted + 1;
       }
   }
   ```

6. **Implement Responsive Media Queries**:
   ```css
   @media (max-width: 768px) {
       .terminal-container {
           flex: 1;
           min-height: 200px;
       }
       
       .game-interface {
           flex-direction: column;
       }
       
       .mission-info {
           flex-direction: column;
       }
       
       .mission-details {
           width: 100%;
       }
   }
   ```

## Troubleshooting

### Common Issues:

1. **WebSocket Connection Fails**:
   - Check if the server is running
   - Verify that port 5000 is not blocked by firewall
   - Ensure browser supports WebSockets

2. **Game Commands Not Working**:
   - Check terminal connection status (green indicator)
   - Try reconnecting with the reconnect button
   - Verify command syntax in help documentation

3. **Mobile Interface Issues**:
   - Enable "Request Desktop Site" for better experience on small screens
   - Try landscape orientation for wider terminal view
   - Clear browser cache if styling appears broken

## Contributing

We welcome contributions to HackSim! Please follow these steps:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Commit your changes (`git commit -m 'Add some amazing feature'`)
5. Push to the branch (`git push origin feature/amazing-feature`)
6. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

---

Created by [Your Name/Organization] - Happy Hacking!