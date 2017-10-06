//
//Copyright [2016] [SnapRoute Inc]
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//       Unless required by applicable law or agreed to in writing, software
//       distributed under the License is distributed on an "AS IS" BASIS,
//       WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//       See the License for the specific language governing permissions and
//       limitations under the License.
//
// _______  __       __________   ___      _______.____    __    ____  __  .___________.  ______  __    __
// |   ____||  |     |   ____\  \ /  /     /       |\   \  /  \  /   / |  | |           | /      ||  |  |  |
// |  |__   |  |     |  |__   \  V  /     |   (----` \   \/    \/   /  |  | `---|  |----`|  ,----'|  |__|  |
// |   __|  |  |     |   __|   >   <       \   \      \            /   |  |     |  |     |  |     |   __   |
// |  |     |  `----.|  |____ /  .  \  .----)   |      \    /\    /    |  |     |  |     |  `----.|  |  |  |
// |__|     |_______||_______/__/ \__\ |_______/        \__/  \__/     |__|     |__|      \______||__|  |__|
//

package objects

type Ospfv2Global struct {
	ConfigObj
	Vrf                string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"w", MULTIPLICITY:"1", AUTOCREATE: "true", DESCRIPTION: "VRF id for OSPF global config", DEFAULT:"default"`
	RouterId           string `DESCRIPTION: A 32-bit integer uniquely identifying the router in the Autonomous System. By convention, to ensure uniqueness, this should default to the value of one of the router's IP interface addresses.  This object is persistent and when written the entity SHOULD save the change to non-volatile storage., DEFAULT:"0.0.0.0"`
	AdminState         string `DESCRIPTION: Indicates if OSPF is enabled globally., DEFAULT:"DOWN"`
	ASBdrRtrStatus     bool   `DESCRIPTION: A flag to note whether this router is configured as an Autonomous System Border Router.  This object is persistent and when written the entity SHOULD save the change to non-volatile storage., DEFAULT:false`
	ReferenceBandwidth uint32 `DESCRIPTION: "Reference bandwidth in kilobits/second for calculating default interface metrics. Unit: Mbps", MIN: 100, MAX: 2147483647, DEFAULT: 100`
}

