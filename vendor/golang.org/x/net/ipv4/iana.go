// go generate gen.go
// Code generated by the command above; DO NOT EDIT.

package ipv4

// Internet Control Message Protocol (ICMP) Parameters, Updated: 2018-02-26
const (
	ICMPTypeEchoReply              ICMPType = 0  // Echo Reply
	ICMPTypeDestinationUnreachable ICMPType = 3  // Destination Unreachable
	ICMPTypeRedirect               ICMPType = 5  // Redirect
	ICMPTypeEcho                   ICMPType = 8  // Echo
	ICMPTypeRouterAdvertisement    ICMPType = 9  // Router Advertisement
	ICMPTypeRouterSolicitation     ICMPType = 10 // Router Solicitation
	ICMPTypeTimeExceeded           ICMPType = 11 // Time Exceeded
	ICMPTypeParameterProblem       ICMPType = 12 // Parameter Problem
	ICMPTypeTimestamp              ICMPType = 13 // Timestamp
	ICMPTypeTimestampReply         ICMPType = 14 // Timestamp Reply
	ICMPTypePhoturis               ICMPType = 40 // Photuris
	ICMPTypeExtendedEchoRequest    ICMPType = 42 // Extended Echo Request
	ICMPTypeExtendedEchoReply      ICMPType = 43 // Extended Echo Reply
)

// Internet Control Message Protocol (ICMP) Parameters, Updated: 2018-02-26
var icmpTypes = map[ICMPType]string{
	0:  "echo reply",
	3:  "destination unreachable",
	5:  "redirect",
	8:  "echo",
	9:  "router advertisement",
	10: "router solicitation",
	11: "time exceeded",
	12: "parameter problem",
	13: "timestamp",
	14: "timestamp reply",
	40: "photuris",
	42: "extended echo request",
	43: "extended echo reply",
}
