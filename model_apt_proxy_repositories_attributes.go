/*
Nexus Repository Manager REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.76.0-03
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nexus

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AptProxyRepositoriesAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AptProxyRepositoriesAttributes{}

// AptProxyRepositoriesAttributes struct for AptProxyRepositoriesAttributes
type AptProxyRepositoriesAttributes struct {
	// Distribution to fetch
	Distribution *string `json:"distribution,omitempty"`
	// Whether this repository is flat
	Flat bool `json:"flat"`
}

type _AptProxyRepositoriesAttributes AptProxyRepositoriesAttributes

// NewAptProxyRepositoriesAttributes instantiates a new AptProxyRepositoriesAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAptProxyRepositoriesAttributes(flat bool) *AptProxyRepositoriesAttributes {
	this := AptProxyRepositoriesAttributes{}
	this.Flat = flat
	return &this
}

// NewAptProxyRepositoriesAttributesWithDefaults instantiates a new AptProxyRepositoriesAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAptProxyRepositoriesAttributesWithDefaults() *AptProxyRepositoriesAttributes {
	this := AptProxyRepositoriesAttributes{}
	return &this
}

// GetDistribution returns the Distribution field value if set, zero value otherwise.
func (o *AptProxyRepositoriesAttributes) GetDistribution() string {
	if o == nil || IsNil(o.Distribution) {
		var ret string
		return ret
	}
	return *o.Distribution
}

// GetDistributionOk returns a tuple with the Distribution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AptProxyRepositoriesAttributes) GetDistributionOk() (*string, bool) {
	if o == nil || IsNil(o.Distribution) {
		return nil, false
	}
	return o.Distribution, true
}

// HasDistribution returns a boolean if a field has been set.
func (o *AptProxyRepositoriesAttributes) HasDistribution() bool {
	if o != nil && !IsNil(o.Distribution) {
		return true
	}

	return false
}

// SetDistribution gets a reference to the given string and assigns it to the Distribution field.
func (o *AptProxyRepositoriesAttributes) SetDistribution(v string) {
	o.Distribution = &v
}

// GetFlat returns the Flat field value
func (o *AptProxyRepositoriesAttributes) GetFlat() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Flat
}

// GetFlatOk returns a tuple with the Flat field value
// and a boolean to check if the value has been set.
func (o *AptProxyRepositoriesAttributes) GetFlatOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Flat, true
}

// SetFlat sets field value
func (o *AptProxyRepositoriesAttributes) SetFlat(v bool) {
	o.Flat = v
}

func (o AptProxyRepositoriesAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AptProxyRepositoriesAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Distribution) {
		toSerialize["distribution"] = o.Distribution
	}
	toSerialize["flat"] = o.Flat
	return toSerialize, nil
}

func (o *AptProxyRepositoriesAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"flat",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAptProxyRepositoriesAttributes := _AptProxyRepositoriesAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAptProxyRepositoriesAttributes)

	if err != nil {
		return err
	}

	*o = AptProxyRepositoriesAttributes(varAptProxyRepositoriesAttributes)

	return err
}

type NullableAptProxyRepositoriesAttributes struct {
	value *AptProxyRepositoriesAttributes
	isSet bool
}

func (v NullableAptProxyRepositoriesAttributes) Get() *AptProxyRepositoriesAttributes {
	return v.value
}

func (v *NullableAptProxyRepositoriesAttributes) Set(val *AptProxyRepositoriesAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAptProxyRepositoriesAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAptProxyRepositoriesAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAptProxyRepositoriesAttributes(val *AptProxyRepositoriesAttributes) *NullableAptProxyRepositoriesAttributes {
	return &NullableAptProxyRepositoriesAttributes{value: val, isSet: true}
}

func (v NullableAptProxyRepositoriesAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAptProxyRepositoriesAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


