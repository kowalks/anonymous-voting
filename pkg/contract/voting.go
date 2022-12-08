// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// Point is an auto generated low-level Go binding around an user-defined struct.
type Point struct {
	X *big.Int
	Y *big.Int
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"name\":\"addCandidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"name\":\"addVoter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"name\":\"announcePublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ballotCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ballots\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structPoint\",\"name\":\"SA\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structPoint\",\"name\":\"R\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"signature\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"candidateCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"candidates\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structPoint\",\"name\":\"pubkey\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPublicKey\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"managerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"managers\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structPoint\",\"name\":\"pubkey\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structPoint\",\"name\":\"announcedKey\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"privateKey\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"phase\",\"outputs\":[{\"internalType\":\"enumPhase\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"name\":\"registerManager\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"}],\"name\":\"revealPublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"SAx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"SAy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Rx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Ry\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"signature\",\"type\":\"string\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voterCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"voters\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structPoint\",\"name\":\"pubkey\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405260008060006101000a81548160ff0219169083600281111561002957610028610054565b5b0217905550600060055560006006556000600755600060085534801561004e57600080fd5b50610083565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b611981806100926000396000f3fe6080604052600436106100f35760003560e01c806375eac54d1161008a578063b1c9fe6e11610059578063b1c9fe6e1461032f578063c00300321461035a578063da58c7d914610385578063e9857485146103c2576100f3565b806375eac54d146102895780639d578eca146102b25780639ffdcc59146102db578063a9a981a314610304576100f3565b80633477ee2e116100c65780633477ee2e146101a15780633d39c260146101df57806342169e481461021f5780635c632b381461024a576100f3565b80630a120bb0146100f8578063103b7397146101215780631970e4461461014c5780632e33445214610175575b600080fd5b34801561010457600080fd5b5061011f600480360381019061011a9190610edc565b6103de565b005b34801561012d57600080fd5b50610136610492565b6040516101439190610f2b565b60405180910390f35b34801561015857600080fd5b50610173600480360381019061016e919061108c565b610498565b005b34801561018157600080fd5b5061018a610569565b6040516101989291906110fb565b60405180910390f35b3480156101ad57600080fd5b506101c860048036038101906101c39190611124565b6106b5565b6040516101d692919061120e565b60405180910390f35b3480156101eb57600080fd5b5061020660048036038101906102019190611124565b61077f565b604051610216949392919061127f565b60405180910390f35b34801561022b57600080fd5b5061023461080b565b6040516102419190610f2b565b60405180910390f35b34801561025657600080fd5b50610271600480360381019061026c9190611124565b610811565b604051610280939291906112c4565b60405180910390f35b34801561029557600080fd5b506102b060048036038101906102ab9190611302565b6108ff565b005b3480156102be57600080fd5b506102d960048036038101906102d49190611124565b610a5d565b005b3480156102e757600080fd5b5061030260048036038101906102fd9190610edc565b610b7c565b005b34801561031057600080fd5b50610319610c8b565b6040516103269190610f2b565b60405180910390f35b34801561033b57600080fd5b50610344610c91565b6040516103519190611410565b60405180910390f35b34801561036657600080fd5b5061036f610ca2565b60405161037c9190610f2b565b60405180910390f35b34801561039157600080fd5b506103ac60048036038101906103a79190611124565b610ca8565b6040516103b9919061142b565b60405180910390f35b6103dc60048036038101906103d79190610edc565b610ce4565b005b600060028111156103f2576103f1611399565b5b60008054906101000a900460ff16600281111561041257610411611399565b5b1461041c57600080fd5b6040518060200160405280604051806040016040528085815260200184815250815250600260006006548152602001908152602001600020600082015181600001600082015181600001556020820151816001015550509050506006600081548092919061048990611475565b91905055505050565b60085481565b600060028111156104ac576104ab611399565b5b60008054906101000a900460ff1660028111156104cc576104cb611399565b5b146104d657600080fd5b6040518060400160405280848152602001604051806040016040528085815260200184815250815250600160006005548152602001908152602001600020600082015181600001908161052991906116c9565b50602082015181600101600082015181600001556020820151816001015550509050506005600081548092919061055f90611475565b9190505550505050565b60008060006004600080815260200190815260200160002060030160000154905060006004600080815260200190815260200160002060030160010154905060005b60085481101561067e5782600460008381526020019081526020016000206003016000015403610610576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610607906117e7565b60405180910390fd5b8160046000838152602001908152602001600020600301600101540361066b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610662906117e7565b60405180910390fd5b808061067690611475565b9150506105ab565b5060016000806101000a81548160ff021916908360028111156106a4576106a3611399565b5b021790555081819350935050509091565b60016020528060005260406000206000915090508060000180546106d8906114ec565b80601f0160208091040260200160405190810160405280929190818152602001828054610704906114ec565b80156107515780601f1061072657610100808354040283529160200191610751565b820191906000526020600020905b81548152906001019060200180831161073457829003601f168201915b5050505050908060010160405180604001604052908160008201548152602001600182015481525050905082565b60046020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160405180604001604052908160008201548152602001600182015481525050908060030160405180604001604052908160008201548152602001600182015481525050908060050154905084565b60065481565b600360205280600052604060002060009150905080600001604051806040016040529081600082015481526020016001820154815250509080600201604051806040016040529081600082015481526020016001820154815250509080600401805461087c906114ec565b80601f01602080910402602001604051908101604052809291908181526020018280546108a8906114ec565b80156108f55780601f106108ca576101008083540402835291602001916108f5565b820191906000526020600020905b8154815290600101906020018083116108d857829003601f168201915b5050505050905083565b60065460075410610945576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161093c90611853565b60405180910390fd5b6001600281111561095957610958611399565b5b60008054906101000a900460ff16600281111561097957610978611399565b5b1461098357600080fd5b600060405180604001604052808781526020018681525090506000604051806040016040528086815260200185815250905060405180606001604052808381526020018281526020018481525060036000600754815260200190815260200160002060008201518160000160008201518160000155602082015181600101555050602082015181600201600082015181600001556020820151816001015550506040820151816004019081610a3891906116c9565b5090505060076000815480929190610a4f90611475565b919050555050505050505050565b60065460075414610aa3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a9a906118bf565b60405180910390fd5b60026000806101000a81548160ff02191690836002811115610ac857610ac7611399565b5b0217905550600033905060005b600854811015610b77578173ffffffffffffffffffffffffffffffffffffffff166004600083815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610b64578260046000838152602001908152602001600020600501819055505b8080610b6f90611475565b915050610ad5565b505050565b60006002811115610b9057610b8f611399565b5b60008054906101000a900460ff166002811115610bb057610baf611399565b5b14610bba57600080fd5b600033905060005b600854811015610c85578173ffffffffffffffffffffffffffffffffffffffff166004600083815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610c72578360046000838152602001908152602001600020600301600001819055508260046000838152602001908152602001600020600301600001819055505b8080610c7d90611475565b915050610bc2565b50505050565b60055481565b60008054906101000a900460ff1681565b60075481565b60026020528060005260406000206000915090508060000160405180604001604052908160008201548152602001600182015481525050905081565b600260085410610d29576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d209061192b565b60405180910390fd5b60006002811115610d3d57610d3c611399565b5b60008054906101000a900460ff166002811115610d5d57610d5c611399565b5b14610d6757600080fd5b600033905060405180608001604052808273ffffffffffffffffffffffffffffffffffffffff168152602001604051806040016040528086815260200185815250815260200160405180604001604052806000815260200160008152508152602001600081525060046000600854815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160008201518160000155602082015181600101555050604082015181600301600082015181600001556020820151816001015550506060820151816005015590505060086000815480929190610e8890611475565b9190505550505050565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b610eb981610ea6565b8114610ec457600080fd5b50565b600081359050610ed681610eb0565b92915050565b60008060408385031215610ef357610ef2610e9c565b5b6000610f0185828601610ec7565b9250506020610f1285828601610ec7565b9150509250929050565b610f2581610ea6565b82525050565b6000602082019050610f406000830184610f1c565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610f9982610f50565b810181811067ffffffffffffffff82111715610fb857610fb7610f61565b5b80604052505050565b6000610fcb610e92565b9050610fd78282610f90565b919050565b600067ffffffffffffffff821115610ff757610ff6610f61565b5b61100082610f50565b9050602081019050919050565b82818337600083830152505050565b600061102f61102a84610fdc565b610fc1565b90508281526020810184848401111561104b5761104a610f4b565b5b61105684828561100d565b509392505050565b600082601f83011261107357611072610f46565b5b813561108384826020860161101c565b91505092915050565b6000806000606084860312156110a5576110a4610e9c565b5b600084013567ffffffffffffffff8111156110c3576110c2610ea1565b5b6110cf8682870161105e565b93505060206110e086828701610ec7565b92505060406110f186828701610ec7565b9150509250925092565b60006040820190506111106000830185610f1c565b61111d6020830184610f1c565b9392505050565b60006020828403121561113a57611139610e9c565b5b600061114884828501610ec7565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561118b578082015181840152602081019050611170565b60008484015250505050565b60006111a282611151565b6111ac818561115c565b93506111bc81856020860161116d565b6111c581610f50565b840191505092915050565b6111d981610ea6565b82525050565b6040820160008201516111f560008501826111d0565b50602082015161120860208501826111d0565b50505050565b600060608201905081810360008301526112288185611197565b905061123760208301846111df565b9392505050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006112698261123e565b9050919050565b6112798161125e565b82525050565b600060c0820190506112946000830187611270565b6112a160208301866111df565b6112ae60608301856111df565b6112bb60a0830184610f1c565b95945050505050565b600060a0820190506112d960008301866111df565b6112e660408301856111df565b81810360808301526112f88184611197565b9050949350505050565b600080600080600060a0868803121561131e5761131d610e9c565b5b600061132c88828901610ec7565b955050602061133d88828901610ec7565b945050604061134e88828901610ec7565b935050606061135f88828901610ec7565b925050608086013567ffffffffffffffff8111156113805761137f610ea1565b5b61138c8882890161105e565b9150509295509295909350565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600381106113d9576113d8611399565b5b50565b60008190506113ea826113c8565b919050565b60006113fa826113dc565b9050919050565b61140a816113ef565b82525050565b60006020820190506114256000830184611401565b92915050565b600060408201905061144060008301846111df565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061148082610ea6565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036114b2576114b1611446565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061150457607f821691505b602082108103611517576115166114bd565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b60006008830261157f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611542565b6115898683611542565b95508019841693508086168417925050509392505050565b6000819050919050565b60006115c66115c16115bc84610ea6565b6115a1565b610ea6565b9050919050565b6000819050919050565b6115e0836115ab565b6115f46115ec826115cd565b84845461154f565b825550505050565b600090565b6116096115fc565b6116148184846115d7565b505050565b5b818110156116385761162d600082611601565b60018101905061161a565b5050565b601f82111561167d5761164e8161151d565b61165784611532565b81016020851015611666578190505b61167a61167285611532565b830182611619565b50505b505050565b600082821c905092915050565b60006116a060001984600802611682565b1980831691505092915050565b60006116b9838361168f565b9150826002028217905092915050565b6116d282611151565b67ffffffffffffffff8111156116eb576116ea610f61565b5b6116f582546114ec565b61170082828561163c565b600060209050601f8311600181146117335760008415611721578287015190505b61172b85826116ad565b865550611793565b601f1984166117418661151d565b60005b8281101561176957848901518255600182019150602085019450602081019050611744565b868310156117865784890151611782601f89168261168f565b8355505b6001600288020188555050505b505050505050565b7f646976657267656e7420616e6e6f756e63656d656e7400000000000000000000600082015250565b60006117d160168361115c565b91506117dc8261179b565b602082019050919050565b60006020820190508181036000830152611800816117c4565b9050919050565b7f746f6f206d616e7920766f746573000000000000000000000000000000000000600082015250565b600061183d600e8361115c565b915061184882611807565b602082019050919050565b6000602082019050818103600083015261186c81611830565b9050919050565b7f6d697373696e672062616c6c6f74730000000000000000000000000000000000600082015250565b60006118a9600f8361115c565b91506118b482611873565b602082019050919050565b600060208201905081810360008301526118d88161189c565b9050919050565b7f546f6f206d616e79206d616e6167657273000000000000000000000000000000600082015250565b600061191560118361115c565b9150611920826118df565b602082019050919050565b6000602082019050818103600083015261194481611908565b905091905056fea26469706673582212209b41c22c980c8f47dffd14102343697c95eff644912ac22139f57eac381be22d64736f6c63430008110033",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// BallotCount is a free data retrieval call binding the contract method 0xc0030032.
//
// Solidity: function ballotCount() view returns(uint256)
func (_Contract *ContractCaller) BallotCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "ballotCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BallotCount is a free data retrieval call binding the contract method 0xc0030032.
//
// Solidity: function ballotCount() view returns(uint256)
func (_Contract *ContractSession) BallotCount() (*big.Int, error) {
	return _Contract.Contract.BallotCount(&_Contract.CallOpts)
}

