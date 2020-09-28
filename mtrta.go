package mtrta

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"google.golang.org/protobuf/proto"
)

func Request(url string, rtaRequest *RtaRequest) (*RtaResponse, error) {
	payload, err := proto.Marshal(rtaRequest)
	if err != nil {
		log.Fatal("RtaRequest marshaling error: ", err)
		return nil, err
	}

	client := &http.Client{
		Timeout: 60 * time.Second,
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		log.Fatal("post request error: ", err)
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-protobuf;charset=UTF-8")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal("send post request error: ", err)
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	rtaResponse := &RtaResponse{}
	err = proto.Unmarshal(body, rtaResponse)
	if err != nil {
		log.Fatal("RtaResponse unmarshaling error: ", err)
		return nil, err
	}

	return rtaResponse, nil
}
