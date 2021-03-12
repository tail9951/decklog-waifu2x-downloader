package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func (tsdb *Tsdb) outputTSDB(outputFolder string) {

	s := fmt.Sprintf("# Save File\ncardsx=%s\ncardsy=%s\ncard-width=%s\ncard-height=%s\nzoom=%s\nbackground-color=%s\n", tsdb.cardsX, tsdb.cardsY, tsdb.cardWidth, tsdb.cardHeight, tsdb.zoom, tsdb.backgroundColor)
	outputpath := fmt.Sprintf("./%s/newdeck.tsdb", outputFolder)

	err := ioutil.WriteFile(outputpath, []byte(s+tsdb.createTsdbLine(outputFolder)), 0700)
	checkErr(err)

}

func getImgFolderPath(outputFolder string) string {
	path, err := os.Getwd()
	checkErr(err)
	path = strings.Replace(path, "\\", "\\\\", -1)
	path = strings.Replace(path, ":", "\\:", -1)
	path = fmt.Sprintf("%s\\\\%s\\\\", path, outputFolder)
	return path
}

func (tsdb *Tsdb) createTsdbLine(outputFolder string) string {
	var buffer bytes.Buffer
	path := getImgFolderPath(outputFolder)
	i := 0
	j := 0

	for _, card := range tsdb.cards {
		fmt.Println("Downloading", fileNameCounter, "image ....")

		if len(apiKey) == 0 {
			err, _ := downloadFile(card.url, outputFolder)
			checkErr(err)
		} else {
			err, waifu2xURL := getWaifu2xImgFromURL(card.url)
			checkErr(err)
			err, _ = downloadFile(waifu2xURL, outputFolder)
			checkErr(err)
		}

		checkAndRotate(outputFolder + "/" + strconv.Itoa(fileNameCounter) + ".png")

		for k := 0; k < card.num; k++ {
			temp := fmt.Sprintf("%d_%d=%s%d.png\n", i, j, path, fileNameCounter)
			buffer.WriteString(temp)
			i++
			if i > 9 {
				j++
				i = 0
			}
		}
		fileNameCounter++
	}

	return buffer.String()
}