// BallotCount is a free data retrieval call binding the contract method 0xc0030032.
//
// Solidity: function ballotCount() view returns(uint256)
func (_Contract *ContractCallerSession) BallotCount() (*big.Int, error) {
	return _Contract.Contract.BallotCount(&_Contract.CallOpts)
}

// Ballots is a free data retrieval call binding the contract method 0x5c632b38.
//
// Solidity: function ballots(uint256 ) view returns((uint256,uint256) SA, (uint256,uint256) R, string signature)
func (_Contract *ContractCaller) Ballots(opts *bind.CallOpts, arg0 *big.Int) (struct {
	SA        Point
	R         Point
	Signature string
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "ballots", arg0)

	outstruct := new(struct {
		SA        Point
		R         Point
		Signature string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SA = *abi.ConvertType(out[0], new(Point)).(*Point)
	outstruct.R = *abi.ConvertType(out[1], new(Point)).(*Point)
	outstruct.Signature = *abi.ConvertType(out[2], new(string)).(*string)

	return *outstruct, err

}

// Ballots is a free data retrieval call binding the contract method 0x5c632b38.
//
// Solidity: function ballots(uint256 ) view returns((uint256,uint256) SA, (uint256,uint256) R, string signature)
func (_Contract *ContractSession) Ballots(arg0 *big.Int) (struct {
	SA        Point
	R         Point
	Signature string
}, error) {
	return _Contract.Contract.Ballots(&_Contract.CallOpts, arg0)
}

// Ballots is a free data retrieval call binding the contract method 0x5c632b38.
//
// Solidity: function ballots(uint256 ) view returns((uint256,uint256) SA, (uint256,uint256) R, string signature)
func (_Contract *ContractCallerSession) Ballots(arg0 *big.Int) (struct {
	SA        Point
	R         Point
	Signature string
}, error) {
	return _Contract.Contract.Ballots(&_Contract.CallOpts, arg0)
}

// CandidateCount is a free data retrieval call binding the contract method 0xa9a981a3.
//
// Solidity: function candidateCount() view returns(uint256)
func (_Contract *ContractCaller) CandidateCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "candidateCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CandidateCount is a free data retrieval call binding the contract method 0xa9a981a3.
//
// Solidity: function candidateCount() view returns(uint256)
func (_Contract *ContractSession) CandidateCount() (*big.Int, error) {
	return _Contract.Contract.CandidateCount(&_Contract.CallOpts)
}

// CandidateCount is a free data retrieval call binding the contract method 0xa9a981a3.
//
// Solidity: function candidateCount() view returns(uint256)
func (_Contract *ContractCallerSession) CandidateCount() (*big.Int, error) {
	return _Contract.Contract.CandidateCount(&_Contract.CallOpts)
}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates(uint256 ) view returns(string name, (uint256,uint256) pubkey)
func (_Contract *ContractCaller) Candidates(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name   string
	Pubkey Point
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "candidates", arg0)

	outstruct := new(struct {
		Name   string
		Pubkey Point
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Pubkey = *abi.ConvertType(out[1], new(Point)).(*Point)

	return *outstruct, err

}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates(uint256 ) view returns(string name, (uint256,uint256) pubkey)
func (_Contract *ContractSession) Candidates(arg0 *big.Int) (struct {
	Name   string
	Pubkey Point
}, error) {
	return _Contract.Contract.Candidates(&_Contract.CallOpts, arg0)
}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates(uint256 ) view returns(string name, (uint256,uint256) pubkey)
func (_Contract *ContractCallerSession) Candidates(arg0 *big.Int) (struct {
	Name   string
	Pubkey Point
}, error) {
	return _Contract.Contract.Candidates(&_Contract.CallOpts, arg0)
}

// ManagerCount is a free data retrieval call binding the contract method 0x103b7397.
//
// Solidity: function managerCount() view returns(uint256)
func (_Contract *ContractCaller) ManagerCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "managerCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ManagerCount is a free data retrieval call binding the contract method 0x103b7397.
//
// Solidity: function managerCount() view returns(uint256)
func (_Contract *ContractSession) ManagerCount() (*big.Int, error) {
	return _Contract.Contract.ManagerCount(&_Contract.CallOpts)
}

// ManagerCount is a free data retrieval call binding the contract method 0x103b7397.
//
// Solidity: function managerCount() view returns(uint256)
func (_Contract *ContractCallerSession) ManagerCount() (*big.Int, error) {
	return _Contract.Contract.ManagerCount(&_Contract.CallOpts)
}

// Managers is a free data retrieval call binding the contract method 0x3d39c260.
//
// Solidity: function managers(uint256 ) view returns(address owner, (uint256,uint256) pubkey, (uint256,uint256) announcedKey, uint256 privateKey)
func (_Contract *ContractCaller) Managers(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Owner        common.Address
	Pubkey       Point
	AnnouncedKey Point
	PrivateKey   *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "managers", arg0)

	outstruct := new(struct {
		Owner        common.Address
		Pubkey       Point
		AnnouncedKey Point
		PrivateKey   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Pubkey = *abi.ConvertType(out[1], new(Point)).(*Point)
	outstruct.AnnouncedKey = *abi.ConvertType(out[2], new(Point)).(*Point)
	outstruct.PrivateKey = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Managers is a free data retrieval call binding the contract method 0x3d39c260.
//
// Solidity: function managers(uint256 ) view returns(address owner, (uint256,uint256) pubkey, (uint256,uint256) announcedKey, uint256 privateKey)
func (_Contract *ContractSession) Managers(arg0 *big.Int) (struct {
	Owner        common.Address
	Pubkey       Point
	AnnouncedKey Point
	PrivateKey   *big.Int
}, error) {
	return _Contract.Contract.Managers(&_Contract.CallOpts, arg0)
}

// Managers is a free data retrieval call binding the contract method 0x3d39c260.
//
// Solidity: function managers(uint256 ) view returns(address owner, (uint256,uint256) pubkey, (uint256,uint256) announcedKey, uint256 privateKey)
func (_Contract *ContractCallerSession) Managers(arg0 *big.Int) (struct {
	Owner        common.Address
	Pubkey       Point
	AnnouncedKey Point
	PrivateKey   *big.Int
}, error) {
	return _Contract.Contract.Managers(&_Contract.CallOpts, arg0)
}

// Phase is a free data retrieval call binding the contract method 0xb1c9fe6e.
//
// Solidity: function phase() view returns(uint8)
func (_Contract *ContractCaller) Phase(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "phase")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Phase is a free data retrieval call binding the contract method 0xb1c9fe6e.
//
// Solidity: function phase() view returns(uint8)
func (_Contract *ContractSession) Phase() (uint8, error) {
	return _Contract.Contract.Phase(&_Contract.CallOpts)
}

// Phase is a free data retrieval call binding the contract method 0xb1c9fe6e.
//
// Solidity: function phase() view returns(uint8)
func (_Contract *ContractCallerSession) Phase() (uint8, error) {
	return _Contract.Contract.Phase(&_Contract.CallOpts)
}

// VoterCount is a free data retrieval call binding the contract method 0x42169e48.
//
// Solidity: function voterCount() view returns(uint256)
func (_Contract *ContractCaller) VoterCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "voterCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VoterCount is a free data retrieval call binding the contract method 0x42169e48.
//
// Solidity: function voterCount() view returns(uint256)
func (_Contract *ContractSession) VoterCount() (*big.Int, error) {
	return _Contract.Contract.VoterCount(&_Contract.CallOpts)
}

// VoterCount is a free data retrieval call binding the contract method 0x42169e48.
//
// Solidity: function voterCount() view returns(uint256)
func (_Contract *ContractCallerSession) VoterCount() (*big.Int, error) {
	return _Contract.Contract.VoterCount(&_Contract.CallOpts)
}

// Voters is a free data retrieval call binding the contract method 0xda58c7d9.
//
// Solidity: function voters(uint256 ) view returns((uint256,uint256) pubkey)
func (_Contract *ContractCaller) Voters(opts *bind.CallOpts, arg0 *big.Int) (Point, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "voters", arg0)

	if err != nil {
		return *new(Point), err
	}

	out0 := *abi.ConvertType(out[0], new(Point)).(*Point)

	return out0, err

}

// Voters is a free data retrieval call binding the contract method 0xda58c7d9.
//
// Solidity: function voters(uint256 ) view returns((uint256,uint256) pubkey)
func (_Contract *ContractSession) Voters(arg0 *big.Int) (Point, error) {
	return _Contract.Contract.Voters(&_Contract.CallOpts, arg0)
}

// Voters is a free data retrieval call binding the contract method 0xda58c7d9.
//
// Solidity: function voters(uint256 ) view returns((uint256,uint256) pubkey)
func (_Contract *ContractCallerSession) Voters(arg0 *big.Int) (Point, error) {
	return _Contract.Contract.Voters(&_Contract.CallOpts, arg0)
}

// AddCandidate is a paid mutator transaction binding the contract method 0x1970e446.
//
// Solidity: function addCandidate(string name, uint256 x, uint256 y) returns()
func (_Contract *ContractTransactor) AddCandidate(opts *bind.TransactOpts, name string, x *big.Int, y *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "addCandidate", name, x, y)
}

// AddCandidate is a paid mutator transaction binding the contract method 0x1970e446.
//
// Solidity: function addCandidate(string name, uint256 x, uint256 y) returns()
func (_Contract *ContractSession) AddCandidate(name string, x *big.Int, y *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.AddCandidate(&_Contract.TransactOpts, name, x, y)
}

// AddCandidate is a paid mutator transaction binding the contract method 0x1970e446.
//
// Solidity: function addCandidate(string name, uint256 x, uint256 y) returns()
func (_Contract *ContractTransactorSession) AddCandidate(name string, x *big.Int, y *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.AddCandidate(&_Contract.TransactOpts, name, x, y)
}

// AddVoter is a paid mutator transaction binding the contract method 0x0a120bb0.
//
// Solidity: function addVoter(uint256 x, uint256 y) returns()
func (_Contract *ContractTransactor) AddVoter(opts *bind.TransactOpts, x *big.Int, y *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "addVoter", x, y)
}

// AddVoter is a paid mutator transaction binding the contract method 0x0a120bb0.
//
// Solidity: function addVoter(uint256 x, uint256 y) returns()
func (_Contract *ContractSession) AddVoter(x *big.Int, y *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.AddVoter(&_Contract.TransactOpts, x, y)
}

// AddVoter is a paid mutator transaction binding the contract method 0x0a120bb0.
//
// Solidity: function addVoter(uint256 x, uint256 y) returns()
func (_Contract *ContractTransactorSession) AddVoter(x *big.Int, y *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.AddVoter(&_Contract.TransactOpts, x, y)
}

// AnnouncePublicKey is a paid mutator transaction binding the contract method 0x9ffdcc59.
//
// Solidity: function announcePublicKey(uint256 x, uint256 y) returns()
func (_Contract *ContractTransactor) AnnouncePublicKey(opts *bind.TransactOpts, x *big.Int, y *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "announcePublicKey", x, y)
}

// AnnouncePublicKey is a paid mutator transaction binding the contract method 0x9ffdcc59.
//
// Solidity: function announcePublicKey(uint256 x, uint256 y) returns()
func (_Contract *ContractSession) AnnouncePublicKey(x *big.Int, y *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.AnnouncePublicKey(&_Contract.TransactOpts, x, y)
}

// AnnouncePublicKey is a paid mutator transaction binding the contract method 0x9ffdcc59.
//
// Solidity: function announcePublicKey(uint256 x, uint256 y) returns()
func (_Contract *ContractTransactorSession) AnnouncePublicKey(x *big.Int, y *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.AnnouncePublicKey(&_Contract.TransactOpts, x, y)
}

// GetPublicKey is a paid mutator transaction binding the contract method 0x2e334452.
//
// Solidity: function getPublicKey() returns(uint256, uint256)
func (_Contract *ContractTransactor) GetPublicKey(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "getPublicKey")
}

