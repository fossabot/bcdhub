package operations

import (
	"github.com/baking-bad/bcdhub/internal/models/account"
	"github.com/baking-bad/bcdhub/internal/models/operation"
	"github.com/baking-bad/bcdhub/internal/models/types"
	"github.com/baking-bad/bcdhub/internal/noderpc"
	"github.com/baking-bad/bcdhub/internal/parsers"
)

// TxRollupOrigination -
type TxRollupOrigination struct {
	*ParseParams
}

// NewTxRollupOrigination -
func NewTxRollupOrigination(params *ParseParams) TxRollupOrigination {
	return TxRollupOrigination{params}
}

// Parse -
func (p TxRollupOrigination) Parse(data noderpc.Operation, store parsers.Store) error {
	source := account.Account{
		Address: data.Source,
		Type:    types.NewAccountType(data.Source),
	}

	txRollupOrigination := operation.Operation{
		Hash:         p.hash,
		ProtocolID:   p.protocol.ID,
		Level:        p.head.Level,
		Timestamp:    p.head.Timestamp,
		Kind:         types.NewOperationKind(data.Kind),
		Initiator:    source,
		Source:       source,
		Fee:          data.Fee,
		Counter:      data.Counter,
		GasLimit:     data.GasLimit,
		StorageLimit: data.StorageLimit,
		Nonce:        data.Nonce,
		ContentIndex: p.contentIdx,
	}

	p.fillInternal(&txRollupOrigination)

	parseOperationResult(data, &txRollupOrigination)

	txRollupOrigination.SetBurned(*p.protocol.Constants)

	p.stackTrace.Add(txRollupOrigination)

	store.AddOperations(&txRollupOrigination)

	return nil
}

func (p TxRollupOrigination) fillInternal(txRollupOrigination *operation.Operation) {
	if p.main == nil {
		return
	}

	txRollupOrigination.Counter = p.main.Counter
	txRollupOrigination.Hash = p.main.Hash
	txRollupOrigination.Level = p.main.Level
	txRollupOrigination.Timestamp = p.main.Timestamp
	txRollupOrigination.Internal = true
	txRollupOrigination.Initiator = p.main.Source
}
