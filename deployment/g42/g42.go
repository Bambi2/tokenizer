// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package g42

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// G42MetaData contains all meta data concerning the G42 contract.
var G42MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_authority\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b50604051611c4b380380611c4b83398181016040528101906100319190610443565b6040518060400160405280600b81526020017f477269676f7269616e34320000000000000000000000000000000000000000008152506040518060400160405280600381526020017f47343200000000000000000000000000000000000000000000000000000000008152508282600390816100ad91906106ab565b5081600490816100bd91906106ab565b508060055f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050506101373361011361013d60201b60201c565b600a61011f91906108e2565b6103e861012c919061092c565b61014560201b60201c565b50610adb565b5f6012905090565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036101b3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101aa906109c7565b60405180910390fd5b6101c45f83836101c860201b60201c565b5050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610218578060025f82825461020c91906109e5565b925050819055506102e1565b805f5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015610297576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161028e90610a62565b60405180910390fd5b805f5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825403925050819055505b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610331578060025f8282546103259190610a80565b9250508190555061037b565b805f5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516103d89190610ac2565b60405180910390a3505050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610412826103e9565b9050919050565b61042281610408565b811461042c575f5ffd5b50565b5f8151905061043d81610419565b92915050565b5f60208284031215610458576104576103e5565b5b5f6104658482850161042f565b91505092915050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806104e957607f821691505b6020821081036104fc576104fb6104a5565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261055e7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610523565b6105688683610523565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6105ac6105a76105a284610580565b610589565b610580565b9050919050565b5f819050919050565b6105c583610592565b6105d96105d1826105b3565b84845461052f565b825550505050565b5f5f905090565b6105f06105e1565b6105fb8184846105bc565b505050565b5b8181101561061e576106135f826105e8565b600181019050610601565b5050565b601f8211156106635761063481610502565b61063d84610514565b8101602085101561064c578190505b61066061065885610514565b830182610600565b50505b505050565b5f82821c905092915050565b5f6106835f1984600802610668565b1980831691505092915050565b5f61069b8383610674565b9150826002028217905092915050565b6106b48261046e565b67ffffffffffffffff8111156106cd576106cc610478565b5b6106d782546104d2565b6106e2828285610622565b5f60209050601f831160018114610713575f8415610701578287015190505b61070b8582610690565b865550610772565b601f19841661072186610502565b5f5b8281101561074857848901518255600182019150602085019450602081019050610723565b868310156107655784890151610761601f891682610674565b8355505b6001600288020188555050505b505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f8160011c9050919050565b5f5f8291508390505b60018511156107fc578086048111156107d8576107d761077a565b5b60018516156107e75780820291505b80810290506107f5856107a7565b94506107bc565b94509492505050565b5f8261081457600190506108cf565b81610821575f90506108cf565b8160018114610837576002811461084157610870565b60019150506108cf565b60ff8411156108535761085261077a565b5b8360020a91508482111561086a5761086961077a565b5b506108cf565b5060208310610133831016604e8410600b84101617156108a55782820a9050838111156108a05761089f61077a565b5b6108cf565b6108b284848460016107b3565b925090508184048111156108c9576108c861077a565b5b81810290505b9392505050565b5f60ff82169050919050565b5f6108ec82610580565b91506108f7836108d6565b92506109247fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484610805565b905092915050565b5f61093682610580565b915061094183610580565b925082820261094f81610580565b915082820484148315176109665761096561077a565b5b5092915050565b5f82825260208201905092915050565b7f696e76616c6964206163636f756e7400000000000000000000000000000000005f82015250565b5f6109b1600f8361096d565b91506109bc8261097d565b602082019050919050565b5f6020820190508181035f8301526109de816109a5565b9050919050565b5f6109ef82610580565b91506109fa83610580565b9250828201905080821115610a1257610a1161077a565b5b92915050565b7f73656e6465722068617320696e73756666696369656e742066756e64730000005f82015250565b5f610a4c601d8361096d565b9150610a5782610a18565b602082019050919050565b5f6020820190508181035f830152610a7981610a40565b9050919050565b5f610a8a82610580565b9150610a9583610580565b9250828203905081811115610aad57610aac61077a565b5b92915050565b610abc81610580565b82525050565b5f602082019050610ad55f830184610ab3565b92915050565b61116380610ae85f395ff3fe608060405234801561000f575f5ffd5b506004361061009c575f3560e01c806340c10f191161006457806340c10f191461015a57806370a082311461017657806395d89b41146101a6578063a9059cbb146101c4578063dd62ed3e146101f45761009c565b806306fdde03146100a0578063095ea7b3146100be57806318160ddd146100ee57806323b872dd1461010c578063313ce5671461013c575b5f5ffd5b6100a8610224565b6040516100b59190610b2e565b60405180910390f35b6100d860048036038101906100d39190610bdf565b6102b4565b6040516100e59190610c37565b60405180910390f35b6100f66102ca565b6040516101039190610c5f565b60405180910390f35b61012660048036038101906101219190610c78565b6102d3565b6040516101339190610c37565b60405180910390f35b6101446102f5565b6040516101519190610ce3565b60405180910390f35b610174600480360381019061016f9190610bdf565b6102fd565b005b610190600480360381019061018b9190610cfc565b610409565b60405161019d9190610c5f565b60405180910390f35b6101ae61044e565b6040516101bb9190610b2e565b60405180910390f35b6101de60048036038101906101d99190610bdf565b6104de565b6040516101eb9190610c37565b60405180910390f35b61020e60048036038101906102099190610d27565b6104f4565b60405161021b9190610c5f565b60405180910390f35b60606003805461023390610d92565b80601f016020809104026020016040519081016040528092919081815260200182805461025f90610d92565b80156102aa5780601f10610281576101008083540402835291602001916102aa565b820191905f5260205f20905b81548152906001019060200180831161028d57829003601f168201915b5050505050905090565b5f6102c0338484610576565b6001905092915050565b5f600254905090565b5f6102df843384610588565b6102ea8484846105ea565b600190509392505050565b5f6012905090565b60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461038c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161038390610e0c565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036103fa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103f190610e74565b60405180910390fd5b6104055f83836106d6565b5050565b5f5f5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b60606004805461045d90610d92565b80601f016020809104026020016040519081016040528092919081815260200182805461048990610d92565b80156104d45780601f106104ab576101008083540402835291602001916104d4565b820191905f5260205f20905b8154815290600101906020018083116104b757829003601f168201915b5050505050905090565b5f6104ea3384846105ea565b6001905092915050565b5f60015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905092915050565b61058383838360016108f3565b505050565b5f61059384846104f4565b9050818110156105d8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105cf90610edc565b60405180910390fd5b6105e48484845f6108f3565b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610658576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161064f90610f44565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036106c6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106bd90610fac565b60405180910390fd5b6106d18383836106d6565b505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610726578060025f82825461071a9190610ff7565b925050819055506107ef565b805f5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205410156107a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161079c90611074565b60405180910390fd5b805f5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825403925050819055505b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361083f578060025f8282546108339190611092565b92505081905550610889565b805f5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516108e69190610c5f565b60405180910390a3505050565b5f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610961576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109589061110f565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036109cf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109c690610f44565b60405180910390fd5b8160015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508015610ab8578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610aaf9190610c5f565b60405180910390a35b50505050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610b0082610abe565b610b0a8185610ac8565b9350610b1a818560208601610ad8565b610b2381610ae6565b840191505092915050565b5f6020820190508181035f830152610b468184610af6565b905092915050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610b7b82610b52565b9050919050565b610b8b81610b71565b8114610b95575f5ffd5b50565b5f81359050610ba681610b82565b92915050565b5f819050919050565b610bbe81610bac565b8114610bc8575f5ffd5b50565b5f81359050610bd981610bb5565b92915050565b5f5f60408385031215610bf557610bf4610b4e565b5b5f610c0285828601610b98565b9250506020610c1385828601610bcb565b9150509250929050565b5f8115159050919050565b610c3181610c1d565b82525050565b5f602082019050610c4a5f830184610c28565b92915050565b610c5981610bac565b82525050565b5f602082019050610c725f830184610c50565b92915050565b5f5f5f60608486031215610c8f57610c8e610b4e565b5b5f610c9c86828701610b98565b9350506020610cad86828701610b98565b9250506040610cbe86828701610bcb565b9150509250925092565b5f60ff82169050919050565b610cdd81610cc8565b82525050565b5f602082019050610cf65f830184610cd4565b92915050565b5f60208284031215610d1157610d10610b4e565b5b5f610d1e84828501610b98565b91505092915050565b5f5f60408385031215610d3d57610d3c610b4e565b5b5f610d4a85828601610b98565b9250506020610d5b85828601610b98565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610da957607f821691505b602082108103610dbc57610dbb610d65565b5b50919050565b7f6e6f74206f776e657200000000000000000000000000000000000000000000005f82015250565b5f610df6600983610ac8565b9150610e0182610dc2565b602082019050919050565b5f6020820190508181035f830152610e2381610dea565b9050919050565b7f696e76616c6964206163636f756e7400000000000000000000000000000000005f82015250565b5f610e5e600f83610ac8565b9150610e6982610e2a565b602082019050919050565b5f6020820190508181035f830152610e8b81610e52565b9050919050565b7f6e6f7420656e6f756768206f6620616c6c6f77616e63650000000000000000005f82015250565b5f610ec6601783610ac8565b9150610ed182610e92565b602082019050919050565b5f6020820190508181035f830152610ef381610eba565b9050919050565b7f696e76616c6964207370656e64657200000000000000000000000000000000005f82015250565b5f610f2e600f83610ac8565b9150610f3982610efa565b602082019050919050565b5f6020820190508181035f830152610f5b81610f22565b9050919050565b7f696e76616c6964207265636569766572000000000000000000000000000000005f82015250565b5f610f96601083610ac8565b9150610fa182610f62565b602082019050919050565b5f6020820190508181035f830152610fc381610f8a565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61100182610bac565b915061100c83610bac565b925082820190508082111561102457611023610fca565b5b92915050565b7f73656e6465722068617320696e73756666696369656e742066756e64730000005f82015250565b5f61105e601d83610ac8565b91506110698261102a565b602082019050919050565b5f6020820190508181035f83015261108b81611052565b9050919050565b5f61109c82610bac565b91506110a783610bac565b92508282039050818111156110bf576110be610fca565b5b92915050565b7f696e76616c6964206f776e6572000000000000000000000000000000000000005f82015250565b5f6110f9600d83610ac8565b9150611104826110c5565b602082019050919050565b5f6020820190508181035f830152611126816110ed565b905091905056fea264697066735822122074dda05ddc521020917f6c6132db969c6074428320b8d6379873ed8b55eb68d464736f6c634300081c0033",
}

