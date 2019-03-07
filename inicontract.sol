pragma solidity >0.4.0 <=0.6.0;
pragma experimental ABIEncoderV2;

contract fairswap{

	bytes32 public CircuitMerkleRoot;
	bytes32 public EncInputRoot;
	uint public price;
	bytes32 public commitmentofkey;
	bytes32 public key;
	/*bytes internal keyToBeHashed;
	bytes public _keybyte;
	bytes32 public _keybyteHash;
	uint public complaintlength;
	uint public index;*/
	bool public complianBool;

	mapping(uint => uint) public circuitGatesOperation;

	address public senderaddrs;
	address public recieveraddrs;

  uint public timeout;

	enum stage{created,initialized,accepted,keyRevealed,finished}

	stage public phase = stage.created;

	function nextStage() public{
		phase=stage(uint(phase)+1);
		timeout=now+3 minutes;
	}

	constructor(uint _price,bytes32 _commitmentofkey,bytes32  _CircuitMerkleRoot,bytes32  _EncInputRoot,uint[] memory	 circuitGatesOperationArray) public{

		senderaddrs=msg.sender;
		price = _price;
		commitmentofkey = _commitmentofkey;
		CircuitMerkleRoot = _CircuitMerkleRoot;
		EncInputRoot = _EncInputRoot;
		for(uint i=0;i<circuitGatesOperationArray.length;i++)
		{
			circuitGatesOperation[i]=circuitGatesOperationArray[i];
		}
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

	function revealKey (bytes32  _key) allowed(senderaddrs, stage.accepted) public {
		key=_key;
		nextStage();
	}

 	function complain(bytes32[5][] memory complaint,uint[] memory indices) public returns(bool)
	{

		bytes32 operationValue;
		uint operation;
		bytes32[] memory inputVectorsToTheGate;
		bytes32 Out;
		uint index;
		index=indices[0];
		inputVectorsToTheGate = new bytes32[] (complaint.length-1);
		require(Mverify(complaint[0],indices[0]));
		operation=circuitGatesOperation[indices[0]];
		Out=Dec(complaint[0][0]);

		for(uint i=0;i<complaint.length-1;i++){
			require(Mverify(complaint[i],indices[i]));
			inputVectorsToTheGate[i]=Dec(complaint[i+1][0]);

		}

		operationValue=Operation(operation,inputVectorsToTheGate);
		if(Out==operationValue)
		{
			complianBool=true;
			return true;
		}
		
		complianBool=false;
		return false;


	}

	function Operation(uint operation,bytes32[] memory inputVectorsToTheGate) internal pure returns(bytes32)
	{
		bytes32 Hashed;
			if(operation==2)
			{
				bytes memory toBeHashed;

				toBeHashed=new bytes (64);
				for (uint i=0;i<32;i++)
				{
					toBeHashed[i]=inputVectorsToTheGate[0][i];
				}
				for (uint i=0;i<32;i++)
				{
					toBeHashed[i+32]=inputVectorsToTheGate[1][i];
				}
				Hashed=keccak256(toBeHashed);
			}
			return Hashed;
	}

	function  Mverify(bytes32[5] memory complaint,uint _index)  public view returns(bool)
	{
		bytes32 _value;
		_value=complaint[0];
		bytes32 Hashed;
		uint i ;
		uint depth=5;
		bytes memory toBeHashed;
		toBeHashed=new bytes (64);

		for(i=0; i<depth-1; i++)
		{
			if ((_index&(1<<i))>>i == 1)
			{

				for (uint j=0;j<32;j++)
				{
					toBeHashed[j]=complaint[i+1][j];
				}
				for (uint j=0;j<32;j++)
				{
					toBeHashed[j+32]=_value[j];
				}


				Hashed=keccak256(toBeHashed);
				_value=Hashed;
			}
			else
			{
				for (uint j=0;j<32;j++)
				{
					toBeHashed[j]=_value[j];
				}
				for (uint j=0;j<32;j++)
				{
					toBeHashed[j+32]=complaint[i+1][j];
				}

				Hashed=keccak256(toBeHashed);
				_value=Hashed;
			}
		}
		return (_value==EncInputRoot);
	}


	function Dec(bytes32 encryptedtext) internal view returns(bytes32)
	{
		bytes memory toBeHashedKey;
		bytes32 keyHashed;
		bytes32 decodedtext;

		toBeHashedKey=new bytes (33);
		for(uint i=0;i<32;i++){

			toBeHashedKey[i]=key[i];

		}
		toBeHashedKey[32]=0;
		keyHashed=keccak256(toBeHashedKey);
		decodedtext=keyHashed^encryptedtext;

		return decodedtext;
	}

}

