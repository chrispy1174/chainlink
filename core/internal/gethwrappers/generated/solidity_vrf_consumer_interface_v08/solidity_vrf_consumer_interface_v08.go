// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package solidity_vrf_consumer_interface_v08

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

const VRFConsumerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vrfCoordinator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_link\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"randomnessOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"randomness\",\"type\":\"uint256\"}],\"name\":\"rawFulfillRandomness\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_keyHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_seed\",\"type\":\"uint256\"}],\"name\":\"testRequestRandomness\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

var VRFConsumerBin = "0x60c060405234801561001057600080fd5b5060405161051a38038061051a83398101604081905261002f91610069565b6001600160601b0319606092831b811660a052911b1660805261009b565b80516001600160a01b038116811461006457600080fd5b919050565b6000806040838503121561007b578182fd5b6100848361004d565b91506100926020840161004d565b90509250929050565b60805160601c60a05160601c61044e6100cc6000396000818160c901526101a501526000610169015261044e6000f3fe608060405234801561001057600080fd5b506004361061004b5760003560e01c80626d6cae146100505780632f47fd861461006b5780638a5f002b1461007457806394985ddd14610087575b600080fd5b61005960025481565b60405190815260200160405180910390f35b61005960015481565b610059610082366004610343565b61009c565b61009a610095366004610322565b6100b1565b005b60006100a9848484610165565b949350505050565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610154576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f4f6e6c7920565246436f6f7264696e61746f722063616e2066756c66696c6c00604482015260640160405180910390fd5b6101618282600155600255565b5050565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16634000aea07f00000000000000000000000000000000000000000000000000000000000000008587866040516020016101e1929190918252602082015260400190565b6040516020818303038152906040526040518463ffffffff1660e01b815260040161020e9392919061036e565b602060405180830381600087803b15801561022857600080fd5b505af115801561023c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061026091906102fb565b5060008481526020818152604080832054815180840189905280830187905230606082015260808082018390528351808303909101815260a0909101909252815191830191909120878452929091526102ba906001610404565b600086815260208181526040918290209290925580518083018890528082018490528151808203830181526060909101909152805191012095945050505050565b60006020828403121561030c578081fd5b8151801515811461031b578182fd5b9392505050565b60008060408385031215610334578081fd5b50508035926020909101359150565b600080600060608486031215610357578081fd5b505081359360208301359350604090920135919050565b73ffffffffffffffffffffffffffffffffffffffff84168152600060208481840152606060408401528351806060850152825b818110156103bd578581018301518582016080015282016103a1565b818111156103ce5783608083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160800195945050505050565b6000821982111561043c577f4e487b710000000000000000000000000000000000000000000000000000000081526011600452602481fd5b50019056fea164736f6c6343000804000a"

func DeployVRFConsumer(auth *bind.TransactOpts, backend bind.ContractBackend, _vrfCoordinator common.Address, _link common.Address) (common.Address, *types.Transaction, *VRFConsumer, error) {
	parsed, err := abi.JSON(strings.NewReader(VRFConsumerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VRFConsumerBin), backend, _vrfCoordinator, _link)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VRFConsumer{VRFConsumerCaller: VRFConsumerCaller{contract: contract}, VRFConsumerTransactor: VRFConsumerTransactor{contract: contract}, VRFConsumerFilterer: VRFConsumerFilterer{contract: contract}}, nil
}

type VRFConsumer struct {
	address common.Address
	abi     abi.ABI
	VRFConsumerCaller
	VRFConsumerTransactor
	VRFConsumerFilterer
}

type VRFConsumerCaller struct {
	contract *bind.BoundContract
}

type VRFConsumerTransactor struct {
	contract *bind.BoundContract
}

type VRFConsumerFilterer struct {
	contract *bind.BoundContract
}

