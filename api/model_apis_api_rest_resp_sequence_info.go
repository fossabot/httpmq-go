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

// ApisAPIRestRespSequenceInfo struct for ApisAPIRestRespSequenceInfo
type ApisAPIRestRespSequenceInfo struct {
	// Consumer is consumer level sequence number
	ConsumerSeq int32 `json:"consumer_seq"`
	// Last timestamp when these values updated
	LastActive *string `json:"last_active,omitempty"`
	// Stream is stream level sequence number
	StreamSeq int32 `json:"stream_seq"`
}

// NewApisAPIRestRespSequenceInfo instantiates a new ApisAPIRestRespSequenceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApisAPIRestRespSequenceInfo(consumerSeq int32, streamSeq int32) *ApisAPIRestRespSequenceInfo {
	this := ApisAPIRestRespSequenceInfo{}
	this.ConsumerSeq = consumerSeq
	this.StreamSeq = streamSeq
	return &this
}

// NewApisAPIRestRespSequenceInfoWithDefaults instantiates a new ApisAPIRestRespSequenceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApisAPIRestRespSequenceInfoWithDefaults() *ApisAPIRestRespSequenceInfo {
	this := ApisAPIRestRespSequenceInfo{}
	return &this
}

// GetConsumerSeq returns the ConsumerSeq field value
func (o *ApisAPIRestRespSequenceInfo) GetConsumerSeq() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ConsumerSeq
}

// GetConsumerSeqOk returns a tuple with the ConsumerSeq field value
// and a boolean to check if the value has been set.
func (o *ApisAPIRestRespSequenceInfo) GetConsumerSeqOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConsumerSeq, true
}

// SetConsumerSeq sets field value
func (o *ApisAPIRestRespSequenceInfo) SetConsumerSeq(v int32) {
	o.ConsumerSeq = v
}

// GetLastActive returns the LastActive field value if set, zero value otherwise.
func (o *ApisAPIRestRespSequenceInfo) GetLastActive() string {
	if o == nil || o.LastActive == nil {
		var ret string
		return ret
	}
	return *o.LastActive
}

// GetLastActiveOk returns a tuple with the LastActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApisAPIRestRespSequenceInfo) GetLastActiveOk() (*string, bool) {
	if o == nil || o.LastActive == nil {
		return nil, false
	}
	return o.LastActive, true
}

// HasLastActive returns a boolean if a field has been set.
func (o *ApisAPIRestRespSequenceInfo) HasLastActive() bool {
	if o != nil && o.LastActive != nil {
		return true
	}

	return false
}

// SetLastActive gets a reference to the given string and assigns it to the LastActive field.
func (o *ApisAPIRestRespSequenceInfo) SetLastActive(v string) {
	o.LastActive = &v
}

// GetStreamSeq returns the StreamSeq field value
func (o *ApisAPIRestRespSequenceInfo) GetStreamSeq() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StreamSeq
}

// GetStreamSeqOk returns a tuple with the StreamSeq field value
// and a boolean to check if the value has been set.
func (o *ApisAPIRestRespSequenceInfo) GetStreamSeqOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StreamSeq, true
}

// SetStreamSeq sets field value
func (o *ApisAPIRestRespSequenceInfo) SetStreamSeq(v int32) {
	o.StreamSeq = v
}

func (o ApisAPIRestRespSequenceInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["consumer_seq"] = o.ConsumerSeq
	}
	if o.LastActive != nil {
		toSerialize["last_active"] = o.LastActive
	}
	if true {
		toSerialize["stream_seq"] = o.StreamSeq
	}
	return json.Marshal(toSerialize)
}

type NullableApisAPIRestRespSequenceInfo struct {
	value *ApisAPIRestRespSequenceInfo
	isSet bool
}

func (v NullableApisAPIRestRespSequenceInfo) Get() *ApisAPIRestRespSequenceInfo {
	return v.value
}

func (v *NullableApisAPIRestRespSequenceInfo) Set(val *ApisAPIRestRespSequenceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableApisAPIRestRespSequenceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableApisAPIRestRespSequenceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApisAPIRestRespSequenceInfo(val *ApisAPIRestRespSequenceInfo) *NullableApisAPIRestRespSequenceInfo {
	return &NullableApisAPIRestRespSequenceInfo{value: val, isSet: true}
}

func (v NullableApisAPIRestRespSequenceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApisAPIRestRespSequenceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
