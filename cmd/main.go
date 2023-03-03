package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-auth/pb"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-auth/pkg/repository"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-auth/pkg/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(&repository.Config{
		Host:     viper.GetString("HOST"),
		Port:     viper.GetString("DB_PORT"),
		Username: viper.GetString("USERNAME"),
		Password: viper.GetString("PASSWORD"),
		DBName:   viper.GetString("DB_NAME"),
		SSLMode:  viper.GetString("SSL_MODE"),
	})

	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	lis, err := net.Listen("tcp", viper.GetString("PORT"))
	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Auth Svc on", viper.GetString("PORT"))

	r := repository.NewRepository(db)
	s := service.NewService(r)

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("pkg/config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
