package main

import "fmt"

type BlockchainClient interface {
	GetNetworkName() string
	GetBalance(address string) float64
	GetLatestBlock() int
}

type Ethereum struct {
	NetworkName string
	LatestBlock int
}
type BSC struct {
	NetworkName string
	LatestBlock int
}
type Arbitrum struct {
	NetworkName string
	LatestBlock int
}

func (e *Ethereum) GetNetworkName() string {	
	return e.NetworkName
}

func (e *Ethereum) GetBalance(address string) float64{
	return 1500.09
}

func (e *Ethereum) GetLatestBlock()int {
	return e.LatestBlock
}
func (b *BSC) GetNetworkName() string {
	return b.NetworkName
}

func (b *BSC) GetBalance(address string) float64{
	return 230.99
}

func (b *BSC) GetLatestBlock()int {
	return b.LatestBlock
}
func (a *Arbitrum) GetNetworkName() string {
	return a.NetworkName
}

func (a *Arbitrum) GetBalance(address string) float64{
	return 999.65
}

func (a *Arbitrum) GetLatestBlock()int {
	return a.LatestBlock
}


func ClientInfo(b BlockchainClient, address string) {
	fmt.Println("Network Name is:", b.GetNetworkName())
	fmt.Println("Address:", address)
	fmt.Println("Balance is:", b.GetBalance(address))
	fmt.Println("Latest block is", b.GetLatestBlock())
}

func main() {
	ethereum := Ethereum{
		NetworkName: "Ethereum",
		LatestBlock: 123,	
	}
	bsc := BSC{
		NetworkName: "BSC",
		LatestBlock: 321,	
	}
	arbitrum := Arbitrum{
		NetworkName: "Arbitrum",
		LatestBlock: 1000,	
	}

	ClientInfo(&ethereum, "eth_sadw_0x09")
	ClientInfo(&bsc, "bsc_sw_0sdx09")
	ClientInfo(&arbitrum, "arb_wosd_d0x09")
}


