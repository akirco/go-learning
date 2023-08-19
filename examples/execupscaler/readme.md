## go cli example

```go
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

func main() {

	// log cwd
	rootdir, error := os.Getwd()
	if error != nil {
		fmt.Println(error)
		return
	}

	bin :=path.Join(rootdir,"bin","realesrgan")
	bin = filepath.FromSlash(bin)

	modelsPath := filepath.FromSlash(path.Join(rootdir,"models"))

	inputFile := filepath.FromSlash(path.Join(rootdir,"input","1.jpg"))
	outputFile := filepath.FromSlash(path.Join(rootdir,"input","1-clear.jpg"))

  args := []string{"-i", inputFile, "-o", outputFile,"-s","4","-m",modelsPath, "-n", "realesrgan-x4plus"}

	// 创建一个Cmd对象
	cmd := exec.Command(bin,args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err :=cmd.Run()
	if err !=nil{
		fmt.Println("Command execution faild:",err)
		os.Exit(1)
	}
}

```
