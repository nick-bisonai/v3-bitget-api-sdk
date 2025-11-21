package v2

import (
	"github.com/bisonai/v3-bitget-api-sdk/pkg/common"
	"github.com/bisonai/v3-bitget-api-sdk/pkg/utils"
)

type SpotMarketClient struct {
	BitgetRestClient *common.BitgetRestClient
}

func (p *SpotMarketClient) Init(opts ...common.ClientOption) *SpotMarketClient {
	p.BitgetRestClient = new(common.BitgetRestClient).Init(opts...)
	return p
}

func (p *SpotMarketClient) Coins() ([]byte, error) {
	params := utils.NewParams()
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/public/coins", params)
	return resp, err
}

func (p *SpotMarketClient) Symbols(params map[string]string) ([]byte, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/public/symbols", params)
	return resp, err
}

func (p *SpotMarketClient) Fills(params map[string]string) ([]byte, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/market/fills", params)
	return resp, err
}

func (p *SpotMarketClient) Orderbook(params map[string]string) ([]byte, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/market/orderbook", params)
	return resp, err
}

func (p *SpotMarketClient) Tickers(params map[string]string) ([]byte, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/market/tickers", params)
	return resp, err
}

func (p *SpotMarketClient) Candles(params map[string]string) ([]byte, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/market/candles", params)
	return resp, err
}
