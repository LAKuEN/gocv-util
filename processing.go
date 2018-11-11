package gocvutil

import (
	"fmt"
	"image"

	"gocv.io/x/gocv"
)

// processing.go はMatを加工する関数をまとめたモジュールです。
// TODO テストコード

// Mozaic はMat全体に対してモザイク処理を掛けます。
func Mozaic(m *gocv.Mat, blockSize image.Point) error {
	if m.Size()[0] < blockSize.Y || m.Size()[1] < blockSize.X {
		return fmt.Errorf("blockSize is bigger than matSize")
	}
	for _, row := range ToBlocks(m, blockSize) {
		for _, block := range row {
			block.SetTo(block.Mean())
		}
	}

	return nil
}
