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
const FairswapABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"EncInputRoot\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"receiverAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"accept\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"key\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"receiverGetEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Now\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"senderAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"complaint\",\"type\":\"bytes32[][]\"},{\"name\":\"gateMerkleProof\",\"type\":\"bytes32[]\"},{\"name\":\"encodedVectors_ThisGate\",\"type\":\"bytes[]\"}],\"name\":\"complain\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"depth\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"timeout\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"validComplain\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EncCircuitRoot\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"complaint\",\"type\":\"bytes32[]\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"Mverify\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxInputLinesToAGate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"gateProof\",\"type\":\"bytes32[]\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"MverifyForGate\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"phase\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"source\",\"type\":\"bytes\"}],\"name\":\"bytesToBytes32\",\"outputs\":[{\"name\":\"result\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"commitmentOfKey\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"revealKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalNumberOfGates\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"nextStage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"senderGetEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_commitmentOfKey\",\"type\":\"bytes32\"},{\"name\":\"_EncInputRoot\",\"type\":\"bytes32\"},{\"name\":\"_EncCircuitRoot\",\"type\":\"bytes32\"},{\"name\":\"_totalNumberOfGates\",\"type\":\"uint256\"},{\"name\":\"_maxInputLinesToAGate\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// FairswapBin is the compiled bytecode used for deploying new contracts.
const FairswapBin = `0x6080604052600c805460ff191690553480156200001b57600080fd5b5060405160c08062001fbe83398101806040526200003d91908101906200030c565b6009805461010060a860020a03191633610100021790556002869055600385905560008490556001839055600782905560088190556200008682640100000000620000ab810204565b6001016006556200009f640100000000620002ae810204565b505050505050620003a3565b604080517ff8f9cbfae6cc78fbefe7cdc3a1793dfcf4f0e8bbd8cec470b6a28a7a5a3e1efd81527ff5ecf1b3e9debc68e1d9cfabc5997135bfb7a7a3938b7b606b5b4b3f2f1f0ffe60208201527ff6e4ed9ff2d6b458eadcdf97bd91692de2d4da8fd2d0ac50c6ae9a8272523616818301527fc8c0b887b0a8a4489c948c7f847c6125746c645c544c444038302820181008ff60608201527ff7cae577eec2a03cf3bad76fb589591debb2dd67e0aa9834bea6925f6a4a2e0e60808201527fe39ed557db96902cd38ed14fad815115c786af479b7e8324736353433727170760a08201527fc976c13bb96e881cb166a933a55e490d9d56952b8d4e801485467d236242260660c08201527f753a6d1b65325d0c552a4d1345224105391a310b29122104190a11030902010060e08201526101008082019092527f0100000000000000000000000000000000000000000000000000000000000000600170010000000000000000000000000000000068010000000000000000640100000000620100006010600460026000198c019081041790810417908104178881041790810417908104179081041790810417017e818283848586878898a8b8c8d8e8f929395969799a9b9d9e9faaeb6bedeeff0281900460ff0390910151047f8000000000000000000000000000000000000000000000000000000000000000909211020190565b600c5460ff166006811115620002c057fe5b6001016006811115620002cf57fe5b600c805460ff19166001836006811115620002e657fe5b0217905550600a4204600a01600b55565b6000620003058251620003a0565b9392505050565b60008060008060008060c087890312156200032657600080fd5b6000620003348989620002f7565b96505060206200034789828a01620002f7565b95505060406200035a89828a01620002f7565b94505060606200036d89828a01620002f7565b93505060806200038089828a01620002f7565b92505060a06200039389828a01620002f7565b9150509295509295509295565b90565b611c0b80620003b36000396000f3fe60806040526004361061015d576000357c0100000000000000000000000000000000000000000000000000000000900480637a514594116100d3578063bfe370d91161008c578063bfe370d91461031a578063ca641f861461033a578063d65472361461034f578063e44b23c81461036f578063ee3743ab14610384578063f6fa345f146103995761015d565b80637a51459414610279578063a035b1fe1461028e578063a2630372146102a3578063aa9b8ac7146102c3578063afa31a42146102d8578063b1c9fe6e146102f85761015d565b806344d4fd191161012557806344d4fd19146101e35780634fc84791146101f85780635c57b3081461020d578063631c56ef1461023a57806370dea79a1461024f57806376aa657f146102645761015d565b8063062a13981461016257806316fed3e21461018d5780632852b71c146101af5780633943380c146101b95780633e37aeb6146101ce575b600080fd5b34801561016e57600080fd5b506101776103ae565b6040516101849190611aff565b60405180910390f35b34801561019957600080fd5b506101a26103b4565b6040516101849190611ae3565b6101b76103c3565b005b3480156101c557600080fd5b5061017761047c565b3480156101da57600080fd5b506101b7610482565b3480156101ef57600080fd5b50610177610531565b34801561020457600080fd5b506101a2610537565b34801561021957600080fd5b5061022d610228366004611979565b61054b565b6040516101849190611af1565b34801561024657600080fd5b50610177610ad1565b34801561025b57600080fd5b50610177610ad7565b34801561027057600080fd5b5061022d610add565b34801561028557600080fd5b50610177610ae6565b34801561029a57600080fd5b50610177610aec565b3480156102af57600080fd5b5061022d6102be366004611a0d565b610af2565b3480156102cf57600080fd5b50610177610d0a565b3480156102e457600080fd5b5061022d6102f3366004611a0d565b610d10565b34801561030457600080fd5b5061030d610f27565b6040516101849190611b0d565b34801561032657600080fd5b50610177610335366004611a84565b610f30565b34801561034657600080fd5b50610177610f54565b34801561035b57600080fd5b506101b761036a366004611a5e565b610f5a565b34801561037b57600080fd5b50610177611098565b34801561039057600080fd5b506101b761109e565b3480156103a557600080fd5b506101b76110e4565b60005481565b600a54600160a060020a031681565b33600180600c5460ff1660068111156103d857fe5b146103e257600080fd5b600b54600a4204106103f357600080fd5b33600160a060020a0383161461040857600080fd5b6002543410156104515760405133903480156108fc02916000818181858888f1935050505015801561043e573d6000803e3d6000fd5b50600c805460ff19166006179055610478565b600a805473ffffffffffffffffffffffffffffffffffffffff19163317905561047861109e565b5050565b60045481565b600a54600160a060020a0316600280600c5460ff1660068111156104a257fe5b146104ac57600080fd5b600b54600a4204106104bd57600080fd5b33600160a060020a038316146104d257600080fd5b600b54600a4204116104e357600080fd5b600a54600254604051600160a060020a039092169181156108fc0291906000818181858888f1935050505015801561051f573d6000803e3d6000fd5b5050600c805460ff1916600517905550565b60055481565b6009546101009004600160a060020a031681565b60008060008060606000606060008060608c518b511415156105cf57600c80546006919060ff19166001835b0217905550600954600254604051610100909204600160a060020a0316916108fc82150291906000818181858888f193505050501580156105bc573d6000803e3d6000fd5b5060009950505050505050505050610aca565b60005b8b518110156106a1578d818151811015156105e957fe5b90602001906020020151600081518110151561060157fe5b9060200190602002015161062b8d8381518110151561061c57fe5b90602001906020020151611193565b1461069957600c80546006919060ff19166001835b0217905550600954600254604051610100909204600160a060020a0316916108fc82150291906000818181858888f19350505050158015610685573d6000803e3d6000fd5b5060009a5050505050505050505050610aca565b6001016105d2565b508b60008151811015156106b157fe5b6020908102909101015160001a60f860020a0260f860020a900460ff1697506106f28d60008151811015156106e257fe5b9060200190602002015189610af2565b151561070c57600c80546006919060ff1916600183610577565b6107168c89610d10565b151561073057600c80546006919060ff1916600183610577565b8b600081518110151561073f57fe5b6020908102909101015160011a60f860020a0260f860020a900498508a600081518110151561076a57fe5b90602001906020020151516040519080825280601f01601f19166020018201604052801561079f576020820181803883390190505b5090506107c68b60008151811015156107b457fe5b90602001906020020151600454611232565b905060025b600854600201811015610829576007548d60008151811015156107ea57fe5b906020019060200201518260208110151561080157fe5b1a60f860020a0260f860020a900460ff16141561081d57610829565b600195860195016107cb565b508460ff16604051908082528060200260200182016040528015610857578160200160208202803883390190505b50955060005b8560ff168110156108c6578c600081518110151561087757fe5b906020019060200201518160020160208110151561089157fe5b1a60f860020a0260f860020a900487828151811015156108ad57fe5b60ff90921660209283029091019091015260010161085d565b508c516000190160ff8616146108e85760009950505050505050505050610aca565b60018d510360405190808252806020026020018201604052801561092057816020015b606081526020019060019003908161090b5790505b50935060005b8651811015610a1b5761096d8e8260010181518110151561094357fe5b90602001906020020151888381518110151561095b57fe5b9060200190602002015160ff16610af2565b151561098757600c80546006919060ff1916600183610640565b8b8160010181518110151561099857fe5b90602001906020020151516040519080825280601f01601f1916602001820160405280156109cd576020820181803883390190505b5085828151811015156109dc57fe5b602090810290910101528b516109fb908d90600184019081106107b457fe5b8582815181101515610a0957fe5b60209081029091010152600101610926565b50610a298960ff1685611329565b9250610a3481610f30565b9150828214610a6c5750506009805460ff199081166001908117909255600c805490911660051790559750610aca9650505050505050565b600c8054600660ff19909116179055600954600254604051610100909204600160a060020a0316916108fc82150291906000818181858888f19350505050158015610abb573d6000803e3d6000fd5b50600099505050505050505050505b9392505050565b60065481565b600b5481565b60095460ff1681565b60015481565b60025481565b600080836000815181101515610b0457fe5b6020908102909101015160408051818152606081810183529293506000928392909190602082018180388339019050509050600091505b600160065403821015610cf857600282900a8681160460011415610c255760005b6020811015610bc2578783600101815181101515610b7657fe5b9060200190602002015181602081101515610b8d57fe5b1a60f860020a028282815181101515610ba257fe5b906020010190600160f860020a031916908160001a905350600101610b5c565b5060005b6020811015610c1357848160208110610bdb57fe5b1a60f860020a028282602001815181101515610bf357fe5b906020010190600160f860020a031916908160001a905350600101610bc6565b50805160208201209350839250610ced565b60005b6020811015610c7257848160208110610c3d57fe5b1a60f860020a028282815181101515610c5257fe5b906020010190600160f860020a031916908160001a905350600101610c28565b5060005b6020811015610cdf578783600101815181101515610c9057fe5b9060200190602002015181602081101515610ca757fe5b1a60f860020a028282602001815181101515610cbf57fe5b906020010190600160f860020a031916908160001a905350600101610c76565b508051602082012093508392505b600190910190610b3b565b60005484149450505050505b92915050565b60085481565b600080836000815181101515610d2257fe5b6020908102909101015160408051818152606081810183529293506000928392909190602082018180388339019050509050600091505b600160065403821015610f1657600282900a8681160460011415610e435760005b6020811015610de0578783600101815181101515610d9457fe5b9060200190602002015181602081101515610dab57fe5b1a60f860020a028282815181101515610dc057fe5b906020010190600160f860020a031916908160001a905350600101610d7a565b5060005b6020811015610e3157848160208110610df957fe5b1a60f860020a028282602001815181101515610e1157fe5b906020010190600160f860020a031916908160001a905350600101610de4565b50805160208201209350839250610f0b565b60005b6020811015610e9057848160208110610e5b57fe5b1a60f860020a028282815181101515610e7057fe5b906020010190600160f860020a031916908160001a905350600101610e46565b5060005b6020811015610efd578783600101815181101515610eae57fe5b9060200190602002015181602081101515610ec557fe5b1a60f860020a028282602001815181101515610edd57fe5b906020010190600160f860020a031916908160001a905350600101610e94565b508051602082012093508392505b600190910190610d59565b600154841494505050505092915050565b600c5460ff1681565b805160009082901515610f47575060009050610f4f565b505060208101515b919050565b60035481565b6009546101009004600160a060020a0316600280600c5460ff166006811115610f7f57fe5b14610f8957600080fd5b600b54600a420410610f9a57600080fd5b33600160a060020a03831614610faf57600080fd5b6040805160208082528183019092526000916060919060208201818038833901905050905060005b602081101561102157858160208110610fec57fe5b1a60f860020a02828281518110151561100157fe5b906020010190600160f860020a031916908160001a905350600101610fd7565b5080516020820120600354909250821461108457600a54600254604051600160a060020a039092169181156108fc0291906000818181858888f19350505050158015611071573d6000803e3d6000fd5b50600c805460ff19166005179055611091565b600485905561109161109e565b5050505050565b60075481565b600c5460ff1660068111156110af57fe5b60010160068111156110bd57fe5b600c805460ff191660018360068111156110d357fe5b0217905550600a4204600a01600b55565b6009546101009004600160a060020a0316600380600c5460ff16600681111561110957fe5b1461111357600080fd5b600b54600a42041061112457600080fd5b33600160a060020a0383161461113957600080fd5b600b54600a42041161114a57600080fd5b600954600254604051610100909204600160a060020a0316916108fc82150291906000818181858888f1935050505015801561118a573d6000803e3d6000fd5b5061047861109e565b600080606083516040519080825280601f01601f1916602001820160405280156111c4576020820181803883390190505b50905060005b84518110156112225784818151811015156111e157fe5b90602001015160f860020a900460f860020a02828281518110151561120257fe5b906020010190600160f860020a031916908160001a9053506001016111ca565b5080516020909101209392505050565b606080606084516040519080825280601f01601f191660200182016040528015611263576020820181803883390190505b5090506112718486516114c4565b905084516040519080825280601f01601f19166020018201604052801561129f576020820181803883390190505b50915060005b855181101561131f5785818151811015156112bc57fe5b90602001015160f860020a900460f860020a0282828151811015156112dd57fe5b90602001015160f860020a900460f860020a021883828151811015156112ff57fe5b906020010190600160f860020a031916908160001a9053506001016112a5565b5090949350505050565b6000808360021415610aca57606083600081518110151561134657fe5b90602001906020020151516002026040519080825280601f01601f19166020018201604052801561137e576020820181803883390190505b50905060005b84600081518110151561139357fe5b906020019060200201515181101561140c578460008151811015156113b457fe5b90602001906020020151818151811015156113cb57fe5b90602001015160f860020a900460f860020a0282828151811015156113ec57fe5b906020010190600160f860020a031916908160001a905350600101611384565b5060005b84600181518110151561141f57fe5b90602001906020020151518110156114b35784600181518110151561144057fe5b906020019060200201518181518110151561145757fe5b90602001015160f860020a900460f860020a028286600081518110151561147a57fe5b9060200190602002015151830181518110151561149357fe5b906020010190600160f860020a031916908160001a905350600101611410565b508051602090910120949350505050565b60608060606000846040519080825280601f01601f1916602001820160405280156114f6576020820181803883390190505b5060408051602080825281830190925291945060208201818038833901905050915060005b6020860481101561176a5780151561165b57604080516021808252606082810190935260208201818038833901905050905060005b602081101561159a5788816020811061156557fe5b1a60f860020a02828281518110151561157a57fe5b906020010190600160f860020a031916908160001a905350600101611550565b508051600090829060209081106115ad57fe5b906020010190600160f860020a031916908160001a90535080516020820120925060005b6020811015611654578381602081106115e657fe5b1a60f860020a0285828151811015156115fb57fe5b906020010190600160f860020a031916908160001a90535083816020811061161f57fe5b1a60f860020a02868281518110151561163457fe5b906020010190600160f860020a031916908160001a9053506001016115d1565b5050611762565b604080516020808252818301909252606094509060208201818038833901905050925060005b60208110156116cb5782816020811061169657fe5b1a60f860020a0284828151811015156116ab57fe5b906020010190600160f860020a031916908160001a905350600101611681565b5082516020840120915060005b6020811015611760578281602081106116ed57fe5b1a60f860020a02848281518110151561170257fe5b906020010190600160f860020a031916908160001a90535082816020811061172657fe5b1a60f860020a028582846020020181518110151561174057fe5b906020010190600160f860020a031916908160001a9053506001016116d8565b505b60010161151b565b509195945050505050565b6000601f8201831361178657600080fd5b813561179961179482611b42565b611b1b565b81815260209384019390925082018360005b838110156117d757813586016117c188826117e1565b84525060209283019291909101906001016117ab565b5050505092915050565b6000601f820183136117f257600080fd5b813561180061179482611b42565b9150818183526020840193506020810190508385602084028201111561182557600080fd5b60005b838110156117d7578161183b888261191e565b8452506020928301929190910190600101611828565b6000601f8201831361186257600080fd5b813561187061179482611b42565b9150818183526020840193506020810190508385602084028201111561189557600080fd5b60005b838110156117d757816118ab888261191e565b8452506020928301929190910190600101611898565b6000601f820183136118d257600080fd5b81356118e061179482611b42565b81815260209384019390925082018360005b838110156117d75781358601611908888261192a565b84525060209283019291909101906001016118f2565b6000610aca8235611b9b565b6000601f8201831361193b57600080fd5b813561194961179482611b63565b9150808252602083016020830185838301111561196557600080fd5b611970838284611bc5565b50505092915050565b60008060006060848603121561198e57600080fd5b833567ffffffffffffffff8111156119a557600080fd5b6119b186828701611775565b935050602084013567ffffffffffffffff8111156119ce57600080fd5b6119da86828701611851565b925050604084013567ffffffffffffffff8111156119f757600080fd5b611a03868287016118c1565b9150509250925092565b60008060408385031215611a2057600080fd5b823567ffffffffffffffff811115611a3757600080fd5b611a4385828601611851565b9250506020611a548582860161191e565b9150509250929050565b600060208284031215611a7057600080fd5b6000611a7c848461191e565b949350505050565b600060208284031215611a9657600080fd5b813567ffffffffffffffff811115611aad57600080fd5b611a7c8482850161192a565b611ac281611b8b565b82525050565b611ac281611b96565b611ac281611b9b565b611ac281611bba565b60208101610d048284611ab9565b60208101610d048284611ac8565b60208101610d048284611ad1565b60208101610d048284611ada565b60405181810167ffffffffffffffff81118282101715611b3a57600080fd5b604052919050565b600067ffffffffffffffff821115611b5957600080fd5b5060209081020190565b600067ffffffffffffffff821115611b7a57600080fd5b506020601f91909101601f19160190565b6000610d0482611bae565b151590565b90565b600060078210611baa57fe5b5090565b600160a060020a031690565b6000610d0482611b9e565b8281833750600091015256fea265627a7a723058204c16715b50a1071e8f1477648f47a2c1b549d9217fc3bf69e93a7bdb2ee9a5fc6c6578706572696d656e74616cf50037`

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

