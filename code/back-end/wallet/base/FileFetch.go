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
	// 使用互斥锁来确保对路径的线程安全访问
	mu sync.Mutex
	// 存储文件路径
	filePath string
)

// SetFilePath 设置文件路径，并可在设置时执行某些验证
func SetFilePath(path string) error {
	mu.Lock()
	defer mu.Unlock()
	fmt.Printf("path:%v\n", path)
	// 这里可以添加路径验证逻辑
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist at path: %s", path)
	}
	// 假设只是读取文件
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("error opening file: %s", err)
	}
	defer file.Close()
	filePath = path
	return nil
}

// GetFilePath 获取存储的文件路径
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

	// 创建目录及其所有父目录
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

	// 创建目录及其所有父目录
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
	defer file.Close() // 确保在函数结束时关闭文件

	// 写入字符串到文件中
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
	// 检查是目录还是文件
	fileType := "File"
	if f.IsDir() {
		fileType = "Directory"
	}
	// 获取并打印文件或目录的权限
	perms := f.Mode().Perm() // 返回文件权限的部分

	fmt.Printf("%s: %s, Permissions: %04o\n", fileType, path, perms)
	return nil
}

func VisitAll() {
	root := GetFilePath()
	// 使用 filepath.Walk 遍历目录
	err := filepath.Walk(root, visit)
	if err != nil {
		fmt.Printf("Error walking the path %q: %v\n", root, err)
	}
}
