#!/usr/bin/make -f

########################################
### Simulations

sim-regen-nondeterminism:
	@echo "Running nondeterminism test..."
	@go test -mod=readonly $(APP_DIR) -run TestAppStateDeterminism -Enabled=true \
		-NumBlocks=5 -BlockSize=10 -Commit=true -Period=0 -v -timeout 24h -tags=experimental

sim-regen-custom-genesis-fast:
	@echo "Running custom genesis simulation..."
	@echo "By default, ${HOME}/.regen/config/genesis.json will be used."
	@go test -mod=readonly $(APP_DIR) -run TestFullAppSimulation -Genesis=${HOME}/.regen/config/genesis.json \
		-Enabled=true -NumBlocks=100 -BlockSize=200 -Commit=true -Seed=99 -Period=5 -v -timeout 24h -tags=experimental

sim-regen-fast:
	@echo "Running quick Regen simulation. This may take several minutes..."
	@go test -mod=readonly $(APP_DIR) -run TestFullAppSimulation -Enabled=true -NumBlocks=50 -BlockSize=100 -Commit=true -Seed=99 -Period=5 -v -timeout 24h -tags=experimental

sim-regen-import-export: runsim
	@echo "Running Regen import/export simulation. This may take several minutes..."
	$(GOPATH)/bin/runsim -Jobs=4 -ExitOnFail 25 5 TestImportExport

sim-regen-after-import: runsim
	@echo "Running application simulation-after-import. This may take several minutes..."
	$(GOPATH)/bin/runsim -Jobs=4 -ExitOnFail 50 5 TestAppSimulationAfterImport

SIM_NUM_BLOCKS ?= 500
SIM_BLOCK_SIZE ?= 200
SIM_COMMIT ?= true

sim-regen-benchmark:
	@echo "Running application benchmark for numBlocks=$(SIM_NUM_BLOCKS), blockSize=$(SIM_BLOCK_SIZE). This may take awhile!"
	@go test -mod=readonly -benchmem -run=^$$ $(APP_DIR) -bench ^BenchmarkFullAppSimulation$$  \
		-Enabled=true -NumBlocks=$(SIM_NUM_BLOCKS) -BlockSize=$(SIM_BLOCK_SIZE) -Commit=$(SIM_COMMIT) -timeout 24h

sim-regen-profile:
	@echo "Running application benchmark for numBlocks=$(SIM_NUM_BLOCKS), blockSize=$(SIM_BLOCK_SIZE). This may take awhile!"
	@go test -mod=readonly -benchmem -run=^$$ $(APP_DIR) -bench ^BenchmarkFullAppSimulation$$ \
		-Enabled=true -NumBlocks=$(SIM_NUM_BLOCKS) -BlockSize=$(SIM_BLOCK_SIZE) -Commit=$(SIM_COMMIT) -timeout 24h -cpuprofile cpu.out -memprofile mem.out


.PHONY: runsim sim-regen-nondeterminism sim-regen-custom-genesis-fast sim-regen-fast sim-regen-import-export \
	sim-regen-after-import sim-regen-benchmark sim-regen-profile
