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

// checks if the ValuationsIREITNODIRECTValuationRatios type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValuationsIREITNODIRECTValuationRatios{}

// ValuationsIREITNODIRECTValuationRatios struct for ValuationsIREITNODIRECTValuationRatios
type ValuationsIREITNODIRECTValuationRatios struct {
	// The Cyclically Adjusted PB Ratio, also known as the CAPB Ratio, is the stock price divided by the average of the inflation adjusted book value per share of a company over the past 10 years.
	CyclicallyAdjustedPbRatio *float32 `json:"cyclically_adjusted_pb_ratio,omitempty"`
	CyclicallyAdjustedPriceToFcf *float32 `json:"cyclically_adjusted_price_to_fcf,omitempty"`
	// The Cyclically Adjusted PS Ratio, also referred to as the CAPS Ratio, is the stock price divided by the average of the inflation adjusted revenue per share of a company over the past 10 years.
	CyclicallyAdjustedPsRatio *float32 `json:"cyclically_adjusted_ps_ratio,omitempty"`
	// The standard definition of earnings yield is the earnings per share divided by the price of a share. It's the inverse of P/E and shows the amount of money earned compared to the price you pay for a share.
	EarningYieldGreenblatt *float32 `json:"earning_yield_greenblatt,omitempty"`
	// EV-to-EBIT is calculated as Enterprise Value divided by its EBIT.
	EnterpriseValueToEbit *float32 `json:"enterprise_value_to_ebit,omitempty"`
	// EV-to-EBITDA is calculated as enterprise value divided by its EBITDA.
	EnterpriseValueToEbitda *float32 `json:"enterprise_value_to_ebitda,omitempty"`
	EnterpriseValueToFcf *float32 `json:"enterprise_value_to_fcf,omitempty"`
	// EV-to-Revenue is calculated as enterprise value divided by its revenue.
	EnterpriseValueToRevenue *float32 `json:"enterprise_value_to_revenue,omitempty"`
	// Free cash flow yield: the free cash flow divided by share price
	FcfYield *float32 `json:"fcf_yield,omitempty"`
	// Companies use the price-to-book ratio to compare a firm's market to book value by dividing price per share by book value per share. Some people know it as the price-equity ratio.
	PbRatio *float32 `json:"pb_ratio,omitempty"`
	// <p>The PE ratio, or Price-to-Earnings ratio, is the most widely used metric in the valuation of stocks. It is calculated as:  <b>PE Ratio = Share Price / {{eps_diluated}}</b>.   It can also be calculated from the numbers for the whole company:  <b>PE Ratio = {{mktcap}} / {{net_income}}</b>.</p>  <p>There are at least three kinds of PE ratios used among different investors. They are Trailing Twelve Month PE Ratio({{pettm}}), {{forwardPE}}, and {{penri}}. A new PE ratio based on inflation-adjusted normalized PE ratio is called {{ShillerPE}}, after Yale professor Robert Shiller.</p>  <p>In the calculation of {{pettm}}, the earnings per share used are the earnings per share over the past 12 months({{ttm_eps}}). For {{forwardPE}}, the earnings are the expected earnings for the next twelve months({{forward_eps}}). In the case of {{penri}}, the reported earnings less the non-recurring items are used({{eps_nri}}).For the {{ShillerPE}}, the earnings of the past 10 years are inflation-adjusted and averaged. Since {{ShillerPE}} looks at the average over the last 10 years, it is also called PE10.</p>
	PeRatio *float32 `json:"pe_ratio,omitempty"`
	// PEG is defined as the PE Ratio without NRI divided by the growth ratio. The growth rate we use is the 5-year average EBITDA growth rate.
	PegRatio *float32 `json:"peg_ratio,omitempty"`
	Penri *float32 `json:"penri,omitempty"`
	// Price to Funds From Operations is an equity valuation metric used to compare a company's per-share market price to its per-share amount of Funds From Operations (FFO).
	PriceToFfo *float32 `json:"price_to_ffo,omitempty"`
	// Price to free cash flow is an equity valuation metric used to compare a company's per-share market price to its per-share amount of free cash flow (FCF).
	PriceToFreeCashFlow *float32 `json:"price_to_free_cash_flow,omitempty"`
	// The price/cash flow ratio (also called price-to-cash flow ratio or P/CF), is a ratio used to compare a company's market value to its cash flow.
	PriceToOperatingCashFlow *float32 `json:"price_to_operating_cash_flow,omitempty"`
	// In 1986 Berkshire Hathaway Shareholder Letter, Warren Buffett defined owner earnings as follows:  \"These represent (a) reported earnings plus (b) depreciation, depletion, amortization, and certain other non-cash charges...less (c) the average annual amount of capitalized expenditures for plant and equipment, etc. that the business requires to fully maintain its long-term competitive position and its unit volume. (If the business requires additional working capital to maintain its competitive position and unit volume, the increment also should be included in (c))...Our owner-earnings equation does not yield the deceptively precise figures provided by GAAP, since (c) must be a guess - and one sometimes very difficult to make. Despite this problem, we consider the owner earnings figure, not the GAAP figure, to be the relevant item for valuation purposes - both for investors in buying stocks and for managers in buying entire businesses...All of this points up the absurdity of the 'cash flow' numbers that are often set forth in Wall Street reports. These numbers routinely include (a) plus (b) - but do not subtract (c).\"
	PriceToOwnerEarnings *float32 `json:"price_to_owner_earnings,omitempty"`
	// The Price to Tangible Book Value ratio (PTBV) expresses share price as a proportion of the company's tangible book value reported on the company's balance sheet.
	PriceToTangibleBook *float32 `json:"price_to_tangible_book,omitempty"`
	// Price–sales ratio, P/S ratio, or PSR, is a valuation metric for stocks.
	PsRatio *float32 `json:"ps_ratio,omitempty"`
	// Yacktman defines forward rate of return as the normalized free cash flow yield plus real growth plus inflation.
	RateOfReturnValue *float32 `json:"rate_of_return_value,omitempty"`
	// Price earnings ratio is based on average inflation-adjusted earnings from the previous 10 years, known as the Cyclically Adjusted PE Ratio (CAPE Ratio), Shiller PE Ratio, or PE 10
	ShillerPeRatio *float32 `json:"shiller_pe_ratio,omitempty"`
	// The dividend yield is the ratio of a company's annual dividend compared to its share price.
	Yield *float32 `json:"yield,omitempty"`
}

