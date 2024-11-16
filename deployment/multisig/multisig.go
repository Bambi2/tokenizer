// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package multisig

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

// MultisigMetaData contains all meta data concerning the Multisig contract.
var MultisigMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_numConfirmationsRequired\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_txIndex\",\"type\":\"uint256\"}],\"name\":\"confirmTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_txIndex\",\"type\":\"uint256\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_txIndex\",\"type\":\"uint256\"}],\"name\":\"getTransaction\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"numConfirmations\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTransactionCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isConfirmed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numConfirmationsRequired\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"owners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_txIndex\",\"type\":\"uint256\"}],\"name\":\"revokeConfirmation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"submitTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transactions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"numConfirmations\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b50604051612292380380612292833981810160405281019061003191906104a5565b5f825111610074576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161006b90610559565b60405180910390fd5b5f81118015610084575081518111155b6100c3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100ba906105e7565b60405180910390fd5b5f5f90505b82518110156102a8575f8382815181106100e5576100e4610605565b5b602002602001015190505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361015d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101549061067c565b60405180910390fd5b60015f8273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16156101e7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101de906106e4565b60405180910390fd5b6001805f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055505f81908060018154018082558091505060019003905f5260205f20015f9091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505080806001019150506100c8565b50806002819055505050610702565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610312826102cc565b810181811067ffffffffffffffff82111715610331576103306102dc565b5b80604052505050565b5f6103436102b7565b905061034f8282610309565b919050565b5f67ffffffffffffffff82111561036e5761036d6102dc565b5b602082029050602081019050919050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6103ac82610383565b9050919050565b6103bc816103a2565b81146103c6575f5ffd5b50565b5f815190506103d7816103b3565b92915050565b5f6103ef6103ea84610354565b61033a565b905080838252602082019050602084028301858111156104125761041161037f565b5b835b8181101561043b578061042788826103c9565b845260208401935050602081019050610414565b5050509392505050565b5f82601f830112610459576104586102c8565b5b81516104698482602086016103dd565b91505092915050565b5f819050919050565b61048481610472565b811461048e575f5ffd5b50565b5f8151905061049f8161047b565b92915050565b5f5f604083850312156104bb576104ba6102c0565b5b5f83015167ffffffffffffffff8111156104d8576104d76102c4565b5b6104e485828601610445565b92505060206104f585828601610491565b9150509250929050565b5f82825260208201905092915050565b7f6f776e65727320726571756972656400000000000000000000000000000000005f82015250565b5f610543600f836104ff565b915061054e8261050f565b602082019050919050565b5f6020820190508181035f83015261057081610537565b9050919050565b7f696e76616c6964206e756d626572206f6620726571756972656420636f6e66695f8201527f726d6174696f6e73000000000000000000000000000000000000000000000000602082015250565b5f6105d16028836104ff565b91506105dc82610577565b604082019050919050565b5f6020820190508181035f8301526105fe816105c5565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f696e76616c6964206f776e6572000000000000000000000000000000000000005f82015250565b5f610666600d836104ff565b915061067182610632565b602082019050919050565b5f6020820190508181035f8301526106938161065a565b9050919050565b7f6f776e6572206e6f7420756e69717565000000000000000000000000000000005f82015250565b5f6106ce6010836104ff565b91506106d98261069a565b602082019050919050565b5f6020820190508181035f8301526106fb816106c2565b9050919050565b611b838061070f5f395ff3fe6080604052600436106100aa575f3560e01c80639ace38c2116100635780639ace38c2146101fb578063a0e67e2b1461023b578063c01a8c8414610265578063c64274741461028d578063d0549b85146102b5578063ee22610b146102df576100b1565b8063025e7c27146100b557806320ea8d86146100f15780632e7700f0146101195780632f54bf6e1461014357806333ea3dc81461017f57806380f59a65146101bf576100b1565b366100b157005b5f5ffd5b3480156100c0575f5ffd5b506100db60048036038101906100d69190610f91565b610307565b6040516100e89190610ffb565b60405180910390f35b3480156100fc575f5ffd5b5061011760048036038101906101129190610f91565b610341565b005b348015610124575f5ffd5b5061012d6105c2565b60405161013a9190611023565b60405180910390f35b34801561014e575f5ffd5b5061016960048036038101906101649190611066565b6105ce565b60405161017691906110ab565b60405180910390f35b34801561018a575f5ffd5b506101a560048036038101906101a09190610f91565b6105eb565b6040516101b6959493929190611134565b60405180910390f35b3480156101ca575f5ffd5b506101e560048036038101906101e0919061118c565b6106f4565b6040516101f291906110ab565b60405180910390f35b348015610206575f5ffd5b50610221600480360381019061021c9190610f91565b61071e565b604051610232959493929190611134565b60405180910390f35b348015610246575f5ffd5b5061024f610810565b60405161025c9190611281565b60405180910390f35b348015610270575f5ffd5b5061028b60048036038101906102869190610f91565b61089a565b005b348015610298575f5ffd5b506102b360048036038101906102ae91906113cd565b610b1f565b005b3480156102c0575f5ffd5b506102c9610ca3565b6040516102d69190611023565b60405180910390f35b3480156102ea575f5ffd5b5061030560048036038101906103009190610f91565b610ca9565b005b5f8181548110610315575f80fd5b905f5260205f20015f915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff166103ca576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103c190611493565b60405180910390fd5b806004805490508110610412576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610409906114fb565b60405180910390fd5b816004818154811061042757610426611519565b5b905f5260205f2090600502016003015f9054906101000a900460ff1615610483576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161047a90611590565b60405180910390fd5b5f6004848154811061049857610497611519565b5b905f5260205f209060050201905060035f8581526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff1661053e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610535906115f8565b60405180910390fd5b6001816004015f8282546105529190611643565b925050819055505f60035f8681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff02191690831515021790555050505050565b5f600480549050905090565b6001602052805f5260405f205f915054906101000a900460ff1681565b5f5f60605f5f5f6004878154811061060657610605611519565b5b905f5260205f2090600502019050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16816001015482600201836003015f9054906101000a900460ff168460040154828054610662906116a3565b80601f016020809104026020016040519081016040528092919081815260200182805461068e906116a3565b80156106d95780601f106106b0576101008083540402835291602001916106d9565b820191905f5260205f20905b8154815290600101906020018083116106bc57829003601f168201915b50505050509250955095509550955095505091939590929450565b6003602052815f5260405f20602052805f5260405f205f915091509054906101000a900460ff1681565b6004818154811061072d575f80fd5b905f5260205f2090600502015f91509050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001015490806002018054610777906116a3565b80601f01602080910402602001604051908101604052809291908181526020018280546107a3906116a3565b80156107ee5780601f106107c5576101008083540402835291602001916107ee565b820191905f5260205f20905b8154815290600101906020018083116107d157829003601f168201915b505050505090806003015f9054906101000a900460ff16908060040154905085565b60605f80548060200260200160405190810160405280929190818152602001828054801561089057602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610847575b5050505050905090565b60015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16610923576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161091a90611493565b60405180910390fd5b80600480549050811061096b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610962906114fb565b60405180910390fd5b81600481815481106109805761097f611519565b5b905f5260205f2090600502016003015f9054906101000a900460ff16156109dc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109d390611590565b60405180910390fd5b8260035f8281526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff1615610a76576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a6d9061171d565b60405180910390fd5b5f60048581548110610a8b57610a8a611519565b5b905f5260205f20906005020190506001816004015f828254610aad919061173b565b92505081905550600160035f8781526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055505050505050565b60015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16610ba8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b9f90611493565b60405180910390fd5b60046040518060a001604052808573ffffffffffffffffffffffffffffffffffffffff1681526020018481526020018381526020015f151581526020015f815250908060018154018082558091505060019003905f5260205f2090600502015f909190919091505f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101556040820151816002019081610c72919061190e565b506060820151816003015f6101000a81548160ff021916908315150217905550608082015181600401555050505050565b60025481565b60015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16610d32576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d2990611493565b60405180910390fd5b806004805490508110610d7a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d71906114fb565b60405180910390fd5b8160048181548110610d8f57610d8e611519565b5b905f5260205f2090600502016003015f9054906101000a900460ff1615610deb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610de290611590565b60405180910390fd5b5f60048481548110610e0057610dff611519565b5b905f5260205f209060050201905060025481600401541015610e57576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e4e90611a27565b60405180910390fd5b6001816003015f6101000a81548160ff0219169083151502179055505f815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16826001015483600201604051610ec39190611acf565b5f6040518083038185875af1925050503d805f8114610efd576040519150601f19603f3d011682016040523d82523d5f602084013e610f02565b606091505b5050905080610f46576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f3d90611b2f565b60405180910390fd5b5050505050565b5f604051905090565b5f5ffd5b5f5ffd5b5f819050919050565b610f7081610f5e565b8114610f7a575f5ffd5b50565b5f81359050610f8b81610f67565b92915050565b5f60208284031215610fa657610fa5610f56565b5b5f610fb384828501610f7d565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610fe582610fbc565b9050919050565b610ff581610fdb565b82525050565b5f60208201905061100e5f830184610fec565b92915050565b61101d81610f5e565b82525050565b5f6020820190506110365f830184611014565b92915050565b61104581610fdb565b811461104f575f5ffd5b50565b5f813590506110608161103c565b92915050565b5f6020828403121561107b5761107a610f56565b5b5f61108884828501611052565b91505092915050565b5f8115159050919050565b6110a581611091565b82525050565b5f6020820190506110be5f83018461109c565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f611106826110c4565b61111081856110ce565b93506111208185602086016110de565b611129816110ec565b840191505092915050565b5f60a0820190506111475f830188610fec565b6111546020830187611014565b818103604083015261116681866110fc565b9050611175606083018561109c565b6111826080830184611014565b9695505050505050565b5f5f604083850312156111a2576111a1610f56565b5b5f6111af85828601610f7d565b92505060206111c085828601611052565b9150509250929050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b6111fc81610fdb565b82525050565b5f61120d83836111f3565b60208301905092915050565b5f602082019050919050565b5f61122f826111ca565b61123981856111d4565b9350611244836111e4565b805f5b8381101561127457815161125b8882611202565b975061126683611219565b925050600181019050611247565b5085935050505092915050565b5f6020820190508181035f8301526112998184611225565b905092915050565b5f5ffd5b5f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6112df826110ec565b810181811067ffffffffffffffff821117156112fe576112fd6112a9565b5b80604052505050565b5f611310610f4d565b905061131c82826112d6565b919050565b5f67ffffffffffffffff82111561133b5761133a6112a9565b5b611344826110ec565b9050602081019050919050565b828183375f83830152505050565b5f61137161136c84611321565b611307565b90508281526020810184848401111561138d5761138c6112a5565b5b611398848285611351565b509392505050565b5f82601f8301126113b4576113b36112a1565b5b81356113c484826020860161135f565b91505092915050565b5f5f5f606084860312156113e4576113e3610f56565b5b5f6113f186828701611052565b935050602061140286828701610f7d565b925050604084013567ffffffffffffffff81111561142357611422610f5a565b5b61142f868287016113a0565b9150509250925092565b5f82825260208201905092915050565b7f6e6f74206f776e657200000000000000000000000000000000000000000000005f82015250565b5f61147d600983611439565b915061148882611449565b602082019050919050565b5f6020820190508181035f8301526114aa81611471565b9050919050565b7f747820646f6573206e6f742065786973740000000000000000000000000000005f82015250565b5f6114e5601183611439565b91506114f0826114b1565b602082019050919050565b5f6020820190508181035f830152611512816114d9565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f747820616c7265616479206578656375746564000000000000000000000000005f82015250565b5f61157a601383611439565b915061158582611546565b602082019050919050565b5f6020820190508181035f8301526115a78161156e565b9050919050565b7f7478206e6f7420636f6e6669726d6564000000000000000000000000000000005f82015250565b5f6115e2601083611439565b91506115ed826115ae565b602082019050919050565b5f6020820190508181035f83015261160f816115d6565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61164d82610f5e565b915061165883610f5e565b92508282039050818111156116705761166f611616565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806116ba57607f821691505b6020821081036116cd576116cc611676565b5b50919050565b7f747820616c726561647920636f6e6669726d65640000000000000000000000005f82015250565b5f611707601483611439565b9150611712826116d3565b602082019050919050565b5f6020820190508181035f830152611734816116fb565b9050919050565b5f61174582610f5e565b915061175083610f5e565b925082820190508082111561176857611767611616565b5b92915050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026117ca7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261178f565b6117d4868361178f565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61180f61180a61180584610f5e565b6117ec565b610f5e565b9050919050565b5f819050919050565b611828836117f5565b61183c61183482611816565b84845461179b565b825550505050565b5f5f905090565b611853611844565b61185e81848461181f565b505050565b5b81811015611881576118765f8261184b565b600181019050611864565b5050565b601f8211156118c6576118978161176e565b6118a084611780565b810160208510156118af578190505b6118c36118bb85611780565b830182611863565b50505b505050565b5f82821c905092915050565b5f6118e65f19846008026118cb565b1980831691505092915050565b5f6118fe83836118d7565b9150826002028217905092915050565b611917826110c4565b67ffffffffffffffff8111156119305761192f6112a9565b5b61193a82546116a3565b611945828285611885565b5f60209050601f831160018114611976575f8415611964578287015190505b61196e85826118f3565b8655506119d5565b601f1984166119848661176e565b5f5b828110156119ab57848901518255600182019150602085019450602081019050611986565b868310156119c857848901516119c4601f8916826118d7565b8355505b6001600288020188555050505b505050505050565b7f63616e6e6f7420657865637574652074780000000000000000000000000000005f82015250565b5f611a11601183611439565b9150611a1c826119dd565b602082019050919050565b5f6020820190508181035f830152611a3e81611a05565b9050919050565b5f81905092915050565b5f8154611a5b816116a3565b611a658186611a45565b9450600182165f8114611a7f5760018114611a9457611ac6565b60ff1983168652811515820286019350611ac6565b611a9d8561176e565b5f5b83811015611abe57815481890152600182019150602081019050611a9f565b838801955050505b50505092915050565b5f611ada8284611a4f565b915081905092915050565b7f7478206661696c656400000000000000000000000000000000000000000000005f82015250565b5f611b19600983611439565b9150611b2482611ae5565b602082019050919050565b5f6020820190508181035f830152611b4681611b0d565b905091905056fea26469706673582212203d9437b2280e87a5e2ffe738aa254c51d22ed5247b7a94fa04d9fbb4ee610b5e64736f6c634300081c0033",
}

