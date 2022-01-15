package network

// import (
// 	"bytes"
// 	"encoding/gob"
// 	"encoding/hex"
// 	"fmt"
// 	"io"
// 	"io/ioutil"
// 	"log"
// 	"net"
// 	"nexus/pkg/block"
// 	"nexus/pkg/chain"
// 	"nexus/utils"
// 	"os"
// 	"runtime"
// 	"syscall"

// 	"github.com/vrecan/death"
// )

// const (
// 	protocol      = "tcp"
// 	version       = 1
// 	commandLength = 12
// )

// // TODO: store it in some kind of redis like storage
// var (
// 	nodeAddr       string // addr of the working node
// 	minerAddr      string
// 	KnownNodes     = []string{"localhost:9099"}
// 	blockInTransit = [][]byte{}
// 	memoryPool     = make(map[string][]byte)
// )

// type Addr struct {
// 	AddrList []string
// }

// type Block struct {
// 	AddrFrom string
// 	Block    []byte
// }

// type GetBlocks struct {
// 	AddrFrom string
// }

// type GetData struct {
// 	AddrFrom string
// 	Type     string
// 	Id       []byte
// }

// type Inventory struct {
// 	AddFrom string
// 	Type    string
// 	Items   [][]byte
// }

// type BlockData struct {
// 	Addfrom string
// 	Data    []byte
// }

// type Version struct {
// 	Version    int
// 	BestHeight int
// 	AddrFrom   string
// }

// func CloseDB(blockchain *chain.Chain) {
// 	d := death.NewDeath(syscall.SIGINT, syscall.SIGALRM, os.Interrupt)
// 	d.WaitForDeathWithFunc(func() {
// 		defer os.Exit(1)
// 		defer runtime.Goexit()
// 		// TODO: close db conn
// 	})
// }

// func SendData(addr string, data []byte) {
// 	conn, err := net.Dial(protocol, addr)
// 	if err != nil {
// 		fmt.Printf("%s not available\n", addr)

// 		var newNodes []string

// 		for _, node := range KnownNodes {
// 			if node != addr {
// 				newNodes = append(newNodes, node)
// 			}
// 		}

// 		KnownNodes = newNodes
// 		return
// 	}

// 	defer conn.Close()

// 	_, err = io.Copy(conn, bytes.NewReader(data))
// 	utils.HandleException(err)
// }

// func ExtractCmd(req []byte) []byte {
// 	return req[:commandLength]
// }

// func SendAddr(address string) {
// 	nodes := Addr{
// 		AddrList: KnownNodes,
// 	}
// 	nodes.AddrList = append(nodes.AddrList, nodeAddr)
// 	payload := utils.Encode(nodes)
// 	req := append(CmdToBytes("addr"), payload...)
// 	SendData(address, req)
// }

// func SendBlock(addr string, b *block.Block) {
// 	serialized, err := b.Serialize()
// 	utils.HandleException(err)

// 	data := Block{
// 		AddrFrom: addr,
// 		Block:    serialized,
// 	}
// 	payload := utils.Encode(data)
// 	req := append(CmdToBytes("block"), payload...)
// 	SendData(addr, req)
// }

// func SendInventory(addr, kind string, items [][]byte) {
// 	inventory := Inventory{
// 		AddFrom: nodeAddr,
// 		Type:    kind,
// 		Items:   items,
// 	}
// 	payload := utils.Encode(inventory)
// 	req := append(CmdToBytes("inventory"), payload...)
// 	SendData(addr, req)
// }

// func SendBlockData(addr string, data []byte) {
// 	blockData := BlockData{
// 		Addfrom: nodeAddr,
// 		Data:    data,
// 	}

// 	payload := utils.Encode(blockData)
// 	req := append(CmdToBytes("block_data"), payload...)
// 	SendData(addr, req)
// }

// func SendVersion(addr string, chain *chain.Chain) {
// 	height := chain.GetBestHeight()

// 	vs := Version{
// 		Version:    version,
// 		BestHeight: height,
// 		AddrFrom:   nodeAddr,
// 	}

// 	payload := utils.Encode(vs)
// 	req := append(CmdToBytes("version"), payload...)
// 	SendData(addr, req)
// }

// func SendGetData(addr, kind string, id []byte) {
// 	payload := utils.Encode(GetData{nodeAddr, kind, id})
// 	req := append(CmdToBytes("get_data"), payload...)
// 	SendData(addr, req)
// }

