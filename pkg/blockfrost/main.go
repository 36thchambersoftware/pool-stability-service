package blockfrost

import (
	"context"
	"encoding/hex"
	"log/slog"
	"os"
	"strings"

	bfg "github.com/blockfrost/blockfrost-go"
)

var (
	client         bfg.APIClient
	APIQueryParams bfg.APIQueryParams
	blockfrostProjectID string
)

const (
	LOVELACE             = 1_000_000
	ADA_HANDLE_PREFIX    = "$"
	ADA_HANDLE_POLICY_ID = "f0ff48bbb7bbe9d59a40f1ce90e9e9d0ff5002ec48f232b49ca0fb9a"
	CIP68v1_NONSENSE     = "000de140"
)

type (
	Lovelace int
	Ada      int
)

type AddressExtended struct {
	Address      string   `json:"address,omitempty"`
	Amount       []Amount `json:"amount,omitempty"`
	StakeAddress string   `json:"stake_address,omitempty"`
	Type         string   `json:"type,omitempty"`
	Script       bool     `json:"script,omitempty"`
}
type Amount struct {
	Unit                  string `json:"unit,omitempty"`
	Quantity              string `json:"quantity,omitempty"`
	Decimals              int    `json:"decimals,omitempty"`
	HasNftOnchainMetadata bool   `json:"has_nft_onchain_metadata,omitempty"`
}

func loadBlockfrostProjectID() string {
	blockfrostProjectID, ok := os.LookupEnv("BLOCKFROST_PROJECT_ID")
	if !ok {
		slog.Error("Could not get blockfrost project id")
	}

	return blockfrostProjectID
}

func init() {
	client = bfg.NewAPIClient(bfg.APIClientOptions{ProjectID: loadBlockfrostProjectID()})
}

// Convert ADA Handle address
func HandleAddress(ctx context.Context, addr string) (string, error) {
	isAdaHandle := strings.HasPrefix(addr, ADA_HANDLE_PREFIX)
	if isAdaHandle {
		hexAddr := hex.EncodeToString([]byte(addr[1:]))
		assetName := ADA_HANDLE_POLICY_ID + CIP68v1_NONSENSE + hexAddr
		addresses, err := client.AssetAddresses(ctx, assetName, APIQueryParams)
		if err != nil {
			return "", err
		}

		if len(addresses) > 0 {
			return addresses[0].Address, nil
		}

	}

	return addr, nil
}

func PoolInfo(ctx context.Context, poolID string) (*bfg.Pool, error) {
	info, err := client.Pool(ctx, poolID)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

func PoolHistory(ctx context.Context, poolID string) ([]bfg.PoolHistory, error) {
	APIQueryParams.Order = "desc"
	history, err := client.PoolHistory(ctx, poolID, APIQueryParams)
	if err != nil {
		return nil, err
	}

	return history, nil
}

func PoolMeta(ctx context.Context, poolID string) (*bfg.PoolMetadata, error) {
	info, err := client.PoolMetadata(ctx, poolID)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

func PoolBlocks(ctx context.Context, poolID string) (bfg.PoolBlocks, error) {
	APIQueryParams.Order = "desc"
	blocks, err := client.PoolBlocks(ctx, poolID, APIQueryParams)
	if err != nil {
		return nil, err
	}

	return blocks, nil
}

func PoolRelays(ctx context.Context, poolID string) ([]bfg.PoolRelay, error) {
	relays, err := client.PoolRelays(ctx, poolID)
	if err != nil {
		return []bfg.PoolRelay{}, err
	}

	return relays, nil
}