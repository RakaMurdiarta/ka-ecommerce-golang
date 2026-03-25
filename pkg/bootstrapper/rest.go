package bootstrapper

import (
	"net/http"

	"github.com/RakaMurdiarta/ka-ecommerce-golang/internal/modules/example/provider"
	examplRepoImpl "github.com/RakaMurdiarta/ka-ecommerce-golang/internal/modules/example/repository/impl"
	examplServiceImpl "github.com/RakaMurdiarta/ka-ecommerce-golang/internal/modules/example/services/impl"
	"github.com/labstack/echo/v5"
)

type Server struct {
	e *echo.Echo
}

func NewServer(e *echo.Echo) *Server {
	return &Server{
		e: e,
	}
}

func (s *Server) InitApi() {
	privateRoute, publicRoute := s.initInternalRoute()

	exampleRepo := examplRepoImpl.NewExampleRepoImpl()
	exampleService := examplServiceImpl.NewExampleService(exampleRepo)

	provider.NewExampleProvider(privateRoute, publicRoute, exampleService)

}

func (s *Server) initInternalRoute() (privateRoute *echo.Group, publicRoute *echo.Group) {

	s.e.GET("/", ping)
	s.e.GET("/health", healthFunc)

	privateRoute = s.e.Group("")
	publicRoute = s.e.Group("")

	return privateRoute, publicRoute

}

func healthFunc(c *echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Server OK",
	})
}

func ping(c *echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
