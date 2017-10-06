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

type AsicGlobalState struct {
	baseObj
	ModuleId   uint8   `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY: "1", DESCRIPTION:"Module identifier"`
	VendorId   string  `DESCRIPTION: "Vendor identification value"`
	PartNumber string  `DESCRIPTION: "Part number of underlying switching asic"`
	RevisionId string  `DESCRIPTION: "Revision ID of underlying switching asic"`
	ModuleTemp float64 `DESCRIPTION: "Current module temperature", UNIT: degC`
}

type PMData struct {
	TimeStamp string  `DESCRIPTION: "Timestamp at which data is collected"`
	Value     float64 `DESCRIPTION: "PM Data Value"`
}

type AsicGlobalPM struct {
	baseObj
	ModuleId           uint8   `SNAPROUTE: "KEY", ACCESS:"rw", MULTIPLICITY: "1", AUTODISCOVER:"true", DESCRIPTION:"Module identifier, DEFAULT: 0"`
	Resource           string  `SNAPROUTE: "KEY", DESCRIPTION: "Resource identifier", SELECTION: "Temperature"`
	PMClassAEnable     bool    `DESCRIPTION: "Enable/Disable control for CLASS-A PM", DEFAULT:true`
	PMClassBEnable     bool    `DESCRIPTION: "Enable/Disable control for CLASS-B PM", DEFAULT:true`
	PMClassCEnable     bool    `DESCRIPTION: "Enable/Disable control for CLASS-C PM", DEFAULT:true`
	HighAlarmThreshold float64 `DESCRIPTION: "High alarm threshold value for this PM", DEFAULT: 100000`
	HighWarnThreshold  float64 `DESCRIPTION: "High warning threshold value for this PM", DEFAULT: 100000`
	LowAlarmThreshold  float64 `DESCRIPTION: "Low alarm threshold value for this PM", DEFAULT: -100000`
	LowWarnThreshold   float64 `DESCRIPTION: "Low warning threshold value for this PM", DEFAULT: -100000`
}

type AsicGlobalPMState struct {
	baseObj
	ModuleId     uint8    `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY: "1", DESCRIPTION:"Module identifier"`
	Resource     string   `SNAPROUTE: "KEY", DESCRIPTION: "Resource identifier"`
	ClassAPMData []PMData `DESCRIPTION: "PM Data corresponding to PM Class A"`
	ClassBPMData []PMData `DESCRIPTION: "PM Data corresponding to PM Class B"`
	ClassCPMData []PMData `DESCRIPTION: "PM Data corresponding to PM Class C"`
}

type EthernetPM struct {
	baseObj
	IntfRef            string  `SNAPROUTE: "KEY", ACCESS:"rw", MULTIPLICITY: "*", AUTODISCOVER:"true", DESCRIPTION: "Interface name of port"`
	Resource           string  `SNAPROUTE: "KEY", DESCRIPTION: "Resource identifier", SELECTION:"StatUnderSizePkts/StatOverSizePkts/StatFragments/StatCRCAlignErrors/StatJabber/StatEtherPkts/StatMCPkts/StatBCPkts/Stat64OctOrLess/Stat65OctTo126Oct/Stat128OctTo255Oct/Stat128OctTo255Oct/Stat256OctTo511Oct/Stat512OctTo1023Oct/Statc1024OctTo1518Oct"`
	PMClassAEnable     bool    `DESCRIPTION: "Enable/Disable control for CLASS-A PM", DEFAULT:true`
	PMClassBEnable     bool    `DESCRIPTION: "Enable/Disable control for CLASS-B PM", DEFAULT:true`
	PMClassCEnable     bool    `DESCRIPTION: "Enable/Disable control for CLASS-C PM", DEFAULT:true`
	HighAlarmThreshold float64 `DESCRIPTION: "High alarm threshold value for this PM", DEFAULT: 100000`
	HighWarnThreshold  float64 `DESCRIPTION: "High warning threshold value for this PM", DEFAULT: 100000`
	LowAlarmThreshold  float64 `DESCRIPTION: "Low alarm threshold value for this PM", DEFAULT: -100000`
	LowWarnThreshold   float64 `DESCRIPTION: "Low warning threshold value for this PM", DEFAULT: -100000`
}

