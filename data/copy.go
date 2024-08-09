package data

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
)

func randWithoutX(size int, x int) int {
	ans := rand.Intn(size)
	if ans == x {
		return randWithoutX(size, x)
	}
	return ans
}

// copyDataToRandomFolder copies a random selection of files from the source folder to a randomly chosen target folder.
// It avoids copying files back to the same source folder. The function reads each selected file, then writes it
// to a new location in the target folder. The number of files copied ranges between 300 and 600.
//
// Parameters:
// - sourceFolder: The folder from which files will be copied.
// - wg: A pointer to a sync.WaitGroup to signal when the operation is complete.
func copyDataToRandomFolder(sourceFolder int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Determine the number of files to copy, between 300 and 600.
	fileCount := rand.Intn(300) + 300

	for i := 0; i < fileCount; i++ {
		// Construct the file path in the source folder.
		filePath := fmt.Sprintf("FileData/%d/%d_%d", sourceFolder, sourceFolder, i)

		// Read the file's data.
		data, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", filePath, err)
			continue
		}

		// Choose a random target folder different from the source folder.
		targetFolder := randWithoutX(5, sourceFolder)
		// Construct the file path in the target folder.
		targetPath := fmt.Sprintf("data/FileData/%d/%d_%d", targetFolder, sourceFolder, i)

		// Write the data to the target file.
		if err := os.WriteFile(targetPath, data, os.ModeAppend); err != nil {
			fmt.Printf("Error writing to file %s: %v\n", targetPath, err)
		}
	}
}

func AddCopyData() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go copyDataToRandomFolder(i, &wg)
	}

	wg.Wait()
}
