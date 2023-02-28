package gocbm

import (
	"encoding/json"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/couchbase/gocbcore/v9"
)

type GetIndexStatsOpts struct {
	MemdAddrs []string
	HTTPAddrs []string
	Username  string
	Password  string

	AgentGroupConfig *gocbcore.AgentGroupConfig

	IndexName string
}

func GetIndexStats(opts GetIndexStatsOpts) (*IndexStatsResponse, error) {
	agentGroup, err := gocbcore.CreateAgentGroup(opts.GetAgentGroupConfig())
	if err != nil {
		return nil, err
	}
	coreReq := &gocbcore.HTTPRequest{
		Service: gocbcore.FtsService,
		Method:  "GET",
		Path:    fmt.Sprintf("/api/stats/index/%s", opts.IndexName),
	}

	var wg = &sync.WaitGroup{}
	var innerErr error
	var innerData []byte

	wg.Add(1)
	_, err = agentGroup.DoHTTPRequest(coreReq, func(response *gocbcore.HTTPResponse, err error) {
		defer wg.Done()
		if err != nil {
			innerErr = err
		}
		data, err := io.ReadAll(response.Body)
		if err != nil {
			innerErr = err
		}
		innerData = data
	})
	if err != nil {
		return nil, err
	}
	wg.Wait()
	if innerErr != nil {
		return nil, innerErr
	}

	var resp IndexStatsResponse
	err = json.Unmarshal(innerData, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (opts GetIndexStatsOpts) GetAgentGroupConfig() *gocbcore.AgentGroupConfig {
	if opts.AgentGroupConfig != nil {
		return opts.AgentGroupConfig
	}

	return &gocbcore.AgentGroupConfig{
		AgentConfig: gocbcore.AgentConfig{
			MemdAddrs: opts.MemdAddrs,
			HTTPAddrs: opts.HTTPAddrs,
			UserAgent: "gocb/v2.2.4", // I don't care
			Auth: gocbcore.PasswordAuthProvider{
				Username: opts.Username,
				Password: opts.Password,
			},
			UseMutationTokens:      true,
			UseDurations:           true,
			UseOutOfOrderResponses: true,
			UseCollections:         true,
			ConnectTimeout:         10000000000,
			KVConnectTimeout:       7 * time.Second,
			NoRootTraceSpans:       true,
			DefaultRetryStrategy:   nil,
			CircuitBreakerConfig:   gocbcore.CircuitBreakerConfig{},
			UseZombieLogger:        true,
		},
	}
}
