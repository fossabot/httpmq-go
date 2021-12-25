/*
httpmq

HTTP/2 based message broker built around NATS JetStream

API version: v0.1.2-rc.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ApisAPIRestRespStreamState struct for ApisAPIRestRespStreamState
type ApisAPIRestRespStreamState struct {
	// Bytes is the number of message bytes in the stream
	Bytes int32 `json:"bytes"`
	// Consumers number of consumers on the stream
	ConsumerCount int32 `json:"consumer_count"`
	// FirstSeq is the oldest message sequence number on the stream
	FirstSeq int32 `json:"first_seq"`
	// FirstTime is the oldest message timestamp on the stream
	FirstTs string `json:"first_ts"`
	// LastSeq is the newest message sequence number on the stream
	LastSeq int32 `json:"last_seq"`
	// LastTime is the newest message timestamp on the stream
	LastTs string `json:"last_ts"`
	// Msgs is the number of messages in the stream
	Messages int32 `json:"messages"`
}

// NewApisAPIRestRespStreamState instantiates a new ApisAPIRestRespStreamState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApisAPIRestRespStreamState(bytes int32, consumerCount int32, firstSeq int32, firstTs string, lastSeq int32, lastTs string, messages int32) *ApisAPIRestRespStreamState {
	this := ApisAPIRestRespStreamState{}
	this.Bytes = bytes
	this.ConsumerCount = consumerCount
	this.FirstSeq = firstSeq
	this.FirstTs = firstTs
	this.LastSeq = lastSeq
	this.LastTs = lastTs
	this.Messages = messages
	return &this
}

// NewApisAPIRestRespStreamStateWithDefaults instantiates a new ApisAPIRestRespStreamState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApisAPIRestRespStreamStateWithDefaults() *ApisAPIRestRespStreamState {
	this := ApisAPIRestRespStreamState{}
	return &this
}

// GetBytes returns the Bytes field value
func (o *ApisAPIRestRespStreamState) GetBytes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Bytes
}

// GetBytesOk returns a tuple with the Bytes field value
// and a boolean to check if the value has been set.
func (o *ApisAPIRestRespStreamState) GetBytesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bytes, true
}

// SetBytes sets field value
func (o *ApisAPIRestRespStreamState) SetBytes(v int32) {
	o.Bytes = v
}

// GetConsumerCount returns the ConsumerCount field value
func (o *ApisAPIRestRespStreamState) GetConsumerCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ConsumerCount
}

// GetConsumerCountOk returns a tuple with the ConsumerCount field value
// and a boolean to check if the value has been set.
func (o *ApisAPIRestRespStreamState) GetConsumerCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConsumerCount, true
}

// SetConsumerCount sets field value
func (o *ApisAPIRestRespStreamState) SetConsumerCount(v int32) {
	o.ConsumerCount = v
}

// GetFirstSeq returns the FirstSeq field value
func (o *ApisAPIRestRespStreamState) GetFirstSeq() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FirstSeq
}

// GetFirstSeqOk returns a tuple with the FirstSeq field value
// and a boolean to check if the value has been set.
func (o *ApisAPIRestRespStreamState) GetFirstSeqOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstSeq, true
}

// SetFirstSeq sets field value
func (o *ApisAPIRestRespStreamState) SetFirstSeq(v int32) {
	o.FirstSeq = v
}

// GetFirstTs returns the FirstTs field value
func (o *ApisAPIRestRespStreamState) GetFirstTs() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstTs
}

// GetFirstTsOk returns a tuple with the FirstTs field value
// and a boolean to check if the value has been set.
func (o *ApisAPIRestRespStreamState) GetFirstTsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstTs, true
}

// SetFirstTs sets field value
func (o *ApisAPIRestRespStreamState) SetFirstTs(v string) {
	o.FirstTs = v
}

// GetLastSeq returns the LastSeq field value
func (o *ApisAPIRestRespStreamState) GetLastSeq() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LastSeq
}

// GetLastSeqOk returns a tuple with the LastSeq field value
// and a boolean to check if the value has been set.
func (o *ApisAPIRestRespStreamState) GetLastSeqOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastSeq, true
}

// SetLastSeq sets field value
func (o *ApisAPIRestRespStreamState) SetLastSeq(v int32) {
	o.LastSeq = v
}

// GetLastTs returns the LastTs field value
func (o *ApisAPIRestRespStreamState) GetLastTs() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastTs
}

// GetLastTsOk returns a tuple with the LastTs field value
// and a boolean to check if the value has been set.
func (o *ApisAPIRestRespStreamState) GetLastTsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastTs, true
}

// SetLastTs sets field value
func (o *ApisAPIRestRespStreamState) SetLastTs(v string) {
	o.LastTs = v
}

// GetMessages returns the Messages field value
func (o *ApisAPIRestRespStreamState) GetMessages() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value
// and a boolean to check if the value has been set.
func (o *ApisAPIRestRespStreamState) GetMessagesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Messages, true
}

// SetMessages sets field value
func (o *ApisAPIRestRespStreamState) SetMessages(v int32) {
	o.Messages = v
}

func (o ApisAPIRestRespStreamState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bytes"] = o.Bytes
	}
	if true {
		toSerialize["consumer_count"] = o.ConsumerCount
	}
	if true {
		toSerialize["first_seq"] = o.FirstSeq
	}
	if true {
		toSerialize["first_ts"] = o.FirstTs
	}
	if true {
		toSerialize["last_seq"] = o.LastSeq
	}
	if true {
		toSerialize["last_ts"] = o.LastTs
	}
	if true {
		toSerialize["messages"] = o.Messages
	}
	return json.Marshal(toSerialize)
}

type NullableApisAPIRestRespStreamState struct {
	value *ApisAPIRestRespStreamState
	isSet bool
}

func (v NullableApisAPIRestRespStreamState) Get() *ApisAPIRestRespStreamState {
	return v.value
}

func (v *NullableApisAPIRestRespStreamState) Set(val *ApisAPIRestRespStreamState) {
	v.value = val
	v.isSet = true
}

func (v NullableApisAPIRestRespStreamState) IsSet() bool {
	return v.isSet
}

func (v *NullableApisAPIRestRespStreamState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApisAPIRestRespStreamState(val *ApisAPIRestRespStreamState) *NullableApisAPIRestRespStreamState {
	return &NullableApisAPIRestRespStreamState{value: val, isSet: true}
}

func (v NullableApisAPIRestRespStreamState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApisAPIRestRespStreamState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