// Mverify is a free data retrieval call binding the contract method 0xa2630372.
//
// Solidity: function Mverify(bytes32[] complaint, uint256 _index) constant returns(bool)
func (_Fairswap *FairswapCaller) Mverify(opts *bind.CallOpts, complaint [][32]byte, _index *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Fairswap.contract.Call(opts, out, "Mverify", complaint, _index)
	return *ret0, err
}

// Mverify is a free data retrieval call binding the contract method 0xa2630372.
//
// Solidity: function Mverify(bytes32[] complaint, uint256 _index) constant returns(bool)
func (_Fairswap *FairswapSession) Mverify(complaint [][32]byte, _index *big.Int) (bool, error) {
	return _Fairswap.Contract.Mverify(&_Fairswap.CallOpts, complaint, _index)
}

// Mverify is a free data retrieval call binding the contract method 0xa2630372.
//
// Solidity: function Mverify(bytes32[] complaint, uint256 _index) constant returns(bool)
func (_Fairswap *FairswapCallerSession) Mverify(complaint [][32]byte, _index *big.Int) (bool, error) {
	return _Fairswap.Contract.Mverify(&_Fairswap.CallOpts, complaint, _index)
}

// MverifyForGate is a free data retrieval call binding the contract method 0xafa31a42.
//
// Solidity: function MverifyForGate(bytes32[] gateProof, uint256 _index) constant returns(bool)
func (_Fairswap *FairswapCaller) MverifyForGate(opts *bind.CallOpts, gateProof [][32]byte, _index *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Fairswap.contract.Call(opts, out, "MverifyForGate", gateProof, _index)
	return *ret0, err
}

