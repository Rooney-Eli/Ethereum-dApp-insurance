import Web3 from 'web3'
import './App.css';
import {useEffect, useState} from 'react';
import {INSURANCE_ABI, INSURANCE_ADDRESS, ORACLE_ADDRESS} from './config';
import {Contract} from "web3-eth-contract";


function App() {

  const [account, setAccount] = useState("0x0000000000000000000000000000000000000000")
  const [insurance, setInsurance] = useState(Contract)
  const [providerAddr, setProviderAddr] = useState("0x0000000000000000000000000000000000000000")
  const [policyHolderAddr, setPolicyHolderAddr] = useState("0x0000000000000000000000000000000000000000")
  const [oracleAddr, setOracleAddr] = useState("0x0000000000000000000000000000000000000000")
  const [contractBalance, setContractBalance] = useState(0)
  const [endTime, setEndTime] = useState(0)


  useEffect(() => {
    loadBlockchainData()
  }, []);


  async function loadBlockchainData() {

    const web3 = new Web3(Web3.givenProvider || "http://localhost:8545")
    const network = await web3.eth.net.getNetworkType()
    const accounts = await web3.eth.getAccounts()
    setAccount(accounts[0])

    const ins = new web3.eth.Contract(INSURANCE_ABI, INSURANCE_ADDRESS)
    setInsurance(ins)
    console.log("insurance: " + ins)

    const oracle = await ins.methods.getOracleAddress().call()
    setOracleAddr(oracle)

    const provider = await ins.methods.getInsuranceProvider().call()
    setProviderAddr(provider)

    const policyHolder = await ins.methods.getPolicyHolder().call()
    setPolicyHolderAddr(policyHolder)


    const balance = await ins.methods.balanceOfContract().call()
    setContractBalance(balance)

    const end = await ins.methods.getPolicyEndTime().call()
    setEndTime(end)
    console.log("network: ", network)

    console.log("current account: " + accounts[0])
    console.log("Policy provider: " + provider)
    console.log("Policy holder: " + policyHolder)
    console.log("Current contract balance: " + balance)

    ins.events.PolicyEnded({})
        .on('data', event => console.log(event))

    ins.events.PolicyClaimed({})
        .on('data', event => console.log(event))

    await new Promise(resolve => setTimeout(() => resolve(), 2000))
  }


  const handleDeposit = async (e) => {
    e.preventDefault();

    const web3 = new Web3(Web3.givenProvider || "http://localhost:8545")
    const network = await web3.eth.net.getNetworkType()
    const accounts = await web3.eth.getAccounts()
    const ins = new web3.eth.Contract(INSURANCE_ABI, INSURANCE_ADDRESS)


    console.log("sending transaction from: " + account)
    await ins.methods.depositPayment().send({
      from: accounts[0],
      value: '2'
    })

  }

  const handleClaimPolicy = async (e) => {
    e.preventDefault();


    const web3 = new Web3(Web3.givenProvider || "http://localhost:8545")
    const network = await web3.eth.net.getNetworkType()
    const accounts = await web3.eth.getAccounts()
    const ins = new web3.eth.Contract(INSURANCE_ABI, INSURANCE_ADDRESS)

    console.log("Claiming policy...")
    await ins.methods.claimPolicy().call()

  }


  return (
      <div className="App">
        <h1>Policy Holder Page</h1>
        <p>Your account: {account}</p>
        <p>Provider account: {providerAddr}</p>
        <p>Holder account: {policyHolderAddr}</p>
        <p>Current contract balance: {contractBalance}</p>
        <p>Policy End time (Unix): {endTime}</p>
        <p>Oracle address: {oracleAddr}</p>

        <button onClick={handleDeposit}>
          <input type="submit" value="Deposit payment (2 wei)" />
        </button>

        <button onClick={handleClaimPolicy}>
          <input type="submit" value="Claim Policy" />
        </button>
      </div>
  );
}

export default App;
