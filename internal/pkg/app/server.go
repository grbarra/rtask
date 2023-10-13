package app

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"gitlab.com/grbarra/rtask/config"
	"gitlab.com/grbarra/rtask/internal/handlers"
	"gitlab.com/grbarra/rtask/internal/repository"
	"gitlab.com/grbarra/rtask/internal/service"
)

type APIServer struct {
	config  *config.Config
	handler handlers.Handler
	router  *http.ServeMux
	server  *http.Server
}

func New(config *config.Config) *APIServer {
	router := http.NewServeMux()
	server := &http.Server{
		Addr:    config.HTTPAddr,
		Handler: router,
	}

	return &APIServer{
		config: config,
		router: router,
		server: server,
	}
}

func (s *APIServer) Run() {
	srv := &http.Server{
		Addr:    s.config.HTTPAddr,
		Handler: s.router,
	}

	dbConf := s.InitConfig()
	db, err := repository.NewRedisDB(dbConf)
	if err != nil {
		logrus.Errorf("Failed to initialize db: %s", err)
	}
	defer db.Close()

	repo, err := repository.NewRepository(db)
	if err != nil {
		logrus.Errorf("Failed to initialize repos: %s", err)
		return
	}

	serv := service.NewService(repo)

	s.handler = *handlers.NewHandler(serv)

	s.confRouter()
	logrus.Infof("Запускаем сервер на порту %s", s.config.HTTPAddr)
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		<-sigint
		ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			logrus.Error(err)
		}
	}()
	if err := srv.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			logrus.Error(err)
		}
	}
	logrus.Info("Всего доброго!")
}

func (s *APIServer) InitConfig() repository.RedisConfig {
	if err := godotenv.Load(); err != nil {
		logrus.Info("Error initializing config db: %s", err)
	}

	dbConfig := repository.RedisConfig{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
	}

	return dbConfig
}
