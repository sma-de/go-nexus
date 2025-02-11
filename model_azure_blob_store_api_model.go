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

// checks if the AzureBlobStoreApiModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureBlobStoreApiModel{}

// AzureBlobStoreApiModel struct for AzureBlobStoreApiModel
type AzureBlobStoreApiModel struct {
	// The name of the Azure blob store.
	Name string `json:"name"`
	SoftQuota *BlobStoreApiSoftQuota `json:"softQuota,omitempty"`
	BucketConfiguration AzureBlobStoreApiBucketConfiguration `json:"bucketConfiguration"`
}

type _AzureBlobStoreApiModel AzureBlobStoreApiModel

// NewAzureBlobStoreApiModel instantiates a new AzureBlobStoreApiModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureBlobStoreApiModel(name string, bucketConfiguration AzureBlobStoreApiBucketConfiguration) *AzureBlobStoreApiModel {
	this := AzureBlobStoreApiModel{}
	this.Name = name
	this.BucketConfiguration = bucketConfiguration
	return &this
}

// NewAzureBlobStoreApiModelWithDefaults instantiates a new AzureBlobStoreApiModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureBlobStoreApiModelWithDefaults() *AzureBlobStoreApiModel {
	this := AzureBlobStoreApiModel{}
	return &this
}

// GetName returns the Name field value
func (o *AzureBlobStoreApiModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AzureBlobStoreApiModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AzureBlobStoreApiModel) SetName(v string) {
	o.Name = v
}

// GetSoftQuota returns the SoftQuota field value if set, zero value otherwise.
func (o *AzureBlobStoreApiModel) GetSoftQuota() BlobStoreApiSoftQuota {
	if o == nil || IsNil(o.SoftQuota) {
		var ret BlobStoreApiSoftQuota
		return ret
	}
	return *o.SoftQuota
}

// GetSoftQuotaOk returns a tuple with the SoftQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureBlobStoreApiModel) GetSoftQuotaOk() (*BlobStoreApiSoftQuota, bool) {
	if o == nil || IsNil(o.SoftQuota) {
		return nil, false
	}
	return o.SoftQuota, true
}

// HasSoftQuota returns a boolean if a field has been set.
func (o *AzureBlobStoreApiModel) HasSoftQuota() bool {
	if o != nil && !IsNil(o.SoftQuota) {
		return true
	}

	return false
}

// SetSoftQuota gets a reference to the given BlobStoreApiSoftQuota and assigns it to the SoftQuota field.
func (o *AzureBlobStoreApiModel) SetSoftQuota(v BlobStoreApiSoftQuota) {
	o.SoftQuota = &v
}

// GetBucketConfiguration returns the BucketConfiguration field value
func (o *AzureBlobStoreApiModel) GetBucketConfiguration() AzureBlobStoreApiBucketConfiguration {
	if o == nil {
		var ret AzureBlobStoreApiBucketConfiguration
		return ret
	}

	return o.BucketConfiguration
}

// GetBucketConfigurationOk returns a tuple with the BucketConfiguration field value
// and a boolean to check if the value has been set.
func (o *AzureBlobStoreApiModel) GetBucketConfigurationOk() (*AzureBlobStoreApiBucketConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BucketConfiguration, true
}

// SetBucketConfiguration sets field value
func (o *AzureBlobStoreApiModel) SetBucketConfiguration(v AzureBlobStoreApiBucketConfiguration) {
	o.BucketConfiguration = v
}

func (o AzureBlobStoreApiModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureBlobStoreApiModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.SoftQuota) {
		toSerialize["softQuota"] = o.SoftQuota
	}
	toSerialize["bucketConfiguration"] = o.BucketConfiguration
	return toSerialize, nil
}

func (o *AzureBlobStoreApiModel) UnmarshalJSON(data []byte) (err error) {
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

	varAzureBlobStoreApiModel := _AzureBlobStoreApiModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAzureBlobStoreApiModel)

	if err != nil {
		return err
	}

	*o = AzureBlobStoreApiModel(varAzureBlobStoreApiModel)

	return err
}

type NullableAzureBlobStoreApiModel struct {
	value *AzureBlobStoreApiModel
	isSet bool
}

func (v NullableAzureBlobStoreApiModel) Get() *AzureBlobStoreApiModel {
	return v.value
}

func (v *NullableAzureBlobStoreApiModel) Set(val *AzureBlobStoreApiModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureBlobStoreApiModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureBlobStoreApiModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureBlobStoreApiModel(val *AzureBlobStoreApiModel) *NullableAzureBlobStoreApiModel {
	return &NullableAzureBlobStoreApiModel{value: val, isSet: true}
}

func (v NullableAzureBlobStoreApiModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureBlobStoreApiModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


