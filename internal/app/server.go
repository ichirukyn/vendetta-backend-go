package app

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
	uuid "github.com/satori/go.uuid"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"time"
	"vendetta/internal/app/config"
	"vendetta/internal/app/initializers"
	"vendetta/internal/domain/store"
	"vendetta/internal/services/postgres"
	"vendetta/internal/services/rabbitmq"
	"vendetta/pkg/utils"
)

type Server struct {
	ctx context.Context
	l   *utils.Logger
	c   *config.Config

	ar *gin.Engine
	mr *gin.Engine

	s store.Store

	rmq  *rabbitmq.Service
	psql *postgres.Service
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.ar.ServeHTTP(w, r)
}

func (s *Server) GetMetricsServer(metricsAddr string) *http.Server {
	return &http.Server{
		Addr:    metricsAddr,
		Handler: s.mr,
	}
}

func (s *Server) GetAppServer(appAddr string) *http.Server {
	return &http.Server{
		Addr:    appAddr,
		Handler: s.ar,
	}
}

func (s *Server) configureRouter() {
	m := ginmetrics.GetMonitor()
	m.UseWithoutExposingEndpoint(s.ar)
	m.SetMetricPath("/metrics")
	m.Expose(s.mr)

	s.ar.Use(s.CORSMiddleware())
	s.ar.Use(s.setRequestID())
	s.ar.Use(s.logger())

	s.ar.Handle("GET", "/", HealthCheckHandler)

	s.ar.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	s.configureUserRouter()
}

func (s *Server) configureUserRouter() {
	user := initializers.NewUserInitializer(s.c, s.l, s.ar, s.s)
	user.Init()
}

func (s *Server) setRequestID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := uuid.NewV4().String()
		ctx.Writer.Header().Set("X-Request-ID", id)

		ctx.Next()
	}
}

func (s *Server) logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := time.Now()

		ctx.Next()

		latency := time.Since(t)

		id := ctx.Writer.Header().Get("X-Request-ID")
		url := ctx.Request.URL
		method := ctx.Request.Method
		body := ctx.Request.Body
		status := ctx.Writer.Status()
		s.l.Info(status, "—", method, url, body, id, "—", latency)
	}
}

type healthCheck struct {
	status string
}

func HealthCheckHandler(ctx *gin.Context) {
	response := healthCheck{
		status: "ok",
	}

	ctx.JSON(200, response)
}

func (s *Server) CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH,OPTIONS,GET,PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func NewServer(
	ctx context.Context,
	l *utils.Logger,
	c *config.Config,

	store store.Store,

	rmq *rabbitmq.Service,
	psql *postgres.Service,
) *Server {
	gin.SetMode(gin.ReleaseMode)

	s := &Server{
		ctx: ctx,
		l:   l,
		c:   c,

		ar: gin.New(),
		mr: gin.New(),

		s: store,

		rmq:  rmq,
		psql: psql,
	}

	s.configureRouter()

	return s
}
