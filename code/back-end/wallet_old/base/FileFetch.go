package base

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	// Use mutexes to ensure thread-safe access to paths
	mu sync.Mutex
	// The path to the storage file
	filePath string
)

// SetFilePath
// Set the file path and perform some validation at setup time
func SetFilePath(path string) error {
	mu.Lock()
	defer mu.Unlock()
	//fmt.Printf("path:%v\n", path)
	// Here you can add the path validation logic
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist at path: %s", path)
	}
	// Let's say it's just reading the file
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("error opening file: %s", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("error closing file")
		}
	}(file)
	filePath = path
	return nil
}

// GetFilePath
// Get the path to the stored file
func GetFilePath() string {
	mu.Lock()
	defer mu.Unlock()
	return filePath
}

func CreateDir(path string) {
	folderPath := filepath.Join(path, "/example")
	err := os.Mkdir(folderPath, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
	} else {
		fmt.Println("Directory created successfully")
	}

	// Create a directory and all of its parent directories
	folderPath1 := filepath.Join(path, "/example/dir1/dir2")
	err = os.MkdirAll(folderPath1, 0755)
	if err != nil {
		fmt.Println("Error creating nested directories:", err)
	} else {
		fmt.Println("Nested directories created successfully")
	}
}

func CreateDir2(path string) {
	folderPath := filepath.Join(path, "/.lit")
	folderPath1 := filepath.Join(folderPath, "/logs")
	err := os.Mkdir(folderPath1, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
	} else {
		fmt.Println("Directory created successfully")
	}

	// Create a directory and all of its parent directories
	folderPath2 := filepath.Join(path, "/.lit")
	folderPath21 := filepath.Join(folderPath2, "/data/logs")
	err = os.MkdirAll(folderPath21, 0755)
	if err != nil {
		fmt.Println("Error creating nested directories:", err)
	} else {
		fmt.Println("Nested directories created successfully")
	}
}

func FileConfig(path string) bool {
	filePath := "example.txt"
	folderPath := path + string(os.PathSeparator) + filePath
	fmt.Printf("filePath: %s\n", folderPath)
	file, err := os.Create(folderPath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return false
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("error closing file")
		}
	}(file)
	// Write the string to a file
	_, err = file.WriteString("Hello, world!\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return false
	}
	fmt.Println("File written successfully")
	return true
}

func ReadConfig(path string) string {
	filePath := "example.txt"
	folderPath := path + string(os.PathSeparator) + filePath
	f, err := os.Open(folderPath)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Printf("f.Close err: %v\n", err)
		}
	}(f)
	if err != nil {
		fmt.Printf("open file err: %v\n", err)
	}
	r := bufio.NewReader(f)
	b, _, err := r.ReadLine()
	if err != nil {
		fmt.Printf("read file err: %v\n", err)
		return ""
	}
	s := strings.TrimSpace(string(b))
	fmt.Printf("%v\n", s)
	return s
}

func ReadConfig1(path string) bool {
	filePath := "config.txt"
	folderPath := path + string(os.PathSeparator) + filePath
	f, err := os.Open(folderPath)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Printf("f.Close err: %v\n", err)
		}
	}(f)
	if err != nil {
		fmt.Printf("open file err: %v\n", err)
		return false
	}
	fmt.Printf("%v\n", ReadConfigFile(folderPath))
	return true
}

func ReadConfig2(path string) bool {
	filePath := "config.txt"
	appName := "lit"
	folderPath := path + string(os.PathSeparator) + filePath
	fileConfig := ReadConfigFile(folderPath)
	dirPath := fileConfig["dirpath"]
	fmt.Printf("read dirpath is :%v\n", dirPath)
	folderPath1 := filepath.Join(dirPath, "."+appName)
	fmt.Printf("read folderPath1 is :%v\n", folderPath1)
	return true
}

func visit(path string, f os.FileInfo, err error) error {
	if err != nil {
		return err // 处理潜在的错误
	}
	// Check if it's a directory or a file
	fileType := "File"
	if f.IsDir() {
		fileType = "Directory"
	}
	// Obtain permission to print a file or directory
	// Go back to the section on file permissions
	perms := f.Mode().Perm()

	fmt.Printf("%s: %s, Permissions: %04o\n", fileType, path, perms)
	return nil
}

func VisitAll() {
	root := GetFilePath()
	// Use filepath.Walk to traverse the directory
	err := filepath.Walk(root, visit)
	if err != nil {
		fmt.Printf("Error walking the path %q: %v\n", root, err)
	}
}
