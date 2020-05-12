package main

import (
	"fmt"
	"kmdgo"

	"github.com/satindergrewal/shurli/sagoutil"
)

func main() {
	var chains = []kmdgo.AppType{"komodo", "PIRATE"}

	var wallets []sagoutil.WInfo
	wallets = sagoutil.WalletInfo(chains)

	fmt.Println(wallets)

}
