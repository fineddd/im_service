package main

import "net/http"
import "encoding/json"
import log "github.com/golang/glog"

type ServerSummary struct {
    nconnections int64
    nclients int64
    in_message_count int64
    out_message_count int64
}

func NewServerSummary() *ServerSummary {
    s := new(ServerSummary)
    return s
}

func Summary(rw http.ResponseWriter, req *http.Request) {
    obj := make(map[string]interface{})
    obj["connection_count"] = server_summary.nconnections
    obj["client_count"] = server_summary.nclients
    obj["in_message_count"] = server_summary.in_message_count
    obj["out_message_count"] = server_summary.out_message_count

    res, err := json.Marshal(obj)
    if err != nil {
        log.Info("json marshal:", err)
        return
    }

    rw.Header().Add("Content-Type", "application/json")
    _, err = rw.Write(res)
    if err != nil {
        log.Info("write err:", err)
    }
	return
}


func StartHttpServer(addr string) {
    go func () {
        http.HandleFunc("/summary", Summary)
        err := http.ListenAndServe(addr, nil)
        log.Info("http server err:", err)
    }()
}