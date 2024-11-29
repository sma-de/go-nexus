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

// checks if the ApiPrivilegeRepositoryViewRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiPrivilegeRepositoryViewRequest{}

// ApiPrivilegeRepositoryViewRequest struct for ApiPrivilegeRepositoryViewRequest
type ApiPrivilegeRepositoryViewRequest struct {
	// The name of the privilege.  This value cannot be changed.
	Name *string `json:"name,omitempty" validate:"regexp=^[a-zA-Z0-9\\\\-]{1}[a-zA-Z0-9_\\\\-\\\\.]*$"`
	Description *string `json:"description,omitempty"`
	// A collection of actions to associate with the privilege, using BREAD syntax (browse,read,edit,add,delete,all) as well as 'run' for script privileges.
	Actions []string `json:"actions,omitempty"`
	// The repository format (i.e 'nuget', 'npm') this privilege will grant access to (or * for all).
	Format *string `json:"format,omitempty"`
	// The name of the repository this privilege will grant access to (or * for all).
	Repository *string `json:"repository,omitempty"`
}

// NewApiPrivilegeRepositoryViewRequest instantiates a new ApiPrivilegeRepositoryViewRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiPrivilegeRepositoryViewRequest() *ApiPrivilegeRepositoryViewRequest {
	this := ApiPrivilegeRepositoryViewRequest{}
	return &this
}

// NewApiPrivilegeRepositoryViewRequestWithDefaults instantiates a new ApiPrivilegeRepositoryViewRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiPrivilegeRepositoryViewRequestWithDefaults() *ApiPrivilegeRepositoryViewRequest {
	this := ApiPrivilegeRepositoryViewRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiPrivilegeRepositoryViewRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPrivilegeRepositoryViewRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiPrivilegeRepositoryViewRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiPrivilegeRepositoryViewRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApiPrivilegeRepositoryViewRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPrivilegeRepositoryViewRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApiPrivilegeRepositoryViewRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApiPrivilegeRepositoryViewRequest) SetDescription(v string) {
	o.Description = &v
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *ApiPrivilegeRepositoryViewRequest) GetActions() []string {
	if o == nil || IsNil(o.Actions) {
		var ret []string
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPrivilegeRepositoryViewRequest) GetActionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *ApiPrivilegeRepositoryViewRequest) HasActions() bool {
	if o != nil && !IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given []string and assigns it to the Actions field.
func (o *ApiPrivilegeRepositoryViewRequest) SetActions(v []string) {
	o.Actions = v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *ApiPrivilegeRepositoryViewRequest) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPrivilegeRepositoryViewRequest) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *ApiPrivilegeRepositoryViewRequest) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *ApiPrivilegeRepositoryViewRequest) SetFormat(v string) {
	o.Format = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *ApiPrivilegeRepositoryViewRequest) GetRepository() string {
	if o == nil || IsNil(o.Repository) {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPrivilegeRepositoryViewRequest) GetRepositoryOk() (*string, bool) {
	if o == nil || IsNil(o.Repository) {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *ApiPrivilegeRepositoryViewRequest) HasRepository() bool {
	if o != nil && !IsNil(o.Repository) {
		return true
	}

	return false
}

// SetRepository gets a reference to the given string and assigns it to the Repository field.
func (o *ApiPrivilegeRepositoryViewRequest) SetRepository(v string) {
	o.Repository = &v
}

func (o ApiPrivilegeRepositoryViewRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiPrivilegeRepositoryViewRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Actions) {
		toSerialize["actions"] = o.Actions
	}
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	if !IsNil(o.Repository) {
		toSerialize["repository"] = o.Repository
	}
	return toSerialize, nil
}

type NullableApiPrivilegeRepositoryViewRequest struct {
	value *ApiPrivilegeRepositoryViewRequest
	isSet bool
}

func (v NullableApiPrivilegeRepositoryViewRequest) Get() *ApiPrivilegeRepositoryViewRequest {
	return v.value
}

func (v *NullableApiPrivilegeRepositoryViewRequest) Set(val *ApiPrivilegeRepositoryViewRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiPrivilegeRepositoryViewRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiPrivilegeRepositoryViewRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiPrivilegeRepositoryViewRequest(val *ApiPrivilegeRepositoryViewRequest) *NullableApiPrivilegeRepositoryViewRequest {
	return &NullableApiPrivilegeRepositoryViewRequest{value: val, isSet: true}
}

func (v NullableApiPrivilegeRepositoryViewRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiPrivilegeRepositoryViewRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


