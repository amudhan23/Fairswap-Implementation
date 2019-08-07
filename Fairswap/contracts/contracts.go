// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// FairswapABI is the input ABI used to generate the binding from.
const FairswapABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"EncInputRoot\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"receiverAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"senderGetEther_ReceiverNoComplainMessage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"accept\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"key\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Now\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"senderAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"complaint\",\"type\":\"bytes32[][]\"},{\"name\":\"gateMerkleProof\",\"type\":\"bytes32[]\"},{\"name\":\"encodedVectors_ThisGate\",\"type\":\"bytes[]\"}],\"name\":\"complain\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"depth\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"timeout\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"gotTheCorrectFiles\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"validComplain\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EncCircuitRoot\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"Nowy2\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxInputLinesToAGate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"phase\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"receiverGetEther_SenderNotRevealedKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"source\",\"type\":\"bytes\"}],\"name\":\"bytesToBytes32\",\"outputs\":[{\"name\":\"result\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"commitmentOfKey\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"revealKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalNumberOfGates\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"nextStage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"Nowy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_commitmentOfKey\",\"type\":\"bytes32\"},{\"name\":\"_EncInputRoot\",\"type\":\"bytes32\"},{\"name\":\"_EncCircuitRoot\",\"type\":\"bytes32\"},{\"name\":\"_totalNumberOfGates\",\"type\":\"uint256\"},{\"name\":\"_maxInputLinesToAGate\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// FairswapBin is the compiled bytecode used for deploying new contracts.
const FairswapBin = `0x6080604052600c805460ff191690553480156200001b57600080fd5b5060405160c080620020b783398101806040526200003d91908101906200030d565b6009805461010060a860020a03191633610100021790556002869055600385905560008490556001839055600782905560088190556200008682640100000000620000af810204565b60010160065542600555620000a3640100000000620002b2810204565b505050505050620003a4565b604080517ff8f9cbfae6cc78fbefe7cdc3a1793dfcf4f0e8bbd8cec470b6a28a7a5a3e1efd81527ff5ecf1b3e9debc68e1d9cfabc5997135bfb7a7a3938b7b606b5b4b3f2f1f0ffe60208201527ff6e4ed9ff2d6b458eadcdf97bd91692de2d4da8fd2d0ac50c6ae9a8272523616818301527fc8c0b887b0a8a4489c948c7f847c6125746c645c544c444038302820181008ff60608201527ff7cae577eec2a03cf3bad76fb589591debb2dd67e0aa9834bea6925f6a4a2e0e60808201527fe39ed557db96902cd38ed14fad815115c786af479b7e8324736353433727170760a08201527fc976c13bb96e881cb166a933a55e490d9d56952b8d4e801485467d236242260660c08201527f753a6d1b65325d0c552a4d1345224105391a310b29122104190a11030902010060e08201526101008082019092527f0100000000000000000000000000000000000000000000000000000000000000600170010000000000000000000000000000000068010000000000000000640100000000620100006010600460026000198c019081041790810417908104178881041790810417908104179081041790810417017e818283848586878898a8b8c8d8e8f929395969799a9b9d9e9faaeb6bedeeff0281900460ff0390910151047f8000000000000000000000000000000000000000000000000000000000000000909211020190565b600c5460ff166004811115620002c457fe5b6001016004811115620002d357fe5b600c805460ff19166001836004811115620002ea57fe5b021790555060144201600b55565b6000620003068251620003a1565b9392505050565b60008060008060008060c087890312156200032757600080fd5b6000620003358989620002f8565b96505060206200034889828a01620002f8565b95505060406200035b89828a01620002f8565b94505060606200036e89828a01620002f8565b93505060806200038189828a01620002f8565b92505060a06200039489828a01620002f8565b9150509295509295509295565b90565b611d0380620003b46000396000f3fe608060405260043610610168576000357c0100000000000000000000000000000000000000000000000000000000900480637a514594116100d3578063bfe370d91161008c578063bfe370d914610324578063ca641f8614610344578063d654723614610359578063e44b23c814610379578063ee3743ab1461038e578063fd6fd6fa146103a357610168565b80637a51459414610299578063a035b1fe146102ae578063a2b26c06146102c3578063aa9b8ac7146102d8578063b1c9fe6e146102ed578063b97773911461030f57610168565b80634fc84791116101255780634fc84791146102035780635c57b30814610218578063631c56ef1461024557806370dea79a1461025a57806375986d641461026f57806376aa657f1461028457610168565b8063062a13981461016d57806316fed3e214610198578063264c57e1146101ba5780632852b71c146101d15780633943380c146101d957806344d4fd19146101ee575b600080fd5b34801561017957600080fd5b506101826103b8565b60405161018f9190611bf7565b60405180910390f35b3480156101a457600080fd5b506101ad6103be565b60405161018f9190611bdb565b3480156101c657600080fd5b506101cf6103cd565b005b6101cf610471565b3480156101e557600080fd5b5061018261051a565b3480156101fa57600080fd5b50610182610520565b34801561020f57600080fd5b506101ad610526565b34801561022457600080fd5b50610238610233366004611ac2565b61053a565b60405161018f9190611be9565b34801561025157600080fd5b50610182610b9d565b34801561026657600080fd5b50610182610ba3565b34801561027b57600080fd5b506101cf610ba9565b34801561029057600080fd5b50610238610c43565b3480156102a557600080fd5b50610182610c4c565b3480156102ba57600080fd5b50610182610c52565b3480156102cf57600080fd5b506101cf610c58565b3480156102e457600080fd5b50610182610c61565b3480156102f957600080fd5b50610302610c67565b60405161018f9190611c05565b34801561031b57600080fd5b506101cf610c70565b34801561033057600080fd5b5061018261033f366004611b7c565b610cf9565b34801561035057600080fd5b50610182610d1d565b34801561036557600080fd5b506101cf610374366004611b56565b610d23565b34801561038557600080fd5b50610182610e5e565b34801561039a57600080fd5b506101cf610e64565b3480156103af57600080fd5b506101cf610ea7565b60005481565b600a54600160a060020a031681565b6009546101009004600160a060020a0316600380600c5460ff1660048111156103f257fe5b146103fc57600080fd5b600b54421161040a57600080fd5b33600160a060020a0383161461041f57600080fd5b600954600254604051610100909204600160a060020a0316916108fc82150291906000818181858888f1935050505015801561045f573d6000803e3d6000fd5b5050600c805460ff1916600417905550565b33600180600c5460ff16600481111561048657fe5b1461049057600080fd5b600b54421061049e57600080fd5b33600160a060020a038316146104b357600080fd5b6002543410156104ef5760405133903480156108fc02916000818181858888f193505050501580156104e9573d6000803e3d6000fd5b50610516565b600a805473ffffffffffffffffffffffffffffffffffffffff191633179055610516610e64565b5050565b60045481565b60055481565b6009546101009004600160a060020a031681565b60006003600c5460ff16600481111561054f57fe5b1461055957600080fd5b600b54421061056757600080fd5b600a54600160a060020a0316331461057e57600080fd5b600080600060606000606060008060606004600c60006101000a81548160ff021916908360048111156105ad57fe5b02179055508c518b511461060e57600954600254604051610100909204600160a060020a0316916108fc82150291906000818181858888f193505050501580156105fb573d6000803e3d6000fd5b5060009950505050505050505050610b96565b60005b8b518110156106cb578d8181518110151561062857fe5b90602001906020020151600081518110151561064057fe5b9060200190602002015161066a8d8381518110151561065b57fe5b90602001906020020151610ead565b146106c357600954600254604051610100909204600160a060020a0316916108fc82150291906000818181858888f193505050501580156106af573d6000803e3d6000fd5b5060009a5050505050505050505050610b96565b600101610611565b508b60008151811015156106db57fe5b6020908102909101015160001a60f860020a0260f860020a900460ff16975061071c8d600081518110151561070c57fe5b9060200190602002015189610f4c565b151561076257600954600254604051610100909204600160a060020a0316916108fc82150291906000818181858888f193505050501580156105fb573d6000803e3d6000fd5b61076c8c89611164565b15156107b257600954600254604051610100909204600160a060020a0316916108fc82150291906000818181858888f193505050501580156105fb573d6000803e3d6000fd5b8b60008151811015156107c157fe5b6020908102909101015160011a60f860020a0260f860020a900498508a60008151811015156107ec57fe5b90602001906020020151516040519080825280601f01601f191660200182016040528015610821576020820181803883390190505b5090506108488b600081518110151561083657fe5b9060200190602002015160045461137b565b905060025b6008546002018110156108ab576007548d600081518110151561086c57fe5b906020019060200201518260208110151561088357fe5b1a60f860020a0260f860020a900460ff16141561089f576108ab565b6001958601950161084d565b508460ff166040519080825280602002602001820160405280156108d9578160200160208202803883390190505b50955060005b8560ff16811015610948578c60008151811015156108f957fe5b906020019060200201518160020160208110151561091357fe5b1a60f860020a0260f860020a9004878281518110151561092f57fe5b60ff9092166020928302909101909101526001016108df565b508c516000190160ff86161461096a5760009950505050505050505050610b96565b60018d51036040519080825280602002602001820160405280156109a257816020015b606081526020019060019003908161098d5790505b50935060005b8651811015610ac9576109ef8e826001018151811015156109c557fe5b9060200190602002015188838151811015156109dd57fe5b9060200190602002015160ff16610f4c565b1515610a3557600954600254604051610100909204600160a060020a0316916108fc82150291906000818181858888f193505050501580156106af573d6000803e3d6000fd5b8b81600101815181101515610a4657fe5b90602001906020020151516040519080825280601f01601f191660200182016040528015610a7b576020820181803883390190505b508582815181101515610a8a57fe5b602090810290910101528b51610aa9908d906001840190811061083657fe5b8582815181101515610ab757fe5b602090810290910101526001016109a8565b50610ad78960ff1685611472565b9250610ae281610cf9565b9150828214610b47576009805460ff19166001179055600a54600254604051600160a060020a039092169181156108fc0291906000818181858888f19350505050158015610b34573d6000803e3d6000fd5b5060019950505050505050505050610b96565b600954600254604051610100909204600160a060020a0316916108fc82150291906000818181858888f19350505050158015610b87573d6000803e3d6000fd5b50600099505050505050505050505b9392505050565b60065481565b600b5481565b600a54600160a060020a0316600380600c5460ff166004811115610bc957fe5b14610bd357600080fd5b600b544210610be157600080fd5b33600160a060020a03831614610bf657600080fd5b610bfe610e64565b600954600254604051610100909204600160a060020a0316916108fc82150291906000818181858888f19350505050158015610c3e573d6000803e3d6000fd5b505050565b60095460ff1681565b60015481565b60025481565b60144201600555565b60085481565b600c5460ff1681565b600a54600160a060020a0316600280600c5460ff166004811115610c9057fe5b14610c9a57600080fd5b600b544211610ca857600080fd5b33600160a060020a03831614610cbd57600080fd5b600a54600254604051600160a060020a039092169181156108fc0291906000818181858888f1935050505015801561045f573d6000803e3d6000fd5b805160009082901515610d10575060009050610d18565b505060208101515b919050565b60035481565b6009546101009004600160a060020a0316600280600c5460ff166004811115610d4857fe5b14610d5257600080fd5b600b544210610d6057600080fd5b33600160a060020a03831614610d7557600080fd5b6040805160208082528183019092526000916060919060208201818038833901905050905060005b6020811015610de757858160208110610db257fe5b1a60f860020a028282815181101515610dc757fe5b906020010190600160f860020a031916908160001a905350600101610d9d565b50805160208201206003549092508214610e4a57600a54600254604051600160a060020a039092169181156108fc0291906000818181858888f19350505050158015610e37573d6000803e3d6000fd5b50600c805460ff19166004179055610e57565b6004859055610e57610e64565b5050505050565b60075481565b600c5460ff166004811115610e7557fe5b6001016004811115610e8357fe5b600c805460ff19166001836004811115610e9957fe5b021790555060144201600b55565b42600555565b600080606083516040519080825280601f01601f191660200182016040528015610ede576020820181803883390190505b50905060005b8451811015610f3c578481815181101515610efb57fe5b90602001015160f860020a900460f860020a028282815181101515610f1c57fe5b906020010190600160f860020a031916908160001a905350600101610ee4565b5080516020909101209392505050565b600080836000815181101515610f5e57fe5b6020908102909101015160408051818152606081810183529293506000928392909190602082018180388339019050509050600091505b60016006540382101561115257600282900a868116046001141561107f5760005b602081101561101c578783600101815181101515610fd057fe5b9060200190602002015181602081101515610fe757fe5b1a60f860020a028282815181101515610ffc57fe5b906020010190600160f860020a031916908160001a905350600101610fb6565b5060005b602081101561106d5784816020811061103557fe5b1a60f860020a02828260200181518110151561104d57fe5b906020010190600160f860020a031916908160001a905350600101611020565b50805160208201209350839250611147565b60005b60208110156110cc5784816020811061109757fe5b1a60f860020a0282828151811015156110ac57fe5b906020010190600160f860020a031916908160001a905350600101611082565b5060005b60208110156111395787836001018151811015156110ea57fe5b906020019060200201518160208110151561110157fe5b1a60f860020a02828260200181518110151561111957fe5b906020010190600160f860020a031916908160001a9053506001016110d0565b508051602082012093508392505b600190910190610f95565b60005484149450505050505b92915050565b60008083600081518110151561117657fe5b6020908102909101015160408051818152606081810183529293506000928392909190602082018180388339019050509050600091505b60016006540382101561136a57600282900a86811604600114156112975760005b60208110156112345787836001018151811015156111e857fe5b90602001906020020151816020811015156111ff57fe5b1a60f860020a02828281518110151561121457fe5b906020010190600160f860020a031916908160001a9053506001016111ce565b5060005b60208110156112855784816020811061124d57fe5b1a60f860020a02828260200181518110151561126557fe5b906020010190600160f860020a031916908160001a905350600101611238565b5080516020820120935083925061135f565b60005b60208110156112e4578481602081106112af57fe5b1a60f860020a0282828151811015156112c457fe5b906020010190600160f860020a031916908160001a90535060010161129a565b5060005b602081101561135157878360010181518110151561130257fe5b906020019060200201518160208110151561131957fe5b1a60f860020a02828260200181518110151561133157fe5b906020010190600160f860020a031916908160001a9053506001016112e8565b508051602082012093508392505b6001909101906111ad565b600154841494505050505092915050565b606080606084516040519080825280601f01601f1916602001820160405280156113ac576020820181803883390190505b5090506113ba84865161160d565b905084516040519080825280601f01601f1916602001820160405280156113e8576020820181803883390190505b50915060005b855181101561146857858181518110151561140557fe5b90602001015160f860020a900460f860020a02828281518110151561142657fe5b90602001015160f860020a900460f860020a0218838281518110151561144857fe5b906020010190600160f860020a031916908160001a9053506001016113ee565b5090949350505050565b6000808360021415610b9657606083600081518110151561148f57fe5b90602001906020020151516002026040519080825280601f01601f1916602001820160405280156114c7576020820181803883390190505b50905060005b8460008151811015156114dc57fe5b9060200190602002015151811015611555578460008151811015156114fd57fe5b906020019060200201518181518110151561151457fe5b90602001015160f860020a900460f860020a02828281518110151561153557fe5b906020010190600160f860020a031916908160001a9053506001016114cd565b5060005b84600181518110151561156857fe5b90602001906020020151518110156115fc5784600181518110151561158957fe5b90602001906020020151818151811015156115a057fe5b90602001015160f860020a900460f860020a02828660008151811015156115c357fe5b906020019060200201515183018151811015156115dc57fe5b906020010190600160f860020a031916908160001a905350600101611559565b508051602090910120949350505050565b60608060606000846040519080825280601f01601f19166020018201604052801561163f576020820181803883390190505b5060408051602080825281830190925291945060208201818038833901905050915060005b602086048110156118b3578015156117a457604080516021808252606082810190935260208201818038833901905050905060005b60208110156116e3578881602081106116ae57fe5b1a60f860020a0282828151811015156116c357fe5b906020010190600160f860020a031916908160001a905350600101611699565b508051600090829060209081106116f657fe5b906020010190600160f860020a031916908160001a90535080516020820120925060005b602081101561179d5783816020811061172f57fe5b1a60f860020a02858281518110151561174457fe5b906020010190600160f860020a031916908160001a90535083816020811061176857fe5b1a60f860020a02868281518110151561177d57fe5b906020010190600160f860020a031916908160001a90535060010161171a565b50506118ab565b604080516020808252818301909252606094509060208201818038833901905050925060005b6020811015611814578281602081106117df57fe5b1a60f860020a0284828151811015156117f457fe5b906020010190600160f860020a031916908160001a9053506001016117ca565b5082516020840120915060005b60208110156118a95782816020811061183657fe5b1a60f860020a02848281518110151561184b57fe5b906020010190600160f860020a031916908160001a90535082816020811061186f57fe5b1a60f860020a028582846020020181518110151561188957fe5b906020010190600160f860020a031916908160001a905350600101611821565b505b600101611664565b509195945050505050565b6000601f820183136118cf57600080fd5b81356118e26118dd82611c3a565b611c13565b81815260209384019390925082018360005b83811015611920578135860161190a888261192a565b84525060209283019291909101906001016118f4565b5050505092915050565b6000601f8201831361193b57600080fd5b81356119496118dd82611c3a565b9150818183526020840193506020810190508385602084028201111561196e57600080fd5b60005b8381101561192057816119848882611a67565b8452506020928301929190910190600101611971565b6000601f820183136119ab57600080fd5b81356119b96118dd82611c3a565b915081818352602084019350602081019050838560208402820111156119de57600080fd5b60005b8381101561192057816119f48882611a67565b84525060209283019291909101906001016119e1565b6000601f82018313611a1b57600080fd5b8135611a296118dd82611c3a565b81815260209384019390925082018360005b838110156119205781358601611a518882611a73565b8452506020928301929190910190600101611a3b565b6000610b968235611c93565b6000601f82018313611a8457600080fd5b8135611a926118dd82611c5b565b91508082526020830160208301858383011115611aae57600080fd5b611ab9838284611cbd565b50505092915050565b600080600060608486031215611ad757600080fd5b833567ffffffffffffffff811115611aee57600080fd5b611afa868287016118be565b935050602084013567ffffffffffffffff811115611b1757600080fd5b611b238682870161199a565b925050604084013567ffffffffffffffff811115611b4057600080fd5b611b4c86828701611a0a565b9150509250925092565b600060208284031215611b6857600080fd5b6000611b748484611a67565b949350505050565b600060208284031215611b8e57600080fd5b813567ffffffffffffffff811115611ba557600080fd5b611b7484828501611a73565b611bba81611c83565b82525050565b611bba81611c8e565b611bba81611c93565b611bba81611cb2565b6020810161115e8284611bb1565b6020810161115e8284611bc0565b6020810161115e8284611bc9565b6020810161115e8284611bd2565b60405181810167ffffffffffffffff81118282101715611c3257600080fd5b604052919050565b600067ffffffffffffffff821115611c5157600080fd5b5060209081020190565b600067ffffffffffffffff821115611c7257600080fd5b506020601f91909101601f19160190565b600061115e82611ca6565b151590565b90565b600060058210611ca257fe5b5090565b600160a060020a031690565b600061115e82611c96565b8281833750600091015256fea265627a7a72305820ac67425267f64a9e5d7f052f4ae3e587db8b4418e8cdec5aa9ff22d692ff136e6c6578706572696d656e74616cf50037`

