package main

import (
	_ "github.com/lib/pq"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-auth/internal/app"
)

func main() {
	app.Run()
}
