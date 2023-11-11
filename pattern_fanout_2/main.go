package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/nfnt/resize"
)

type Task struct {
	FilePath string
}

func main() {
	// List of image file paths to process
	imagePaths := []string{"image1.png", "image2.png", "image3.png", "image4.png", "image5.png", "image6.png", "image7.png"}

	numWorkers := 4
	taskCh := make(chan Task, len(imagePaths))
	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg, taskCh)
	}

	fmt.Println("Sleeping after spawning workers")
	time.Sleep(2 * time.Second)

	// Send tasks to the workers
	for _, path := range imagePaths {
		taskCh <- Task{FilePath: path}
	}
	close(taskCh)

	// Wait for all workers to finish
	wg.Wait()
	fmt.Println("All images processed.")
}

func worker(wId int, wg *sync.WaitGroup, taskCh <-chan Task) {
	fmt.Printf("Worker %v is ready!\n", wId)
	defer wg.Done()
	for task := range taskCh {
		fmt.Printf("Task [%v] is picked by worker [%v]\n", task.FilePath, wId)
		processImage(task)
	}
}

func processImage(task Task) {
	fmt.Printf("Started processing task [%v]\n", task.FilePath)
	time.Sleep(1 * time.Second)
	img, err := loadImage(task.FilePath)
	if err != nil {
		fmt.Printf("Failed to load image %s: %v\n", task.FilePath, err)
		return
	}

	// Resize image to 100x100 pixels
	resizedImg := resize.Resize(100, 100, img, resize.Lanczos3)

	// Save the resized image
	savePath := filepath.Join("resized", task.FilePath)
	if err := saveImage(resizedImg, savePath); err != nil {
		fmt.Printf("Failed to save image %s: %v\n", savePath, err)
	}
	fmt.Printf("Processed image: %s\n", task.FilePath)
}

func loadImage(filePath string) (image.Image, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	return img, err
}

func saveImage(img image.Image, filePath string) error {
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	fileExtension := strings.ToLower(filepath.Ext(filePath))
	switch fileExtension {
	case ".jpg", ".jpeg":
		return jpeg.Encode(out, img, nil)
	case ".png":
		return png.Encode(out, img)
	default:
		return fmt.Errorf("unsupported file extension: %s", fileExtension)
	}
}
