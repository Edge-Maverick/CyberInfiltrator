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
    
    // Debug panel elements
    const wsStatus = document.getElementById('ws-status');
    const sessionIdElement = document.getElementById('session-id');
    const lastMessage = document.getElementById('last-message');
    const reconnectBtn = document.getElementById('reconnect-btn');
    
    // Terminal instance
    let terminal;
    let fitAddon;
    
    // WebSocket connection
    let socket;
    let currentSession = null;
    
    // WebSocket ready state constants (for browsers that don't expose these)
    const WS_CONNECTING = 0;
    const WS_OPEN = 1;
    const WS_CLOSING = 2;
    const WS_CLOSED = 3;
    
    // Initialize the terminal
    function initTerminal() {
        // Determine appropriate font size based on device width
        const isMobile = window.innerWidth < 768;
        const fontSize = isMobile ? 12 : 14;
        
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
            fontSize: fontSize,
            scrollback: 1000,
            allowTransparency: true
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
                if (commandBuffer.trim()) {
                    if (socket && socket.readyState === WS_OPEN) {
                        terminal.write(`Sending command: ${commandBuffer}\r\n`);
                        socket.send(commandBuffer);
                        commandBuffer = '';
                    } else {
                        let errorMsg = '';
                        if (!socket) {
                            errorMsg = 'Socket not initialized';
                        } else {
                            switch(socket.readyState) {
                                case WS_CONNECTING:
                                    errorMsg = 'Socket is still connecting...';
                                    break;
                                case WS_CLOSED:
                                    errorMsg = 'Socket is closed';
                                    break;
                                case WS_CLOSING:
                                    errorMsg = 'Socket is closing';
                                    break;
                                default:
                                    errorMsg = `Unknown socket state: ${socket.readyState}`;
                            }
                        }
                        terminal.write(`Not connected to server (${errorMsg})\r\n`);
                        showPrompt();
                        commandBuffer = '';
                    }
                } else {
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
        
        // Additional fit handler for mobile orientation changes
        window.addEventListener('orientationchange', () => {
            setTimeout(() => {
                fitAddon.fit();
            }, 100);
        });
        
        // Force fit after a short delay to ensure proper sizing on all devices
        setTimeout(() => {
            fitAddon.fit();
        }, 200);
        
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
            // Update objective counters
            document.getElementById('total-count').textContent = objectives.length;
            document.getElementById('completed-count').textContent = '0';
            
            objectives.forEach((objective, index) => {
                const li = document.createElement('li');
                li.setAttribute('data-id', index);
                li.textContent = objective.Description || objective.description;
                objectivesList.appendChild(li);
            });
            
            // Add event listeners for objectives panel
            const toggleBtn = document.getElementById('toggle-objectives-btn');
            const closeBtn = document.getElementById('close-objectives-btn');
            const objectivesPanel = document.getElementById('objectives-panel');
            
            toggleBtn.addEventListener('click', () => {
                objectivesPanel.style.display = 'block';
            });
            
            closeBtn.addEventListener('click', () => {
                objectivesPanel.style.display = 'none';
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
            
            // Update completed objectives counter
            const completedCount = document.getElementById('completed-count');
            const currentCompleted = parseInt(completedCount.textContent);
            completedCount.textContent = currentCompleted + 1;
            
            // Flash the counter to draw attention
            const objectiveCounter = document.querySelector('.objective-counter');
            objectiveCounter.style.color = '#ffffff';
            setTimeout(() => {
                objectiveCounter.style.color = 'var(--matrix-green)';
            }, 300);
        }
    }
    
    // Connect to the WebSocket server
    function connectWebSocket(sessionId, scenario) {
        // Determine the appropriate WebSocket protocol
        const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:';
        const wsUrl = `${protocol}//${window.location.host}/ws?session=${sessionId}&scenario=${scenario || 'network-breach'}`;
        
        // Update debug panel (invisible to user)
        wsStatus.textContent = 'Connecting...';
        sessionIdElement.textContent = sessionId;
        lastMessage.textContent = 'None yet';
        
        terminal.write(`Connecting to mission server...\r\n`);
        
        try {
            // Create WebSocket connection
            socket = new WebSocket(wsUrl);
            
            // Connection opened
            socket.addEventListener('open', (event) => {
                connectionStatus.classList.add('connected');
                statusText.textContent = 'Connected';
                wsStatus.textContent = 'Connected (OPEN)';
                terminal.write('\r\nConnected to mission server.\r\n');
                showPrompt();
            });
            
            // Listen for messages
            socket.addEventListener('message', (event) => {
                // Update debug panel with raw message (hidden from user)
                lastMessage.textContent = event.data.substring(0, 50) + (event.data.length > 50 ? '...' : '');
                
                try {
                    const data = JSON.parse(event.data);
                    
                    switch (data.type) {
                        case 'info':
                            // Update mission info without debug message
                            updateMissionInfo(data.scenario, data.description, data.objectives);
                            terminal.write('\r\nMission information received. Check objectives to see your tasks.\r\n');
                            showPrompt();
                            break;
                        
                        case 'command_output':
                            // First echo the command that was sent
                            if (data.command) {
                                terminal.write('\r\n\x1b[1;32m$ \x1b[0m' + data.command + '\r\n');
                            }
                            
                            // Then display command output
                            terminal.write(data.output + '\r\n');
                            
                            // Show a new prompt when done
                            showPrompt();
                            
                            // Update progress
                            updateProgress(data.progress);
                            
                            // Check if objective was completed
                            if (data.objective_completed) {
                                const index = data.objective_index || 0;
                                markObjectiveCompleted(index);
                                terminal.write('\r\n\x1b[1;32m✓ Objective completed!\x1b[0m\r\n');
                                showPrompt();
                            }
                            break;
                        
                        case 'game_complete':
                            // Display game complete message
                            terminal.write('\r\n\x1b[1;32m' + data.message + '\x1b[0m\r\n');
                            terminal.write('\r\n\x1b[1;32mMission Complete! All objectives achieved.\x1b[0m\r\n');
                            break;
                        
                        default:
                            terminal.write(`Unknown message type: ${data.type}\r\n`);
                            console.log('Unknown message type:', data.type);
                    }
                } catch (error) {
                    terminal.write(`\r\n\x1b[1;31mError parsing message: ${error.message}\x1b[0m\r\n`);
                    console.error('Error parsing message:', error);
                    lastMessage.textContent = `Error: ${error.message}`;
                }
            });
            
            // Connection closed
            socket.addEventListener('close', (event) => {
                connectionStatus.classList.remove('connected');
                statusText.textContent = 'Disconnected';
                wsStatus.textContent = `Closed (${event.code}: ${event.reason})`;
                terminal.write(`\r\n\x1b[1;31mWebSocket closed with code: ${event.code}, reason: ${event.reason}\x1b[0m\r\n`);
                currentSession = null;
            });
            
            // Connection error
            socket.addEventListener('error', (event) => {
                console.error('WebSocket error:', event);
                wsStatus.textContent = 'Error';
                terminal.write('\r\n\x1b[1;31mWebSocket connection error. Please try again.\x1b[0m\r\n');
            });
        } catch (error) {
            wsStatus.textContent = `Error: ${error.message}`;
            terminal.write(`\r\n\x1b[1;31mFailed to create WebSocket: ${error.message}\x1b[0m\r\n`);
            console.error('Failed to create WebSocket:', error);
        }
    }
    
    // Start a new game session
    function startGameSession(scenario) {
        // Show loading state
        scenarioSelector.style.display = 'none';
        gameInterface.style.display = 'flex';
        
        // Initialize the terminal
        initTerminal();
        
        // Show debug message in terminal
        terminal.write('\r\nConnecting to session API...\r\n');
        
        // Create a new session via API
        fetch('/api/session', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ scenario: scenario })
        })
        .then(response => {
            if (!response.ok) {
                throw new Error(`Server responded with status: ${response.status}`);
            }
            return response.json();
        })
        .then(data => {
            terminalTitle.textContent = 'Connected: ' + data.scenario;
            currentSession = data.session_id;
            
            // Connect to WebSocket with the session ID and scenario
            terminal.write(`Initializing mission "${scenario}"...\r\n`);
            connectWebSocket(data.session_id, scenario);
        })
        .catch(error => {
            console.error('Error creating session:', error);
            terminal.write(`\r\n\x1b[1;31mError: ${error.message}\x1b[0m\r\n`);
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
    
    // Add event listener for reconnect button
    reconnectBtn.addEventListener('click', () => {
        if (currentSession) {
            const scenarioName = terminalTitle.textContent.replace('Connected: ', '');
            terminal.write(`\r\nAttempting to reconnect to ${scenarioName}...\r\n`);
            connectWebSocket(currentSession, scenarioName);
        } else {
            terminal.write('\r\nNo active session to reconnect to. Please start a new game.\r\n');
        }
    });
    
    // Handle mobile input
    const mobileInput = document.getElementById('mobile-input');
    const sendCommandBtn = document.getElementById('send-command-btn');
    const commandBtns = document.querySelectorAll('.command-btn');
    
    function sendMobileCommand() {
        const command = mobileInput.value.trim();
        if (command && socket && socket.readyState === WS_OPEN) {
            socket.send(command);
            
            // We don't write the command to the terminal here anymore
            // The server will echo it back as part of the response
            
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
});