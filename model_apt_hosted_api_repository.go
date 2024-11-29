/*
Nexus Repository Manager REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.70.3-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nexus

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AptHostedApiRepository type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AptHostedApiRepository{}

// AptHostedApiRepository struct for AptHostedApiRepository
type AptHostedApiRepository struct {
	// A unique identifier for this repository
	Name *string `json:"name,omitempty" validate:"regexp=^[a-zA-Z0-9\\\\-]{1}[a-zA-Z0-9_\\\\-\\\\.]*$"`
	// Whether this repository accepts incoming requests
	Online bool `json:"online"`
	Storage HostedStorageAttributes `json:"storage"`
	Cleanup *CleanupPolicyAttributes `json:"cleanup,omitempty"`
	Component *ComponentAttributes `json:"component,omitempty"`
	Apt AptHostedRepositoriesAttributes `json:"apt"`
	AptSigning AptSigningRepositoriesAttributes `json:"aptSigning"`
}

type _AptHostedApiRepository AptHostedApiRepository

// NewAptHostedApiRepository instantiates a new AptHostedApiRepository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAptHostedApiRepository(online bool, storage HostedStorageAttributes, apt AptHostedRepositoriesAttributes, aptSigning AptSigningRepositoriesAttributes) *AptHostedApiRepository {
	this := AptHostedApiRepository{}
	this.Online = online
	this.Storage = storage
	this.Apt = apt
	this.AptSigning = aptSigning
	return &this
}

// NewAptHostedApiRepositoryWithDefaults instantiates a new AptHostedApiRepository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAptHostedApiRepositoryWithDefaults() *AptHostedApiRepository {
	this := AptHostedApiRepository{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AptHostedApiRepository) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AptHostedApiRepository) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AptHostedApiRepository) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AptHostedApiRepository) SetName(v string) {
	o.Name = &v
}

// GetOnline returns the Online field value
func (o *AptHostedApiRepository) GetOnline() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Online
}

// GetOnlineOk returns a tuple with the Online field value
// and a boolean to check if the value has been set.
func (o *AptHostedApiRepository) GetOnlineOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Online, true
}

// SetOnline sets field value
func (o *AptHostedApiRepository) SetOnline(v bool) {
	o.Online = v
}

// GetStorage returns the Storage field value
func (o *AptHostedApiRepository) GetStorage() HostedStorageAttributes {
	if o == nil {
		var ret HostedStorageAttributes
		return ret
	}

	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value
// and a boolean to check if the value has been set.
func (o *AptHostedApiRepository) GetStorageOk() (*HostedStorageAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Storage, true
}

// SetStorage sets field value
func (o *AptHostedApiRepository) SetStorage(v HostedStorageAttributes) {
	o.Storage = v
}

// GetCleanup returns the Cleanup field value if set, zero value otherwise.
func (o *AptHostedApiRepository) GetCleanup() CleanupPolicyAttributes {
	if o == nil || IsNil(o.Cleanup) {
		var ret CleanupPolicyAttributes
		return ret
	}
	return *o.Cleanup
}

// GetCleanupOk returns a tuple with the Cleanup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AptHostedApiRepository) GetCleanupOk() (*CleanupPolicyAttributes, bool) {
	if o == nil || IsNil(o.Cleanup) {
		return nil, false
	}
	return o.Cleanup, true
}

// HasCleanup returns a boolean if a field has been set.
func (o *AptHostedApiRepository) HasCleanup() bool {
	if o != nil && !IsNil(o.Cleanup) {
		return true
	}

	return false
}

// SetCleanup gets a reference to the given CleanupPolicyAttributes and assigns it to the Cleanup field.
func (o *AptHostedApiRepository) SetCleanup(v CleanupPolicyAttributes) {
	o.Cleanup = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *AptHostedApiRepository) GetComponent() ComponentAttributes {
	if o == nil || IsNil(o.Component) {
		var ret ComponentAttributes
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AptHostedApiRepository) GetComponentOk() (*ComponentAttributes, bool) {
	if o == nil || IsNil(o.Component) {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *AptHostedApiRepository) HasComponent() bool {
	if o != nil && !IsNil(o.Component) {
		return true
	}

	return false
}

// SetComponent gets a reference to the given ComponentAttributes and assigns it to the Component field.
func (o *AptHostedApiRepository) SetComponent(v ComponentAttributes) {
	o.Component = &v
}

// GetApt returns the Apt field value
func (o *AptHostedApiRepository) GetApt() AptHostedRepositoriesAttributes {
	if o == nil {
		var ret AptHostedRepositoriesAttributes
		return ret
	}

	return o.Apt
}

// GetAptOk returns a tuple with the Apt field value
// and a boolean to check if the value has been set.
func (o *AptHostedApiRepository) GetAptOk() (*AptHostedRepositoriesAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Apt, true
}

// SetApt sets field value
func (o *AptHostedApiRepository) SetApt(v AptHostedRepositoriesAttributes) {
	o.Apt = v
}

// GetAptSigning returns the AptSigning field value
func (o *AptHostedApiRepository) GetAptSigning() AptSigningRepositoriesAttributes {
	if o == nil {
		var ret AptSigningRepositoriesAttributes
		return ret
	}

	return o.AptSigning
}

// GetAptSigningOk returns a tuple with the AptSigning field value
// and a boolean to check if the value has been set.
func (o *AptHostedApiRepository) GetAptSigningOk() (*AptSigningRepositoriesAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AptSigning, true
}

// SetAptSigning sets field value
func (o *AptHostedApiRepository) SetAptSigning(v AptSigningRepositoriesAttributes) {
	o.AptSigning = v
}

func (o AptHostedApiRepository) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AptHostedApiRepository) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
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

func (o *AptHostedApiRepository) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varAptHostedApiRepository := _AptHostedApiRepository{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAptHostedApiRepository)

	if err != nil {
		return err
	}

	*o = AptHostedApiRepository(varAptHostedApiRepository)

	return err
}

type NullableAptHostedApiRepository struct {
	value *AptHostedApiRepository
	isSet bool
}

func (v NullableAptHostedApiRepository) Get() *AptHostedApiRepository {
	return v.value
}

func (v *NullableAptHostedApiRepository) Set(val *AptHostedApiRepository) {
	v.value = val
	v.isSet = true
}

func (v NullableAptHostedApiRepository) IsSet() bool {
	return v.isSet
}

func (v *NullableAptHostedApiRepository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAptHostedApiRepository(val *AptHostedApiRepository) *NullableAptHostedApiRepository {
	return &NullableAptHostedApiRepository{value: val, isSet: true}
}

func (v NullableAptHostedApiRepository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAptHostedApiRepository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


