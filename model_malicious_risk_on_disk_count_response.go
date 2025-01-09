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

// checks if the MaliciousRiskOnDiskCountResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MaliciousRiskOnDiskCountResponse{}

// MaliciousRiskOnDiskCountResponse struct for MaliciousRiskOnDiskCountResponse
type MaliciousRiskOnDiskCountResponse struct {
	TotalCount *int64 `json:"totalCount,omitempty"`
	HdsExceptionThrown *bool `json:"hdsExceptionThrown,omitempty"`
}

// NewMaliciousRiskOnDiskCountResponse instantiates a new MaliciousRiskOnDiskCountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaliciousRiskOnDiskCountResponse() *MaliciousRiskOnDiskCountResponse {
	this := MaliciousRiskOnDiskCountResponse{}
	return &this
}

// NewMaliciousRiskOnDiskCountResponseWithDefaults instantiates a new MaliciousRiskOnDiskCountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaliciousRiskOnDiskCountResponseWithDefaults() *MaliciousRiskOnDiskCountResponse {
	this := MaliciousRiskOnDiskCountResponse{}
	return &this
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *MaliciousRiskOnDiskCountResponse) GetTotalCount() int64 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int64
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaliciousRiskOnDiskCountResponse) GetTotalCountOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *MaliciousRiskOnDiskCountResponse) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *MaliciousRiskOnDiskCountResponse) SetTotalCount(v int64) {
	o.TotalCount = &v
}

// GetHdsExceptionThrown returns the HdsExceptionThrown field value if set, zero value otherwise.
func (o *MaliciousRiskOnDiskCountResponse) GetHdsExceptionThrown() bool {
	if o == nil || IsNil(o.HdsExceptionThrown) {
		var ret bool
		return ret
	}
	return *o.HdsExceptionThrown
}

// GetHdsExceptionThrownOk returns a tuple with the HdsExceptionThrown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaliciousRiskOnDiskCountResponse) GetHdsExceptionThrownOk() (*bool, bool) {
	if o == nil || IsNil(o.HdsExceptionThrown) {
		return nil, false
	}
	return o.HdsExceptionThrown, true
}

// HasHdsExceptionThrown returns a boolean if a field has been set.
func (o *MaliciousRiskOnDiskCountResponse) HasHdsExceptionThrown() bool {
	if o != nil && !IsNil(o.HdsExceptionThrown) {
		return true
	}

	return false
}

// SetHdsExceptionThrown gets a reference to the given bool and assigns it to the HdsExceptionThrown field.
func (o *MaliciousRiskOnDiskCountResponse) SetHdsExceptionThrown(v bool) {
	o.HdsExceptionThrown = &v
}

func (o MaliciousRiskOnDiskCountResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MaliciousRiskOnDiskCountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	if !IsNil(o.HdsExceptionThrown) {
		toSerialize["hdsExceptionThrown"] = o.HdsExceptionThrown
	}
	return toSerialize, nil
}

type NullableMaliciousRiskOnDiskCountResponse struct {
	value *MaliciousRiskOnDiskCountResponse
	isSet bool
}

func (v NullableMaliciousRiskOnDiskCountResponse) Get() *MaliciousRiskOnDiskCountResponse {
	return v.value
}

func (v *NullableMaliciousRiskOnDiskCountResponse) Set(val *MaliciousRiskOnDiskCountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMaliciousRiskOnDiskCountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMaliciousRiskOnDiskCountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaliciousRiskOnDiskCountResponse(val *MaliciousRiskOnDiskCountResponse) *NullableMaliciousRiskOnDiskCountResponse {
	return &NullableMaliciousRiskOnDiskCountResponse{value: val, isSet: true}
}

func (v NullableMaliciousRiskOnDiskCountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaliciousRiskOnDiskCountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


