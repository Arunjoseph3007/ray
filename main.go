package main

import (
	"fmt"
	"log"
	"os"
	"ray-tracing/vec"
)

func main() {
	ratio := 2.0
	height := 100
	width := int(float64(height) * ratio)

	out := fmt.Sprintf("P3\n%d %d\n255\n", width, height)

	for i := 0; i < height; i++ {
		fmt.Printf("\rRemaining: %f /", float64(i)/float64(height)*100)
		for j := 0; j < width; j++ {
			v := vec.New(float64(i)/float64(height)*255, float64(j)/float64(width)*255, 0)
			out += v.ToStr()
		}
	}

	file, err := os.Create("image.ppm")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err2 := file.WriteString(out)
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("Done!")
}