// G42ABI is the input ABI used to generate the binding from.
// Deprecated: Use G42MetaData.ABI instead.
var G42ABI = G42MetaData.ABI

// G42Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use G42MetaData.Bin instead.
var G42Bin = G42MetaData.Bin

// DeployG42 deploys a new Ethereum contract, binding an instance of G42 to it.
func DeployG42(auth *bind.TransactOpts, backend bind.ContractBackend, _authority common.Address) (common.Address, *types.Transaction, *G42, error) {
	parsed, err := G42MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(G42Bin), backend, _authority)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &G42{G42Caller: G42Caller{contract: contract}, G42Transactor: G42Transactor{contract: contract}, G42Filterer: G42Filterer{contract: contract}}, nil
}

// G42 is an auto generated Go binding around an Ethereum contract.
type G42 struct {
	G42Caller     // Read-only binding to the contract
	G42Transactor // Write-only binding to the contract
	G42Filterer   // Log filterer for contract events
}

// G42Caller is an auto generated read-only Go binding around an Ethereum contract.
type G42Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// G42Transactor is an auto generated write-only Go binding around an Ethereum contract.
type G42Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// G42Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type G42Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// G42Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type G42Session struct {
	Contract     *G42              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// G42CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type G42CallerSession struct {
	Contract *G42Caller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// G42TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type G42TransactorSession struct {
	Contract     *G42Transactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// G42Raw is an auto generated low-level Go binding around an Ethereum contract.
