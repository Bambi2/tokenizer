# Multi Signature

## Description
MultiSig is just a solidity pattern that allows users to make multiple confirmations on any transaction.
You can add multisig pattern to any function on your contract.
It will require multiple parties (chosen beforehand) to sign the call of the function before it is actually executed
This way MultiSig creates an additional level of security and creates a space for other creative logic of the contract

## Implementation
MultiSig is usually implemented using a separate contract.
When deployed the address of the contact will be saved and chosen as an owner of the main contract,
so the chosen function of the contact can be called only by MultiSig contact.
After that you will have no choice but to call your function only from multisig contract, but how to do so?
For that you will implement functions in the multisig contact that allow you to call other contracts from it.
In Solidity it's done by using `call` method on `address` data type

## Address
Main address of my MultiSig contract on Holesky chain is `0x53A365d956fEA96bF11A916b2cb86cF3BCF73414`  