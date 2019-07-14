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
const FairswapABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"EncInputRoot\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"receiverAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"accept\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"key\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"receiverGetEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Now\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"senderAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"complaint\",\"type\":\"bytes32[][]\"},{\"name\":\"gateMerkleProof\",\"type\":\"bytes32[]\"},{\"name\":\"encodedVectors_ThisGate\",\"type\":\"bytes[]\"}],\"name\":\"complain\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"depth\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"timeout\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"validComplain\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EncCircuitRoot\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recEntryKey\",\"type\":\"bytes32\"}],\"name\":\"initializeRecieverAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"complaint\",\"type\":\"bytes32[]\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"Mverify\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxInputLinesToAGate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"gateProof\",\"type\":\"bytes32[]\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"MverifyForGate\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"phase\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"source\",\"type\":\"bytes\"}],\"name\":\"bytesToBytes32\",\"outputs\":[{\"name\":\"result\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"commitmentOfKey\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"revealKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalNumberOfGates\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"receiverEntryKey\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"nextStage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"senderGetEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_commitmentOfKey\",\"type\":\"bytes32\"},{\"name\":\"_EncInputRoot\",\"type\":\"bytes32\"},{\"name\":\"_EncCircuitRoot\",\"type\":\"bytes32\"},{\"name\":\"_receiverEntryKey\",\"type\":\"bytes32\"},{\"name\":\"_totalNumberOfGates\",\"type\":\"uint256\"},{\"name\":\"_maxInputLinesToAGate\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// FairswapBin is the compiled bytecode used for deploying new contracts.
const FairswapBin = `0x6080604052600d805460ff191690553480156200001b57600080fd5b5060405160e0806200205883398101806040526200003d919081019062000312565b600a805461010060a860020a031916336101000217905560028790556003869055600085905560018490556005839055600882905560098190556200008b82640100000000620000b1810204565b600101600755620000a4640100000000620002b4810204565b50505050505050620003c0565b604080517ff8f9cbfae6cc78fbefe7cdc3a1793dfcf4f0e8bbd8cec470b6a28a7a5a3e1efd81527ff5ecf1b3e9debc68e1d9cfabc5997135bfb7a7a3938b7b606b5b4b3f2f1f0ffe60208201527ff6e4ed9ff2d6b458eadcdf97bd91692de2d4da8fd2d0ac50c6ae9a8272523616818301527fc8c0b887b0a8a4489c948c7f847c6125746c645c544c444038302820181008ff60608201527ff7cae577eec2a03cf3bad76fb589591debb2dd67e0aa9834bea6925f6a4a2e0e60808201527fe39ed557db96902cd38ed14fad815115c786af479b7e8324736353433727170760a08201527fc976c13bb96e881cb166a933a55e490d9d56952b8d4e801485467d236242260660c08201527f753a6d1b65325d0c552a4d1345224105391a310b29122104190a11030902010060e08201526101008082019092527f0100000000000000000000000000000000000000000000000000000000000000600170010000000000000000000000000000000068010000000000000000640100000000620100006010600460026000198c019081041790810417908104178881041790810417908104179081041790810417017e818283848586878898a8b8c8d8e8f929395969799a9b9d9e9faaeb6bedeeff0281900460ff0390910151047f8000000000000000000000000000000000000000000000000000000000000000909211020190565b600d5460ff166006811115620002c657fe5b6001016006811115620002d557fe5b600d805460ff19166001836006811115620002ec57fe5b0217905550600a4204600a01600c55565b60006200030b8251620003bd565b9392505050565b600080600080600080600060e0888a0312156200032e57600080fd5b60006200033c8a8a620002fd565b97505060206200034f8a828b01620002fd565b9650506040620003628a828b01620002fd565b9550506060620003758a828b01620002fd565b9450506080620003888a828b01620002fd565b93505060a06200039b8a828b01620002fd565b92505060c0620003ae8a828b01620002fd565b91505092959891949750929550565b90565b611c8880620003d06000396000f3fe608060405260043610610183576000357c0100000000000000000000000000000000000000000000000000000000900480639ab164e5116100ee578063bfe370d9116100a7578063e44b23c811610081578063e44b23c8146103b5578063eaa94830146103ca578063ee3743ab146103df578063f6fa345f146103f457610183565b8063bfe370d914610360578063ca641f8614610380578063d65472361461039557610183565b80639ab164e5146102b4578063a035b1fe146102d4578063a2630372146102e9578063aa9b8ac714610309578063afa31a421461031e578063b1c9fe6e1461033e57610183565b80634fc84791116101405780634fc847911461021e5780635c57b30814610233578063631c56ef1461026057806370dea79a1461027557806376aa657f1461028a5780637a5145941461029f57610183565b8063062a13981461018857806316fed3e2146101b35780632852b71c146101d55780633943380c146101df5780633e37aeb6146101f457806344d4fd1914610209575b600080fd5b34801561019457600080fd5b5061019d610409565b6040516101aa9190611b7c565b60405180910390f35b3480156101bf57600080fd5b506101c861040f565b6040516101aa9190611b60565b6101dd61041e565b005b3480156101eb57600080fd5b5061019d6104c3565b34801561020057600080fd5b506101dd6104c9565b34801561021557600080fd5b5061019d610578565b34801561022a57600080fd5b506101c861057e565b34801561023f57600080fd5b5061025361024e3660046119f6565b610592565b6040516101aa9190611b6e565b34801561026c57600080fd5b5061019d610b18565b34801561028157600080fd5b5061019d610b1e565b34801561029657600080fd5b50610253610b24565b3480156102ab57600080fd5b5061019d610b2d565b3480156102c057600080fd5b506101dd6102cf366004611adb565b610b33565b3480156102e057600080fd5b5061019d610b63565b3480156102f557600080fd5b50610253610304366004611a8a565b610b69565b34801561031557600080fd5b5061019d610d81565b34801561032a57600080fd5b50610253610339366004611a8a565b610d87565b34801561034a57600080fd5b50610353610f9e565b6040516101aa9190611b8a565b34801561036c57600080fd5b5061019d61037b366004611b01565b610fa7565b34801561038c57600080fd5b5061019d610fcb565b3480156103a157600080fd5b506101dd6103b0366004611adb565b610fd1565b3480156103c157600080fd5b5061019d61110f565b3480156103d657600080fd5b5061019d611115565b3480156103eb57600080fd5b506101dd61111b565b34801561040057600080fd5b506101dd611161565b60005481565b600b54600160a060020a031681565b600b54600160a060020a0316600180600d5460ff16600681111561043e57fe5b1461044857600080fd5b600c54600a42041061045957600080fd5b33600160a060020a0383161461046e57600080fd5b6002543410156104b75760405133903480156108fc02916000818181858888f193505050501580156104a4573d6000803e3d6000fd5b50600d805460ff191660061790556104bf565b6104bf61111b565b5050565b60045481565b600b54600160a060020a0316600280600d5460ff1660068111156104e957fe5b146104f357600080fd5b600c54600a42041061050457600080fd5b33600160a060020a0383161461051957600080fd5b600c54600a42041161052a57600080fd5b600b54600254604051600160a060020a039092169181156108fc0291906000818181858888f19350505050158015610566573d6000803e3d6000fd5b5050600d805460ff1916600517905550565b60065481565b600a546101009004600160a060020a031681565b60008060008060606000606060008060608c518b5114151561061657600d80546006919060ff19166001835b0217905550600a54600254604051610100909204600160a060020a0316916108fc82150291906000818181858888f19350505050158015610603573d6000803e3d6000fd5b5060009950505050505050505050610b11565b60005b8b518110156106e8578d8181518110151561063057fe5b90602001906020020151600081518110151561064857fe5b906020019060200201516106728d8381518110151561066357fe5b90602001906020020151611210565b146106e057600d80546006919060ff19166001835b0217905550600a54600254604051610100909204600160a060020a0316916108fc82150291906000818181858888f193505050501580156106cc573d6000803e3d6000fd5b5060009a5050505050505050505050610b11565b600101610619565b508b60008151811015156106f857fe5b6020908102909101015160001a60f860020a0260f860020a900460ff1697506107398d600081518110151561072957fe5b9060200190602002015189610b69565b151561075357600d80546006919060ff19166001836105be565b61075d8c89610d87565b151561077757600d80546006919060ff19166001836105be565b8b600081518110151561078657fe5b6020908102909101015160011a60f860020a0260f860020a900498508a60008151811015156107b157fe5b90602001906020020151516040519080825280601f01601f1916602001820160405280156107e6576020820181803883390190505b50905061080d8b60008151811015156107fb57fe5b906020019060200201516004546112af565b905060025b600954600201811015610870576008548d600081518110151561083157fe5b906020019060200201518260208110151561084857fe5b1a60f860020a0260f860020a900460ff16141561086457610870565b60019586019501610812565b508460ff1660405190808252806020026020018201604052801561089e578160200160208202803883390190505b50955060005b8560ff1681101561090d578c60008151811015156108be57fe5b90602001906020020151816002016020811015156108d857fe5b1a60f860020a0260f860020a900487828151811015156108f457fe5b60ff9092166020928302909101909101526001016108a4565b508c516000190160ff86161461092f5760009950505050505050505050610b11565b60018d510360405190808252806020026020018201604052801561096757816020015b60608152602001906001900390816109525790505b50935060005b8651811015610a62576109b48e8260010181518110151561098a57fe5b9060200190602002015188838151811015156109a257fe5b9060200190602002015160ff16610b69565b15156109ce57600d80546006919060ff1916600183610687565b8b816001018151811015156109df57fe5b90602001906020020151516040519080825280601f01601f191660200182016040528015610a14576020820181803883390190505b508582815181101515610a2357fe5b602090810290910101528b51610a42908d90600184019081106107fb57fe5b8582815181101515610a5057fe5b6020908102909101015260010161096d565b50610a708960ff16856113a6565b9250610a7b81610fa7565b9150828214610ab3575050600a805460ff199081166001908117909255600d805490911660051790559750610b119650505050505050565b600d8054600660ff19909116179055600a54600254604051610100909204600160a060020a0316916108fc82150291906000818181858888f19350505050158015610b02573d6000803e3d6000fd5b50600099505050505050505050505b9392505050565b60075481565b600c5481565b600a5460ff1681565b60015481565b6005548114610b4157600080fd5b50600b805473ffffffffffffffffffffffffffffffffffffffff191633179055565b60025481565b600080836000815181101515610b7b57fe5b6020908102909101015160408051818152606081810183529293506000928392909190602082018180388339019050509050600091505b600160075403821015610d6f57600282900a8681160460011415610c9c5760005b6020811015610c39578783600101815181101515610bed57fe5b9060200190602002015181602081101515610c0457fe5b1a60f860020a028282815181101515610c1957fe5b906020010190600160f860020a031916908160001a905350600101610bd3565b5060005b6020811015610c8a57848160208110610c5257fe5b1a60f860020a028282602001815181101515610c6a57fe5b906020010190600160f860020a031916908160001a905350600101610c3d565b50805160208201209350839250610d64565b60005b6020811015610ce957848160208110610cb457fe5b1a60f860020a028282815181101515610cc957fe5b906020010190600160f860020a031916908160001a905350600101610c9f565b5060005b6020811015610d56578783600101815181101515610d0757fe5b9060200190602002015181602081101515610d1e57fe5b1a60f860020a028282602001815181101515610d3657fe5b906020010190600160f860020a031916908160001a905350600101610ced565b508051602082012093508392505b600190910190610bb2565b60005484149450505050505b92915050565b60095481565b600080836000815181101515610d9957fe5b6020908102909101015160408051818152606081810183529293506000928392909190602082018180388339019050509050600091505b600160075403821015610f8d57600282900a8681160460011415610eba5760005b6020811015610e57578783600101815181101515610e0b57fe5b9060200190602002015181602081101515610e2257fe5b1a60f860020a028282815181101515610e3757fe5b906020010190600160f860020a031916908160001a905350600101610df1565b5060005b6020811015610ea857848160208110610e7057fe5b1a60f860020a028282602001815181101515610e8857fe5b906020010190600160f860020a031916908160001a905350600101610e5b565b50805160208201209350839250610f82565b60005b6020811015610f0757848160208110610ed257fe5b1a60f860020a028282815181101515610ee757fe5b906020010190600160f860020a031916908160001a905350600101610ebd565b5060005b6020811015610f74578783600101815181101515610f2557fe5b9060200190602002015181602081101515610f3c57fe5b1a60f860020a028282602001815181101515610f5457fe5b906020010190600160f860020a031916908160001a905350600101610f0b565b508051602082012093508392505b600190910190610dd0565b600154841494505050505092915050565b600d5460ff1681565b805160009082901515610fbe575060009050610fc6565b505060208101515b919050565b60035481565b600a546101009004600160a060020a0316600280600d5460ff166006811115610ff657fe5b1461100057600080fd5b600c54600a42041061101157600080fd5b33600160a060020a0383161461102657600080fd5b6040805160208082528183019092526000916060919060208201818038833901905050905060005b60208110156110985785816020811061106357fe5b1a60f860020a02828281518110151561107857fe5b906020010190600160f860020a031916908160001a90535060010161104e565b508051602082012060035490925082146110fb57600b54600254604051600160a060020a039092169181156108fc0291906000818181858888f193505050501580156110e8573d6000803e3d6000fd5b50600d805460ff19166005179055611108565b600485905561110861111b565b5050505050565b60085481565b60055481565b600d5460ff16600681111561112c57fe5b600101600681111561113a57fe5b600d805460ff1916600183600681111561115057fe5b0217905550600a4204600a01600c55565b600a546101009004600160a060020a0316600380600d5460ff16600681111561118657fe5b1461119057600080fd5b600c54600a4204106111a157600080fd5b33600160a060020a038316146111b657600080fd5b600c54600a4204116111c757600080fd5b600a54600254604051610100909204600160a060020a0316916108fc82150291906000818181858888f19350505050158015611207573d6000803e3d6000fd5b506104bf61111b565b600080606083516040519080825280601f01601f191660200182016040528015611241576020820181803883390190505b50905060005b845181101561129f57848181518110151561125e57fe5b90602001015160f860020a900460f860020a02828281518110151561127f57fe5b906020010190600160f860020a031916908160001a905350600101611247565b5080516020909101209392505050565b606080606084516040519080825280601f01601f1916602001820160405280156112e0576020820181803883390190505b5090506112ee848651611541565b905084516040519080825280601f01601f19166020018201604052801561131c576020820181803883390190505b50915060005b855181101561139c57858181518110151561133957fe5b90602001015160f860020a900460f860020a02828281518110151561135a57fe5b90602001015160f860020a900460f860020a0218838281518110151561137c57fe5b906020010190600160f860020a031916908160001a905350600101611322565b5090949350505050565b6000808360021415610b115760608360008151811015156113c357fe5b90602001906020020151516002026040519080825280601f01601f1916602001820160405280156113fb576020820181803883390190505b50905060005b84600081518110151561141057fe5b90602001906020020151518110156114895784600081518110151561143157fe5b906020019060200201518181518110151561144857fe5b90602001015160f860020a900460f860020a02828281518110151561146957fe5b906020010190600160f860020a031916908160001a905350600101611401565b5060005b84600181518110151561149c57fe5b9060200190602002015151811015611530578460018151811015156114bd57fe5b90602001906020020151818151811015156114d457fe5b90602001015160f860020a900460f860020a02828660008151811015156114f757fe5b9060200190602002015151830181518110151561151057fe5b906020010190600160f860020a031916908160001a90535060010161148d565b508051602090910120949350505050565b60608060606000846040519080825280601f01601f191660200182016040528015611573576020820181803883390190505b5060408051602080825281830190925291945060208201818038833901905050915060005b602086048110156117e7578015156116d857604080516021808252606082810190935260208201818038833901905050905060005b6020811015611617578881602081106115e257fe5b1a60f860020a0282828151811015156115f757fe5b906020010190600160f860020a031916908160001a9053506001016115cd565b5080516000908290602090811061162a57fe5b906020010190600160f860020a031916908160001a90535080516020820120925060005b60208110156116d15783816020811061166357fe5b1a60f860020a02858281518110151561167857fe5b906020010190600160f860020a031916908160001a90535083816020811061169c57fe5b1a60f860020a0286828151811015156116b157fe5b906020010190600160f860020a031916908160001a90535060010161164e565b50506117df565b604080516020808252818301909252606094509060208201818038833901905050925060005b60208110156117485782816020811061171357fe5b1a60f860020a02848281518110151561172857fe5b906020010190600160f860020a031916908160001a9053506001016116fe565b5082516020840120915060005b60208110156117dd5782816020811061176a57fe5b1a60f860020a02848281518110151561177f57fe5b906020010190600160f860020a031916908160001a9053508281602081106117a357fe5b1a60f860020a02858284602002018151811015156117bd57fe5b906020010190600160f860020a031916908160001a905350600101611755565b505b600101611598565b509195945050505050565b6000601f8201831361180357600080fd5b813561181661181182611bbf565b611b98565b81815260209384019390925082018360005b83811015611854578135860161183e888261185e565b8452506020928301929190910190600101611828565b5050505092915050565b6000601f8201831361186f57600080fd5b813561187d61181182611bbf565b915081818352602084019350602081019050838560208402820111156118a257600080fd5b60005b8381101561185457816118b8888261199b565b84525060209283019291909101906001016118a5565b6000601f820183136118df57600080fd5b81356118ed61181182611bbf565b9150818183526020840193506020810190508385602084028201111561191257600080fd5b60005b838110156118545781611928888261199b565b8452506020928301929190910190600101611915565b6000601f8201831361194f57600080fd5b813561195d61181182611bbf565b81815260209384019390925082018360005b83811015611854578135860161198588826119a7565b845250602092830192919091019060010161196f565b6000610b118235611c18565b6000601f820183136119b857600080fd5b81356119c661181182611be0565b915080825260208301602083018583830111156119e257600080fd5b6119ed838284611c42565b50505092915050565b600080600060608486031215611a0b57600080fd5b833567ffffffffffffffff811115611a2257600080fd5b611a2e868287016117f2565b935050602084013567ffffffffffffffff811115611a4b57600080fd5b611a57868287016118ce565b925050604084013567ffffffffffffffff811115611a7457600080fd5b611a808682870161193e565b9150509250925092565b60008060408385031215611a9d57600080fd5b823567ffffffffffffffff811115611ab457600080fd5b611ac0858286016118ce565b9250506020611ad18582860161199b565b9150509250929050565b600060208284031215611aed57600080fd5b6000611af9848461199b565b949350505050565b600060208284031215611b1357600080fd5b813567ffffffffffffffff811115611b2a57600080fd5b611af9848285016119a7565b611b3f81611c08565b82525050565b611b3f81611c13565b611b3f81611c18565b611b3f81611c37565b60208101610d7b8284611b36565b60208101610d7b8284611b45565b60208101610d7b8284611b4e565b60208101610d7b8284611b57565b60405181810167ffffffffffffffff81118282101715611bb757600080fd5b604052919050565b600067ffffffffffffffff821115611bd657600080fd5b5060209081020190565b600067ffffffffffffffff821115611bf757600080fd5b506020601f91909101601f19160190565b6000610d7b82611c2b565b151590565b90565b600060078210611c2757fe5b5090565b600160a060020a031690565b6000610d7b82611c1b565b8281833750600091015256fea265627a7a72305820eac7f45f33f826016baa2c5d4a8f9c6dffc0c57fed57e6763aff96399df24b0c6c6578706572696d656e74616cf50037`

