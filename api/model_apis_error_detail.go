/*
httpmq

HTTP/2 based message broker built around NATS JetStream

API version: v0.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ApisErrorDetail struct for ApisErrorDetail
type ApisErrorDetail struct {
	// Code is the response code
	Code int32 `json:"code"`
	// Msg is an optional descriptive message
	Message *string `json:"message,omitempty"`
}

// NewApisErrorDetail instantiates a new ApisErrorDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApisErrorDetail(code int32) *ApisErrorDetail {
	this := ApisErrorDetail{}
	this.Code = code
	return &this
}

// NewApisErrorDetailWithDefaults instantiates a new ApisErrorDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApisErrorDetailWithDefaults() *ApisErrorDetail {
	this := ApisErrorDetail{}
	return &this
}

// GetCode returns the Code field value
func (o *ApisErrorDetail) GetCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ApisErrorDetail) GetCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ApisErrorDetail) SetCode(v int32) {
	o.Code = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ApisErrorDetail) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApisErrorDetail) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ApisErrorDetail) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ApisErrorDetail) SetMessage(v string) {
	o.Message = &v
}

func (o ApisErrorDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["code"] = o.Code
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableApisErrorDetail struct {
	value *ApisErrorDetail
	isSet bool
}

func (v NullableApisErrorDetail) Get() *ApisErrorDetail {
	return v.value
}

func (v *NullableApisErrorDetail) Set(val *ApisErrorDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableApisErrorDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableApisErrorDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApisErrorDetail(val *ApisErrorDetail) *NullableApisErrorDetail {
	return &NullableApisErrorDetail{value: val, isSet: true}
}

func (v NullableApisErrorDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApisErrorDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
