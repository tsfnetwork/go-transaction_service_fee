// Copyright 2014 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY+ADs- without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see +ADw-http://www.gnu.org/licenses/+AD4-.

package core

import (
	+ACI-bytes+ACI-
	+ACI-encoding/hex+ACI-
	+ACI-encoding/json+ACI-
	+ACI-errors+ACI-
	+ACI-fmt+ACI-
	+ACI-math/big+ACI-
	+ACI-strings+ACI-

	+ACI-github.com/tsfnetwork/go-transaction+AF8-service+AF8-fee/common+ACI-
	+ACI-github.com/tsfnetwork/go-transaction+AF8-service+AF8-fee/common/hexutil+ACI-
	+ACI-github.com/tsfnetwork/go-transaction+AF8-service+AF8-fee/common/math+ACI-
	+ACI-github.com/tsfnetwork/go-transaction+AF8-service+AF8-fee/core/state+ACI-
	+ACI-github.com/tsfnetwork/go-transaction+AF8-service+AF8-fee/core/types+ACI-
	+ACI-github.com/tsfnetwork/go-transaction+AF8-service+AF8-fee/ethdb+ACI-
	+ACI-github.com/tsfnetwork/go-transaction+AF8-service+AF8-fee/log+ACI-
	+ACI-github.com/tsfnetwork/go-transaction+AF8-service+AF8-fee/params+ACI-
	+ACI-github.com/tsfnetwork/go-transaction+AF8-service+AF8-fee/rlp+ACI-
)

//go:generate gencodec -type Genesis -field-override genesisSpecMarshaling -out gen+AF8-genesis.go
//go:generate gencodec -type GenesisAccount -field-override genesisAccountMarshaling -out gen+AF8-genesis+AF8-account.go

var errGenesisNoConfig +AD0- errors.New(+ACI-genesis has no chain configuration+ACI-)

// Genesis specifies the header fields, state of a genesis block. It also defines hard
// fork switch-over blocks through the chain configuration.
type Genesis struct +AHs-
	Config     +ACo-params.ChainConfig +AGA-json:+ACI-config+ACIAYA-
	Nonce      uint64              +AGA-json:+ACI-nonce+ACIAYA-
	Timestamp  uint64              +AGA-json:+ACI-timestamp+ACIAYA-
	ExtraData  +AFsAXQ-byte              +AGA-json:+ACI-extraData+ACIAYA-
	GasLimit   uint64              +AGA-json:+ACI-gasLimit+ACI-   gencodec:+ACI-required+ACIAYA-
	Difficulty +ACo-big.Int            +AGA-json:+ACI-difficulty+ACI- gencodec:+ACI-required+ACIAYA-
	Mixhash    common.Hash         +AGA-json:+ACI-mixHash+ACIAYA-
	Coinbase   common.Address      +AGA-json:+ACI-coinbase+ACIAYA-
	Alloc      GenesisAlloc        +AGA-json:+ACI-alloc+ACI-      gencodec:+ACI-required+ACIAYA-

	// These fields are used for consensus tests. Please don't use them
	// in actual genesis blocks.
	Number     uint64      +AGA-json:+ACI-number+ACIAYA-
	GasUsed    uint64      +AGA-json:+ACI-gasUsed+ACIAYA-
	ParentHash common.Hash +AGA-json:+ACI-parentHash+ACIAYA-
+AH0-

// GenesisAlloc specifies the initial state that is part of the genesis block.
type GenesisAlloc map+AFs-common.Address+AF0-GenesisAccount

func (ga +ACo-GenesisAlloc) UnmarshalJSON(data +AFsAXQ-byte) error +AHs-
	m :+AD0- make(map+AFs-common.UnprefixedAddress+AF0-GenesisAccount)
	if err :+AD0- json.Unmarshal(data, +ACY-m)+ADs- err +ACEAPQ- nil +AHs-
		return err
	+AH0-
	+ACo-ga +AD0- make(GenesisAlloc)
	for addr, a :+AD0- range m +AHs-
		(+ACo-ga)+AFs-common.Address(addr)+AF0- +AD0- a
	+AH0-
	return nil
+AH0-

// GenesisAccount is an account in the state of the genesis block.
type GenesisAccount struct +AHs-
	Code       +AFsAXQ-byte                      +AGA-json:+ACI-code,omitempty+ACIAYA-
	Storage    map+AFs-common.Hash+AF0-common.Hash +AGA-json:+ACI-storage,omitempty+ACIAYA-
	Balance    +ACo-big.Int                    +AGA-json:+ACI-balance+ACI- gencodec:+ACI-required+ACIAYA-
	Nonce      uint64                      +AGA-json:+ACI-nonce,omitempty+ACIAYA-
	PrivateKey +AFsAXQ-byte                      +AGA-json:+ACI-secretKey,omitempty+ACIAYA- // for tests
