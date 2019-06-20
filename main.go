package main

import (
	"flag"
	"log"
	"math/rand"
	"time"

	"github.com/dhconnelly/rtreego"
)

var (
	numCellsXY = flag.Int("dim", 1000, "Number of cells along X and Y axis")
)

type box struct {
	box *rtreego.Rect
}

func (b *box) Bounds() *rtreego.Rect {
	return b.box
}

func (b *box) String() {
	return
}

func main() {
	flag.Parse()

	numCells := *numCellsXY * *numCellsXY

	log.Printf("Making array of size %d", numCells)
	start := time.Now()
	data := make([]rtreego.Spatial, numCells)
	elapsed := time.Since(start)
	log.Printf("Done making array elapsed = %s", elapsed)

	log.Printf("Filling array")
	count := 0
	start = time.Now()
	for i := 0; i < *numCellsXY; i++ {
		for j := 0; j < *numCellsXY; j++ {
			r, _ := rtreego.NewRect(rtreego.Point{float64(i), float64(j)}, []float64{1.0, 1.0})
			data[count] = &box{box: r}
			count++
		}
	}
	elapsed = time.Since(start)
	log.Printf("Filled array elapsed = %s", elapsed)
	log.Printf("Inserting into tree")
	start = time.Now()
	rt := rtreego.NewTree(2, 25, 50, data[:numCells]...)
	elapsed = time.Since(start)
	log.Printf("Done inserting into tree elapsed = %s", elapsed)

	for i := 0; i < 100; i++ {
		r, _ := rtreego.NewRect(rtreego.Point{rand.Float64() * float64(*numCellsXY), rand.Float64() * float64(*numCellsXY)}, []float64{0.00000000001, 0.00000000001})
		start := time.Now()
		match := rt.SearchIntersectWithLimit(1, r)
		elapsed := time.Since(start)
		log.Printf("elapsed = %s, r = %s match = %s", elapsed, r, match[0].Bounds().String())
	}
}
