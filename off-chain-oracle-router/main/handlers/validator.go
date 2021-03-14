package handlers

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"net/http"
	"off-chain-oracle-router/main/config"
	InsuranceOracle "off-chain-oracle-router/main/contracts"
	"off-chain-oracle-router/main/data"
)

type Validation struct {
	l *log.Logger
}

func NewValidation(l *log.Logger) *Validation {
	return &Validation{l}
}

func (v *Validation) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	client, err := ethclient.Dial(config.NetworkUrl)
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress(config.InsuranceOracleAddress)

	if r.Method == http.MethodPost {
		clm := v.addValidation(rw, r)
		setPolicyClaimable(client, contractAddress, clm)
		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)

}


func (v *Validation) addValidation(rw http.ResponseWriter, r *http.Request) bool {
	v.l.Println("Handle POST greeting")

	vld := &data.Validation{}

	err := vld.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	return vld.Validation

}


func setPolicyClaimable(client *ethclient.Client, addr common.Address, newClaimable bool) {

	privateKey, err := crypto.HexToECDSA(config.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	instance, err := InsuranceOracle.NewInsuranceOracle(addr, client)
	if err != nil {
		log.Fatal(err)
	}


	tx, err := instance.SetPolicyIsClaimable(auth, newClaimable)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("transaction ID: %s \n", tx.Hash().Hex())

	result, err := instance.GetPolicyIsClaimable(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result from oracle: %t", result)
}