+AH0-

// field type overrides for gencodec
type genesisSpecMarshaling struct +AHs-
	Nonce      math.HexOrDecimal64
	Timestamp  math.HexOrDecimal64
	ExtraData  hexutil.Bytes
	GasLimit   math.HexOrDecimal64
	GasUsed    math.HexOrDecimal64
	Number     math.HexOrDecimal64
	Difficulty +ACo-math.HexOrDecimal256
	Alloc      map+AFs-common.UnprefixedAddress+AF0-GenesisAccount
+AH0-

type genesisAccountMarshaling struct +AHs-
	Code       hexutil.Bytes
	Balance    +ACo-math.HexOrDecimal256
	Nonce      math.HexOrDecimal64
	Storage    map+AFs-storageJSON+AF0-storageJSON
	PrivateKey hexutil.Bytes
+AH0-

// storageJSON represents a 256 bit byte array, but allows less than 256 bits when
// unmarshaling from hex.
type storageJSON common.Hash

func (h +ACo-storageJSON) UnmarshalText(text +AFsAXQ-byte) error +AHs-
	text +AD0- bytes.TrimPrefix(text, +AFsAXQ-byte(+ACI-0x+ACI-))
	if len(text) +AD4- 64 +AHs-
		return fmt.Errorf(+ACI-too many hex characters in storage key/value +ACU-q+ACI-, text)
	+AH0-
	offset :+AD0- len(h) - len(text)/2 // pad on the left
	if +AF8-, err :+AD0- hex.Decode(h+AFs-offset:+AF0-, text)+ADs- err +ACEAPQ- nil +AHs-
		fmt.Println(err)
		return fmt.Errorf(+ACI-invalid hex storage key/value +ACU-q+ACI-, text)
	+AH0-
	return nil
+AH0-

func (h storageJSON) MarshalText() (+AFsAXQ-byte, error) +AHs-
	return hexutil.Bytes(h+AFs-:+AF0-).MarshalText()
+AH0-

// GenesisMismatchError is raised when trying to overwrite an existing
// genesis block with an incompatible one.
type GenesisMismatchError struct +AHs-
	Stored, New common.Hash
+AH0-

func (e +ACo-GenesisMismatchError) Error() string +AHs-
	return fmt.Sprintf(+ACI-database already contains an incompatible genesis block (have +ACU-x, new +ACU-x)+ACI-, e.Stored+AFs-:8+AF0-, e.New+AFs-:8+AF0-)
+AH0-

// SetupGenesisBlock writes or updates the genesis block in db.
// The block that will be used is:
//
//                          genesis +AD0APQ- nil       genesis +ACEAPQ- nil
//                       +------------------------------------------
//     db has no genesis +AHw-  main-net default  +AHw-  genesis
//     db has genesis    +AHw-  from DB           +AHw-  genesis (if compatible)
//
// The stored chain configuration will be updated if it is compatible (i.e. does not
// specify a fork block below the local head block). In case of a conflict, the
// error is a +ACo-params.ConfigCompatError and the new, unwritten config is returned.
//
// The returned chain configuration is never nil.
func SetupGenesisBlock(db ethdb.Database, genesis +ACo-Genesis) (+ACo-params.ChainConfig, common.Hash, error) +AHs-
	if genesis +ACEAPQ- nil +ACYAJg- genesis.Config +AD0APQ- nil +AHs-
		return params.AllProtocolChanges, common.Hash+AHsAfQ-, errGenesisNoConfig
	+AH0-

	// Just commit the new block if there is no stored genesis block.
	stored :+AD0- GetCanonicalHash(db, 0)
	if (stored +AD0APQ- common.Hash+AHsAfQ-) +AHs-
		if genesis +AD0APQ- nil +AHs-
			log.Info(+ACI-Writing default main-net genesis block+ACI-)
			genesis +AD0- DefaultGenesisBlock()
		+AH0- else +AHs-
			log.Info(+ACI-Writing custom genesis block+ACI-)
		+AH0-
		block, err :+AD0- genesis.Commit(db)
		return genesis.Config, block.Hash(), err
	+AH0-

	// Check whether the genesis block is already written.
	if genesis +ACEAPQ- nil +AHs-
		block, +AF8- :+AD0- genesis.ToBlock()
		hash :+AD0- block.Hash()
		if hash +ACEAPQ- stored +AHs-
			return genesis.Config, block.Hash(), +ACY-GenesisMismatchError+AHs-stored, hash+AH0-
		+AH0-
	+AH0-

	// Get the existing chain configuration.
	newcfg :+AD0- genesis.configOrDefault(stored)
	storedcfg, err :+AD0- GetChainConfig(db, stored)
	if err +ACEAPQ- nil +AHs-
		if err +AD0APQ- ErrChainConfigNotFound +AHs-
			// This case happens if a genesis write was interrupted.
			log.Warn(+ACI-Found genesis block without chain config+ACI-)
			err +AD0- WriteChainConfig(db, stored, newcfg)
		+AH0-
		return newcfg, stored, err
	+AH0-
	// Special case: don't change the existing config of a non-mainnet chain if no new
	// config is supplied. These chains would get AllProtocolChanges (and a compat error)
	// if we just continued here.
	if genesis +AD0APQ- nil +ACYAJg- stored +ACEAPQ- params.MainnetGenesisHash +AHs-
		return storedcfg, stored, nil
	+AH0-

	// Check config compatibility and write the config. Compatibility errors
	// are returned to the caller unless we're already at block zero.
	height :+AD0- GetBlockNumber(db, GetHeadHeaderHash(db))
	if height +AD0APQ- missingNumber +AHs-
		return newcfg, stored, fmt.Errorf(+ACI-missing block number for head header hash+ACI-)
	+AH0-
	compatErr :+AD0- storedcfg.CheckCompatible(newcfg, height)
	if compatErr +ACEAPQ- nil +ACYAJg- height +ACEAPQ- 0 +ACYAJg- compatErr.RewindTo +ACEAPQ- 0 +AHs-
		return newcfg, stored, compatErr
	+AH0-
	return newcfg, stored, WriteChainConfig(db, stored, newcfg)
