pragma solidity ^0.7.4;

import "./OracleInterface.sol";
import "./Ownable.sol";


contract Insurance is Ownable {
    
    uint policyAmount;
    uint policyPrice;
    uint policyEndTime;

    uint policyBalance;
    uint paymentBalance;

    bool ended = false;
    bool active = false;
    bool policyIsSet = false;

    address payable insuranceProvider;
    address payable policyHolder;
    address internal onChainOracle;

    event PolicyEnded(address insuranceProvider, uint policyAmount, uint payment);
    event PolicyClaimed(address policyHolder, uint policyAmount, uint payment);

    OracleInterface internal insuranceOracle = OracleInterface(onChainOracle);

    constructor (
        uint _policyAmount,      // wei
        uint _policyPrice,       // wei
        uint _policyDuration     // UNIX time
    ) {
        policyAmount = _policyAmount;
        policyPrice = _policyPrice;
        policyEndTime = block.timestamp + _policyDuration;

        insuranceProvider = msg.sender;
    }

    function depositPolicyAmount() external payable {
        require(msg.sender == insuranceProvider, "You are not the provider!");
        require(msg.value == policyAmount, "Invalid policy amount.");
        require(!ended, "Policy has ended");
        require (!policyIsSet, "Policy amount already deposited!");
        policyIsSet = true;
        policyBalance = msg.value; 
    }

    function depositPayment() external payable {
        require(msg.value == policyPrice, "Not the correct payment amount!");
        require(!active, "Policy is already active.");
        require(!ended, "Policy has ended");
        policyHolder = msg.sender;
        active = true;
        policyBalance += msg.value;
    }


    function endPolicy() public {
        require( msg.sender ==  insuranceProvider, "Only the provider is allowed to end.");
        require(block.timestamp >= policyEndTime, "Period not yet expired.");
        require(!ended, "endPolicy has already been called.");

        ended = true;

        emit PolicyEnded(insuranceProvider, policyAmount, policyPrice);

        insuranceProvider.transfer(policyBalance + paymentBalance);
    }

    function claimPolicy() public {
        require(block.timestamp <= policyEndTime, "Period expired.");
        require(!ended, "Policy has already ended.");
        require(insuranceOracle.getPolicyIsClaimable(), "Not yet a valid claim.");

        ended = true;

        emit PolicyClaimed(policyHolder, policyAmount, policyPrice);

        policyHolder.transfer(policyBalance + paymentBalance);
    }

    function setOracleAddress(address _onChainOracle) external onlyOwner returns (bool) {
        onChainOracle = _onChainOracle;
        insuranceOracle = OracleInterface(onChainOracle); 
        return insuranceOracle.testConnection();
    }

    function getOracleAddress() external view returns (address) {
        return onChainOracle;
    }

    function testOracleConnection() public view returns (bool) {
        return insuranceOracle.testConnection(); 
    }

    function balanceOfContract() external view returns (uint) {
        return address(this).balance;
    }

    function getInsuranceProvider() external view returns (address) {
        return insuranceProvider;
    }
    
    function getPolicyHolder() external view returns (address) {
        return policyHolder;
    }

    function getPolicyAmount() external view returns (uint){
        return policyAmount;
    }

    function getPolicyPrice() external view returns (uint){
        return policyPrice;
    }

    function getPolicyEndTime() external view returns (uint){
        return policyEndTime;
    }

}