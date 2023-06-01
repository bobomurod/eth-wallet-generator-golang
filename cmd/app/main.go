package main

import "fmt"
import "walletGenerate/internal/ethereum"
import "walletGenerate/internal/collector"

func main() {
	fmt.Println("Start generating wallet")
	collector.WriterCsv(ethereum.Ethereum())
}
