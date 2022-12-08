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
	Bin: "0x608060405260008060006101000a81548160ff0219169083600281111561002957610028610054565b5b0217905550600060055560006006556000600755600060085534801561004e57600080fd5b50610083565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b611aad806100926000396000f3fe6080604052600436106100f35760003560e01c806375eac54d1161008a578063b1c9fe6e11610059578063b1c9fe6e1461032f578063c00300321461035a578063da58c7d914610385578063e9857485146103c2576100f3565b806375eac54d146102895780639d578eca146102b25780639ffdcc59146102db578063a9a981a314610304576100f3565b80633477ee2e116100c65780633477ee2e146101a15780633d39c260146101df57806342169e481461021f5780635c632b381461024a576100f3565b80630a120bb0146100f8578063103b7397146101215780631970e4461461014c5780632e33445214610175575b600080fd5b34801561010457600080fd5b5061011f600480360381019061011a9190610f9c565b6103de565b005b34801561012d57600080fd5b50610136610492565b6040516101439190610feb565b60405180910390f35b34801561015857600080fd5b50610173600480360381019061016e919061114c565b610498565b005b34801561018157600080fd5b5061018a61061a565b6040516101989291906111bb565b60405180910390f35b3480156101ad57600080fd5b506101c860048036038101906101c391906111e4565b610766565b6040516101d69291906112ce565b60405180910390f35b3480156101eb57600080fd5b50610206600480360381019061020191906111e4565b610830565b604051610216949392919061133f565b60405180910390f35b34801561022b57600080fd5b506102346108bc565b6040516102419190610feb565b60405180910390f35b34801561025657600080fd5b50610271600480360381019061026c91906111e4565b6108c2565b60405161028093929190611384565b60405180910390f35b34801561029557600080fd5b506102b060048036038101906102ab91906113c2565b6109b0565b005b3480156102be57600080fd5b506102d960048036038101906102d491906111e4565b610b0e565b005b3480156102e757600080fd5b5061030260048036038101906102fd9190610f9c565b610c2d565b005b34801561031057600080fd5b50610319610d4b565b6040516103269190610feb565b60405180910390f35b34801561033b57600080fd5b50610344610d51565b60405161035191906114d0565b60405180910390f35b34801561036657600080fd5b5061036f610d62565b60405161037c9190610feb565b60405180910390f35b34801561039157600080fd5b506103ac60048036038101906103a791906111e4565b610d68565b6040516103b991906114eb565b60405180910390f35b6103dc60048036038101906103d79190610f9c565b610da4565b005b600060028111156103f2576103f1611459565b5b60008054906101000a900460ff16600281111561041257610411611459565b5b1461041c57600080fd5b6040518060200160405280604051806040016040528085815260200184815250815250600260006006548152602001908152602001600020600082015181600001600082015181600001556020820151816001015550509050506006600081548092919061048990611535565b91905055505050565b60085481565b600060028111156104ac576104ab611459565b5b60008054906101000a900460ff1660028111156104cc576104cb611459565b5b146104d657600080fd5b60005b60085481101561058657600060016000838152602001908152602001600020600101604051806040016040529081600082015481526020016001820154815250509050838160000151141580610533575082816020015114155b610572576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610569906115c9565b60405180910390fd5b50808061057e90611535565b9150506104d9565b50604051806040016040528084815260200160405180604001604052808581526020018481525081525060016000600554815260200190815260200160002060008201518160000190816105da91906117f5565b50602082015181600101600082015181600001556020820151816001015550509050506005600081548092919061061090611535565b9190505550505050565b60008060006004600080815260200190815260200160002060030160000154905060006004600080815260200190815260200160002060030160010154905060005b60085481101561072f57826004600083815260200190815260200160002060030160000154036106c1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106b890611913565b60405180910390fd5b8160046000838152602001908152602001600020600301600101540361071c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161071390611913565b60405180910390fd5b808061072790611535565b91505061065c565b5060016000806101000a81548160ff0219169083600281111561075557610754611459565b5b021790555081819350935050509091565b600160205280600052604060002060009150905080600001805461078990611618565b80601f01602080910402602001604051908101604052809291908181526020018280546107b590611618565b80156108025780601f106107d757610100808354040283529160200191610802565b820191906000526020600020905b8154815290600101906020018083116107e557829003601f168201915b5050505050908060010160405180604001604052908160008201548152602001600182015481525050905082565b60046020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160405180604001604052908160008201548152602001600182015481525050908060030160405180604001604052908160008201548152602001600182015481525050908060050154905084565b60065481565b600360205280600052604060002060009150905080600001604051806040016040529081600082015481526020016001820154815250509080600201604051806040016040529081600082015481526020016001820154815250509080600401805461092d90611618565b80601f016020809104026020016040519081016040528092919081815260200182805461095990611618565b80156109a65780601f1061097b576101008083540402835291602001916109a6565b820191906000526020600020905b81548152906001019060200180831161098957829003601f168201915b5050505050905083565b600654600754106109f6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109ed9061197f565b60405180910390fd5b60016002811115610a0a57610a09611459565b5b60008054906101000a900460ff166002811115610a2a57610a29611459565b5b14610a3457600080fd5b600060405180604001604052808781526020018681525090506000604051806040016040528086815260200185815250905060405180606001604052808381526020018281526020018481525060036000600754815260200190815260200160002060008201518160000160008201518160000155602082015181600101555050602082015181600201600082015181600001556020820151816001015550506040820151816004019081610ae991906117f5565b5090505060076000815480929190610b0090611535565b919050555050505050505050565b60065460075414610b54576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b4b906119eb565b60405180910390fd5b60026000806101000a81548160ff02191690836002811115610b7957610b78611459565b5b0217905550600033905060005b600854811015610c28578173ffffffffffffffffffffffffffffffffffffffff166004600083815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610c15578260046000838152602001908152602001600020600501819055505b8080610c2090611535565b915050610b86565b505050565b60006002811115610c4157610c40611459565b5b60008054906101000a900460ff166002811115610c6157610c60611459565b5b14610c6b57600080fd5b600260085414610c7a57600080fd5b600033905060005b600854811015610d45578173ffffffffffffffffffffffffffffffffffffffff166004600083815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610d32578360046000838152602001908152602001600020600301600001819055508260046000838152602001908152602001600020600301600001819055505b8080610d3d90611535565b915050610c82565b50505050565b60055481565b60008054906101000a900460ff1681565b60075481565b60026020528060005260406000206000915090508060000160405180604001604052908160008201548152602001600182015481525050905081565b600260085410610de9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610de090611a57565b60405180910390fd5b60006002811115610dfd57610dfc611459565b5b60008054906101000a900460ff166002811115610e1d57610e1c611459565b5b14610e2757600080fd5b600033905060405180608001604052808273ffffffffffffffffffffffffffffffffffffffff168152602001604051806040016040528086815260200185815250815260200160405180604001604052806000815260200160008152508152602001600081525060046000600854815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160008201518160000155602082015181600101555050604082015181600301600082015181600001556020820151816001015550506060820151816005015590505060086000815480929190610f4890611535565b9190505550505050565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b610f7981610f66565b8114610f8457600080fd5b50565b600081359050610f9681610f70565b92915050565b60008060408385031215610fb357610fb2610f5c565b5b6000610fc185828601610f87565b9250506020610fd285828601610f87565b9150509250929050565b610fe581610f66565b82525050565b60006020820190506110006000830184610fdc565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61105982611010565b810181811067ffffffffffffffff8211171561107857611077611021565b5b80604052505050565b600061108b610f52565b90506110978282611050565b919050565b600067ffffffffffffffff8211156110b7576110b6611021565b5b6110c082611010565b9050602081019050919050565b82818337600083830152505050565b60006110ef6110ea8461109c565b611081565b90508281526020810184848401111561110b5761110a61100b565b5b6111168482856110cd565b509392505050565b600082601f83011261113357611132611006565b5b81356111438482602086016110dc565b91505092915050565b60008060006060848603121561116557611164610f5c565b5b600084013567ffffffffffffffff81111561118357611182610f61565b5b61118f8682870161111e565b93505060206111a086828701610f87565b92505060406111b186828701610f87565b9150509250925092565b60006040820190506111d06000830185610fdc565b6111dd6020830184610fdc565b9392505050565b6000602082840312156111fa576111f9610f5c565b5b600061120884828501610f87565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561124b578082015181840152602081019050611230565b60008484015250505050565b600061126282611211565b61126c818561121c565b935061127c81856020860161122d565b61128581611010565b840191505092915050565b61129981610f66565b82525050565b6040820160008201516112b56000850182611290565b5060208201516112c86020850182611290565b50505050565b600060608201905081810360008301526112e88185611257565b90506112f7602083018461129f565b9392505050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000611329826112fe565b9050919050565b6113398161131e565b82525050565b600060c0820190506113546000830187611330565b611361602083018661129f565b61136e606083018561129f565b61137b60a0830184610fdc565b95945050505050565b600060a082019050611399600083018661129f565b6113a6604083018561129f565b81810360808301526113b88184611257565b9050949350505050565b600080600080600060a086880312156113de576113dd610f5c565b5b60006113ec88828901610f87565b95505060206113fd88828901610f87565b945050604061140e88828901610f87565b935050606061141f88828901610f87565b925050608086013567ffffffffffffffff8111156114405761143f610f61565b5b61144c8882890161111e565b9150509295509295909350565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6003811061149957611498611459565b5b50565b60008190506114aa82611488565b919050565b60006114ba8261149c565b9050919050565b6114ca816114af565b82525050565b60006020820190506114e560008301846114c1565b92915050565b6000604082019050611500600083018461129f565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061154082610f66565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361157257611571611506565b5b600182019050919050565b7f6475706c69636174652063616e64696461746500000000000000000000000000600082015250565b60006115b360138361121c565b91506115be8261157d565b602082019050919050565b600060208201905081810360008301526115e2816115a6565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061163057607f821691505b602082108103611643576116426115e9565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026116ab7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261166e565b6116b5868361166e565b95508019841693508086168417925050509392505050565b6000819050919050565b60006116f26116ed6116e884610f66565b6116cd565b610f66565b9050919050565b6000819050919050565b61170c836116d7565b611720611718826116f9565b84845461167b565b825550505050565b600090565b611735611728565b611740818484611703565b505050565b5b818110156117645761175960008261172d565b600181019050611746565b5050565b601f8211156117a95761177a81611649565b6117838461165e565b81016020851015611792578190505b6117a661179e8561165e565b830182611745565b50505b505050565b600082821c905092915050565b60006117cc600019846008026117ae565b1980831691505092915050565b60006117e583836117bb565b9150826002028217905092915050565b6117fe82611211565b67ffffffffffffffff81111561181757611816611021565b5b6118218254611618565b61182c828285611768565b600060209050601f83116001811461185f576000841561184d578287015190505b61185785826117d9565b8655506118bf565b601f19841661186d86611649565b60005b8281101561189557848901518255600182019150602085019450602081019050611870565b868310156118b257848901516118ae601f8916826117bb565b8355505b6001600288020188555050505b505050505050565b7f646976657267656e7420616e6e6f756e63656d656e7400000000000000000000600082015250565b60006118fd60168361121c565b9150611908826118c7565b602082019050919050565b6000602082019050818103600083015261192c816118f0565b9050919050565b7f746f6f206d616e7920766f746573000000000000000000000000000000000000600082015250565b6000611969600e8361121c565b915061197482611933565b602082019050919050565b600060208201905081810360008301526119988161195c565b9050919050565b7f6d697373696e672062616c6c6f74730000000000000000000000000000000000600082015250565b60006119d5600f8361121c565b91506119e08261199f565b602082019050919050565b60006020820190508181036000830152611a04816119c8565b9050919050565b7f546f6f206d616e79206d616e6167657273000000000000000000000000000000600082015250565b6000611a4160118361121c565b9150611a4c82611a0b565b602082019050919050565b60006020820190508181036000830152611a7081611a34565b905091905056fea2646970667358221220d127decd5a0bde860e08d8c04bd44ca0fde07ac63f7d60faf539e211b696b15764736f6c63430008110033",
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
