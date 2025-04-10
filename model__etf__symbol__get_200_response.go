/*
Gurufocus Data Package API

API for accessing Gurufocus data packages, please go to [https://www.gurufocus.com/user/me?tab=account&subtab=api-token](https://www.gurufocus.com/user/me?tab=account&subtab=api-token) to view or generate authorization keys.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the EtfSymbolGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EtfSymbolGet200Response{}

// EtfSymbolGet200Response struct for EtfSymbolGet200Response
type EtfSymbolGet200Response struct {
	BasicInformation *EtfEtfBasicInformation `json:"basic_information,omitempty"`
	Dividends *EtfEtfDividends `json:"dividends,omitempty"`
	Fundamental *EtfEtfFundamental `json:"fundamental,omitempty"`
	KeyStatistics *EtfEtfKeyStatistics `json:"key_statistics,omitempty"`
	PortfolioHoldings *EtfEtfPortfolioHoldings `json:"portfolio_holdings,omitempty"`
	SectorBreakdowns *EtfEtfSectorBreakdowns `json:"sector_breakdowns,omitempty"`
}

// NewEtfSymbolGet200Response instantiates a new EtfSymbolGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEtfSymbolGet200Response() *EtfSymbolGet200Response {
	this := EtfSymbolGet200Response{}
	return &this
}

// NewEtfSymbolGet200ResponseWithDefaults instantiates a new EtfSymbolGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEtfSymbolGet200ResponseWithDefaults() *EtfSymbolGet200Response {
	this := EtfSymbolGet200Response{}
	return &this
}

// GetBasicInformation returns the BasicInformation field value if set, zero value otherwise.
func (o *EtfSymbolGet200Response) GetBasicInformation() EtfEtfBasicInformation {
	if o == nil || IsNil(o.BasicInformation) {
		var ret EtfEtfBasicInformation
		return ret
	}
	return *o.BasicInformation
}

// GetBasicInformationOk returns a tuple with the BasicInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtfSymbolGet200Response) GetBasicInformationOk() (*EtfEtfBasicInformation, bool) {
	if o == nil || IsNil(o.BasicInformation) {
		return nil, false
	}
	return o.BasicInformation, true
}

// HasBasicInformation returns a boolean if a field has been set.
func (o *EtfSymbolGet200Response) HasBasicInformation() bool {
	if o != nil && !IsNil(o.BasicInformation) {
		return true
	}

	return false
}

// SetBasicInformation gets a reference to the given EtfEtfBasicInformation and assigns it to the BasicInformation field.
func (o *EtfSymbolGet200Response) SetBasicInformation(v EtfEtfBasicInformation) {
	o.BasicInformation = &v
}

// GetDividends returns the Dividends field value if set, zero value otherwise.
func (o *EtfSymbolGet200Response) GetDividends() EtfEtfDividends {
	if o == nil || IsNil(o.Dividends) {
		var ret EtfEtfDividends
		return ret
	}
	return *o.Dividends
}

// GetDividendsOk returns a tuple with the Dividends field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtfSymbolGet200Response) GetDividendsOk() (*EtfEtfDividends, bool) {
	if o == nil || IsNil(o.Dividends) {
		return nil, false
	}
	return o.Dividends, true
}

// HasDividends returns a boolean if a field has been set.
func (o *EtfSymbolGet200Response) HasDividends() bool {
	if o != nil && !IsNil(o.Dividends) {
		return true
	}

	return false
}

// SetDividends gets a reference to the given EtfEtfDividends and assigns it to the Dividends field.
func (o *EtfSymbolGet200Response) SetDividends(v EtfEtfDividends) {
	o.Dividends = &v
}

// GetFundamental returns the Fundamental field value if set, zero value otherwise.
func (o *EtfSymbolGet200Response) GetFundamental() EtfEtfFundamental {
	if o == nil || IsNil(o.Fundamental) {
		var ret EtfEtfFundamental
		return ret
	}
	return *o.Fundamental
}

// GetFundamentalOk returns a tuple with the Fundamental field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtfSymbolGet200Response) GetFundamentalOk() (*EtfEtfFundamental, bool) {
	if o == nil || IsNil(o.Fundamental) {
		return nil, false
	}
	return o.Fundamental, true
}

// HasFundamental returns a boolean if a field has been set.
func (o *EtfSymbolGet200Response) HasFundamental() bool {
	if o != nil && !IsNil(o.Fundamental) {
		return true
	}

	return false
}

// SetFundamental gets a reference to the given EtfEtfFundamental and assigns it to the Fundamental field.
func (o *EtfSymbolGet200Response) SetFundamental(v EtfEtfFundamental) {
	o.Fundamental = &v
}

// GetKeyStatistics returns the KeyStatistics field value if set, zero value otherwise.
func (o *EtfSymbolGet200Response) GetKeyStatistics() EtfEtfKeyStatistics {
	if o == nil || IsNil(o.KeyStatistics) {
		var ret EtfEtfKeyStatistics
		return ret
	}
	return *o.KeyStatistics
}

// GetKeyStatisticsOk returns a tuple with the KeyStatistics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtfSymbolGet200Response) GetKeyStatisticsOk() (*EtfEtfKeyStatistics, bool) {
	if o == nil || IsNil(o.KeyStatistics) {
		return nil, false
	}
	return o.KeyStatistics, true
}

// HasKeyStatistics returns a boolean if a field has been set.
func (o *EtfSymbolGet200Response) HasKeyStatistics() bool {
	if o != nil && !IsNil(o.KeyStatistics) {
		return true
	}

	return false
}

// SetKeyStatistics gets a reference to the given EtfEtfKeyStatistics and assigns it to the KeyStatistics field.
func (o *EtfSymbolGet200Response) SetKeyStatistics(v EtfEtfKeyStatistics) {
	o.KeyStatistics = &v
}

// GetPortfolioHoldings returns the PortfolioHoldings field value if set, zero value otherwise.
func (o *EtfSymbolGet200Response) GetPortfolioHoldings() EtfEtfPortfolioHoldings {
	if o == nil || IsNil(o.PortfolioHoldings) {
		var ret EtfEtfPortfolioHoldings
		return ret
	}
	return *o.PortfolioHoldings
}

// GetPortfolioHoldingsOk returns a tuple with the PortfolioHoldings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtfSymbolGet200Response) GetPortfolioHoldingsOk() (*EtfEtfPortfolioHoldings, bool) {
	if o == nil || IsNil(o.PortfolioHoldings) {
		return nil, false
	}
	return o.PortfolioHoldings, true
}

// HasPortfolioHoldings returns a boolean if a field has been set.
func (o *EtfSymbolGet200Response) HasPortfolioHoldings() bool {
	if o != nil && !IsNil(o.PortfolioHoldings) {
		return true
	}

	return false
}

// SetPortfolioHoldings gets a reference to the given EtfEtfPortfolioHoldings and assigns it to the PortfolioHoldings field.
func (o *EtfSymbolGet200Response) SetPortfolioHoldings(v EtfEtfPortfolioHoldings) {
	o.PortfolioHoldings = &v
}

// GetSectorBreakdowns returns the SectorBreakdowns field value if set, zero value otherwise.
func (o *EtfSymbolGet200Response) GetSectorBreakdowns() EtfEtfSectorBreakdowns {
	if o == nil || IsNil(o.SectorBreakdowns) {
		var ret EtfEtfSectorBreakdowns
		return ret
	}
	return *o.SectorBreakdowns
}

// GetSectorBreakdownsOk returns a tuple with the SectorBreakdowns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtfSymbolGet200Response) GetSectorBreakdownsOk() (*EtfEtfSectorBreakdowns, bool) {
	if o == nil || IsNil(o.SectorBreakdowns) {
		return nil, false
	}
	return o.SectorBreakdowns, true
}

// HasSectorBreakdowns returns a boolean if a field has been set.
func (o *EtfSymbolGet200Response) HasSectorBreakdowns() bool {
	if o != nil && !IsNil(o.SectorBreakdowns) {
		return true
	}

	return false
}

// SetSectorBreakdowns gets a reference to the given EtfEtfSectorBreakdowns and assigns it to the SectorBreakdowns field.
func (o *EtfSymbolGet200Response) SetSectorBreakdowns(v EtfEtfSectorBreakdowns) {
	o.SectorBreakdowns = &v
}

func (o EtfSymbolGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EtfSymbolGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BasicInformation) {
		toSerialize["basic_information"] = o.BasicInformation
	}
	if !IsNil(o.Dividends) {
		toSerialize["dividends"] = o.Dividends
	}
	if !IsNil(o.Fundamental) {
		toSerialize["fundamental"] = o.Fundamental
	}
	if !IsNil(o.KeyStatistics) {
		toSerialize["key_statistics"] = o.KeyStatistics
	}
	if !IsNil(o.PortfolioHoldings) {
		toSerialize["portfolio_holdings"] = o.PortfolioHoldings
	}
	if !IsNil(o.SectorBreakdowns) {
		toSerialize["sector_breakdowns"] = o.SectorBreakdowns
	}
	return toSerialize, nil
}

type NullableEtfSymbolGet200Response struct {
	value *EtfSymbolGet200Response
	isSet bool
}

func (v NullableEtfSymbolGet200Response) Get() *EtfSymbolGet200Response {
	return v.value
}

func (v *NullableEtfSymbolGet200Response) Set(val *EtfSymbolGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEtfSymbolGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEtfSymbolGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEtfSymbolGet200Response(val *EtfSymbolGet200Response) *NullableEtfSymbolGet200Response {
	return &NullableEtfSymbolGet200Response{value: val, isSet: true}
}

func (v NullableEtfSymbolGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEtfSymbolGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


