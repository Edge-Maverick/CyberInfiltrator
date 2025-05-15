package game

// Import packages as needed

// Scenario represents a game scenario
type Scenario struct {
        Name                string
        Description         string
        Objectives          []Objective
        NetworkNodes        []Node
        FileSystems         map[string]*Directory
        InitialSecurityLevel int
}

// Objective represents a mission objective
type Objective struct {
        Description string
        Type        string // connect, data, command
        Target      string
        Completed   bool
}

// NetworkBreachScenario creates the Network Breach scenario
func NetworkBreachScenario() Scenario {
        // Create objectives
        objectives := []Objective{
                {
                        Description: "Discover available nodes on the network",
                        Type:        "command",
                        Target:      "scan",
                        Completed:   false,
                },
                {
                        Description: "Connect to the gateway server at 192.168.1.1",
                        Type:        "connect",
                        Target:      "192.168.1.1",
                        Completed:   false,
                },
                {
                        Description: "Find and read the network configuration file",
                        Type:        "data",
                        Target:      "network.conf",
                        Completed:   false,
                },
                {
                        Description: "Connect to the database server at 192.168.1.10",
                        Type:        "connect",
                        Target:      "192.168.1.10",
                        Completed:   false,
                },
                {
                        Description: "Crack the admin credentials",
                        Type:        "command",
                        Target:      "crack",
                        Completed:   false,
                },
                {
                        Description: "Download the user database file",
                        Type:        "command",
                        Target:      "download",
                        Completed:   false,
                },
        }
        
        // Create network nodes
        nodes := []Node{
                {
                        Name: "Entry Point",
                        IP:   "192.168.1.5",
                        Ports: map[string]Service{
                                "22": {
                                        Name:    "SSH",
                                        Version: "OpenSSH 7.9",
                                        Vulnerable: false,
                                        Credentials: map[string]string{
                                                "user": "password123",
                                        },
                                },
                                "80": {
                                        Name:    "HTTP",
                                        Version: "Apache 2.4.38",
                                        Vulnerable: true,
                                        Credentials: map[string]string{},
                                },
                        },
                        AccessLevel: 1,
                        Type:        "Workstation",
                },
                {
                        Name: "Gateway Server",
                        IP:   "192.168.1.1",
                        Ports: map[string]Service{
                                "22": {
                                        Name:    "SSH",
                                        Version: "OpenSSH 7.4",
                                        Vulnerable: false,
                                        Credentials: map[string]string{
                                                "admin": "gateway2023",
                                        },
                                },
                                "80": {
                                        Name:    "HTTP",
                                        Version: "Nginx 1.14.2",
                                        Vulnerable: false,
                                        Credentials: map[string]string{},
                                },
                                "443": {
                                        Name:    "HTTPS",
                                        Version: "Nginx 1.14.2",
                                        Vulnerable: false,
                                        Credentials: map[string]string{},
                                },
                        },
                        AccessLevel: 0,
                        Type:        "Router",
                },
                {
                        Name: "Database Server",
                        IP:   "192.168.1.10",
                        Ports: map[string]Service{
                                "22": {
                                        Name:    "SSH",
                                        Version: "OpenSSH 8.0",
                                        Vulnerable: false,
                                        Credentials: map[string]string{
                                                "dbadmin": "db$ecureP@ss",
                                        },
                                },
                                "3306": {
                                        Name:    "MySQL",
                                        Version: "8.0.21",
                                        Vulnerable: true,
                                        Credentials: map[string]string{
                                                "root": "mysql2023!",
                                        },
                                },
                        },
                        AccessLevel: 0,
                        Type:        "Database Server",
                },
        }
        
        // Create file systems
        fileSystems := make(map[string]*Directory)
        
        // Root directory
        fileSystems["/"] = &Directory{
                Name:        "/",
                Files:       map[string]string{
                        "welcome.txt": "Welcome to HackSim Terminal.\nYou are on a mission to breach the corporate network.\nUse 'help' command to see available tools.",
                },
                Subdirs:     []string{"home", "etc"},
                Permissions: "rwx",
        }
        
        // Home directory
        fileSystems["/home"] = &Directory{
                Name:        "home",
                Files:       map[string]string{
                        "user_notes.txt": "Remember to change default passwords on all servers.\nNeed to patch the database server vulnerability next week.",
                },
                Subdirs:     []string{},
                Permissions: "rwx",
        }
        
        // Etc directory
        fileSystems["/etc"] = &Directory{
                Name:        "etc",
                Files:       map[string]string{
                        "network.conf": "# Network Configuration\n\nGATEWAY=192.168.1.1\nSUBNET=255.255.255.0\nDNS=8.8.8.8\n\n# Servers\nDBSERVER=192.168.1.10\nWEBSERVER=192.168.1.20\nMAILSERVER=192.168.1.30",
                },
                Subdirs:     []string{},
                Permissions: "r-x",
        }
        
        // Create the scenario
        return Scenario{
                Name:        "Network Breach",
                Description: "Infiltrate a corporate network and access sensitive data",
                Objectives:  objectives,
                NetworkNodes: nodes,
                FileSystems: fileSystems,
                InitialSecurityLevel: 3,
        }
}

