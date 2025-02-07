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

// checks if the NugetGroupRepositoryApiRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NugetGroupRepositoryApiRequest{}

// NugetGroupRepositoryApiRequest struct for NugetGroupRepositoryApiRequest
type NugetGroupRepositoryApiRequest struct {
	// A unique identifier for this repository
	Name string `json:"name" validate:"regexp=^[a-zA-Z0-9\\\\-]{1}[a-zA-Z0-9_\\\\-\\\\.]*$"`
	// Whether this repository accepts incoming requests
	Online bool `json:"online"`
	Storage StorageAttributes `json:"storage"`
	Group GroupAttributes `json:"group"`
}

type _NugetGroupRepositoryApiRequest NugetGroupRepositoryApiRequest

// NewNugetGroupRepositoryApiRequest instantiates a new NugetGroupRepositoryApiRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNugetGroupRepositoryApiRequest(name string, online bool, storage StorageAttributes, group GroupAttributes) *NugetGroupRepositoryApiRequest {
	this := NugetGroupRepositoryApiRequest{}
	this.Name = name
	this.Online = online
	this.Storage = storage
	this.Group = group
	return &this
}

// NewNugetGroupRepositoryApiRequestWithDefaults instantiates a new NugetGroupRepositoryApiRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNugetGroupRepositoryApiRequestWithDefaults() *NugetGroupRepositoryApiRequest {
	this := NugetGroupRepositoryApiRequest{}
	return &this
}

// GetName returns the Name field value
func (o *NugetGroupRepositoryApiRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NugetGroupRepositoryApiRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NugetGroupRepositoryApiRequest) SetName(v string) {
	o.Name = v
}

// GetOnline returns the Online field value
func (o *NugetGroupRepositoryApiRequest) GetOnline() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Online
}

// GetOnlineOk returns a tuple with the Online field value
// and a boolean to check if the value has been set.
func (o *NugetGroupRepositoryApiRequest) GetOnlineOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Online, true
}

// SetOnline sets field value
func (o *NugetGroupRepositoryApiRequest) SetOnline(v bool) {
	o.Online = v
}

// GetStorage returns the Storage field value
func (o *NugetGroupRepositoryApiRequest) GetStorage() StorageAttributes {
	if o == nil {
		var ret StorageAttributes
		return ret
	}

	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value
// and a boolean to check if the value has been set.
func (o *NugetGroupRepositoryApiRequest) GetStorageOk() (*StorageAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Storage, true
}

// SetStorage sets field value
func (o *NugetGroupRepositoryApiRequest) SetStorage(v StorageAttributes) {
	o.Storage = v
}

// GetGroup returns the Group field value
func (o *NugetGroupRepositoryApiRequest) GetGroup() GroupAttributes {
	if o == nil {
		var ret GroupAttributes
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *NugetGroupRepositoryApiRequest) GetGroupOk() (*GroupAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *NugetGroupRepositoryApiRequest) SetGroup(v GroupAttributes) {
	o.Group = v
}

func (o NugetGroupRepositoryApiRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NugetGroupRepositoryApiRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["online"] = o.Online
	toSerialize["storage"] = o.Storage
	toSerialize["group"] = o.Group
	return toSerialize, nil
}

func (o *NugetGroupRepositoryApiRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"online",
		"storage",
		"group",
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

	varNugetGroupRepositoryApiRequest := _NugetGroupRepositoryApiRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNugetGroupRepositoryApiRequest)

	if err != nil {
		return err
	}

	*o = NugetGroupRepositoryApiRequest(varNugetGroupRepositoryApiRequest)

	return err
}

type NullableNugetGroupRepositoryApiRequest struct {
	value *NugetGroupRepositoryApiRequest
	isSet bool
}

func (v NullableNugetGroupRepositoryApiRequest) Get() *NugetGroupRepositoryApiRequest {
	return v.value
}

func (v *NullableNugetGroupRepositoryApiRequest) Set(val *NugetGroupRepositoryApiRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNugetGroupRepositoryApiRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNugetGroupRepositoryApiRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNugetGroupRepositoryApiRequest(val *NugetGroupRepositoryApiRequest) *NullableNugetGroupRepositoryApiRequest {
	return &NullableNugetGroupRepositoryApiRequest{value: val, isSet: true}
}

func (v NullableNugetGroupRepositoryApiRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNugetGroupRepositoryApiRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


