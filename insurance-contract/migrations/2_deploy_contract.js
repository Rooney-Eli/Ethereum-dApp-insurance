const Insurance = artifacts.require("./Insurance.sol");

module.exports = function(deployer) {

    var policyAmount = 100; //wei
    var requiredPayment = 2; //wei
    var time = 2592000; //2,592,000 = 30 days in UNIX time 
    
    deployer.deploy(Insurance, policyAmount, requiredPayment, time)
}