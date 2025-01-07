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

// checks if the ConanProxyAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConanProxyAttributes{}

// ConanProxyAttributes struct for ConanProxyAttributes
type ConanProxyAttributes struct {
	// Conan protocol version;
	ConanVersion *string `json:"conanVersion,omitempty"`
}

// NewConanProxyAttributes instantiates a new ConanProxyAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConanProxyAttributes() *ConanProxyAttributes {
	this := ConanProxyAttributes{}
	return &this
}

// NewConanProxyAttributesWithDefaults instantiates a new ConanProxyAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConanProxyAttributesWithDefaults() *ConanProxyAttributes {
	this := ConanProxyAttributes{}
	return &this
}

// GetConanVersion returns the ConanVersion field value if set, zero value otherwise.
func (o *ConanProxyAttributes) GetConanVersion() string {
	if o == nil || IsNil(o.ConanVersion) {
		var ret string
		return ret
	}
	return *o.ConanVersion
}

// GetConanVersionOk returns a tuple with the ConanVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConanProxyAttributes) GetConanVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ConanVersion) {
		return nil, false
	}
	return o.ConanVersion, true
}

// HasConanVersion returns a boolean if a field has been set.
func (o *ConanProxyAttributes) HasConanVersion() bool {
	if o != nil && !IsNil(o.ConanVersion) {
		return true
	}

	return false
}

// SetConanVersion gets a reference to the given string and assigns it to the ConanVersion field.
func (o *ConanProxyAttributes) SetConanVersion(v string) {
	o.ConanVersion = &v
}

func (o ConanProxyAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConanProxyAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConanVersion) {
		toSerialize["conanVersion"] = o.ConanVersion
	}
	return toSerialize, nil
}

type NullableConanProxyAttributes struct {
	value *ConanProxyAttributes
	isSet bool
}

func (v NullableConanProxyAttributes) Get() *ConanProxyAttributes {
	return v.value
}

func (v *NullableConanProxyAttributes) Set(val *ConanProxyAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableConanProxyAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableConanProxyAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConanProxyAttributes(val *ConanProxyAttributes) *NullableConanProxyAttributes {
	return &NullableConanProxyAttributes{value: val, isSet: true}
}

func (v NullableConanProxyAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConanProxyAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


