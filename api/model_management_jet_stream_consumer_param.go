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

// ManagementJetStreamConsumerParam struct for ManagementJetStreamConsumerParam
type ManagementJetStreamConsumerParam struct {
	// AckWait when specified, the number of ns to wait for ACK before retry
	AckWait *int64 `json:"ack_wait,omitempty"`
	// DeliveryGroup creates a consumer using a delivery group name.  A consumer using delivery group allows multiple clients to subscribe under the same consumer and group name tuple. For subjects this consumer listens to, the messages will be shared amongst the connected clients.
	DeliveryGroup *string `json:"delivery_group,omitempty"`
	// FilterSubject sets the consumer to filter for subjects matching this NATs subject string  See https://docs.nats.io/nats-concepts/subjects
	FilterSubject *string `json:"filter_subject,omitempty"`
	// MaxInflight is max number of un-ACKed message permitted in-flight (must be >= 1)
	MaxInflight int64 `json:"max_inflight"`
	// MaxRetry max number of times an un-ACKed message is resent (-1: infinite)
	MaxRetry *int64 `json:"max_retry,omitempty"`
	// Mode whether the consumer is push or pull consumer
	Mode string `json:"mode"`
	// Name is the consumer name
	Name string `json:"name"`
	// Notes are descriptions regarding this consumer
	Notes *string `json:"notes,omitempty"`
}

// NewManagementJetStreamConsumerParam instantiates a new ManagementJetStreamConsumerParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementJetStreamConsumerParam(maxInflight int64, mode string, name string) *ManagementJetStreamConsumerParam {
	this := ManagementJetStreamConsumerParam{}
	this.MaxInflight = maxInflight
	this.Mode = mode
	this.Name = name
	return &this
}

// NewManagementJetStreamConsumerParamWithDefaults instantiates a new ManagementJetStreamConsumerParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementJetStreamConsumerParamWithDefaults() *ManagementJetStreamConsumerParam {
	this := ManagementJetStreamConsumerParam{}
	return &this
}

// GetAckWait returns the AckWait field value if set, zero value otherwise.
func (o *ManagementJetStreamConsumerParam) GetAckWait() int64 {
	if o == nil || o.AckWait == nil {
		var ret int64
		return ret
	}
	return *o.AckWait
}

// GetAckWaitOk returns a tuple with the AckWait field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementJetStreamConsumerParam) GetAckWaitOk() (*int64, bool) {
	if o == nil || o.AckWait == nil {
		return nil, false
	}
	return o.AckWait, true
}

// HasAckWait returns a boolean if a field has been set.
func (o *ManagementJetStreamConsumerParam) HasAckWait() bool {
	if o != nil && o.AckWait != nil {
		return true
	}

	return false
}

// SetAckWait gets a reference to the given int64 and assigns it to the AckWait field.
func (o *ManagementJetStreamConsumerParam) SetAckWait(v int64) {
	o.AckWait = &v
}

// GetDeliveryGroup returns the DeliveryGroup field value if set, zero value otherwise.
func (o *ManagementJetStreamConsumerParam) GetDeliveryGroup() string {
	if o == nil || o.DeliveryGroup == nil {
		var ret string
		return ret
	}
	return *o.DeliveryGroup
}

// GetDeliveryGroupOk returns a tuple with the DeliveryGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementJetStreamConsumerParam) GetDeliveryGroupOk() (*string, bool) {
	if o == nil || o.DeliveryGroup == nil {
		return nil, false
	}
	return o.DeliveryGroup, true
}

// HasDeliveryGroup returns a boolean if a field has been set.
func (o *ManagementJetStreamConsumerParam) HasDeliveryGroup() bool {
	if o != nil && o.DeliveryGroup != nil {
		return true
	}

	return false
}

// SetDeliveryGroup gets a reference to the given string and assigns it to the DeliveryGroup field.
func (o *ManagementJetStreamConsumerParam) SetDeliveryGroup(v string) {
	o.DeliveryGroup = &v
}

// GetFilterSubject returns the FilterSubject field value if set, zero value otherwise.
func (o *ManagementJetStreamConsumerParam) GetFilterSubject() string {
	if o == nil || o.FilterSubject == nil {
		var ret string
		return ret
	}
	return *o.FilterSubject
}

