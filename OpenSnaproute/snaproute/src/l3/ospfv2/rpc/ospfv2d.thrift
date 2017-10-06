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
//   This is a auto-generated file, please do not edit!
// _______   __       __________   ___      _______.____    __    ____  __  .___________.  ______  __    __  
// |   ____||  |     |   ____\  \ /  /     /       |\   \  /  \  /   / |  | |           | /      ||  |  |  | 
// |  |__   |  |     |  |__   \  V  /     |   (----  \   \/    \/   /  |  |  ---|  |---- |  ,---- |  |__|  | 
// |   __|  |  |     |   __|   >   <       \   \      \            /   |  |     |  |     |  |     |   __   | 
// |  |     |  `----.|  |____ /  .  \  .----)   |      \    /\    /    |  |     |  |     |  `----.|  |  |  | 
// |__|     |_______||_______/__/ \__\ |_______/        \__/  \__/     |__|     |__|      \______||__|  |__| 
//                                                                                                           
		
namespace go ospfv2d
typedef i32 int
typedef i16 uint16
struct Ospfv2Area {
	1 : string AreaId
	2 : string AdminState
	3 : string AuthType
	4 : bool ImportASExtern
}
struct Ospfv2RouteState {
	1 : string DestId
	2 : string AddrMask
	3 : string DestType
	4 : i32 OptCapabilities
	5 : string AreaId
	6 : string PathType
	7 : i32 Cost
	8 : i32 Type2Cost
	9 : i16 NumOfPaths
	10 : string LSOriginLSType
	11 : string LSOriginLSId
	12 : string LSOriginAdvRouter
	13 : list<Ospfv2NextHop> NextHops
}
struct Ospfv2RouteStateGetInfo {
	1: int StartIdx
	2: int EndIdx
	3: int Count
	4: bool More
	5: list<Ospfv2RouteState> Ospfv2RouteStateList
}
struct Ospfv2Intf {
	1 : string IpAddress
	2 : i32 AddressLessIfIdx
	3 : string AdminState
	4 : string AreaId
	5 : string Type
	6 : byte RtrPriority
	7 : i16 TransitDelay
	8 : i16 RetransInterval
	9 : i16 HelloInterval
	10 : i32 RtrDeadInterval
	11 : i16 MetricValue
}
struct Ospfv2NbrState {
	1 : string IpAddr
	2 : i32 AddressLessIfIdx
	3 : string RtrId
	4 : i32 Options
	5 : string State
}
struct Ospfv2NbrStateGetInfo {
	1: int StartIdx
	2: int EndIdx
	3: int Count
	4: bool More
	5: list<Ospfv2NbrState> Ospfv2NbrStateList
}
struct Ospfv2AreaState {
	1 : string AreaId
	2 : i32 NumOfRouterLSA
	3 : i32 NumOfNetworkLSA
	4 : i32 NumOfSummary3LSA
	5 : i32 NumOfSummary4LSA
	6 : i32 NumOfASExternalLSA
	7 : i32 NumOfIntfs
	8 : i32 NumOfLSA
	9 : i32 NumOfNbrs
	10 : i32 NumOfRoutes
}
struct Ospfv2AreaStateGetInfo {
	1: int StartIdx
	2: int EndIdx
	3: int Count
	4: bool More
	5: list<Ospfv2AreaState> Ospfv2AreaStateList
}
struct Ospfv2LsdbState {
	1 : string LSType
	2 : string LSId
	3 : string AreaId
	4 : string AdvRouterId
	5 : string SequenceNum
	6 : i16 Age
	7 : i16 Checksum
	8 : byte Options
	9 : i16 Length
	10 : string Advertisement
}
struct Ospfv2LsdbStateGetInfo {
	1: int StartIdx
	2: int EndIdx
	3: int Count
	4: bool More
	5: list<Ospfv2LsdbState> Ospfv2LsdbStateList
}
struct Ospfv2GlobalState {
	1 : string Vrf
	2 : bool AreaBdrRtrStatus
	3 : i32 NumOfAreas
	4 : i32 NumOfIntfs
	5 : i32 NumOfNbrs
	6 : i32 NumOfLSA
	7 : i32 NumOfRouterLSA
	8 : i32 NumOfNetworkLSA
	9 : i32 NumOfSummary3LSA
	10 : i32 NumOfSummary4LSA
	11 : i32 NumOfASExternalLSA
	12 : i32 NumOfRoutes
}
struct Ospfv2GlobalStateGetInfo {
	1: int StartIdx
	2: int EndIdx
	3: int Count
	4: bool More
	5: list<Ospfv2GlobalState> Ospfv2GlobalStateList
}
struct Ospfv2IntfState {
	1 : string IpAddress
	2 : i32 AddressLessIfIdx
	3 : string State
	4 : string DesignatedRouter
	5 : string DesignatedRouterId
	6 : string BackupDesignatedRouter
	7 : string BackupDesignatedRouterId
	8 : i32 NumOfRouterLSA
	9 : i32 NumOfNetworkLSA
	10 : i32 NumOfSummary3LSA
	11 : i32 NumOfSummary4LSA
	12 : i32 NumOfASExternalLSA
	13 : i32 NumOfLSA
	14 : i32 NumOfNbrs
	15 : i32 NumOfRoutes
	16 : i32 Mtu
	17 : i32 Cost
	18 : i32 NumOfStateChange
	19 : string TimeOfStateChange
}
struct Ospfv2IntfStateGetInfo {
	1: int StartIdx
	2: int EndIdx
	3: int Count
	4: bool More
	5: list<Ospfv2IntfState> Ospfv2IntfStateList
}
struct Ospfv2Global {
	1 : string Vrf
	2 : string RouterId
	3 : string AdminState
	4 : bool ASBdrRtrStatus
	5 : i32 ReferenceBandwidth
}
struct Ospfv2NextHop {
	1 : string IntfIPAddr
	2 : i32 IntfIdx
	3 : string NextHopIPAddr
	4 : string AdvRtrId
}