type Ospfv2GlobalState struct {
	ConfigObj
	Vrf                string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"r", MULTIPLICITY:"1", DESCRIPTION: "VRF id for OSPF global config", DEFAULT:"Default"`
	AreaBdrRtrStatus   bool   `DESCRIPTION: A flag to note whether this router is an Area Border Router.`
	NumOfAreas         uint32 `DESCRIPTION: Number of OSPF Areas.`
	NumOfIntfs         uint32 `DESCRIPTION: Number of OSPF interfaces.`
	NumOfNbrs          uint32 `DESCRIPTION: Number of Neighbors.`
	NumOfLSA           uint32 `DESCRIPTION: Number of LSAs.`
	NumOfRouterLSA     uint32 `DESCRIPTION: Number of Router LSAs.`
	NumOfNetworkLSA    uint32 `DESCRIPTION: Number of Network LSAs.`
	NumOfSummary3LSA   uint32 `DESCRIPTION: Number of Summary 3 LSAs.`
	NumOfSummary4LSA   uint32 `DESCRIPTION: Number of Summary 4 LSAs.`
	NumOfASExternalLSA uint32 `DESCRIPTION: Number of ASExternal LSAs.`
	NumOfRoutes        uint32 `DESCRIPTION: Number of Routes (Unsupported).`
}

type Ospfv2Area struct {
	ConfigObj
	AreaId         string `SNAPROUTE: "KEY", CATEGORY:"L3",  ACCESS:"w", MULTIPLICITY:"*", DESCRIPTION: A 32-bit integer uniquely identifying an area. Area ID 0.0.0.0 is used for the OSPF backbone.`
	AdminState     string `DESCRIPTION: Indicates if OSPF is enabled on this area, DEFAULT:"DOWN"`
	AuthType       string `DESCRIPTION: The authentication type specified for an area., SELECTION: none(0)/simplePassword(1)/md5(2), DEFAULT:"None"`
	ImportASExtern bool   `DESCRIPTION: ExternalRoutingCapability if false AS External LSA will not be flooded into this area, DEFAULT: true`
}

type Ospfv2AreaState struct {
	ConfigObj
	AreaId string `SNAPROUTE: "KEY", CATEGORY:"L3",  ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: A 32-bit integer uniquely identifying an area. Area ID 0.0.0.0 is used for the OSPF backbone.`
	//NumSpfRuns       uint32 `DESCRIPTION: The number of times that the intra-area route table has been calculated using this area's link state database.  This is typically done using Dijkstra's algorithm.  Discontinuities in the value of this counter can occur at re-initialization of the management system, and at other times as indicated by the value of ospfDiscontinuityTime.`
	//NumBdrRtr        uint32 `DESCRIPTION: The total number of Area Border Routers reachable within this area.  This is initially zero and is calculated in each Shortest Path First (SPF) pass.`
	//NumAsBdrRtr      uint32 `DESCRIPTION: The total number of Autonomous System Border Routers reachable within this area.  This is initially zero and is calculated in each SPF pass.`
	NumOfRouterLSA     uint32 `DESCRIPTION: Number of Router LSA in a given Area`
	NumOfNetworkLSA    uint32 `DESCRIPTION: Number of Network LSA in a given Area`
	NumOfSummary3LSA   uint32 `DESCRIPTION: Number of Summary3 LSA in a given Area`
	NumOfSummary4LSA   uint32 `DESCRIPTION: Number of Summary4 LSA in a given Area`
	NumOfASExternalLSA uint32 `DESCRIPTION: Number of ASExternal LSA in a given Area`
	NumOfIntfs         uint32 `DESCRIPTION: Number of Interfaces in a given Area.`
	NumOfLSA           uint32 `DESCRIPTION: Number of LSAs in a given Area.`
	NumOfNbrs          uint32 `DESCRIPTION: Number of Neighbors in a given Area`
	NumOfRoutes        uint32 `DESCRIPTION: Number of Routes in a given Area (Unsupported).`
}
type Ospfv2Intf struct {
	ConfigObj
	IpAddress        string `SNAPROUTE: "KEY", CATEGORY:"L3",  ACCESS:"w", MULTIPLICITY:"*", DESCRIPTION: The IP address of this OSPF interface., RELTN:"DEP:[Vlan, Port]`
	AddressLessIfIdx uint32 `SNAPROUTE: "KEY", CATEGORY:"L3",  DESCRIPTION: For the purpose of easing the instancing of addressed and addressless interfaces; this variable takes the value 0 on interfaces with IP addresses and the corresponding value of ifIndex for interfaces having no IP address., MIN: 0, MAX: 2147483647`
	AdminState       string `DESCRIPTION: Indiacates if OSPF is enabled on this interface, DEFAULT:"DOWN"`
	AreaId           string `DESCRIPTION: A 32-bit integer uniquely identifying the area to which the interface connects.  Area ID 0.0.0.0 is used for the OSPF backbone., DEFAULT:"0.0.0.0"`
	Type             string `DESCRIPTION: The OSPF interface type. By way of a default, this field may be intuited from the corresponding value of ifType. Broadcast LANs, such as Ethernet and IEEE 802.5, take the value 'broadcast', X.25 and similar technologies take the value 'nbma', and links that are definitively point to point take the value 'pointToPoint'., SELECTION: Broadcast/PointToPoint, DEFAULT:"Broadcast"`
	RtrPriority      uint8  `DESCRIPTION: The priority of this interface.  Used in multi-access networks, this field is used in the designated router election algorithm.  The value 0 signifies that the router is not eligible to become the designated router on this particular network.  In the event of a tie in this value, routers will use their Router ID as a tie breaker., MIN: 0, MAX: 255, DEFAULT:1`
	TransitDelay     uint16 `DESCRIPTION: The estimated number of seconds it takes to transmit a link state update packet over this interface.  Note that the minimal value SHOULD be 1 second., MIN: 0, MAX: 3600, DEFAULT:1`
	RetransInterval  uint16 `DESCRIPTION: The number of seconds between link state advertisement retransmissions, for adjacencies belonging to this interface.  This value is also used when retransmitting  database description and Link State request packets. Note that minimal value SHOULD be 1 second., MIN: 0, MAX:3600, DEFAULT:5`
	HelloInterval    uint16 `DESCRIPTION: The length of time, in seconds, between the Hello packets that the router sends on the interface.  This value must be the same for all routers attached to a common network., MIN: 1, MAX: 65535, DEFAULT:10`
	RtrDeadInterval  uint32 `DESCRIPTION: The number of seconds that a router's Hello packets have not been seen before its neighbors declare the router down. This should be some multiple of the Hello interval.  This value must be the same for all routers attached to a common network., MIN: 0, MAX: 2147483647, DEFAULT:40`
	MetricValue      uint16 `DESCRIPTION: The metric of using this Type of Service on this interface.  The default value of the TOS 0 metric is 10^8 / ifSpeed., MIN: 0, MAX: 65535, DEFAULT:10`
}

