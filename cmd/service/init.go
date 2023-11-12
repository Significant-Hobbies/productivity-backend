package service

import (
	"net/http"
	"todo/pkg/metrics"
	validators "todo/pkg/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func CreateService() {
	e := echo.New()
	e.Use(middleware.CORS())

	e.Logger.SetLevel(1)
	logger, _ := zap.NewProduction()

	s := metrics.NewStats()
	e.Use(s.Process)
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(_ echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info("request",
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
			)
			return nil
		},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/api/test", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Hello, World!",
		})
	})

	e.POST("/api/todo", CreateTodo, validators.CreateTaskValidator)
	e.GET("/api/todo", GetTodo, validators.GetTasksValidator)
	e.DELETE("/api/todo", DeleteTodo, validators.DeleteTaskValidator)
	e.PATCH("/api/todo", UpdateTodo, validators.UpdateTaskValidator)

	e.POST("/api/admin/db_migrate", migrateDB)
	e.POST("/api/admin/db_delete_all", deleteAllTasks)
	e.GET("/api/admin/metrics", s.Handle)
	e.POST("/api/admin/db_seed", seedTasks)

	e.Logger.Fatal(e.Start(":1323"))
}
