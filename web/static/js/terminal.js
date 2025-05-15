// Wait for the DOM to be fully loaded
document.addEventListener('DOMContentLoaded', function() {
    // Elements
    const scenarioSelector = document.getElementById('scenario-selector');
    const gameInterface = document.getElementById('game-interface');
    const terminalContainer = document.getElementById('terminal-container');
    const terminalTitle = document.getElementById('terminal-title');
    const connectionStatus = document.getElementById('connection-status');
    const statusText = document.getElementById('status-text');
    const missionName = document.getElementById('mission-name');
    const missionDescription = document.getElementById('mission-description');
    const objectivesList = document.getElementById('objectives-list');
    const progressFill = document.getElementById('progress-fill');
    const progressText = document.getElementById('progress-text');
    
    // Terminal instance
    let terminal;
    let fitAddon;
    
    // WebSocket connection
    let socket;
    let currentSession = null;
    
    // Initialize the terminal
    function initTerminal() {
        // Create the terminal
        terminal = new Terminal({
            cursorBlink: true,
            theme: {
                background: '#0C0C0C',
                foreground: '#33FF33',
                cursor: '#33FF33',
                selectionBackground: '#33FF33',
                selectionForeground: '#0C0C0C'
            },
            fontFamily: 'Courier New, monospace',
            fontSize: 14,
            scrollback: 1000
        });
        
        // Create and attach the fit addon
        fitAddon = new FitAddon.FitAddon();
        terminal.loadAddon(fitAddon);
        
        // Open the terminal
        terminal.open(terminalContainer);
        fitAddon.fit();
        
        // Handle terminal input
        let commandBuffer = '';
        
        terminal.onKey(({ key, domEvent }) => {
            const printable = !domEvent.altKey && !domEvent.ctrlKey && !domEvent.metaKey;
            
            // Handle Enter key
            if (domEvent.keyCode === 13) {
                terminal.write('\r\n');
                
                // Send command to server
                if (socket && socket.readyState === WebSocket.OPEN && commandBuffer.trim()) {
                    socket.send(commandBuffer);
                    commandBuffer = '';
                } else {
                    terminal.write('Not connected to server\r\n');
                    showPrompt();
                    commandBuffer = '';
                }
            } 
            // Handle Backspace key
            else if (domEvent.keyCode === 8) {
                if (commandBuffer.length > 0) {
                    commandBuffer = commandBuffer.slice(0, -1);
                    terminal.write('\b \b');
                }
            } 
            // Handle printable characters
            else if (printable && !domEvent.altKey && !domEvent.ctrlKey && !domEvent.metaKey) {
                commandBuffer += key;
                terminal.write(key);
            }
        });
        
        // Handle terminal resize
        window.addEventListener('resize', () => {
            fitAddon.fit();
        });
        
        // Set initial prompt
        showWelcomeMessage();
    }
    
    // Show welcome message
    function showWelcomeMessage() {
        terminal.write('\x1b[1;32m==========================================\r\n');
        terminal.write('     HackSim - Terminal Hacking Simulator\r\n');
        terminal.write('==========================================\x1b[0m\r\n\r\n');
        terminal.write('Connecting to mission server...\r\n');
    }
    
    // Show command prompt
    function showPrompt() {
        terminal.write('\r\n\x1b[1;32m$ \x1b[0m');
    }
    
    // Update the mission info panel
    function updateMissionInfo(scenario, description, objectives) {
        missionName.textContent = 'Mission: ' + scenario;
        missionDescription.textContent = description;
        
        // Clear existing objectives
        objectivesList.innerHTML = '';
        
        // Add objectives
        if (objectives && objectives.length) {
            objectives.forEach((objective, index) => {
                const li = document.createElement('li');
                li.setAttribute('data-id', index);
                li.textContent = objective.Description || objective.description;
                objectivesList.appendChild(li);
            });
        }
    }
    
    // Update the progress bar
    function updateProgress(progress) {
        const percentage = Math.floor(progress * 100);
        progressFill.style.width = percentage + '%';
        progressText.textContent = percentage + '%';
    }
    
    // Mark an objective as completed
    function markObjectiveCompleted(index) {
        const objective = objectivesList.querySelector(`li[data-id="${index}"]`);
        if (objective) {
            objective.classList.add('completed');
        }
    }
    
    // Connect to the WebSocket server
    function connectWebSocket(sessionId) {
        // Determine the appropriate WebSocket protocol
        const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:';
        const wsUrl = `${protocol}//${window.location.host}/ws?session=${sessionId}`;
        
        // Create WebSocket connection
        socket = new WebSocket(wsUrl);
        
        // Connection opened
        socket.addEventListener('open', (event) => {
            connectionStatus.classList.add('connected');
            statusText.textContent = 'Connected';
            terminal.write('\r\nConnected to mission server.\r\n');
            showPrompt();
        });
        
        // Listen for messages
        socket.addEventListener('message', (event) => {
            const data = JSON.parse(event.data);
            
            switch (data.type) {
                case 'info':
                    // Update mission info
                    updateMissionInfo(data.scenario, data.description, data.objectives);
                    break;
                
                case 'command_output':
                    // Display command output
                    terminal.write('\r\n' + data.output + '\r\n');
                    showPrompt();
                    
                    // Update progress
                    updateProgress(data.progress);
                    
                    // Check if objective was completed
                    if (data.objective_completed) {
                        const index = data.objective_index || 0;
                        markObjectiveCompleted(index);
                        terminal.write('\r\n\x1b[1;32m✓ Objective completed!\x1b[0m\r\n');
                    }
                    break;
                
                case 'game_complete':
                    // Display game complete message
                    terminal.write('\r\n\x1b[1;32m' + data.message + '\x1b[0m\r\n');
                    terminal.write('\r\n\x1b[1;32mMission Complete! All objectives achieved.\x1b[0m\r\n');
                    break;
                
                default:
                    console.log('Unknown message type:', data.type);
            }
        });
        
        // Connection closed
        socket.addEventListener('close', (event) => {
            connectionStatus.classList.remove('connected');
            statusText.textContent = 'Disconnected';
            terminal.write('\r\n\x1b[1;31mDisconnected from mission server.\x1b[0m\r\n');
            currentSession = null;
        });
        
        // Connection error
        socket.addEventListener('error', (event) => {
            console.error('WebSocket error:', event);
            terminal.write('\r\n\x1b[1;31mConnection error. Please try again.\x1b[0m\r\n');
        });
    }
    
    // Start a new game session
    function startGameSession(scenario) {
        // Show loading state
        scenarioSelector.style.display = 'none';
        gameInterface.style.display = 'flex';
        
        // Initialize the terminal
        initTerminal();
        
        // Create a new session via API
        fetch('/api/session', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ scenario: scenario })
        })
        .then(response => response.json())
        .then(data => {
            terminalTitle.textContent = 'Connected: ' + data.scenario;
            currentSession = data.session_id;
            
            // Connect to WebSocket with the session ID
            connectWebSocket(data.session_id);
        })
        .catch(error => {
            console.error('Error creating session:', error);
            terminal.write('\r\n\x1b[1;31mFailed to connect to mission server. Please try again.\x1b[0m\r\n');
            
            // Show scenario selector again after a brief delay
            setTimeout(() => {
                gameInterface.style.display = 'none';
                scenarioSelector.style.display = 'block';
            }, 3000);
        });
    }
    
    // Add event listeners to scenario cards
    document.querySelectorAll('.scenario-card').forEach(card => {
        card.querySelector('.start-btn').addEventListener('click', () => {
            const scenario = card.getAttribute('data-scenario');
            startGameSession(scenario);
        });
    });
});