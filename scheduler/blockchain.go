package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"nimbusx/bindings"
)

// BlockchainClient wraps the Ethereum client and smart contract bindings.
type BlockchainClient struct {
	client      *ethclient.Client
	marketplace *bindings.Marketplace
	privateKey  *ecdsa.PrivateKey
	chainID     *big.Int
}

// NewBlockchainClient initializes connection to the Ethereum node and contract bindings.
func NewBlockchainClient() (*BlockchainClient, error) {
	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		rpcURL = "http://127.0.0.1:8545" // Hardhat local network default
	}

	privKeyHex := os.Getenv("PRIVATE_KEY")
	if privKeyHex == "" {
		// Hardhat account[0] private key as fallback for local dev
		privKeyHex = "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	}
	// Strip "0x" prefix if present
	if len(privKeyHex) >= 2 && privKeyHex[:2] == "0x" {
		privKeyHex = privKeyHex[2:]
	}

	marketplaceAddrHex := os.Getenv("MARKETPLACE_ADDRESS")
	if marketplaceAddrHex == "" {
		return nil, fmt.Errorf("MARKETPLACE_ADDRESS environment variable is not set")
	}

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum node: %v", err)
	}

	privateKey, err := crypto.HexToECDSA(privKeyHex)
	if err != nil {
		return nil, fmt.Errorf("invalid private key: %v", err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID: %v", err)
	}

	marketplaceAddr := common.HexToAddress(marketplaceAddrHex)
	marketplace, err := bindings.NewMarketplace(marketplaceAddr, client)
	if err != nil {
		return nil, fmt.Errorf("failed to bind to Marketplace contract: %v", err)
	}

	return &BlockchainClient{
		client:      client,
		marketplace: marketplace,
		privateKey:  privateKey,
		chainID:     chainID,
	}, nil
}

// getAuth creates a transactor with the correct nonce and gas price.
func (b *BlockchainClient) getAuth(ctx context.Context) (*bind.TransactOpts, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(b.privateKey, b.chainID)
	if err != nil {
		return nil, err
	}

	// Fetch current gas price
	gasPrice, err := b.client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}
	auth.GasPrice = gasPrice

	return auth, nil
}

// AssignJobOnChain sends a transaction to assign a job to a provider on the Marketplace contract.
func (b *BlockchainClient) AssignJobOnChain(ctx context.Context, jobID string, providerAddrHex string) error {
	auth, err := b.getAuth(ctx)
	if err != nil {
		return fmt.Errorf("failed to create transactor: %v", err)
	}

	providerAddr := common.HexToAddress(providerAddrHex)
	
	// We need to pass the on-chain numeric Job ID, for now we map string UUID to big.Int 
	// Or we should assume the string ID is actually a numeric string representation.
	// Since our struct uses UUIDs for `id`, but the contract uses `uint256 jobId`,
	// we will hash the UUID to create a uint256 jobId, or change the Job ID logic.
	// For this basic integration, let's assume `jobID` can be parsed to a BigInt,
	// or we will use a hash. Let's just create a dummy ID for now or use the jobID string as big int if possible.
	numericJobID := new(big.Int)
	_, ok := numericJobID.SetString(jobID, 10)
	if !ok {
		// If it's a UUID, we'll hash it to a big Int to fit into uint256
		hash := crypto.Keccak256Hash([]byte(jobID))
		numericJobID.SetBytes(hash.Bytes())
	}

	tx, err := b.marketplace.AssignJob(auth, numericJobID, providerAddr)
	if err != nil {
		return fmt.Errorf("failed to send AssignJob tx: %v", err)
	}

	fmt.Printf("[blockchain] Assigned job %s on-chain. TX Hash: %s\n", jobID, tx.Hash().Hex())
	return nil
}
