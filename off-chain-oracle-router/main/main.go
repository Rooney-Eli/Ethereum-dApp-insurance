package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"net/http"
	"off-chain-oracle-router/main/config"
	InsuranceOracle "off-chain-oracle-router/main/contracts"
	"off-chain-oracle-router/main/handlers"
	"os"
	"os/signal"
	"strings"
	"time"
)


func main() {

	client, err := ethclient.Dial(config.NetworkUrl)
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress(config.InsuranceOracleAddress)


	l := log.New(os.Stdout, "insurance-api", log.LstdFlags)
	gh := handlers.NewValidation(l)

	l.Println("Starting Endpoint.")

	sm := http.NewServeMux()
	sm.Handle("/", gh)

	s:= &http.Server{
		Addr: ":9090",
		Handler: sm,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,
	}


	go func() {
		l.Println("Starting Router on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting router: %s\n", err)
			os.Exit(1)
		}

	}()

	go func() {
		fmt.Printf("Getting events")
		getEvents(client, contractAddress)
	}()


	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <- sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)

}

func getEvents(client *ethclient.Client, addr common.Address) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{addr},
	}

	contractAbi, err := abi.JSON(strings.NewReader(InsuranceOracle.InsuranceOracleABI))
	if err != nil {
		log.Fatal(err)
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:

			datas , err := contractAbi.Unpack("PolicyInitialized", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			if datas[0] != nil {
				message := datas[0].(bool)
				fmt.Printf("Policy activation: %t\n", message)
			}
		}
	}
}