type Ospfv2IntfState struct {
	ConfigObj
	IpAddress                string `SNAPROUTE: "KEY", CATEGORY:"L3",   ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: The IP address of this OSPF interface.`
	AddressLessIfIdx         uint32 `SNAPROUTE: "KEY", CATEGORY:"L3",  DESCRIPTION: For the purpose of easing the instancing of addressed and addressless interfaces; this variable takes the value 0 on interfaces with IP addresses and the corresponding value of ifIndex for interfaces having no IP address., MIN: 0, MAX: 2147483647`
	State                    string `DESCRIPTION: The OSPF Interface State., SELECTION: otherDesignatedRouter(7)/backupDesignatedRouter(6)/loopback(2)/down(1)/designatedRouter(5)/waiting(3)/pointToPoint(4)`
	DesignatedRouter         string `DESCRIPTION: The IP address of the designated router.`
	DesignatedRouterId       string `DESCRIPTION: The Router ID of the designated router.`
	BackupDesignatedRouter   string `DESCRIPTION: The IP address of the backup designated router.`
	BackupDesignatedRouterId string `DESCRIPTION: The Router ID of the backup designated router.`
	NumOfRouterLSA           uint32 `DESCRIPTION: Number of Router LSA in a given Area corresponding to Interface`
	NumOfNetworkLSA          uint32 `DESCRIPTION: Number of Network LSA in a given Area corresponding to Interface`
	NumOfSummary3LSA         uint32 `DESCRIPTION: Number of Summary3 LSA in a given Area corresponding to Interface`
	NumOfSummary4LSA         uint32 `DESCRIPTION: Number of Summary4 LSA in a given Area corresponding to Interfaces`
	NumOfASExternalLSA       uint32 `DESCRIPTION: Number of ASExternal LSA in a given Area`
	NumOfLSA                 uint32 `DESCRIPTION: Number of LSAs in a given Area corresponding to Interface.`
	NumOfNbrs                uint32 `DESCRIPTION: Number of Neighbors in a given Interface`
	NumOfRoutes              uint32 `DESCRIPTION: Number of Routes in a given Interface (Unsupported).`
	Mtu                      uint32 `DESCRIPTION: MTU for a given Interface.`
	Cost                     uint32 `DESCRIPTION: Cost for a given Interface.`
	NumOfStateChange         uint32 `DESCRIPTION: Number of FSM State Change.`
	TimeOfStateChange        string `DESCRIPTION: Last time stamp Intf FSM State Change.`
}

type Ospfv2NbrState struct {
	ConfigObj
	IpAddr           string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"r", MULTIPLICITY:"*",  DESCRIPTION: The IP address this neighbor is using in its IP source address.  Note that, on addressless links, this will not be 0.0.0.0 but the  address of another of the neighbor's interfaces.`
	AddressLessIfIdx uint32 `SNAPROUTE: "KEY", CATEGORY:"L3",  DESCRIPTION: On an interface having an IP address, zero. On addressless interfaces, the corresponding value of ifIndex in the Internet Standard MIB. On row creation, this can be derived from the instance., MIN:0, MAX: 2147483647`
	RtrId            string `DESCRIPTION: A 32-bit integer (represented as a type IpAddress) uniquely identifying the neighboring router in the Autonomous System.`
	Options          int32  `DESCRIPTION: A bit mask corresponding to the neighbor's options field.  Bit 0, if set, indicates that the system will operate on Type of Service metrics other than TOS 0.  If zero, the neighbor will ignore all metrics except the TOS 0 metric.  Bit 1, if set, indicates that the associated area accepts and operates on external information; if zero, it is a stub area.  Bit 2, if set, indicates that the system is capable of routing IP multicast datagrams, that is that it implements the multicast extensions to OSPF.  Bit 3, if set, indicates that the associated area is an NSSA.  These areas are capable of carrying type-7 external advertisements, which are translated into type-5 external advertisements at NSSA borders.`
	State            string `DESCRIPTION: The state of the relationship with this neighbor., SELECTION: exchangeStart(5)/loading(7)/attempt(2)/exchange(6)/down(1)/init(3)/full(8)/twoWay(4)`
}

