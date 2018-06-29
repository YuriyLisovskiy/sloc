package cli

import (
	"os"
	"log"
	"fmt"
)

type CLI struct{}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Print("  createblockchain\n    -address string\n\tThe address to send genesis block reward to\n\n")
	fmt.Print("  createwallet\n\tGenerates a new key-pair and saves it into the wallet file\n\n")
	fmt.Print("  getbalance\n    -address string\n\tThe address to get balance for\n\n")
	fmt.Print("  listaddresses\n\tLists all addresses from the wallet file\n\n")
	fmt.Print("  printchain\n\tPrint all the blocks of the blockchain\n\n")
	fmt.Print("  reindexutxo\n\tRebuilds the UTXO set\n\n")
	fmt.Print("  send\n    -from string\n\tSource wallet address\n    -to string\n\tDestination wallet address\n    -amount int\n\tAmount to send\n    -mine\n\tMine on the same node\n\n")
	fmt.Print("  startnode\n    -miner string\n\tStart a node with ID specified in NODE_ID env. var. -miner enables mining\n\n")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
	}
}

func (cli *CLI) Run() {
	cli.validateArgs()

//	f := countCmd.String("f", "", "set file(s) to count lines")
//	d := countCmd.String("d", "", "set folder(s) to count lines")
//	e := countCmd.String("e", "", "set file(s) and(or) folder(s) to exclude from counting lines")
//	j := countCmd.Bool("j", false, "write result to json file")
//	x := countCmd.Bool("x", false, "write result to xml file")
//	y := countCmd.Bool("y", false, "write result to yaml file")

	switch os.Args[1] {
	case "count":
		err := countCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "help":
		cli.printUsage()
		return
	default:
		cli.printUsage()
		return
	}
	if countCmd.Parsed() {

	}
}
