package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"path/filepath"
)
	var batchFiles []string

func main() {
	fileCount := 0
	batchCount := 1
	const batchSize = 120
	var currentBatch []string

	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".MP4" {
			currentBatch = append(currentBatch, file.Name())
			fileCount++

			if fileCount >= batchSize {
				writeBatchToFile(currentBatch)
				batchCount++
				fileCount = 0
				currentBatch = nil // Reset for the next batch
			}
		}
	}

	// Write any remaining files
	if len(currentBatch) > 0 {
		writeBatchToFile(currentBatch)
	}
	createVids()
}

func writeBatchToFile(batch []string) {
	// filename := fmt.Sprintf("batch_%d.txt", batchNum)
	firstDate := getVidTime(batch[0])
	filename := fmt.Sprintf("%s_trailvid.txt", firstDate)
	batchFiles = append(batchFiles, filename)
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	for _, filename := range batch {
		_, err := file.WriteString("file " + "'" + filename + "'" + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	fmt.Printf("Wrote %d filenames to %s\n", len(batch), filename)
}

func createVids() {
	for _, batchfile := range batchFiles {
		cmd := exec.Command("ffmpeg", "-f", "concat", "-safe", "0", "-i", batchfile, "-c", "copy", strings.TrimSuffix(batchfile, "txt") + "mp4")

		fmt.Println("Executing Command:", cmd.String())

		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Error executing ffmpeg:", err)
			fmt.Println("ffmpeg output:", string(out)) // Print output for debugging
		return
		}

	fmt.Println("ffmpeg output:", string(out))
	fmt.Println("Successfully concatenated videos!")
	}
}

func getVidTime(f string) string {
	
	fileInfo, err := os.Stat(f)
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return "Error"
	}

	creationTime := fileInfo.ModTime() // Get the last modified time
	date := creationTime.Format("2006-01-02_15-04-05")
	return date

}
