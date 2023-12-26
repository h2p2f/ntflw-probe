package model

import (
	"bytes"
	"encoding/binary"
	"strconv"
)

type Packet struct {
	SourceID                  string `json:"source_id,omitempty"`
	InBytes                   string `json:"in_bytes,omitempty"`
	InPkts                    string `json:"in_pkts,omitempty"`
	Flows                     string `json:"flows,omitempty"`
	Protocol                  string `json:"protocol,omitempty"`
	SrcTos                    string `json:"src_tos,omitempty"`
	TCPFlags                  string `json:"tcp_flags,omitempty"`
	L4SrcPort                 string `json:"l4_src_port,omitempty"`
	IPv4SrcAddr               string `json:"ipv4_src_addr,omitempty"`
	SrcMask                   string `json:"src_mask,omitempty"`
	InputSnmp                 string `json:"input_snmp,omitempty"`
	L4DstPort                 string `json:"l4_dst_port,omitempty"`
	IPv4DstAddr               string `json:"ipv4_dst_addr,omitempty"`
	DstMask                   string `json:"dst_mask,omitempty"`
	OutputSnmp                string `json:"output_snmp,omitempty"`
	IPv4NextHop               string `json:"ipv4_next_hop,omitempty"`
	SrcAs                     string `json:"src_as,omitempty"`
	DstAs                     string `json:"dst_as,omitempty"`
	BgpIPv4NextHop            string `json:"bgp_ipv4_next_hop,omitempty"`
	MulDstPkts                string `json:"mul_dst_pkts,omitempty"`
	MulDstBytes               string `json:"mul_dst_bytes,omitempty"`
	LastSwitched              string `json:"last_switched,omitempty"`
	FirstSwitched             string `json:"first_switched,omitempty"`
	OutBytes                  string `json:"out_bytes,omitempty"`
	OutPkts                   string `json:"out_pkts,omitempty"`
	MinPktLngth               string `json:"min_pkt_lngth,omitempty"`
	MaxPktLngth               string `json:"max_pkt_lngth,omitempty"`
	IPv6SrcAddr               string `json:"ipv6_src_addr,omitempty"`
	IPv6DstAddr               string `json:"ipv6_dst_addr,omitempty"`
	IPv6SrcMask               string `json:"ipv6_src_mask,omitempty"`
	IPv6DstMask               string `json:"ipv6_dst_mask,omitempty"`
	IPv6FlowLabel             string `json:"ipv6_flow_label,omitempty"`
	IcmpType                  string `json:"icmp_type,omitempty"`
	MulIgmpType               string `json:"mul_igmp_type,omitempty"`
	SamplingInterval          string `json:"sampling_interval,omitempty"`
	SamplingAlgorithm         string `json:"sampling_algorithm,omitempty"`
	FlowActiveTimeout         string `json:"flow_active_timeout,omitempty"`
	FlowInactiveTimeout       string `json:"flow_inactive_timeout,omitempty"`
	EngineType                string `json:"engine_type,omitempty"`
	EngineID                  string `json:"engine_id,omitempty"`
	TotalBytesExp             string `json:"total_bytes_exp,omitempty"`
	TotalPktsExp              string `json:"total_pkts_exp,omitempty"`
	TotalFlowsExp             string `json:"total_flows_exp,omitempty"`
	VendorProprietary43       string `json:"vendor_proprietary_43,omitempty"`
	IPv4SrcPrefix             string `json:"ipv4_src_prefix,omitempty"`
	IPv4DstPrefix             string `json:"ipv4_dst_prefix,omitempty"`
	MplsTopLabelType          string `json:"mpls_top_label_type,omitempty"`
	MplsTopLabelIPv4Addr      string `json:"mpls_top_label_ipv4_addr,omitempty"`
	FlowSamplerID             string `json:"flow_sampler_id,omitempty"`
	FlowSamplerMode           string `json:"flow_sampler_mode,omitempty"`
	FlowSamplerRandomInterval string `json:"flow_sampler_random_interval,omitempty"`
	VendorProprietary51       string `json:"vendor_proprietary_51,omitempty"`
	MinTTL                    string `json:"min_ttl,omitempty"`
	MaxTTL                    string `json:"max_ttl,omitempty"`
	IPv4Ident                 string `json:"ipv4_ident,omitempty"`
	DstTos                    string `json:"dst_tos,omitempty"`
	InSrcMac                  string `json:"in_src_mac,omitempty"`
	OutDstMac                 string `json:"out_dst_mac,omitempty"`
	SrcVlan                   string `json:"src_vlan,omitempty"`
	DstVlan                   string `json:"dst_vlan,omitempty"`
	IpProtocolVersion         string `json:"ip_protocol_version,omitempty"`
	Direction                 string `json:"direction,omitempty"`
	IPv6NextHop               string `json:"ipv6_next_hop,omitempty"`
	BgpIPv6NextHop            string `json:"bgp_ipv6_next_hop,omitempty"`
	IPv6OptionHeaders         string `json:"ipv6_option_headers,omitempty"`
	VendorProprietary65       string `json:"vendor_proprietary_65,omitempty"`
	VendorProprietary66       string `json:"vendor_proprietary_66,omitempty"`
	VendorProprietary67       string `json:"vendor_proprietary_67,omitempty"`
	VendorProprietary68       string `json:"vendor_proprietary_68,omitempty"`
	VendorProprietary69       string `json:"vendor_proprietary_69,omitempty"`
	MplsLabel1                string `json:"mpls_label_1,omitempty"`
	MplsLabel2                string `json:"mpls_label_2,omitempty"`
	MplsLabel3                string `json:"mpls_label_3,omitempty"`
	MplsLabel4                string `json:"mpls_label_4,omitempty"`
	MplsLabel5                string `json:"mpls_label_5,omitempty"`
	MplsLabel6                string `json:"mpls_label_6,omitempty"`
	MplsLabel7                string `json:"mpls_label_7,omitempty"`
	MplsLabel8                string `json:"mpls_label_8,omitempty"`
	MplsLabel9                string `json:"mpls_label_9,omitempty"`
	MplsLabel10               string `json:"mpls_label_10,omitempty"`
	InDstMac                  string `json:"in_dst_mac,omitempty"`
	OutSrcMac                 string `json:"out_src_mac,omitempty"`
	IfName                    string `json:"if_name,omitempty"`
	IfDesc                    string `json:"if_desc,omitempty"`
	SamplerName               string `json:"sampler_name,omitempty"`
	InPermanentBytes          string `json:"in_permanent_bytes,omitempty"`
	InPermanentPkts           string `json:"in_permanent_pkts,omitempty"`
	VendorProprietary87       string `json:"vendor_proprietary_87,omitempty"`
	FragmentOffset            string `json:"fragment_offset,omitempty"`
	ForwardingStatus          string `json:"forwarding_status,omitempty"`
	MplsPalRD                 string `json:"mpls_pal_rd,omitempty"`
	MplsPrefixLen             string `json:"mpls_prefix_len,omitempty"`
	SrcTrafficIndex           string `json:"src_traffic_index,omitempty"`
	DstTrafficIndex           string `json:"dst_traffic_index,omitempty"`
	ApplicationDescription    string `json:"application_description,omitempty"`
	ApplicationTag            string `json:"application_tag,omitempty"`
	ApplicationName           string `json:"application_name,omitempty"`
	PostipDiffServCodePoint   string `json:"postip_diff_serv_code_point,omitempty"`
	ReplicationFactor         string `json:"replication_factor,omitempty"`
	Deprecated100             string `json:"deprecated_100,omitempty"`
	Layer2packetSectionOffset string `json:"layer2packet_section_offset,omitempty"`
	Layer2packetSectionSize   string `json:"layer2packet_section_size,omitempty"`
	Layer2packetSectionData   string `json:"layer2packet_section_data,omitempty"`
}