// func SendGetBlocks(addr string) {
// 	payload := utils.Encode(GetBlocks{AddrFrom: nodeAddr})
// 	req := append(CmdToBytes("get_blocks"), payload...)
// 	SendData(addr, req)
// }

// func HandleAddr(req []byte) {
// 	var buff bytes.Buffer
// 	var payload Addr

// 	buff.Write(req[commandLength:])
// 	dec := gob.NewDecoder(&buff)
// 	err := dec.Decode(&payload)
// 	utils.HandleException(err)

// 	KnownNodes = append(KnownNodes, payload.AddrList...)
// 	fmt.Printf("there are %d known nodes\n", len(KnownNodes))
// 	RequestBlocks()
// }

// func RequestBlocks() {
// 	for _, node := range KnownNodes {
// 		SendGetBlocks(node)
// 	}
// }

// func HandleBlock(req []byte, c *chain.Chain) {
// 	var buff bytes.Buffer
// 	var payload Block

// 	buff.Write(req[commandLength:])
// 	dec := gob.NewDecoder(&buff)
// 	err := dec.Decode(&payload)
// 	utils.HandleException(err)

// 	blockData := payload.Block
// 	b, err := block.Deserialize(blockData)
// 	utils.HandleException(err)

// 	fmt.Println("received a new block")
// 	c.AddBlock(b)

// 	fmt.Printf("added block %x\n", b.Hash)

// 	if len(blockInTransit) > 0 {
// 		blockHash := blockInTransit[0]
// 		SendGetData(payload.AddrFrom, "block", blockHash)
// 		blockInTransit = blockInTransit[1:]
// 	} else {
// 		// utxoSet := c.utsxset(c)
// 		// utxoSet.ReIndex()
// 	}
// }

// func HandleGetBlocks(request []byte, c *chain.Chain) {
// 	var buff bytes.Buffer
// 	var payload GetBlocks

// 	buff.Write(request[commandLength:])
// 	dec := gob.NewDecoder(&buff)
// 	err := dec.Decode(&payload)
// 	utils.HandleException(err)

// 	blocks := chain.GetBlockHashes()
// 	SendInventory(payload.AddrFrom, "block", blocks)
// }

// func HandleGetData(request []byte, c *chain.Chain) {
// 	var buff bytes.Buffer
// 	var payload GetData

// 	buff.Write(request[commandLength:])
// 	dec := gob.NewDecoder(&buff)
// 	err := dec.Decode(&payload)
// 	utils.HandleException(err)

// 	if payload.Type == "block" {
// 		block, err := c.GetBlock([]byte(payload.Id))
// 		if err != nil {
// 			return
// 		}

// 		SendBlock(payload.AddrFrom, &block)
// 	}

// 	if payload.Type == "block_data" {
// 		dataId := hex.EncodeToString(payload.Id)
// 		data := memoryPool[dataId]

// 		SendBlockData(payload.AddrFrom, data)
// 	}
// }

// func HandleVersion(request []byte, c *chain.Chain) {
// 	var buff bytes.Buffer
// 	var payload Version

// 	buff.Write(request[commandLength:])
// 	dec := gob.NewDecoder(&buff)
// 	err := dec.Decode(&payload)
// 	utils.HandleException(err)

// 	bestHeight := chain.GetBestHeight()
// 	otherHeight := payload.BestHeight

// 	if bestHeight < otherHeight {
// 		SendGetBlocks(payload.AddrFrom)
// 	} else if bestHeight > otherHeight {
// 		SendVersion(payload.AddrFrom, c)
// 	}

// 	if !NodeIsKnown(payload.AddrFrom) {
// 		KnownNodes = append(KnownNodes, payload.AddrFrom)
// 	}
// }

// func NodeIsKnown(addr string) bool {
// 	for _, node := range KnownNodes {
// 		if node == addr {
// 			return true
// 		}
// 	}

// 	return false
// }

// func HandleBlockData(request []byte, c *chain.Chain) {
// 	var buff bytes.Buffer
// 	var payload BlockData

// 	buff.Write(request[commandLength:])
// 	dec := gob.NewDecoder(&buff)
// 	err := dec.Decode(&payload)
// 	utils.HandleException(err)

// 	txData := payload.Data
// 	tx := c.DeserializeData(txData)
// 	memoryPool[hex.EncodeToString(tx.ID)] = tx

