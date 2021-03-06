// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// TSF/DEV Go Bootnodes 
	"enode://853c50c949bcf4cf6c5371c11f8fc9adff6e2a2533074c0a760125cbc91ea30b52da06951d38865012ff7d8a30b6dcb50b409a34848201a92417d4ee3a3203ec@185.141.62.215:59997",
	"enode://10a4991c6b5b5eb831a07e7c35724d8266893c7986b526729348d1bf37b37a61c19c895013065e348b124b6431d7c9a9e8e970b3786eb95af2f921e7f93514f0@46.101.187.225:59997",
	"enode://f70c7ea4bd48dc7116a95b87b6e42fe61f0dc5f8a70e7157f68c99d13f393ec1bb3a0089701f6657a861e1e174e1aab650bc585246ac429b9cbe84b8077930db@165.232.99.209:59997",
	"enode://6f490097cfe0f49f5d311a86038dda4b259dfe889b11f63bffefbb8c7e81851ebf941ca7c875aae2fa28864a2c0611affd72212b278df74210db1bbf3f62a534@94.156.189.141:59997",
	"enode://37317e7a885aeaf703a4a38224a4c3bf6bb9df45acb9f6a409f30296dd58c5e7b28b908e534c85631a9a91544e0e69fb8d522496d5cf58a879275a459d4ff665@185.206.146.57:59997",
	"enode://a3cb002422cd9496a7c156940ee1b4caccbbd3cd7e34204ddb2cebdb72f5dd91b699082bd823f61bf05bc62b5c0ecfc2fc458fd4d20e5de6bfe8644d3212a275@78.46.67.207:59997",
	"enode://853c50c949bcf4cf6c5371c11f8fc9adff6e2a2533074c0a760125cbc91ea30b52da06951d38865012ff7d8a30b6dcb50b409a34848201a92417d4ee3a3203ec@78.46.67.207:59997",
	

}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
     "enode://9fb686d265a7d6e3c5cd4de44e8b50eee40a618fed41ac6640174225a8bdaa1b7248f92369b9299e4576be8a532eae601353114b8443f274f7263a3e3d0b4c50@104.248.32.50:59997",
	
	
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	
}

// RinkebyV5Bootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network for the experimental RLPx v5 topic-discovery network.
var RinkebyV5Bootnodes = []string{
	
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{

}
