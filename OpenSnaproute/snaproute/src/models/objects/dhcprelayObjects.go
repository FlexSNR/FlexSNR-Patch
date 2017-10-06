//
//Copyright [2016] [SnapRoute Inc]
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//	 Unless required by applicable law or agreed to in writing, software
//	 distributed under the License is distributed on an "AS IS" BASIS,
//	 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//	 See the License for the specific language governing permissions and
//	 limitations under the License.
//
// _______  __       __________   ___      _______.____    __    ____  __  .___________.  ______  __    __
// |   ____||  |     |   ____\  \ /  /     /       |\   \  /  \  /   / |  | |           | /      ||  |  |  |
// |  |__   |  |     |  |__   \  V  /     |   (----` \   \/    \/   /  |  | `---|  |----`|  ,----'|  |__|  |
// |   __|  |  |     |   __|   >   <       \   \      \            /   |  |     |  |     |  |     |   __   |
// |  |     |  `----.|  |____ /  .  \  .----)   |      \    /\    /    |  |     |  |     |  `----.|  |  |  |
// |__|     |_______||_______/__/ \__\ |_______/        \__/  \__/     |__|     |__|      \______||__|  |__|
//

package objects

import ()

type DHCPRelayGlobal struct {
	baseObj
	Vrf           string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"w", MULTIPLICITY:"1", AUTOCREATE: "true", DESCRIPTION: "VRF id for DHCPv4 Relay Agent global config", DEFAULT:"default"`
	Enable        bool   `DESCRIPTION: "Global level config for enabling/disabling the Relay Agent", DEFAULT:"false"`
	HopCountLimit int32  `DESCRIPTION: "Hop Count Limit", DEFAULT:"32"`
}

type DHCPRelayIntf struct {
	baseObj
	IntfRef  string   `SNAPROUTE: "KEY", CATEGORY:"L3",  ACCESS:"w", MULTIPLICITY:"*", DESCRIPTION:"DHCP Client facing interface reference for which Relay Agent needs to be configured"`
	Enable   bool     `DESCRIPTION: "Interface level config for enabling/disabling the relay agent"`
	ServerIp []string `DESCRIPTION: "DHCP Server(s) where relay agent can relay client dhcp requests"`
}

type DHCPRelayClientState struct {
	baseObj
	MacAddr         string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Host Hardware/Mac Address"`
	ServerIp        string `DESCRIPTION: "Configured DHCP Server"`
	OfferedIp       string `DESCRIPTION: "Ip Address offered by DHCP Server"`
	GatewayIp       string `DESCRIPTION: "Ip Address which client needs to use"`
	AcceptedIp      string `DESCRIPTION: "Ip Address which client accepted"`
	RequestedIp     string `DESCRIPTION: "Ip Address request from client"`
	ClientDiscover  string `DESCRIPTION: "Most recent time stamp of client discover message to dhcp server"`
	ClientRequest   string `DESCRIPTION: "Most recent time stamp of client request message"`
	ClientRequests  int32  `DESCRIPTION: "Total Number of client request message relayed to server"`
	ClientResponses int32  `DESCRIPTION: "Total Number of server response relayed to client"`
	ServerOffer     string `DESCRIPTION: "Most recent time stamp of server offer message"`
	ServerAck       string `DESCRIPTION: "Most recent time stamp of server ack message"`
	ServerRequests  int32  `DESCRIPTION: "Total Number of requests relayed to server"`
	ServerResponses int32  `DESCRIPTION: "Total Number of responses received from server"`
}

type DHCPRelayIntfState struct {
	baseObj
	IntfRef           string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Interface for which state is required to be collected"`
	TotalDrops        int32  `DESCRIPTION: "Total number of DHCP Packets dropped by relay agent"`
	TotalDhcpClientRx int32  `DESCRIPTION: "Total number of client requests that came to relay agent"`
	TotalDhcpClientTx int32  `DESCRIPTION: "Total number of client responses send out by relay agent"`
	TotalDhcpServerRx int32  `DESCRIPTION: "Total number of server requests made by relay agent"`
	TotalDhcpServerTx int32  `DESCRIPTION: "Total number of server responses received by relay agent"`
}

type DHCPRelayIntfServerState struct {
	baseObj
	IntfRef   string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"r", MULTIPLICITY:"1", DESCRIPTION: "Interface Index for which state is required to be collected"`
	ServerIp  string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"r", MULTIPLICITY:"1", DESCRIPTION: "Server IP on the interface for which state is required to be collected"`
	Request   int32  `DESCRIPTION: "Total number of requests to Server"`
	Responses int32  `DESCRIPTION: "Total number of responses from Server"`
}

