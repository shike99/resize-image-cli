package main

import (
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strconv"

	"github.com/nfnt/resize"
)

func main() {
	resizeImage()
}

func resizeImage() {
	// 元の画像の読み込み
	filepath := os.Args[1]
	fileData, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}

	// 画像をimage.Image型にdecodeします
	img, data, err := image.Decode(fileData)
	if err != nil {
		log.Fatal(err)
	}
	fileData.Close()

	// ここでリサイズします
	// 片方のサイズを0にするとアスペクト比固定してくれます
	width, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatal(err)
	}
	height, err := strconv.Atoi(os.Args[4])
	if err != nil {
		log.Fatal(err)
	}
	w := &width
	h := &height
	resizedImg := resize.Resize(uint(*w), uint(*h), img, resize.NearestNeighbor)

	// 書き出すファイル名を指定します
	createFilePath := os.Args[2] + "." + data
	output, err := os.Create(createFilePath)
	if err != nil {
		log.Fatal(err)
	}
	// 最後にファイルを閉じる
	defer output.Close()

	// 画像のエンコード(書き込み)
	switch data {
	case "png":
		if err := png.Encode(output, resizedImg); err != nil {
			log.Fatal(err)
		}
	case "jpeg", "jpg":
		opts := &jpeg.Options{Quality: 100}
		if err := jpeg.Encode(output, resizedImg, opts); err != nil {
			log.Fatal(err)
		}
	default:
		if err := png.Encode(output, resizedImg); err != nil {
			log.Fatal(err)
		}
	}
}
