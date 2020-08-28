package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/baking-bad/bcdhub/internal/config"
	"github.com/baking-bad/bcdhub/internal/helpers"
	"github.com/baking-bad/bcdhub/internal/logger"
	"github.com/baking-bad/bcdhub/internal/mq"
	"github.com/baking-bad/bcdhub/internal/verifier/compilation"
	"github.com/streadway/amqp"
)

// Context -
type Context struct {
	*config.Context
}

func main() {
	cfg, err := config.LoadDefaultConfig()
	if err != nil {
		logger.Fatal(err)
	}

	if cfg.Metrics.Sentry.Enabled {
		helpers.InitSentry(cfg.Sentry.Debug, cfg.Sentry.Environment, cfg.Sentry.URI)
		helpers.SetTagSentry("project", cfg.Metrics.Sentry.Project)
		defer helpers.CatchPanicSentry()
	}

	context := &Context{
		config.NewContext(
			config.WithRPC(cfg.RPC),
			config.WithDatabase(cfg.DB),
			config.WithRabbitReceiver(cfg.RabbitMQ, "verifier"),
		),
	}

	msgs, err := context.MQ.Consume(mq.QueueCompilations)
	if err != nil {
		logger.Fatal(err)
	}

	defer context.MQ.Close()

	logger.Info("Connected to %s queue", mq.QueueCompilations)

	for msg := range msgs {
		if err := context.handleMessage(msg); err != nil {
			logger.Error(err)
		}
	}

}

func (ctx *Context) handleMessage(data amqp.Delivery) error {
	defer func(d amqp.Delivery) {
		if err := data.Ack(false); err != nil {
			logger.Errorf("Error acknowledging message: %s", err)
		}
	}(data)

	return ctx.parseData(data)
}

func (ctx *Context) parseData(data amqp.Delivery) error {
	if data.RoutingKey != mq.QueueCompilations {
		return fmt.Errorf("[parseData] Unknown data routing key %s", data.RoutingKey)
	}

	var ct compilation.Task
	if err := json.Unmarshal(data.Body, &ct); err != nil {
		return fmt.Errorf("[parseData] Unmarshal message body error: %s", err)
	}

	defer os.RemoveAll(ct.Dir) // clean up

	switch ct.Kind {
	case compilation.KindVerification:
		return ctx.verification(ct)
	case compilation.KindCompilation:
		log.Fatal("not implemented", compilation.KindCompilation)
	case compilation.KindDeployment:
		log.Fatal("not implemented", compilation.KindDeployment)
	}

	return fmt.Errorf("[parseData] Unknown compilation task kind %s", ct.Kind)
}
