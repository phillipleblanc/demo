package main

import (
	"database/sql"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/spiceai/gospice/v2"
	"github.com/spiceai/gospice/v2/pkg/function"
)

func ShareToUnderlying(ctx *function.FunctionCtx, duckDb *sql.DB, client *gospice.SpiceClient) error {
	// Temporary step
	_, err := duckDb.ExecContext(ctx, `
	CREATE TABLE IF NOT EXISTS output.strategy_manager_share_to_token (
		block_number BIGINT PRIMARY KEY,
		shares_to_underlying DOUBLE PRECISION,
		strategy_address TEXT
	);`)
	if err != nil {
		return err
	}

	apiKey := os.Getenv("SPICE_API_KEY")
	spiceEndpoint := fmt.Sprintf("https://dev-data.spiceai.io/goerli?api_key=%s", apiKey)
	rpcClient, err := rpc.DialContext(ctx, spiceEndpoint)
	if err != nil {
		return err
	}
	defer rpcClient.Close()
	ethClient := ethclient.NewClient(rpcClient)

	strategyAddress := "0xb613e78e2068d7489bb66419fb1cfa11275d14da"
	strategyCaller, err := NewStrategy(common.HexToAddress(strategyAddress), ethClient)
	if err != nil {
		return err
	}

	underlying, err := strategyCaller.SharesToUnderlying(nil, big.NewInt(100))
	if err != nil {
		return err
	}

	sharesToUnderlyingRatio := new(big.Float).SetInt(underlying)
	sharesToUnderlyingRatio = sharesToUnderlyingRatio.Quo(sharesToUnderlyingRatio, big.NewFloat(100))
	var sharesToUnderlying float64
	sharesToUnderlying, _ = sharesToUnderlyingRatio.Float64()

	_, err = duckDb.ExecContext(ctx, `
	INSERT INTO output.strategy_manager_share_to_token 
	(block_number, shares_to_underlying, strategy_address) 
	VALUES ($1, $2, $3);`,
		ctx.BlockNumber(), sharesToUnderlying, strategyAddress)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	function.RunFunction(ShareToUnderlying)
}
