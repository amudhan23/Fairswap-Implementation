
pragma solidity >0.4.0 <=0.6.0;
pragma experimental ABIEncoderV2;

contract protocol2{

	uint public price;
	bytes public key;			//key
	uint public Now; //to get the current time
	uint public noOfFileChunks;
	bytes32[] public hashedFileChunks;
	bytes32[] public hashedEncryptedFileChunks;

	//address of the two parties
	address payable public  senderAddress;
	address payable public  receiverAddress;

	//deadline for completing a stage
  uint public timeout;

	//stages of fairswap
	enum stage{created,initialized,active,revealed,finalized}

	stage public phase = stage.created;

	//modifier to check the eligibility of party, time and phase for a function execution
	modifier allowedBefore(address _party,stage _stage){
		require(phase==_stage);
		require(now<timeout);
		require(msg.sender == _party);
		_;
	}

	modifier allowedAfter(address _party,stage _stage){
		require(phase==_stage);
		require(now>timeout);
		require(msg.sender == _party);
		_;
	}

	//fn used to move to the next stage fron the current stage
	function nextStage() public{
		phase=stage(uint(phase)+1);
		timeout=now + 60 seconds;
	}

	function Nowy() public {
		Now=now;
	}

	//state variable initilization including the circuit gate's operation
	//stage goes from created to initialized
	constructor(uint _price,bytes32[] memory _hashedFileChunks,bytes32[] memory _hashedEncryptedFileChunks) payable public{

		senderAddress=msg.sender;
		require(_hashedFileChunks.length ==_hashedEncryptedFileChunks.length);
		noOfFileChunks=_hashedFileChunks.length;
		hashedFileChunks = new bytes32[] (noOfFileChunks);
		for(uint i=0;i<noOfFileChunks;i++)
		{
			hashedFileChunks[i]=_hashedFileChunks[i];
		}
		hashedEncryptedFileChunks = new bytes32[] (noOfFileChunks);
		for(uint i=0;i<noOfFileChunks;i++)
		{
			hashedEncryptedFileChunks[i]=_hashedEncryptedFileChunks[i];
		}
		price = _price;
		nextStage();
	}

/* allowedBefore(senderAddress,stage.active) */
	//receiver pays the price for the digital good
	//If receiver pays the required ether,then the stage goes from initialized to accepted.
	//In the case when receiver pays lesser amount,the ehter is transferred back to the receiver
	//and the stage goes from initialized to finishedAndMaliciousReciever
	/*  allowedBefore(msg.sender,stage.initialized) */
	function accept() allowedBefore(msg.sender,stage.initialized)  payable public {

		if(msg.value < price)
		{
			msg.sender.transfer(msg.value);
		}
		else
		{
			receiverAddress=msg.sender;
			nextStage();
		 }
	}


	function keyReveal(bytes memory _key) allowedBefore(senderAddress,stage.active)  public {
		key=_key;
		nextStage();
	}

	function complain(bytes memory EncryptedFileChunk,uint index) allowedBefore(receiverAddress,stage.revealed) public returns(bool){

		if(hashedEncryptedFileChunks[index]!=Hash(EncryptedFileChunk))
		{
			phase=stage.finalized;
			senderAddress.transfer(price);
			return false;
		}
		if(hashedFileChunks[index]!=Hash(Dec(EncryptedFileChunk)))
		{
			phase=stage.finalized;
			receiverAddress.transfer(price);
			return true;
		}

		phase=stage.finalized;
		senderAddress.transfer(price);
		return false;
	}

	//Operation function is circuit specific
	//This function is implemented for the circuit given in the repository.
	//pure functono
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

	//used to Dccrypt the encryptedtext with the key
	function Dec(bytes memory encryptedtext) public returns(bytes memory)
	{
		bytes memory decoded;
		decoded=new bytes (encryptedtext.length);
	/*	encrypted=encryptedtext;
		for(uint i=0;i<32;i++)
		{
			encrypted[i]=encryptedtext[i];
		}
		decoded=new bytes (32);

		//keyMemory=key;*/
		for(uint i=0;i<encryptedtext.length;i++){
			decoded[i]=key[i]^encryptedtext[i];
		}
		return decoded;
	}

	//fn is called by sender after the faiswap completion to get the ether
	/* allowedAfter(senderAddress,stage.revealed) */
	function senderGetEther() allowedAfter(senderAddress,stage.revealed)   public{
		/* require(now>timeout); */
		senderAddress.transfer(price);
		phase=stage.finalized;

	}

	//function is called by receiver in the accepted stage if the receiver doesn't reveal the key within the time. The
	//ether is transferred to the receiver
	function receiverGetEther() allowedAfter(receiverAddress,stage.active) public{
		/* require(now>timeout); */
		receiverAddress.transfer(price);
		phase=stage.finalized;
	}

}
