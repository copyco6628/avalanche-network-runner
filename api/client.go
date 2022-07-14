package api

import (
	"github.com/ava-labs/coreth/plugin/evm"
	"github.com/copyco6628/avalanchego/api/admin"
	"github.com/copyco6628/avalanchego/api/health"
	"github.com/copyco6628/avalanchego/api/info"
	"github.com/copyco6628/avalanchego/api/ipcs"
	"github.com/copyco6628/avalanchego/api/keystore"
	"github.com/copyco6628/avalanchego/indexer"
	"github.com/copyco6628/avalanchego/vms/avm"
	"github.com/copyco6628/avalanchego/vms/platformvm"
)

// Issues API calls to a node
// TODO: byzantine api. check if appropiate. improve implementation.
type Client interface {
	PChainAPI() platformvm.Client
	XChainAPI() avm.Client
	XChainWalletAPI() avm.WalletClient
	CChainAPI() evm.Client
	CChainEthAPI() EthClient // ethclient websocket wrapper that adds mutexed calls, and lazy conn init (on first call)
	InfoAPI() info.Client
	HealthAPI() health.Client
	IpcsAPI() ipcs.Client
	KeystoreAPI() keystore.Client
	AdminAPI() admin.Client
	PChainIndexAPI() indexer.Client
	CChainIndexAPI() indexer.Client
	// TODO add methods
}
