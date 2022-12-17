package day17

import (
	"strconv"
	"time"
)

var num int

func (g *game) debugStage() {
	// dc.DrawPoint(3, 10000, 0)
	// dc.SavePNG("day17.png")
	// dc.DrawCircle(500, 500, 400)
	// dc.DrawPoint(3, 10000, 0)
	dc.SetRGBA255(255, 255, 255, 255)
	dc.DrawRectangle(0, 0, 100, 100_000)
	dc.Fill()
	g.drawStage()
	if g.hasRock {
		g.drawRock()
	}
	num++
	// filename := fmt.Sprintf("./day17/pic/day17_%v.png", num)
	// filename := fmt.Sprintf("day17_%v.png", num)
	filename := "day17.png"
	// fmt.Println(filename)
	dc.SavePNG(filename)
	time.Sleep(100 * time.Millisecond)
}

func (g *game) drawStage() {
	for j := g.height - 7; j < 10_000; j++ {
		// fmt.Println(g.stage[j], " y: ", j)
		dc.SetRGBA255(245, 142, 65, 255)
		dc.DrawString(strconv.Itoa(j), 70, float64(j)*10+10)
		dc.Fill()
		for i := 0; i < 7; i++ {
			if g.stage[j][i] == 2 {
				dc.SetRGBA255(245, 142, 65, 255)
				dc.DrawRectangle(float64(i)*10, float64(j)*10, 9, 9)
				dc.Fill()
			}
			if g.stage[j][i] == 3 {
				dc.SetRGB255(105, 105, 105)
				dc.DrawRectangle(float64(i)*10, float64(j)*10, 9, 9)
				dc.Fill()
			}
		}
	}
}

func (g *game) drawRock() {
	rockX := g.rock1.position[0]
	rockY := g.rock1.position[1]
	dc.SetRGBA255(245, 0, 0, 255)
	for j := 0; j < 4; j++ {
		for i := 0; i < 4; i++ {
			if g.rock1.body[j][i] == 1 {
				dc.DrawRectangle(float64(i+rockX)*10, float64(j+rockY)*10, 9, 9)
			}
		}
	}
	dc.Fill()

}
