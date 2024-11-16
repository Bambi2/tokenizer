CONTRACTS = code/*.sol
CONTRACT_BINS = code/bin/g42/*.abi code/bin/g42/*.bin code/bin/multisig/*.abi code/bin/multisig/*.bin
GENERATED_GO_CODE = deployment/g42/*.go deployment/multisig/*.go

$(CONTRACT_BINS): $(CONTRACTS)
	$(info $(M) building contracts...)
	@solc --abi code/G42.sol -o code/bin/g42 --overwrite
	@solc --abi code/MultiSig.sol -o code/bin/multisig --overwrite
	@solc --bin code/G42.sol -o code/bin/g42 --overwrite
	@solc --bin code/MultiSig.sol -o code/bin/multisig --overwrite

$(GENERATED_GO_CODE): $(CONTRACT_BINS)
	$(info $(M) tranfroming contracts to golang code...)
	mkdir -p deployment/g42
	mkdir -p deployment/multisig
	abigen --bin=code/bin/g42/G42.bin --abi=code/bin/g42/G42.abi --pkg=g42 --out=deployment/g42/g42.go
	abigen --bin=code/bin/multisig/MultiSigWallet.bin --abi=code/bin/multisig/MultiSigWallet.abi --pkg=multisig --out=deployment/multisig/multisig.go

generate: $(GENERATED_GO_CODE)

build:
	docker build -t tokenizer .