// DeployFairswap deploys a new Ethereum contract, binding an instance of Fairswap to it.
func DeployFairswap(auth *bind.TransactOpts, backend bind.ContractBackend, _price *big.Int, _commitmentOfKey [32]byte, _EncInputRoot [32]byte, _EncCircuitRoot [32]byte, _totalNumberOfGates *big.Int, _maxInputLinesToAGate *big.Int) (common.Address, *types.Transaction, *Fairswap, error) {
	parsed, err := abi.JSON(strings.NewReader(FairswapABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FairswapBin), backend, _price, _commitmentOfKey, _EncInputRoot, _EncCircuitRoot, _totalNumberOfGates, _maxInputLinesToAGate)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Fairswap{FairswapCaller: FairswapCaller{contract: contract}, FairswapTransactor: FairswapTransactor{contract: contract}, FairswapFilterer: FairswapFilterer{contract: contract}}, nil
}

// Fairswap is an auto generated Go binding around an Ethereum contract.
type Fairswap struct {
	FairswapCaller     // Read-only binding to the contract
	FairswapTransactor // Write-only binding to the contract
	FairswapFilterer   // Log filterer for contract events
}

// FairswapCaller is an auto generated read-only Go binding around an Ethereum contract.
type FairswapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FairswapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FairswapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FairswapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FairswapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FairswapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FairswapSession struct {
	Contract     *Fairswap         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FairswapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FairswapCallerSession struct {
	Contract *FairswapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// FairswapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FairswapTransactorSession struct {
	Contract     *FairswapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// FairswapRaw is an auto generated low-level Go binding around an Ethereum contract.
type FairswapRaw struct {
	Contract *Fairswap // Generic contract binding to access the raw methods on
}

// FairswapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FairswapCallerRaw struct {
	Contract *FairswapCaller // Generic read-only contract binding to access the raw methods on
}

// FairswapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FairswapTransactorRaw struct {
	Contract *FairswapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFairswap creates a new instance of Fairswap, bound to a specific deployed contract.
func NewFairswap(address common.Address, backend bind.ContractBackend) (*Fairswap, error) {
	contract, err := bindFairswap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Fairswap{FairswapCaller: FairswapCaller{contract: contract}, FairswapTransactor: FairswapTransactor{contract: contract}, FairswapFilterer: FairswapFilterer{contract: contract}}, nil
}

// NewFairswapCaller creates a new read-only instance of Fairswap, bound to a specific deployed contract.
func NewFairswapCaller(address common.Address, caller bind.ContractCaller) (*FairswapCaller, error) {
	contract, err := bindFairswap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FairswapCaller{contract: contract}, nil
}

// NewFairswapTransactor creates a new write-only instance of Fairswap, bound to a specific deployed contract.
func NewFairswapTransactor(address common.Address, transactor bind.ContractTransactor) (*FairswapTransactor, error) {
	contract, err := bindFairswap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FairswapTransactor{contract: contract}, nil
}

// NewFairswapFilterer creates a new log filterer instance of Fairswap, bound to a specific deployed contract.
func NewFairswapFilterer(address common.Address, filterer bind.ContractFilterer) (*FairswapFilterer, error) {
	contract, err := bindFairswap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FairswapFilterer{contract: contract}, nil
}

// bindFairswap binds a generic wrapper to an already deployed contract.
func bindFairswap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FairswapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fairswap *FairswapRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Fairswap.Contract.FairswapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fairswap *FairswapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fairswap.Contract.FairswapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fairswap *FairswapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fairswap.Contract.FairswapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fairswap *FairswapCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Fairswap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fairswap *FairswapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fairswap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fairswap *FairswapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fairswap.Contract.contract.Transact(opts, method, params...)
}

