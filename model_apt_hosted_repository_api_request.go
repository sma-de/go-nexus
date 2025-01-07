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

// checks if the AptHostedRepositoryApiRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AptHostedRepositoryApiRequest{}

// AptHostedRepositoryApiRequest struct for AptHostedRepositoryApiRequest
type AptHostedRepositoryApiRequest struct {
	// A unique identifier for this repository
	Name string `json:"name" validate:"regexp=^[a-zA-Z0-9\\\\-]{1}[a-zA-Z0-9_\\\\-\\\\.]*$"`
	// Whether this repository accepts incoming requests
	Online bool `json:"online"`
	Storage HostedStorageAttributes `json:"storage"`
	Cleanup *CleanupPolicyAttributes `json:"cleanup,omitempty"`
	Component *ComponentAttributes `json:"component,omitempty"`
	Apt AptHostedRepositoriesAttributes `json:"apt"`
	AptSigning AptSigningRepositoriesAttributes `json:"aptSigning"`
}

type _AptHostedRepositoryApiRequest AptHostedRepositoryApiRequest

// NewAptHostedRepositoryApiRequest instantiates a new AptHostedRepositoryApiRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAptHostedRepositoryApiRequest(name string, online bool, storage HostedStorageAttributes, apt AptHostedRepositoriesAttributes, aptSigning AptSigningRepositoriesAttributes) *AptHostedRepositoryApiRequest {
	this := AptHostedRepositoryApiRequest{}
	this.Name = name
	this.Online = online
	this.Storage = storage
	this.Apt = apt
	this.AptSigning = aptSigning
	return &this
}

// NewAptHostedRepositoryApiRequestWithDefaults instantiates a new AptHostedRepositoryApiRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAptHostedRepositoryApiRequestWithDefaults() *AptHostedRepositoryApiRequest {
	this := AptHostedRepositoryApiRequest{}
	return &this
}

// GetName returns the Name field value
func (o *AptHostedRepositoryApiRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AptHostedRepositoryApiRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AptHostedRepositoryApiRequest) SetName(v string) {
	o.Name = v
}

// GetOnline returns the Online field value
func (o *AptHostedRepositoryApiRequest) GetOnline() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Online
}

// GetOnlineOk returns a tuple with the Online field value
// and a boolean to check if the value has been set.
func (o *AptHostedRepositoryApiRequest) GetOnlineOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Online, true
}

// SetOnline sets field value
func (o *AptHostedRepositoryApiRequest) SetOnline(v bool) {
	o.Online = v
}

// GetStorage returns the Storage field value
func (o *AptHostedRepositoryApiRequest) GetStorage() HostedStorageAttributes {
	if o == nil {
		var ret HostedStorageAttributes
		return ret
	}

	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value
// and a boolean to check if the value has been set.
func (o *AptHostedRepositoryApiRequest) GetStorageOk() (*HostedStorageAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Storage, true
}

// SetStorage sets field value
func (o *AptHostedRepositoryApiRequest) SetStorage(v HostedStorageAttributes) {
	o.Storage = v
}

// GetCleanup returns the Cleanup field value if set, zero value otherwise.
func (o *AptHostedRepositoryApiRequest) GetCleanup() CleanupPolicyAttributes {
	if o == nil || IsNil(o.Cleanup) {
		var ret CleanupPolicyAttributes
		return ret
	}
	return *o.Cleanup
}

// GetCleanupOk returns a tuple with the Cleanup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AptHostedRepositoryApiRequest) GetCleanupOk() (*CleanupPolicyAttributes, bool) {
	if o == nil || IsNil(o.Cleanup) {
		return nil, false
	}
	return o.Cleanup, true
}

// HasCleanup returns a boolean if a field has been set.
func (o *AptHostedRepositoryApiRequest) HasCleanup() bool {
	if o != nil && !IsNil(o.Cleanup) {
		return true
	}

	return false
}

