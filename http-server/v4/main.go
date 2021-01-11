package main

import (
	"log"
	"net/http"
)

/**
TODO
很好！你创建了一个 RESTful 风格的服务。为了实现这一目标，你需要选择一个数据存储来持久化得分数据。
选择一种存储机制（Bolt? Mongo? Postgres? File system?）
用一个 PostgresPlayerStore 函数来实现 PlayerStore
通过 TDD 来确保它能正常工作
接入集成测试中，检查它是否依然正常工作
最终接入到主程序中
*/
func main() {
	server := &PlayerServer{NewInMemoryPlayerStore()}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