// EncCircuitRoot is a free data retrieval call binding the contract method 0x7a514594.
//
// Solidity: function EncCircuitRoot() constant returns(bytes32)
func (_Fairswap *FairswapCaller) EncCircuitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Fairswap.contract.Call(opts, out, "EncCircuitRoot")
	return *ret0, err
}

// EncCircuitRoot is a free data retrieval call binding the contract method 0x7a514594.
//
// Solidity: function EncCircuitRoot() constant returns(bytes32)
func (_Fairswap *FairswapSession) EncCircuitRoot() ([32]byte, error) {
	return _Fairswap.Contract.EncCircuitRoot(&_Fairswap.CallOpts)
}

// EncCircuitRoot is a free data retrieval call binding the contract method 0x7a514594.
//
// Solidity: function EncCircuitRoot() constant returns(bytes32)
func (_Fairswap *FairswapCallerSession) EncCircuitRoot() ([32]byte, error) {
	return _Fairswap.Contract.EncCircuitRoot(&_Fairswap.CallOpts)
}

// EncInputRoot is a free data retrieval call binding the contract method 0x062a1398.
//
// Solidity: function EncInputRoot() constant returns(bytes32)
func (_Fairswap *FairswapCaller) EncInputRoot(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Fairswap.contract.Call(opts, out, "EncInputRoot")
	return *ret0, err
}