// SetCleanup gets a reference to the given CleanupPolicyAttributes and assigns it to the Cleanup field.
func (o *AptHostedRepositoryApiRequest) SetCleanup(v CleanupPolicyAttributes) {
	o.Cleanup = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *AptHostedRepositoryApiRequest) GetComponent() ComponentAttributes {
	if o == nil || IsNil(o.Component) {
		var ret ComponentAttributes
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AptHostedRepositoryApiRequest) GetComponentOk() (*ComponentAttributes, bool) {
	if o == nil || IsNil(o.Component) {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *AptHostedRepositoryApiRequest) HasComponent() bool {
	if o != nil && !IsNil(o.Component) {
		return true
	}

	return false
}

// SetComponent gets a reference to the given ComponentAttributes and assigns it to the Component field.
func (o *AptHostedRepositoryApiRequest) SetComponent(v ComponentAttributes) {
	o.Component = &v
}

// GetApt returns the Apt field value
func (o *AptHostedRepositoryApiRequest) GetApt() AptHostedRepositoriesAttributes {
	if o == nil {
		var ret AptHostedRepositoriesAttributes
		return ret
	}

	return o.Apt
}

// GetAptOk returns a tuple with the Apt field value
// and a boolean to check if the value has been set.
func (o *AptHostedRepositoryApiRequest) GetAptOk() (*AptHostedRepositoriesAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Apt, true
}

// SetApt sets field value
func (o *AptHostedRepositoryApiRequest) SetApt(v AptHostedRepositoriesAttributes) {
	o.Apt = v
}

// GetAptSigning returns the AptSigning field value
func (o *AptHostedRepositoryApiRequest) GetAptSigning() AptSigningRepositoriesAttributes {
	if o == nil {
		var ret AptSigningRepositoriesAttributes
		return ret
	}

	return o.AptSigning
}

// GetAptSigningOk returns a tuple with the AptSigning field value
// and a boolean to check if the value has been set.
func (o *AptHostedRepositoryApiRequest) GetAptSigningOk() (*AptSigningRepositoriesAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AptSigning, true
}

// SetAptSigning sets field value
func (o *AptHostedRepositoryApiRequest) SetAptSigning(v AptSigningRepositoriesAttributes) {
	o.AptSigning = v
}

func (o AptHostedRepositoryApiRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AptHostedRepositoryApiRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["online"] = o.Online
	toSerialize["storage"] = o.Storage
	if !IsNil(o.Cleanup) {
		toSerialize["cleanup"] = o.Cleanup
	}
	if !IsNil(o.Component) {
		toSerialize["component"] = o.Component
	}
	toSerialize["apt"] = o.Apt
	toSerialize["aptSigning"] = o.AptSigning
	return toSerialize, nil
}

func (o *AptHostedRepositoryApiRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"online",
		"storage",
		"apt",
		"aptSigning",
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

	varAptHostedRepositoryApiRequest := _AptHostedRepositoryApiRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAptHostedRepositoryApiRequest)

	if err != nil {
		return err
	}

	*o = AptHostedRepositoryApiRequest(varAptHostedRepositoryApiRequest)

	return err
}

type NullableAptHostedRepositoryApiRequest struct {
	value *AptHostedRepositoryApiRequest
	isSet bool
}

func (v NullableAptHostedRepositoryApiRequest) Get() *AptHostedRepositoryApiRequest {
	return v.value
}

func (v *NullableAptHostedRepositoryApiRequest) Set(val *AptHostedRepositoryApiRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAptHostedRepositoryApiRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAptHostedRepositoryApiRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAptHostedRepositoryApiRequest(val *AptHostedRepositoryApiRequest) *NullableAptHostedRepositoryApiRequest {
	return &NullableAptHostedRepositoryApiRequest{value: val, isSet: true}
}

func (v NullableAptHostedRepositoryApiRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAptHostedRepositoryApiRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