struct PatchOpInfo {
    1 : string Op
    2 : string Path
    3 : string Value
}
			        
service OSPFV2DServices {
	bool CreateOspfv2Area(1: Ospfv2Area config);
	bool UpdateOspfv2Area(1: Ospfv2Area origconfig, 2: Ospfv2Area newconfig, 3: list<bool> attrset, 4: list<PatchOpInfo> op);
	bool DeleteOspfv2Area(1: Ospfv2Area config);

	Ospfv2RouteStateGetInfo GetBulkOspfv2RouteState(1: int fromIndex, 2: int count);
	Ospfv2RouteState GetOspfv2RouteState(1: string DestId, 2: string AddrMask, 3: string DestType);
	bool CreateOspfv2Intf(1: Ospfv2Intf config);
	bool UpdateOspfv2Intf(1: Ospfv2Intf origconfig, 2: Ospfv2Intf newconfig, 3: list<bool> attrset, 4: list<PatchOpInfo> op);
	bool DeleteOspfv2Intf(1: Ospfv2Intf config);

	Ospfv2NbrStateGetInfo GetBulkOspfv2NbrState(1: int fromIndex, 2: int count);
	Ospfv2NbrState GetOspfv2NbrState(1: string IpAddr, 2: i32 AddressLessIfIdx);
	Ospfv2AreaStateGetInfo GetBulkOspfv2AreaState(1: int fromIndex, 2: int count);
	Ospfv2AreaState GetOspfv2AreaState(1: string AreaId);
	Ospfv2LsdbStateGetInfo GetBulkOspfv2LsdbState(1: int fromIndex, 2: int count);
	Ospfv2LsdbState GetOspfv2LsdbState(1: string LSType, 2: string LSId, 3: string AreaId, 4: string AdvRouterId);
	Ospfv2GlobalStateGetInfo GetBulkOspfv2GlobalState(1: int fromIndex, 2: int count);
	Ospfv2GlobalState GetOspfv2GlobalState(1: string Vrf);
	Ospfv2IntfStateGetInfo GetBulkOspfv2IntfState(1: int fromIndex, 2: int count);
	Ospfv2IntfState GetOspfv2IntfState(1: string IpAddress, 2: i32 AddressLessIfIdx);
	bool CreateOspfv2Global(1: Ospfv2Global config);
	bool UpdateOspfv2Global(1: Ospfv2Global origconfig, 2: Ospfv2Global newconfig, 3: list<bool> attrset, 4: list<PatchOpInfo> op);
	bool DeleteOspfv2Global(1: Ospfv2Global config);

}