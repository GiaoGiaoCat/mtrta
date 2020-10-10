package mtrta

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

func Request(cfg *Config, url string, req *RtaRequest) (resp *RtaResponse, err error) {
	if cfg != nil && cfg.Mock {
		reqID := ""
		if req.Id != nil {
			reqID = *req.Id
		}
		resp = &RtaResponse{
			RequestId:         proto.String("mock_" + reqID),
			Code:              proto.Uint32(0),
			PromotionTargetId: []int64{20005, 20006, 20007},
		}
		return
	}
	payload, err := proto.Marshal(req)
	if err != nil {
		return resp, errors.WithStack(err)
	}

	res, err := send(cfg, url, bytes.NewBuffer(payload))
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

func send(cfg *Config, url string, body io.Reader) (*http.Response, error) {
	// Pass context.Background() to SendWithContext
	to := 60 * time.Millisecond
	if cfg != nil && cfg.Timeout != 0 {
		to = cfg.Timeout
	}
	ctx, cancel := context.WithTimeout(context.Background(), to)
	defer cancel()

	return sendWithContext(ctx, cfg, url, body)
}

// Sending an HTTP request and accepting context.
func sendWithContext(ctx context.Context, cfg *Config, url string, body io.Reader) (*http.Response, error) {
	// Change NewRequest to NewRequestWithContext and pass context it
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	req.Header.Add("Content-Type", "application/x-protobuf;charset=UTF-8")

	res, err := GetHTTPClient(cfg).Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return res, nil
}