+AH0-

func (g +ACo-Genesis) configOrDefault(ghash common.Hash) +ACo-params.ChainConfig +AHs-
	switch +AHs-
	case g +ACEAPQ- nil:
		return g.Config
	case ghash +AD0APQ- params.MainnetGenesisHash:
		return params.MainnetChainConfig
	case ghash +AD0APQ- params.TestnetGenesisHash:
		return params.TestnetChainConfig
	default:
		return params.AllProtocolChanges
	+AH0-
+AH0-

// ToBlock creates the block and state of a genesis specification.
func (g +ACo-Genesis) ToBlock() (+ACo-types.Block, +ACo-state.StateDB) +AHs-
	db, +AF8- :+AD0- ethdb.NewMemDatabase()
	statedb, +AF8- :+AD0- state.New(common.Hash+AHsAfQ-, state.NewDatabase(db))
	for addr, account :+AD0- range g.Alloc +AHs-
		statedb.AddBalance(addr, account.Balance)
		statedb.SetCode(addr, account.Code)
		statedb.SetNonce(addr, account.Nonce)
		for key, value :+AD0- range account.Storage +AHs-
			statedb.SetState(addr, key, value)
		+AH0-
	+AH0-
	root :+AD0- statedb.IntermediateRoot(false)
	head :+AD0- +ACY-types.Header+AHs-
		Number:     new(big.Int).SetUint64(g.Number),
		Nonce:      types.EncodeNonce(g.Nonce),
		Time:       new(big.Int).SetUint64(g.Timestamp),
		ParentHash: g.ParentHash,
		Extra:      g.ExtraData,
		GasLimit:   new(big.Int).SetUint64(g.GasLimit),
		GasUsed:    new(big.Int).SetUint64(g.GasUsed),
		Difficulty: g.Difficulty,
		MixDigest:  g.Mixhash,
		Coinbase:   g.Coinbase,
		Root:       root,
	+AH0-
	if g.GasLimit +AD0APQ- 0 +AHs-
		head.GasLimit +AD0- params.GenesisGasLimit
	+AH0-
	if g.Difficulty +AD0APQ- nil +AHs-
		head.Difficulty +AD0- params.GenesisDifficulty
	+AH0-
	return types.NewBlock(head, nil, nil, nil), statedb
+AH0-

