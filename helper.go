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
