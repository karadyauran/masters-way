/*
Masters way chat API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SchemasGetChatPreviewResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemasGetChatPreviewResponse{}

// SchemasGetChatPreviewResponse struct for SchemasGetChatPreviewResponse
type SchemasGetChatPreviewResponse struct {
	UnreadMessagesAmount int32
}

type _SchemasGetChatPreviewResponse SchemasGetChatPreviewResponse

// NewSchemasGetChatPreviewResponse instantiates a new SchemasGetChatPreviewResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasGetChatPreviewResponse(unreadMessagesAmount int32) *SchemasGetChatPreviewResponse {
	this := SchemasGetChatPreviewResponse{}
	this.UnreadMessagesAmount = unreadMessagesAmount
	return &this
}

// NewSchemasGetChatPreviewResponseWithDefaults instantiates a new SchemasGetChatPreviewResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasGetChatPreviewResponseWithDefaults() *SchemasGetChatPreviewResponse {
	this := SchemasGetChatPreviewResponse{}
	return &this
}

// GetUnreadMessagesAmount returns the UnreadMessagesAmount field value
func (o *SchemasGetChatPreviewResponse) GetUnreadMessagesAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UnreadMessagesAmount
}

// GetUnreadMessagesAmountOk returns a tuple with the UnreadMessagesAmount field value
// and a boolean to check if the value has been set.
func (o *SchemasGetChatPreviewResponse) GetUnreadMessagesAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnreadMessagesAmount, true
}

// SetUnreadMessagesAmount sets field value
func (o *SchemasGetChatPreviewResponse) SetUnreadMessagesAmount(v int32) {
	o.UnreadMessagesAmount = v
}

func (o SchemasGetChatPreviewResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemasGetChatPreviewResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["unreadMessagesAmount"] = o.UnreadMessagesAmount
	return toSerialize, nil
}

func (o *SchemasGetChatPreviewResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"unreadMessagesAmount",
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

	varSchemasGetChatPreviewResponse := _SchemasGetChatPreviewResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSchemasGetChatPreviewResponse)

	if err != nil {
		return err
	}

	*o = SchemasGetChatPreviewResponse(varSchemasGetChatPreviewResponse)

	return err
}

type NullableSchemasGetChatPreviewResponse struct {
	value *SchemasGetChatPreviewResponse
	isSet bool
}

func (v NullableSchemasGetChatPreviewResponse) Get() *SchemasGetChatPreviewResponse {
	return v.value
}

func (v *NullableSchemasGetChatPreviewResponse) Set(val *SchemasGetChatPreviewResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemasGetChatPreviewResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemasGetChatPreviewResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemasGetChatPreviewResponse(val *SchemasGetChatPreviewResponse) *NullableSchemasGetChatPreviewResponse {
	return &NullableSchemasGetChatPreviewResponse{value: val, isSet: true}
}

func (v NullableSchemasGetChatPreviewResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemasGetChatPreviewResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


