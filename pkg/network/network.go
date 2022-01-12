package network

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"nexus/pkg/block"
	"nexus/pkg/chain"
	"nexus/utils"
	"os"
	"runtime"
	"syscall"

	"github.com/vrecan/death"
)

const (
	protocol      = "tcp"
	version       = 1
	commandLength = 12
)

// TODO: store it in some kind of redis like storage
var (
	nodeAddr       string // addr of the working node
	minerAddr      string
	KnownNodes     = []string{"localhost:9099"}
	blockInTransit = [][]byte{}
	memoryPool     = make(map[string][]byte)
)

type Addr struct {
	AddrList []string
}

type Block struct {
	AddrFrom string
	Block    []byte
}

type GetBlocks struct {
	AddrFrom string
}

type GetData struct {
	AddrFrom string
	Type     string
	Id       []byte
}

type Inventory struct {
	AddFrom string
	Type    string
	Items   [][]byte
}

type Data struct {
	Addfrom     string
	Transaction []byte
}

type Version struct {
	Version    int
	BestHeight int
	AddrFrom   string
}

func CloseDB(blockchain *chain.Chain) {
	d := death.NewDeath(syscall.SIGINT, syscall.SIGALRM, os.Interrupt)
	d.WaitForDeathWithFunc(func() {
		defer os.Exit(1)
		defer runtime.Goexit()
		// TODO: close db conn
	})
}

func HandleConnection(conn net.Conn, blockchain *chain.Chain) {
	req, err := ioutil.ReadAll(conn)
	defer conn.Close()

	utils.HandleException(err)

	command := BytesToCmd(req[:commandLength])

	switch command {
	default:
		fmt.Println("unknown command")

	}
}

func SendData(addr string, data []byte) {
	conn, err := net.Dial(protocol, addr)
	if err != nil {
		fmt.Printf("%s not available\n", addr)

		var newNodes []string

		for _, node := range KnownNodes {
			if node != addr {
				newNodes = append(newNodes, node)
			}
		}

		KnownNodes = newNodes
		return
	}

	defer conn.Close()

	_, err = io.Copy(conn, bytes.NewReader(data))
	utils.HandleException(err)
}

func ExtractCmd(req []byte) []byte {
	return req[:commandLength]
}

func SendAddr(address string) {
	nodes := Addr{
		AddrList: KnownNodes,
	}
	nodes.AddrList = append(nodes.AddrList, nodeAddr)
	payload := utils.Encode(nodes)
	req := append(CmdToBytes("addr"), payload...)
	SendData(address, req)
}

func SendBlock(addr string, b *block.Block) {
	serialized, err := b.Serialize()
	utils.HandleException(err)

	data := Block{
		AddrFrom: addr,
		Block:    serialized,
	}
	payload := utils.Encode(data)
	req := append(CmdToBytes("block"), payload...)
	SendData(addr, req)
}