// NewValuationsIREITNODIRECTValuationRatios instantiates a new ValuationsIREITNODIRECTValuationRatios object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValuationsIREITNODIRECTValuationRatios() *ValuationsIREITNODIRECTValuationRatios {
	this := ValuationsIREITNODIRECTValuationRatios{}
	return &this
}

// NewValuationsIREITNODIRECTValuationRatiosWithDefaults instantiates a new ValuationsIREITNODIRECTValuationRatios object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValuationsIREITNODIRECTValuationRatiosWithDefaults() *ValuationsIREITNODIRECTValuationRatios {
	this := ValuationsIREITNODIRECTValuationRatios{}
	return &this
}

// GetCyclicallyAdjustedPbRatio returns the CyclicallyAdjustedPbRatio field value if set, zero value otherwise.
func (o *ValuationsIREITNODIRECTValuationRatios) GetCyclicallyAdjustedPbRatio() float32 {
	if o == nil || IsNil(o.CyclicallyAdjustedPbRatio) {
		var ret float32
		return ret
	}
	return *o.CyclicallyAdjustedPbRatio
}

// GetCyclicallyAdjustedPbRatioOk returns a tuple with the CyclicallyAdjustedPbRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsIREITNODIRECTValuationRatios) GetCyclicallyAdjustedPbRatioOk() (*float32, bool) {
	if o == nil || IsNil(o.CyclicallyAdjustedPbRatio) {
		return nil, false
	}
	return o.CyclicallyAdjustedPbRatio, true
}

// HasCyclicallyAdjustedPbRatio returns a boolean if a field has been set.
func (o *ValuationsIREITNODIRECTValuationRatios) HasCyclicallyAdjustedPbRatio() bool {
	if o != nil && !IsNil(o.CyclicallyAdjustedPbRatio) {
		return true
	}

	return false
}

// SetCyclicallyAdjustedPbRatio gets a reference to the given float32 and assigns it to the CyclicallyAdjustedPbRatio field.
func (o *ValuationsIREITNODIRECTValuationRatios) SetCyclicallyAdjustedPbRatio(v float32) {
	o.CyclicallyAdjustedPbRatio = &v
}

// GetCyclicallyAdjustedPriceToFcf returns the CyclicallyAdjustedPriceToFcf field value if set, zero value otherwise.
func (o *ValuationsIREITNODIRECTValuationRatios) GetCyclicallyAdjustedPriceToFcf() float32 {
	if o == nil || IsNil(o.CyclicallyAdjustedPriceToFcf) {
		var ret float32
		return ret
	}
	return *o.CyclicallyAdjustedPriceToFcf
}

// GetCyclicallyAdjustedPriceToFcfOk returns a tuple with the CyclicallyAdjustedPriceToFcf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsIREITNODIRECTValuationRatios) GetCyclicallyAdjustedPriceToFcfOk() (*float32, bool) {
	if o == nil || IsNil(o.CyclicallyAdjustedPriceToFcf) {
		return nil, false
	}
	return o.CyclicallyAdjustedPriceToFcf, true
}

