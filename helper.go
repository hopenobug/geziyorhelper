package geziyorhelper

import (
	"log"

	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
)

func SaveFileCallback(filename string) func(*geziyor.Geziyor, *client.Response) {
	return func(g *geziyor.Geziyor, r *client.Response) {
		err := saveFile(filename, r.Body)
		if !g.Opt.LogDisabled {
			if err == nil {
				log.Printf("save %s ok\n", filename)
			} else {
				log.Printf("save %s error: %s\n", filename, err)
			}
		}
	}
}

type SaveFileOption struct {
	NeedDecode      bool
	SkipExistedFile bool
}

var DefaultSaveFileOption = SaveFileOption{
	NeedDecode:      false,
	SkipExistedFile: true,
}

// SaveFile saves the result of request to a file with given name.
func SaveFile(g *geziyor.Geziyor, url, filename string, option ...*SaveFileOption) (err error) {
	opt := &DefaultSaveFileOption
	if len(option) > 0 && option[0] != nil {
		opt = option[0]
	}

	if opt.SkipExistedFile {
		var exists bool
		exists, err = pathExists(filename)
		if err != nil || exists {
			return
		}
	}

	req, _ := client.NewRequest("GET", url, nil)
	if !opt.NeedDecode {
		req.Encoding = "."
	}
	g.Do(req, SaveFileCallback(filename))
	return
}