// 	fmt.Printf("%s, %d", nodeAddr, len(memoryPool))

// 	if nodeAddr == KnownNodes[0] {
// 		for _, node := range KnownNodes {
// 			if node != nodeAddr && node != payload.Addfrom {
// 				SendInventory(node, "block_data", [][]byte{tx.ID})
// 			}
// 		}
// 	} else {
// 		if len(memoryPool) >= 2 && len(minerAddr) > 0 {
// 			MineBlockDAta(c)
// 		}
// 	}
// }

// func MineBlockDAta(c *chain.Chain) {
// 	var data []*chain.Transactions

// 	for id := range memoryPool {
// 		fmt.Printf("tx: %s\n", memoryPool[id].Id)
// 		tx := memoryPool[id]
// 		if chain.VerifyTransaction(&tx) {
// 			data = append(data, &tx)
// 		}
// 	}

// 	if len(data) == 0 {
// 		fmt.Println("All Transactions are invalid")
// 		return
// 	}

// 	cbTx := c.CoinbaseTx(mineAddress, "")
// 	data = append(data, cbTx)

// 	newBlock := c.MineBlock(data)
// 	UTXOSet := blockchain.UTXOSet{chain}
// 	UTXOSet.Reindex()

// 	fmt.Println("New Block mined")

// 	for _, tx := range data {
// 		txID := hex.EncodeToString(tx.ID)
// 		delete(memoryPool, txID)
// 	}

// 	for _, node := range KnownNodes {
// 		if node != nodeAddr {
// 			SendInventory(node, "block", [][]byte{newBlock.Hash})
// 		}
// 	}

// 	if len(memoryPool) > 0 {
// 		MineBlockDAta(c)
// 	}
// }

// func HandleInv(request []byte, c *chain.Chain) {
// 	var buff bytes.Buffer
// 	var payload Inventory

// 	buff.Write(request[commandLength:])
// 	dec := gob.NewDecoder(&buff)
// 	err := dec.Decode(&payload)
// 	utils.HandleException(err)

// 	fmt.Printf("Received inventory with %d %s\n", len(payload.Items), payload.Type)

// 	if payload.Type == "block" {
// 		blockInTransit = payload.Items

// 		blockHash := payload.Items[0]
// 		SendGetData(payload.AddFrom, "block", blockHash)

// 		newInTransit := [][]byte{}
// 		for _, b := range blockInTransit {
// 			if bytes.Compare(b, blockHash) != 0 {
// 				newInTransit = append(newInTransit, b)
// 			}
// 		}
// 		blockInTransit = newInTransit
// 	}

// 	if payload.Type == "tx" {
// 		txID := payload.Items[0]

// 		if memoryPool[hex.EncodeToString(txID)].ID == nil {
// 			SendGetData(payload.AddFrom, "tx", txID)
// 		}
// 	}
// }

// func HandleConnection(conn net.Conn, c *chain.Chain) {
// 	req, err := ioutil.ReadAll(conn)
// 	defer conn.Close()

// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	command := BytesToCmd(req[:commandLength])
// 	fmt.Printf("Received %s command\n", command)

// 	switch command {
// 	case "addr":
// 		HandleAddr(req)
// 	case "block":
// 		HandleBlock(req, c)
// 	case "inventory":
// 		HandleInv(req, c)
// 	case "get_blocks":
// 		HandleGetBlocks(req, c)
// 	case "get_data":
// 		HandleGetData(req, c)
// 	case "block_Data":
// 		HandleBlockData(req, c)
// 	case "version":
// 		HandleVersion(req, c)
// 	default:
// 		fmt.Println("Unknown command")
// 	}

// }

// func StartServer(nodeID, minerAddr string) {
// 	nodeAddr = fmt.Sprintf("localhost:%s", nodeID)
// 	minerAddr = minerAddr
// 	ln, err := net.Listen(protocol, nodeAddr)
// 	utils.HandleException(err)

// 	defer ln.Close()

// 	chain := chain.ContinueBlockChain(nodeID)
// 	defer chain.Database.Close()
// 	go CloseDB(chain)

// 	if nodeAddr != KnownNodes[0] {
// 		SendVersion(KnownNodes[0], chain)
// 	}
// 	for {
// 		conn, err := ln.Accept()
// 		if err != nil {
// 			log.Panic(err)
// 		}
// 		go HandleConnection(conn, chain)

// 	}
// }
