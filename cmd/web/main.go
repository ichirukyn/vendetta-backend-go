package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
	"vendetta/internal/adapters/store/sql"
	"vendetta/internal/app"
	"vendetta/internal/app/config"
	"vendetta/internal/services/postgres"
	"vendetta/internal/services/rabbitmq"
	"vendetta/pkg/utils"
)

var traceName = "[APP]"

// Точка входа в приложение
func main() {
	time.Local = time.UTC

	// Инициализация конфига
	c := config.NewConfig()

	// Инициализация логгера
	l := utils.NewDefaultLogger(c.AppMode)

	// Инициализация глобального контекста
	ctx := utils.NewContextWithCancel(context.Background())

	// Инициализация базы данных
	psql := postgres.NewPostgresService(c, l)
	if err := psql.Init(); err != nil {
		l.FatalT(traceName, "init postgres service error", err)
		return
	}

	// Отключение от бд в случае завершения работы приложения
	defer func(psql *postgres.Service) {
		if err := psql.Disconnect(); err != nil {
			l.FatalT(traceName, "disconnect postgres service error", err)
		}
	}(psql)

	// Инициализация раббита
	rmq := rabbitmq.NewRabbitMQService(l, c)
	if err := rmq.Init(); err != nil {
		l.ErrorT(traceName, "Error connecting to RabbitMQ", err)
		return
	}

	// Отключение от раббита в случае завершения работы приложения
	defer func(rmq *rabbitmq.Service) {
		err := rmq.Disconnect()
		if err != nil {
			l.ErrorT(traceName, "Error when disconnecting from RabbitMQ", err)
		}
	}(rmq)

	// Инициализация стора и сервера
	store := sql.New(psql.GetDatabase())
	srv := app.NewServer(ctx.GetContext(), l, c, store, rmq, psql)

	// Инициализация серверов для метрик и самого приложения
	ms := srv.GetMetricsServer(c.MetricsBindAddress)
	as := srv.GetAppServer(c.BindAddress)

	// Запуск серверов
	go func() {
		go func() {
			l.InfoT(traceName, "The Metrics Server has been started", c.MetricsBindAddress)
			if err := ms.ListenAndServe(); err != nil {
				l.ErrorT(traceName, fmt.Sprintf("listen metrics server: %s\n", err))
			}
		}()

		go func() {
			l.InfoT(traceName, "The Server has been started", c.BindAddress)
			if err := as.ListenAndServe(); err != nil {
				l.ErrorT(traceName, fmt.Sprintf("listen app server: %s\n", err))
			}
		}()
	}()

	// Канал для отслеживания сигналов системы о прекращении работы приложения
	quit := make(chan os.Signal, 1)

	// Канал для отслеживания, завершились ли все пост-процессы для правильного завершения работы
	done := make(chan bool)

	// Отслеживание сигнала
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Прослушивание канала сигналов
	<-quit

	// В случае завершения контекста (завершения работы приложения), выполняется коллбек
	ctx.OnCancel(func() {
		l.InfoT(traceName, "Shutdown Server ...")

		ctx := utils.NewContextWithTimeout(ctx.GetContext(), 5*time.Second)
		defer ctx.Cancel()

		{
			if err := ms.Shutdown(ctx.GetContext()); err != nil {
				l.FatalT(traceName, fmt.Sprintf("Metrics server Shutdown: %s\n", err))
			}

			l.InfoT(traceName, "Metrics server has been shutdown")

			if err := as.Shutdown(ctx.GetContext()); err != nil {
				l.FatalT(traceName, fmt.Sprintf("App server Shutdown: %s\n", err))
			}

			l.InfoT(traceName, "App server has been shutdown")
		}

		func(rmq *rabbitmq.Service) {
			err := rmq.Disconnect()
			if err != nil {
				l.ErrorT(traceName, "Error when disconnecting from RabbitMQ", err)
			}
		}(rmq)

		done <- true
	})

	// Завершаем контекст, чтобы выполнился коллбек onCancel
	ctx.Cancel()

	// Слушаем канал, чтобы знать, завершился коллбек или нет
	<-done

	l.InfoT(traceName, "Server exiting")
}
