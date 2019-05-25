package main

import (
	"net/http"

	_ "net/http/pprof"

	"github.com/pion/sfu/conf"
	"github.com/pion/sfu/gslb"
	"github.com/pion/sfu/log"
	"github.com/pion/sfu/service"
)

func main() {
	if conf.SFU.Pprof != "" {
		go func() {
			log.Infof("Start pprof on %s", conf.SFU.Pprof)
			http.ListenAndServe(conf.SFU.Pprof, nil)
		}()
	}

	if !conf.SFU.Single {
		g, err := gslb.New()
		if err != nil {
			log.Errorf("gslb err => %v", err)
			return
		}
		g.KeepAlive()
	}
	service.Start()
	select {}
}
