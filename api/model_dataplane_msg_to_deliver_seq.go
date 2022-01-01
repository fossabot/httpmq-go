/*
httpmq

HTTP/2 based message broker built around NATS JetStream

API version: v0.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// DataplaneMsgToDeliverSeq struct for DataplaneMsgToDeliverSeq
type DataplaneMsgToDeliverSeq struct {
	// Consumer is the message sequence number for this consumer
	Consumer int64 `json:"consumer"`
	// Stream is the message sequence number within the stream
	Stream int64 `json:"stream"`
}

// NewDataplaneMsgToDeliverSeq instantiates a new DataplaneMsgToDeliverSeq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataplaneMsgToDeliverSeq(consumer int64, stream int64) *DataplaneMsgToDeliverSeq {
	this := DataplaneMsgToDeliverSeq{}
	this.Consumer = consumer
	this.Stream = stream
	return &this
}

// NewDataplaneMsgToDeliverSeqWithDefaults instantiates a new DataplaneMsgToDeliverSeq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataplaneMsgToDeliverSeqWithDefaults() *DataplaneMsgToDeliverSeq {
	this := DataplaneMsgToDeliverSeq{}
	return &this
}

// GetConsumer returns the Consumer field value
func (o *DataplaneMsgToDeliverSeq) GetConsumer() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Consumer
}

// GetConsumerOk returns a tuple with the Consumer field value
// and a boolean to check if the value has been set.
func (o *DataplaneMsgToDeliverSeq) GetConsumerOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Consumer, true
}

// SetConsumer sets field value
func (o *DataplaneMsgToDeliverSeq) SetConsumer(v int64) {
	o.Consumer = v
}

// GetStream returns the Stream field value
func (o *DataplaneMsgToDeliverSeq) GetStream() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Stream
}

// GetStreamOk returns a tuple with the Stream field value
// and a boolean to check if the value has been set.
func (o *DataplaneMsgToDeliverSeq) GetStreamOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stream, true
}

// SetStream sets field value
func (o *DataplaneMsgToDeliverSeq) SetStream(v int64) {
	o.Stream = v
}

func (o DataplaneMsgToDeliverSeq) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["consumer"] = o.Consumer
	}
	if true {
		toSerialize["stream"] = o.Stream
	}
	return json.Marshal(toSerialize)
}

type NullableDataplaneMsgToDeliverSeq struct {
	value *DataplaneMsgToDeliverSeq
	isSet bool
}

func (v NullableDataplaneMsgToDeliverSeq) Get() *DataplaneMsgToDeliverSeq {
	return v.value
}

func (v *NullableDataplaneMsgToDeliverSeq) Set(val *DataplaneMsgToDeliverSeq) {
	v.value = val
	v.isSet = true
}

func (v NullableDataplaneMsgToDeliverSeq) IsSet() bool {
	return v.isSet
}

func (v *NullableDataplaneMsgToDeliverSeq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataplaneMsgToDeliverSeq(val *DataplaneMsgToDeliverSeq) *NullableDataplaneMsgToDeliverSeq {
	return &NullableDataplaneMsgToDeliverSeq{value: val, isSet: true}
}

func (v NullableDataplaneMsgToDeliverSeq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataplaneMsgToDeliverSeq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