type EthernetPMState struct {
	baseObj
	IntfRef      string   `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION:"Interface name of port"`
	Resource     string   `SNAPROUTE: "KEY", DESCRIPTION: "Resource identifier"`
	ClassAPMData []PMData `DESCRIPTION: "PM Data corresponding to PM Class A"`
	ClassBPMData []PMData `DESCRIPTION: "PM Data corresponding to PM Class B"`
	ClassCPMData []PMData `DESCRIPTION: "PM Data corresponding to PM Class C"`
}

type AsicSummaryState struct {
	baseObj
	ModuleId      uint8 `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY: "1", DESCRIPTION:"Module identifier"`
	NumPortsUp    int32 `DESCRIPTION: Summary stating number of ports that have operstate UP`
	NumPortsDown  int32 `DESCRIPTION: Summary stating number of ports that have operstate DOWN`
	NumVlans      int32 `DESCRIPTION: Summary stating number of vlans configured in the asic`
	NumV4Intfs    int32 `DESCRIPTION: Summary stating number of IPv4 interfaces configured in the asic`
	NumV6Intfs    int32 `DESCRIPTION: Summary stating number of IPv6 interfaces configured in the asic`
	NumV4Adjs     int32 `DESCRIPTION: Summary stating number of IPv4 adjacencies configured in the asic`
	NumV6Adjs     int32 `DESCRIPTION: Summary stating number of IPv6 adjacencies configured in the asic`
	NumV4Routes   int32 `DESCRIPTION: Summary stating number of IPv4 routes configured in the asic`
	NumV6Routes   int32 `DESCRIPTION: Summary stating number of IPv6 routes configured in the asic`
	NumECMPRoutes int32 `DESCRIPTION: Summary stating number of ECMP routes configured in the asic`
}

type Vlan struct {
	baseObj
	VlanId        int32    `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY: "*", MIN:"1", MAX: "4094", DESCRIPTION: "802.1Q tag/Vlan ID for vlan being provisioned"`
	IntfList      []string `DESCRIPTION: "List of interface names or ifindex values to  be added as tagged members of the vlan"`
	UntagIntfList []string `DESCRIPTION: "List of interface names or ifindex values to  be added as untagged members of the vlan"`
	AdminState    string   `DESCRIPTION: "Administrative state of this vlan interface", SELECTION:"UP/DOWN", DEFAULT:"UP"`
}

type VlanState struct {
	baseObj
	VlanId                 int32  `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY: "*", DESCRIPTION: "802.1Q tag/Vlan ID for vlan being provisioned"`
	VlanName               string `DESCRIPTION: "System assigned vlan name"`
	OperState              string `DESCRIPTION: "Operational state of vlan interface"`
	IfIndex                int32  `DESCRIPTION: "System assigned interface id for this vlan interface"`
	SysInternalDescription string `DESCRIPTION: "This is a system generated string that explains the operstate value"`
}

type IPv4Intf struct {
	baseObj
	IntfRef    string `SNAPROUTE: "KEY", ACCESS:"w", DESCRIPTION: "Interface name or ifindex of port/lag or vlan on which this IPv4 object is configured", RELTN:"DEP:[Vlan, Port]`
	IpAddr     string `DESCRIPTION: "Interface IP/Net mask in CIDR format to provision on switch interface", STRLEN:"18"`
	AdminState string `DESCRIPTION: "Administrative state of this IP interface", SELECTION:"UP/DOWN", DEFAULT:"UP"`
}

type IPv4IntfState struct {
	baseObj
	IntfRef           string `SNAPROUTE: "KEY", ACCESS:"r", DESCRIPTION: "System assigned interface id of L2 interface (port/lag/vlan) to which this IPv4 object is linked"`
	IfIndex           int32  `DESCRIPTION: "System assigned interface id for this IPv4 interface"`
	IpAddr            string `DESCRIPTION: "Interface IP/Net mask in CIDR format to provision on switch interface"`
	OperState         string `DESCRIPTION: "Operational state of this IP interface"`
	NumUpEvents       int32  `DESCRIPTION: "Number of times the operational state transitioned from DOWN to UP"`
	LastUpEventTime   string `DESCRIPTION: "Timestamp corresponding to the last DOWN to UP operational state change event"`
	NumDownEvents     int32  `DESCRIPTION: "Number of times the operational state transitioned from UP to DOWN"`
	LastDownEventTime string `DESCRIPTION: "Timestamp corresponding to the last UP to DOWN operational state change event"`
	L2IntfType        string `DESCRIPTION: "Type of L2 interface on which IP has been configured (Port/Lag/Vlan)"`
	L2IntfId          int32  `DESCRIPTION: "Id of the L2 interface. Port number/lag id/vlan id."`
}

