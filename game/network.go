package game

import (
        "fmt"
        "strings"
)

// Network represents the network topology
type Network struct {
        Nodes       map[string]Node
        CurrentNode string
        Discovered  map[string]bool
        Routes      map[string][]string // Maps IPs to possible routes
}

// Node represents a network node
type Node struct {
        Name        string
        IP          string
        Ports       map[string]Service
        AccessLevel int // 0=none, 1=user, 2=admin, 3=root
        Type        string // server, router, workstation, etc.
}

// Service represents a network service
type Service struct {
        Name        string
        Version     string
        Vulnerable  bool
        Credentials map[string]string // username -> password
}

// NewNetwork creates a new network with the given nodes
func NewNetwork(nodes []Node) Network {
        // Create node map for efficient lookup
        nodeMap := make(map[string]Node)
        for _, node := range nodes {
                nodeMap[node.IP] = node
        }
        
        // Initialize discovery map - start with no nodes discovered
        discovered := make(map[string]bool)
        
        // Initialize routes
        routes := make(map[string][]string)
        
        // Set some basic routes
        for _, node := range nodes {
                // For this simple model, all nodes can see each other
                var nodeRoutes []string
                for _, otherNode := range nodes {
                        if node.IP != otherNode.IP {
                                nodeRoutes = append(nodeRoutes, otherNode.IP)
                        }
                }
                routes[node.IP] = nodeRoutes
        }
        
        // Set initial node
        initialNode := "local"
        if len(nodes) > 0 {
                initialNode = nodes[0].IP
                discovered[initialNode] = true
        }
        
        return Network{
                Nodes:       nodeMap,
                CurrentNode: initialNode,
                Discovered:  discovered,
                Routes:      routes,
        }
}

// Scan performs a network scan and returns available nodes
func (n *Network) Scan(targetIP string) string {
        // If specific IP is given, scan just that IP
        if targetIP != "" {
                if node, exists := n.Nodes[targetIP]; exists {
                        n.Discovered[targetIP] = true
                        return n.formatNodeScan(node)
                }
                return "No host found at " + targetIP
        }
        
        // If no IP given, scan current network
        var results []string
        
        // Get routes from current node
        routes, exists := n.Routes[n.CurrentNode]
        if !exists || len(routes) == 0 {
                return "No network routes found from current node."
        }
        
        results = append(results, "Network scan results:")
        results = append(results, "-----------------------")
        
        // Scan each route
        for _, ip := range routes {
                if node, exists := n.Nodes[ip]; exists {
                        n.Discovered[ip] = true
                        results = append(results, fmt.Sprintf("%s (%s)", ip, node.Name))
                }
        }
        
        return strings.Join(results, "\n")
}

// Connect attempts to connect to a node
func (n *Network) Connect(ip, port string) (string, bool) {
        // Check if node exists
        node, exists := n.Nodes[ip]
        if !exists {
                return "Error: Host not found: " + ip, false
        }
        
        // Check if node is discoverable from current position
        routes, routesExist := n.Routes[n.CurrentNode]
        if !routesExist {
                return "Error: No route to host: " + ip, false
        }
        
        routeExists := false
        for _, route := range routes {
                if route == ip {
                        routeExists = true
                        break
                }
        }
        
        if !routeExists {
                return "Error: No route to host: " + ip, false
        }
        
        // Check if port exists
        service, portExists := node.Ports[port]
        if !portExists {
                return fmt.Sprintf("Error: Connection refused to %s on port %s", ip, port), false
        }
        
        // For this demo, we'll allow connection without credentials
        // In a more complex game, you'd check credentials here
        
        // Set current node to the connected node
        n.CurrentNode = ip
        
        // Show connection success message
        return fmt.Sprintf("Connected to %s (%s) on port %s\nService: %s %s", 
                ip, node.Name, port, service.Name, service.Version), true
}

// formatNodeScan returns a formatted string with node scan results
func (n *Network) formatNodeScan(node Node) string {
        result := fmt.Sprintf("Scan results for %s (%s):\n", node.IP, node.Name)
        result += "----------------------------------\n"
        result += fmt.Sprintf("Type: %s\n", node.Type)
        result += "Open Ports:\n"
        
        for port, service := range node.Ports {
                result += fmt.Sprintf("  %s: %s %s\n", port, service.Name, service.Version)
        }
        
        return result
}

// GetNetworkMap returns a text-based map of the network
func (n *Network) GetNetworkMap() string {
        var discovered []string
        var undiscovered []string
        
        for ip, node := range n.Nodes {
                if n.Discovered[ip] {
                        discovered = append(discovered, fmt.Sprintf("%s (%s) - %s", ip, node.Name, node.Type))
                } else {
                        undiscovered = append(undiscovered, "Unknown node")
                }
        }
        
        result := "DISCOVERED NODES:\n"
        result += "-----------------\n"
        result += strings.Join(discovered, "\n")
        
        if len(undiscovered) > 0 {
                result += "\n\nUNDISCOVERED NODES:\n"
                result += "-------------------\n"
                result += fmt.Sprintf("%d nodes not yet discovered", len(undiscovered))
        }
        
        return result
}
