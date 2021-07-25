all: node
	@echo  "all"

node: 
	geth \
		-datadir test/geth/data \
		--dev \
		--ws  \
		--http \
		--mine \
		--metrics \
		--pprof \
		--graphql \
		--http.api admin,eth,debug,miner,net,txpool,personal,web3 \
		console 2 >> test/geth/log/geth.log

console: 
	geth attach test/geth/data/geth.ipc

clean:
	rm -rf test/geth/data
init:
	geth \
		-datadir test/geth/data \
		--dev \
		--ws  \
		--http \
		--mine \
		--http.api admin,eth,debug,miner,net,txpool,personal,web3 \
		init test/geth/genesis.json