// HasCyclicallyAdjustedPriceToFcf returns a boolean if a field has been set.
func (o *ValuationsIREITNODIRECTValuationRatios) HasCyclicallyAdjustedPriceToFcf() bool {
	if o != nil && !IsNil(o.CyclicallyAdjustedPriceToFcf) {
		return true
	}

	return false
}

// SetCyclicallyAdjustedPriceToFcf gets a reference to the given float32 and assigns it to the CyclicallyAdjustedPriceToFcf field.
func (o *ValuationsIREITNODIRECTValuationRatios) SetCyclicallyAdjustedPriceToFcf(v float32) {
	o.CyclicallyAdjustedPriceToFcf = &v
}

// GetCyclicallyAdjustedPsRatio returns the CyclicallyAdjustedPsRatio field value if set, zero value otherwise.
func (o *ValuationsIREITNODIRECTValuationRatios) GetCyclicallyAdjustedPsRatio() float32 {
	if o == nil || IsNil(o.CyclicallyAdjustedPsRatio) {
		var ret float32
		return ret
	}
	return *o.CyclicallyAdjustedPsRatio
}

// GetCyclicallyAdjustedPsRatioOk returns a tuple with the CyclicallyAdjustedPsRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsIREITNODIRECTValuationRatios) GetCyclicallyAdjustedPsRatioOk() (*float32, bool) {
	if o == nil || IsNil(o.CyclicallyAdjustedPsRatio) {
		return nil, false
	}
	return o.CyclicallyAdjustedPsRatio, true
}

// HasCyclicallyAdjustedPsRatio returns a boolean if a field has been set.
func (o *ValuationsIREITNODIRECTValuationRatios) HasCyclicallyAdjustedPsRatio() bool {
	if o != nil && !IsNil(o.CyclicallyAdjustedPsRatio) {
		return true
	}

	return false
}

// SetCyclicallyAdjustedPsRatio gets a reference to the given float32 and assigns it to the CyclicallyAdjustedPsRatio field.
func (o *ValuationsIREITNODIRECTValuationRatios) SetCyclicallyAdjustedPsRatio(v float32) {
	o.CyclicallyAdjustedPsRatio = &v
}

// GetEarningYieldGreenblatt returns the EarningYieldGreenblatt field value if set, zero value otherwise.
func (o *ValuationsIREITNODIRECTValuationRatios) GetEarningYieldGreenblatt() float32 {
	if o == nil || IsNil(o.EarningYieldGreenblatt) {
		var ret float32
		return ret
	}
	return *o.EarningYieldGreenblatt
}

// GetEarningYieldGreenblattOk returns a tuple with the EarningYieldGreenblatt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsIREITNODIRECTValuationRatios) GetEarningYieldGreenblattOk() (*float32, bool) {
	if o == nil || IsNil(o.EarningYieldGreenblatt) {
		return nil, false
	}
	return o.EarningYieldGreenblatt, true
}

// HasEarningYieldGreenblatt returns a boolean if a field has been set.
func (o *ValuationsIREITNODIRECTValuationRatios) HasEarningYieldGreenblatt() bool {
	if o != nil && !IsNil(o.EarningYieldGreenblatt) {
		return true
	}

	return false
}

// SetEarningYieldGreenblatt gets a reference to the given float32 and assigns it to the EarningYieldGreenblatt field.
func (o *ValuationsIREITNODIRECTValuationRatios) SetEarningYieldGreenblatt(v float32) {
	o.EarningYieldGreenblatt = &v
}

// GetEnterpriseValueToEbit returns the EnterpriseValueToEbit field value if set, zero value otherwise.
func (o *ValuationsIREITNODIRECTValuationRatios) GetEnterpriseValueToEbit() float32 {
	if o == nil || IsNil(o.EnterpriseValueToEbit) {
		var ret float32
		return ret
	}
	return *o.EnterpriseValueToEbit
}

// GetEnterpriseValueToEbitOk returns a tuple with the EnterpriseValueToEbit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsIREITNODIRECTValuationRatios) GetEnterpriseValueToEbitOk() (*float32, bool) {
	if o == nil || IsNil(o.EnterpriseValueToEbit) {
		return nil, false
	}
	return o.EnterpriseValueToEbit, true
}

// HasEnterpriseValueToEbit returns a boolean if a field has been set.
func (o *ValuationsIREITNODIRECTValuationRatios) HasEnterpriseValueToEbit() bool {
	if o != nil && !IsNil(o.EnterpriseValueToEbit) {
		return true
	}

	return false
}

// SetEnterpriseValueToEbit gets a reference to the given float32 and assigns it to the EnterpriseValueToEbit field.
func (o *ValuationsIREITNODIRECTValuationRatios) SetEnterpriseValueToEbit(v float32) {
	o.EnterpriseValueToEbit = &v
}

