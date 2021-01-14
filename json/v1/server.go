package main

import (
	"fmt"
	"net/http"
	"strings"
)

//玩家信息接口
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

//给HTTP使用的接口 用于调用玩家信息接口查询玩家信息
type PlayerServer struct {
	store PlayerStore
	http.Handler
	//嵌入了http.Handler，便可以访问http.Handler的所有公共接口和字段
	//可以用嵌入将接口组装成新的接口，也可以嵌入类型，混合类型 目前个人理解有点类似js的混入吧
	//嵌入类型时，真正要考虑的是对你公开的 API 有什么影响。
	//不要嵌入过多，会污染你的接口
}

func NewPlayerServer(store PlayerStore) *PlayerServer {

	p := new(PlayerServer)

	p.store = store

	//Go 有一个内置的路由机制叫做 ServeMux（request multiplexer，多路请求复用器）
	//它允许你将 http.Handler 附加到特定的请求路径。
	router := http.NewServeMux()

	router.Handle("/players/", http.HandlerFunc(p.playerHandler))
	router.Handle("/league", http.HandlerFunc(p.leagueHandle))

	p.Handler = router

	return p
}

func (p *PlayerServer) playerHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodGet:
		p.showScore(w, player)
	case http.MethodPost:
		p.processWin(w, player)
	}
}

func (p *PlayerServer) leagueHandle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) showScore(w http.ResponseWriter, name string) {
	score := p.store.GetPlayerScore(name)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, name string) {
	p.store.RecordWin(name)
	w.WriteHeader(http.StatusAccepted)
}
