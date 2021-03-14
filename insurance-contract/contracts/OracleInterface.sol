pragma solidity ^0.7.4;

interface OracleInterface {

    function declarePolicyActive(bool value) external;

    function setPolicyIsClaimable(bool value) external;

    function getPolicyIsClaimable() external view returns (bool);

    function getAddress() external view returns (address);

    function testConnection() external pure returns (bool);
}