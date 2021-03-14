const InsuranceOracle = artifacts.require("./InsuranceOracle.sol");

module.exports = function(deployer) {
    deployer.deploy(InsuranceOracle);
}