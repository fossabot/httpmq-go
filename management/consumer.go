package management

import (
	"context"
	"fmt"

	"github.com/alwitt/httpmq-go/api"
	"github.com/alwitt/httpmq-go/common"
)

/*
CreateConsumerForStream define a new consumer on a stream

 @param ctxt context.Context - the caller context
 @param stream string - the stream to create the consumer on
 @param params api.ManagementJetStreamConsumerParam - consumer parameters
 @return request ID (to reference logs)
         an error message if request failed
*/
func (c *mgmtAPIWrapperImpl) CreateConsumerForStream(
	ctxt context.Context, stream string, params api.ManagementJetStreamConsumerParam,
) (string, error) {
	baseRequest := c.client.ManagementApi.V1AdminStreamStreamNameConsumerPost(ctxt, stream)
	request := baseRequest.ConsumerParam(params)

	response, httpResp, err :=
		c.client.ManagementApi.V1AdminStreamStreamNameConsumerPostExecute(request)
	if err != nil {
		return "", err
	}

	errorDetail, ok := response.GetErrorOk()
	errorMsg := ""
	if ok {
		msg, ok := errorDetail.GetMessageOk()
		if ok {
			errorMsg = *msg
		}
	}

	requestID := httpResp.Header.Get(common.RequestIDHeader)

	if !response.Success {
		return requestID, fmt.Errorf(
			"failed to define new consumer %s on stream %s: %s", stream, params.Name, errorMsg,
		)
	}

	return requestID, nil
}

/*
ListAllConsumerOfStream query for list of all known consumers on a stream

 @param ctxt context.Context - the caller context
 @param stream string - the stream to query for
 @return request ID (to reference logs)
         list of known consumer of a stream,
         or an error message if request failed
*/
func (c *mgmtAPIWrapperImpl) ListAllConsumerOfStream(ctxt context.Context, stream string) (
	string, map[string]api.ApisAPIRestRespConsumerInfo, error,
) {
	request := c.client.ManagementApi.V1AdminStreamStreamNameConsumerGet(ctxt, stream)

	response, httpResp, err :=
		c.client.ManagementApi.V1AdminStreamStreamNameConsumerGetExecute(request)
	if err != nil {
		return "", map[string]api.ApisAPIRestRespConsumerInfo{}, err
	}

	errorDetail, ok := response.GetErrorOk()
	errorMsg := ""
	if ok {
		msg, ok := errorDetail.GetMessageOk()
		if ok {
			errorMsg = *msg
		}
	}

	requestID := httpResp.Header.Get(common.RequestIDHeader)

	if !response.Success {
		return requestID, map[string]api.ApisAPIRestRespConsumerInfo{}, fmt.Errorf(
			"failed to list all consumers on stream %s: %s", stream, errorMsg,
		)
	}

	return requestID, response.GetConsumers(), nil
}

/*
GetConsumerOfStream query for a particular consumer on a stream

 @param ctxt context.Context - the caller context
 @param stream string - name of the stream
 @param consumer string - name of the consumer
 @return request ID (to reference logs)
         information on that consumer
         or an error message if request failed
*/
func (c *mgmtAPIWrapperImpl) GetConsumerOfStream(ctxt context.Context, stream, consumer string) (
	string, *api.ApisAPIRestRespConsumerInfo, error,
) {
	request := c.client.ManagementApi.V1AdminStreamStreamNameConsumerConsumerNameGet(
		ctxt, stream, consumer,
	)

	response, httpResp, err :=
		c.client.ManagementApi.V1AdminStreamStreamNameConsumerConsumerNameGetExecute(request)
	if err != nil {
		return "", nil, err
	}

	errorDetail, ok := response.GetErrorOk()
	errorMsg := ""
	if ok {
		msg, ok := errorDetail.GetMessageOk()
		if ok {
			errorMsg = *msg
		}
	}

	requestID := httpResp.Header.Get(common.RequestIDHeader)

	if !response.Success {
		return requestID, nil, fmt.Errorf(
			"failed to query consumer %s on stream %s: %s", consumer, stream, errorMsg,
		)
	}

	return requestID, response.Consumer, nil
}

/*
DeleteConsumerOnStream delete a consumer on a stream

 @param ctxt context.Context - the caller context
 @param stream string - name of the stream
 @param consumer string - name of the consumer
 @return request ID (to reference logs)
         an error message if request failed
*/
func (c *mgmtAPIWrapperImpl) DeleteConsumerOnStream(
	ctxt context.Context, stream, consumer string,
) (string, error) {
	request := c.client.ManagementApi.V1AdminStreamStreamNameConsumerConsumerNameDelete(
		ctxt, stream, consumer,
	)

	response, httpResp, err :=
		c.client.ManagementApi.V1AdminStreamStreamNameConsumerConsumerNameDeleteExecute(request)
	if err != nil {
		return "", err
	}

	errorDetail, ok := response.GetErrorOk()
	errorMsg := ""
	if ok {
		msg, ok := errorDetail.GetMessageOk()
		if ok {
			errorMsg = *msg
		}
	}

	requestID := httpResp.Header.Get(common.RequestIDHeader)

	if !response.Success {
		return requestID, fmt.Errorf(
			"failed to delete consumer %s on stream %s: %s", consumer, stream, errorMsg,
		)
	}

	return requestID, nil
}