// MverifyForGate is a free data retrieval call binding the contract method 0xafa31a42.
//
// Solidity: function MverifyForGate(bytes32[] gateProof, uint256 _index) constant returns(bool)
func (_Fairswap *FairswapSession) MverifyForGate(gateProof [][32]byte, _index *big.Int) (bool, error) {
	return _Fairswap.Contract.MverifyForGate(&_Fairswap.CallOpts, gateProof, _index)
}

// MverifyForGate is a free data retrieval call binding the contract method 0xafa31a42.
//
// Solidity: function MverifyForGate(bytes32[] gateProof, uint256 _index) constant returns(bool)
func (_Fairswap *FairswapCallerSession) MverifyForGate(gateProof [][32]byte, _index *big.Int) (bool, error) {
	return _Fairswap.Contract.MverifyForGate(&_Fairswap.CallOpts, gateProof, _index)
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

// ReceiverGetEther is a paid mutator transaction binding the contract method 0x3e37aeb6.
//
// Solidity: function receiverGetEther() returns()
func (_Fairswap *FairswapTransactor) ReceiverGetEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fairswap.contract.Transact(opts, "receiverGetEther")
}

// ReceiverGetEther is a paid mutator transaction binding the contract method 0x3e37aeb6.
//
// Solidity: function receiverGetEther() returns()
func (_Fairswap *FairswapSession) ReceiverGetEther() (*types.Transaction, error) {
	return _Fairswap.Contract.ReceiverGetEther(&_Fairswap.TransactOpts)
}