type Ospfv2LsdbState struct {
	baseObj
	LSType        string `SNAPROUTE: "KEY", CATEGORY:"L3",  ACCESS:"r",  MULTIPLICITY:"*", DESCRIPTION: "The type of the link state advertisement. Each link state type has a separate advertisement format.  Note: External link state advertisements are permitted for backward compatibility, but should be displayed in the AsLsdbTable rather than here., SELECTION: router/network/summary3/summary4/asexternal"`
	LSId          string `SNAPROUTE: "KEY", CATEGORY:"L3",  DESCRIPTION: The Link State ID is an LS Type Specific field containing either a Router ID or an IP address; it identifies the piece of the routing domain that is being described by the advertisement.`
	AreaId        string `SNAPROUTE: "KEY", CATEGORY:"L3",  DESCRIPTION: The 32-bit identifier of the area from which the LSA was received.`
	AdvRouterId   string `SNAPROUTE: "KEY", CATEGORY:"L3",  DESCRIPTION: The 32-bit number that uniquely identifies the originating router in the Autonomous System.`
	SequenceNum   string `DESCRIPTION: The sequence number field is a signed 32-bit integer.  It starts with the value '80000001'h, or -'7FFFFFFF'h, and increments until '7FFFFFFF'h. Thus, a typical sequence number will be very negative. It is used to detect old and duplicate Link State Advertisements.  The space of sequence numbers is linearly ordered.  The larger the sequence number, the more recent the advertisement.`
	Age           uint16 `DESCRIPTION: This field is the age of the link state advertisement in seconds.`
	Checksum      uint16 `DESCRIPTION: This field is the checksum of the complete contents of the advertisement, excepting the age field.  The age field is excepted so that an advertisement's age can be incremented without updating the checksum.  The checksum used is the same that is used for ISO connectionless  datagrams; it is commonly referred to as the Fletcher checksum.`
	Options       uint8  `DESCRIPTION: Options field in LSA.`
	Length        uint16 `DESCRIPTION: Lenght of LSA including LSA header.`
	Advertisement string `DESCRIPTION: The entire link state advertisement, including its header.  Note that for variable length LSAs, SNMP agents may not be able to return the largest string size.`
}

type Ospfv2NextHop struct {
	IntfIPAddr    string `DESCRIPTION: O/P interface IP address`
	IntfIdx       uint32 `DESCRIPTION: Interface index `
	NextHopIPAddr string `DESCRIPTION: Nexthop ip address`
	AdvRtrId      string `DESCRIPTION: Advertising router id`
}

type Ospfv2RouteState struct {
	baseObj
	DestId            string          `SNAPROUTE: "KEY", CATEGORY:"L3",  ACCESS:"r",  MULTIPLICITY:"*", DESCRIPTION: "Dest ip" , USESTATEDB:"true"`
	AddrMask          string          ` SNAPROUTE: "KEY", CATEGORY:"L3", DESCRIPTION: "netmask"`
	DestType          string          `SNAPROUTE: "KEY", CATEGORY:"L3", DESCRIPTION: destination type`
	OptCapabilities   int32           `DESCRIPTION: "capabilities", MIN: 0, MAX:2147483647`
	AreaId            string          `DESCRIPTION: area id for the route`
	PathType          string          `DESCRIPTION: "Path type such as direct / connected / ext"`
	Cost              uint32          `DESCRIPTION: "Cost to reach the destination", MIN: 0, MAX:2147483647`
	Type2Cost         uint32          `DESCRIPTION: "Type2 cost used for external routes.", MIN: 0, MAX:2147483647`
	NumOfPaths        uint16          `DESCRIPTION: "Total number of paths", MIN: 0, MAX: 2147483647`
	LSOriginLSType    string          `DESCRIPTION: "Link State Type only valid for Intra Area"`
	LSOriginLSId      string          `DESCRIPTION: "Link State Id only valid for Intra Area"`
	LSOriginAdvRouter string          `DESCRIPTION: "Advertising router Id only valid for Intra Area"`
	NextHops          []Ospfv2NextHop `DESCRIPTION: "Nexthops for this route"`
}