// GetEnterpriseValueToEbitda returns the EnterpriseValueToEbitda field value if set, zero value otherwise.
func (o *ValuationsIREITNODIRECTValuationRatios) GetEnterpriseValueToEbitda() float32 {
	if o == nil || IsNil(o.EnterpriseValueToEbitda) {
		var ret float32
		return ret
	}
	return *o.EnterpriseValueToEbitda
}

// GetEnterpriseValueToEbitdaOk returns a tuple with the EnterpriseValueToEbitda field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsIREITNODIRECTValuationRatios) GetEnterpriseValueToEbitdaOk() (*float32, bool) {
	if o == nil || IsNil(o.EnterpriseValueToEbitda) {
		return nil, false
	}
	return o.EnterpriseValueToEbitda, true
}

// HasEnterpriseValueToEbitda returns a boolean if a field has been set.
func (o *ValuationsIREITNODIRECTValuationRatios) HasEnterpriseValueToEbitda() bool {
	if o != nil && !IsNil(o.EnterpriseValueToEbitda) {
		return true
	}

	return false
}

// SetEnterpriseValueToEbitda gets a reference to the given float32 and assigns it to the EnterpriseValueToEbitda field.
func (o *ValuationsIREITNODIRECTValuationRatios) SetEnterpriseValueToEbitda(v float32) {
	o.EnterpriseValueToEbitda = &v
}

// GetEnterpriseValueToFcf returns the EnterpriseValueToFcf field value if set, zero value otherwise.
func (o *ValuationsIREITNODIRECTValuationRatios) GetEnterpriseValueToFcf() float32 {
	if o == nil || IsNil(o.EnterpriseValueToFcf) {
		var ret float32
		return ret
	}
	return *o.EnterpriseValueToFcf
}

// GetEnterpriseValueToFcfOk returns a tuple with the EnterpriseValueToFcf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsIREITNODIRECTValuationRatios) GetEnterpriseValueToFcfOk() (*float32, bool) {
	if o == nil || IsNil(o.EnterpriseValueToFcf) {
		return nil, false
	}
	return o.EnterpriseValueToFcf, true
}

// HasEnterpriseValueToFcf returns a boolean if a field has been set.
func (o *ValuationsIREITNODIRECTValuationRatios) HasEnterpriseValueToFcf() bool {
	if o != nil && !IsNil(o.EnterpriseValueToFcf) {
		return true
	}

	return false
}

// SetEnterpriseValueToFcf gets a reference to the given float32 and assigns it to the EnterpriseValueToFcf field.
func (o *ValuationsIREITNODIRECTValuationRatios) SetEnterpriseValueToFcf(v float32) {
	o.EnterpriseValueToFcf = &v
}

// GetEnterpriseValueToRevenue returns the EnterpriseValueToRevenue field value if set, zero value otherwise.
func (o *ValuationsIREITNODIRECTValuationRatios) GetEnterpriseValueToRevenue() float32 {
	if o == nil || IsNil(o.EnterpriseValueToRevenue) {
		var ret float32
		return ret
	}
	return *o.EnterpriseValueToRevenue
}

// GetEnterpriseValueToRevenueOk returns a tuple with the EnterpriseValueToRevenue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsIREITNODIRECTValuationRatios) GetEnterpriseValueToRevenueOk() (*float32, bool) {
	if o == nil || IsNil(o.EnterpriseValueToRevenue) {
		return nil, false
	}
	return o.EnterpriseValueToRevenue, true
}

// HasEnterpriseValueToRevenue returns a boolean if a field has been set.
func (o *ValuationsIREITNODIRECTValuationRatios) HasEnterpriseValueToRevenue() bool {
	if o != nil && !IsNil(o.EnterpriseValueToRevenue) {
		return true
	}

	return false
}

// SetEnterpriseValueToRevenue gets a reference to the given float32 and assigns it to the EnterpriseValueToRevenue field.
func (o *ValuationsIREITNODIRECTValuationRatios) SetEnterpriseValueToRevenue(v float32) {
	o.EnterpriseValueToRevenue = &v
}

// GetFcfYield returns the FcfYield field value if set, zero value otherwise.
func (o *ValuationsIREITNODIRECTValuationRatios) GetFcfYield() float32 {
	if o == nil || IsNil(o.FcfYield) {
		var ret float32
		return ret
	}
	return *o.FcfYield
}

// GetFcfYieldOk returns a tuple with the FcfYield field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsIREITNODIRECTValuationRatios) GetFcfYieldOk() (*float32, bool) {
	if o == nil || IsNil(o.FcfYield) {
		return nil, false
	}
	return o.FcfYield, true
}

// HasFcfYield returns a boolean if a field has been set.
func (o *ValuationsIREITNODIRECTValuationRatios) HasFcfYield() bool {
	if o != nil && !IsNil(o.FcfYield) {
		return true
	}

	return false
}