// DeployFairswap deploys a new Ethereum contract, binding an instance of Fairswap to it.
func DeployFairswap(auth *bind.TransactOpts, backend bind.ContractBackend, _price *big.Int, _commitmentOfKey [32]byte, _EncInputRoot [32]byte, _EncCircuitRoot [32]byte, _receiverEntryKey [32]byte, _totalNumberOfGates *big.Int, _maxInputLinesToAGate *big.Int) (common.Address, *types.Transaction, *Fairswap, error) {
	parsed, err := abi.JSON(strings.NewReader(FairswapABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FairswapBin), backend, _price, _commitmentOfKey, _EncInputRoot, _EncCircuitRoot, _receiverEntryKey, _totalNumberOfGates, _maxInputLinesToAGate)
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

// ReceiverEntryKey is a free data retrieval call binding the contract method 0xeaa94830.
//
// Solidity: function receiverEntryKey() constant returns(bytes32)
func (_Fairswap *FairswapCaller) ReceiverEntryKey(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Fairswap.contract.Call(opts, out, "receiverEntryKey")
	return *ret0, err
}

// ReceiverEntryKey is a free data retrieval call binding the contract method 0xeaa94830.
//
// Solidity: function receiverEntryKey() constant returns(bytes32)
func (_Fairswap *FairswapSession) ReceiverEntryKey() ([32]byte, error) {
	return _Fairswap.Contract.ReceiverEntryKey(&_Fairswap.CallOpts)
}

// ReceiverEntryKey is a free data retrieval call binding the contract method 0xeaa94830.
//
// Solidity: function receiverEntryKey() constant returns(bytes32)
func (_Fairswap *FairswapCallerSession) ReceiverEntryKey() ([32]byte, error) {
	return _Fairswap.Contract.ReceiverEntryKey(&_Fairswap.CallOpts)
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

// InitializeRecieverAddress is a paid mutator transaction binding the contract method 0x9ab164e5.
//
// Solidity: function initializeRecieverAddress(bytes32 recEntryKey) returns()
func (_Fairswap *FairswapTransactor) InitializeRecieverAddress(opts *bind.TransactOpts, recEntryKey [32]byte) (*types.Transaction, error) {
	return _Fairswap.contract.Transact(opts, "initializeRecieverAddress", recEntryKey)
}

// InitializeRecieverAddress is a paid mutator transaction binding the contract method 0x9ab164e5.
//
// Solidity: function initializeRecieverAddress(bytes32 recEntryKey) returns()
func (_Fairswap *FairswapSession) InitializeRecieverAddress(recEntryKey [32]byte) (*types.Transaction, error) {
	return _Fairswap.Contract.InitializeRecieverAddress(&_Fairswap.TransactOpts, recEntryKey)
}

// InitializeRecieverAddress is a paid mutator transaction binding the contract method 0x9ab164e5.
//
// Solidity: function initializeRecieverAddress(bytes32 recEntryKey) returns()
func (_Fairswap *FairswapTransactorSession) InitializeRecieverAddress(recEntryKey [32]byte) (*types.Transaction, error) {
	return _Fairswap.Contract.InitializeRecieverAddress(&_Fairswap.TransactOpts, recEntryKey)
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