type Port struct {
	baseObj
	IntfRef        string `SNAPROUTE: "KEY", ACCESS:"rw", MULTIPLICITY:"*", AUTODISCOVER:"true", DESCRIPTION: "Front panel port name or system assigned interface id"`
	IfIndex        int32  `DESCRIPTION: "System assigned interface id for this port. Read only attribute"`
	Description    string `DESCRIPTION: "User provided string description", DEFAULT:"FP Port", STRLEN:"64"`
	PhyIntfType    string `DESCRIPTION: "Type of internal phy interface", STRLEN:"16" SELECTION:"GMII/SGMII/QSMII/SFI/XFI/XAUI/XLAUI/RXAUI/CR/CR2/CR4/KR/KR2/KR4/SR/SR2/SR4/SR10/LR/LR4"`
	AdminState     string `DESCRIPTION: "Administrative state of this port", STRLEN:"4" SELECTION:"UP/DOWN", DEFAULT:"DOWN"`
	MacAddr        string `DESCRIPTION: "Mac address associated with this port", STRLEN:"17"`
	Speed          int32  `DESCRIPTION: "Port speed in Mbps", MIN: 10, MAX: "100000"`
	Duplex         string `DESCRIPTION: "Duplex setting for this port", STRLEN:"16" SELECTION:"Half Duplex/Full Duplex", DEFAULT:"Full Duplex"`
	Autoneg        string `DESCRIPTION: "Autonegotiation setting for this port", STRLEN:"4" SELECTION:"ON/OFF", DEFAULT:"OFF"`
	MediaType      string `DESCRIPTION: "Type of media inserted into this port", STRLEN:"16"`
	Mtu            int32  `DESCRIPTION: "Maximum transmission unit size for this port"`
	BreakOutMode   string `DESCRIPTION: "Break out mode for the port. Only applicable on ports that support breakout. Valid modes - 1x40, 4x10", STRLEN:"6" SELECTION:"1x40(1)/4x10(2)"`
	LoopbackMode   string `DESCRIPTION: "Desired loopback setting for this port", SELECTION:"NONE/MAC/PHY/RMT", DEFAULT:"NONE"`
	EnableFEC      bool   `DESCRIPTION: "Enable/Disable 802.3bj FEC on this interface", DEFAULT: false`
	PRBSTxEnable   bool   `DESCRIPTION: "Enable/Disable generation of PRBS on this port", DEFAULT: false`
	PRBSRxEnable   bool   `DESCRIPTION: "Enable/Disable PRBS checker on this port", DEFAULT: false`
	PRBSPolynomial string `DESCRIPTION: "PRBS polynomial to use for generation/checking", DEFAULT:2^7, SELECTION:"2^7/2^23/2^31"`
}

