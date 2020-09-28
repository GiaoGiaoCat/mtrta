package main

import (
	"fmt"

	"github.com/GiaoGiaoCat/mtrta"
	"google.golang.org/protobuf/proto"
)

func main() {
	device := &mtrta.RtaRequest_Device{
		Os:              mtrta.RtaRequest_OperatingSystem.Enum(mtrta.RtaRequest_OS_ANDROID),
		IdfaMd5Sum:      proto.String(""),
		ImeiMd5Sum:      proto.String("725e5fa0ebf879a99da9eed8636ef7d2"),
		AndroidIdMd5Sum: proto.String(""),
		MacMd5Sum:       proto.String(""),
		OaidMd5Sum:      proto.String(""),
		Ip:              proto.String("106.121.177.97"),
		Oaid:            proto.String("cfdfdcca-ffb7-cd28-f6df-777ff1cfc250"),
		Ipv6:            proto.String(""),
	}

	rtaRequest := &mtrta.RtaRequest{
		Id:     proto.String("187123abcde66661601193447"),
		IsPing: proto.Bool(true),
		IsTest: proto.Bool(true),
		Device: device,
		SiteId: proto.String("20009"),
	}

	rtaurl := "https://gdtrtbdsp.meituan.com/rta?rta_site_param=netunion_rta"

	rtaResponse, err := mtrta.Request(rtaurl, rtaRequest)
	if err != nil {
		panic(err)
	}

	fmt.Printf("RequestId %v, Code %v, ProcessingTimeMs %d \n", rtaResponse.GetRequestId(), rtaResponse.GetCode(), rtaResponse.GetProcessingTimeMs())
}
