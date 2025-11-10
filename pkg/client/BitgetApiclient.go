package client

import (
	"github.com/nick-bisonai/v3-bitget-api-sdk/pkg/common"
	"github.com/nick-bisonai/v3-bitget-api-sdk/pkg/utils"
)

type BitgetApiClient struct {
	BitgetRestClient *common.BitgetRestClient
}

func (p *BitgetApiClient) Init() *BitgetApiClient {
	p.BitgetRestClient = new(common.BitgetRestClient).Init()
	return p
}

func (p *BitgetApiClient) Post(url string, params map[string]string) ([]byte, error) {
	postBody, jsonErr := utils.ToJson(params)
	if jsonErr != nil {
		return nil, jsonErr
	}
	resp, err := p.BitgetRestClient.DoPost(url, postBody)
	return resp, err
}

func (p *BitgetApiClient) Get(url string, params map[string]string) ([]byte, error) {
	resp, err := p.BitgetRestClient.DoGet(url, params)
	return resp, err
}
