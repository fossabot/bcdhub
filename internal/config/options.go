package config

import (
	"time"

	"github.com/baking-bad/bcdhub/internal/bcd/tezerrors"
	"github.com/baking-bad/bcdhub/internal/postgres"
	"github.com/baking-bad/bcdhub/internal/postgres/account"
	"github.com/baking-bad/bcdhub/internal/postgres/bigmapdiff"
	"github.com/baking-bad/bcdhub/internal/postgres/contract"
	"github.com/baking-bad/bcdhub/internal/postgres/domains"
	"github.com/baking-bad/bcdhub/internal/postgres/global_constant"
	"github.com/baking-bad/bcdhub/internal/postgres/migration"
	"github.com/baking-bad/bcdhub/internal/postgres/operation"
	"github.com/baking-bad/bcdhub/internal/postgres/protocol"
	smartrollup "github.com/baking-bad/bcdhub/internal/postgres/smart_rollup"
	"github.com/baking-bad/bcdhub/internal/postgres/ticket"
	"github.com/baking-bad/bcdhub/internal/services/mempool"

	"github.com/baking-bad/bcdhub/internal/postgres/bigmapaction"
	"github.com/baking-bad/bcdhub/internal/postgres/block"
	pgCore "github.com/baking-bad/bcdhub/internal/postgres/core"

	"github.com/baking-bad/bcdhub/internal/noderpc"
)

// ContextOption -
type ContextOption func(ctx *Context)

// WithRPC -
func WithRPC(rpcConfig map[string]RPCConfig) ContextOption {
	return func(ctx *Context) {
		if rpcProvider, ok := rpcConfig[ctx.Network.String()]; ok {
			if rpcProvider.URI == "" {
				return
			}
			opts := []noderpc.NodeOption{
				noderpc.WithTimeout(time.Second * time.Duration(rpcProvider.Timeout)),
				noderpc.WithRateLimit(rpcProvider.RequestsPerSecond),
			}
			if rpcProvider.Log {
				opts = append(opts, noderpc.WithLog())
			}

			ctx.RPC = noderpc.NewNodeRPC(rpcProvider.URI, opts...)
		}
	}
}

// WithWaitRPC -
func WithWaitRPC(rpcConfig map[string]RPCConfig) ContextOption {
	return func(ctx *Context) {
		if rpcProvider, ok := rpcConfig[ctx.Network.String()]; ok {
			if rpcProvider.URI == "" {
				return
			}
			opts := []noderpc.NodeOption{
				noderpc.WithTimeout(time.Second * time.Duration(rpcProvider.Timeout)),
				noderpc.WithRateLimit(rpcProvider.RequestsPerSecond),
			}
			if rpcProvider.Log {
				opts = append(opts, noderpc.WithLog())
			}

			ctx.RPC = noderpc.NewWaitNodeRPC(rpcProvider.URI, opts...)
		}
	}
}

// WithStorage -
func WithStorage(cfg StorageConfig, appName string, maxPageSize int64, maxConnCount, idleConnCount int, createDatabaseIfNotExists bool) ContextOption {
	return func(ctx *Context) {
		if len(cfg.Postgres.Host) == 0 {
			panic("Please set connection strings to storage in config")
		}

		opts := []pgCore.PostgresOption{
			pgCore.WithPageSize(maxPageSize),
			pgCore.WithIdleConnections(idleConnCount),
			pgCore.WithMaxConnections(maxConnCount),
		}

		if cfg.LogQueries {
			opts = append(opts, pgCore.WithQueryLogging())
		}

		conn := pgCore.WaitNew(
			cfg.Postgres.ConnectionString(), ctx.Network.String(),
			appName, cfg.Timeout, opts...,
		)

		contractStorage := contract.NewStorage(conn)
		ctx.StorageDB = conn
		ctx.Storage = conn
		ctx.Accounts = account.NewStorage(conn)
		ctx.BigMapActions = bigmapaction.NewStorage(conn)
		ctx.Blocks = block.NewStorage(conn)
		ctx.BigMapDiffs = bigmapdiff.NewStorage(conn)
		ctx.Contracts = contractStorage
		ctx.Migrations = migration.NewStorage(conn)
		ctx.Operations = operation.NewStorage(conn)
		ctx.Protocols = protocol.NewStorage(conn)
		ctx.GlobalConstants = global_constant.NewStorage(conn)
		ctx.Domains = domains.NewStorage(conn)
		ctx.TicketUpdates = ticket.NewStorage(conn)
		ctx.Scripts = contractStorage
		ctx.SmartRollups = smartrollup.NewStorage(conn)
		ctx.Partitions = postgres.NewPartitionManager(conn)
	}
}

// WithConfigCopy -
func WithConfigCopy(cfg Config) ContextOption {
	return func(ctx *Context) {
		ctx.Config = cfg
	}
}

// WithMempool -
func WithMempool(cfg map[string]ServiceConfig) ContextOption {
	return func(ctx *Context) {
		if svcCfg, ok := cfg[ctx.Network.String()]; ok {
			if svcCfg.MempoolURI == "" {
				return
			}
			ctx.Mempool = mempool.NewMempool(svcCfg.MempoolURI)
		}
	}
}

// WithLoadErrorDescriptions -
func WithLoadErrorDescriptions() ContextOption {
	return func(ctx *Context) {
		if err := tezerrors.LoadErrorDescriptions(); err != nil {
			panic(err)
		}
	}
}