type PortState struct {
	baseObj
	IntfRef                     string `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Front panel port name or system assigned interface id"`
	IfIndex                     int32  `DESCRIPTION: "System assigned interface id for this port"`
	Name                        string `DESCRIPTION: "System assigned vlan name"`
	OperState                   string `DESCRIPTION: "Operational state of front panel port"`
	NumUpEvents                 int32  `DESCRIPTION: "Number of times the operational state transitioned from DOWN to UP"`
	LastUpEventTime             string `DESCRIPTION: "Timestamp corresponding to the last DOWN to UP operational state change event"`
	NumDownEvents               int32  `DESCRIPTION: "Number of times the operational state transitioned from UP to DOWN"`
	LastDownEventTime           string `DESCRIPTION: "Timestamp corresponding to the last UP to DOWN operational state change event"`
	Pvid                        int32  `DESCRIPTION: "The vlanid assigned to untagged traffic ingressing this port"`
	IfInOctets                  int64  `DESCRIPTION: "RFC2233 Total number of octets received on this port"`
	IfInUcastPkts               int64  `DESCRIPTION: "RFC2233 Total number of unicast packets received on this port"`
	IfInDiscards                int64  `DESCRIPTION: "RFC2233 Total number of inbound packets that were discarded"`
	IfInErrors                  int64  `DESCRIPTION: "RFC2233 Total number of inbound packets that contained an error"`
	IfInUnknownProtos           int64  `DESCRIPTION: "RFC2233 Total number of inbound packets discarded due to unknown protocol"`
	IfOutOctets                 int64  `DESCRIPTION: "RFC2233 Total number of octets transmitted on this port"`
	IfOutUcastPkts              int64  `DESCRIPTION: "RFC2233 Total number of unicast packets transmitted on this port"`
	IfOutDiscards               int64  `DESCRIPTION: "RFC2233 Total number of error free packets discarded and not transmitted"`
	IfOutErrors                 int64  `DESCRIPTION: "RFC2233 Total number of packets discarded and not transmitted due to packet errors"`
	IfEtherUnderSizePktCnt      int64  `DESCRIPTION: "RFC 1757 Total numbe of undersized packets received and transmitted"`
	IfEtherOverSizePktCnt       int64  `DESCRIPTION: "RFC 1757 Total number of oversized packets received and transmitted"`
	IfEtherFragments            int64  `DESCRIPTION: "RFC1757 Total number of ethernet fragments received and transmitted"`
	IfEtherCRCAlignError        int64  `DESCRIPTION: "RFC 1757 Total number of CRC alignment errors"`
	IfEtherJabber               int64  `DESCRIPTION: "RFC 1757 Total number of jabber frames received and transmitted"`
	IfEtherPkts                 int64  `DESCRIPTION: "RFC 1757 Total number of ethernet packets received and transmitted"`
	IfEtherMCPkts               int64  `DESCRIPTION: "RFC 1757 Total number of multicast packets received and transmitted"`
	IfEtherBcastPkts            int64  `DESCRIPTION: "RFC 1757 Total number of ethernet broadcast packets received and transmitted"`
	IfEtherPkts64OrLessOctets   int64  `DESCRIPTION: "RFC1757 Total number of ethernet packets sized 64 bytes or lesser"`
	IfEtherPkts65To127Octets    int64  `DESCRIPTION: "RFC 1757 Total number of ethernet packets sized between 65 and 127 bytes"`
	IfEtherPkts128To255Octets   int64  `DESCRIPTION: "RFC 1757 Total number of ethernet packets sized between 128 and 255 bytes"`
	IfEtherPkts256To511Octets   int64  `DESCRIPTION: "RFC 1757 Total number of ethernet packets sized between 256 and 511 bytes"`
	IfEtherPkts512To1023Octets  int64  `DESCRIPTION: "RFC 1757 Total number of ethernet packets sized between 512 and 1023 bytes"`
	IfEtherPkts1024To1518Octets int64  `DESCRIPTION: "RFC 1757 Total number of ethernet packets sized between 1024 and 1518 bytes"`
	ErrDisableReason            string `DESCRIPTION: "Reason explaining why port has been disabled by protocol code"`
	PresentInHW                 string `DESCRIPTION: "Indication of whether this port object maps to a physical port. Set to 'No' for ports that are not broken out."`
	ConfigMode                  string `DESCRIPTION: "The current mode of configuration on this port (L2/L3/Internal)"`
	PRBSRxErrCnt                int64  `DESCRIPTION: "Receive error count reported by PRBS checker"`
}

type MacTableEntryState struct {
	baseObj
	MacAddr string `SNAPROUTE: "KEY", ACCESS:"r", DESCRIPTION: "MAC Address", USESTATEDB:"true"`
	VlanId  int32  `DESCRIPTION: "Vlan id corresponding to which mac was learned", DEFAULT:0`
	Port    int32  `DESCRIPTION: "Port number on which mac was learned", DEFAULT:0`
}

type IPv4RouteHwState struct {
	baseObj
	DestinationNw    string `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "IP address of the route in CIDR format"`
	NextHopIps       string `DESCRIPTION: "next hop ip list for the route"`
	RouteCreatedTime string `DESCRIPTION :"Time when the route was added"`
	RouteUpdatedTime string `DESCRIPTION :"Time when the route was last updated"`
}