// MultisigABI is the input ABI used to generate the binding from.
// Deprecated: Use MultisigMetaData.ABI instead.
var MultisigABI = MultisigMetaData.ABI

// MultisigBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MultisigMetaData.Bin instead.
var MultisigBin = MultisigMetaData.Bin

// DeployMultisig deploys a new Ethereum contract, binding an instance of Multisig to it.
func DeployMultisig(auth *bind.TransactOpts, backend bind.ContractBackend, _owners []common.Address, _numConfirmationsRequired *big.Int) (common.Address, *types.Transaction, *Multisig, error) {
	parsed, err := MultisigMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MultisigBin), backend, _owners, _numConfirmationsRequired)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Multisig{MultisigCaller: MultisigCaller{contract: contract}, MultisigTransactor: MultisigTransactor{contract: contract}, MultisigFilterer: MultisigFilterer{contract: contract}}, nil
}

// Multisig is an auto generated Go binding around an Ethereum contract.
type Multisig struct {
	MultisigCaller     // Read-only binding to the contract
	MultisigTransactor // Write-only binding to the contract
	MultisigFilterer   // Log filterer for contract events
}

// MultisigCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultisigCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultisigTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultisigTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultisigFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultisigFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultisigSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultisigSession struct {
	Contract     *Multisig         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MultisigCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultisigCallerSession struct {
	Contract *MultisigCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// MultisigTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultisigTransactorSession struct {
	Contract     *MultisigTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MultisigRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultisigRaw struct {
	Contract *Multisig // Generic contract binding to access the raw methods on
}

// MultisigCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultisigCallerRaw struct {
	Contract *MultisigCaller // Generic read-only contract binding to access the raw methods on
}

// MultisigTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultisigTransactorRaw struct {
	Contract *MultisigTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMultisig creates a new instance of Multisig, bound to a specific deployed contract.
func NewMultisig(address common.Address, backend bind.ContractBackend) (*Multisig, error) {
	contract, err := bindMultisig(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Multisig{MultisigCaller: MultisigCaller{contract: contract}, MultisigTransactor: MultisigTransactor{contract: contract}, MultisigFilterer: MultisigFilterer{contract: contract}}, nil
}

// NewMultisigCaller creates a new read-only instance of Multisig, bound to a specific deployed contract.
func NewMultisigCaller(address common.Address, caller bind.ContractCaller) (*MultisigCaller, error) {
	contract, err := bindMultisig(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultisigCaller{contract: contract}, nil
}

// NewMultisigTransactor creates a new write-only instance of Multisig, bound to a specific deployed contract.
func NewMultisigTransactor(address common.Address, transactor bind.ContractTransactor) (*MultisigTransactor, error) {
	contract, err := bindMultisig(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultisigTransactor{contract: contract}, nil
}

// NewMultisigFilterer creates a new log filterer instance of Multisig, bound to a specific deployed contract.
func NewMultisigFilterer(address common.Address, filterer bind.ContractFilterer) (*MultisigFilterer, error) {
	contract, err := bindMultisig(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultisigFilterer{contract: contract}, nil
}

// bindMultisig binds a generic wrapper to an already deployed contract.
func bindMultisig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MultisigMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multisig *MultisigRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multisig.Contract.MultisigCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multisig *MultisigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multisig.Contract.MultisigTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multisig *MultisigRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multisig.Contract.MultisigTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multisig *MultisigCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multisig.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multisig *MultisigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multisig.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multisig *MultisigTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multisig.Contract.contract.Transact(opts, method, params...)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_Multisig *MultisigCaller) GetOwners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Multisig.contract.Call(opts, &out, "getOwners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_Multisig *MultisigSession) GetOwners() ([]common.Address, error) {
	return _Multisig.Contract.GetOwners(&_Multisig.CallOpts)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_Multisig *MultisigCallerSession) GetOwners() ([]common.Address, error) {
	return _Multisig.Contract.GetOwners(&_Multisig.CallOpts)
}

// GetTransaction is a free data retrieval call binding the contract method 0x33ea3dc8.
//
// Solidity: function getTransaction(uint256 _txIndex) view returns(address to, uint256 value, bytes data, bool executed, uint256 numConfirmations)
func (_Multisig *MultisigCaller) GetTransaction(opts *bind.CallOpts, _txIndex *big.Int) (struct {
	To               common.Address
	Value            *big.Int
	Data             []byte
	Executed         bool
	NumConfirmations *big.Int
}, error) {
	var out []interface{}
	err := _Multisig.contract.Call(opts, &out, "getTransaction", _txIndex)

	outstruct := new(struct {
		To               common.Address
		Value            *big.Int
		Data             []byte
		Executed         bool
		NumConfirmations *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.To = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Value = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Data = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.Executed = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.NumConfirmations = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetTransaction is a free data retrieval call binding the contract method 0x33ea3dc8.
//
// Solidity: function getTransaction(uint256 _txIndex) view returns(address to, uint256 value, bytes data, bool executed, uint256 numConfirmations)
func (_Multisig *MultisigSession) GetTransaction(_txIndex *big.Int) (struct {
	To               common.Address
	Value            *big.Int
	Data             []byte
	Executed         bool
	NumConfirmations *big.Int
}, error) {
	return _Multisig.Contract.GetTransaction(&_Multisig.CallOpts, _txIndex)
}

// GetTransaction is a free data retrieval call binding the contract method 0x33ea3dc8.
//
// Solidity: function getTransaction(uint256 _txIndex) view returns(address to, uint256 value, bytes data, bool executed, uint256 numConfirmations)
func (_Multisig *MultisigCallerSession) GetTransaction(_txIndex *big.Int) (struct {
	To               common.Address
	Value            *big.Int
	Data             []byte
	Executed         bool
	NumConfirmations *big.Int
}, error) {
	return _Multisig.Contract.GetTransaction(&_Multisig.CallOpts, _txIndex)
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x2e7700f0.
//
// Solidity: function getTransactionCount() view returns(uint256)
func (_Multisig *MultisigCaller) GetTransactionCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Multisig.contract.Call(opts, &out, "getTransactionCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTransactionCount is a free data retrieval call binding the contract method 0x2e7700f0.
//
// Solidity: function getTransactionCount() view returns(uint256)
func (_Multisig *MultisigSession) GetTransactionCount() (*big.Int, error) {
	return _Multisig.Contract.GetTransactionCount(&_Multisig.CallOpts)
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x2e7700f0.
//
// Solidity: function getTransactionCount() view returns(uint256)
func (_Multisig *MultisigCallerSession) GetTransactionCount() (*big.Int, error) {
	return _Multisig.Contract.GetTransactionCount(&_Multisig.CallOpts)
}

// IsConfirmed is a free data retrieval call binding the contract method 0x80f59a65.
//
// Solidity: function isConfirmed(uint256 , address ) view returns(bool)
func (_Multisig *MultisigCaller) IsConfirmed(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Multisig.contract.Call(opts, &out, "isConfirmed", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsConfirmed is a free data retrieval call binding the contract method 0x80f59a65.
//
// Solidity: function isConfirmed(uint256 , address ) view returns(bool)
func (_Multisig *MultisigSession) IsConfirmed(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _Multisig.Contract.IsConfirmed(&_Multisig.CallOpts, arg0, arg1)
}

// IsConfirmed is a free data retrieval call binding the contract method 0x80f59a65.
//
// Solidity: function isConfirmed(uint256 , address ) view returns(bool)
func (_Multisig *MultisigCallerSession) IsConfirmed(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _Multisig.Contract.IsConfirmed(&_Multisig.CallOpts, arg0, arg1)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) view returns(bool)
func (_Multisig *MultisigCaller) IsOwner(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Multisig.contract.Call(opts, &out, "isOwner", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) view returns(bool)
func (_Multisig *MultisigSession) IsOwner(arg0 common.Address) (bool, error) {
	return _Multisig.Contract.IsOwner(&_Multisig.CallOpts, arg0)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) view returns(bool)
func (_Multisig *MultisigCallerSession) IsOwner(arg0 common.Address) (bool, error) {
	return _Multisig.Contract.IsOwner(&_Multisig.CallOpts, arg0)
}

// NumConfirmationsRequired is a free data retrieval call binding the contract method 0xd0549b85.
//
// Solidity: function numConfirmationsRequired() view returns(uint256)
func (_Multisig *MultisigCaller) NumConfirmationsRequired(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Multisig.contract.Call(opts, &out, "numConfirmationsRequired")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumConfirmationsRequired is a free data retrieval call binding the contract method 0xd0549b85.
//
// Solidity: function numConfirmationsRequired() view returns(uint256)
func (_Multisig *MultisigSession) NumConfirmationsRequired() (*big.Int, error) {
	return _Multisig.Contract.NumConfirmationsRequired(&_Multisig.CallOpts)
}

// NumConfirmationsRequired is a free data retrieval call binding the contract method 0xd0549b85.
//
// Solidity: function numConfirmationsRequired() view returns(uint256)
func (_Multisig *MultisigCallerSession) NumConfirmationsRequired() (*big.Int, error) {
	return _Multisig.Contract.NumConfirmationsRequired(&_Multisig.CallOpts)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_Multisig *MultisigCaller) Owners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Multisig.contract.Call(opts, &out, "owners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_Multisig *MultisigSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _Multisig.Contract.Owners(&_Multisig.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_Multisig *MultisigCallerSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _Multisig.Contract.Owners(&_Multisig.CallOpts, arg0)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(address to, uint256 value, bytes data, bool executed, uint256 numConfirmations)
func (_Multisig *MultisigCaller) Transactions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	To               common.Address
	Value            *big.Int
	Data             []byte
	Executed         bool
	NumConfirmations *big.Int
}, error) {
	var out []interface{}
	err := _Multisig.contract.Call(opts, &out, "transactions", arg0)

	outstruct := new(struct {
		To               common.Address
		Value            *big.Int
		Data             []byte
		Executed         bool
		NumConfirmations *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.To = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Value = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Data = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.Executed = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.NumConfirmations = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(address to, uint256 value, bytes data, bool executed, uint256 numConfirmations)
func (_Multisig *MultisigSession) Transactions(arg0 *big.Int) (struct {
	To               common.Address
	Value            *big.Int
	Data             []byte
	Executed         bool
	NumConfirmations *big.Int
}, error) {
	return _Multisig.Contract.Transactions(&_Multisig.CallOpts, arg0)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(address to, uint256 value, bytes data, bool executed, uint256 numConfirmations)
func (_Multisig *MultisigCallerSession) Transactions(arg0 *big.Int) (struct {
	To               common.Address
	Value            *big.Int
	Data             []byte
	Executed         bool
	NumConfirmations *big.Int
}, error) {
	return _Multisig.Contract.Transactions(&_Multisig.CallOpts, arg0)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 _txIndex) returns()
func (_Multisig *MultisigTransactor) ConfirmTransaction(opts *bind.TransactOpts, _txIndex *big.Int) (*types.Transaction, error) {
	return _Multisig.contract.Transact(opts, "confirmTransaction", _txIndex)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 _txIndex) returns()
func (_Multisig *MultisigSession) ConfirmTransaction(_txIndex *big.Int) (*types.Transaction, error) {
	return _Multisig.Contract.ConfirmTransaction(&_Multisig.TransactOpts, _txIndex)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 _txIndex) returns()
func (_Multisig *MultisigTransactorSession) ConfirmTransaction(_txIndex *big.Int) (*types.Transaction, error) {
	return _Multisig.Contract.ConfirmTransaction(&_Multisig.TransactOpts, _txIndex)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 _txIndex) returns()
func (_Multisig *MultisigTransactor) ExecuteTransaction(opts *bind.TransactOpts, _txIndex *big.Int) (*types.Transaction, error) {
	return _Multisig.contract.Transact(opts, "executeTransaction", _txIndex)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 _txIndex) returns()
func (_Multisig *MultisigSession) ExecuteTransaction(_txIndex *big.Int) (*types.Transaction, error) {
	return _Multisig.Contract.ExecuteTransaction(&_Multisig.TransactOpts, _txIndex)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 _txIndex) returns()
func (_Multisig *MultisigTransactorSession) ExecuteTransaction(_txIndex *big.Int) (*types.Transaction, error) {
	return _Multisig.Contract.ExecuteTransaction(&_Multisig.TransactOpts, _txIndex)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 _txIndex) returns()
func (_Multisig *MultisigTransactor) RevokeConfirmation(opts *bind.TransactOpts, _txIndex *big.Int) (*types.Transaction, error) {
	return _Multisig.contract.Transact(opts, "revokeConfirmation", _txIndex)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 _txIndex) returns()
func (_Multisig *MultisigSession) RevokeConfirmation(_txIndex *big.Int) (*types.Transaction, error) {
	return _Multisig.Contract.RevokeConfirmation(&_Multisig.TransactOpts, _txIndex)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 _txIndex) returns()
func (_Multisig *MultisigTransactorSession) RevokeConfirmation(_txIndex *big.Int) (*types.Transaction, error) {
	return _Multisig.Contract.RevokeConfirmation(&_Multisig.TransactOpts, _txIndex)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address _to, uint256 _value, bytes _data) returns()
func (_Multisig *MultisigTransactor) SubmitTransaction(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _Multisig.contract.Transact(opts, "submitTransaction", _to, _value, _data)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address _to, uint256 _value, bytes _data) returns()
func (_Multisig *MultisigSession) SubmitTransaction(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _Multisig.Contract.SubmitTransaction(&_Multisig.TransactOpts, _to, _value, _data)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address _to, uint256 _value, bytes _data) returns()
func (_Multisig *MultisigTransactorSession) SubmitTransaction(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _Multisig.Contract.SubmitTransaction(&_Multisig.TransactOpts, _to, _value, _data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Multisig *MultisigTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multisig.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Multisig *MultisigSession) Receive() (*types.Transaction, error) {
	return _Multisig.Contract.Receive(&_Multisig.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Multisig *MultisigTransactorSession) Receive() (*types.Transaction, error) {
	return _Multisig.Contract.Receive(&_Multisig.TransactOpts)
}
