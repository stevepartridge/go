package locale

import (
	"github.com/stevepartridge/go/log"
)

var ref = map[string]string{}

func main() {

}

func Add(lang, key, text string) {
	k := lang + "." + key
	ref[k] = text
}

func Get(lang, key string) string {
	k := lang + "." + key
	if ref[k] != "" {
		return ref[k]
	} else {
		log.Notice("locale.Get", "Reference not found for "+key+" in "+lang)
		return k
	}
}