type IPv6RouteHwState struct {
	baseObj
	DestinationNw    string `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "IP address of the route in CIDR format"`
	NextHopIps       string `DESCRIPTION: "next hop ip list for the route"`
	RouteCreatedTime string `DESCRIPTION :"Time when the route was added"`
	RouteUpdatedTime string `DESCRIPTION :"Time when the route was last updated"`
}

type ArpEntryHwState struct {
	baseObj
	IpAddr  string `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", QPARAM: "optional" ,DESCRIPTION: "Neighbor's IP Address"`
	MacAddr string `DESCRIPTION: "MAC address of the neighbor machine with corresponding IP Address", QPARAM: "optional" `
	Vlan    string `DESCRIPTION: "Vlan ID of the Router Interface to which neighbor is attached to", QPARAM: "optional" `
	Port    string `DESCRIPTION: "Router Interface to which neighbor is attached to", QPARAM: "optional" `
}

type NdpEntryHwState struct {
	baseObj
	IpAddr  string `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", QPARAM: "optional" ,DESCRIPTION: "Neighbor's IP Address"`
	MacAddr string `DESCRIPTION: "MAC address of the neighbor machine with corresponding IP Address", QPARAM: "optional" `
	Vlan    string `DESCRIPTION: "Vlan ID of the Router Interface to which neighbor is attached to", QPARAM: "optional" `
	Port    string `DESCRIPTION: "Router Interface to which neighbor is attached to", QPARAM: "optional" `
}

type LogicalIntf struct {
	baseObj
	Name string `SNAPROUTE: "KEY", ACCESS:"w", DESCRIPTION: "Name of logical interface"`
	Type string `DESCRIPTION: "Type of logical interface (e.x. loopback)", SELECTION:"Loopback", DEFAULT:"Loopback", STRLEN:"16"`
}

type LogicalIntfState struct {
	baseObj
	Name              string `SNAPROUTE: "KEY", ACCESS:"r", DESCRIPTION: "Name of logical interface"`
	IfIndex           int32  `DESCRIPTION: "System assigned interface id for this logical interface"`
	SrcMac            string `DESCRIPTION: "Source Mac assigned to the interface"`
	OperState         string `DESCRIPTION: "Operational state of logical interface"`
	IfInOctets        int64  `DESCRIPTION: "RFC2233 Total number of octets received on this port"`
	IfInUcastPkts     int64  `DESCRIPTION: "RFC2233 Total number of unicast packets received on this port"`
	IfInDiscards      int64  `DESCRIPTION: "RFC2233 Total number of inbound packets that were discarded"`
	IfInErrors        int64  `DESCRIPTION: "RFC2233 Total number of inbound packets that contained an error"`
	IfInUnknownProtos int64  `DESCRIPTION: "RFC2233 Total number of inbound packets discarded due to unknown protocol"`
	IfOutOctets       int64  `DESCRIPTION: "RFC2233 Total number of octets transmitted on this port"`
	IfOutUcastPkts    int64  `DESCRIPTION: "RFC2233 Total number of unicast packets transmitted on this port"`
	IfOutDiscards     int64  `DESCRIPTION: "RFC2233 Total number of error free packets discarded and not transmitted"`
	IfOutErrors       int64  `DESCRIPTION: "RFC2233 Total number of packets discarded and not transmitted due to packet errors"`
}

type SubIPv4Intf struct {
	baseObj
	IpAddr  string `SNAPROUTE: "KEY", ACCESS:"w", DESCRIPTION:"Ip Address for the interface"`
	IntfRef string `SNAPROUTE: "KEY", ACCESS:"w", DESCRIPTION:"Intf name of system generated id (ifindex) of the ipv4Intf where sub interface is to be configured"`
	Type    string `DESCRIPTION:"Type of interface, e.g. Secondary or Virtual", STRLEN:"16"`
	MacAddr string `DESCRIPTION:"Mac address to be used for the sub interface. If none specified IPv4Intf mac address will be used", STRLEN:"17"`
	Enable  bool   `DESCRIPTION:"Enable or disable this interface", DEFAULT:false`
}

