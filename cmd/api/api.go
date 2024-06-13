package api

import (
	"database/sql"

	_ "github.com/delapaska/1C/cmd/docs"
	"github.com/delapaska/1C/service/orders"
	"github.com/delapaska/1C/service/products"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type APIServer struct {
	engine *gin.Engine
	db     *sql.DB
}

func NewApiServer(db *sql.DB) *APIServer {

	engine := gin.New()
	engine.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	orderStore := orders.NewStore(db)
	orderHandler := orders.NewHandler(orderStore)
	orderHandler.RegisterRoutes(engine)
	productStore := products.NewStore(db)
	productHandler := products.NewHandler(productStore)
	productHandler.RegisterRoutes(engine)
	/*
		userStore := user.NewStore(db)
		userHandler := user.NewHandler(userStore)
		userHandler.RegisterRoutes(engine)

		projectStore := project.NewStore(db)
		projectHandler := project.NewHandler(projectStore, userStore)
		projectHandler.RegisterRoutes(engine)

		folderStore := folder.NewStore(db)
		folderHandler := folder.NewHandler(folderStore, userStore)
		folderHandler.RegisterRoutes(engine)
	*/
	return &APIServer{
		engine: engine,
		db:     db,
	}
}

func (s *APIServer) Run() {

	s.engine.Run(":8000")
}
