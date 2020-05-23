package docconv

import (
	"io"
	"io/ioutil"
	"log"
	"sync"

	"github.com/otiai10/gosseract"
)

var langs = struct {
	sync.RWMutex
	lang string
}{lang: "eng"}

// ConvertImage converts images to text.
// Requires gosseract.
func ConvertImage(r io.Reader) (string, map[string]string, error) {

	b, err := ioutil.ReadAll(r)

	client := gosseract.NewClient()
	defer func() {
		_ = client.Close()
	}()

	err = client.SetLanguage("hye", "eng")
	if err != nil {
		log.Println("error setting language", err)
	}

	err = client.SetImageFromBytes(b)
	if err != nil {
		log.Println("error setting file", err)
	}

	text, _ := client.Text()

	return text, map[string]string{}, nil
}

// SetImageLanguages sets the languages parameter passed to gosseract.
func SetImageLanguages(l string) {
	langs.Lock()
	langs.lang = l
	langs.Unlock()
}