func (p *Packet) SetField(key int, value []uint8) {
	switch key {
	case 1:
		p.InBytes = StringDefault(value)
	case 2:
		p.InPkts = StringDefault(value)
	case 3:
		p.Flows = StringDefault(value)
	case 4:
		p.Protocol = StringIPProtocol(value)
	case 5:
		p.SrcTos = StringDefault(value)
	case 6:
		p.TCPFlags = StringDefault(value)
	case 7:
		p.L4SrcPort = StringDefault(value)
	case 8:
		p.IPv4SrcAddr = StringIPv4(value)
	case 9:
		p.SrcMask = StringDefault(value)
	case 10:
		p.InputSnmp = StringDefault(value)
	case 11:
		p.L4DstPort = StringDefault(value)
	case 12:
		p.IPv4DstAddr = StringIPv4(value)
	case 13:
		p.DstMask = StringDefault(value)
	case 14:
		p.OutputSnmp = StringDefault(value)
	case 15:
		p.IPv4NextHop = StringIPv4(value)
	case 16:
		p.SrcAs = StringDefault(value)
	case 17:
		p.DstAs = StringDefault(value)
	case 18:
		p.BgpIPv4NextHop = StringIPv4(value)
	case 19:
		p.MulDstPkts = StringDefault(value)
	case 20:
		p.MulDstBytes = StringDefault(value)
	case 21:
		p.LastSwitched = StringDefault(value)
	case 22:
		p.FirstSwitched = StringDefault(value)
	case 23:
		p.OutBytes = StringDefault(value)
	case 24:
		p.OutPkts = StringDefault(value)
	case 25:
		p.MinPktLngth = StringDefault(value)
	case 26:
		p.MaxPktLngth = StringDefault(value)
	case 27:
		p.IPv6SrcAddr = StringDefault(value)
	case 28:
		p.IPv6DstAddr = StringDefault(value)
	case 29:
		p.IPv6SrcMask = StringDefault(value)
	case 30:
		p.IPv6DstMask = StringDefault(value)
	case 31:
		p.IPv6FlowLabel = StringDefault(value)
	case 32:
		p.IcmpType = StringDefault(value)
	case 33:
		p.MulIgmpType = StringDefault(value)
	case 34:
		p.SamplingInterval = StringDefault(value)
	case 35:
		p.SamplingAlgorithm = StringDefault(value)
	case 36:
		p.FlowActiveTimeout = StringDefault(value)
	case 37:
		p.FlowInactiveTimeout = StringDefault(value)
	case 38:
		p.EngineType = StringDefault(value)
	case 39:
		p.EngineID = StringDefault(value)
	case 40:
		p.TotalBytesExp = StringDefault(value)
	case 41:
		p.TotalPktsExp = StringDefault(value)
	case 42:
		p.TotalFlowsExp = StringDefault(value)
	case 43:
		p.VendorProprietary43 = StringDefault(value)
	case 44:
		p.IPv4SrcPrefix = StringDefault(value)
	case 45:
		p.IPv4DstPrefix = StringDefault(value)
	case 46:
		p.MplsTopLabelType = StringDefault(value)
	case 47:
		p.MplsTopLabelIPv4Addr = StringIPv4(value)
	case 48:
		p.FlowSamplerID = StringDefault(value)
	case 49:
		p.FlowSamplerMode = StringDefault(value)
	case 50:
		p.FlowSamplerRandomInterval = StringDefault(value)
	case 51:
		p.VendorProprietary51 = StringDefault(value)
	case 52:
		p.MinTTL = StringDefault(value)
	case 53:
		p.MaxTTL = StringDefault(value)
	case 54:
		p.IPv4Ident = StringDefault(value)
	case 55:
		p.DstTos = StringDefault(value)
	case 56:
		p.InSrcMac = StringMAC(value)
	case 57:
		p.OutDstMac = StringMAC(value)
	case 58:
		p.SrcVlan = StringDefault(value)
	case 59:
		p.DstVlan = StringDefault(value)
	case 60:
		p.IpProtocolVersion = StringDefault(value)
	case 61:
		p.Direction = StringDefault(value)
	case 62:
		p.IPv6NextHop = StringDefault(value)
	case 63:
		p.BgpIPv6NextHop = StringDefault(value)
	case 64:
		p.IPv6OptionHeaders = StringDefault(value)
	case 65:
		p.VendorProprietary65 = StringDefault(value)
	case 66:
		p.VendorProprietary66 = StringDefault(value)
	case 67:
		p.VendorProprietary67 = StringDefault(value)
	case 68:
		p.VendorProprietary68 = StringDefault(value)
	case 69:
		p.VendorProprietary69 = StringDefault(value)
	case 70:
		p.MplsLabel1 = StringDefault(value)
	case 71:
		p.MplsLabel2 = StringDefault(value)
	case 72:
		p.MplsLabel3 = StringDefault(value)
	case 73:
		p.MplsLabel4 = StringDefault(value)
	case 74:
		p.MplsLabel5 = StringDefault(value)
	case 75:
		p.MplsLabel6 = StringDefault(value)
	case 76:
		p.MplsLabel7 = StringDefault(value)
	case 77:
		p.MplsLabel8 = StringDefault(value)
	case 78:
		p.MplsLabel9 = StringDefault(value)
	case 79:
		p.MplsLabel10 = StringDefault(value)
	case 80:
		p.InDstMac = StringMAC(value)
	case 81:
		p.OutSrcMac = StringMAC(value)
	case 82:
		p.IfName = StringDefault(value)
	case 83:
		p.IfDesc = StringDefault(value)
	case 84:
		p.SamplerName = StringDefault(value)
	case 85:
		p.InPermanentBytes = StringDefault(value)
	case 86:
		p.InPermanentPkts = StringDefault(value)
	case 87:
		p.VendorProprietary87 = StringDefault(value)
	case 88:
		p.FragmentOffset = StringDefault(value)
	case 89:
		p.ForwardingStatus = StringDefault(value)
	case 90:
		p.MplsPalRD = StringDefault(value)
	case 91:
		p.MplsPrefixLen = StringDefault(value)
	case 92:
		p.SrcTrafficIndex = StringDefault(value)
	case 93:
		p.DstTrafficIndex = StringDefault(value)
	case 94:
		p.ApplicationDescription = StringDefault(value)
	case 95:
		p.ApplicationTag = StringDefault(value)
	case 96:
		p.ApplicationName = StringDefault(value)
	case 98:
		p.PostipDiffServCodePoint = StringDefault(value)
	case 99:
		p.ReplicationFactor = StringDefault(value)
	case 100:
		p.Deprecated100 = StringDefault(value)
	case 102:
		p.Layer2packetSectionOffset = StringDefault(value)
	case 103:
		p.Layer2packetSectionSize = StringDefault(value)
	case 104:
		p.Layer2packetSectionData = StringDefault(value)
	}
}