type DHCPv6RelayGlobal struct {
	baseObj
	Vrf           string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"w", MULTIPLICITY:"1", AUTOCREATE: "true", DESCRIPTION: "VRF id for DHCPv6 Relay Agent global config", DEFAULT:"default"`
	Enable        bool   `DESCRIPTION: "Global level config for enabling/disabling the Relay Agent", DEFAULT:"false"`
	HopCountLimit int32  `DESCRIPTION: "Hop Count Limit", DEFAULT:"32"`
}

type DHCPv6RelayIntf struct {
	baseObj
	IntfRef       string   `SNAPROUTE: "KEY", CATEGORY:"L3",  ACCESS:"w", MULTIPLICITY:"*", DESCRIPTION:"DHCP Client facing interface reference for which Relay Agent needs to be configured""`
	Enable        bool     `DESCRIPTION: "Interface level config for enabling/disabling the relay agent"`
	ServerIp      []string `DESCRIPTION: "DHCP Server(s) where relay agent can relay client dhcp requests"`
	UpstreamIntfs []string `DESCRIPTION: "DHCP Server facing interfaces where Relay Forward messages are multicasted"`
}

type DHCPv6RelayClientState struct {
	baseObj
	MacAddr           string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Host Hardware/Mac Address"`
	ServerIp          string `DESCRIPTION: "Configured DHCP Server"`
	OfferedIp         string `DESCRIPTION: "Ip Address offered by DHCP Server"`
	GatewayIp         string `DESCRIPTION: "Ip Address which client needs to use"`
	AcceptedIp        string `DESCRIPTION: "Ip Address which client accepted"`
	RequestedIp       string `DESCRIPTION: "Ip Address request from client"`
	ClientSolicit     string `DESCRIPTION: "Most recent time stamp of client solicit message to dhcp server"`
	ClientRequest     string `DESCRIPTION: "Most recent time stamp of client request message"`
	ClientConfirm     string `DESCRIPTION: "Most recent time stamp of client confirm message to dhcp server"`
	ClientRenew       string `DESCRIPTION: "Most recent time stamp of client renew message to dhcp server"`
	ClientRebind      string `DESCRIPTION: "Most recent time stamp of client rebind message to dhcp server"`
	ClientRelease     string `DESCRIPTION: "Most recent time stamp of client release message"`
	ClientDecline     string `DESCRIPTION: "Most recent time stamp of client decline message"`
	ClientInfoRequest string `DESCRIPTION: "Most recent time stamp of client info-request message"`
	ClientRequests    int32  `DESCRIPTION: "Total Number of client request message relayed to server"`
	ClientResponses   int32  `DESCRIPTION: "Total Number of server response relayed to client"`
	ServerAdvertise   string `DESCRIPTION: "Most recent time stamp of server advertise message"`
	ServerReply       string `DESCRIPTION: "Most recent time stamp of server reply message"`
	ServerReconfigure string `DESCRIPTION: "Most recent time stamp of server reconfigure message"`
	ServerRequests    int32  `DESCRIPTION: "Total Number of requests relayed to server"`
	ServerResponses   int32  `DESCRIPTION: "Total Number of responses received from server"`
}

type DHCPv6RelayIntfState struct {
	baseObj
	IntfRef           string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Interface for which state is required to be collected"`
	TotalDrops        int32  `DESCRIPTION: "Total number of DHCP Packets dropped by relay agent"`
	TotalDhcpClientRx int32  `DESCRIPTION: "Total number of client requests that came to relay agent"`
	TotalDhcpClientTx int32  `DESCRIPTION: "Total number of client responses send out by relay agent"`
	TotalDhcpServerRx int32  `DESCRIPTION: "Total number of server requests made by relay agent"`
	TotalDhcpServerTx int32  `DESCRIPTION: "Total number of server responses received by relay agent"`
}

type DHCPv6RelayIntfServerState struct {
	baseObj
	IntfRef   string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"r", MULTIPLICITY:"1", DESCRIPTION: "Interface Index for which state is required to be collected"`
	ServerIp  string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"r", MULTIPLICITY:"1", DESCRIPTION: "Server IP on the interface for which state is required to be collected"`
	Request   int32  `DESCRIPTION: "Total number of requests to Server"`
	Responses int32  `DESCRIPTION: "Total number of responses from Server"`
}
