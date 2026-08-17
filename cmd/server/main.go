package main

import (
	"context"
	"errors"
	"github.com/wyw14/cry038/internal/application"
	"github.com/wyw14/cry038/internal/config"
	"github.com/wyw14/cry038/internal/domain"
	"github.com/wyw14/cry038/internal/repository"
	classhttp "github.com/wyw14/cry038/internal/transport/http"
	"go.uber.org/zap"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := config.Load()
	log, _ := zap.NewProduction()
	defer log.Sync()
	seed := domain.GenerateSession("demo", "版画入门", "Art-203", time.Now().Add(2*time.Hour), []domain.TemplateItem{{Key: "press", Name: "压印机安全检查", Kind: "safety", Critical: true, DefaultOwner: "助教"}})
	app := application.NewPreparation(repository.NewSessionMemory(seed), time.Now)
	srv := &http.Server{Addr: cfg.Listen, Handler: classhttp.Routes(app, log), ReadHeaderTimeout: 3 * time.Second}
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("server", zap.Error(err))
		}
	}()
	<-ctx.Done()
	shutdown, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()
	_ = srv.Shutdown(shutdown)
}
