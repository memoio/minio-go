package minio

import (
	"context"
	"net/http"
	"net/url"
)

func (c *Client) QueryPrice(ctx context.Context) (string, error) {
	urlValues := make(url.Values)
	urlValues.Set("queryprice", "")

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

	reqMetadata := requestMetadata{
		bucketName:       addr,
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

func (c *Client) GetDCAndPC(ctx context.Context, bucket string) (uint32, uint32, error) {
	urlValues := make(url.Values)
	urlValues.Set("getbucketdcandpc", "")

	reqMetadata := requestMetadata{
		bucketName:       bucket,
		queryValues:      urlValues,
		contentSHA256Hex: emptySHA256Hex,
	}

	resp, err := c.executeMethod(ctx, http.MethodGet, reqMetadata)

	defer closeResponse(resp)
	if err != nil {
		return 0, 0, err
	}
	if resp != nil {
		if resp.StatusCode != http.StatusOK {
			return 0, 0, httpRespToErrorResponse(resp, bucket, "")
		}
	}
	getDCPCResult := getDCPCResult{}
	err = xmlDecoder(resp.Body, &getDCPCResult)
	if err != nil {
		return 0, 0, err
	}
	return getDCPCResult.DC, getDCPCResult.PC, nil
}
