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

// checks if the DockerGroupApiRepository type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DockerGroupApiRepository{}

// DockerGroupApiRepository struct for DockerGroupApiRepository
type DockerGroupApiRepository struct {
	// A unique identifier for this repository
	Name *string `json:"name,omitempty" validate:"regexp=^[a-zA-Z0-9\\\\-]{1}[a-zA-Z0-9_\\\\-\\\\.]*$"`
	// Whether this repository accepts incoming requests
	Online bool `json:"online"`
	Storage StorageAttributes `json:"storage"`
	Group GroupDeployAttributes `json:"group"`
	Docker DockerAttributes `json:"docker"`
}

type _DockerGroupApiRepository DockerGroupApiRepository

// NewDockerGroupApiRepository instantiates a new DockerGroupApiRepository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDockerGroupApiRepository(online bool, storage StorageAttributes, group GroupDeployAttributes, docker DockerAttributes) *DockerGroupApiRepository {
	this := DockerGroupApiRepository{}
	this.Online = online
	this.Storage = storage
	this.Group = group
	this.Docker = docker
	return &this
}

// NewDockerGroupApiRepositoryWithDefaults instantiates a new DockerGroupApiRepository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDockerGroupApiRepositoryWithDefaults() *DockerGroupApiRepository {
	this := DockerGroupApiRepository{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DockerGroupApiRepository) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DockerGroupApiRepository) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DockerGroupApiRepository) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DockerGroupApiRepository) SetName(v string) {
	o.Name = &v
}

// GetOnline returns the Online field value
func (o *DockerGroupApiRepository) GetOnline() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Online
}

// GetOnlineOk returns a tuple with the Online field value
// and a boolean to check if the value has been set.
func (o *DockerGroupApiRepository) GetOnlineOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Online, true
}

// SetOnline sets field value
func (o *DockerGroupApiRepository) SetOnline(v bool) {
	o.Online = v
}

// GetStorage returns the Storage field value
func (o *DockerGroupApiRepository) GetStorage() StorageAttributes {
	if o == nil {
		var ret StorageAttributes
		return ret
	}

	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value
// and a boolean to check if the value has been set.
func (o *DockerGroupApiRepository) GetStorageOk() (*StorageAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Storage, true
}

// SetStorage sets field value
func (o *DockerGroupApiRepository) SetStorage(v StorageAttributes) {
	o.Storage = v
}

// GetGroup returns the Group field value
func (o *DockerGroupApiRepository) GetGroup() GroupDeployAttributes {
	if o == nil {
		var ret GroupDeployAttributes
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *DockerGroupApiRepository) GetGroupOk() (*GroupDeployAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *DockerGroupApiRepository) SetGroup(v GroupDeployAttributes) {
	o.Group = v
}

// GetDocker returns the Docker field value
func (o *DockerGroupApiRepository) GetDocker() DockerAttributes {
	if o == nil {
		var ret DockerAttributes
		return ret
	}

	return o.Docker
}

// GetDockerOk returns a tuple with the Docker field value
// and a boolean to check if the value has been set.
func (o *DockerGroupApiRepository) GetDockerOk() (*DockerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Docker, true
}

// SetDocker sets field value
func (o *DockerGroupApiRepository) SetDocker(v DockerAttributes) {
	o.Docker = v
}

func (o DockerGroupApiRepository) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DockerGroupApiRepository) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["online"] = o.Online
	toSerialize["storage"] = o.Storage
	toSerialize["group"] = o.Group
	toSerialize["docker"] = o.Docker
	return toSerialize, nil
}

func (o *DockerGroupApiRepository) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"online",
		"storage",
		"group",
		"docker",
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

	varDockerGroupApiRepository := _DockerGroupApiRepository{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDockerGroupApiRepository)

	if err != nil {
		return err
	}

	*o = DockerGroupApiRepository(varDockerGroupApiRepository)

	return err
}

type NullableDockerGroupApiRepository struct {
	value *DockerGroupApiRepository
	isSet bool
}

func (v NullableDockerGroupApiRepository) Get() *DockerGroupApiRepository {
	return v.value
}

func (v *NullableDockerGroupApiRepository) Set(val *DockerGroupApiRepository) {
	v.value = val
	v.isSet = true
}

func (v NullableDockerGroupApiRepository) IsSet() bool {
	return v.isSet
}

func (v *NullableDockerGroupApiRepository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDockerGroupApiRepository(val *DockerGroupApiRepository) *NullableDockerGroupApiRepository {
	return &NullableDockerGroupApiRepository{value: val, isSet: true}
}

func (v NullableDockerGroupApiRepository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDockerGroupApiRepository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


