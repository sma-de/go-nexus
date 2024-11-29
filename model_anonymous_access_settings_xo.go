/*
Nexus Repository Manager REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.70.3-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nexus

import (
	"encoding/json"
)

// checks if the AnonymousAccessSettingsXO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnonymousAccessSettingsXO{}

// AnonymousAccessSettingsXO struct for AnonymousAccessSettingsXO
type AnonymousAccessSettingsXO struct {
	// Whether or not Anonymous Access is enabled
	Enabled *bool `json:"enabled,omitempty"`
	// The username of the anonymous account
	UserId *string `json:"userId,omitempty"`
	// The name of the authentication realm for the anonymous account
	RealmName *string `json:"realmName,omitempty"`
}

// NewAnonymousAccessSettingsXO instantiates a new AnonymousAccessSettingsXO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnonymousAccessSettingsXO() *AnonymousAccessSettingsXO {
	this := AnonymousAccessSettingsXO{}
	return &this
}

// NewAnonymousAccessSettingsXOWithDefaults instantiates a new AnonymousAccessSettingsXO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnonymousAccessSettingsXOWithDefaults() *AnonymousAccessSettingsXO {
	this := AnonymousAccessSettingsXO{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AnonymousAccessSettingsXO) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnonymousAccessSettingsXO) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AnonymousAccessSettingsXO) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AnonymousAccessSettingsXO) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *AnonymousAccessSettingsXO) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnonymousAccessSettingsXO) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *AnonymousAccessSettingsXO) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *AnonymousAccessSettingsXO) SetUserId(v string) {
	o.UserId = &v
}

// GetRealmName returns the RealmName field value if set, zero value otherwise.
func (o *AnonymousAccessSettingsXO) GetRealmName() string {
	if o == nil || IsNil(o.RealmName) {
		var ret string
		return ret
	}
	return *o.RealmName
}

// GetRealmNameOk returns a tuple with the RealmName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnonymousAccessSettingsXO) GetRealmNameOk() (*string, bool) {
	if o == nil || IsNil(o.RealmName) {
		return nil, false
	}
	return o.RealmName, true
}

// HasRealmName returns a boolean if a field has been set.
func (o *AnonymousAccessSettingsXO) HasRealmName() bool {
	if o != nil && !IsNil(o.RealmName) {
		return true
	}

	return false
}

// SetRealmName gets a reference to the given string and assigns it to the RealmName field.
func (o *AnonymousAccessSettingsXO) SetRealmName(v string) {
	o.RealmName = &v
}

func (o AnonymousAccessSettingsXO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnonymousAccessSettingsXO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !IsNil(o.RealmName) {
		toSerialize["realmName"] = o.RealmName
	}
	return toSerialize, nil
}

type NullableAnonymousAccessSettingsXO struct {
	value *AnonymousAccessSettingsXO
	isSet bool
}

func (v NullableAnonymousAccessSettingsXO) Get() *AnonymousAccessSettingsXO {
	return v.value
}

func (v *NullableAnonymousAccessSettingsXO) Set(val *AnonymousAccessSettingsXO) {
	v.value = val
	v.isSet = true
}

func (v NullableAnonymousAccessSettingsXO) IsSet() bool {
	return v.isSet
}

func (v *NullableAnonymousAccessSettingsXO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnonymousAccessSettingsXO(val *AnonymousAccessSettingsXO) *NullableAnonymousAccessSettingsXO {
	return &NullableAnonymousAccessSettingsXO{value: val, isSet: true}
}

func (v NullableAnonymousAccessSettingsXO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnonymousAccessSettingsXO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