// GetFilterSubjectOk returns a tuple with the FilterSubject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementJetStreamConsumerParam) GetFilterSubjectOk() (*string, bool) {
	if o == nil || o.FilterSubject == nil {
		return nil, false
	}
	return o.FilterSubject, true
}

// HasFilterSubject returns a boolean if a field has been set.
func (o *ManagementJetStreamConsumerParam) HasFilterSubject() bool {
	if o != nil && o.FilterSubject != nil {
		return true
	}

	return false
}

// SetFilterSubject gets a reference to the given string and assigns it to the FilterSubject field.
func (o *ManagementJetStreamConsumerParam) SetFilterSubject(v string) {
	o.FilterSubject = &v
}

// GetMaxInflight returns the MaxInflight field value
func (o *ManagementJetStreamConsumerParam) GetMaxInflight() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MaxInflight
}

// GetMaxInflightOk returns a tuple with the MaxInflight field value
// and a boolean to check if the value has been set.
func (o *ManagementJetStreamConsumerParam) GetMaxInflightOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxInflight, true
}

// SetMaxInflight sets field value
func (o *ManagementJetStreamConsumerParam) SetMaxInflight(v int64) {
	o.MaxInflight = v
}

// GetMaxRetry returns the MaxRetry field value if set, zero value otherwise.
func (o *ManagementJetStreamConsumerParam) GetMaxRetry() int64 {
	if o == nil || o.MaxRetry == nil {
		var ret int64
		return ret
	}
	return *o.MaxRetry
}

// GetMaxRetryOk returns a tuple with the MaxRetry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementJetStreamConsumerParam) GetMaxRetryOk() (*int64, bool) {
	if o == nil || o.MaxRetry == nil {
		return nil, false
	}
	return o.MaxRetry, true
}

// HasMaxRetry returns a boolean if a field has been set.
func (o *ManagementJetStreamConsumerParam) HasMaxRetry() bool {
	if o != nil && o.MaxRetry != nil {
		return true
	}

	return false
}

// SetMaxRetry gets a reference to the given int64 and assigns it to the MaxRetry field.
func (o *ManagementJetStreamConsumerParam) SetMaxRetry(v int64) {
	o.MaxRetry = &v
}

// GetMode returns the Mode field value
func (o *ManagementJetStreamConsumerParam) GetMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *ManagementJetStreamConsumerParam) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *ManagementJetStreamConsumerParam) SetMode(v string) {
	o.Mode = v
}

// GetName returns the Name field value
func (o *ManagementJetStreamConsumerParam) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ManagementJetStreamConsumerParam) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ManagementJetStreamConsumerParam) SetName(v string) {
	o.Name = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *ManagementJetStreamConsumerParam) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementJetStreamConsumerParam) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *ManagementJetStreamConsumerParam) HasNotes() bool {
	if o != nil && o.Notes != nil {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *ManagementJetStreamConsumerParam) SetNotes(v string) {
	o.Notes = &v
}

func (o ManagementJetStreamConsumerParam) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AckWait != nil {
		toSerialize["ack_wait"] = o.AckWait
	}
	if o.DeliveryGroup != nil {
		toSerialize["delivery_group"] = o.DeliveryGroup
	}
	if o.FilterSubject != nil {
		toSerialize["filter_subject"] = o.FilterSubject
	}
	if true {
		toSerialize["max_inflight"] = o.MaxInflight
	}
	if o.MaxRetry != nil {
		toSerialize["max_retry"] = o.MaxRetry
	}
	if true {
		toSerialize["mode"] = o.Mode
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	return json.Marshal(toSerialize)
}

type NullableManagementJetStreamConsumerParam struct {
	value *ManagementJetStreamConsumerParam
	isSet bool
}

func (v NullableManagementJetStreamConsumerParam) Get() *ManagementJetStreamConsumerParam {
	return v.value
}

func (v *NullableManagementJetStreamConsumerParam) Set(val *ManagementJetStreamConsumerParam) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementJetStreamConsumerParam) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementJetStreamConsumerParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementJetStreamConsumerParam(val *ManagementJetStreamConsumerParam) *NullableManagementJetStreamConsumerParam {
	return &NullableManagementJetStreamConsumerParam{value: val, isSet: true}
}

func (v NullableManagementJetStreamConsumerParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementJetStreamConsumerParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
