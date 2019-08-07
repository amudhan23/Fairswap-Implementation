pragma solidity >0.4.0 <=0.6.0;
pragma experimental ABIEncoderV2;

contract fairswap{


	bytes32 public EncInputRoot;  //root of encryptedGateOutputs
	bytes32 public EncCircuitRoot;
	uint public price;						//price of the digital good
	bytes32 public commitmentOfKey; // commitment of the key
	bytes32 public key;			//key
	uint public Now; //to get the current time
	uint public depth; //the depth of the merkle tree of encoded inputs
	uint public totalNumberOfGates;
	uint public maxInputLinesToAGate;
	bool public validComplain;

	address payable public  senderAddress;
	address payable public  receiverAddress;

	uint public timeout;

	enum stage{created,initialized,accepted,keyRevealed,finishedAndBothHonestParties,finishedAndMaliciousSender,finishedAndMaliciousReciever}

	stage public phase = stage.created;

	//fn used to move to the next stage fron the current stage
	function nextStage() public{
		phase=stage(uint(phase)+1);
		timeout=now/10+ 10 seconds;
	}

	constructor(uint _price, bytes32 _commitmentOfKey, bytes32  _EncInputRoot, bytes32 _EncCircuitRoot,uint _totalNumberOfGates,uint _maxInputLinesToAGate) public{

		senderAddress=msg.sender;
		price = _price;
		commitmentOfKey = _commitmentOfKey;
		EncInputRoot = _EncInputRoot;
		EncCircuitRoot = _EncCircuitRoot;
		totalNumberOfGates=_totalNumberOfGates;
		maxInputLinesToAGate=_maxInputLinesToAGate;
		depth=log2(totalNumberOfGates)+1;
		nextStage();
	}



//modifier to check the eligibility of party, time and phase for a function execution
modifier allowed(address _party,stage _stage){
	require(phase==_stage);
	require(now/10<timeout);
	require(msg.sender == _party);
	_;
}

//receiver pays the price for the digital good
//If receiver pays the required ether,then the stage goes from initialized to accepted.
//In the case when receiver pays lesser amount,the ehter is transferred back to the receiver
//and the stage goes from initialized to finishedAndMaliciousReciever
function accept() allowed(msg.sender,stage.initialized) payable public {

	if(msg.value < price)
	{
		msg.sender.transfer(msg.value);
		phase=stage.finishedAndMaliciousReciever;
	}
	else
	{
		receiverAddress=msg.sender;
		nextStage();
	}
}

//sender reveals the key
//If the sender reveals the right key then the contract stage goes from accepted to keyRevealed
//If the sender reveals the wrong key, then the ether is transferred to the receiver and the stage
//goes from accepted to finishedAndMaliciousSender
function revealKey (bytes32  _key) allowed(senderAddress, stage.accepted) public
{

	bytes32 Hashed;
	bytes memory toBeHashed;
	toBeHashed=new bytes (32);
	for (uint i=0;i<32;i++)
	{
			toBeHashed[i]=_key[i];
	}
	Hashed=keccak256(toBeHashed);

	if(Hashed!=commitmentOfKey){
		receiverAddress.transfer(price);
		phase=stage.finishedAndMaliciousSender;
	}
	else
	{
		key=_key;
		nextStage();
	}
}



	function Dec(bytes memory encryptedText,bytes32 key) internal pure returns(bytes memory)
	{
		bytes memory decoded;
		bytes memory finalKey;
		finalKey=new bytes (encryptedText.length);
		finalKey=generateKey(key,encryptedText.length);
		decoded=new bytes (encryptedText.length);
		for(uint i=0;i<encryptedText.length;i++){
			decoded[i]=finalKey[i]^encryptedText[i];
		}
		return decoded;

	}

	function generateKey(bytes32 key,uint keySize) internal pure  returns(bytes memory)
	{
		bytes memory finalKey;
		bytes memory keyHashed;
		bytes32 temp;
		finalKey=new bytes(keySize);
		keyHashed=new bytes(32);

		for(uint i=0;i<keySize/32;i++){

			if(i==0)
			{
				bytes memory keyPlusIndex;
				keyPlusIndex=new bytes (33);
				for(uint j=0;j<32;j++){
					keyPlusIndex[j]=key[j];
				}
				keyPlusIndex[32]=0;
				temp=keccak256(keyPlusIndex);

				for(uint j=0;j<32;j++)
				{
					keyHashed[j]=temp[j];
					finalKey[j]=temp[j];
				}
			}
			else
			{
				delete keyHashed;
				keyHashed=new bytes(32);

				for(uint j=0;j<32;j++)
				{
						keyHashed[j]=temp[j];
				}
				temp=keccak256(keyHashed);

				for(uint j=0;j<32;j++)
				{
					keyHashed[j]=temp[j];
					finalKey[i*32+j]=temp[j];
				}

			}

		}
		return finalKey;
	}

/*	function Enc(bytes memory plainText,bytes32 key) public returns(bytes memory)
	{
		bytes memory encryptedText;
		bytes memory finalKey;
		finalKey=new bytes (plainText.length);
		finalKey=generateKey(key,plainText.length);
		encryptedText=new bytes (plainText.length);
		for(uint i=0;i<plainText.length;i++){
			encryptedText[i]=finalKey[i]^plainText[i];
		}
		return encryptedText;

	}*/

	function Operation(uint operation,bytes[] memory inputVectorsToTheGate) internal pure returns(bytes32)
	{
			bytes32 Hashed;
			if(operation==2)
			{
				bytes memory toBeHashed;

				toBeHashed=new bytes (2*inputVectorsToTheGate[0].length);
				for (uint i=0;i<inputVectorsToTheGate[0].length;i++)
				{
					toBeHashed[i]=inputVectorsToTheGate[0][i];
				}
				for (uint i=0;i<inputVectorsToTheGate[1].length;i++)
				{
					toBeHashed[i+inputVectorsToTheGate[0].length]=inputVectorsToTheGate[1][i];
				}
				Hashed=keccak256(toBeHashed);
			}
			return Hashed;
	}


function Hash(bytes memory _toBeHashed) internal pure returns(bytes32)
	{
			bytes32 Hashed;
			bytes memory toBeHashed;
			toBeHashed=new bytes (_toBeHashed.length);
			for (uint i=0;i<_toBeHashed.length;i++)
			{
				toBeHashed[i]=_toBeHashed[i];
			}
			Hashed=keccak256(toBeHashed);
			return Hashed;
	}

	function bytesToBytes32(bytes memory source) public returns (bytes32 result)
	{
    bytes memory tempEmptyStringTest = bytes(source);
    if (tempEmptyStringTest.length == 0) {
        return 0x0;
    }

    assembly {
        result := mload(add(source, 32))
    }
	}


	function complain(bytes32[][] memory complaint, bytes32[] memory gateMerkleProof,bytes[] memory encodedVectors_ThisGate)  public returns(bool)
	{


		uint8 operation;
		uint gateIndex;
		uint8 index;
		uint8[] memory indicesOfInputGates;
		uint8 noOFinputToThisParticularGate;
		bytes[] memory inputVectorsToTheGate;
		bytes32 OperationValue;
		bytes32 Out;

		bytes memory Outy;

		if(encodedVectors_ThisGate.length!=complaint.length)
		{
			phase=stage.finishedAndMaliciousReciever;
			senderAddress.transfer(price);
			return false;
		}

		for(uint i=0;i<encodedVectors_ThisGate.length;i++)
		{
			if(Hash(encodedVectors_ThisGate[i])!=complaint[i][0])
			{
				phase=stage.finishedAndMaliciousReciever;
				senderAddress.transfer(price);
				return false;
			}
		}

		gateIndex=uint8(gateMerkleProof[0][0]);

		if(!Mverify(complaint[0],gateIndex))
		{
			phase=stage.finishedAndMaliciousReciever;
			senderAddress.transfer(price);
			return false;
		}

		if(!MverifyForGate(gateMerkleProof,gateIndex))
		{

			phase=stage.finishedAndMaliciousReciever;
			senderAddress.transfer(price);
			return false;
		}

		operation=uint8(gateMerkleProof[0][1]);

		Outy=new bytes(encodedVectors_ThisGate[0].length);
		Outy=Dec(encodedVectors_ThisGate[0],key);

		for(uint i=2;i<2+maxInputLinesToAGate;i++)    //noOFinputToThisParticularGate
		{
			if(uint8(gateMerkleProof[0][i])==totalNumberOfGates){   //totalNumberOfGates
			  break;
			}
			noOFinputToThisParticularGate++;
		}

		indicesOfInputGates=new uint8[] (noOFinputToThisParticularGate);

		for(uint i=0;i<noOFinputToThisParticularGate;i++)   //noOFinputToThisParticularGate
		{
			indicesOfInputGates[i]=(uint8(gateMerkleProof[0][i+2]));
		}

		if(complaint.length-1!=noOFinputToThisParticularGate)
		{
		  return false;
		}

		//uptoComplaintLength=true;
		inputVectorsToTheGate = new bytes[] (complaint.length-1);

		for(uint i=0;i<indicesOfInputGates.length;i++)
		{
				if(!Mverify(complaint[i+1],indicesOfInputGates[i]))
				{
					phase=stage.finishedAndMaliciousReciever;
					senderAddress.transfer(price);
					return false;
				}
				inputVectorsToTheGate[i]=new bytes (encodedVectors_ThisGate[i+1].length);
				inputVectorsToTheGate[i]=Dec(encodedVectors_ThisGate[i+1],key);
		}

		OperationValue=Operation(operation,inputVectorsToTheGate);
		Out=bytesToBytes32(Outy);

		if(Out!=OperationValue)
		{
			validComplain=true;
			phase=stage.finishedAndMaliciousSender;
			return true;
		}
		phase=stage.finishedAndMaliciousReciever;
		senderAddress.transfer(price);
		return false;
	}


	function  Mverify(bytes32[] memory complaint,uint _index)  public view returns(bool)
	{
		bytes32 _value;
		_value=complaint[0];
		bytes32 Hashed;
		uint i ;
		bytes memory toBeHashed;
		toBeHashed=new bytes (64);

		for(i=0; i<depth-1; i++)
		{
			//if condition is used to decide whether to append the _value with the complaint vector or append complaint vector
			//with the _value
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

	function  MverifyForGate(bytes32[] memory gateProof,uint _index)  public view returns(bool)
	{
		bytes32 _value;
		_value=gateProof[0];
		bytes32 Hashed;
		uint i ;
		bytes memory toBeHashed;
		toBeHashed=new bytes (64);

		for(i=0; i<depth-1; i++)
		{
			//if condition is used to decide whether to append the _value with the gateProof vector or append gateProof vector
			//with the _value
			if ((_index&(1<<i))>>i == 1)
			{

				for (uint j=0;j<32;j++)
				{
					toBeHashed[j]=gateProof[i+1][j];
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
					toBeHashed[j+32]=gateProof[i+1][j];
				}

				Hashed=keccak256(toBeHashed);
				_value=Hashed;
			}
		}
		return (_value==EncCircuitRoot);
	}



	//fn is called by sender after the faiswap completion to get the ether
	function senderGetEther() allowed(senderAddress,stage.keyRevealed) public{
		require(now/10>timeout);
		senderAddress.transfer(price);
		nextStage();

	}

	//function is called by receiver in the accepted stage if the receiver doesn't reveal the key within the time. The
	//ether is transferred to the receiver
	function receiverGetEther() allowed(receiverAddress,stage.accepted) public{
		require(now/10>timeout);
		receiverAddress.transfer(price);
		phase=stage.finishedAndMaliciousSender;

	}

	//log function
	function log2(uint x) private pure returns (uint y)
  {
     assembly
     {
	     let arg := x
	     x := sub(x,1)
	     x := or(x, div(x, 0x02))
	     x := or(x, div(x, 0x04))
	     x := or(x, div(x, 0x10))
	     x := or(x, div(x, 0x100))
	     x := or(x, div(x, 0x10000))
	     x := or(x, div(x, 0x100000000))
	     x := or(x, div(x, 0x10000000000000000))
	     x := or(x, div(x, 0x100000000000000000000000000000000))
	     x := add(x, 1)
	     let m := mload(0x40)
	     mstore(m,           0xf8f9cbfae6cc78fbefe7cdc3a1793dfcf4f0e8bbd8cec470b6a28a7a5a3e1efd)
	     mstore(add(m,0x20), 0xf5ecf1b3e9debc68e1d9cfabc5997135bfb7a7a3938b7b606b5b4b3f2f1f0ffe)
	     mstore(add(m,0x40), 0xf6e4ed9ff2d6b458eadcdf97bd91692de2d4da8fd2d0ac50c6ae9a8272523616)
	     mstore(add(m,0x60), 0xc8c0b887b0a8a4489c948c7f847c6125746c645c544c444038302820181008ff)
	     mstore(add(m,0x80), 0xf7cae577eec2a03cf3bad76fb589591debb2dd67e0aa9834bea6925f6a4a2e0e)
	     mstore(add(m,0xa0), 0xe39ed557db96902cd38ed14fad815115c786af479b7e83247363534337271707)
	     mstore(add(m,0xc0), 0xc976c13bb96e881cb166a933a55e490d9d56952b8d4e801485467d2362422606)
	     mstore(add(m,0xe0), 0x753a6d1b65325d0c552a4d1345224105391a310b29122104190a110309020100)
	     mstore(0x40, add(m, 0x100))
	     let magic := 0x818283848586878898a8b8c8d8e8f929395969799a9b9d9e9faaeb6bedeeff
	     let shift := 0x100000000000000000000000000000000000000000000000000000000000000
	     let a := div(mul(x, magic), shift)
	     y := div(mload(add(m,sub(255,a))), shift)
	     y := add(y, mul(256, gt(arg, 0x8000000000000000000000000000000000000000000000000000000000000000)))
    }
	}

}
