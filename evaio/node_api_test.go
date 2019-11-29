package evaio

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/shopspring/decimal"
)

func Test_tmp2(t *testing.T) {
	address := "evaio1z9k73l7trgshqpgg7m6hk9ehe4gphea5ch9dyh"
	c := NewClient("http://47.112.139.225:20001", false)

	for {
		height, _ := c.getBlockHeight()
		fmt.Println("height : ", height)
		balance, _ := c.getBalance(address, "uatom")
		fmt.Println("balance : ", balance.Balance.String())
	}
}
func Test_getBlockHeight(t *testing.T) {
	c := NewClient("http://18.163.89.47:20181", false)

	r, err := c.getBlockHeight()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}
}

func Test_getBlockByHash(t *testing.T) {
	hash := "3Uvb87ukKKwVeU6BFsZ21hy9sSbSd3Rd5QZTWbNop1d3TaY9ZzceJAT54vuY8XXQmw6nDx8ZViPV3cVznAHTtiVE"

	c := NewClient("http://localhost:9922/", false)

	r, err := c.Call("blocks/signature/"+hash, nil, "GET")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}
}

func Test_getBlockHash(t *testing.T) {
	c := NewClient("https://stargate.evaio.network", false)

	height := uint64(184952)

	r, err := c.getBlockHash(height)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}

}
func Test_tmp(t *testing.T) {
	test, err := time.Parse(time.RFC3339Nano, "2019-05-08T02:13:41.937681458Z")
	fmt.Println(err)
	fmt.Println(test.Unix())
}
func Test_getBalance(t *testing.T) {
	c := NewClient("https://stargate.evaio.net", false)

	address := "eva1pn80qt83wzk9w4gs3muc8hw26cexlgav75mar0"

	r, err := c.getBalance(address, "neva")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}

}

func Test_getTransaction(t *testing.T) {
	c := NewClient("https://stargate.evaio.net", false)
	txid := "23A170E1D868AFED47EE4FB3294335E8B8BECAD048486C54A310EA45EC1ABEBA" //"9KBoALfTjvZLJ6CAuJCGyzRA1aWduiNFMvbqTchfBVpF"

	path := "/txs/" + txid
	r, err := c.Call(path, nil, "GET")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}

	trx := NewTransaction(r, "auth/StdTx", "cosmos-sdk/MsgSend", "muon")

	fmt.Println(trx)
}

func Test_convert(t *testing.T) {

	amount := uint64(5000000001)

	amountStr := fmt.Sprintf("%d", amount)

	fmt.Println(amountStr)

	d, _ := decimal.NewFromString(amountStr)

	w, _ := decimal.NewFromString("100000000")

	d = d.Div(w)

	fmt.Println(d.String())

	d = d.Mul(w)

	fmt.Println(d.String())

	r, _ := strconv.ParseInt(d.String(), 10, 64)

	fmt.Println(r)

	fmt.Println(time.Now().UnixNano())
}

func Test_getTransactionByAddresses(t *testing.T) {
	addrs := "ARAA8AnUYa4kWwWkiZTTyztG5C6S9MFTx11"

	c := NewClient("http://localhost:9922/", false)
	result, err := c.getMultiAddrTransactions("auth/StdTx", "cosmos-sdk/MsgSend", "uatom", 0, -1, addrs)

	if err != nil {
		t.Error("get transactions failed!")
	} else {
		for _, tx := range result {
			fmt.Println(tx.TxID)
		}
	}
}

func Test_getBlockByHeight(t *testing.T) {
	height := uint64(429734)
	c := NewClient("http://127.0.0.1:1317", false)
	result, err := c.getBlockByHeight(height)
	if err != nil {
		t.Error("get block failed!")
	} else {
		fmt.Println(result)
	}
}

func Test_sequence(t *testing.T) {
	addr := "eva1pn80qt83wzk9w4gs3muc8hw26cexlgav75mar0"
	c := NewClient("http://47.91.232.118:8888", false)
	accountnumber, sequence, err := c.getAccountNumberAndSequence(addr)
	fmt.Println(err)
	fmt.Println(accountnumber)
	fmt.Println(sequence)
}
