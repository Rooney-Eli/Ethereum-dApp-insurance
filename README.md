This is a dApp mocking an insurance policy using the ethereum blockchain and solidity at it's core.    
  
**[Demo Video](https://www.youtube.com/watch?v=KQNK9ax8Z08)**
  
Insurance Policy Smart Contract Order of Operations:  

Oracle Service Steps:  
truffle: deploy oracle service using account[2]  
Request router: set oracle address  
Request router: set private key to account[2]  
await 'real world' data  

Insurance Provider Steps:  
truffle: deploy insurance contract using account[0]  
react: set oracle address using account[0]  
react: deposit coverage amount into contract using account[0]  
react: End policy?  
  
Insurance Client Steps:  
react: Deposit payment using account[1]  
react: claim policy?  
  