// EncInputRoot is a free data retrieval call binding the contract method 0x062a1398.
//
// Solidity: function EncInputRoot() constant returns(bytes32)
func (_Fairswap *FairswapSession) EncInputRoot() ([32]byte, error) {
	return _Fairswap.Contract.EncInputRoot(&_Fairswap.CallOpts)
}

// EncInputRoot is a free data retrieval call binding the contract method 0x062a1398.
//
// Solidity: function EncInputRoot() constant returns(bytes32)
func (_Fairswap *FairswapCallerSession) EncInputRoot() ([32]byte, error) {
	return _Fairswap.Contract.EncInputRoot(&_Fairswap.CallOpts)
}

// Now is a free data retrieval call binding the contract method 0x44d4fd19.
//
// Solidity: function Now() constant returns(uint256)
func (_Fairswap *FairswapCaller) Now(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Fairswap.contract.Call(opts, out, "Now")
	return *ret0, err
}

// Now is a free data retrieval call binding the contract method 0x44d4fd19.
//
// Solidity: function Now() constant returns(uint256)
func (_Fairswap *FairswapSession) Now() (*big.Int, error) {
	return _Fairswap.Contract.Now(&_Fairswap.CallOpts)
}

// Now is a free data retrieval call binding the contract method 0x44d4fd19.
//
// Solidity: function Now() constant returns(uint256)
func (_Fairswap *FairswapCallerSession) Now() (*big.Int, error) {
	return _Fairswap.Contract.Now(&_Fairswap.CallOpts)
}