// SetFcfYield gets a reference to the given float32 and assigns it to the FcfYield field.
func (o *ValuationsIREITNODIRECTValuationRatios) SetFcfYield(v float32) {
	o.FcfYield = &v
}

// GetPbRatio returns the PbRatio field value if set, zero value otherwise.
func (o *ValuationsIREITNODIRECTValuationRatios) GetPbRatio() float32 {
	if o == nil || IsNil(o.PbRatio) {
		var ret float32
		return ret
	}
	return *o.PbRatio
}

// GetPbRatioOk returns a tuple with the PbRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsIREITNODIRECTValuationRatios) GetPbRatioOk() (*float32, bool) {
	if o == nil || IsNil(o.PbRatio) {
		return nil, false
	}
	return o.PbRatio, true
}

// HasPbRatio returns a boolean if a field has been set.
func (o *ValuationsIREITNODIRECTValuationRatios) HasPbRatio() bool {
	if o != nil && !IsNil(o.PbRatio) {
		return true
	}

	return false
}

// SetPbRatio gets a reference to the given float32 and assigns it to the PbRatio field.
func (o *ValuationsIREITNODIRECTValuationRatios) SetPbRatio(v float32) {
	o.PbRatio = &v
}

// GetPeRatio returns the PeRatio field value if set, zero value otherwise.
func (o *ValuationsIREITNODIRECTValuationRatios) GetPeRatio() float32 {
	if o == nil || IsNil(o.PeRatio) {
		var ret float32
		return ret
	}
	return *o.PeRatio
}

// GetPeRatioOk returns a tuple with the PeRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsIREITNODIRECTValuationRatios) GetPeRatioOk() (*float32, bool) {
	if o == nil || IsNil(o.PeRatio) {
		return nil, false
	}
	return o.PeRatio, true
}

// HasPeRatio returns a boolean if a field has been set.
func (o *ValuationsIREITNODIRECTValuationRatios) HasPeRatio() bool {
	if o != nil && !IsNil(o.PeRatio) {
		return true
	}

	return false
}

// SetPeRatio gets a reference to the given float32 and assigns it to the PeRatio field.
func (o *ValuationsIREITNODIRECTValuationRatios) SetPeRatio(v float32) {
	o.PeRatio = &v
}

// GetPegRatio returns the PegRatio field value if set, zero value otherwise.
func (o *ValuationsIREITNODIRECTValuationRatios) GetPegRatio() float32 {
	if o == nil || IsNil(o.PegRatio) {
		var ret float32
		return ret
	}
	return *o.PegRatio
}

// GetPegRatioOk returns a tuple with the PegRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsIREITNODIRECTValuationRatios) GetPegRatioOk() (*float32, bool) {
	if o == nil || IsNil(o.PegRatio) {
		return nil, false
	}
	return o.PegRatio, true
}

// HasPegRatio returns a boolean if a field has been set.
func (o *ValuationsIREITNODIRECTValuationRatios) HasPegRatio() bool {
	if o != nil && !IsNil(o.PegRatio) {
		return true
	}

	return false
}

// SetPegRatio gets a reference to the given float32 and assigns it to the PegRatio field.
func (o *ValuationsIREITNODIRECTValuationRatios) SetPegRatio(v float32) {
	o.PegRatio = &v
}

// GetPenri returns the Penri field value if set, zero value otherwise.
func (o *ValuationsIREITNODIRECTValuationRatios) GetPenri() float32 {
	if o == nil || IsNil(o.Penri) {
		var ret float32
		return ret
	}
	return *o.Penri
}

// GetPenriOk returns a tuple with the Penri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsIREITNODIRECTValuationRatios) GetPenriOk() (*float32, bool) {
	if o == nil || IsNil(o.Penri) {
		return nil, false
	}
	return o.Penri, true
}

// HasPenri returns a boolean if a field has been set.
func (o *ValuationsIREITNODIRECTValuationRatios) HasPenri() bool {
	if o != nil && !IsNil(o.Penri) {
		return true
	}

	return false
}

// SetPenri gets a reference to the given float32 and assigns it to the Penri field.
func (o *ValuationsIREITNODIRECTValuationRatios) SetPenri(v float32) {
	o.Penri = &v
}

// GetPriceToFfo returns the PriceToFfo field value if set, zero value otherwise.
func (o *ValuationsIREITNODIRECTValuationRatios) GetPriceToFfo() float32 {
	if o == nil || IsNil(o.PriceToFfo) {
		var ret float32
		return ret
	}
	return *o.PriceToFfo
}

// GetPriceToFfoOk returns a tuple with the PriceToFfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsIREITNODIRECTValuationRatios) GetPriceToFfoOk() (*float32, bool) {
	if o == nil || IsNil(o.PriceToFfo) {
		return nil, false
	}
	return o.PriceToFfo, true
}

