package model

var CiscoFields = map[int]string{
	1:   "IN_BYTES",
	2:   "IN_PKTS",
	3:   "FLOWS",
	4:   "PROTOCOL",
	5:   "SRC_TOS",
	6:   "TCP_FLAGS",
	7:   "L4_SRC_PORT",
	8:   "IPV4_SRC_ADDR",
	9:   "SRC_MASK",
	10:  "INPUT_SNMP",
	11:  "L4_DST_PORT",
	12:  "IPV4_DST_ADDR",
	13:  "DST_MASK",
	14:  "OUTPUT_SNMP",
	15:  "IPV4_NEXT_HOP",
	16:  "SRC_AS",
	17:  "DST_AS",
	18:  "BGP_IPV4_NEXT_HOP",
	19:  "MUL_DST_PKTS",
	20:  "MUL_DST_BYTES",
	21:  "LAST_SWITCHED",
	22:  "FIRST_SWITCHED",
	23:  "OUT_BYTES",
	24:  "OUT_PKTS",
	25:  "MIN_PKT_LNGTH",
	26:  "MAX_PKT_LNGTH",
	27:  "IPV6_SRC_ADDR",
	28:  "IPV6_DST_ADDR",
	29:  "IPV6_SRC_MASK",
	30:  "IPV6_DST_MASK",
	31:  "IPV6_FLOW_LABEL",
	32:  "ICMP_TYPE",
	33:  "MUL_IGMP_TYPE",
	34:  "SAMPLING_INTERVAL",
	35:  "SAMPLING_ALGORITHM",
	36:  "FLOW_ACTIVE_TIMEOUT",
	37:  "FLOW_INACTIVE_TIMEOUT",
	38:  "ENGINE_TYPE",
	39:  "ENGINE_ID",
	40:  "TOTAL_BYTES_EXP",
	41:  "TOTAL_PKTS_EXP",
	42:  "TOTAL_FLOWS_EXP",
	43:  "*Vendor Proprietary*",
	44:  "IPV4_SRC_PREFIX",
	45:  "IPV4_DST_PREFIX",
	46:  "MPLS_TOP_LABEL_TYPE",
	47:  "MPLS_TOP_LABEL_IP_ADDR",
	48:  "FLOW_SAMPLER_ID",
	49:  "FLOW_SAMPLER_MODE",
	50:  "FLOW_SAMPLER_RANDOM_INTERVAL",
	51:  "*Vendor Proprietary*",
	52:  "MIN_TTL",
	53:  "MAX_TTL",
	54:  "IPV4_IDENT",
	55:  "DST_TOS",
	56:  "IN_SRC_MAC",
	57:  "OUT_DST_MAC",
	58:  "SRC_VLAN",
	59:  "DST_VLAN",
	60:  "IP_PROTOCOL_VERSION",
	61:  "DIRECTION",
	62:  "IPV6_NEXT_HOP",
	63:  "BGP_IPV6_NEXT_HOP",
	64:  "IPV6_OPTION_HEADERS",
	65:  "*Vendor Proprietary*",
	66:  "*Vendor Proprietary*",
	67:  "*Vendor Proprietary*",
	68:  "*Vendor Proprietary*",
	69:  "*Vendor Proprietary*",
	70:  "MPLS_LABEL_1",
	71:  "MPLS_LABEL_2",
	72:  "MPLS_LABEL_3",
	73:  "MPLS_LABEL_4",
	74:  "MPLS_LABEL_5",
	75:  "MPLS_LABEL_6",
	76:  "MPLS_LABEL_7",
	77:  "MPLS_LABEL_8",
	78:  "MPLS_LABEL_9",
	79:  "MPLS_LABEL_10",
	80:  "IN_DST_MAC",
	81:  "OUT_SRC_MAC",
	82:  "IF_NAME",
	83:  "IF_DESC",
	84:  "SAMPLER_NAME",
	85:  "IN_PERMANENT_BYTES",
	86:  "IN_PERMANENT_PKTS",
	87:  "*Vendor Proprietary*",
	88:  "FRAGMENT_OFFSET",
	89:  "FORWARDING_STATUS",
	90:  "MPLS_PAL_RD",
	91:  "MPLS_PREFIX_LEN",
	92:  "SRC_TRAFFIC_INDEX",
	93:  "DST_TRAFFIC_INDEX",
	94:  "APPLICATION_DESCRIPTION",
	95:  "APPLICATION_TAG",
	96:  "APPLICATION_NAME",
	98:  "postipDiffServCodePoint",
	99:  "replication_factor",
	100: "DEPRECATED",
	102: "layer2packetSectionOffset",
	103: "layer2packetSectionSize",
	104: "layer2packetSectionData",
}