// Commit writes the block and state of a genesis specification to the database.
// The block is committed as the canonical head block.
func (g +ACo-Genesis) Commit(db ethdb.Database) (+ACo-types.Block, error) +AHs-
	block, statedb :+AD0- g.ToBlock()
	if block.Number().Sign() +ACEAPQ- 0 +AHs-
		return nil, fmt.Errorf(+ACI-can't commit genesis block with number +AD4- 0+ACI-)
	+AH0-
	if +AF8-, err :+AD0- statedb.CommitTo(db, false)+ADs- err +ACEAPQ- nil +AHs-
		return nil, fmt.Errorf(+ACI-cannot write state: +ACU-v+ACI-, err)
	+AH0-
	if err :+AD0- WriteTd(db, block.Hash(), block.NumberU64(), g.Difficulty)+ADs- err +ACEAPQ- nil +AHs-
		return nil, err
	+AH0-
	if err :+AD0- WriteBlock(db, block)+ADs- err +ACEAPQ- nil +AHs-
		return nil, err
	+AH0-
	if err :+AD0- WriteBlockReceipts(db, block.Hash(), block.NumberU64(), nil)+ADs- err +ACEAPQ- nil +AHs-
		return nil, err
	+AH0-
	if err :+AD0- WriteCanonicalHash(db, block.Hash(), block.NumberU64())+ADs- err +ACEAPQ- nil +AHs-
		return nil, err
	+AH0-
	if err :+AD0- WriteHeadBlockHash(db, block.Hash())+ADs- err +ACEAPQ- nil +AHs-
		return nil, err
	+AH0-
	if err :+AD0- WriteHeadHeaderHash(db, block.Hash())+ADs- err +ACEAPQ- nil +AHs-
		return nil, err
	+AH0-
	config :+AD0- g.Config
	if config +AD0APQ- nil +AHs-
		config +AD0- params.AllProtocolChanges
	+AH0-
	return block, WriteChainConfig(db, block.Hash(), config)
+AH0-

// MustCommit writes the genesis block and state to db, panicking on error.
// The block is committed as the canonical head block.
func (g +ACo-Genesis) MustCommit(db ethdb.Database) +ACo-types.Block +AHs-
	block, err :+AD0- g.Commit(db)
	if err +ACEAPQ- nil +AHs-
		panic(err)
	+AH0-
	return block
+AH0-

// GenesisBlockForTesting creates and writes a block in which addr has the given wei balance.
func GenesisBlockForTesting(db ethdb.Database, addr common.Address, balance +ACo-big.Int) +ACo-types.Block +AHs-
	g :+AD0- Genesis+AHs-Alloc: GenesisAlloc+AHs-addr: +AHs-Balance: balance+AH0AfQB9-
	return g.MustCommit(db)
+AH0-

// DefaultGenesisBlock returns the Ethereum main net genesis block.
func DefaultGenesisBlock() +ACo-Genesis +AHs-
	return +ACY-Genesis+AHs-
		Config:     params.MainnetChainConfig,
		Nonce:      0x242c819c166b9b13,
		ExtraData:  hexutil.MustDecode(+ACI-0x00000000000000000000000000000000+ACI-),
		GasLimit:   3141592,
		Difficulty: big.NewInt(34359738368),
		Coinbase:   common.HexToAddress(+ACI-0x4ad372afa90a0d828166eb3bed9552707757a3d0+ACI-),
		Alloc:      decodePrealloc(mainnetAllocData),
	+AH0-
+AH0-

// DefaultTestnetGenesisBlock returns the Ropsten network genesis block.
func DefaultTestnetGenesisBlock() +ACo-Genesis +AHs-
	return +ACY-Genesis+AHs-
		Config:     params.TestnetChainConfig,
		Nonce:      66,
		ExtraData:  hexutil.MustDecode(+ACI-0x3535353535353535353535353535353535353535353535353535353535353535+ACI-),
		GasLimit:   16777216,
		Difficulty: big.NewInt(1048576),
		Alloc:      decodePrealloc(testnetAllocData),
	+AH0-
+AH0-

// DefaultRinkebyGenesisBlock returns the Rinkeby network genesis block.
func DefaultRinkebyGenesisBlock() +ACo-Genesis +AHs-
	return +ACY-Genesis+AHs-
		Config:     params.RinkebyChainConfig,
		Timestamp:  1492009146,
		ExtraData:  hexutil.MustDecode(+ACI-0x52657370656374206d7920617574686f7269746168207e452e436172746d616e42eb768f2244c8811c63729a21a3569731535f067ffc57839b00206d1ad20c69a1981b489f772031b279182d99e65703f0076e4812653aab85fca0f00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000+ACI-),
		GasLimit:   4700000,
		Difficulty: big.NewInt(1),
		Alloc:      decodePrealloc(rinkebyAllocData),
	+AH0-
+AH0-

// DevGenesisBlock returns the 'geth --dev' genesis block.
func DevGenesisBlock() +ACo-Genesis +AHs-
	return +ACY-Genesis+AHs-
		Config:     params.AllProtocolChanges,
		Nonce:      42,
		GasLimit:   4712388,
		Difficulty: big.NewInt(131072),
		Alloc:      decodePrealloc(devAllocData),
	+AH0-
+AH0-

func decodePrealloc(data string) GenesisAlloc +AHs-
	var p +AFsAXQ-struct+AHs- Addr, Balance +ACo-big.Int +AH0-
	if err :+AD0- rlp.NewStream(strings.NewReader(data), 0).Decode(+ACY-p)+ADs- err +ACEAPQ- nil +AHs-
		panic(err)
	+AH0-
	ga :+AD0- make(GenesisAlloc, len(p))
	for +AF8-, account :+AD0- range p +AHs-
		ga+AFs-common.BigToAddress(account.Addr)+AF0- +AD0- GenesisAccount+AHs-Balance: account.Balance+AH0-
	+AH0-
	return ga
+AH0-