// CommitmentOfKey is a free data retrieval call binding the contract method 0xca641f86.
//
// Solidity: function commitmentOfKey() constant returns(bytes32)
func (_Fairswap *FairswapCaller) CommitmentOfKey(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Fairswap.contract.Call(opts, out, "commitmentOfKey")
	return *ret0, err
}

// CommitmentOfKey is a free data retrieval call binding the contract method 0xca641f86.
//
// Solidity: function commitmentOfKey() constant returns(bytes32)
func (_Fairswap *FairswapSession) CommitmentOfKey() ([32]byte, error) {
	return _Fairswap.Contract.CommitmentOfKey(&_Fairswap.CallOpts)
}

// CommitmentOfKey is a free data retrieval call binding the contract method 0xca641f86.
//
// Solidity: function commitmentOfKey() constant returns(bytes32)
func (_Fairswap *FairswapCallerSession) CommitmentOfKey() ([32]byte, error) {
	return _Fairswap.Contract.CommitmentOfKey(&_Fairswap.CallOpts)
}

// Depth is a free data retrieval call binding the contract method 0x631c56ef.
//
// Solidity: function depth() constant returns(uint256)
func (_Fairswap *FairswapCaller) Depth(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Fairswap.contract.Call(opts, out, "depth")
	return *ret0, err
}

// Depth is a free data retrieval call binding the contract method 0x631c56ef.
//
// Solidity: function depth() constant returns(uint256)
func (_Fairswap *FairswapSession) Depth() (*big.Int, error) {
	return _Fairswap.Contract.Depth(&_Fairswap.CallOpts)
}

// Depth is a free data retrieval call binding the contract method 0x631c56ef.
//
// Solidity: function depth() constant returns(uint256)
func (_Fairswap *FairswapCallerSession) Depth() (*big.Int, error) {
	return _Fairswap.Contract.Depth(&_Fairswap.CallOpts)
}

// Key is a free data retrieval call binding the contract method 0x3943380c.
//
// Solidity: function key() constant returns(bytes32)
func (_Fairswap *FairswapCaller) Key(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Fairswap.contract.Call(opts, out, "key")
	return *ret0, err
}

// Key is a free data retrieval call binding the contract method 0x3943380c.
//
// Solidity: function key() constant returns(bytes32)
func (_Fairswap *FairswapSession) Key() ([32]byte, error) {
	return _Fairswap.Contract.Key(&_Fairswap.CallOpts)
}

// Key is a free data retrieval call binding the contract method 0x3943380c.
//
// Solidity: function key() constant returns(bytes32)
func (_Fairswap *FairswapCallerSession) Key() ([32]byte, error) {
	return _Fairswap.Contract.Key(&_Fairswap.CallOpts)
}

// MaxInputLinesToAGate is a free data retrieval call binding the contract method 0xaa9b8ac7.
//
// Solidity: function maxInputLinesToAGate() constant returns(uint256)
func (_Fairswap *FairswapCaller) MaxInputLinesToAGate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Fairswap.contract.Call(opts, out, "maxInputLinesToAGate")
	return *ret0, err
}

