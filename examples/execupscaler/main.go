package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"time"

	"github.com/AlecAivazis/survey/v2"
)

func main() {

	// log cwd
	rootdir, error := os.Getwd()
	if error != nil {
		fmt.Println(error)
		return
	}

	bin := filepath.Join(rootdir, "bin", "realesrgan")
	bin = filepath.FromSlash(bin)
	print(bin)

	modelsPath := filepath.FromSlash(filepath.Join(rootdir, "models"))

	outputDir := filepath.FromSlash(filepath.Join(rootdir, "output"))

	var inputDir string
	fmt.Print("Please input the file dir: ")
	fmt.Scanln(&inputDir)

	// 读取用户输入的目录路径
	var inputMode string

	prompt := &survey.Select{
		Message: "Please choose the mode:",
		Options: []string{"parallel", "sequential"},
	}
	survey.AskOne(prompt, &inputMode)

	if inputMode == "parallel" {
		// 创建一个chan，用于存储需要处理的图像文件路径
		fileChan := make(chan string)

		// 开启多个goroutine来处理图像
		for i := 0; i < 4; i++ {
			go func() {
				for path := range fileChan {
					// 生成随机文件名
					randStr := fmt.Sprintf("%d", time.Now().Unix())
					outputFile := filepath.Join(outputDir, randStr+".jpg")

					// 调用realesrgan进行图像处理
					args := []string{"-i", path, "-o", outputFile, "-s", "4", "-m", modelsPath, "-n", "realesrgan-x4plus"}
					cmd := exec.Command(bin, args...)
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr
					err := cmd.Run()

					if err != nil {
						fmt.Println("Command execution failed:", err)
					}
				}
			}()
		}

		// 遍历输入目录下的所有图像文件，并将它们存储到fileChan中
		files, err := os.ReadDir(inputDir)
		if err != nil {
			fmt.Println("Error reading input directory:", err)
			return
		}

		for _, file := range files {
			if !file.IsDir() {
				filePath := filepath.Join(inputDir, file.Name())
				fileChan <- filePath
			}
		}

		// 关闭chan，等待所有goroutine处理完毕
		close(fileChan)
		wg := sync.WaitGroup{}
		wg.Add(4)
		for i := 0; i < 4; i++ {
			go func() {
				defer wg.Done()
				for range fileChan {
				}
			}()
		}
		wg.Wait()

	} else if inputMode == "sequential" {
		// 遍历输入目录下的所有图像文件，并依次进行处理
		files, err := os.ReadDir(inputDir)
		if err != nil {
			fmt.Println("Error reading input directory:", err)
			return
		}

		for _, file := range files {
			if !file.IsDir() {
				filePath := filepath.Join(inputDir, file.Name())

				// 生成随机文件名

				randStr := fmt.Sprintf("%d", time.Now().Unix())
				outputFile := filepath.Join(outputDir, randStr+".jpg")

				// 调用realesrgan进行图像处理
				args := []string{"-i", filePath, "-o", outputFile, "-s", "4", "-m", modelsPath, "-n", "realesrgan-x4plus"}
				cmd := exec.Command(bin, args...)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err := cmd.Run()

				if err != nil {
					fmt.Println("Command execution failed:", err)
				}

			}
		}
	} else {
		fmt.Println("Invalid mode input, please input 'parallel' or 'sequential'.")
		return
	}

	fmt.Println("All images processed successfully!")
}
