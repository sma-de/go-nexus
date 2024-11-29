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

// checks if the GroupBlobStoreApiResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupBlobStoreApiResponse{}

// GroupBlobStoreApiResponse struct for GroupBlobStoreApiResponse
type GroupBlobStoreApiResponse struct {
	SoftQuota *BlobStoreApiSoftQuota `json:"softQuota,omitempty"`
	// List of the names of blob stores that are members of this group
	Members []string `json:"members,omitempty"`
	FillPolicy *string `json:"fillPolicy,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewGroupBlobStoreApiResponse instantiates a new GroupBlobStoreApiResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupBlobStoreApiResponse() *GroupBlobStoreApiResponse {
	this := GroupBlobStoreApiResponse{}
	return &this
}

// NewGroupBlobStoreApiResponseWithDefaults instantiates a new GroupBlobStoreApiResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupBlobStoreApiResponseWithDefaults() *GroupBlobStoreApiResponse {
	this := GroupBlobStoreApiResponse{}
	return &this
}

// GetSoftQuota returns the SoftQuota field value if set, zero value otherwise.
func (o *GroupBlobStoreApiResponse) GetSoftQuota() BlobStoreApiSoftQuota {
	if o == nil || IsNil(o.SoftQuota) {
		var ret BlobStoreApiSoftQuota
		return ret
	}
	return *o.SoftQuota
}

// GetSoftQuotaOk returns a tuple with the SoftQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupBlobStoreApiResponse) GetSoftQuotaOk() (*BlobStoreApiSoftQuota, bool) {
	if o == nil || IsNil(o.SoftQuota) {
		return nil, false
	}
	return o.SoftQuota, true
}

// HasSoftQuota returns a boolean if a field has been set.
func (o *GroupBlobStoreApiResponse) HasSoftQuota() bool {
	if o != nil && !IsNil(o.SoftQuota) {
		return true
	}

	return false
}

// SetSoftQuota gets a reference to the given BlobStoreApiSoftQuota and assigns it to the SoftQuota field.
func (o *GroupBlobStoreApiResponse) SetSoftQuota(v BlobStoreApiSoftQuota) {
	o.SoftQuota = &v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *GroupBlobStoreApiResponse) GetMembers() []string {
	if o == nil || IsNil(o.Members) {
		var ret []string
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupBlobStoreApiResponse) GetMembersOk() ([]string, bool) {
	if o == nil || IsNil(o.Members) {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *GroupBlobStoreApiResponse) HasMembers() bool {
	if o != nil && !IsNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []string and assigns it to the Members field.
func (o *GroupBlobStoreApiResponse) SetMembers(v []string) {
	o.Members = v
}

// GetFillPolicy returns the FillPolicy field value if set, zero value otherwise.
func (o *GroupBlobStoreApiResponse) GetFillPolicy() string {
	if o == nil || IsNil(o.FillPolicy) {
		var ret string
		return ret
	}
	return *o.FillPolicy
}

// GetFillPolicyOk returns a tuple with the FillPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupBlobStoreApiResponse) GetFillPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.FillPolicy) {
		return nil, false
	}
	return o.FillPolicy, true
}

// HasFillPolicy returns a boolean if a field has been set.
func (o *GroupBlobStoreApiResponse) HasFillPolicy() bool {
	if o != nil && !IsNil(o.FillPolicy) {
		return true
	}

	return false
}

// SetFillPolicy gets a reference to the given string and assigns it to the FillPolicy field.
func (o *GroupBlobStoreApiResponse) SetFillPolicy(v string) {
	o.FillPolicy = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GroupBlobStoreApiResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupBlobStoreApiResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GroupBlobStoreApiResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GroupBlobStoreApiResponse) SetName(v string) {
	o.Name = &v
}

func (o GroupBlobStoreApiResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupBlobStoreApiResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SoftQuota) {
		toSerialize["softQuota"] = o.SoftQuota
	}
	if !IsNil(o.Members) {
		toSerialize["members"] = o.Members
	}
	if !IsNil(o.FillPolicy) {
		toSerialize["fillPolicy"] = o.FillPolicy
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableGroupBlobStoreApiResponse struct {
	value *GroupBlobStoreApiResponse
	isSet bool
}

func (v NullableGroupBlobStoreApiResponse) Get() *GroupBlobStoreApiResponse {
	return v.value
}

func (v *NullableGroupBlobStoreApiResponse) Set(val *GroupBlobStoreApiResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupBlobStoreApiResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupBlobStoreApiResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupBlobStoreApiResponse(val *GroupBlobStoreApiResponse) *NullableGroupBlobStoreApiResponse {
	return &NullableGroupBlobStoreApiResponse{value: val, isSet: true}
}

func (v NullableGroupBlobStoreApiResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupBlobStoreApiResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