// MaxInputLinesToAGate is a free data retrieval call binding the contract method 0xaa9b8ac7.
//
// Solidity: function maxInputLinesToAGate() constant returns(uint256)
func (_Fairswap *FairswapSession) MaxInputLinesToAGate() (*big.Int, error) {
	return _Fairswap.Contract.MaxInputLinesToAGate(&_Fairswap.CallOpts)
}

// MaxInputLinesToAGate is a free data retrieval call binding the contract method 0xaa9b8ac7.
//
// Solidity: function maxInputLinesToAGate() constant returns(uint256)
func (_Fairswap *FairswapCallerSession) MaxInputLinesToAGate() (*big.Int, error) {
	return _Fairswap.Contract.MaxInputLinesToAGate(&_Fairswap.CallOpts)
}

// Phase is a free data retrieval call binding the contract method 0xb1c9fe6e.
//
// Solidity: function phase() constant returns(uint8)
func (_Fairswap *FairswapCaller) Phase(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Fairswap.contract.Call(opts, out, "phase")
	return *ret0, err
}

// Phase is a free data retrieval call binding the contract method 0xb1c9fe6e.
//
// Solidity: function phase() constant returns(uint8)
func (_Fairswap *FairswapSession) Phase() (uint8, error) {
	return _Fairswap.Contract.Phase(&_Fairswap.CallOpts)
}

// Phase is a free data retrieval call binding the contract method 0xb1c9fe6e.
//
// Solidity: function phase() constant returns(uint8)
func (_Fairswap *FairswapCallerSession) Phase() (uint8, error) {
	return _Fairswap.Contract.Phase(&_Fairswap.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_Fairswap *FairswapCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Fairswap.contract.Call(opts, out, "price")
	return *ret0, err
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_Fairswap *FairswapSession) Price() (*big.Int, error) {
	return _Fairswap.Contract.Price(&_Fairswap.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_Fairswap *FairswapCallerSession) Price() (*big.Int, error) {
	return _Fairswap.Contract.Price(&_Fairswap.CallOpts)
}

// ReceiverAddress is a free data retrieval call binding the contract method 0x16fed3e2.
//
// Solidity: function receiverAddress() constant returns(address)
func (_Fairswap *FairswapCaller) ReceiverAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Fairswap.contract.Call(opts, out, "receiverAddress")
	return *ret0, err
}

// ReceiverAddress is a free data retrieval call binding the contract method 0x16fed3e2.
//
// Solidity: function receiverAddress() constant returns(address)
func (_Fairswap *FairswapSession) ReceiverAddress() (common.Address, error) {
	return _Fairswap.Contract.ReceiverAddress(&_Fairswap.CallOpts)
}

// ReceiverAddress is a free data retrieval call binding the contract method 0x16fed3e2.
//
// Solidity: function receiverAddress() constant returns(address)
func (_Fairswap *FairswapCallerSession) ReceiverAddress() (common.Address, error) {
	return _Fairswap.Contract.ReceiverAddress(&_Fairswap.CallOpts)
}

// SenderAddress is a free data retrieval call binding the contract method 0x4fc84791.
//
// Solidity: function senderAddress() constant returns(address)
func (_Fairswap *FairswapCaller) SenderAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Fairswap.contract.Call(opts, out, "senderAddress")
	return *ret0, err
}

// SenderAddress is a free data retrieval call binding the contract method 0x4fc84791.
//
// Solidity: function senderAddress() constant returns(address)
func (_Fairswap *FairswapSession) SenderAddress() (common.Address, error) {
	return _Fairswap.Contract.SenderAddress(&_Fairswap.CallOpts)
}

// SenderAddress is a free data retrieval call binding the contract method 0x4fc84791.
//
// Solidity: function senderAddress() constant returns(address)
func (_Fairswap *FairswapCallerSession) SenderAddress() (common.Address, error) {
	return _Fairswap.Contract.SenderAddress(&_Fairswap.CallOpts)
}

// Timeout is a free data retrieval call binding the contract method 0x70dea79a.
//
// Solidity: function timeout() constant returns(uint256)
func (_Fairswap *FairswapCaller) Timeout(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Fairswap.contract.Call(opts, out, "timeout")
	return *ret0, err
}

// Timeout is a free data retrieval call binding the contract method 0x70dea79a.
//
// Solidity: function timeout() constant returns(uint256)
func (_Fairswap *FairswapSession) Timeout() (*big.Int, error) {
	return _Fairswap.Contract.Timeout(&_Fairswap.CallOpts)
}

// Timeout is a free data retrieval call binding the contract method 0x70dea79a.
//
// Solidity: function timeout() constant returns(uint256)
func (_Fairswap *FairswapCallerSession) Timeout() (*big.Int, error) {
	return _Fairswap.Contract.Timeout(&_Fairswap.CallOpts)
}

// TotalNumberOfGates is a free data retrieval call binding the contract method 0xe44b23c8.
//
// Solidity: function totalNumberOfGates() constant returns(uint256)
func (_Fairswap *FairswapCaller) TotalNumberOfGates(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Fairswap.contract.Call(opts, out, "totalNumberOfGates")
	return *ret0, err
}

// TotalNumberOfGates is a free data retrieval call binding the contract method 0xe44b23c8.
//
// Solidity: function totalNumberOfGates() constant returns(uint256)
func (_Fairswap *FairswapSession) TotalNumberOfGates() (*big.Int, error) {
	return _Fairswap.Contract.TotalNumberOfGates(&_Fairswap.CallOpts)
}

// TotalNumberOfGates is a free data retrieval call binding the contract method 0xe44b23c8.
//
// Solidity: function totalNumberOfGates() constant returns(uint256)
func (_Fairswap *FairswapCallerSession) TotalNumberOfGates() (*big.Int, error) {
	return _Fairswap.Contract.TotalNumberOfGates(&_Fairswap.CallOpts)
}

// ValidComplain is a free data retrieval call binding the contract method 0x76aa657f.
//
// Solidity: function validComplain() constant returns(bool)
func (_Fairswap *FairswapCaller) ValidComplain(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Fairswap.contract.Call(opts, out, "validComplain")
	return *ret0, err
}