type VRFConsumerSession struct {
	Contract     *VRFConsumer
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type VRFConsumerCallerSession struct {
	Contract *VRFConsumerCaller
	CallOpts bind.CallOpts
}

type VRFConsumerTransactorSession struct {
	Contract     *VRFConsumerTransactor
	TransactOpts bind.TransactOpts
}

type VRFConsumerRaw struct {
	Contract *VRFConsumer
}

type VRFConsumerCallerRaw struct {
	Contract *VRFConsumerCaller
}

type VRFConsumerTransactorRaw struct {
	Contract *VRFConsumerTransactor
}

func NewVRFConsumer(address common.Address, backend bind.ContractBackend) (*VRFConsumer, error) {
	abi, err := abi.JSON(strings.NewReader(VRFConsumerABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindVRFConsumer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRFConsumer{address: address, abi: abi, VRFConsumerCaller: VRFConsumerCaller{contract: contract}, VRFConsumerTransactor: VRFConsumerTransactor{contract: contract}, VRFConsumerFilterer: VRFConsumerFilterer{contract: contract}}, nil
}

func NewVRFConsumerCaller(address common.Address, caller bind.ContractCaller) (*VRFConsumerCaller, error) {
	contract, err := bindVRFConsumer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFConsumerCaller{contract: contract}, nil
}

func NewVRFConsumerTransactor(address common.Address, transactor bind.ContractTransactor) (*VRFConsumerTransactor, error) {
	contract, err := bindVRFConsumer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFConsumerTransactor{contract: contract}, nil
}

func NewVRFConsumerFilterer(address common.Address, filterer bind.ContractFilterer) (*VRFConsumerFilterer, error) {
	contract, err := bindVRFConsumer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFConsumerFilterer{contract: contract}, nil
}

func bindVRFConsumer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VRFConsumerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_VRFConsumer *VRFConsumerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFConsumer.Contract.VRFConsumerCaller.contract.Call(opts, result, method, params...)
}

func (_VRFConsumer *VRFConsumerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFConsumer.Contract.VRFConsumerTransactor.contract.Transfer(opts)
}

func (_VRFConsumer *VRFConsumerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFConsumer.Contract.VRFConsumerTransactor.contract.Transact(opts, method, params...)
}

func (_VRFConsumer *VRFConsumerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFConsumer.Contract.contract.Call(opts, result, method, params...)
}

func (_VRFConsumer *VRFConsumerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFConsumer.Contract.contract.Transfer(opts)
}

func (_VRFConsumer *VRFConsumerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFConsumer.Contract.contract.Transact(opts, method, params...)
}

func (_VRFConsumer *VRFConsumerCaller) RandomnessOutput(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFConsumer.contract.Call(opts, &out, "randomnessOutput")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFConsumer *VRFConsumerSession) RandomnessOutput() (*big.Int, error) {
	return _VRFConsumer.Contract.RandomnessOutput(&_VRFConsumer.CallOpts)
}

func (_VRFConsumer *VRFConsumerCallerSession) RandomnessOutput() (*big.Int, error) {
	return _VRFConsumer.Contract.RandomnessOutput(&_VRFConsumer.CallOpts)
}

func (_VRFConsumer *VRFConsumerCaller) RequestId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VRFConsumer.contract.Call(opts, &out, "requestId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_VRFConsumer *VRFConsumerSession) RequestId() ([32]byte, error) {
	return _VRFConsumer.Contract.RequestId(&_VRFConsumer.CallOpts)
}

func (_VRFConsumer *VRFConsumerCallerSession) RequestId() ([32]byte, error) {
	return _VRFConsumer.Contract.RequestId(&_VRFConsumer.CallOpts)
}

func (_VRFConsumer *VRFConsumerTransactor) RawFulfillRandomness(opts *bind.TransactOpts, requestId [32]byte, randomness *big.Int) (*types.Transaction, error) {
	return _VRFConsumer.contract.Transact(opts, "rawFulfillRandomness", requestId, randomness)
}

func (_VRFConsumer *VRFConsumerSession) RawFulfillRandomness(requestId [32]byte, randomness *big.Int) (*types.Transaction, error) {
	return _VRFConsumer.Contract.RawFulfillRandomness(&_VRFConsumer.TransactOpts, requestId, randomness)
}

func (_VRFConsumer *VRFConsumerTransactorSession) RawFulfillRandomness(requestId [32]byte, randomness *big.Int) (*types.Transaction, error) {
	return _VRFConsumer.Contract.RawFulfillRandomness(&_VRFConsumer.TransactOpts, requestId, randomness)
}

func (_VRFConsumer *VRFConsumerTransactor) TestRequestRandomness(opts *bind.TransactOpts, _keyHash [32]byte, _fee *big.Int, _seed *big.Int) (*types.Transaction, error) {
	return _VRFConsumer.contract.Transact(opts, "testRequestRandomness", _keyHash, _fee, _seed)
}

func (_VRFConsumer *VRFConsumerSession) TestRequestRandomness(_keyHash [32]byte, _fee *big.Int, _seed *big.Int) (*types.Transaction, error) {
	return _VRFConsumer.Contract.TestRequestRandomness(&_VRFConsumer.TransactOpts, _keyHash, _fee, _seed)
}

func (_VRFConsumer *VRFConsumerTransactorSession) TestRequestRandomness(_keyHash [32]byte, _fee *big.Int, _seed *big.Int) (*types.Transaction, error) {
	return _VRFConsumer.Contract.TestRequestRandomness(&_VRFConsumer.TransactOpts, _keyHash, _fee, _seed)
}

func (_VRFConsumer *VRFConsumer) Address() common.Address {
	return _VRFConsumer.address
}

type VRFConsumerInterface interface {
	RandomnessOutput(opts *bind.CallOpts) (*big.Int, error)

	RequestId(opts *bind.CallOpts) ([32]byte, error)

	RawFulfillRandomness(opts *bind.TransactOpts, requestId [32]byte, randomness *big.Int) (*types.Transaction, error)

	TestRequestRandomness(opts *bind.TransactOpts, _keyHash [32]byte, _fee *big.Int, _seed *big.Int) (*types.Transaction, error)

	Address() common.Address
}
