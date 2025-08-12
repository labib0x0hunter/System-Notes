package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

var signatures = map[string]string{
	"425a68":   "bzip2",
	"ffd8ff":   "jpeg/jpg",
	"25504446": "pdf",
}

func main() {

	filename := "whatami1"

	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	firstNBytes := make([]byte, 16)

	n, err := file.Read(firstNBytes)
	if err != nil {
		return
	}

	magicBytes := hex.EncodeToString(firstNBytes[:n])

	for mg, fileType := range signatures {
		if strings.HasPrefix(magicBytes, mg) {
			fmt.Println(filename, ":", fileType)
			return
		}
	}

}