func StringDefault(b []uint8) string {
	switch len(b) {
	case 1:
		return strconv.Itoa(int(b[0]))
	case 2:
		var n uint16
		binary.Read(bytes.NewBuffer(b), binary.BigEndian, &n)
		return strconv.Itoa(int(n))
	case 4:
		var n uint32
		binary.Read(bytes.NewBuffer(b), binary.BigEndian, &n)
		return strconv.Itoa(int(n))
	case 8:
		var n uint64
		binary.Read(bytes.NewBuffer(b), binary.BigEndian, &n)
		return strconv.Itoa(int(n))
	}
	s := ""
	for _, n := range b {
		s += strconv.Itoa(int(n)) + " "
	}
	return s
}

func StringIPv4(bytes []uint8) string {
	return strconv.Itoa(int(bytes[0])) + "." +
		strconv.Itoa(int(bytes[1])) + "." +
		strconv.Itoa(int(bytes[2])) + "." +
		strconv.Itoa(int(bytes[3]))
}

func StringMAC(bytes []uint8) string {
	const hexDigit = "0123456789abcdef"
	buf := make([]byte, 0, len(bytes)*3-1)
	for i, b := range bytes {
		if i > 0 {
			buf = append(buf, ':')
		}
		buf = append(buf, hexDigit[b>>4])
		buf = append(buf, hexDigit[b&0xF])
	}
	return string(buf)
}

func StringIPProtocol(b []uint8) string {
	var protoNumber int
	switch len(b) {
	case 1:
		protoNumber = int(b[0])
	case 2:
		var n uint16
		binary.Read(bytes.NewBuffer(b), binary.BigEndian, &n)
		protoNumber = int(n)
	case 4:
		var n uint32
		binary.Read(bytes.NewBuffer(b), binary.BigEndian, &n)
		protoNumber = int(n)
	}
	protoName, ok := IPProtocolNames[protoNumber]
	if !ok {
		return strconv.Itoa(protoNumber)
	}
	if protoName == "" {
		return strconv.Itoa(protoNumber)
	}
	return protoName
}
