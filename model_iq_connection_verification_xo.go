/*
Nexus Repository Manager REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.76.0-03
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nexus

import (
	"encoding/json"
)

// checks if the IqConnectionVerificationXo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IqConnectionVerificationXo{}

// IqConnectionVerificationXo struct for IqConnectionVerificationXo
type IqConnectionVerificationXo struct {
	Success *bool `json:"success,omitempty"`
	Reason *string `json:"reason,omitempty"`
}

// NewIqConnectionVerificationXo instantiates a new IqConnectionVerificationXo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIqConnectionVerificationXo() *IqConnectionVerificationXo {
	this := IqConnectionVerificationXo{}
	return &this
}

// NewIqConnectionVerificationXoWithDefaults instantiates a new IqConnectionVerificationXo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIqConnectionVerificationXoWithDefaults() *IqConnectionVerificationXo {
	this := IqConnectionVerificationXo{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *IqConnectionVerificationXo) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IqConnectionVerificationXo) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *IqConnectionVerificationXo) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *IqConnectionVerificationXo) SetSuccess(v bool) {
	o.Success = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *IqConnectionVerificationXo) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IqConnectionVerificationXo) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *IqConnectionVerificationXo) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *IqConnectionVerificationXo) SetReason(v string) {
	o.Reason = &v
}

func (o IqConnectionVerificationXo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IqConnectionVerificationXo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return toSerialize, nil
}

type NullableIqConnectionVerificationXo struct {
	value *IqConnectionVerificationXo
	isSet bool
}

func (v NullableIqConnectionVerificationXo) Get() *IqConnectionVerificationXo {
	return v.value
}

func (v *NullableIqConnectionVerificationXo) Set(val *IqConnectionVerificationXo) {
	v.value = val
	v.isSet = true
}

func (v NullableIqConnectionVerificationXo) IsSet() bool {
	return v.isSet
}

func (v *NullableIqConnectionVerificationXo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIqConnectionVerificationXo(val *IqConnectionVerificationXo) *NullableIqConnectionVerificationXo {
	return &NullableIqConnectionVerificationXo{value: val, isSet: true}
}

func (v NullableIqConnectionVerificationXo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIqConnectionVerificationXo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


