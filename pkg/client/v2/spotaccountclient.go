package v2

import (
	"github.com/nick-bisonai/v3-bitget-api-sdk/internal"
	"github.com/nick-bisonai/v3-bitget-api-sdk/internal/common"
)

type SpotAccountClient struct {
	BitgetRestClient *common.BitgetRestClient
}

func (p *SpotAccountClient) Init(opts ...common.ClientOption) *SpotAccountClient {
	p.BitgetRestClient = new(common.BitgetRestClient).Init(opts...)
	return p
}

func (p *SpotAccountClient) Info() ([]byte, error) {
	params := internal.NewParams()
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/account/info", params)
	return resp, err
}

func (p *SpotAccountClient) Assets(params map[string]string) ([]byte, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/account/assets", params)
	return resp, err
}

func (p *SpotAccountClient) Bills(params map[string]string) ([]byte, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/account/bills", params)
	return resp, err
}

func (p *SpotAccountClient) TransferRecords(params map[string]string) ([]byte, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/account/transferRecords", params)
	return resp, err
}
