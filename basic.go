package gocvutil

import (
	"image"

	"gocv.io/x/gocv"
)

// basic.go は中間的な処理のみを実装したモジュールです。
// TODO テストコード

// ToBlocks はMatを小領域に分割した配列を返します。
// 小領域は元のMatへの参照を持つので、小領域に対する操作は元のMatに反映されます。
func ToBlocks(m *gocv.Mat, size image.Point) [][]*gocv.Mat {
	var blocks [][]*gocv.Mat
	for r := 0; r < m.Rows(); r = r + 1 {
		var row []*gocv.Mat
		maxY := (r + 1) * size.Y
		if maxY > m.Rows()-2 {
			maxY = m.Rows() - 1
		}

		for c := 0; c < m.Cols(); c = c + 1 {
			maxX := (c + 1) * size.X
			if maxX > m.Cols()-2 {
				maxX = m.Cols() - 1
			}
			rect := image.Rectangle{
				Min: image.Point{X: c * size.X, Y: r * size.Y},
				Max: image.Point{X: maxX, Y: maxY},
			}
			reg := m.Region(rect)
			row = append(row, &reg)

			if rect.Max.X > m.Cols()-2 {
				break
			}
		}
		blocks = append(blocks, row)
		if maxY > m.Rows()-2 {
			break
		}
	}

	return blocks
}
