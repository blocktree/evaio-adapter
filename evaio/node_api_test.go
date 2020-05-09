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
	c := NewClient("https://stargate.evaio.net", false)

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
	c := NewClient("https://stargate.evaio.net", false)

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



/*
2020/01/02 10:08:27.995 [I] [manager_test.go:142]  address[ 0 ] : eva1dqhtv85u4haxs73x8nntqttpy62k658hwev5k3 440
2020/01/02 10:08:27.995 [I] [manager_test.go:142]  address[ 1 ] : eva1ks3qnkq8753stprwvthvppnd57jw20n739ql2k 390
2020/01/02 10:08:27.995 [I] [manager_test.go:142]  address[ 2 ] : eva1pn80qt83wzk9w4gs3muc8hw26cexlgav75mar0 30
2020/01/02 10:08:27.995 [I] [manager_test.go:142]  address[ 3 ] : eva1td0y392rcufal8vdl8pa9cjnkxjyzm8qryqwz9 121
*/
func Test_getBalance(t *testing.T) {
	c := NewClient("https://stargate.evaio.net", false)

	address := "eva17hyme9cqufyqhc7ywzs7v6tv3xg5v5e7fu60u7"

	r, err := c.getBalance(address, "neva")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}

}

func Test_getTransaction(t *testing.T) {
	c := NewClient("https://stargate.evaio.net", false)
	txid := "4D1124D32BFB8F724F3570BD02B264F8BA22BB3E0D906DBCFA3E6EBBC7301B41" //"9KBoALfTjvZLJ6CAuJCGyzRA1aWduiNFMvbqTchfBVpF"

	path := "/txs/" + txid
	r, err := c.Call(path, nil, "GET")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}

	trx := NewTransaction(r, "cosmos-sdk/StdTx", "cosmos-sdk/MsgSend", "neva")

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
	addr := "eva1td0y392rcufal8vdl8pa9cjnkxjyzm8qryqwz9"
	c := NewClient("https://stargate.evaio.net", false)
	accountnumber, sequence, err := c.getAccountNumberAndSequence(addr)
	fmt.Println(err)
	fmt.Println(accountnumber)
	fmt.Println(sequence)
}
