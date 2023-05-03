package client

func GetDefaultConfPool() *ConfPool {
	conf := &ConfPool{
		AppID:   "web3.app_id.default",
		Zone:    "web3.zone.default",
		Cluster: "web3.cluster.default",
		Chains: map[int64]*ConfChain{
			1: {
				ChainID:         1,
				ChainName:       "Ethereum Mainnet",
				OfficialWebsite: "https://ethereum.org",
				ExplorerURL:     "https://etherscan.io",
				Faucets:         []string{},
				Clients: []*ConfClient{
					{Provider: "LlamaNodes", ProviderWebsite: "https://llamanodes.com", TransportSchema: "https", TransportURL: "https://eth.llamarpc.com"},
					{Provider: "OMNIA", ProviderWebsite: "https://omniatech.io", TransportSchema: "https", TransportURL: "https://endpoints.omniatech.io/v1/eth/mainnet/public"},
					{Provider: "Ankr", ProviderWebsite: "https://www.ankr.com", TransportSchema: "https", TransportURL: "https://rpc.ankr.com/eth"},
					{Provider: "PublicNode", ProviderWebsite: "https://ethereum.publicnode.com", TransportSchema: "https", TransportURL: "https://ethereum.publicnode.com"},
					{Provider: "1RPC", ProviderWebsite: "https://www.1rpc.io", TransportSchema: "https", TransportURL: "https://1rpc.io/eth"},
					{Provider: "MEV Blocker", ProviderWebsite: "https://mevblocker.io", TransportSchema: "https", TransportURL: "https://rpc.mevblocker.io"},
					{Provider: "FlashBots", ProviderWebsite: "https://www.flashbots.net", TransportSchema: "https", TransportURL: "https://rpc.flashbots.net"},
					{Provider: "CloudFlare", ProviderWebsite: "https://www.cloudflare.com/web3", TransportSchema: "https", TransportURL: "https://cloudflare-eth.com"},
					{Provider: "SecureRpc", ProviderWebsite: "https://securerpc.com", TransportSchema: "https", TransportURL: "https://api.securerpc.com/v1"},
					{Provider: "BlockPI", ProviderWebsite: "https://public.blockpi.io", TransportSchema: "https", TransportURL: "https://ethereum.blockpi.network/v1/rpc/public"},
					{Provider: "Payload.De", ProviderWebsite: "https://payload.de", TransportSchema: "https", TransportURL: "https://rpc.payload.de"},
					{Provider: "Alchemy", ProviderWebsite: "https://www.alchemy.com", TransportSchema: "https", TransportURL: "https://eth-mainnet.g.alchemy.com/v2/demo"},
					{Provider: "GasHawk", ProviderWebsite: "https://gashawk.io", TransportSchema: "https", TransportURL: "https://core.gashawk.io/rpc"},
				},
			},
			11155111: {
				ChainID:         11155111,
				ChainName:       "Ethereum Sepolia",
				OfficialWebsite: "https://ethereum.org",
				ExplorerURL:     "https://sepolia.etherscan.io",
				Faucets:         []string{},
				Clients: []*ConfClient{
					{Provider: "BlockPI", ProviderWebsite: "https://public.blockpi.io/", TransportSchema: "https", TransportURL: "https://ethereum-sepolia.blockpi.network/v1/rpc/public"},
				},
			},
			5: {
				ChainID:         5,
				ChainName:       "Ethereum Goerli",
				OfficialWebsite: "https://ethereum.org",
				ExplorerURL:     "https://goerli.etherscan.io",
				Faucets:         []string{},
				Clients: []*ConfClient{
					{Provider: "BlockPI", ProviderWebsite: "https://public.blockpi.io/", TransportSchema: "https", TransportURL: "https://goerli.blockpi.network/v1/rpc/public"},
				},
			},
			137: {
				ChainID:         137,
				ChainName:       "Polygon PoS Chain",
				OfficialWebsite: "https://polygon.technology",
				ExplorerURL:     "https://polygonscan.com",
				Faucets:         []string{},
				Clients: []*ConfClient{
					{Provider: "LlamaNodes", ProviderWebsite: "https://llamanodes.com", TransportSchema: "https", TransportURL: "https://polygon.llamarpc.com"},
					{Provider: "Ankr", ProviderWebsite: "https://polygon-rpc.com", TransportSchema: "https", TransportURL: "https://polygon-rpc.com"},
					{Provider: "QuickNode", ProviderWebsite: "https://www.quicknode.com", TransportSchema: "https", TransportURL: "https://rpc-mainnet.matic.quiknode.pro"},
				},
			},
			80001: {
				ChainID:         80001,
				ChainName:       "Polygon PoS Chain Testnet",
				OfficialWebsite: "https://polygon.technology",
				ExplorerURL:     "https://mumbai.polygonscan.com",
				Faucets:         []string{},
				Clients: []*ConfClient{
					{Provider: "BlockPI", ProviderWebsite: "https://public.blockpi.io/", TransportSchema: "https", TransportURL: "https://polygon-mumbai.blockpi.network/v1/rpc/public"},
				},
			},
			10: {
				ChainID:         10,
				ChainName:       "Optimism",
				OfficialWebsite: "https://www.optimism.io",
				ExplorerURL:     "https://optimistic.etherscan.io",
				Faucets:         []string{},
				Clients: []*ConfClient{
					{Provider: "BlockPI", ProviderWebsite: "https://public.blockpi.io/", TransportSchema: "https", TransportURL: "https://optimism.blockpi.network/v1/rpc/public"},
					{Provider: "Alchemy", ProviderWebsite: "https://www.alchemy.com/", TransportSchema: "https", TransportURL: "https://mainnet.optimism.io"},
				},
			},
			420: {
				ChainID:         420,
				ChainName:       "Optimism Goerli",
				OfficialWebsite: "https://www.optimism.io",
				ExplorerURL:     "https://goerli-explorer.optimism.io",
				Faucets:         []string{},
				Clients: []*ConfClient{
					{Provider: "BlockPI", ProviderWebsite: "https://public.blockpi.io/", TransportSchema: "https", TransportURL: "https://goerli.optimism.io"},
				},
			},
			42161: {
				ChainID:         42161,
				ChainName:       "Arbitrum One",
				OfficialWebsite: "https://arbitrum.io",
				ExplorerURL:     "https://arbiscan.io",
				Faucets:         []string{},
				Clients: []*ConfClient{
					{Provider: "BlockPI", ProviderWebsite: "https://public.blockpi.io/", TransportSchema: "https", TransportURL: "https://arbitrum.blockpi.network/v1/rpc/public"},
					{Provider: "Ankr", ProviderWebsite: "https://www.ankr.com", TransportSchema: "https", TransportURL: "https://rpc.ankr.com/arbitrum"},
					{Provider: "Arbitrum", ProviderWebsite: "https://arbitrum.io", TransportSchema: "https", TransportURL: "https://arb1.arbitrum.io/rpc"},
				},
			},
			42170: {
				ChainID:         42170,
				ChainName:       "Arbitrum Nova",
				OfficialWebsite: "https://arbitrum.io",
				ExplorerURL:     "https://nova.arbiscan.io",
				Faucets:         []string{},
				Clients: []*ConfClient{
					{Provider: "Arbitrum", ProviderWebsite: "https://developer.arbitrum.io/public-chains", TransportSchema: "https", TransportURL: "https://nova.arbitrum.io/rpc"},
				},
			},
			421613: {
				ChainID:         421613,
				ChainName:       "Arbitrum Goerli",
				OfficialWebsite: "https://arbitrum.io",
				ExplorerURL:     "https://goerli.arbiscan.io",
				Faucets:         []string{},
				Clients: []*ConfClient{
					{Provider: "Arbitrum", ProviderWebsite: "https://developer.arbitrum.io/public-chains", TransportSchema: "https", TransportURL: "https://goerli-rollup.arbitrum.io/rpc"},
				},
			},
			43114: {
				ChainID:         43114,
				ChainName:       "Avalanche Fuji",
				OfficialWebsite: "https://www.avax.network",
				ExplorerURL:     "https://testnet.snowtrace.io",
				Faucets:         []string{},
				Clients: []*ConfClient{
					{Provider: "Arbitrum", ProviderWebsite: "https://developer.arbitrum.io/public-chains", TransportSchema: "https", TransportURL: "https://rpc.ankr.com/avalanche"},
				},
			},
			43113: {
				ChainID:         43113,
				ChainName:       "Avalanche Fuji",
				OfficialWebsite: "https://www.avax.network",
				ExplorerURL:     "https://testnet.snowtrace.io",
				Faucets:         []string{},
				Clients: []*ConfClient{
					{Provider: "Ankr", ProviderWebsite: "https://developer.arbitrum.io/public-chains", TransportSchema: "https", TransportURL: "https://rpc.ankr.com/avalanche_fuji"},
				},
			},
			100: {
				ChainID:         100,
				ChainName:       "Gnosis",
				OfficialWebsite: "https://www.gnosis.io",
				ExplorerURL:     "https://testnet.snowtrace.io",
				Faucets:         []string{},
				Clients: []*ConfClient{
					{Provider: "BlockPI", ProviderWebsite: "https://blockscout.com/xdai/mainnet", TransportSchema: "https", TransportURL: "https://gnosis.blockpi.network/v1/rpc/public"},
				},
			},
			56: {
				ChainID:         56,
				ChainName:       "BNB Smart Chain",
				OfficialWebsite: "https://bscscan.com",
				ExplorerURL:     "https://bscscan.com",
				Faucets:         []string{},
				Clients: []*ConfClient{
					{Provider: "BlockPI", ProviderWebsite: "https://public.blockpi.io/", TransportSchema: "https", TransportURL: "https://bsc.blockpi.network/v1/rpc/public"},
				},
			},
			97: {
				ChainID:         97,
				ChainName:       "BSC Testnet",
				OfficialWebsite: "https://bscscan.com",
				ExplorerURL:     "https://testnet.bscscan.com",
				Faucets:         []string{},
				Clients: []*ConfClient{
					{Provider: "Blast", ProviderWebsite: "https://blastapi.io", TransportSchema: "https", TransportURL: "https://bsc-testnet.public.blastapi.io"},
				},
			},
		},
	}
	return conf
}