// HasPriceToFfo returns a boolean if a field has been set.
func (o *ValuationsIREITNODIRECTValuationRatios) HasPriceToFfo() bool {
	if o != nil && !IsNil(o.PriceToFfo) {
		return true
	}

	return false
}

// SetPriceToFfo gets a reference to the given float32 and assigns it to the PriceToFfo field.
func (o *ValuationsIREITNODIRECTValuationRatios) SetPriceToFfo(v float32) {
	o.PriceToFfo = &v
}

// GetPriceToFreeCashFlow returns the PriceToFreeCashFlow field value if set, zero value otherwise.
func (o *ValuationsIREITNODIRECTValuationRatios) GetPriceToFreeCashFlow() float32 {
	if o == nil || IsNil(o.PriceToFreeCashFlow) {
		var ret float32
		return ret
	}
	return *o.PriceToFreeCashFlow
}

// GetPriceToFreeCashFlowOk returns a tuple with the PriceToFreeCashFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsIREITNODIRECTValuationRatios) GetPriceToFreeCashFlowOk() (*float32, bool) {
	if o == nil || IsNil(o.PriceToFreeCashFlow) {
		return nil, false
	}
	return o.PriceToFreeCashFlow, true
}

// HasPriceToFreeCashFlow returns a boolean if a field has been set.
func (o *ValuationsIREITNODIRECTValuationRatios) HasPriceToFreeCashFlow() bool {
	if o != nil && !IsNil(o.PriceToFreeCashFlow) {
		return true
	}

	return false
}

// SetPriceToFreeCashFlow gets a reference to the given float32 and assigns it to the PriceToFreeCashFlow field.
func (o *ValuationsIREITNODIRECTValuationRatios) SetPriceToFreeCashFlow(v float32) {
	o.PriceToFreeCashFlow = &v
}

// GetPriceToOperatingCashFlow returns the PriceToOperatingCashFlow field value if set, zero value otherwise.
func (o *ValuationsIREITNODIRECTValuationRatios) GetPriceToOperatingCashFlow() float32 {
	if o == nil || IsNil(o.PriceToOperatingCashFlow) {
		var ret float32
		return ret
	}
	return *o.PriceToOperatingCashFlow
}

// GetPriceToOperatingCashFlowOk returns a tuple with the PriceToOperatingCashFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsIREITNODIRECTValuationRatios) GetPriceToOperatingCashFlowOk() (*float32, bool) {
	if o == nil || IsNil(o.PriceToOperatingCashFlow) {
		return nil, false
	}
	return o.PriceToOperatingCashFlow, true
}

// HasPriceToOperatingCashFlow returns a boolean if a field has been set.
func (o *ValuationsIREITNODIRECTValuationRatios) HasPriceToOperatingCashFlow() bool {
	if o != nil && !IsNil(o.PriceToOperatingCashFlow) {
		return true
	}

	return false
}

// SetPriceToOperatingCashFlow gets a reference to the given float32 and assigns it to the PriceToOperatingCashFlow field.
func (o *ValuationsIREITNODIRECTValuationRatios) SetPriceToOperatingCashFlow(v float32) {
	o.PriceToOperatingCashFlow = &v
}

// GetPriceToOwnerEarnings returns the PriceToOwnerEarnings field value if set, zero value otherwise.
func (o *ValuationsIREITNODIRECTValuationRatios) GetPriceToOwnerEarnings() float32 {
	if o == nil || IsNil(o.PriceToOwnerEarnings) {
		var ret float32
		return ret
	}
	return *o.PriceToOwnerEarnings
}

// GetPriceToOwnerEarningsOk returns a tuple with the PriceToOwnerEarnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsIREITNODIRECTValuationRatios) GetPriceToOwnerEarningsOk() (*float32, bool) {
	if o == nil || IsNil(o.PriceToOwnerEarnings) {
		return nil, false
	}
	return o.PriceToOwnerEarnings, true
}

// HasPriceToOwnerEarnings returns a boolean if a field has been set.
func (o *ValuationsIREITNODIRECTValuationRatios) HasPriceToOwnerEarnings() bool {
	if o != nil && !IsNil(o.PriceToOwnerEarnings) {
		return true
	}

	return false
}

// SetPriceToOwnerEarnings gets a reference to the given float32 and assigns it to the PriceToOwnerEarnings field.
func (o *ValuationsIREITNODIRECTValuationRatios) SetPriceToOwnerEarnings(v float32) {
	o.PriceToOwnerEarnings = &v
}

// GetPriceToTangibleBook returns the PriceToTangibleBook field value if set, zero value otherwise.
func (o *ValuationsIREITNODIRECTValuationRatios) GetPriceToTangibleBook() float32 {
	if o == nil || IsNil(o.PriceToTangibleBook) {
		var ret float32
		return ret
	}
	return *o.PriceToTangibleBook
}