type IPv6Intf struct {
	baseObj
	IntfRef    string `SNAPROUTE: "KEY", ACCESS:"w", DESCRIPTION: "Interface name or ifindex of port/lag or vlan on which this IPv4 object is configured", RELTN:"DEP:[Vlan, Port]`
	IpAddr     string `DESCRIPTION: "Interface Global Scope IP Address/Prefix-Length to provision on switch interface", STRLEN:"43", DEFAULT:""`
	LinkIp     bool   `DESCRIPTION: "Interface Link Scope IP Address auto-configured", DEFAULT:true`
	AdminState string `DESCRIPTION: "Administrative state of this IP interface", SELECTION:"UP/DOWN", DEFAULT:"UP"`
}

type IPv6IntfState struct {
	baseObj
	IntfRef           string `SNAPROUTE: "KEY", ACCESS:"r", DESCRIPTION: "System assigned interface id of L2 interface (port/lag/vlan) to which this IPv4 object is linked"`
	IfIndex           int32  `DESCRIPTION: "System assigned interface id for this IPv4 interface"`
	IpAddr            string `DESCRIPTION: "Interface IP Address/Prefix-Lenght to provisioned on switch interface", STRLEN:"43"`
	OperState         string `DESCRIPTION: "Operational state of this IP interface"`
	NumUpEvents       int32  `DESCRIPTION: "Number of times the operational state transitioned from DOWN to UP"`
	LastUpEventTime   string `DESCRIPTION: "Timestamp corresponding to the last DOWN to UP operational state change event"`
	NumDownEvents     int32  `DESCRIPTION: "Number of times the operational state transitioned from UP to DOWN"`
	LastDownEventTime string `DESCRIPTION: "Timestamp corresponding to the last UP to DOWN operational state change event"`
	L2IntfType        string `DESCRIPTION: "Type of L2 interface on which IP has been configured (Port/Lag/Vlan)"`
	L2IntfId          int32  `DESCRIPTION: "Id of the L2 interface. Port number/lag id/vlan id."`
}

type SubIPv6Intf struct {
	baseObj
	IpAddr  string `SNAPROUTE: "KEY", ACCESS:"w", DESCRIPTION:"Ip Address for the interface", STRLEN:"43"`
	IntfRef string `SNAPROUTE: "KEY", ACCESS:"w", DESCRIPTION:"Intf name of system generated id (ifindex) of the ipv4Intf where sub interface is to be configured"`
	Type    string `DESCRIPTION:"Type of interface, e.g. Secondary or Virtual", STRLEN:"16"`
	MacAddr string `DESCRIPTION:"Mac address to be used for the sub interface. If none specified IPv4Intf mac address will be used", STRLEN:"17"`
	Enable  bool   `DESCRIPTION:"Enable or disable this interface", DEFAULT:false`
}

type BufferPortStatState struct {
	baseObj
	IntfRef        string `SNAPROUTE: "KEY", ACCESS:"r", DESCRIPTION: "Front panel port name interface id"`
	IfIndex        int32  `DESCRIPTION: "System assigned interface id for this port. Read only attribute"`
	EgressPort     uint64 `DESCRIPTION: "Egress port buffer stats "`
	IngressPort    uint64 `DESCRIPTION: "Ingress port buffer stats "`
	PortBufferStat uint64 `DESCRIPTION: "Per port buffer stats"`
}

type BufferGlobalStatState struct {
	baseObj
	DeviceId          uint32 `SNAPROUTE: "KEY", ACCESS:"r", DESCRIPTION: "Device id"`
	BufferStat        uint64 `DESCRIPTION: "Buffer stats for the device "`
	EgressBufferStat  uint64 `DESCRIPTION: "Egress Buffer stats "`
	IngressBufferStat uint64 `DESCRIPTION: "Ingress buffer stats "`
}

type Acl struct {
	baseObj
	AclName      string   `SNAPROUTE: "KEY", ACCESS:"w",MULTIPLICITY: "*", DESCRIPTION: "Acl name to be used to refer to this ACL"`
	AclType      string   `DESCRIPTION: "Type can be IP/MAC/SVI"`
	IntfList     []string `DESCRIPTION: "list of IntfRef can be port/lag object"`
	RuleNameList []string `DESCRIPTION: "List of rules to be applied to this ACL. This should match with AclRule RuleName"`
	Direction    string   `SNAPROUTE: "IN/OUT direction in which ACL to be applied"`
}

