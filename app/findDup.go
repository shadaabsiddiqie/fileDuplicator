package app

import (
	"crypto/sha256"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type DuplicationContext struct {
	mu     sync.Mutex
	hashes map[string][]string
	fileCh chan string
	doneCh chan bool
	wg     sync.WaitGroup
}

func (ctx *DuplicationContext) hashAndStore() {
	for {
		select {
		case <-ctx.doneCh:
			return
		case filePath, ok := <-ctx.fileCh:
			if !ok {
				return
			}
			data, err := os.ReadFile(filePath)
			if err != nil {
				panic(err)
			}

			hash := sha256.Sum256(data)
			hashKey := string(hash[:])

			ctx.mu.Lock()
			ctx.hashes[hashKey] = append(ctx.hashes[hashKey], filePath)
			ctx.mu.Unlock()

			ctx.wg.Done()
		}
	}
}

func (ctx *DuplicationContext) startConsumers(workers int) {
	for i := 0; i < workers; i++ {
		go ctx.hashAndStore()
	}
}

func (ctx *DuplicationContext) producerScan(rootDir string) error {
	err := filepath.Walk(rootDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			ctx.wg.Add(1)
			ctx.fileCh <- path
		}
		return nil
	})

	ctx.wg.Wait()
	close(ctx.doneCh)
	return err
}

func printDuplicates(hashes map[string][]string) {
	for _, files := range hashes {
		if len(files) > 1 {
			fmt.Println("Duplicate files:", files)
		}
	}
}

func RunFindDup(rootDir string) {
	ctx := &DuplicationContext{
		hashes: make(map[string][]string),
		fileCh: make(chan string, 300),
		doneCh: make(chan bool),
	}

	start := time.Now()

	go ctx.producerScan(rootDir)
	ctx.startConsumers(1000)

	<-ctx.doneCh
	close(ctx.fileCh)

	printDuplicates(ctx.hashes)

	fmt.Printf("Total unique hashes: %d\n", len(ctx.hashes))
	fmt.Printf("Time taken: %.2f seconds\n", time.Since(start).Seconds())
}
