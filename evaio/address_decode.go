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

package evaio

import (
	"github.com/blocktree/go-owcdrivers/addressEncoder"
	"github.com/blocktree/go-owcdrivers/evaioTransaction"
	owcrypt "github.com/blocktree/go-owcrypt"
	"github.com/blocktree/openwallet/v2/openwallet"
)


type AddressDecoderV2 struct {

	openwallet.AddressDecoderV2Base
	//ScriptPubKeyToBech32Address(scriptPubKey []byte) (string, error)
}
type addressDecoder struct {
	wm *WalletManager //钱包管理者
}

//NewAddressDecoder 地址解析器
func NewAddressDecoderV2(wm *WalletManager) *AddressDecoderV2 {
	decoder := AddressDecoderV2{}
	return &decoder
}


//AddressDecode 地址解析
func (dec *AddressDecoderV2) AddressDecode(addr string, opts ...interface{}) ([]byte, error) {
	decodeHash, err := evaioTransaction.Bech32Decode(addr)
	if err != nil {
		return nil, err
	}
	return decodeHash, nil
}

//AddressEncode 地址编码
func (dec *AddressDecoderV2) AddressEncode(hash []byte, opts ...interface{}) (string, error) {

	cfg := addressEncoder.EVA_mainnetAddress

	pkHash := owcrypt.Hash(hash, 32, owcrypt.HASH_ALG_HASH160)

	address := addressEncoder.AddressEncode(pkHash, cfg)
	return address, nil
}

// AddressVerify 地址校验
func (dec *AddressDecoderV2) AddressVerify(address string, opts ...interface{}) bool {
	_, err := evaioTransaction.Bech32Decode(address)
	if err != nil {
		return false
	}
	return true
}
