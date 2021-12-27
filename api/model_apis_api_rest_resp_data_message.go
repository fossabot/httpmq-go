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

// ApisAPIRestRespDataMessage struct for ApisAPIRestRespDataMessage
type ApisAPIRestRespDataMessage struct {
	// Message is the message body
	B64Msg []byte `json:"b64_msg"`
	// Consumer is the name of the consumer
	Consumer string                   `json:"consumer"`
	Error    *ApisErrorDetail         `json:"error,omitempty"`
	Sequence DataplaneMsgToDeliverSeq `json:"sequence"`
	// Stream is the name of the stream
	Stream string `json:"stream"`
	// Subject is the name of the subject / subject filter
	Subject string `json:"subject"`
	// Success indicates whether the request was successful
	Success bool `json:"success"`
}

// NewApisAPIRestRespDataMessage instantiates a new ApisAPIRestRespDataMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApisAPIRestRespDataMessage(b64Msg []byte, consumer string, sequence DataplaneMsgToDeliverSeq, stream string, subject string, success bool) *ApisAPIRestRespDataMessage {
	this := ApisAPIRestRespDataMessage{}
	this.B64Msg = b64Msg
	this.Consumer = consumer
	this.Sequence = sequence
	this.Stream = stream
	this.Subject = subject
	this.Success = success
	return &this
}

// NewApisAPIRestRespDataMessageWithDefaults instantiates a new ApisAPIRestRespDataMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApisAPIRestRespDataMessageWithDefaults() *ApisAPIRestRespDataMessage {
	this := ApisAPIRestRespDataMessage{}
	return &this
}

// GetB64Msg returns the B64Msg field value
func (o *ApisAPIRestRespDataMessage) GetB64Msg() []byte {
	if o == nil {
		var ret []byte
		return ret
	}

	return o.B64Msg
}

// GetB64MsgOk returns a tuple with the B64Msg field value
// and a boolean to check if the value has been set.
func (o *ApisAPIRestRespDataMessage) GetB64MsgOk() (*[]byte, bool) {
	if o == nil {
		return nil, false
	}
	return &o.B64Msg, true
}

// SetB64Msg sets field value
func (o *ApisAPIRestRespDataMessage) SetB64Msg(v []byte) {
	o.B64Msg = v
}

// GetConsumer returns the Consumer field value
func (o *ApisAPIRestRespDataMessage) GetConsumer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Consumer
}

// GetConsumerOk returns a tuple with the Consumer field value
// and a boolean to check if the value has been set.
func (o *ApisAPIRestRespDataMessage) GetConsumerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Consumer, true
}

// SetConsumer sets field value
func (o *ApisAPIRestRespDataMessage) SetConsumer(v string) {
	o.Consumer = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ApisAPIRestRespDataMessage) GetError() ApisErrorDetail {
	if o == nil || o.Error == nil {
		var ret ApisErrorDetail
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApisAPIRestRespDataMessage) GetErrorOk() (*ApisErrorDetail, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ApisAPIRestRespDataMessage) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given ApisErrorDetail and assigns it to the Error field.
func (o *ApisAPIRestRespDataMessage) SetError(v ApisErrorDetail) {
	o.Error = &v
}

// GetSequence returns the Sequence field value
func (o *ApisAPIRestRespDataMessage) GetSequence() DataplaneMsgToDeliverSeq {
	if o == nil {
		var ret DataplaneMsgToDeliverSeq
		return ret
	}

	return o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value
// and a boolean to check if the value has been set.
func (o *ApisAPIRestRespDataMessage) GetSequenceOk() (*DataplaneMsgToDeliverSeq, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sequence, true
}

// SetSequence sets field value
func (o *ApisAPIRestRespDataMessage) SetSequence(v DataplaneMsgToDeliverSeq) {
	o.Sequence = v
}

// GetStream returns the Stream field value
func (o *ApisAPIRestRespDataMessage) GetStream() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Stream
}

// GetStreamOk returns a tuple with the Stream field value
// and a boolean to check if the value has been set.
func (o *ApisAPIRestRespDataMessage) GetStreamOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stream, true
}

// SetStream sets field value
func (o *ApisAPIRestRespDataMessage) SetStream(v string) {
	o.Stream = v
}

// GetSubject returns the Subject field value
func (o *ApisAPIRestRespDataMessage) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *ApisAPIRestRespDataMessage) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *ApisAPIRestRespDataMessage) SetSubject(v string) {
	o.Subject = v
}

// GetSuccess returns the Success field value
func (o *ApisAPIRestRespDataMessage) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *ApisAPIRestRespDataMessage) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *ApisAPIRestRespDataMessage) SetSuccess(v bool) {
	o.Success = v
}

func (o ApisAPIRestRespDataMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["b64_msg"] = o.B64Msg
	}
	if true {
		toSerialize["consumer"] = o.Consumer
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if true {
		toSerialize["sequence"] = o.Sequence
	}
	if true {
		toSerialize["stream"] = o.Stream
	}
	if true {
		toSerialize["subject"] = o.Subject
	}
	if true {
		toSerialize["success"] = o.Success
	}
	return json.Marshal(toSerialize)
}

type NullableApisAPIRestRespDataMessage struct {
	value *ApisAPIRestRespDataMessage
	isSet bool
}

func (v NullableApisAPIRestRespDataMessage) Get() *ApisAPIRestRespDataMessage {
	return v.value
}

func (v *NullableApisAPIRestRespDataMessage) Set(val *ApisAPIRestRespDataMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableApisAPIRestRespDataMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableApisAPIRestRespDataMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApisAPIRestRespDataMessage(val *ApisAPIRestRespDataMessage) *NullableApisAPIRestRespDataMessage {
	return &NullableApisAPIRestRespDataMessage{value: val, isSet: true}
}

func (v NullableApisAPIRestRespDataMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApisAPIRestRespDataMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
