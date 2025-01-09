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

// checks if the PageComponentXO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PageComponentXO{}

// PageComponentXO struct for PageComponentXO
type PageComponentXO struct {
	Items []ComponentXO `json:"items,omitempty"`
	ContinuationToken *string `json:"continuationToken,omitempty"`
}

// NewPageComponentXO instantiates a new PageComponentXO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPageComponentXO() *PageComponentXO {
	this := PageComponentXO{}
	return &this
}

// NewPageComponentXOWithDefaults instantiates a new PageComponentXO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageComponentXOWithDefaults() *PageComponentXO {
	this := PageComponentXO{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *PageComponentXO) GetItems() []ComponentXO {
	if o == nil || IsNil(o.Items) {
		var ret []ComponentXO
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageComponentXO) GetItemsOk() ([]ComponentXO, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *PageComponentXO) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ComponentXO and assigns it to the Items field.
func (o *PageComponentXO) SetItems(v []ComponentXO) {
	o.Items = v
}

// GetContinuationToken returns the ContinuationToken field value if set, zero value otherwise.
func (o *PageComponentXO) GetContinuationToken() string {
	if o == nil || IsNil(o.ContinuationToken) {
		var ret string
		return ret
	}
	return *o.ContinuationToken
}

// GetContinuationTokenOk returns a tuple with the ContinuationToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageComponentXO) GetContinuationTokenOk() (*string, bool) {
	if o == nil || IsNil(o.ContinuationToken) {
		return nil, false
	}
	return o.ContinuationToken, true
}

// HasContinuationToken returns a boolean if a field has been set.
func (o *PageComponentXO) HasContinuationToken() bool {
	if o != nil && !IsNil(o.ContinuationToken) {
		return true
	}

	return false
}

// SetContinuationToken gets a reference to the given string and assigns it to the ContinuationToken field.
func (o *PageComponentXO) SetContinuationToken(v string) {
	o.ContinuationToken = &v
}

func (o PageComponentXO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PageComponentXO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.ContinuationToken) {
		toSerialize["continuationToken"] = o.ContinuationToken
	}
	return toSerialize, nil
}

type NullablePageComponentXO struct {
	value *PageComponentXO
	isSet bool
}

func (v NullablePageComponentXO) Get() *PageComponentXO {
	return v.value
}

func (v *NullablePageComponentXO) Set(val *PageComponentXO) {
	v.value = val
	v.isSet = true
}

func (v NullablePageComponentXO) IsSet() bool {
	return v.isSet
}

func (v *NullablePageComponentXO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePageComponentXO(val *PageComponentXO) *NullablePageComponentXO {
	return &NullablePageComponentXO{value: val, isSet: true}
}

func (v NullablePageComponentXO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePageComponentXO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