// GetPublicKey is a paid mutator transaction binding the contract method 0x2e334452.
//
// Solidity: function getPublicKey() returns(uint256, uint256)
func (_Contract *ContractSession) GetPublicKey() (*types.Transaction, error) {
	return _Contract.Contract.GetPublicKey(&_Contract.TransactOpts)
}

// GetPublicKey is a paid mutator transaction binding the contract method 0x2e334452.
//
// Solidity: function getPublicKey() returns(uint256, uint256)
func (_Contract *ContractTransactorSession) GetPublicKey() (*types.Transaction, error) {
	return _Contract.Contract.GetPublicKey(&_Contract.TransactOpts)
}

// RegisterManager is a paid mutator transaction binding the contract method 0xe9857485.
//
// Solidity: function registerManager(uint256 x, uint256 y) payable returns()
func (_Contract *ContractTransactor) RegisterManager(opts *bind.TransactOpts, x *big.Int, y *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "registerManager", x, y)
}

// RegisterManager is a paid mutator transaction binding the contract method 0xe9857485.
//
// Solidity: function registerManager(uint256 x, uint256 y) payable returns()
func (_Contract *ContractSession) RegisterManager(x *big.Int, y *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RegisterManager(&_Contract.TransactOpts, x, y)
}

// RegisterManager is a paid mutator transaction binding the contract method 0xe9857485.
//
// Solidity: function registerManager(uint256 x, uint256 y) payable returns()
func (_Contract *ContractTransactorSession) RegisterManager(x *big.Int, y *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RegisterManager(&_Contract.TransactOpts, x, y)
}