// ReceiverGetEther is a paid mutator transaction binding the contract method 0x3e37aeb6.
//
// Solidity: function receiverGetEther() returns()
func (_Fairswap *FairswapTransactorSession) ReceiverGetEther() (*types.Transaction, error) {
	return _Fairswap.Contract.ReceiverGetEther(&_Fairswap.TransactOpts)
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

// SenderGetEther is a paid mutator transaction binding the contract method 0xf6fa345f.
//
// Solidity: function senderGetEther() returns()
func (_Fairswap *FairswapTransactor) SenderGetEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fairswap.contract.Transact(opts, "senderGetEther")
}

// SenderGetEther is a paid mutator transaction binding the contract method 0xf6fa345f.
//
// Solidity: function senderGetEther() returns()
func (_Fairswap *FairswapSession) SenderGetEther() (*types.Transaction, error) {
	return _Fairswap.Contract.SenderGetEther(&_Fairswap.TransactOpts)
}

// SenderGetEther is a paid mutator transaction binding the contract method 0xf6fa345f.
//
// Solidity: function senderGetEther() returns()
func (_Fairswap *FairswapTransactorSession) SenderGetEther() (*types.Transaction, error) {
	return _Fairswap.Contract.SenderGetEther(&_Fairswap.TransactOpts)
}