// GetPriceToTangibleBookOk returns a tuple with the PriceToTangibleBook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsIREITNODIRECTValuationRatios) GetPriceToTangibleBookOk() (*float32, bool) {
	if o == nil || IsNil(o.PriceToTangibleBook) {
		return nil, false
	}
	return o.PriceToTangibleBook, true
}

// HasPriceToTangibleBook returns a boolean if a field has been set.
func (o *ValuationsIREITNODIRECTValuationRatios) HasPriceToTangibleBook() bool {
	if o != nil && !IsNil(o.PriceToTangibleBook) {
		return true
	}

	return false
}

// SetPriceToTangibleBook gets a reference to the given float32 and assigns it to the PriceToTangibleBook field.
func (o *ValuationsIREITNODIRECTValuationRatios) SetPriceToTangibleBook(v float32) {
	o.PriceToTangibleBook = &v
}

// GetPsRatio returns the PsRatio field value if set, zero value otherwise.
func (o *ValuationsIREITNODIRECTValuationRatios) GetPsRatio() float32 {
	if o == nil || IsNil(o.PsRatio) {
		var ret float32
		return ret
	}
	return *o.PsRatio
}

// GetPsRatioOk returns a tuple with the PsRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsIREITNODIRECTValuationRatios) GetPsRatioOk() (*float32, bool) {
	if o == nil || IsNil(o.PsRatio) {
		return nil, false
	}
	return o.PsRatio, true
}

// HasPsRatio returns a boolean if a field has been set.
func (o *ValuationsIREITNODIRECTValuationRatios) HasPsRatio() bool {
	if o != nil && !IsNil(o.PsRatio) {
		return true
	}

	return false
}

// SetPsRatio gets a reference to the given float32 and assigns it to the PsRatio field.
func (o *ValuationsIREITNODIRECTValuationRatios) SetPsRatio(v float32) {
	o.PsRatio = &v
}

// GetRateOfReturnValue returns the RateOfReturnValue field value if set, zero value otherwise.
func (o *ValuationsIREITNODIRECTValuationRatios) GetRateOfReturnValue() float32 {
	if o == nil || IsNil(o.RateOfReturnValue) {
		var ret float32
		return ret
	}
	return *o.RateOfReturnValue
}

// GetRateOfReturnValueOk returns a tuple with the RateOfReturnValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsIREITNODIRECTValuationRatios) GetRateOfReturnValueOk() (*float32, bool) {
	if o == nil || IsNil(o.RateOfReturnValue) {
		return nil, false
	}
	return o.RateOfReturnValue, true
}

// HasRateOfReturnValue returns a boolean if a field has been set.
func (o *ValuationsIREITNODIRECTValuationRatios) HasRateOfReturnValue() bool {
	if o != nil && !IsNil(o.RateOfReturnValue) {
		return true
	}

	return false
}

// SetRateOfReturnValue gets a reference to the given float32 and assigns it to the RateOfReturnValue field.
func (o *ValuationsIREITNODIRECTValuationRatios) SetRateOfReturnValue(v float32) {
	o.RateOfReturnValue = &v
}

// GetShillerPeRatio returns the ShillerPeRatio field value if set, zero value otherwise.
func (o *ValuationsIREITNODIRECTValuationRatios) GetShillerPeRatio() float32 {
	if o == nil || IsNil(o.ShillerPeRatio) {
		var ret float32
		return ret
	}
	return *o.ShillerPeRatio
}

// GetShillerPeRatioOk returns a tuple with the ShillerPeRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsIREITNODIRECTValuationRatios) GetShillerPeRatioOk() (*float32, bool) {
	if o == nil || IsNil(o.ShillerPeRatio) {
		return nil, false
	}
	return o.ShillerPeRatio, true
}

// HasShillerPeRatio returns a boolean if a field has been set.
func (o *ValuationsIREITNODIRECTValuationRatios) HasShillerPeRatio() bool {
	if o != nil && !IsNil(o.ShillerPeRatio) {
		return true
	}

	return false
}

// SetShillerPeRatio gets a reference to the given float32 and assigns it to the ShillerPeRatio field.
func (o *ValuationsIREITNODIRECTValuationRatios) SetShillerPeRatio(v float32) {
	o.ShillerPeRatio = &v
}

// GetYield returns the Yield field value if set, zero value otherwise.
func (o *ValuationsIREITNODIRECTValuationRatios) GetYield() float32 {
	if o == nil || IsNil(o.Yield) {
		var ret float32
		return ret
	}
	return *o.Yield
}

// GetYieldOk returns a tuple with the Yield field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsIREITNODIRECTValuationRatios) GetYieldOk() (*float32, bool) {
	if o == nil || IsNil(o.Yield) {
		return nil, false
	}
	return o.Yield, true
}

