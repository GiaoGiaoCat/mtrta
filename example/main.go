package main

import (
	"fmt"
	"time"

	"github.com/GiaoGiaoCat/mtrta"
	"google.golang.org/protobuf/proto"
)

func main() {
	c := &mtrta.Config{
		Dial:      10 * time.Second,
		KeepAlive: 1 * time.Second,
		MaxConns:  10,
		MaxIdle:   10,
		Version:   0,
	}

	device := &mtrta.RtaRequest_Device{
		Os:              mtrta.RtaRequest_OperatingSystem.Enum(mtrta.RtaRequest_OS_ANDROID),
		IdfaMd5Sum:      proto.String(""),
		ImeiMd5Sum:      proto.String("0a50d917250da0101444e165b0d83bae"),
		AndroidIdMd5Sum: proto.String(""),
		MacMd5Sum:       proto.String(""),
		OaidMd5Sum:      proto.String(""),
		Ip:              proto.String("106.121.177.97"),
		Oaid:            proto.String(""),
		Ipv6:            proto.String(""),
	}

	rtaRequest := &mtrta.RtaRequest{
		Id:     proto.String("99ace15790f118cee840a07967b04d24"),
		IsPing: proto.Bool(false),
		IsTest: proto.Bool(false),
		Device: device,
		SiteId: proto.String("xxx"), // 这里用你申请的渠道号替换
	}

	rtaurl := "https://gdtrtbdsp.meituan.com/rta?rta_site_param=netunion_rta"

	rtaResponse, err := mtrta.Request(c, rtaurl, rtaRequest)
	if err != nil {
		panic(err)
	}

	fmt.Printf("RequestId %v, Code %v, GetPromotionTargetId %v \n", rtaResponse.GetRequestId(), rtaResponse.GetCode(), rtaResponse.GetPromotionTargetId())

	rtaResponse, err = mtrta.Request(c, rtaurl, rtaRequest)
	if err != nil {
		panic(err)
	}

	fmt.Printf("RequestId %v, Code %v, GetPromotionTargetId %v \n", rtaResponse.GetRequestId(), rtaResponse.GetCode(), rtaResponse.GetPromotionTargetId())

	c.Version++
	rtaResponse, err = mtrta.Request(c, rtaurl, rtaRequest)
	if err != nil {
		panic(err)
	}

	fmt.Printf("RequestId %v, Code %v, GetPromotionTargetId %v \n", rtaResponse.GetRequestId(), rtaResponse.GetCode(), rtaResponse.GetPromotionTargetId())
}
