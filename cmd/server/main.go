package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/RakaMurdiarta/ka-ecommerce-golang/pkg/bootstrapper"
	"github.com/labstack/echo/v5"
)

func main() {

	ctx := context.Background()

	echo := echo.New()

	apiServer := bootstrapper.NewServer(echo)

	apiServer.InitApi()

	listen := net.JoinHostPort("localhost", "8080")

	srv := &http.Server{
		Addr:    listen,
		Handler: echo,
	}

	fmt.Printf("[Server] Running on : http://localhost:8080")

	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(ctx, fmt.Sprintf("[Server] Error running server : %v", err))
	}

}
