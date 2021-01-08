package coinbasepro

import "fmt"

type Fees struct {
	MakerFeeRate string `json:"maker_fee_rate"`
	TakerFeeRate string `json:"taker_fee_rate"`
	UsdVolume    string `json:"usd_volume"`
}

func (c *Client) GetFees() (Fees, error) {
	var fees []Fees
	_, err := c.Request("GET", "/fees", nil, &fees)
	if len(fees) == 0 {
		return Fees{}, fmt.Errorf("no fees returned")
	}
	return fees[0], err
}
