package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/jessevdk/go-flags"

	"github.com/moon-dragon-dev/go-gen-ip-networks/internal/ipv4"
	"github.com/moon-dragon-dev/go-gen-ip-networks/internal/pow"
	"github.com/moon-dragon-dev/go-gen-ip-networks/internal/weighter"
)

type Options struct {
	NetworksCount int    `long:"networks-count" description:"Number of networks to generate"`
	IpsCount      int    `long:"ips-count" description:"Number of IPs to check"`
	ResultsDir    string `long:"results-dir" description:"Directory to store results"`
}

type Network struct {
	from uint32
	to   uint32
}

func selectIpFromNetwork(network Network) string {
	if network.from == network.to {
		return ipv4.Dec2ip(network.from)
	} else {
		return ipv4.Dec2ip(network.from + uint32(rand.Uint32())%(network.to-network.from))
	}
}

type MaskWeight struct {
	Mask   uint32
	Weight uint32
}

func main() {
	var opts Options
	parser := flags.NewParser(&opts, flags.Default)
	_, err := parser.Parse()
	if err != nil {
		log.Fatal(err)
	}

	if opts.ResultsDir == "" || opts.NetworksCount == 0 || opts.IpsCount == 0 {
		parser.WriteHelp(os.Stdout)
		os.Exit(1)
	}

	err = os.MkdirAll(opts.ResultsDir, 0755)
	if err != nil {
		log.Fatal(err)
	}

	networksFile, err := os.Create(opts.ResultsDir + "/networks.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer networksFile.Close()

	containsFile, err := os.Create(opts.ResultsDir + "/contains.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer containsFile.Close()

	maskWeight := []MaskWeight{
		{Mask: 32, Weight: 87},
		{Mask: 31, Weight: 1},
		{Mask: 30, Weight: 1},
		{Mask: 28, Weight: 1},
		{Mask: 24, Weight: 9},
		{Mask: 23, Weight: 1},
	}

	mw := make([]uint32, len(maskWeight), len(maskWeight))
	for i, v := range maskWeight {
		mw[i] = v.Weight
	}

	maskSelector := weighter.CreateSelector(mw)

	networks := make([]Network, opts.NetworksCount, opts.NetworksCount)
	nw := make([]uint32, opts.NetworksCount, opts.NetworksCount)

	rand.Seed(42)

	for i := 0; i < opts.NetworksCount; i++ {
		maskPos := maskSelector(rand.Uint32())
		mask := maskWeight[maskPos].Mask
		ip := rand.Uint32()
		net := ipv4.Mask2dec(mask)

		from := ip & net
		to := ip | ^net

		nw[i] = pow.Pow(2, 32-mask)
		networks[i] = Network{from: from, to: to}
		networksFile.WriteString(fmt.Sprintf("%s/%d\n", ipv4.Dec2ip(from), mask))
	}

	networkSelector := weighter.CreateSelector(nw)

	for i := 0; i < opts.IpsCount; i++ {
		networkPos := networkSelector(rand.Uint32())
		network := networks[networkPos]
		randomIp := selectIpFromNetwork(network)
		containsFile.WriteString(randomIp + "\n")
	}
}
