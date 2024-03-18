package base

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestReadConfigFile(t *testing.T) {
	config := ReadConfigFile("D:\\wallet\\api\\file\\babbab.txt")
	dirpath := config["dirpath"]
	folderPath := filepath.Join(dirpath, ".lnd")
	fmt.Print(folderPath)
}