// HasYield returns a boolean if a field has been set.
func (o *ValuationsIREITNODIRECTValuationRatios) HasYield() bool {
	if o != nil && !IsNil(o.Yield) {
		return true
	}

	return false
}

// SetYield gets a reference to the given float32 and assigns it to the Yield field.
func (o *ValuationsIREITNODIRECTValuationRatios) SetYield(v float32) {
	o.Yield = &v
}

func (o ValuationsIREITNODIRECTValuationRatios) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValuationsIREITNODIRECTValuationRatios) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CyclicallyAdjustedPbRatio) {
		toSerialize["cyclically_adjusted_pb_ratio"] = o.CyclicallyAdjustedPbRatio
	}
	if !IsNil(o.CyclicallyAdjustedPriceToFcf) {
		toSerialize["cyclically_adjusted_price_to_fcf"] = o.CyclicallyAdjustedPriceToFcf
	}
	if !IsNil(o.CyclicallyAdjustedPsRatio) {
		toSerialize["cyclically_adjusted_ps_ratio"] = o.CyclicallyAdjustedPsRatio
	}
	if !IsNil(o.EarningYieldGreenblatt) {
		toSerialize["earning_yield_greenblatt"] = o.EarningYieldGreenblatt
	}
	if !IsNil(o.EnterpriseValueToEbit) {
		toSerialize["enterprise_value_to_ebit"] = o.EnterpriseValueToEbit
	}
	if !IsNil(o.EnterpriseValueToEbitda) {
		toSerialize["enterprise_value_to_ebitda"] = o.EnterpriseValueToEbitda
	}
	if !IsNil(o.EnterpriseValueToFcf) {
		toSerialize["enterprise_value_to_fcf"] = o.EnterpriseValueToFcf
	}
	if !IsNil(o.EnterpriseValueToRevenue) {
		toSerialize["enterprise_value_to_revenue"] = o.EnterpriseValueToRevenue
	}
	if !IsNil(o.FcfYield) {
		toSerialize["fcf_yield"] = o.FcfYield
	}
	if !IsNil(o.PbRatio) {
		toSerialize["pb_ratio"] = o.PbRatio
	}
	if !IsNil(o.PeRatio) {
		toSerialize["pe_ratio"] = o.PeRatio
	}
	if !IsNil(o.PegRatio) {
		toSerialize["peg_ratio"] = o.PegRatio
	}
	if !IsNil(o.Penri) {
		toSerialize["penri"] = o.Penri
	}
	if !IsNil(o.PriceToFfo) {
		toSerialize["price_to_ffo"] = o.PriceToFfo
	}
	if !IsNil(o.PriceToFreeCashFlow) {
		toSerialize["price_to_free_cash_flow"] = o.PriceToFreeCashFlow
	}
	if !IsNil(o.PriceToOperatingCashFlow) {
		toSerialize["price_to_operating_cash_flow"] = o.PriceToOperatingCashFlow
	}
	if !IsNil(o.PriceToOwnerEarnings) {
		toSerialize["price_to_owner_earnings"] = o.PriceToOwnerEarnings
	}
	if !IsNil(o.PriceToTangibleBook) {
		toSerialize["price_to_tangible_book"] = o.PriceToTangibleBook
	}
	if !IsNil(o.PsRatio) {
		toSerialize["ps_ratio"] = o.PsRatio
	}
	if !IsNil(o.RateOfReturnValue) {
		toSerialize["rate_of_return_value"] = o.RateOfReturnValue
	}
	if !IsNil(o.ShillerPeRatio) {
		toSerialize["shiller_pe_ratio"] = o.ShillerPeRatio
	}
	if !IsNil(o.Yield) {
		toSerialize["yield"] = o.Yield
	}
	return toSerialize, nil
}

type NullableValuationsIREITNODIRECTValuationRatios struct {
	value *ValuationsIREITNODIRECTValuationRatios
	isSet bool
}

func (v NullableValuationsIREITNODIRECTValuationRatios) Get() *ValuationsIREITNODIRECTValuationRatios {
	return v.value
}

func (v *NullableValuationsIREITNODIRECTValuationRatios) Set(val *ValuationsIREITNODIRECTValuationRatios) {
	v.value = val
	v.isSet = true
}

func (v NullableValuationsIREITNODIRECTValuationRatios) IsSet() bool {
	return v.isSet
}

func (v *NullableValuationsIREITNODIRECTValuationRatios) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValuationsIREITNODIRECTValuationRatios(val *ValuationsIREITNODIRECTValuationRatios) *NullableValuationsIREITNODIRECTValuationRatios {
	return &NullableValuationsIREITNODIRECTValuationRatios{value: val, isSet: true}
}

func (v NullableValuationsIREITNODIRECTValuationRatios) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValuationsIREITNODIRECTValuationRatios) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