type AclRule struct {
	baseObj
	RuleName    string `SNAPROUTE: "KEY", MULTIPLICITY: "*", ACCESS:"w", DESCRIPTION: "Acl rule name"`
	SourceMac   string `DESCRIPTION: "Source MAC address."`
	DestMac     string `DESCRIPTION: "Destination MAC address"`
	SourceIp    string `DESCRIPTION: "Source IP address"`
	DestIp      string `DESCRIPTION: "Destination IP address"`
	SourceMask  string `DESCRIPTION: "Network mask for source IP"`
	DestMask    string `DESCRIPTION: "Network mark for dest IP"`
	Action      string `DESCRIPTION: "Type of action (Allow/Deny)", DEFAULT:"Allow", STRLEN:"16"`
	Proto       string `DESCRIPTION: "Protocol type"`
	SrcPort     string `DESCRIPTION: "Source Port", DEFAULT:0`
	DstPort     string `DESCRIPTION: "Dest Port", DEFAULT:0`
	L4SrcPort   int32  `DESCRIPTION: "TCP/UDP source port", DEFAULT:0`
	L4DstPort   int32  `DESCRIPTION: "TCP/UDP destionation port", DEFAULT:0`
	L4PortMatch string `DESCRIPTION: "match condition can be EQ(equal) , NEQ(not equal), LT(larger than), GT(greater than), RANGE(port range)", DEFAULT:"NA"`
	L4MinPort   int32  `DESCRIPTION: "Min port when l4 port is specified as range", DEFAULT:0`
	L4MaxPort   int32  `DESCRIPTION: "Max port when l4 port is specified as range", DEFAULT:0`
}

type AclState struct {
	baseObj
	AclName      string   `SNAPROUTE: "KEY", ACCESS:"r",MULTIPLICITY: "*", DESCRIPTION: "Acl name to be used to refer to this ACL", USESTATEDB:"true"`
	RuleNameList []string `DESCRIPTION: "List of acl rules  to be applied to this ACL. This should match with Acl rule key"`
	IntfList     []string `DESCRIPTION: "list of IntfRef can be port/lag object"`
	Direction    string   `SNAPROUTE: "IN/OUT direction in which ACL to be applied"`
}

type AclRuleState struct {
	baseObj
	RuleName   string   `SNAPROUTE: "KEY", MULTIPLICITY: "*", ACCESS:"r", DESCRIPTION: "Acl rule name"`
	AclType    string   `DESCRIPTION: "Type can be IP/MAC/SVI"`
	IntfList   []string `DESCRIPTION: "list of IntfRef can be port/lag object"`
	HwPresence string   `DESCRIPTION: "Check if the rule is installed in hardware. Applied/Not Applied/Failed"`
	HitCount   uint64   `DESCRIPTION: "No of  packets hit the rule if applied."`
}

// NEED TO ADD SUPPORT TO MAKE THIS INTERNAL ONLY
type LinkScopeIpState struct {
	baseObj
	LinkScopeIp string `SNAPROUTE: "KEY", MULTIPLICITY: "*", ACCESS:"r", DESCRIPTION:"Link scope IP Address", USESTATEDB:"true"`
	IntfRef     string `DESCRIPTION: "Interface where the link scope ip is configured"`
	IfIndex     int32  `DESCRIPTION: "System Generated Unique Interface Id"`
	Used        bool   `DESCRIPTION : "states whether the ip being used"`
}

type CoppStatState struct {
	baseObj
	Protocol     string `SNAPROUTE: "KEY", MULTIPLICITY: "*", ACCESS:"r", DESCRIPTION:"Protocol type for which CoPP is configured."`
	PeakRate     int32  `DESCRIPTION:"Peak rate (packets) for policer."`
	BurstRate    int32  `DESCRIPTION:"Burst rate (packets) for policer."`
	GreenPackets int64  `DESCRIPTION:"Packets marked with green for tri color policer."`
	RedPackets   int64  `DESCRIPTION:"Dropped packets. Packets marked with red for tri color policer. "`
}
