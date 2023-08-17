package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// -i = imagens
	// -v = videos
	// -f = fps

	var (
		imgPath   string
		folPath   string
		fpsFlag   float64
		threshold int
	)

	flag.StringVar(&imgPath, "i", "", "Image path")
	flag.StringVar(&folPath, "v", "", "Folder path")
	flag.Float64Var(&fpsFlag, "f", 30.0, "Video FPS")
	flag.IntVar(&threshold, "t", 127, "Threshold")

	flag.Parse()

	if imgPath != "" && folPath != "" {
		fmt.Println("Error: can't use -i and -v togheter")
		os.Exit(1)
	}

	if folPath != "" {
		fileInfo, err := os.Stat(folPath)
		if err != nil {
			fmt.Println("Can't read directory")
			os.Exit(1)
		}

		if fileInfo.IsDir() {
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
