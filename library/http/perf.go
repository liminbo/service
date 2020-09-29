package http

import (
	"github.com/micro/go-micro/v2/web"
	"net/http/pprof"
)

func StartPerf(service web.Service){
	service.HandleFunc("/debug/pprof/", pprof.Index)
	service.HandleFunc("/debug/pprof/trace/", pprof.Trace)
	service.HandleFunc("/debug/pprof/profile/", pprof.Profile)
	service.HandleFunc("/debug/pprof/cmdline/", pprof.Cmdline)
}