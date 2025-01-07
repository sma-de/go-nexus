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

// checks if the ApiUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiUser{}

// ApiUser struct for ApiUser
type ApiUser struct {
	// The userid which is required for login. This value cannot be changed.
	UserId *string `json:"userId,omitempty"`
	// The first name of the user.
	FirstName *string `json:"firstName,omitempty"`
	// The last name of the user.
	LastName *string `json:"lastName,omitempty"`
	// The email address associated with the user.
	EmailAddress *string `json:"emailAddress,omitempty"`
	// The user source which is the origin of this user. This value cannot be changed.
	Source *string `json:"source,omitempty"`
	// The user's status, e.g. active or disabled.
	Status string `json:"status"`
	// Indicates whether the user's properties could be modified by the Nexus Repository Manager. When false only roles are considered during update.
	ReadOnly *bool `json:"readOnly,omitempty"`
	// The roles which the user has been assigned within Nexus.
	Roles []string `json:"roles,omitempty"`
	// The roles which the user has been assigned in an external source, e.g. LDAP group. These cannot be changed within the Nexus Repository Manager.
	ExternalRoles []string `json:"externalRoles,omitempty"`
}

type _ApiUser ApiUser

// NewApiUser instantiates a new ApiUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiUser(status string) *ApiUser {
	this := ApiUser{}
	this.Status = status
	return &this
}

// NewApiUserWithDefaults instantiates a new ApiUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiUserWithDefaults() *ApiUser {
	this := ApiUser{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *ApiUser) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUser) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *ApiUser) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *ApiUser) SetUserId(v string) {
	o.UserId = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *ApiUser) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUser) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *ApiUser) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *ApiUser) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *ApiUser) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUser) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *ApiUser) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *ApiUser) SetLastName(v string) {
	o.LastName = &v
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *ApiUser) GetEmailAddress() string {
	if o == nil || IsNil(o.EmailAddress) {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUser) GetEmailAddressOk() (*string, bool) {
	if o == nil || IsNil(o.EmailAddress) {
		return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *ApiUser) HasEmailAddress() bool {
	if o != nil && !IsNil(o.EmailAddress) {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *ApiUser) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ApiUser) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUser) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ApiUser) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *ApiUser) SetSource(v string) {
	o.Source = &v
}

// GetStatus returns the Status field value
func (o *ApiUser) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ApiUser) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ApiUser) SetStatus(v string) {
	o.Status = v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *ApiUser) GetReadOnly() bool {
	if o == nil || IsNil(o.ReadOnly) {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUser) GetReadOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.ReadOnly) {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *ApiUser) HasReadOnly() bool {
	if o != nil && !IsNil(o.ReadOnly) {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *ApiUser) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *ApiUser) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUser) GetRolesOk() ([]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *ApiUser) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *ApiUser) SetRoles(v []string) {
	o.Roles = v
}

// GetExternalRoles returns the ExternalRoles field value if set, zero value otherwise.
func (o *ApiUser) GetExternalRoles() []string {
	if o == nil || IsNil(o.ExternalRoles) {
		var ret []string
		return ret
	}
	return o.ExternalRoles
}

// GetExternalRolesOk returns a tuple with the ExternalRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUser) GetExternalRolesOk() ([]string, bool) {
	if o == nil || IsNil(o.ExternalRoles) {
		return nil, false
	}
	return o.ExternalRoles, true
}

// HasExternalRoles returns a boolean if a field has been set.
func (o *ApiUser) HasExternalRoles() bool {
	if o != nil && !IsNil(o.ExternalRoles) {
		return true
	}

	return false
}

// SetExternalRoles gets a reference to the given []string and assigns it to the ExternalRoles field.
func (o *ApiUser) SetExternalRoles(v []string) {
	o.ExternalRoles = v
}

func (o ApiUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !IsNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	if !IsNil(o.EmailAddress) {
		toSerialize["emailAddress"] = o.EmailAddress
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	toSerialize["status"] = o.Status
	if !IsNil(o.ReadOnly) {
		toSerialize["readOnly"] = o.ReadOnly
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !IsNil(o.ExternalRoles) {
		toSerialize["externalRoles"] = o.ExternalRoles
	}
	return toSerialize, nil
}

func (o *ApiUser) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
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

	varApiUser := _ApiUser{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiUser)

	if err != nil {
		return err
	}

	*o = ApiUser(varApiUser)

	return err
}

type NullableApiUser struct {
	value *ApiUser
	isSet bool
}

func (v NullableApiUser) Get() *ApiUser {
	return v.value
}

func (v *NullableApiUser) Set(val *ApiUser) {
	v.value = val
	v.isSet = true
}

func (v NullableApiUser) IsSet() bool {
	return v.isSet
}

func (v *NullableApiUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiUser(val *ApiUser) *NullableApiUser {
	return &NullableApiUser{value: val, isSet: true}
}

func (v NullableApiUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


