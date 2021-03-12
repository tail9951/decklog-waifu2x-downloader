package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"gopkg.in/ini.v1"
)

var (
	fileNameCounter = 1
	configFileName  = "config.ini"
	apiURL          = "https://api.deepai.org/api/waifu2x"
	method          = "POST"
	outputFolder    string
	apiKey          string
	tsdb            Tsdb
)

type Tsdb struct {
	cardsX          string
	cardsY          string
	cardWidth       string
	cardHeight      string
	zoom            string
	backgroundColor string
	cards           []Card
}

type Card struct {
	num int
	url string
}

func createFolderIfNotExist(imageFolder string) {
	if _, err := os.Stat(imageFolder); os.IsNotExist(err) {
		os.Mkdir(imageFolder, 0700)
	}
}

func downloadFile(imgURL string, outputFolder string) (err error, imgPath string) {

	//Get the response bytes from the url
	response, err := http.Get(imgURL)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	//Create a empty file
	imgPath = fmt.Sprintf("%s/%s", outputFolder, strconv.Itoa(fileNameCounter)+".png")
	file, err := os.Create(imgPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//Write the bytes to the file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		panic(err)
	}
	return nil, imgPath
}

func getFileNameList(folder string) []string {
	var fileNameList []string
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fileNameList = append(fileNameList, folder+file.Name())
	}

	return fileNameList
}

func downloadDeck(cards *[]Card, filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		if len(scanner.Text()) > 0 {
			var card Card

			temp := strings.Split(scanner.Text(), " ")
			card.url = temp[0]
			card.num, err = strconv.Atoi(temp[1])
			// fmt.Println(card)
			*cards = append(*cards, card)
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func initConfig(configFileName string) {
	fmt.Println("Loading config...")
	cfg, err := ini.Load(configFileName)
	checkErr(err)
	tsdb.cardsX = cfg.Section("tsdb").Key("cardsX").String()
	tsdb.cardsY = cfg.Section("tsdb").Key("cardsY").String()
	tsdb.cardWidth = cfg.Section("tsdb").Key("cardWidth").String()
	tsdb.cardHeight = cfg.Section("tsdb").Key("cardHeight").String()
	tsdb.zoom = cfg.Section("tsdb").Key("zoom").String()
	tsdb.backgroundColor = cfg.Section("tsdb").Key("backgroundColor").String()

	outputFolder = cfg.Section("output").Key("folder-name").String()
	apiKey = cfg.Section("waifu2x").Key("api-key").String()
}

func main() {

	initConfig(configFileName)

	createFolderIfNotExist(outputFolder)

	downloadDeck(&tsdb.cards, os.Args[1])
	tsdb.outputTSDB(outputFolder)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
