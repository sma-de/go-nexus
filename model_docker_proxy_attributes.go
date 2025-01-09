/*
Nexus Repository Manager REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.76.0-03
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nexus

import (
	"encoding/json"
)

// checks if the DockerProxyAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DockerProxyAttributes{}

// DockerProxyAttributes struct for DockerProxyAttributes
type DockerProxyAttributes struct {
	// Type of Docker Index
	IndexType *string `json:"indexType,omitempty"`
	// Url of Docker Index to use
	IndexUrl *string `json:"indexUrl,omitempty"`
	// Allow Nexus Repository Manager to download and cache foreign layers
	CacheForeignLayers *bool `json:"cacheForeignLayers,omitempty"`
	// Regular expressions used to identify URLs that are allowed for foreign layer requests
	ForeignLayerUrlWhitelist []string `json:"foreignLayerUrlWhitelist,omitempty"`
}

// NewDockerProxyAttributes instantiates a new DockerProxyAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDockerProxyAttributes() *DockerProxyAttributes {
	this := DockerProxyAttributes{}
	return &this
}

// NewDockerProxyAttributesWithDefaults instantiates a new DockerProxyAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDockerProxyAttributesWithDefaults() *DockerProxyAttributes {
	this := DockerProxyAttributes{}
	return &this
}

// GetIndexType returns the IndexType field value if set, zero value otherwise.
func (o *DockerProxyAttributes) GetIndexType() string {
	if o == nil || IsNil(o.IndexType) {
		var ret string
		return ret
	}
	return *o.IndexType
}

// GetIndexTypeOk returns a tuple with the IndexType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DockerProxyAttributes) GetIndexTypeOk() (*string, bool) {
	if o == nil || IsNil(o.IndexType) {
		return nil, false
	}
	return o.IndexType, true
}

// HasIndexType returns a boolean if a field has been set.
func (o *DockerProxyAttributes) HasIndexType() bool {
	if o != nil && !IsNil(o.IndexType) {
		return true
	}

	return false
}

// SetIndexType gets a reference to the given string and assigns it to the IndexType field.
func (o *DockerProxyAttributes) SetIndexType(v string) {
	o.IndexType = &v
}

// GetIndexUrl returns the IndexUrl field value if set, zero value otherwise.
func (o *DockerProxyAttributes) GetIndexUrl() string {
	if o == nil || IsNil(o.IndexUrl) {
		var ret string
		return ret
	}
	return *o.IndexUrl
}

// GetIndexUrlOk returns a tuple with the IndexUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DockerProxyAttributes) GetIndexUrlOk() (*string, bool) {
	if o == nil || IsNil(o.IndexUrl) {
		return nil, false
	}
	return o.IndexUrl, true
}

// HasIndexUrl returns a boolean if a field has been set.
func (o *DockerProxyAttributes) HasIndexUrl() bool {
	if o != nil && !IsNil(o.IndexUrl) {
		return true
	}

	return false
}

// SetIndexUrl gets a reference to the given string and assigns it to the IndexUrl field.
func (o *DockerProxyAttributes) SetIndexUrl(v string) {
	o.IndexUrl = &v
}

// GetCacheForeignLayers returns the CacheForeignLayers field value if set, zero value otherwise.
func (o *DockerProxyAttributes) GetCacheForeignLayers() bool {
	if o == nil || IsNil(o.CacheForeignLayers) {
		var ret bool
		return ret
	}
	return *o.CacheForeignLayers
}

// GetCacheForeignLayersOk returns a tuple with the CacheForeignLayers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DockerProxyAttributes) GetCacheForeignLayersOk() (*bool, bool) {
	if o == nil || IsNil(o.CacheForeignLayers) {
		return nil, false
	}
	return o.CacheForeignLayers, true
}

// HasCacheForeignLayers returns a boolean if a field has been set.
func (o *DockerProxyAttributes) HasCacheForeignLayers() bool {
	if o != nil && !IsNil(o.CacheForeignLayers) {
		return true
	}

	return false
}

// SetCacheForeignLayers gets a reference to the given bool and assigns it to the CacheForeignLayers field.
func (o *DockerProxyAttributes) SetCacheForeignLayers(v bool) {
	o.CacheForeignLayers = &v
}

// GetForeignLayerUrlWhitelist returns the ForeignLayerUrlWhitelist field value if set, zero value otherwise.
func (o *DockerProxyAttributes) GetForeignLayerUrlWhitelist() []string {
	if o == nil || IsNil(o.ForeignLayerUrlWhitelist) {
		var ret []string
		return ret
	}
	return o.ForeignLayerUrlWhitelist
}

// GetForeignLayerUrlWhitelistOk returns a tuple with the ForeignLayerUrlWhitelist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DockerProxyAttributes) GetForeignLayerUrlWhitelistOk() ([]string, bool) {
	if o == nil || IsNil(o.ForeignLayerUrlWhitelist) {
		return nil, false
	}
	return o.ForeignLayerUrlWhitelist, true
}

// HasForeignLayerUrlWhitelist returns a boolean if a field has been set.
func (o *DockerProxyAttributes) HasForeignLayerUrlWhitelist() bool {
	if o != nil && !IsNil(o.ForeignLayerUrlWhitelist) {
		return true
	}

	return false
}

// SetForeignLayerUrlWhitelist gets a reference to the given []string and assigns it to the ForeignLayerUrlWhitelist field.
func (o *DockerProxyAttributes) SetForeignLayerUrlWhitelist(v []string) {
	o.ForeignLayerUrlWhitelist = v
}

func (o DockerProxyAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DockerProxyAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IndexType) {
		toSerialize["indexType"] = o.IndexType
	}
	if !IsNil(o.IndexUrl) {
		toSerialize["indexUrl"] = o.IndexUrl
	}
	if !IsNil(o.CacheForeignLayers) {
		toSerialize["cacheForeignLayers"] = o.CacheForeignLayers
	}
	if !IsNil(o.ForeignLayerUrlWhitelist) {
		toSerialize["foreignLayerUrlWhitelist"] = o.ForeignLayerUrlWhitelist
	}
	return toSerialize, nil
}

type NullableDockerProxyAttributes struct {
	value *DockerProxyAttributes
	isSet bool
}

func (v NullableDockerProxyAttributes) Get() *DockerProxyAttributes {
	return v.value
}

func (v *NullableDockerProxyAttributes) Set(val *DockerProxyAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableDockerProxyAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableDockerProxyAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDockerProxyAttributes(val *DockerProxyAttributes) *NullableDockerProxyAttributes {
	return &NullableDockerProxyAttributes{value: val, isSet: true}
}

func (v NullableDockerProxyAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDockerProxyAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


