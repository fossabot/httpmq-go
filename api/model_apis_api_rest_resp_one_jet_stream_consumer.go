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

// ApisAPIRestRespOneJetStreamConsumer struct for ApisAPIRestRespOneJetStreamConsumer
type ApisAPIRestRespOneJetStreamConsumer struct {
	Consumer *ApisAPIRestRespConsumerInfo `json:"consumer,omitempty"`
	Error    *ApisErrorDetail             `json:"error,omitempty"`
	// Success indicates whether the request was successful
	Success bool `json:"success"`
}

// NewApisAPIRestRespOneJetStreamConsumer instantiates a new ApisAPIRestRespOneJetStreamConsumer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApisAPIRestRespOneJetStreamConsumer(success bool) *ApisAPIRestRespOneJetStreamConsumer {
	this := ApisAPIRestRespOneJetStreamConsumer{}
	this.Success = success
	return &this
}

// NewApisAPIRestRespOneJetStreamConsumerWithDefaults instantiates a new ApisAPIRestRespOneJetStreamConsumer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApisAPIRestRespOneJetStreamConsumerWithDefaults() *ApisAPIRestRespOneJetStreamConsumer {
	this := ApisAPIRestRespOneJetStreamConsumer{}
	return &this
}

// GetConsumer returns the Consumer field value if set, zero value otherwise.
func (o *ApisAPIRestRespOneJetStreamConsumer) GetConsumer() ApisAPIRestRespConsumerInfo {
	if o == nil || o.Consumer == nil {
		var ret ApisAPIRestRespConsumerInfo
		return ret
	}
	return *o.Consumer
}

// GetConsumerOk returns a tuple with the Consumer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApisAPIRestRespOneJetStreamConsumer) GetConsumerOk() (*ApisAPIRestRespConsumerInfo, bool) {
	if o == nil || o.Consumer == nil {
		return nil, false
	}
	return o.Consumer, true
}

// HasConsumer returns a boolean if a field has been set.
func (o *ApisAPIRestRespOneJetStreamConsumer) HasConsumer() bool {
	if o != nil && o.Consumer != nil {
		return true
	}

	return false
}

// SetConsumer gets a reference to the given ApisAPIRestRespConsumerInfo and assigns it to the Consumer field.
func (o *ApisAPIRestRespOneJetStreamConsumer) SetConsumer(v ApisAPIRestRespConsumerInfo) {
	o.Consumer = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ApisAPIRestRespOneJetStreamConsumer) GetError() ApisErrorDetail {
	if o == nil || o.Error == nil {
		var ret ApisErrorDetail
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApisAPIRestRespOneJetStreamConsumer) GetErrorOk() (*ApisErrorDetail, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ApisAPIRestRespOneJetStreamConsumer) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given ApisErrorDetail and assigns it to the Error field.
func (o *ApisAPIRestRespOneJetStreamConsumer) SetError(v ApisErrorDetail) {
	o.Error = &v
}

// GetSuccess returns the Success field value
func (o *ApisAPIRestRespOneJetStreamConsumer) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *ApisAPIRestRespOneJetStreamConsumer) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *ApisAPIRestRespOneJetStreamConsumer) SetSuccess(v bool) {
	o.Success = v
}

func (o ApisAPIRestRespOneJetStreamConsumer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Consumer != nil {
		toSerialize["consumer"] = o.Consumer
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if true {
		toSerialize["success"] = o.Success
	}
	return json.Marshal(toSerialize)
}

type NullableApisAPIRestRespOneJetStreamConsumer struct {
	value *ApisAPIRestRespOneJetStreamConsumer
	isSet bool
}

func (v NullableApisAPIRestRespOneJetStreamConsumer) Get() *ApisAPIRestRespOneJetStreamConsumer {
	return v.value
}

func (v *NullableApisAPIRestRespOneJetStreamConsumer) Set(val *ApisAPIRestRespOneJetStreamConsumer) {
	v.value = val
	v.isSet = true
}

func (v NullableApisAPIRestRespOneJetStreamConsumer) IsSet() bool {
	return v.isSet
}

func (v *NullableApisAPIRestRespOneJetStreamConsumer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApisAPIRestRespOneJetStreamConsumer(val *ApisAPIRestRespOneJetStreamConsumer) *NullableApisAPIRestRespOneJetStreamConsumer {
	return &NullableApisAPIRestRespOneJetStreamConsumer{value: val, isSet: true}
}

func (v NullableApisAPIRestRespOneJetStreamConsumer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApisAPIRestRespOneJetStreamConsumer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