// ValidComplain is a free data retrieval call binding the contract method 0x76aa657f.
//
// Solidity: function validComplain() constant returns(bool)
func (_Fairswap *FairswapSession) ValidComplain() (bool, error) {
	return _Fairswap.Contract.ValidComplain(&_Fairswap.CallOpts)
}

// ValidComplain is a free data retrieval call binding the contract method 0x76aa657f.
//
// Solidity: function validComplain() constant returns(bool)
func (_Fairswap *FairswapCallerSession) ValidComplain() (bool, error) {
	return _Fairswap.Contract.ValidComplain(&_Fairswap.CallOpts)
}

// Nowy is a paid mutator transaction binding the contract method 0xfd6fd6fa.
//
// Solidity: function Nowy() returns()
func (_Fairswap *FairswapTransactor) Nowy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fairswap.contract.Transact(opts, "Nowy")
}

// Nowy is a paid mutator transaction binding the contract method 0xfd6fd6fa.
//
// Solidity: function Nowy() returns()
func (_Fairswap *FairswapSession) Nowy() (*types.Transaction, error) {
	return _Fairswap.Contract.Nowy(&_Fairswap.TransactOpts)
}

// Nowy is a paid mutator transaction binding the contract method 0xfd6fd6fa.
//
// Solidity: function Nowy() returns()
func (_Fairswap *FairswapTransactorSession) Nowy() (*types.Transaction, error) {
	return _Fairswap.Contract.Nowy(&_Fairswap.TransactOpts)
}

// Nowy2 is a paid mutator transaction binding the contract method 0xa2b26c06.
//
// Solidity: function Nowy2() returns()
func (_Fairswap *FairswapTransactor) Nowy2(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fairswap.contract.Transact(opts, "Nowy2")
}

// Nowy2 is a paid mutator transaction binding the contract method 0xa2b26c06.
//
// Solidity: function Nowy2() returns()
func (_Fairswap *FairswapSession) Nowy2() (*types.Transaction, error) {
	return _Fairswap.Contract.Nowy2(&_Fairswap.TransactOpts)
}

// Nowy2 is a paid mutator transaction binding the contract method 0xa2b26c06.
//
// Solidity: function Nowy2() returns()
func (_Fairswap *FairswapTransactorSession) Nowy2() (*types.Transaction, error) {
	return _Fairswap.Contract.Nowy2(&_Fairswap.TransactOpts)
}

// Accept is a paid mutator transaction binding the contract method 0x2852b71c.
//
// Solidity: function accept() returns()
func (_Fairswap *FairswapTransactor) Accept(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fairswap.contract.Transact(opts, "accept")
}

// Accept is a paid mutator transaction binding the contract method 0x2852b71c.
//
// Solidity: function accept() returns()
func (_Fairswap *FairswapSession) Accept() (*types.Transaction, error) {
	return _Fairswap.Contract.Accept(&_Fairswap.TransactOpts)
}

// Accept is a paid mutator transaction binding the contract method 0x2852b71c.
//
// Solidity: function accept() returns()
func (_Fairswap *FairswapTransactorSession) Accept() (*types.Transaction, error) {
	return _Fairswap.Contract.Accept(&_Fairswap.TransactOpts)
}

// BytesToBytes32 is a paid mutator transaction binding the contract method 0xbfe370d9.
//
// Solidity: function bytesToBytes32(bytes source) returns(bytes32 result)
func (_Fairswap *FairswapTransactor) BytesToBytes32(opts *bind.TransactOpts, source []byte) (*types.Transaction, error) {
	return _Fairswap.contract.Transact(opts, "bytesToBytes32", source)
}

// BytesToBytes32 is a paid mutator transaction binding the contract method 0xbfe370d9.
//
// Solidity: function bytesToBytes32(bytes source) returns(bytes32 result)
func (_Fairswap *FairswapSession) BytesToBytes32(source []byte) (*types.Transaction, error) {
	return _Fairswap.Contract.BytesToBytes32(&_Fairswap.TransactOpts, source)
}

// BytesToBytes32 is a paid mutator transaction binding the contract method 0xbfe370d9.
//
// Solidity: function bytesToBytes32(bytes source) returns(bytes32 result)
func (_Fairswap *FairswapTransactorSession) BytesToBytes32(source []byte) (*types.Transaction, error) {
	return _Fairswap.Contract.BytesToBytes32(&_Fairswap.TransactOpts, source)
}

// Complain is a paid mutator transaction binding the contract method 0x5c57b308.
//
// Solidity: function complain(bytes32[][] complaint, bytes32[] gateMerkleProof, bytes[] encodedVectors_ThisGate) returns(bool)
func (_Fairswap *FairswapTransactor) Complain(opts *bind.TransactOpts, complaint [][][32]byte, gateMerkleProof [][32]byte, encodedVectors_ThisGate [][]byte) (*types.Transaction, error) {
	return _Fairswap.contract.Transact(opts, "complain", complaint, gateMerkleProof, encodedVectors_ThisGate)
}

// Complain is a paid mutator transaction binding the contract method 0x5c57b308.
//
// Solidity: function complain(bytes32[][] complaint, bytes32[] gateMerkleProof, bytes[] encodedVectors_ThisGate) returns(bool)
func (_Fairswap *FairswapSession) Complain(complaint [][][32]byte, gateMerkleProof [][32]byte, encodedVectors_ThisGate [][]byte) (*types.Transaction, error) {
	return _Fairswap.Contract.Complain(&_Fairswap.TransactOpts, complaint, gateMerkleProof, encodedVectors_ThisGate)
}

