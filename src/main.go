package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
	"time"
)

func main() {
	// -i = imagens
	// -v = videos
	// -f = fps

	var (
		imgPath    string
		folPath    string
		fpsFlag    float64
		threshold  int
		audioPath  string
		audioDelay float64
	)

	flag.StringVar(&imgPath, "i", "", "Image path")
	flag.StringVar(&folPath, "v", "", "Folder path")
	flag.Float64Var(&fpsFlag, "f", 30.0, "Video FPS")
	flag.IntVar(&threshold, "t", 127, "Threshold")
	flag.StringVar(&audioPath, "a", "", "Audio path")
	flag.Float64Var(&audioDelay, "d", -1, "Audio delay in milliseconds")

	flag.Parse()

	if imgPath != "" && folPath != "" {
		fmt.Println("Error: can't use -i and -v togheter")
		os.Exit(1)
	}

	if audioPath != "" {
		fileInfo, err := os.Stat(audioPath)
		if err != nil {
			fmt.Println("Can't read directory")
			os.Exit(1)
		}

		if fileInfo.IsDir() {
			fmt.Println("Can't pass directory as file")
			os.Exit(1)
		}

		if audioDelay == -1 {
			fmt.Println("Need to set the delay flag")
			os.Exit(1)
		}

	}

	if folPath != "" {
		fileInfo, err := os.Stat(folPath)
		if err != nil {
			fmt.Println("Can't read directory")
			os.Exit(1)
		}

		if fileInfo.IsDir() {
			if audioPath != "" {
				_, err := os.Open(audioPath)
				if err != nil {
					log.Fatal(err)
				}

				command := "mpg123" // Substitua "ls" pelo comando que você deseja verificar

				// Verifica se o comando está disponível no sistema
				_, err = exec.LookPath(command)
				if err != nil {
					fmt.Printf("É necessário ter o comando mpg123 instalado para usar essa feature\n")
					return
				}

				var wg sync.WaitGroup
				wg.Add(2)
				go func() {
					time.Sleep(time.Second)
					defer wg.Done() // Decrementa o contador do WaitGroup quando a goroutine terminar
					play_audio(audioPath)
				}()

				go func() {
					time.Sleep(time.Duration(audioDelay * float64(time.Millisecond)))
					defer wg.Done() // Decrementa o contador do WaitGroup quando a goroutine terminar
					play_video(folPath, threshold, fpsFlag)
				}()

				wg.Wait()

			}
			play_video(folPath, threshold, fpsFlag)
		} else {
			fmt.Println("Can't pass file as directory")
		}
	}

	if imgPath != "" {
		fileInfo, err := os.Stat(imgPath)
		if err != nil {
			fmt.Println("Can't read directory")
			os.Exit(1)
		}

		if fileInfo.IsDir() {
			fmt.Println("Can't pass directory as file")
		} else {

			fmt.Print(generate(imgPath, threshold))
			//generate(imgPath, threshold)
		}
	}

}

func play_audio(audioPath string) {
	cmd := exec.Command("mpg123", audioPath)
	_, err := cmd.Output()
	if err != nil {
		fmt.Printf("Erro ao executar o comando: %s\n", err)
		return
	}
}
