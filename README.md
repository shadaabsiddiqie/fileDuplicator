# File Duplication Finder

## Overview
The File Duplication Finder is a Go application designed to identify and manage duplicate files in a specified directory. It scans files in a directory, calculates their SHA-256 hashes, and identifies duplicates based on these hashes. This tool can be useful for cleaning up storage and organizing files.

## Features
- **File Scanning**: Recursively scans a specified directory for files.
- **Duplicate Detection**: Calculates SHA-256 hashes for each file and identifies duplicates.
- **Data Generation**: Optionally creates test data with random content for testing purposes.

## Installation

1. **Clone the repository**:
   ```bash
   git clone <repository-url>
   ```
2. **Install Dependencies**:
   This project does not have any external dependencies except Go itself.

3. **Run the Application**:
   ```bash
   go run main.go
   ```

## Project Structure

- **app**: Contains the core functionality to find duplicate files.
- **data**: Contains functions to create test data and copy files to random folders for testing.

## Functions

### app

#### `RunFindDup(rootDir string)`
Runs the entire process to find duplicate files in the specified directory.

- **Parameters**:
  - `rootDir`: The directory to be scanned for duplicates.

#### `hashAndStore()`
Hashes each file's content and stores it in a map for comparison.

#### `startConsumers(workers int)`
Starts the specified number of workers that hash and store files.

- **Parameters**:
  - `workers`: The number of worker goroutines to start.

#### `producerScan(rootDir string) error`
Recursively scans the specified directory, sending file paths to the consumer channel.

- **Parameters**:
  - `rootDir`: The root directory to start scanning from.

### data

#### `CreateData()`
Creates random text files in multiple directories for testing purposes. Each directory contains 1000 files with randomly generated content.

#### `AddCopyData()`
Copies a random selection of files from a source folder to a randomly chosen target folder, avoiding copying files back to the same source folder.

#### `copyDataToRandomFolder(sourceFolder int, wg *sync.WaitGroup)`
Copies a random selection of files from the source folder to a randomly chosen target folder. Avoids copying files back to the same source folder.

- **Parameters**:
  - `sourceFolder`: The folder from which files will be copied.
  - `wg`: A pointer to a `sync.WaitGroup` to signal when the operation is complete.

#### `randWithoutX(size int, x int) int`
Generates a random integer between 0 and `size - 1` but ensures the value is not `x`.

- **Parameters**:
  - `size`: The upper limit for the random number.
  - `x`: The number to exclude from the result.

## Usage

1. **Create Test Data**:
   ```go
   data.CreateData()
   ```
   This will generate 5 folders, each containing 1000 files with random content.

2. **Copy Data**:
   ```go
   data.AddCopyData()
   ```
   This will copy random files to random folders to create duplicates.

3. **Find Duplicates**:
   ```go
   app.RunFindDup("<your-directory-path>")
   ```
   This will scan the specified directory and print the duplicate files found.

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