// DataHeistScenario creates the Data Heist scenario
func DataHeistScenario() Scenario {
        // Create objectives
        objectives := []Objective{
                {
                        Description: "Scan the target network",
                        Type:        "command",
                        Target:      "scan",
                        Completed:   false,
                },
                {
                        Description: "Connect to the web server at 10.0.1.20",
                        Type:        "connect",
                        Target:      "10.0.1.20",
                        Completed:   false,
                },
                {
                        Description: "Exploit the web server vulnerability",
                        Type:        "command",
                        Target:      "exploit",
                        Completed:   false,
                },
                {
                        Description: "Connect to the file server at 10.0.1.30",
                        Type:        "connect",
                        Target:      "10.0.1.30",
                        Completed:   false,
                },
                {
                        Description: "Find and read the security policy document",
                        Type:        "data",
                        Target:      "security_policy.pdf",
                        Completed:   false,
                },
                {
                        Description: "Download the customer database",
                        Type:        "data",
                        Target:      "customers.db",
                        Completed:   false,
                },
                {
                        Description: "Cover your tracks by clearing log files",
                        Type:        "command",
                        Target:      "clear",
                        Completed:   false,
                },
        }
        
        // Create network nodes
        nodes := []Node{
                {
                        Name: "Entry Point",
                        IP:   "10.0.1.5",
                        Ports: map[string]Service{
                                "22": {
                                        Name:    "SSH",
                                        Version: "OpenSSH 8.2",
                                        Vulnerable: false,
                                        Credentials: map[string]string{
                                                "guest": "guest123",
                                        },
                                },
                        },
                        AccessLevel: 1,
                        Type:        "Workstation",
                },
                {
                        Name: "Web Server",
                        IP:   "10.0.1.20",
                        Ports: map[string]Service{
                                "22": {
                                        Name:    "SSH",
                                        Version: "OpenSSH 7.6",
                                        Vulnerable: false,
                                        Credentials: map[string]string{},
                                },
                                "80": {
                                        Name:    "HTTP",
                                        Version: "Apache 2.4.29",
                                        Vulnerable: true,
                                        Credentials: map[string]string{},
                                },
                                "443": {
                                        Name:    "HTTPS",
                                        Version: "Apache 2.4.29",
                                        Vulnerable: true,
                                        Credentials: map[string]string{},
                                },
                        },
                        AccessLevel: 0,
                        Type:        "Web Server",
                },
                {
                        Name: "File Server",
                        IP:   "10.0.1.30",
                        Ports: map[string]Service{
                                "22": {
                                        Name:    "SSH",
                                        Version: "OpenSSH 7.9",
                                        Vulnerable: false,
                                        Credentials: map[string]string{},
                                },
                                "21": {
                                        Name:    "FTP",
                                        Version: "vsftpd 3.0.3",
                                        Vulnerable: false,
                                        Credentials: map[string]string{},
                                },
                                "445": {
                                        Name:    "SMB",
                                        Version: "Samba 4.9.5",
                                        Vulnerable: true,
                                        Credentials: map[string]string{},
                                },
                        },
                        AccessLevel: 0,
                        Type:        "File Server",
                },
        }
        
        // Create file systems
        fileSystems := make(map[string]*Directory)
        
        // Root directory
        fileSystems["/"] = &Directory{
                Name:        "/",
                Files:       map[string]string{
                        "mission.txt": "MISSION: DATA HEIST\n\nYour objective is to extract customer data from the target company.\nTheir file server contains valuable information that our client needs.\nRemember to cover your tracks when exiting the system.",
                },
                Subdirs:     []string{"var", "mnt"},
                Permissions: "rwx",
        }
        
        // Var directory
        fileSystems["/var"] = &Directory{
                Name:        "var",
                Files:       map[string]string{
                        "log.txt": "System log file\n2023-06-15 12:42:13 - System startup\n2023-06-15 13:15:27 - Security scan completed\n2023-06-15 14:30:01 - Backup completed",
                },
                Subdirs:     []string{},
                Permissions: "rwx",
        }
        
        // Mnt directory
        fileSystems["/mnt"] = &Directory{
                Name:        "mnt",
                Files:       map[string]string{
                        "security_policy.pdf": "[SECURITY POLICY DOCUMENT]\n\nCompany: Acme Corporation\nClassification: CONFIDENTIAL\n\nPasswords must be changed every 30 days.\nTwo-factor authentication required for all admin accounts.\nFile server access restricted to IT department only.\n\nEmergency contact: security@acmecorp.com",
                        "customers.db": "[CUSTOMER DATABASE - ENCRYPTED]\n\nThis file appears to be a database containing customer information.\nIt includes names, addresses, emails, and possibly payment information.\n\nThe data is valuable but needs to be decrypted with proper credentials.",
                },
                Subdirs:     []string{},
                Permissions: "r-x",
        }
        
        // Create the scenario
        return Scenario{
                Name:        "Data Heist",
                Description: "Extract valuable customer data from a corporate file server",
                Objectives:  objectives,
                NetworkNodes: nodes,
                FileSystems: fileSystems,
                InitialSecurityLevel: 5,
        }
}

