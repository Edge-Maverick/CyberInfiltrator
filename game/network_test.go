package game

import (
        "strings"
        "testing"
)

func TestNetworkScan(t *testing.T) {
        // Create a test network with some nodes
        nodes := []Node{
                {
                        Name: "Test Node 1",
                        IP:   "10.0.0.1",
                        Type: "Server",
                        Ports: map[string]Service{
                                "22": {Name: "SSH", Version: "1.0"},
                        },
                },
                {
                        Name: "Test Node 2",
                        IP:   "10.0.0.2",
                        Type: "Workstation",
                        Ports: map[string]Service{
                                "80": {Name: "HTTP", Version: "1.0"},
                        },
                },
        }

        network := NewNetwork(nodes)
        
        // Set the current node to the first test node to establish routes
        network.CurrentNode = "10.0.0.1"
        
        // Make sure all nodes are in the routes
        if _, exists := network.Routes[network.CurrentNode]; !exists {
                network.Routes[network.CurrentNode] = []string{"10.0.0.2"}
        }

        // Test general scan
        scanResult := network.Scan("")
        if !strings.Contains(scanResult, "10.0.0.2") {
                t.Errorf("Network scan should show connected nodes, got: %s", scanResult)
        }

        // Test specific IP scan
        specificScan := network.Scan("10.0.0.1")
        if !strings.Contains(specificScan, "Test Node 1") || !strings.Contains(specificScan, "Type: Server") {
                t.Errorf("Specific scan should show node details, got: %s", specificScan)
        }

        // Test non-existent IP
        nonExistentScan := network.Scan("10.0.0.99")
        if !strings.Contains(nonExistentScan, "No host found") {
                t.Errorf("Scan for non-existent IP should return 'No host found', got: %s", nonExistentScan)
        }
}

func TestNetworkConnect(t *testing.T) {
        // Create a test network with some nodes
        nodes := []Node{
                {
                        Name: "Test Node 1",
                        IP:   "10.0.0.1",
                        Type: "Server",
                        Ports: map[string]Service{
                                "22": {Name: "SSH", Version: "1.0"},
                        },
                },
                {
                        Name: "Test Node 2",
                        IP:   "10.0.0.2",
                        Type: "Workstation",
                        Ports: map[string]Service{
                                "80": {Name: "HTTP", Version: "1.0"},
                        },
                },
        }

        network := NewNetwork(nodes)
        
        // Set up routes for testing
        network.CurrentNode = "10.0.0.1"
        network.Routes["10.0.0.1"] = []string{"10.0.0.2"}
        network.Routes["10.0.0.2"] = []string{"10.0.0.1"}
        
        // Discover nodes for testing
        network.Discovered["10.0.0.1"] = true
        network.Discovered["10.0.0.2"] = true

        // Test successful connection
        result, success := network.Connect("10.0.0.2", "80")
        if !success || !strings.Contains(result, "Connected to 10.0.0.2") {
                t.Errorf("Connection to valid node should succeed, got: %s, success: %v", result, success)
        }

        // Test connection to valid IP but invalid port
        result, success = network.Connect("10.0.0.2", "22")
        if success || !strings.Contains(result, "Connection refused") {
                t.Errorf("Connection to invalid port should fail, got: %s, success: %v", result, success)
        }

        // Test connection to invalid IP
        result, success = network.Connect("10.0.0.99", "22")
        if success || !strings.Contains(result, "Host not found") {
                t.Errorf("Connection to invalid IP should fail, got: %s, success: %v", result, success)
        }
}