package operation

import (
	"encoding/hex"
	"time"

	"github.com/baking-bad/bcdhub/internal/bcd"
	"github.com/baking-bad/bcdhub/internal/bcd/ast"
	"github.com/baking-bad/bcdhub/internal/bcd/tezerrors"
	"github.com/baking-bad/bcdhub/internal/models/account"
	"github.com/baking-bad/bcdhub/internal/models/bigmapaction"
	"github.com/baking-bad/bcdhub/internal/models/bigmapdiff"
	"github.com/baking-bad/bcdhub/internal/models/protocol"
	"github.com/baking-bad/bcdhub/internal/models/ticket"
	"github.com/baking-bad/bcdhub/internal/models/types"
	"github.com/go-pg/pg/v10"
)

// Operation -
type Operation struct {
	// nolint
	tableName struct{} `pg:"operations,partition_by:RANGE(timestamp)"`

	ID                                 int64      `pg:",pk"`
	ContentIndex                       int64      `pg:",use_zero"`
	Level                              int64      `pg:",use_zero"`
	Counter                            int64      `pg:",use_zero"`
	Fee                                int64      `pg:",use_zero"`
	GasLimit                           int64      `pg:",use_zero"`
	StorageLimit                       int64      `pg:",use_zero"`
	Amount                             int64      `pg:",use_zero"`
	ConsumedGas                        int64      `pg:",use_zero"`
	StorageSize                        int64      `pg:",use_zero"`
	PaidStorageSizeDiff                int64      `pg:",use_zero"`
	Burned                             int64      `pg:",use_zero"`
	AllocatedDestinationContractBurned int64      `pg:",use_zero"`
	ProtocolID                         int64      `pg:",type:SMALLINT"`
	TicketUpdatesCount                 int        `pg:",use_zero"`
	BigMapDiffsCount                   int        `pg:",use_zero"`
	Tags                               types.Tags `pg:",use_zero"`
	Nonce                              *int64

	InitiatorID   int64
	Initiator     account.Account `pg:",rel:has-one"`
	SourceID      int64
	Source        account.Account `pg:",rel:has-one"`
	DestinationID int64
	Destination   account.Account `pg:",rel:has-one"`
	DelegateID    int64
	Delegate      account.Account `pg:",rel:has-one"`

	Timestamp time.Time             `pg:",pk,notnull"`
	Status    types.OperationStatus `pg:",type:SMALLINT"`
	Kind      types.OperationKind   `pg:",type:SMALLINT"`

	Entrypoint      types.NullString `pg:",type:text"`
	Tag             types.NullString `pg:",type:text"`
	Hash            []byte
	Parameters      []byte
	DeffatedStorage []byte
	Payload         []byte
	PayloadType     []byte
	Script          []byte `pg:"-"`

	Errors tezerrors.Errors `pg:",type:bytea"`

	AST *ast.Script `pg:"-"`

	BigMapDiffs   []*bigmapdiff.BigMapDiff     `pg:"rel:has-many"`
	BigMapActions []*bigmapaction.BigMapAction `pg:"rel:has-many"`
	TickerUpdates []*ticket.TicketUpdate       `pg:"rel:has-many"`

	AllocatedDestinationContract bool `pg:",use_zero"`
	Internal                     bool `pg:",use_zero"`
}

// GetID -
func (o *Operation) GetID() int64 {
	return o.ID
}

// GetIndex -
func (o *Operation) GetIndex() string {
	return "operations"
}

// Save -
func (o *Operation) Save(tx pg.DBI) error {
	_, err := tx.Model(o).Returning("id").Insert()
	return err
}

// LogFields -
func (o *Operation) LogFields() map[string]interface{} {
	return map[string]interface{}{
		"hash":  hex.EncodeToString(o.Hash),
		"block": o.Level,
	}
}

// SetAllocationBurn -
func (o *Operation) SetAllocationBurn(constants protocol.Constants) {
	o.AllocatedDestinationContractBurned = 257 * constants.CostPerByte
}

// SetBurned -
func (o *Operation) SetBurned(constants protocol.Constants) {
	if o.Status != types.OperationStatusApplied {
		return
	}
	var burned int64

	if o.PaidStorageSizeDiff != 0 {
		burned += o.PaidStorageSizeDiff * constants.CostPerByte
	}

	if o.AllocatedDestinationContract {
		o.SetAllocationBurn(constants)
		burned += o.AllocatedDestinationContractBurned
	}

	o.Burned = burned
}

// IsEntrypoint -
func (o *Operation) IsEntrypoint(entrypoint string) bool {
	return o.Entrypoint.EqualString(entrypoint)
}

// IsOrigination -
func (o *Operation) IsOrigination() bool {
	return o.Kind == types.OperationKindOrigination || o.Kind == types.OperationKindOriginationNew
}

// IsTransaction -
func (o *Operation) IsTransaction() bool {
	return o.Kind == types.OperationKindTransaction
}

// IsImplicit  -
func (o *Operation) IsImplicit() bool {
	return len(o.Hash) == 0
}

// IsApplied -
func (o *Operation) IsApplied() bool {
	return o.Status == types.OperationStatusApplied
}

// IsCall -
func (o *Operation) IsCall() bool {
	return (bcd.IsContract(o.Destination.Address) || bcd.IsSmartRollupHash(o.Destination.Address)) && len(o.Parameters) > 0
}

// Result -
type Result struct {
	Status                       string
	ConsumedGas                  int64
	StorageSize                  int64
	PaidStorageSizeDiff          int64
	AllocatedDestinationContract bool
	Originated                   string
	Errors                       []*tezerrors.Error
}

// Stats -
type Stats struct {
	Count      int64
	LastAction time.Time
}

// Pageable -
type Pageable struct {
	Operations []Operation
	LastID     string
}

// TokenMethodUsageStats -
type TokenMethodUsageStats struct {
	Count       int64
	ConsumedGas int64
}

// TokenUsageStats -
type TokenUsageStats map[string]TokenMethodUsageStats
