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

// checks if the GoogleCloudBlobstoreApiModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GoogleCloudBlobstoreApiModel{}

// GoogleCloudBlobstoreApiModel struct for GoogleCloudBlobstoreApiModel
type GoogleCloudBlobstoreApiModel struct {
	// The name of the GC Storage blob store.
	Name string `json:"name"`
	SoftQuota *BlobStoreApiSoftQuota `json:"softQuota,omitempty"`
	BucketConfiguration GoogleCloudBlobStoreApiBucketConfiguration `json:"bucketConfiguration"`
	// The blob store type.
	Type *string `json:"type,omitempty"`
}

type _GoogleCloudBlobstoreApiModel GoogleCloudBlobstoreApiModel

// NewGoogleCloudBlobstoreApiModel instantiates a new GoogleCloudBlobstoreApiModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleCloudBlobstoreApiModel(name string, bucketConfiguration GoogleCloudBlobStoreApiBucketConfiguration) *GoogleCloudBlobstoreApiModel {
	this := GoogleCloudBlobstoreApiModel{}
	this.Name = name
	this.BucketConfiguration = bucketConfiguration
	return &this
}

// NewGoogleCloudBlobstoreApiModelWithDefaults instantiates a new GoogleCloudBlobstoreApiModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleCloudBlobstoreApiModelWithDefaults() *GoogleCloudBlobstoreApiModel {
	this := GoogleCloudBlobstoreApiModel{}
	return &this
}

// GetName returns the Name field value
func (o *GoogleCloudBlobstoreApiModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GoogleCloudBlobstoreApiModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GoogleCloudBlobstoreApiModel) SetName(v string) {
	o.Name = v
}

// GetSoftQuota returns the SoftQuota field value if set, zero value otherwise.
func (o *GoogleCloudBlobstoreApiModel) GetSoftQuota() BlobStoreApiSoftQuota {
	if o == nil || IsNil(o.SoftQuota) {
		var ret BlobStoreApiSoftQuota
		return ret
	}
	return *o.SoftQuota
}

// GetSoftQuotaOk returns a tuple with the SoftQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleCloudBlobstoreApiModel) GetSoftQuotaOk() (*BlobStoreApiSoftQuota, bool) {
	if o == nil || IsNil(o.SoftQuota) {
		return nil, false
	}
	return o.SoftQuota, true
}

// HasSoftQuota returns a boolean if a field has been set.
func (o *GoogleCloudBlobstoreApiModel) HasSoftQuota() bool {
	if o != nil && !IsNil(o.SoftQuota) {
		return true
	}

	return false
}

// SetSoftQuota gets a reference to the given BlobStoreApiSoftQuota and assigns it to the SoftQuota field.
func (o *GoogleCloudBlobstoreApiModel) SetSoftQuota(v BlobStoreApiSoftQuota) {
	o.SoftQuota = &v
}

// GetBucketConfiguration returns the BucketConfiguration field value
func (o *GoogleCloudBlobstoreApiModel) GetBucketConfiguration() GoogleCloudBlobStoreApiBucketConfiguration {
	if o == nil {
		var ret GoogleCloudBlobStoreApiBucketConfiguration
		return ret
	}

	return o.BucketConfiguration
}

// GetBucketConfigurationOk returns a tuple with the BucketConfiguration field value
// and a boolean to check if the value has been set.
func (o *GoogleCloudBlobstoreApiModel) GetBucketConfigurationOk() (*GoogleCloudBlobStoreApiBucketConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BucketConfiguration, true
}

// SetBucketConfiguration sets field value
func (o *GoogleCloudBlobstoreApiModel) SetBucketConfiguration(v GoogleCloudBlobStoreApiBucketConfiguration) {
	o.BucketConfiguration = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GoogleCloudBlobstoreApiModel) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleCloudBlobstoreApiModel) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GoogleCloudBlobstoreApiModel) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GoogleCloudBlobstoreApiModel) SetType(v string) {
	o.Type = &v
}

func (o GoogleCloudBlobstoreApiModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GoogleCloudBlobstoreApiModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.SoftQuota) {
		toSerialize["softQuota"] = o.SoftQuota
	}
	toSerialize["bucketConfiguration"] = o.BucketConfiguration
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

func (o *GoogleCloudBlobstoreApiModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"bucketConfiguration",
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

	varGoogleCloudBlobstoreApiModel := _GoogleCloudBlobstoreApiModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGoogleCloudBlobstoreApiModel)

	if err != nil {
		return err
	}

	*o = GoogleCloudBlobstoreApiModel(varGoogleCloudBlobstoreApiModel)

	return err
}

type NullableGoogleCloudBlobstoreApiModel struct {
	value *GoogleCloudBlobstoreApiModel
	isSet bool
}

func (v NullableGoogleCloudBlobstoreApiModel) Get() *GoogleCloudBlobstoreApiModel {
	return v.value
}

func (v *NullableGoogleCloudBlobstoreApiModel) Set(val *GoogleCloudBlobstoreApiModel) {
	v.value = val
	v.isSet = true
}

func (v NullableGoogleCloudBlobstoreApiModel) IsSet() bool {
	return v.isSet
}

func (v *NullableGoogleCloudBlobstoreApiModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoogleCloudBlobstoreApiModel(val *GoogleCloudBlobstoreApiModel) *NullableGoogleCloudBlobstoreApiModel {
	return &NullableGoogleCloudBlobstoreApiModel{value: val, isSet: true}
}

func (v NullableGoogleCloudBlobstoreApiModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoogleCloudBlobstoreApiModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


