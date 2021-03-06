/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 0.15
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package oscgo

import (
	"bytes"
	"encoding/json"
)

// ResourceTag Information about the tag.
type ResourceTag struct {
	// The key of the tag, with a minimum of 1 character.
	Key string `json:"Key"`
	// The value of the tag, between 0 and 255 characters.
	Value string `json:"Value"`
}

// GetKey returns the Key field value
func (o *ResourceTag) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// SetKey sets field value
func (o *ResourceTag) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value
func (o *ResourceTag) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// SetValue sets field value
func (o *ResourceTag) SetValue(v string) {
	o.Value = v
}

type NullableResourceTag struct {
	Value        ResourceTag
	ExplicitNull bool
}

func (v NullableResourceTag) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableResourceTag) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
