# How does G42 work?

## Introduction to Blockchain
A **blockchain** is a decentralized, distributed ledger that records transactions in a secure and transparent manner. Each record, or "block," is linked to the previous one, creating a "chain" of records. This structure ensures that once data is recorded, it cannot be altered without altering every subsequent block, providing high levels of security and immutability.

## Types of Blockchains
Blockchains are typically classified based on their functionality:

1. **Transaction-based Blockchains**: The first blockchain, Bitcoin, was designed solely for secure peer-to-peer transfer of currency (Bitcoin) by maintaining a ledger of transactions.
2. **Smart Contract-based Blockchains**: Ethereum introduced the concept of smart contracts, which allows more complex, programmable transactions that can automate various tasks, making Ethereum and similar blockchains suitable for decentralized applications (DApps).

## Smart Contracts
A smart contract is basically a program withing the blockchain network. Technically it's just another address that has functions to be called.
Withing the functions you can implement different logic, like storing and transferring native currency, providing different checks etc.

## What is ERC20?
ERC20 is an interface standard that defines a set of functions that your smart contract needs to implement to be recognized as a token.

## Key Functions in ERC20
The ERC20 standard outlines several key functions:

* **totalSupply**: Returns the total supply of tokens.
* **balanceOf**: Checks the balance of a specific address.
* **transfer**: Transfers a specified amount of tokens from the sender’s address to another.
* **approve**: Approves another address to spend a certain amount of tokens on behalf of the owner.
* **transferFrom**: Allows approved addresses to transfer tokens on behalf of the owner.
* **allowance**: Checks the amount an address is approved to spend on behalf of another.

These functions provide the foundation for a secure and consistent way to interact with ERC20 tokens.

## The G42 Token
G42 is just a solidity class that implements the ERC20 interface which was deployed on Holesky network 