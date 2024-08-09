package data

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"

	"gopkg.in/loremipsum.v1"
)

func CreateData() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(dirIndex int) {
			defer wg.Done()

			lorem := loremipsum.NewWithSeed(rand.Int63())
			dirPath := "data/FileData/" + strconv.Itoa(dirIndex)

			if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
				fmt.Printf("Error creating directory %s: %v\n", dirPath, err)
				return
			}

			for fileIndex := 0; fileIndex < 1000; fileIndex++ {
				filePath := fmt.Sprintf("%s/%d_%d", dirPath, dirIndex, fileIndex)
				file, err := os.Create(filePath)

				if err != nil {
					fmt.Printf("Error creating file %s: %v\n", filePath, err)
					continue
				}

				file.WriteString(lorem.Paragraph())
				file.Close()
			}
		}(i)
	}

	wg.Wait()
}
