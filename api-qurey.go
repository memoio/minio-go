package minio

import (
	"context"
	"net/http"
	"net/url"
)

func (c *Client) QueryPrice(ctx context.Context, bucket, size, time string) (string, error) {
	urlValues := make(url.Values)
	urlValues.Set("queryprice", "")

	urlValues.Set("bucket", bucket)
	urlValues.Set("ssize", size)
	urlValues.Set("stime", time)

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
	queryPriceResult := queryPriceResult{}
	err = xmlDecoder(resp.Body, &queryPriceResult)
	if err != nil {
		return "", err
	}
	return queryPriceResult.Price, nil
}

func (c *Client) GetBalanceInfo(ctx context.Context, addr string) (string, error) {
	urlValues := make(url.Values)
	urlValues.Set("getbalance", "")

	urlValues.Set("addr", addr)

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

func (c *Client) GetTokenAddress(ctx context.Context) (string, error) {
	urlValues := make(url.Values)
	urlValues.Set("gettokenaddress", "")

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
	getTokenAddressResult := getTokenAddressResult{}
	err = xmlDecoder(resp.Body, &getTokenAddressResult)
	if err != nil {
		return "", err
	}
	return getTokenAddressResult.Addr, nil
}

func (c *Client) GetGatewayAddress(ctx context.Context) (string, error) {
	urlValues := make(url.Values)
	urlValues.Set("getgatewayaddress", "")

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
	getGatewayAddressResult := getTokenAddressResult{}
	err = xmlDecoder(resp.Body, &getGatewayAddressResult)
	if err != nil {
		return "", err
	}
	return getGatewayAddressResult.Addr, nil
}

func (c *Client) Approve(ctx context.Context, ts, faddr string) error {
	urlValues := make(url.Values)
	urlValues.Set("approve", "")

	urlValues.Set("ts", ts)
	urlValues.Set("faddr", faddr)

	reqMetadata := requestMetadata{
		queryValues:      urlValues,
		contentSHA256Hex: emptySHA256Hex,
	}

	resp, err := c.executeMethod(ctx, http.MethodGet, reqMetadata)
	defer closeResponse(resp)
	if err != nil {
		return err
	}
	if resp != nil {
		if resp.StatusCode != http.StatusOK {
			return httpRespToErrorResponse(resp, "", "")
		}
	}

	return nil
}
