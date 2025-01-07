/*
Nexus Repository Manager REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.75.1-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nexus

import (
	"encoding/json"
)

// checks if the BaseTagXO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseTagXO{}

// BaseTagXO struct for BaseTagXO
type BaseTagXO struct {
	Attributes map[string]map[string]interface{} `json:"attributes,omitempty"`
}

// NewBaseTagXO instantiates a new BaseTagXO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseTagXO() *BaseTagXO {
	this := BaseTagXO{}
	return &this
}

// NewBaseTagXOWithDefaults instantiates a new BaseTagXO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseTagXOWithDefaults() *BaseTagXO {
	this := BaseTagXO{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *BaseTagXO) GetAttributes() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Attributes) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseTagXO) GetAttributesOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Attributes) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *BaseTagXO) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]map[string]interface{} and assigns it to the Attributes field.
func (o *BaseTagXO) SetAttributes(v map[string]map[string]interface{}) {
	o.Attributes = v
}

func (o BaseTagXO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseTagXO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableBaseTagXO struct {
	value *BaseTagXO
	isSet bool
}

func (v NullableBaseTagXO) Get() *BaseTagXO {
	return v.value
}

func (v *NullableBaseTagXO) Set(val *BaseTagXO) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseTagXO) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseTagXO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseTagXO(val *BaseTagXO) *NullableBaseTagXO {
	return &NullableBaseTagXO{value: val, isSet: true}
}

func (v NullableBaseTagXO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseTagXO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