// SystemTakeoverScenario creates the System Takeover scenario
func SystemTakeoverScenario() Scenario {
        // Create objectives
        objectives := []Objective{
                {
                        Description: "Discover network layout with scan",
                        Type:        "command",
                        Target:      "scan",
                        Completed:   false,
                },
                {
                        Description: "Connect to the control server at 172.16.1.100",
                        Type:        "connect",
                        Target:      "172.16.1.100",
                        Completed:   false,
                },
                {
                        Description: "Locate system credentials file",
                        Type:        "data",
                        Target:      "credentials.enc",
                        Completed:   false,
                },
                {
                        Description: "Crack the encrypted credentials file",
                        Type:        "command",
                        Target:      "crack",
                        Completed:   false,
                },
                {
                        Description: "Exploit vulnerability in SCADA controller",
                        Type:        "command",
                        Target:      "exploit",
                        Completed:   false,
                },
                {
                        Description: "Connect to primary control unit at 172.16.1.200",
                        Type:        "connect",
                        Target:      "172.16.1.200",
                        Completed:   false,
                },
                {
                        Description: "Upload control override software",
                        Type:        "command",
                        Target:      "upload",
                        Completed:   false,
                },
                {
                        Description: "Execute system takeover sequence",
                        Type:        "command",
                        Target:      "execute",
                        Completed:   false,
                },
        }
        
        // Create network nodes
        nodes := []Node{
                {
                        Name: "Access Terminal",
                        IP:   "172.16.1.5",
                        Ports: map[string]Service{
                                "22": {
                                        Name:    "SSH",
                                        Version: "OpenSSH 8.1",
                                        Vulnerable: false,
                                        Credentials: map[string]string{
                                                "operator": "terminal123",
                                        },
                                },
                        },
                        AccessLevel: 1,
                        Type:        "Terminal",
                },
                {
                        Name: "Network Control Server",
                        IP:   "172.16.1.100",
                        Ports: map[string]Service{
                                "22": {
                                        Name:    "SSH",
                                        Version: "OpenSSH 7.5",
                                        Vulnerable: false,
                                        Credentials: map[string]string{},
                                },
                                "80": {
                                        Name:    "HTTP",
                                        Version: "IIS 10.0",
                                        Vulnerable: true,
                                        Credentials: map[string]string{},
                                },
                                "1433": {
                                        Name:    "SQL Server",
                                        Version: "MS SQL 2016",
                                        Vulnerable: false,
                                        Credentials: map[string]string{},
                                },
                        },
                        AccessLevel: 0,
                        Type:        "Control Server",
                },
                {
                        Name: "SCADA Controller",
                        IP:   "172.16.1.150",
                        Ports: map[string]Service{
                                "22": {
                                        Name:    "SSH",
                                        Version: "OpenSSH 6.8",
                                        Vulnerable: true,
                                        Credentials: map[string]string{},
                                },
                                "502": {
                                        Name:    "Modbus",
                                        Version: "Modbus TCP",
                                        Vulnerable: true,
                                        Credentials: map[string]string{},
                                },
                        },
                        AccessLevel: 0,
                        Type:        "SCADA Controller",
                },
                {
                        Name: "Primary Control Unit",
                        IP:   "172.16.1.200",
                        Ports: map[string]Service{
                                "22": {
                                        Name:    "SSH",
                                        Version: "OpenSSH 7.2",
                                        Vulnerable: false,
                                        Credentials: map[string]string{},
                                },
                                "5000": {
                                        Name:    "Control Interface",
                                        Version: "PCU 2.5",
                                        Vulnerable: false,
                                        Credentials: map[string]string{},
                                },
                        },
                        AccessLevel: 0,
                        Type:        "Control Unit",
                },
        }
        
        // Create file systems
        fileSystems := make(map[string]*Directory)
        
        // Root directory
        fileSystems["/"] = &Directory{
                Name:        "/",
                Files:       map[string]string{
                        "takeover.txt": "MISSION: SYSTEM TAKEOVER\n\nObjective: Gain control of the target facility's operational systems.\nThis is a high-security industrial control system.\nProceed with caution - detection will trigger immediate countermeasures.",
                },
                Subdirs:     []string{"sys", "data"},
                Permissions: "rwx",
        }
        
        // Sys directory
        fileSystems["/sys"] = &Directory{
                Name:        "sys",
                Files:       map[string]string{
                        "control.cfg": "# Control System Configuration\n\nMAIN_CONTROLLER=172.16.1.200\nBACKUP_CONTROLLER=172.16.1.201\nSCADA_INTERFACE=172.16.1.150\n\nSECURITY_LEVEL=MAXIMUM\nAUTO_LOCKDOWN=TRUE",
                },
                Subdirs:     []string{},
                Permissions: "r-x",
        }
        
        // Data directory
        fileSystems["/data"] = &Directory{
                Name:        "data",
                Files:       map[string]string{
                        "credentials.enc": "[ENCRYPTED FILE]\n\nThis file contains encrypted system credentials.\nFormat: AES-256\nAccess Level Required: Administrator\n\nDecryption will require advanced tools.",
                        "facility_map.txt": "FACILITY LAYOUT\n\nLevel 1: Administration\nLevel 2: Operations\nLevel 3: Control Room [RESTRICTED]\nLevel 4: Primary Systems [RESTRICTED]\n\nSecurity checkpoints at all level transitions.",
                },
                Subdirs:     []string{},
                Permissions: "r--",
        }
        
        // Create the scenario
        return Scenario{
                Name:        "System Takeover",
                Description: "Gain control of a critical infrastructure facility",
                Objectives:  objectives,
                NetworkNodes: nodes,
                FileSystems: fileSystems,
                InitialSecurityLevel: 8,
        }
}