type G42Raw struct {
	Contract *G42 // Generic contract binding to access the raw methods on
}

// G42CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type G42CallerRaw struct {
	Contract *G42Caller // Generic read-only contract binding to access the raw methods on
}

// G42TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type G42TransactorRaw struct {
	Contract *G42Transactor // Generic write-only contract binding to access the raw methods on
}

// NewG42 creates a new instance of G42, bound to a specific deployed contract.
func NewG42(address common.Address, backend bind.ContractBackend) (*G42, error) {
	contract, err := bindG42(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &G42{G42Caller: G42Caller{contract: contract}, G42Transactor: G42Transactor{contract: contract}, G42Filterer: G42Filterer{contract: contract}}, nil
}

// NewG42Caller creates a new read-only instance of G42, bound to a specific deployed contract.
func NewG42Caller(address common.Address, caller bind.ContractCaller) (*G42Caller, error) {
	contract, err := bindG42(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &G42Caller{contract: contract}, nil
}

// NewG42Transactor creates a new write-only instance of G42, bound to a specific deployed contract.
func NewG42Transactor(address common.Address, transactor bind.ContractTransactor) (*G42Transactor, error) {
	contract, err := bindG42(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &G42Transactor{contract: contract}, nil
}

// NewG42Filterer creates a new log filterer instance of G42, bound to a specific deployed contract.
func NewG42Filterer(address common.Address, filterer bind.ContractFilterer) (*G42Filterer, error) {
	contract, err := bindG42(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &G42Filterer{contract: contract}, nil
}

// bindG42 binds a generic wrapper to an already deployed contract.
func bindG42(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := G42MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_G42 *G42Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _G42.Contract.G42Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_G42 *G42Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _G42.Contract.G42Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_G42 *G42Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _G42.Contract.G42Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_G42 *G42CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _G42.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_G42 *G42TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _G42.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_G42 *G42TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _G42.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_G42 *G42Caller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _G42.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_G42 *G42Session) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _G42.Contract.Allowance(&_G42.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_G42 *G42CallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _G42.Contract.Allowance(&_G42.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_G42 *G42Caller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _G42.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_G42 *G42Session) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _G42.Contract.BalanceOf(&_G42.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_G42 *G42CallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _G42.Contract.BalanceOf(&_G42.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_G42 *G42Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _G42.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_G42 *G42Session) Decimals() (uint8, error) {
	return _G42.Contract.Decimals(&_G42.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_G42 *G42CallerSession) Decimals() (uint8, error) {
	return _G42.Contract.Decimals(&_G42.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_G42 *G42Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _G42.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_G42 *G42Session) Name() (string, error) {
	return _G42.Contract.Name(&_G42.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_G42 *G42CallerSession) Name() (string, error) {
	return _G42.Contract.Name(&_G42.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_G42 *G42Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _G42.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_G42 *G42Session) Symbol() (string, error) {
	return _G42.Contract.Symbol(&_G42.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_G42 *G42CallerSession) Symbol() (string, error) {
	return _G42.Contract.Symbol(&_G42.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_G42 *G42Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _G42.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_G42 *G42Session) TotalSupply() (*big.Int, error) {
	return _G42.Contract.TotalSupply(&_G42.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_G42 *G42CallerSession) TotalSupply() (*big.Int, error) {
	return _G42.Contract.TotalSupply(&_G42.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_G42 *G42Transactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _G42.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_G42 *G42Session) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _G42.Contract.Approve(&_G42.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_G42 *G42TransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _G42.Contract.Approve(&_G42.TransactOpts, _spender, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _value) returns()
func (_G42 *G42Transactor) Mint(opts *bind.TransactOpts, _account common.Address, _value *big.Int) (*types.Transaction, error) {
	return _G42.contract.Transact(opts, "mint", _account, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _value) returns()
func (_G42 *G42Session) Mint(_account common.Address, _value *big.Int) (*types.Transaction, error) {
	return _G42.Contract.Mint(&_G42.TransactOpts, _account, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _value) returns()
func (_G42 *G42TransactorSession) Mint(_account common.Address, _value *big.Int) (*types.Transaction, error) {
	return _G42.Contract.Mint(&_G42.TransactOpts, _account, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_G42 *G42Transactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _G42.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_G42 *G42Session) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _G42.Contract.Transfer(&_G42.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_G42 *G42TransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _G42.Contract.Transfer(&_G42.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_G42 *G42Transactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _G42.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_G42 *G42Session) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _G42.Contract.TransferFrom(&_G42.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_G42 *G42TransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _G42.Contract.TransferFrom(&_G42.TransactOpts, _from, _to, _value)
}

// G42ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the G42 contract.
type G42ApprovalIterator struct {
	Event *G42Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *G42ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(G42Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(G42Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *G42ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *G42ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// G42Approval represents a Approval event raised by the G42 contract.
type G42Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_G42 *G42Filterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*G42ApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _G42.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &G42ApprovalIterator{contract: _G42.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_G42 *G42Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *G42Approval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _G42.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(G42Approval)
				if err := _G42.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_G42 *G42Filterer) ParseApproval(log types.Log) (*G42Approval, error) {
	event := new(G42Approval)
	if err := _G42.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// G42TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the G42 contract.
type G42TransferIterator struct {
	Event *G42Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *G42TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(G42Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(G42Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *G42TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *G42TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// G42Transfer represents a Transfer event raised by the G42 contract.
type G42Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_G42 *G42Filterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*G42TransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _G42.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &G42TransferIterator{contract: _G42.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_G42 *G42Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *G42Transfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _G42.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(G42Transfer)
				if err := _G42.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_G42 *G42Filterer) ParseTransfer(log types.Log) (*G42Transfer, error) {
	event := new(G42Transfer)
	if err := _G42.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
