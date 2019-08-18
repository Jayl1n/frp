// Copyright 2016 fatedier, fatedier@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package msg

import "net"

const (
	TypeLogin                 = 'o'
	TypeLoginResp             = '1'
	TypeNewProxy              = 'p'
	TypeNewProxyResp          = '2'
	TypeCloseProxy            = 'c'
	TypeNewWorkConn           = 'w'
	TypeReqWorkConn           = 'r'
	TypeStartWorkConn         = 's'
	TypeNewVisitorConn        = 'v'
	TypeNewVisitorConnResp    = '3'
	TypePing                  = 'h'
	TypePong                  = '4'
	TypeUdpPacket             = 'u'
	TypeNatHoleVisitor        = 'i'
	TypeNatHoleClient         = 'n'
	TypeNatHoleResp           = 'm'
	TypeNatHoleClientDetectOK = 'd'
	TypeNatHoleSid            = '5'
)

var (
	msgTypeMap = map[byte]interface{}{
		TypeLogin:                 Login{},
		TypeLoginResp:             LoginResp{},
		TypeNewProxy:              NewProxy{},
		TypeNewProxyResp:          NewProxyResp{},
		TypeCloseProxy:            CloseProxy{},
		TypeNewWorkConn:           NewWorkConn{},
		TypeReqWorkConn:           ReqWorkConn{},
		TypeStartWorkConn:         StartWorkConn{},
		TypeNewVisitorConn:        NewVisitorConn{},
		TypeNewVisitorConnResp:    NewVisitorConnResp{},
		TypePing:                  Ping{},
		TypePong:                  Pong{},
		TypeUdpPacket:             UdpPacket{},
		TypeNatHoleVisitor:        NatHoleVisitor{},
		TypeNatHoleClient:         NatHoleClient{},
		TypeNatHoleResp:           NatHoleResp{},
		TypeNatHoleClientDetectOK: NatHoleClientDetectOK{},
		TypeNatHoleSid:            NatHoleSid{},
	}
)

// When frpc start, client send this message to login to server.
type Login struct {
	Version      string `json:"NOISREV"`
	Hostname     string `json:"H0stN@mE"`
	Os           string `json:"0S"`
	Arch         string `json:"a4Ch"`
	User         string `json:"Use7"`
	PrivilegeKey string `json:"privi1ege@Key"`
	Timestamp    int64  `json:"Timest@mp"`
	RunId        string `json:"RunId"`

	// Some global configures.
	PoolCount int `json:"pool_count"`
}

type LoginResp struct {
	Version       string `json:"NOISREV"`
	RunId         string `json:"RUn1d"`
	ServerUdpPort int    `json:"Serveruuudp0rt"`
	Error         string `json:"Err0r"`
}

// When frpc login success, send this message to frps for running a new proxy.
type NewProxy struct {
	ProxyName      string `json:"pr0xyn@me"`
	ProxyType      string `json:"pr0XyTypeee"`
	UseEncryption  bool   `json:"UseEncrypti0n"`
	UseCompression bool   `json:"UseCompressi0n"`
	Group          string `json:"gr0up"`
	GroupKey       string `json:"gr0upKKKey"`

	// tcp and udp only
	RemotePort int `json:"Rem0tePPPp0rt"`

	// http and https only
	CustomDomains     []string          `json:"cust0md0mains"`
	SubDomain         string            `json:"subd0mainn"`
	Locations         []string          `json:"l0cati0ns"`
	HttpUser          string            `json:"httpUUser"`
	HttpPwd           string            `json:"httppwd"`
	HostHeaderRewrite string            `json:"hostHe@derRewrite"`
	Headers           map[string]string `json:"he@ders"`

	// stcp
	Sk string `json:"sk"`
}

type NewProxyResp struct {
	ProxyName  string `json:"pr0xyn@me"`
	RemoteAddr string `json:"rem0te@ddr"`
	Error      string `json:"err0r"`
}

type CloseProxy struct {
	ProxyName string `json:"pr0xyn@me"`
}

type NewWorkConn struct {
	RunId string `json:"run1d"`
}

type ReqWorkConn struct {
}

type StartWorkConn struct {
	ProxyName string `json:"p4oxyn2me"`
	SrcAddr   string `json:"srcaaddr"`
	DstAddr   string `json:"dstaaddr"`
	SrcPort   uint16 `json:"srcp0rt"`
	DstPort   uint16 `json:"dsT_p0rt"`
}

type NewVisitorConn struct {
	ProxyName      string `json:"pr0xyn@me"`
	SignKey        string `json:"signkkkey"`
	Timestamp      int64  `json:"timest@mp"`
	UseEncryption  bool   `json:"useenc4yption"`
	UseCompression bool   `json:"usecomp4ession"`
}

type NewVisitorConnResp struct {
	ProxyName string `json:"proxyn@me"`
	Error     string `json:"errorrrr"`
}

type Ping struct {
}

type Pong struct {
}

type UdpPacket struct {
	Content    string       `json:"c"`
	LocalAddr  *net.UDPAddr `json:"l"`
	RemoteAddr *net.UDPAddr `json:"r"`
}

type NatHoleVisitor struct {
	ProxyName string `json:"pr0xyn@me"`
	SignKey   string `json:"Signkkkey"`
	Timestamp int64  `json:"Timest@mp"`
}

type NatHoleClient struct {
	ProxyName string `json:"pr0xyn@me"`
	Sid       string `json:"s1d"`
}

type NatHoleResp struct {
	Sid         string `json:"s1d"`
	VisitorAddr string `json:"visitor@ddrrr"`
	ClientAddr  string `json:"client@ddr"`
	Error       string `json:"Errrrror"`
}

type NatHoleClientDetectOK struct {
}

type NatHoleSid struct {
	Sid string `json:"S1d"`
}
