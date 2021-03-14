pragma solidity ^0.5.16;

import "./Ownable.sol";

contract InsuranceOracle is Ownable {

    bool policyIsClaimable = false;

    event PolicyInitialized(
        bool isActive
    );

    function declarePolicyActive(bool value) external {
        emit PolicyInitialized(value);
    }

    function setPolicyIsClaimable(bool value) onlyOwner public {
        policyIsClaimable = value;
    }

    function getPolicyIsClaimable() public view returns (bool) {
        return policyIsClaimable;
    }


    function getAddress() public view returns (address) {
        return address(this);
    }

    function testConnection() public pure returns (bool) {
        return true; 
    }

}
