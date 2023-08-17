package functions

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func Play_video(folderPath string, threshold int, targetFPS float64) {

	callCount := 0

	filePaths, err := os.ReadDir(folderPath)
	if err != nil {
		fmt.Println("Failed reading files:", err)
		return
	}

	frameDuration := time.Second / time.Duration(targetFPS)

	for {
		if callCount >= len(filePaths) {
			break
		}

		startTime := time.Now()

		filePath := filepath.Join(folderPath, filePaths[callCount].Name())
		//fmt.Printf("Chamada da função #%d com arquivo: %s\n", callCount+1, filePath)

		frame := Generate(filePath, threshold) // Passe o caminho do arquivo como parâmetro
		clearTerminal()
		fmt.Print(frame)

		elapsedTime := time.Since(startTime)
		sleepDuration := frameDuration - elapsedTime

		if sleepDuration > 0 {
			time.Sleep(sleepDuration)
		}

		callCount++
	}
}

func clearTerminal() {
	cmd := exec.Command("clear") // Para Windows, use "cls"
	cmd.Stdout = os.Stdout
	cmd.Run()
}
