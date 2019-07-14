
pragma solidity >0.4.0 <=0.6.0;
pragma experimental ABIEncoderV2;

contract protocol2{

	uint public price;
	bytes public key;			//key
	uint public Now; //to get the current time
	uint public noOfFileChunks;
	bytes32[] public hashedFileChunks;
	bytes32[] public hashedEncryptedFileChunks;
	bool public complaintInside;
	bytes32 public decrypted;
	bytes public encrypted;
	bytes32 public hashDecrypted;

	//address of the two parties
	address payable public  senderAddress;
	address payable public  receiverAddress;

	//deadline for completing a stage
  uint public timeout;

	//stages of fairswap
	enum stage{created,initialized,accepted,keyRevealed,finishedAndBothHonestParties,finishedAndMaliciousSender,finishedAndMaliciousReciever}
	//created - sender creates the Contract
	//initialized - sender initializes the state variables
	//stages created and initialized happens together when contract is deployed
	//accepted - receiver has paid the price
	//keyRevealed - sender has revealed the key
	//finishedAndBothHonestParties - fairswap completed and both parties were honest
	//finishedAndMaliciousSender - fairswap not completed as sender was malicious
	//finishedAndMaliciousReciever - fairswap not completed as receiver was malicious

	stage public phase = stage.created;

	//fn used to move to the next stage fron the current stage
	function nextStage() public{
		phase=stage(uint(phase)+1);
		timeout=now/10+ 10 seconds;
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

	//modifier to check the eligibility of party, time and phase for a function execution
	modifier allowed(address _party,stage _stage){
		require(phase==_stage);
		require(now/10<timeout);
		require(msg.sender == _party);
		_;
	}

	function keyReveal(bytes memory _key) allowed(senderAddress,stage.accepted) public {
		key=_key;
		nextStage();
	}
	//receiver pays the price for the digital good
	//If receiver pays the required ether,then the stage goes from initialized to accepted.
	//In the case when receiver pays lesser amount,the ehter is transferred back to the receiver
	//and the stage goes from initialized to finishedAndMaliciousReciever
	function accept() allowed(msg.sender,stage.initialized) payable public {

		receiverAddress=msg.sender;
		if(msg.value < price)
		{
			msg.sender.transfer(msg.value);
			phase=stage.finishedAndMaliciousReciever;
		}
		else
		{
			nextStage();
		}
	}

	function complain(bytes memory EncryptedFileChunk,uint index) allowed(receiverAddress,stage.keyRevealed) public returns(bool){

		complaintInside=true;
		encrypted=EncryptedFileChunk;
		decrypted=Hash(EncryptedFileChunk);
		hashDecrypted=hashedEncryptedFileChunks[index];
		if(hashedEncryptedFileChunks[index]!=Hash(EncryptedFileChunk))
		{
			phase=stage.finishedAndMaliciousReciever;
			senderAddress.transfer(price);
			return false;
		}
		if(hashedFileChunks[index]!=Hash(Dec(EncryptedFileChunk)))
		{
			phase=stage.finishedAndMaliciousSender;
			receiverAddress.transfer(price);
			return true;
		}
		else
		{
			phase=stage.finishedAndMaliciousReciever;
			senderAddress.transfer(price);
			return false;
		}
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

}