// Complain is a paid mutator transaction binding the contract method 0x5c57b308.
//
// Solidity: function complain(bytes32[][] complaint, bytes32[] gateMerkleProof, bytes[] encodedVectors_ThisGate) returns(bool)
func (_Fairswap *FairswapTransactorSession) Complain(complaint [][][32]byte, gateMerkleProof [][32]byte, encodedVectors_ThisGate [][]byte) (*types.Transaction, error) {
	return _Fairswap.Contract.Complain(&_Fairswap.TransactOpts, complaint, gateMerkleProof, encodedVectors_ThisGate)
}

// GotTheCorrectFiles is a paid mutator transaction binding the contract method 0x75986d64.
//
// Solidity: function gotTheCorrectFiles() returns()
func (_Fairswap *FairswapTransactor) GotTheCorrectFiles(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fairswap.contract.Transact(opts, "gotTheCorrectFiles")
}

// GotTheCorrectFiles is a paid mutator transaction binding the contract method 0x75986d64.
//
// Solidity: function gotTheCorrectFiles() returns()
func (_Fairswap *FairswapSession) GotTheCorrectFiles() (*types.Transaction, error) {
	return _Fairswap.Contract.GotTheCorrectFiles(&_Fairswap.TransactOpts)
}

// GotTheCorrectFiles is a paid mutator transaction binding the contract method 0x75986d64.
//
// Solidity: function gotTheCorrectFiles() returns()
func (_Fairswap *FairswapTransactorSession) GotTheCorrectFiles() (*types.Transaction, error) {
	return _Fairswap.Contract.GotTheCorrectFiles(&_Fairswap.TransactOpts)
}

// NextStage is a paid mutator transaction binding the contract method 0xee3743ab.
//
// Solidity: function nextStage() returns()
func (_Fairswap *FairswapTransactor) NextStage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fairswap.contract.Transact(opts, "nextStage")
}

// NextStage is a paid mutator transaction binding the contract method 0xee3743ab.
//
// Solidity: function nextStage() returns()
func (_Fairswap *FairswapSession) NextStage() (*types.Transaction, error) {
	return _Fairswap.Contract.NextStage(&_Fairswap.TransactOpts)
}

// NextStage is a paid mutator transaction binding the contract method 0xee3743ab.
//
// Solidity: function nextStage() returns()
func (_Fairswap *FairswapTransactorSession) NextStage() (*types.Transaction, error) {
	return _Fairswap.Contract.NextStage(&_Fairswap.TransactOpts)
}

// ReceiverGetEtherSenderNotRevealedKey is a paid mutator transaction binding the contract method 0xb9777391.
//
// Solidity: function receiverGetEther_SenderNotRevealedKey() returns()
func (_Fairswap *FairswapTransactor) ReceiverGetEtherSenderNotRevealedKey(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fairswap.contract.Transact(opts, "receiverGetEther_SenderNotRevealedKey")
}

// ReceiverGetEtherSenderNotRevealedKey is a paid mutator transaction binding the contract method 0xb9777391.
//
// Solidity: function receiverGetEther_SenderNotRevealedKey() returns()
func (_Fairswap *FairswapSession) ReceiverGetEtherSenderNotRevealedKey() (*types.Transaction, error) {
	return _Fairswap.Contract.ReceiverGetEtherSenderNotRevealedKey(&_Fairswap.TransactOpts)
}

// ReceiverGetEtherSenderNotRevealedKey is a paid mutator transaction binding the contract method 0xb9777391.
//
// Solidity: function receiverGetEther_SenderNotRevealedKey() returns()
func (_Fairswap *FairswapTransactorSession) ReceiverGetEtherSenderNotRevealedKey() (*types.Transaction, error) {
	return _Fairswap.Contract.ReceiverGetEtherSenderNotRevealedKey(&_Fairswap.TransactOpts)
}

// RevealKey is a paid mutator transaction binding the contract method 0xd6547236.
//
// Solidity: function revealKey(bytes32 _key) returns()
func (_Fairswap *FairswapTransactor) RevealKey(opts *bind.TransactOpts, _key [32]byte) (*types.Transaction, error) {
	return _Fairswap.contract.Transact(opts, "revealKey", _key)
}

// RevealKey is a paid mutator transaction binding the contract method 0xd6547236.
//
// Solidity: function revealKey(bytes32 _key) returns()
func (_Fairswap *FairswapSession) RevealKey(_key [32]byte) (*types.Transaction, error) {
	return _Fairswap.Contract.RevealKey(&_Fairswap.TransactOpts, _key)
}

// RevealKey is a paid mutator transaction binding the contract method 0xd6547236.
//
// Solidity: function revealKey(bytes32 _key) returns()
func (_Fairswap *FairswapTransactorSession) RevealKey(_key [32]byte) (*types.Transaction, error) {
	return _Fairswap.Contract.RevealKey(&_Fairswap.TransactOpts, _key)
}

// SenderGetEtherReceiverNoComplainMessage is a paid mutator transaction binding the contract method 0x264c57e1.
//
// Solidity: function senderGetEther_ReceiverNoComplainMessage() returns()
func (_Fairswap *FairswapTransactor) SenderGetEtherReceiverNoComplainMessage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fairswap.contract.Transact(opts, "senderGetEther_ReceiverNoComplainMessage")
}

// SenderGetEtherReceiverNoComplainMessage is a paid mutator transaction binding the contract method 0x264c57e1.
//
// Solidity: function senderGetEther_ReceiverNoComplainMessage() returns()
func (_Fairswap *FairswapSession) SenderGetEtherReceiverNoComplainMessage() (*types.Transaction, error) {
	return _Fairswap.Contract.SenderGetEtherReceiverNoComplainMessage(&_Fairswap.TransactOpts)
}

// SenderGetEtherReceiverNoComplainMessage is a paid mutator transaction binding the contract method 0x264c57e1.
//
// Solidity: function senderGetEther_ReceiverNoComplainMessage() returns()
func (_Fairswap *FairswapTransactorSession) SenderGetEtherReceiverNoComplainMessage() (*types.Transaction, error) {
	return _Fairswap.Contract.SenderGetEtherReceiverNoComplainMessage(&_Fairswap.TransactOpts)
}
