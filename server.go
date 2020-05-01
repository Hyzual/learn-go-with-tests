package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	player := strings.TrimPrefix(request.URL.Path, "/players/")

	switch request.Method {
	case http.MethodGet:
		p.showScore(writer, player)
	case http.MethodPost:
		p.processWin(writer, player)
	}
}

func (p *PlayerServer) showScore(writer http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		writer.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(writer, score)
}

func (p *PlayerServer) processWin(writer http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	writer.WriteHeader(http.StatusAccepted)
}

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}
