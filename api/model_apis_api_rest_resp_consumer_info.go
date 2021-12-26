/*
httpmq

HTTP/2 based message broker built around NATS JetStream

API version: v0.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ApisAPIRestRespConsumerInfo struct for ApisAPIRestRespConsumerInfo
type ApisAPIRestRespConsumerInfo struct {
	AckFloor ApisAPIRestRespSequenceInfo   `json:"ack_floor"`
	Config   ApisAPIRestRespConsumerConfig `json:"config"`
	// Created is when this consumer was defined
	Created   string                      `json:"created"`
	Delivered ApisAPIRestRespSequenceInfo `json:"delivered"`
	// Name is the name of the consumer
	Name string `json:"name"`
	// NumAckPending is the number of ACK pending / messages in-flight
	NumAckPending int32 `json:"num_ack_pending"`
	// NumPending is the number of message to be delivered for this consumer
	NumPending int32 `json:"num_pending"`
	// NumRedelivered is the number of messages redelivered
	NumRedelivered int32 `json:"num_redelivered"`
	// NumWaiting NATS JetStream does not clearly document this
	NumWaiting int32 `json:"num_waiting"`
	// Stream is the name of the stream
	StreamName string `json:"stream_name"`
}

// NewApisAPIRestRespConsumerInfo instantiates a new ApisAPIRestRespConsumerInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApisAPIRestRespConsumerInfo(ackFloor ApisAPIRestRespSequenceInfo, config ApisAPIRestRespConsumerConfig, created string, delivered ApisAPIRestRespSequenceInfo, name string, numAckPending int32, numPending int32, numRedelivered int32, numWaiting int32, streamName string) *ApisAPIRestRespConsumerInfo {
	this := ApisAPIRestRespConsumerInfo{}
	this.AckFloor = ackFloor
	this.Config = config
	this.Created = created
	this.Delivered = delivered
	this.Name = name
	this.NumAckPending = numAckPending
	this.NumPending = numPending
	this.NumRedelivered = numRedelivered
	this.NumWaiting = numWaiting
	this.StreamName = streamName
	return &this
}

// NewApisAPIRestRespConsumerInfoWithDefaults instantiates a new ApisAPIRestRespConsumerInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApisAPIRestRespConsumerInfoWithDefaults() *ApisAPIRestRespConsumerInfo {
	this := ApisAPIRestRespConsumerInfo{}
	return &this
}

// GetAckFloor returns the AckFloor field value
func (o *ApisAPIRestRespConsumerInfo) GetAckFloor() ApisAPIRestRespSequenceInfo {
	if o == nil {
		var ret ApisAPIRestRespSequenceInfo
		return ret
	}

	return o.AckFloor
}

// GetAckFloorOk returns a tuple with the AckFloor field value
// and a boolean to check if the value has been set.
func (o *ApisAPIRestRespConsumerInfo) GetAckFloorOk() (*ApisAPIRestRespSequenceInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AckFloor, true
}

// SetAckFloor sets field value
func (o *ApisAPIRestRespConsumerInfo) SetAckFloor(v ApisAPIRestRespSequenceInfo) {
	o.AckFloor = v
}

// GetConfig returns the Config field value
func (o *ApisAPIRestRespConsumerInfo) GetConfig() ApisAPIRestRespConsumerConfig {
	if o == nil {
		var ret ApisAPIRestRespConsumerConfig
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *ApisAPIRestRespConsumerInfo) GetConfigOk() (*ApisAPIRestRespConsumerConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *ApisAPIRestRespConsumerInfo) SetConfig(v ApisAPIRestRespConsumerConfig) {
	o.Config = v
}

// GetCreated returns the Created field value
func (o *ApisAPIRestRespConsumerInfo) GetCreated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *ApisAPIRestRespConsumerInfo) GetCreatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *ApisAPIRestRespConsumerInfo) SetCreated(v string) {
	o.Created = v
}

// GetDelivered returns the Delivered field value
func (o *ApisAPIRestRespConsumerInfo) GetDelivered() ApisAPIRestRespSequenceInfo {
	if o == nil {
		var ret ApisAPIRestRespSequenceInfo
		return ret
	}

	return o.Delivered
}

// GetDeliveredOk returns a tuple with the Delivered field value
// and a boolean to check if the value has been set.
func (o *ApisAPIRestRespConsumerInfo) GetDeliveredOk() (*ApisAPIRestRespSequenceInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Delivered, true
}

// SetDelivered sets field value
func (o *ApisAPIRestRespConsumerInfo) SetDelivered(v ApisAPIRestRespSequenceInfo) {
	o.Delivered = v
}

// GetName returns the Name field value
func (o *ApisAPIRestRespConsumerInfo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApisAPIRestRespConsumerInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApisAPIRestRespConsumerInfo) SetName(v string) {
	o.Name = v
}

// GetNumAckPending returns the NumAckPending field value
func (o *ApisAPIRestRespConsumerInfo) GetNumAckPending() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumAckPending
}

// GetNumAckPendingOk returns a tuple with the NumAckPending field value
// and a boolean to check if the value has been set.
func (o *ApisAPIRestRespConsumerInfo) GetNumAckPendingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumAckPending, true
}

// SetNumAckPending sets field value
func (o *ApisAPIRestRespConsumerInfo) SetNumAckPending(v int32) {
	o.NumAckPending = v
}

// GetNumPending returns the NumPending field value
func (o *ApisAPIRestRespConsumerInfo) GetNumPending() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumPending
}

// GetNumPendingOk returns a tuple with the NumPending field value
// and a boolean to check if the value has been set.
func (o *ApisAPIRestRespConsumerInfo) GetNumPendingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumPending, true
}

// SetNumPending sets field value
func (o *ApisAPIRestRespConsumerInfo) SetNumPending(v int32) {
	o.NumPending = v
}

// GetNumRedelivered returns the NumRedelivered field value
func (o *ApisAPIRestRespConsumerInfo) GetNumRedelivered() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumRedelivered
}

// GetNumRedeliveredOk returns a tuple with the NumRedelivered field value
// and a boolean to check if the value has been set.
func (o *ApisAPIRestRespConsumerInfo) GetNumRedeliveredOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumRedelivered, true
}

// SetNumRedelivered sets field value
func (o *ApisAPIRestRespConsumerInfo) SetNumRedelivered(v int32) {
	o.NumRedelivered = v
}

// GetNumWaiting returns the NumWaiting field value
func (o *ApisAPIRestRespConsumerInfo) GetNumWaiting() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumWaiting
}

// GetNumWaitingOk returns a tuple with the NumWaiting field value
// and a boolean to check if the value has been set.
func (o *ApisAPIRestRespConsumerInfo) GetNumWaitingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumWaiting, true
}

// SetNumWaiting sets field value
func (o *ApisAPIRestRespConsumerInfo) SetNumWaiting(v int32) {
	o.NumWaiting = v
}

// GetStreamName returns the StreamName field value
func (o *ApisAPIRestRespConsumerInfo) GetStreamName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StreamName
}

// GetStreamNameOk returns a tuple with the StreamName field value
// and a boolean to check if the value has been set.
func (o *ApisAPIRestRespConsumerInfo) GetStreamNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StreamName, true
}

// SetStreamName sets field value
func (o *ApisAPIRestRespConsumerInfo) SetStreamName(v string) {
	o.StreamName = v
}

func (o ApisAPIRestRespConsumerInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ack_floor"] = o.AckFloor
	}
	if true {
		toSerialize["config"] = o.Config
	}
	if true {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["delivered"] = o.Delivered
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["num_ack_pending"] = o.NumAckPending
	}
	if true {
		toSerialize["num_pending"] = o.NumPending
	}
	if true {
		toSerialize["num_redelivered"] = o.NumRedelivered
	}
	if true {
		toSerialize["num_waiting"] = o.NumWaiting
	}
	if true {
		toSerialize["stream_name"] = o.StreamName
	}
	return json.Marshal(toSerialize)
}

type NullableApisAPIRestRespConsumerInfo struct {
	value *ApisAPIRestRespConsumerInfo
	isSet bool
}

func (v NullableApisAPIRestRespConsumerInfo) Get() *ApisAPIRestRespConsumerInfo {
	return v.value
}

func (v *NullableApisAPIRestRespConsumerInfo) Set(val *ApisAPIRestRespConsumerInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableApisAPIRestRespConsumerInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableApisAPIRestRespConsumerInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApisAPIRestRespConsumerInfo(val *ApisAPIRestRespConsumerInfo) *NullableApisAPIRestRespConsumerInfo {
	return &NullableApisAPIRestRespConsumerInfo{value: val, isSet: true}
}

func (v NullableApisAPIRestRespConsumerInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApisAPIRestRespConsumerInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
