/*
Nexus Repository Manager REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.75.1-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nexus

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the IqAuditXo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IqAuditXo{}

// IqAuditXo struct for IqAuditXo
type IqAuditXo struct {
	// is audit enabled
	Enabled *bool `json:"enabled,omitempty"`
	// repository name
	RepositoryName string `json:"repositoryName"`
	// is quarantine enabled
	EnabledQuarantine bool `json:"enabledQuarantine"`
}

type _IqAuditXo IqAuditXo

// NewIqAuditXo instantiates a new IqAuditXo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIqAuditXo(repositoryName string, enabledQuarantine bool) *IqAuditXo {
	this := IqAuditXo{}
	this.RepositoryName = repositoryName
	this.EnabledQuarantine = enabledQuarantine
	return &this
}

// NewIqAuditXoWithDefaults instantiates a new IqAuditXo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIqAuditXoWithDefaults() *IqAuditXo {
	this := IqAuditXo{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *IqAuditXo) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IqAuditXo) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *IqAuditXo) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *IqAuditXo) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetRepositoryName returns the RepositoryName field value
func (o *IqAuditXo) GetRepositoryName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RepositoryName
}

// GetRepositoryNameOk returns a tuple with the RepositoryName field value
// and a boolean to check if the value has been set.
func (o *IqAuditXo) GetRepositoryNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RepositoryName, true
}

// SetRepositoryName sets field value
func (o *IqAuditXo) SetRepositoryName(v string) {
	o.RepositoryName = v
}

// GetEnabledQuarantine returns the EnabledQuarantine field value
func (o *IqAuditXo) GetEnabledQuarantine() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EnabledQuarantine
}

// GetEnabledQuarantineOk returns a tuple with the EnabledQuarantine field value
// and a boolean to check if the value has been set.
func (o *IqAuditXo) GetEnabledQuarantineOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnabledQuarantine, true
}

// SetEnabledQuarantine sets field value
func (o *IqAuditXo) SetEnabledQuarantine(v bool) {
	o.EnabledQuarantine = v
}

func (o IqAuditXo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IqAuditXo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	toSerialize["repositoryName"] = o.RepositoryName
	toSerialize["enabledQuarantine"] = o.EnabledQuarantine
	return toSerialize, nil
}

func (o *IqAuditXo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"repositoryName",
		"enabledQuarantine",
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

	varIqAuditXo := _IqAuditXo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIqAuditXo)

	if err != nil {
		return err
	}

	*o = IqAuditXo(varIqAuditXo)

	return err
}

type NullableIqAuditXo struct {
	value *IqAuditXo
	isSet bool
}

func (v NullableIqAuditXo) Get() *IqAuditXo {
	return v.value
}

func (v *NullableIqAuditXo) Set(val *IqAuditXo) {
	v.value = val
	v.isSet = true
}

func (v NullableIqAuditXo) IsSet() bool {
	return v.isSet
}

func (v *NullableIqAuditXo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIqAuditXo(val *IqAuditXo) *NullableIqAuditXo {
	return &NullableIqAuditXo{value: val, isSet: true}
}

func (v NullableIqAuditXo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIqAuditXo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


