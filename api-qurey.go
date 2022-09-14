package minio

import (
	"context"
	"net/http"
	"net/url"
)

func (c *Client) QueryPrice(ctx context.Context) (int, error) {
	urlValues := make(url.Values)
	urlValues.Set("queryprice", "")

	reqMetadata := requestMetadata{
		queryValues:      urlValues,
		contentSHA256Hex: emptySHA256Hex,
	}

	resp, err := c.executeMethod(ctx, http.MethodGet, reqMetadata)

	defer closeResponse(resp)
	if err != nil {
		return -1, err
	}
	if resp != nil {
		if resp.StatusCode != http.StatusOK {
			return -1, httpRespToErrorResponse(resp, "", "")
		}
	}
	queryPriceResult := queryPriceResult{}
	err = xmlDecoder(resp.Body, &queryPriceResult)
	if err != nil {
		return -1, err
	}
	return queryPriceResult.Price, nil
}

func (c *Client) GetBalanceInfo(ctx context.Context) (string, error) {
	urlValues := make(url.Values)
	urlValues.Set("getbalance", "")

	reqMetadata := requestMetadata{
		queryValues:      urlValues,
		contentSHA256Hex: emptySHA256Hex,
	}

	resp, err := c.executeMethod(ctx, http.MethodGet, reqMetadata)

	defer closeResponse(resp)
	if err != nil {
		return "", err
	}
	if resp != nil {
		if resp.StatusCode != http.StatusOK {
			return "", httpRespToErrorResponse(resp, "", "")
		}
	}
	getBalanceResult := getBalanceResult{}
	err = xmlDecoder(resp.Body, &getBalanceResult)
	if err != nil {
		return "", err
	}
	return getBalanceResult.Balance, nil
}
