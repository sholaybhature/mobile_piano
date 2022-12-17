package main

// downloaded notes from https://github.com/fuhton/piano-mp3 but renamed them to be in a serial numbering (0-12)*octave
// 1 is A note, this sounds bad but it is what it is
// copy paste the notes to /notes from og-notes before running the script
import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

func main() {
	files, err := ioutil.ReadDir("../notes")
	if err != nil {
		log.Fatal(err)
	}
	counter := 1
	note_key := 1
	last_val := 1
	var file_name string
	var src_file string
	for _, file := range files {
		if counter%8 == 0 {
			counter = 1
			note_key += 1
			last_val = note_key
		}
		// 1, 8, 15, 22, 29, 36, 42, 2, 9, 15
		file_name = "../notes/" + strconv.Itoa(last_val) + ".mp3"
		src_file = fmt.Sprintf("../notes/%s", file.Name())
		log.Print(src_file, file_name)
		// os.Rename(src_file, file_name)
		last_val = last_val + 7
		counter += 1
	}
}
