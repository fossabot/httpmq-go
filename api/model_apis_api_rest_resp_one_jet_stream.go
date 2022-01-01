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

// ApisAPIRestRespOneJetStream struct for ApisAPIRestRespOneJetStream
type ApisAPIRestRespOneJetStream struct {
	Error  *ApisErrorDetail           `json:"error,omitempty"`
	Stream *ApisAPIRestRespStreamInfo `json:"stream,omitempty"`
	// Success indicates whether the request was successful
	Success bool `json:"success"`
}

// NewApisAPIRestRespOneJetStream instantiates a new ApisAPIRestRespOneJetStream object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApisAPIRestRespOneJetStream(success bool) *ApisAPIRestRespOneJetStream {
	this := ApisAPIRestRespOneJetStream{}
	this.Success = success
	return &this
}

// NewApisAPIRestRespOneJetStreamWithDefaults instantiates a new ApisAPIRestRespOneJetStream object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApisAPIRestRespOneJetStreamWithDefaults() *ApisAPIRestRespOneJetStream {
	this := ApisAPIRestRespOneJetStream{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ApisAPIRestRespOneJetStream) GetError() ApisErrorDetail {
	if o == nil || o.Error == nil {
		var ret ApisErrorDetail
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApisAPIRestRespOneJetStream) GetErrorOk() (*ApisErrorDetail, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ApisAPIRestRespOneJetStream) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given ApisErrorDetail and assigns it to the Error field.
func (o *ApisAPIRestRespOneJetStream) SetError(v ApisErrorDetail) {
	o.Error = &v
}

// GetStream returns the Stream field value if set, zero value otherwise.
func (o *ApisAPIRestRespOneJetStream) GetStream() ApisAPIRestRespStreamInfo {
	if o == nil || o.Stream == nil {
		var ret ApisAPIRestRespStreamInfo
		return ret
	}
	return *o.Stream
}

// GetStreamOk returns a tuple with the Stream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApisAPIRestRespOneJetStream) GetStreamOk() (*ApisAPIRestRespStreamInfo, bool) {
	if o == nil || o.Stream == nil {
		return nil, false
	}
	return o.Stream, true
}

// HasStream returns a boolean if a field has been set.
func (o *ApisAPIRestRespOneJetStream) HasStream() bool {
	if o != nil && o.Stream != nil {
		return true
	}

	return false
}

// SetStream gets a reference to the given ApisAPIRestRespStreamInfo and assigns it to the Stream field.
func (o *ApisAPIRestRespOneJetStream) SetStream(v ApisAPIRestRespStreamInfo) {
	o.Stream = &v
}

// GetSuccess returns the Success field value
func (o *ApisAPIRestRespOneJetStream) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *ApisAPIRestRespOneJetStream) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *ApisAPIRestRespOneJetStream) SetSuccess(v bool) {
	o.Success = v
}

func (o ApisAPIRestRespOneJetStream) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.Stream != nil {
		toSerialize["stream"] = o.Stream
	}
	if true {
		toSerialize["success"] = o.Success
	}
	return json.Marshal(toSerialize)
}

type NullableApisAPIRestRespOneJetStream struct {
	value *ApisAPIRestRespOneJetStream
	isSet bool
}

func (v NullableApisAPIRestRespOneJetStream) Get() *ApisAPIRestRespOneJetStream {
	return v.value
}

func (v *NullableApisAPIRestRespOneJetStream) Set(val *ApisAPIRestRespOneJetStream) {
	v.value = val
	v.isSet = true
}

func (v NullableApisAPIRestRespOneJetStream) IsSet() bool {
	return v.isSet
}

func (v *NullableApisAPIRestRespOneJetStream) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApisAPIRestRespOneJetStream(val *ApisAPIRestRespOneJetStream) *NullableApisAPIRestRespOneJetStream {
	return &NullableApisAPIRestRespOneJetStream{value: val, isSet: true}
}

func (v NullableApisAPIRestRespOneJetStream) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApisAPIRestRespOneJetStream) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
