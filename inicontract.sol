pragma solidity >0.4.0 <=0.6.0;

contract fairswap{

	bytes32 public CircuitMerkleRoot;
	bytes32 public EncInputRoot;
	uint public price;
	bytes32 public commitmentofkey;

	address public senderaddrs;
	address public recieveraddrs;

  uint public timeout;

	enum stage{created,initialized,accepted,keyRevealed,finished}

	stage public phase = stage.created;

	function nextStage() public{
		phase=stage(uint(phase)+1);
		timeout=now+3 minutes;
	}

	constructor(uint _price,bytes32 _commitmentofkey,bytes32  _CircuitMerkleRoot,bytes32  _EncInputRoot) public{
  	senderaddrs=msg.sender;
		price = _price;
		commitmentofkey = _commitmentofkey;
		CircuitMerkleRoot = _CircuitMerkleRoot;
		EncInputRoot = _EncInputRoot;
		nextStage();
	}


  function initializerecieveraddrs() public{
      recieveraddrs=msg.sender;
  }

	modifier allowed(address _party,stage _s){

		require(phase==_s);
		require(now<timeout);
		require(msg.sender == _party);
		_;

	}

	function accept() allowed(recieveraddrs,stage.initialized) payable public{

		require(msg.value >= price);
		nextStage();
	}

}
