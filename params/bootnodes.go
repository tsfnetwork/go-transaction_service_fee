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
	"enode://810c9a99dd53f0e8136432f31cfcf1f5925fe8dff4adfa53c462591516447ffc4d0b16e7b954e0b9d1a312c67876d8a16e563b29afe587596d28bd6ff0225e26@167.99.192.153:59997",
	"enode://10a4991c6b5b5eb831a07e7c35724d8266893c7986b526729348d1bf37b37a61c19c895013065e348b124b6431d7c9a9e8e970b3786eb95af2f921e7f93514f0@46.101.187.225:59997",
	"enode://f70c7ea4bd48dc7116a95b87b6e42fe61f0dc5f8a70e7157f68c99d13f393ec1bb3a0089701f6657a861e1e174e1aab650bc585246ac429b9cbe84b8077930db@165.232.99.209:59997",
	"enode://a126d9ce098720d00435745cfec981c6da1b2d04a65401639a36873ee1262ce9525ddc76b5668ec5aa8369e98f54d7814ad4b49b73a26fe516eb8aa1cf7a6b1e@167.71.54.120:59997",
	"enode://891aa70a41e8410aa9c6a10a3e18f53bb8d31283d3e40cde0f66dc4dab506e7e6f415e76651ce66fdb98af0b126b560f7e44d5e0a6e3a6cc920b3192bba231f4@68.183.214.221:59997",
	"enode://ee4a8ceeb6c68c4a62b4f97f2ff6b2f7d80a89d6527310d74d811f79299cafea8a71cc8370aa481de8d5f6ed2b0c01a897f07e42ab5e72b6d327b42927004c91@167.172.97.130:59997",
	"enode://666d2daffc08ddae814a799946637a521a830ff2e1b327119e315507f546c174d27013d94a690b1ffff846ea437d667c7e92dd85440e0b7719952056109f7c3b@167.172.174.6:59997",

}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
    "enode://9fb686d265a7d6e3c5cd4de44e8b50eee40a618fed41ac6640174225a8bdaa1b7248f92369b9299e4576be8a532eae601353114b8443f274f7263a3e3d0b4c50@104.248.32.50:59997",
	"enode://d43e62049e8ca3028c6479ae408cc614f07f84d813947988c9a71b463a99e891746433deb9b03e1950b1bd091612d0cab624a45dcca07f4ecd7f05cfeb04faee@134.122.87.200:59997",
	
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
