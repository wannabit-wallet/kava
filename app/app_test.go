package app

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"

	abci "github.com/tendermint/tendermint/abci/types"
)

func TestExport(t *testing.T) {
	db := db.NewMemDB()
	app := NewApp(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, 0)
	setGenesis(app)

	// Making a new app object with the db, so that initchain hasn't been called
	newApp := NewApp(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, 0)
	_, _, err := newApp.ExportAppStateAndValidators(false, []string{})
	require.NoError(t, err, "ExportAppStateAndValidators should not have an error")
}

func setGenesis(app *App) error {
	genesisState := NewDefaultGenesisState()

	stateBytes, err := codec.MarshalJSONIndent(app.cdc, genesisState)
	if err != nil {
		return err
	}

	// Initialize the chain
	app.InitChain(
		abci.RequestInitChain{
			Validators:    []abci.ValidatorUpdate{},
			AppStateBytes: stateBytes,
		},
	)
	app.Commit()

	return nil
}