// RevealPublicKey is a paid mutator transaction binding the contract method 0x9d578eca.
//
// Solidity: function revealPublicKey(uint256 r) returns()
func (_Contract *ContractTransactor) RevealPublicKey(opts *bind.TransactOpts, r *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "revealPublicKey", r)
}

// RevealPublicKey is a paid mutator transaction binding the contract method 0x9d578eca.
//
// Solidity: function revealPublicKey(uint256 r) returns()
func (_Contract *ContractSession) RevealPublicKey(r *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RevealPublicKey(&_Contract.TransactOpts, r)
}

// RevealPublicKey is a paid mutator transaction binding the contract method 0x9d578eca.
//
// Solidity: function revealPublicKey(uint256 r) returns()
func (_Contract *ContractTransactorSession) RevealPublicKey(r *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RevealPublicKey(&_Contract.TransactOpts, r)
}

// Vote is a paid mutator transaction binding the contract method 0x75eac54d.
//
// Solidity: function vote(uint256 SAx, uint256 SAy, uint256 Rx, uint256 Ry, string signature) returns()
func (_Contract *ContractTransactor) Vote(opts *bind.TransactOpts, SAx *big.Int, SAy *big.Int, Rx *big.Int, Ry *big.Int, signature string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "vote", SAx, SAy, Rx, Ry, signature)
}

// Vote is a paid mutator transaction binding the contract method 0x75eac54d.
//
// Solidity: function vote(uint256 SAx, uint256 SAy, uint256 Rx, uint256 Ry, string signature) returns()
func (_Contract *ContractSession) Vote(SAx *big.Int, SAy *big.Int, Rx *big.Int, Ry *big.Int, signature string) (*types.Transaction, error) {
	return _Contract.Contract.Vote(&_Contract.TransactOpts, SAx, SAy, Rx, Ry, signature)
}

// Vote is a paid mutator transaction binding the contract method 0x75eac54d.
//
// Solidity: function vote(uint256 SAx, uint256 SAy, uint256 Rx, uint256 Ry, string signature) returns()
func (_Contract *ContractTransactorSession) Vote(SAx *big.Int, SAy *big.Int, Rx *big.Int, Ry *big.Int, signature string) (*types.Transaction, error) {
	return _Contract.Contract.Vote(&_Contract.TransactOpts, SAx, SAy, Rx, Ry, signature)
}
