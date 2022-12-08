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
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"name\":\"addCandidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"Px\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Py\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Ix\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Iy\",\"type\":\"uint256\"}],\"name\":\"addVoter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"name\":\"announcePublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ballotCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ballots\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structPoint\",\"name\":\"SA\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structPoint\",\"name\":\"R\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"signature\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"candidateCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"candidates\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structPoint\",\"name\":\"pubkey\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getKey\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"managerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"managers\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structPoint\",\"name\":\"pubkey\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structPoint\",\"name\":\"announcedKey\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"privateKey\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"phase\",\"outputs\":[{\"internalType\":\"enumPhase\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"name\":\"registerManager\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"}],\"name\":\"revealPublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"SAx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"SAy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Rx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Ry\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"signature\",\"type\":\"string\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voterCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"voters\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structPoint\",\"name\":\"P\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structPoint\",\"name\":\"I\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405260008060006101000a81548160ff0219169083600281111561002957610028610054565b5b0217905550600060055560006006556000600755600060085534801561004e57600080fd5b50610083565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b611da0806100926000396000f3fe6080604052600436106100f35760003560e01c8063865c26701161008a578063b1c9fe6e11610059578063b1c9fe6e1461032f578063c00300321461035a578063da58c7d914610385578063e9857485146103c3576100f3565b8063865c2670146102895780639d578eca146102b25780639ffdcc59146102db578063a9a981a314610304576100f3565b806342169e48116100c657806342169e48146101ca5780635c632b38146101f557806375eac54d1461023457806382678dd61461025d576100f3565b8063103b7397146100f85780631970e446146101235780633477ee2e1461014c5780633d39c2601461018a575b600080fd5b34801561010457600080fd5b5061010d6103df565b60405161011a9190611111565b60405180910390f35b34801561012f57600080fd5b5061014a600480360381019061014591906112b2565b6103e5565b005b34801561015857600080fd5b50610173600480360381019061016e9190611321565b610567565b60405161018192919061140b565b60405180910390f35b34801561019657600080fd5b506101b160048036038101906101ac9190611321565b610631565b6040516101c1949392919061147c565b60405180910390f35b3480156101d657600080fd5b506101df6106bd565b6040516101ec9190611111565b60405180910390f35b34801561020157600080fd5b5061021c60048036038101906102179190611321565b6106c3565b60405161022b939291906114c1565b60405180910390f35b34801561024057600080fd5b5061025b600480360381019061025691906114ff565b6107b1565b005b34801561026957600080fd5b50610272610945565b604051610280929190611596565b60405180910390f35b34801561029557600080fd5b506102b060048036038101906102ab91906115bf565b610a91565b005b3480156102be57600080fd5b506102d960048036038101906102d49190611321565b610c31565b005b3480156102e757600080fd5b5061030260048036038101906102fd9190611626565b610dc3565b005b34801561031057600080fd5b50610319610ecd565b6040516103269190611111565b60405180910390f35b34801561033b57600080fd5b50610344610ed3565b60405161035191906116dd565b60405180910390f35b34801561036657600080fd5b5061036f610ee4565b60405161037c9190611111565b60405180910390f35b34801561039157600080fd5b506103ac60048036038101906103a79190611321565b610eea565b6040516103ba9291906116f8565b60405180910390f35b6103dd60048036038101906103d89190611626565b610f4a565b005b60085481565b600060028111156103f9576103f8611666565b5b60008054906101000a900460ff16600281111561041957610418611666565b5b1461042357600080fd5b60005b6005548110156104d357600060016000838152602001908152602001600020600101604051806040016040529081600082015481526020016001820154815250509050838160000151141580610480575082816020015114155b6104bf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104b69061176d565b60405180910390fd5b5080806104cb906117bc565b915050610426565b50604051806040016040528084815260200160405180604001604052808581526020018481525081525060016000600554815260200190815260200160002060008201518160000190816105279190611a10565b50602082015181600101600082015181600001556020820151816001015550509050506005600081548092919061055d906117bc565b9190505550505050565b600160205280600052604060002060009150905080600001805461058a90611833565b80601f01602080910402602001604051908101604052809291908181526020018280546105b690611833565b80156106035780601f106105d857610100808354040283529160200191610603565b820191906000526020600020905b8154815290600101906020018083116105e657829003601f168201915b5050505050908060010160405180604001604052908160008201548152602001600182015481525050905082565b60046020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160405180604001604052908160008201548152602001600182015481525050908060030160405180604001604052908160008201548152602001600182015481525050908060050154905084565b60065481565b600360205280600052604060002060009150905080600001604051806040016040529081600082015481526020016001820154815250509080600201604051806040016040529081600082015481526020016001820154815250509080600401805461072e90611833565b80601f016020809104026020016040519081016040528092919081815260200182805461075a90611833565b80156107a75780601f1061077c576101008083540402835291602001916107a7565b820191906000526020600020905b81548152906001019060200180831161078a57829003601f168201915b5050505050905083565b600654600754106107f7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107ee90611b2e565b60405180910390fd5b6001600281111561080b5761080a611666565b5b60008054906101000a900460ff16600281111561082b5761082a611666565b5b1461083557600080fd5b6000604051806040016040528087815260200186815250905060006040518060400160405280868152602001858152509050604051806060016040528083815260200182815260200184815250600360006007548152602001908152602001600020600082015181600001600082015181600001556020820151816001015550506020820151816002016000820151816000015560208201518160010155505060408201518160040190816108ea9190611a10565b5090505060076000815480929190610901906117bc565b91905055506006546007540361093c5760026000806101000a81548160ff0219169083600281111561093657610935611666565b5b02179055505b50505050505050565b60008060006004600080815260200190815260200160002060030160000154905060006004600080815260200190815260200160002060030160010154905060005b600854811015610a5a57826004600083815260200190815260200160002060030160000154146109ec576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109e390611b9a565b60405180910390fd5b81600460008381526020019081526020016000206003016001015414610a47576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a3e90611b9a565b60405180910390fd5b8080610a52906117bc565b915050610987565b5060016000806101000a81548160ff02191690836002811115610a8057610a7f611666565b5b021790555081819350935050509091565b60006002811115610aa557610aa4611666565b5b60008054906101000a900460ff166002811115610ac557610ac4611666565b5b14610acf57600080fd5b60005b600654811015610b7f57600060026000838152602001908152602001600020600001604051806040016040529081600082015481526020016001820154815250509050858160000151141580610b2c575084816020015114155b610b6b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b6290611c06565b60405180910390fd5b508080610b77906117bc565b915050610ad2565b5060405180604001604052806040518060400160405280878152602001868152508152602001604051806040016040528085815260200184815250815250600260006006548152602001908152602001600020600082015181600001600082015181600001556020820151816001015550506020820151816002016000820151816000015560208201518160010155505090505060066000815480929190610c26906117bc565b919050555050505050565b60065460075414610c77576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c6e90611c72565b60405180910390fd5b600280811115610c8a57610c89611666565b5b60008054906101000a900460ff166002811115610caa57610ca9611666565b5b14610cea576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ce190611cde565b60405180910390fd5b60026000806101000a81548160ff02191690836002811115610d0f57610d0e611666565b5b0217905550600033905060005b600854811015610dbe578173ffffffffffffffffffffffffffffffffffffffff166004600083815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610dab578260046000838152602001908152602001600020600501819055505b8080610db6906117bc565b915050610d1c565b505050565b600260085414610dd257600080fd5b600033905060005b600854811015610e9d578173ffffffffffffffffffffffffffffffffffffffff166004600083815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610e8a578360046000838152602001908152602001600020600301600001819055508260046000838152602001908152602001600020600301600101819055505b8080610e95906117bc565b915050610dda565b5060016000806101000a81548160ff02191690836002811115610ec357610ec2611666565b5b0217905550505050565b60055481565b60008054906101000a900460ff1681565b60075481565b60026020528060005260406000206000915090508060000160405180604001604052908160008201548152602001600182015481525050908060020160405180604001604052908160008201548152602001600182015481525050905082565b600260085410610f8f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f8690611d4a565b60405180910390fd5b60006002811115610fa357610fa2611666565b5b60008054906101000a900460ff166002811115610fc357610fc2611666565b5b14610fcd57600080fd5b600033905060405180608001604052808273ffffffffffffffffffffffffffffffffffffffff168152602001604051806040016040528086815260200185815250815260200160405180604001604052806000815260200160008152508152602001600081525060046000600854815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101600082015181600001556020820151816001015550506040820151816003016000820151816000015560208201518160010155505060608201518160050155905050600860008154809291906110ee906117bc565b9190505550505050565b6000819050919050565b61110b816110f8565b82525050565b60006020820190506111266000830184611102565b92915050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6111938261114a565b810181811067ffffffffffffffff821117156111b2576111b161115b565b5b80604052505050565b60006111c561112c565b90506111d1828261118a565b919050565b600067ffffffffffffffff8211156111f1576111f061115b565b5b6111fa8261114a565b9050602081019050919050565b82818337600083830152505050565b6000611229611224846111d6565b6111bb565b90508281526020810184848401111561124557611244611145565b5b611250848285611207565b509392505050565b600082601f83011261126d5761126c611140565b5b813561127d848260208601611216565b91505092915050565b61128f816110f8565b811461129a57600080fd5b50565b6000813590506112ac81611286565b92915050565b6000806000606084860312156112cb576112ca611136565b5b600084013567ffffffffffffffff8111156112e9576112e861113b565b5b6112f586828701611258565b93505060206113068682870161129d565b92505060406113178682870161129d565b9150509250925092565b60006020828403121561133757611336611136565b5b60006113458482850161129d565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561138857808201518184015260208101905061136d565b60008484015250505050565b600061139f8261134e565b6113a98185611359565b93506113b981856020860161136a565b6113c28161114a565b840191505092915050565b6113d6816110f8565b82525050565b6040820160008201516113f260008501826113cd565b50602082015161140560208501826113cd565b50505050565b600060608201905081810360008301526114258185611394565b905061143460208301846113dc565b9392505050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006114668261143b565b9050919050565b6114768161145b565b82525050565b600060c082019050611491600083018761146d565b61149e60208301866113dc565b6114ab60608301856113dc565b6114b860a0830184611102565b95945050505050565b600060a0820190506114d660008301866113dc565b6114e360408301856113dc565b81810360808301526114f58184611394565b9050949350505050565b600080600080600060a0868803121561151b5761151a611136565b5b60006115298882890161129d565b955050602061153a8882890161129d565b945050604061154b8882890161129d565b935050606061155c8882890161129d565b925050608086013567ffffffffffffffff81111561157d5761157c61113b565b5b61158988828901611258565b9150509295509295909350565b60006040820190506115ab6000830185611102565b6115b86020830184611102565b9392505050565b600080600080608085870312156115d9576115d8611136565b5b60006115e78782880161129d565b94505060206115f88782880161129d565b93505060406116098782880161129d565b925050606061161a8782880161129d565b91505092959194509250565b6000806040838503121561163d5761163c611136565b5b600061164b8582860161129d565b925050602061165c8582860161129d565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600381106116a6576116a5611666565b5b50565b60008190506116b782611695565b919050565b60006116c7826116a9565b9050919050565b6116d7816116bc565b82525050565b60006020820190506116f260008301846116ce565b92915050565b600060808201905061170d60008301856113dc565b61171a60408301846113dc565b9392505050565b7f6475706c69636174652063616e64696461746500000000000000000000000000600082015250565b6000611757601383611359565b915061176282611721565b602082019050919050565b600060208201905081810360008301526117868161174a565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006117c7826110f8565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036117f9576117f861178d565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061184b57607f821691505b60208210810361185e5761185d611804565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026118c67fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611889565b6118d08683611889565b95508019841693508086168417925050509392505050565b6000819050919050565b600061190d611908611903846110f8565b6118e8565b6110f8565b9050919050565b6000819050919050565b611927836118f2565b61193b61193382611914565b848454611896565b825550505050565b600090565b611950611943565b61195b81848461191e565b505050565b5b8181101561197f57611974600082611948565b600181019050611961565b5050565b601f8211156119c45761199581611864565b61199e84611879565b810160208510156119ad578190505b6119c16119b985611879565b830182611960565b50505b505050565b600082821c905092915050565b60006119e7600019846008026119c9565b1980831691505092915050565b6000611a0083836119d6565b9150826002028217905092915050565b611a198261134e565b67ffffffffffffffff811115611a3257611a3161115b565b5b611a3c8254611833565b611a47828285611983565b600060209050601f831160018114611a7a5760008415611a68578287015190505b611a7285826119f4565b865550611ada565b601f198416611a8886611864565b60005b82811015611ab057848901518255600182019150602085019450602081019050611a8b565b86831015611acd5784890151611ac9601f8916826119d6565b8355505b6001600288020188555050505b505050505050565b7f746f6f206d616e7920766f746573000000000000000000000000000000000000600082015250565b6000611b18600e83611359565b9150611b2382611ae2565b602082019050919050565b60006020820190508181036000830152611b4781611b0b565b9050919050565b7f646976657267656e7420616e6e6f756e63656d656e7400000000000000000000600082015250565b6000611b84601683611359565b9150611b8f82611b4e565b602082019050919050565b60006020820190508181036000830152611bb381611b77565b9050919050565b7f6475706c696361746520766f7465720000000000000000000000000000000000600082015250565b6000611bf0600f83611359565b9150611bfb82611bba565b602082019050919050565b60006020820190508181036000830152611c1f81611be3565b9050919050565b7f6d697373696e672062616c6c6f74730000000000000000000000000000000000600082015250565b6000611c5c600f83611359565b9150611c6782611c26565b602082019050919050565b60006020820190508181036000830152611c8b81611c4f565b9050919050565b7f6e6f742074616c6c792070686173650000000000000000000000000000000000600082015250565b6000611cc8600f83611359565b9150611cd382611c92565b602082019050919050565b60006020820190508181036000830152611cf781611cbb565b9050919050565b7f546f6f206d616e79206d616e6167657273000000000000000000000000000000600082015250565b6000611d34601183611359565b9150611d3f82611cfe565b602082019050919050565b60006020820190508181036000830152611d6381611d27565b905091905056fea264697066735822122018028ed601f3f9f297aee691567507285688b2cb32b3d07e52a819e44bbd699f64736f6c63430008110033",
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
// Solidity: function voters(uint256 ) view returns((uint256,uint256) P, (uint256,uint256) I)
func (_Contract *ContractCaller) Voters(opts *bind.CallOpts, arg0 *big.Int) (struct {
	P Point
	I Point
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "voters", arg0)

	outstruct := new(struct {
		P Point
		I Point
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.P = *abi.ConvertType(out[0], new(Point)).(*Point)
	outstruct.I = *abi.ConvertType(out[1], new(Point)).(*Point)

	return *outstruct, err

}

// Voters is a free data retrieval call binding the contract method 0xda58c7d9.
//
// Solidity: function voters(uint256 ) view returns((uint256,uint256) P, (uint256,uint256) I)
func (_Contract *ContractSession) Voters(arg0 *big.Int) (struct {
	P Point
	I Point
}, error) {
	return _Contract.Contract.Voters(&_Contract.CallOpts, arg0)
}

// Voters is a free data retrieval call binding the contract method 0xda58c7d9.
//
// Solidity: function voters(uint256 ) view returns((uint256,uint256) P, (uint256,uint256) I)
func (_Contract *ContractCallerSession) Voters(arg0 *big.Int) (struct {
	P Point
	I Point
}, error) {
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

// AddVoter is a paid mutator transaction binding the contract method 0x865c2670.
//
// Solidity: function addVoter(uint256 Px, uint256 Py, uint256 Ix, uint256 Iy) returns()
func (_Contract *ContractTransactor) AddVoter(opts *bind.TransactOpts, Px *big.Int, Py *big.Int, Ix *big.Int, Iy *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "addVoter", Px, Py, Ix, Iy)
}

// AddVoter is a paid mutator transaction binding the contract method 0x865c2670.
//
// Solidity: function addVoter(uint256 Px, uint256 Py, uint256 Ix, uint256 Iy) returns()
func (_Contract *ContractSession) AddVoter(Px *big.Int, Py *big.Int, Ix *big.Int, Iy *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.AddVoter(&_Contract.TransactOpts, Px, Py, Ix, Iy)
}

// AddVoter is a paid mutator transaction binding the contract method 0x865c2670.
//
// Solidity: function addVoter(uint256 Px, uint256 Py, uint256 Ix, uint256 Iy) returns()
func (_Contract *ContractTransactorSession) AddVoter(Px *big.Int, Py *big.Int, Ix *big.Int, Iy *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.AddVoter(&_Contract.TransactOpts, Px, Py, Ix, Iy)
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

// GetKey is a paid mutator transaction binding the contract method 0x82678dd6.
//
// Solidity: function getKey() returns(uint256, uint256)
func (_Contract *ContractTransactor) GetKey(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "getKey")
}

// GetKey is a paid mutator transaction binding the contract method 0x82678dd6.
//
// Solidity: function getKey() returns(uint256, uint256)
func (_Contract *ContractSession) GetKey() (*types.Transaction, error) {
	return _Contract.Contract.GetKey(&_Contract.TransactOpts)
}

// GetKey is a paid mutator transaction binding the contract method 0x82678dd6.
//
// Solidity: function getKey() returns(uint256, uint256)
func (_Contract *ContractTransactorSession) GetKey() (*types.Transaction, error) {
	return _Contract.Contract.GetKey(&_Contract.TransactOpts)
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
