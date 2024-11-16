# Tokenizer

In this project I implemented **ERC20** and **MultiSig** smart-contracts patterns
to create my own decentralized token with the feature of multiple signature confirmation.
In addition, I created **Go** app to deploy the contact to the blockchain with some extra features.

## Choices Made
I decided to write deployment application in Golang, thanks to ethereum official Go module
for ethereum protocol interactions which I found very handy.

As for the blockchain, I have chosen Holesky testnet as it was pretty easy to get free assets there for the deployment.
It's also recommended by Ethereum itself for the testing.

## How it works
I first created the contracts themselves. Then I used `solc` to compile the contracts
to `.bin` and `.abi` extensions. Using `abigen` tool, I generated Go code for deployment and interactions with the contracts.
Then, with `github.com/ethereum/go-ethereum` module you can easily deploy the contracts and send transactions to it.
All you need is to find a node for your chain.

## Usage

Build container

`make build`

Create a wallet

`docker run wallet ./g42 createWallet`

Deployment with example of possible parameters
`docker run wallet ./g42  deploy --norc=2 --owners="address1,address2" --pkey="private key"`
