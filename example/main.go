package main

import (
	"fmt"

	"github.com/GiaoGiaoCat/mtrta"
	"google.golang.org/protobuf/proto"
)

func main() {
	device := &mtrta.RtaRequest_Device{
		Os: mtrta.RtaRequest_OperatingSystem.Enum(mtrta.RtaRequest_OS_ANDROID),
		// IdfaMd5Sum:      proto.String(""),
		ImeiMd5Sum: proto.String("0a50d917250da0101444e165b0d83bae"),
		// AndroidIdMd5Sum: proto.String(""),
		// MacMd5Sum:       proto.String(""),
		// OaidMd5Sum:      proto.String(""),
		Ip: proto.String("106.121.177.97"),
		// Oaid:            proto.String(""),
		// Ipv6:            proto.String(""),
	}

	rtaRequest := &mtrta.RtaRequest{
		Id:     proto.String("99ace15790f118cee840a07967b04d24"),
		IsPing: proto.Bool(true),
		IsTest: proto.Bool(true),
		Device: device,
		SiteId: proto.String("xxx"),
	}

	rtaurl := "https://gdtrtbdsp.meituan.com/rta?rta_site_param=netunion_rta"

	rtaResponse, err := mtrta.Request(rtaurl, rtaRequest)
	if err != nil {
		panic(err)
	}

	fmt.Printf("RequestId %v, Code %v, GetPromotionTargetId %d \n", rtaResponse.GetRequestId(), rtaResponse.GetCode(), rtaResponse.GetPromotionTargetId())
}
