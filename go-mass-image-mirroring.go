package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/disintegration/imaging"
)

var ext = []string{".jpg", ".png", ".jpeg"}
var sep = string(os.PathSeparator)

func main() {

	var sourcePath string
	var savePath string

	flag.StringVar(&sourcePath, "path", "", "-path /path/to/folder/that/contain/images")
	flag.StringVar(&savePath, "save", "", "-save /path/to/save/mirror/result")

	flag.Parse()

	// check is source folder exist
	if _, err := os.Stat(sourcePath); os.IsNotExist(err) {
		fmt.Println("source folder not found !, just like your future...")
		panic(err)
	}

	if _, err := os.Stat(savePath); os.IsNotExist(err) {
		err = os.Mkdir(savePath, os.ModeDir)
		if err != nil {
			panic(err)
		}
	}

	files, err := ioutil.ReadDir(sourcePath)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if !f.IsDir() && stringInSlice(filepath.Ext(f.Name()), ext) {
			fullPath := sourcePath + sep + f.Name()
			fmt.Println("processing ", fullPath)

			img, err := imaging.Open(fullPath)
			if err != nil {
				log.Println(err)
				continue
			}

			saveTo := savePath + sep + f.Name()
			// img, degrees, set a color to the background
			img = imaging.FlipH(img)
			err = imaging.Save(img, saveTo)
			if err != nil {
				log.Println(err)
			}
		}
	}

}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
