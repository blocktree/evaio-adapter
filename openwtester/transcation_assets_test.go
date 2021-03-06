/*
 * Copyright 2018 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package openwtester

import (
	"fmt"
	"testing"

	"github.com/blocktree/evaio-adapter/evaio"
	"github.com/blocktree/openwallet/v2/openw"

	"github.com/blocktree/openwallet/v2/log"
	"github.com/blocktree/openwallet/v2/openwallet"
)

func testGetAssetsAccountBalance(tm *openw.WalletManager, walletID, accountID string) {
	balance, err := tm.GetAssetsAccountBalance(testApp, walletID, accountID)
	if err != nil {
		log.Error("GetAssetsAccountBalance failed, unexpected error:", err)
		return
	}
	log.Info("balance:", balance)
}

func testGetAssetsAccountTokenBalance(tm *openw.WalletManager, walletID, accountID string, contract openwallet.SmartContract) {
	balance, err := tm.GetAssetsAccountTokenBalance(testApp, walletID, accountID, contract)
	if err != nil {
		log.Error("GetAssetsAccountTokenBalance failed, unexpected error:", err)
		return
	}
	log.Info("token balance:", balance.Balance)
}

func testCreateTransactionStep(tm *openw.WalletManager, walletID, accountID, to, amount, feeRate string, contract *openwallet.SmartContract) (*openwallet.RawTransaction, error) {

	//err := tm.RefreshAssetsAccountBalance(testApp, accountID)
	//if err != nil {
	//	log.Error("RefreshAssetsAccountBalance failed, unexpected error:", err)
	//	return nil, err
	//}

	rawTx, err := tm.CreateTransaction(testApp, walletID, accountID, amount, to, feeRate, "", contract)

	if err != nil {
		log.Error("CreateTransaction failed, unexpected error:", err)
		return nil, err
	}

	return rawTx, nil
}

func testCreateSummaryTransactionStep(
	tm *openw.WalletManager,
	walletID, accountID, summaryAddress, minTransfer, retainedBalance, feeRate string,
	start, limit int,
	contract *openwallet.SmartContract) ([]*openwallet.RawTransactionWithError, error) {

	rawTxArray, err := tm.CreateSummaryRawTransactionWithError(testApp, walletID, accountID, summaryAddress, minTransfer,
		retainedBalance, feeRate, start, limit, contract, nil)

	if err != nil {
		log.Error("CreateSummaryTransaction failed, unexpected error:", err)
		return nil, err
	}

	return rawTxArray, nil
}

func testSignTransactionStep(tm *openw.WalletManager, rawTx *openwallet.RawTransaction) (*openwallet.RawTransaction, error) {

	_, err := tm.SignTransaction(testApp, rawTx.Account.WalletID, rawTx.Account.AccountID, "12345678", rawTx)
	if err != nil {
		log.Error("SignTransaction failed, unexpected error:", err)
		return nil, err
	}

	log.Infof("rawTx: %+v", rawTx)
	return rawTx, nil
}

func testVerifyTransactionStep(tm *openw.WalletManager, rawTx *openwallet.RawTransaction) (*openwallet.RawTransaction, error) {

	//log.Info("rawTx.Signatures:", rawTx.Signatures)

	_, err := tm.VerifyTransaction(testApp, rawTx.Account.WalletID, rawTx.Account.AccountID, rawTx)
	if err != nil {
		log.Error("VerifyTransaction failed, unexpected error:", err)
		return nil, err
	}

	log.Infof("rawTx: %+v", rawTx)
	return rawTx, nil
}

func testSubmitTransactionStep(tm *openw.WalletManager, rawTx *openwallet.RawTransaction) (*openwallet.RawTransaction, error) {

	tx, err := tm.SubmitTransaction(testApp, rawTx.Account.WalletID, rawTx.Account.AccountID, rawTx)
	if err != nil {
		log.Error("SubmitTransaction failed, unexpected error:", err)
		return nil, err
	}

	log.Std.Info("tx: %+v", tx)
	log.Info("wxID:", tx.WxID)
	log.Info("txID:", rawTx.TxID)

	return rawTx, nil
}

func TestTransfer(t *testing.T) {
	tm := testInitWalletManager()
	walletID := "W6oBg6JSaoEMs46u2tfWh4rreu2Dwt8Cf6"
	//accountID := "E1Yea5xyvHk4EUcyex3aNkgPSpAs8R3mSjNC23qeqj4r"
	accountID := "6Sasy5wG4bHt4KbrVeevsRmk3BrmZvMR9u9ogUtgzL4Z"
	to := "eva18jg8gs9mc2gguwcw73csm03amqkspk6nhmr0qr"

	testGetAssetsAccountBalance(tm, walletID, accountID)

	//for index := 0; index < 4; index++ {

	rawTx, err := testCreateTransactionStep(tm, walletID, accountID, to, "20", "", nil)
	if err != nil {
		return
	}

	log.Std.Info("rawTx: %+v", rawTx)

	_, err = testSignTransactionStep(tm, rawTx)
	if err != nil {
		return
	}

	_, err = testVerifyTransactionStep(tm, rawTx)
	if err != nil {
		return
	}

	_, err = testSubmitTransactionStep(tm, rawTx)
	if err != nil {
		return
	}
	//}
	// rawTx, err = testCreateTransactionStep(tm, walletID, accountID, to, "0.02", "", nil)
	// if err != nil {
	// 	return
	// }

	// log.Std.Info("rawTx: %+v", rawTx)

	// _, err = testSignTransactionStep(tm, rawTx)
	// if err != nil {
	// 	return
	// }

	// _, err = testVerifyTransactionStep(tm, rawTx)
	// if err != nil {
	// 	return
	// }

	// _, err = testSubmitTransactionStep(tm, rawTx)
	// if err != nil {
	// 	return
	// }

	// rawTx, err = testCreateTransactionStep(tm, walletID, accountID, to, "0.06", "", nil)
	// if err != nil {
	// 	return
	// }

	// log.Std.Info("rawTx: %+v", rawTx)

	// _, err = testSignTransactionStep(tm, rawTx)
	// if err != nil {
	// 	return
	// }

	// _, err = testVerifyTransactionStep(tm, rawTx)
	// if err != nil {
	// 	return
	// }

	// _, err = testSubmitTransactionStep(tm, rawTx)
	// if err != nil {
	// 	return
	// }

	// rawTx, err = testCreateTransactionStep(tm, walletID, accountID, to, "0.06", "", nil)
	// if err != nil {
	// 	return
	// }

	// log.Std.Info("rawTx: %+v", rawTx)

	// _, err = testSignTransactionStep(tm, rawTx)
	// if err != nil {
	// 	return
	// }

	// _, err = testVerifyTransactionStep(tm, rawTx)
	// if err != nil {
	// 	return
	// }

	// _, err = testSubmitTransactionStep(tm, rawTx)
	// if err != nil {
	// 	return
	// }
	//	getdata("evaio1z9k73l7trgshqpgg7m6hk9ehe4gphea5ch9dyh")

}
func getdata(addr string) {
	c := evaio.NewClient("http://47.112.139.225:20001", false)
	path := "/auth/accounts/" + addr
	for {
		r, _ := c.Call(path, nil, "GET")

		accountNumber := int(r.Get("value").Get("account_number").Uint())
		sequence := int(r.Get("value").Get("sequence").Uint())
		fmt.Println("accountNumber : ", accountNumber)
		fmt.Println("sequence : ", sequence)

		resp, _ := c.Call("/blocks/latest", nil, "GET")

		height := resp.Get("block_meta").Get("header").Get("height").Uint()
		fmt.Println("height : ", height)
	}

}

func TestSummary(t *testing.T) {
	tm := testInitWalletManager()
	walletID := "W6oBg6JSaoEMs46u2tfWh4rreu2Dwt8Cf6"
	accountID := "E1Yea5xyvHk4EUcyex3aNkgPSpAs8R3mSjNC23qeqj4r"
	summaryAddress := "eva17hyme9cqufyqhc7ywzs7v6tv3xg5v5e7fu60u7"

	testGetAssetsAccountBalance(tm, walletID, accountID)

	rawTxArray, err := testCreateSummaryTransactionStep(tm, walletID, accountID,
		summaryAddress, "", "", "",
		0, 100, nil)
	if err != nil {
		log.Errorf("CreateSummaryTransaction failed, unexpected error: %v", err)
		return
	}

	//执行汇总交易
	for _, rawTxWithErr := range rawTxArray {

		if rawTxWithErr.Error != nil {
			log.Error(rawTxWithErr.Error.Error())
			continue
		}

		_, err = testSignTransactionStep(tm, rawTxWithErr.RawTx)
		if err != nil {
			return
		}

		_, err = testVerifyTransactionStep(tm, rawTxWithErr.RawTx)
		if err != nil {
			return
		}

		_, err = testSubmitTransactionStep(tm, rawTxWithErr.RawTx)
		if err != nil {
			return
		}
	}

}
