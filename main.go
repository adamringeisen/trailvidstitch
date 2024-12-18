package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)
var batchFiles []string
var lastBatchFileName string

func main() {
	lastVid := isLastVidAppend()
	fileCount := 0
	batchCount := 1
	batchSize := 120
	var currentBatch []string

	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".MP4" || filepath.Ext(file.Name()) == ".mp4"  {
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
		if lastVid {
			writeLastBatchToFile(currentBatch, lastBatchFileName)
			
		} else {
			writeBatchToFile(currentBatch)
		}
		
	}
	createVids()
}

func writeLastBatchToFile(batch []string, filename string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("Error opening last batch file:", err)
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
func writeBatchToFile(batch []string) {
	// filename := fmt.Sprintf("batch_%d.txt", batchNum)
	firstDate := getVidTime(batch[0])
	filename := fmt.Sprintf("%s_trailvid.txt", firstDate)
	lastBatchFileName = filename
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
	for i, batchfile := range batchFiles {
		cmd := exec.Command("ffmpeg", "-f", "concat", "-safe", "0", "-i", batchfile, "-c", "copy", strings.TrimSuffix(batchfile, "txt") + "mp4")

		fmt.Printf("Creating Video %d of %d...\n", i+1, len(batchFiles))

		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Error executing ffmpeg:", err)
			fmt.Println("ffmpeg output:", string(out)) // Print output for debugging
		return
		}

	// fmt.Println("ffmpeg output:", string(out))
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

func isLastVidAppend() bool {
	fileCount := 0
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println(err)
	}
	
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".MP4" || filepath.Ext(file.Name()) == ".mp4" {
			fileCount++
			}
	}
	if fileCount == 0 {
		fmt.Println("There are no mp4 files in this directory")
		os.Exit(1)
	}
	if fileCount % 120  > 30 {
		// More than 15 minutes on last vid is fine
		return false
	} else {
		return true
	}
}
