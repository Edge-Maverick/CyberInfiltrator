package game

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
)

// FileSystem represents a simulated file system
type FileSystem struct {
	Directories map[string]*Directory
	CurrentPath string
}

// Directory represents a directory in the file system
type Directory struct {
	Name     string
	Files    map[string]string // filename -> content
	Subdirs  []string          // names of subdirectories
	Permissions string         // simplified permissions like "rwx", "r--", etc.
}

// NewFileSystem creates a new file system with the given directories
func NewFileSystem(dirs map[string]*Directory) FileSystem {
	if dirs == nil {
		dirs = make(map[string]*Directory)
		
		// Create root directory if none provided
		dirs["/"] = &Directory{
			Name:     "/",
			Files:    make(map[string]string),
			Subdirs:  []string{},
			Permissions: "rwx",
		}
	}
	
	return FileSystem{
		Directories: dirs,
		CurrentPath: "/",
	}
}

// ListFiles lists files and directories in the specified path
func (fs *FileSystem) ListFiles(path string) string {
	// Resolve path
	fullPath, err := fs.resolvePath(path)
	if err != nil {
		return "Error: " + err.Error()
	}
	
	// Get directory
	dir, exists := fs.Directories[fullPath]
	if !exists {
		return "Error: Directory not found: " + fullPath
	}
	
	// Format listing
	var result []string
	result = append(result, fmt.Sprintf("Directory listing of %s:", fullPath))
	result = append(result, "")
	
	// Add directories
	if len(dir.Subdirs) > 0 {
		for _, subdir := range dir.Subdirs {
			subdirPath := filepath.Join(fullPath, subdir)
			if subdir, exists := fs.Directories[subdirPath]; exists {
				result = append(result, fmt.Sprintf("d%s  %s/", subdir.Permissions, subdir.Name))
			}
		}
	}
	
	// Add files
	if len(dir.Files) > 0 {
		for filename := range dir.Files {
			result = append(result, fmt.Sprintf("-rw-  %s", filename))
		}
	}
	
	// If empty directory
	if len(dir.Subdirs) == 0 && len(dir.Files) == 0 {
		result = append(result, "<empty directory>")
	}
	
	return strings.Join(result, "\n")
}

// ReadFile reads the content of a file
func (fs *FileSystem) ReadFile(filename string) (string, error) {
	// If it's a full path, split it
	dir, file := filepath.Split(filename)
	if dir == "" {
		dir = fs.CurrentPath
	}
	
	// Resolve the directory path
	fullPath, err := fs.resolvePath(dir)
	if err != nil {
		return "", err
	}
	
	// Get directory
	directory, exists := fs.Directories[fullPath]
	if !exists {
		return "", errors.New("directory not found: " + fullPath)
	}
	
	// Get file content
	content, exists := directory.Files[file]
	if !exists {
		return "", errors.New("file not found: " + file)
	}
	
	return content, nil
}

// resolvePath resolves a path to its absolute form
func (fs *FileSystem) resolvePath(path string) (string, error) {
	if path == "" {
		return fs.CurrentPath, nil
	}
	
	// If it's already absolute
	if strings.HasPrefix(path, "/") {
		return filepath.Clean(path), nil
	}
	
	// It's relative, join with current path
	return filepath.Clean(filepath.Join(fs.CurrentPath, path)), nil
}
