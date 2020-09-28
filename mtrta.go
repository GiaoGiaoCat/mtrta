package mtrta

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

func Request(url string, req *RtaRequest) (resp *RtaResponse, err error) {
	payload, err := proto.Marshal(req)
	if err != nil {
		return resp, errors.WithStack(err)
	}

	res, err := send(url, bytes.NewBuffer(payload))
	// Check for response error
	if err != nil {
		return resp, errors.WithStack(err)
	}
	// Close response body
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return resp, errors.WithStack(err)
	}

	rtaResponse := &RtaResponse{}
	err = proto.Unmarshal(body, rtaResponse)
	if err != nil {
		return resp, errors.WithStack(err)
	}

	return rtaResponse, nil
}

func send(url string, body io.Reader) (*http.Response, error) {
	// Pass context.Background() to SendWithContext
	return sendWithContext(context.Background(), url, body)
}

// Sending an HTTP request and accepting context.
func sendWithContext(ctx context.Context, url string, body io.Reader) (*http.Response, error) {
	// Change NewRequest to NewRequestWithContext and pass context it
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	req.Header.Add("Content-Type", "application/x-protobuf;charset=UTF-8")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return res, nil
}
