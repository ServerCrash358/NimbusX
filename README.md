# NimbusX

A decentralized cloud cost optimization protocol that schedules container workloads across compute providers using blockchain escrow, priority-based scheduling, and real-time pricing oracles.

---

## Architecture

```
User
  │
  ▼
API Gateway  (Go · port 8080)
  │
  ▼
Scheduler    (Go · port 9090)
  │
  ├──► Price Oracle  (Go · port 7070)
  │
  ├──► Smart Contracts (Solidity / Hardhat)
  │      ├── ProviderRegistry
  │      ├── Marketplace
  │      └── Escrow
  │
  └──► Node Agents  (Linux dev · Go)
           │
           ▼
       Containers (Docker / K8s)
```

---

## Repository Structure

```
NimbusX/
├── contracts/              # Solidity smart contracts
│   ├── ProviderRegistry.sol
│   ├── Marketplace.sol
│   └── Escrow.sol
├── scheduler/              # Go scheduler engine
│   ├── main.go
│   ├── scheduler.go
│   ├── scoring.go
│   ├── provider_manager.go
│   └── types.go
├── oracle/                 # Go price oracle service
│   ├── main.go
│   ├── oracle.go
│   └── pricing.go
├── api/                    # Go API gateway
│   ├── main.go
│   ├── handlers.go
│   └── middleware.go
├── ignition/modules/       # Hardhat Ignition deploy modules
│   └── NimbusX.ts
├── scripts/
│   └── deploy.ts           # Manual deploy script
├── test/
│   ├── NimbusX.ts          # Contract integration tests
│   └── integration_test.go # Go service integration tests
├── docker-compose.yml      # Local multi-service stack
├── go.mod
└── hardhat.config.ts
```

---

## Smart Contracts

### ProviderRegistry
Tracks compute providers on-chain. Providers register their CPU/memory/price and earn a reputation score (0–100) that affects scheduling priority.

### Marketplace
The job board. Clients post jobs with escrowed ETH; the off-chain scheduler assigns matching providers. Job lifecycle: `Open → Assigned → Running → Completed | Failed | Cancelled`.

### Escrow
Holds per-job ETH. Releases to provider on completion, refunds client on failure/cancellation. Only callable by the Marketplace contract.

---

## Getting Started

### Prerequisites
- Node.js ≥ 20
- Go ≥ 1.21
- `npx` / `npm`

### Install dependencies
```bash
npm install
```

### Run contract tests
```bash
npx hardhat test mocha
```

### Deploy to local simulated network
```bash
npx hardhat ignition deploy ignition/modules/NimbusX.ts
```

### Deploy to Sepolia testnet
```bash
npx hardhat keystore set SEPOLIA_PRIVATE_KEY
npx hardhat keystore set SEPOLIA_RPC_URL
npx hardhat ignition deploy ignition/modules/NimbusX.ts --network sepolia
```

### Run Go services locally
```bash
# API Gateway
go run ./api

# Scheduler
go run ./scheduler

# Price Oracle
go run ./oracle
```

### Run full local stack
```bash
docker-compose up
```

---

## Development Progress

| Step | Component | Status |
|------|-----------|--------|
| 1 | Smart Contracts | ✅ Done |
| 2 | Scheduler base | ⬜ |
| 3 | Oracle service | ⬜ |
| 4 | API Gateway | ⬜ |
| 5 | Scoring algorithm | ⬜ |
| 6 | Integration testing | ⬜ |

---

## License
MIT
