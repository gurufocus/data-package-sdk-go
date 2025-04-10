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

// checks if the ValuationsIREITNODIRECTRatios type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValuationsIREITNODIRECTRatios{}

// ValuationsIREITNODIRECTRatios struct for ValuationsIREITNODIRECTRatios
type ValuationsIREITNODIRECTRatios struct {
	CapexToOperatingCashFlow *float32 `json:"capex_to_operating_cash_flow,omitempty"`
	CapexToRevenue *float32 `json:"capex_to_revenue,omitempty"`
	// Debt to assets is a leverage ratio that defines the total amount of debt relative to assets
	DebtToAsset *float32 `json:"debt_to_asset,omitempty"`
	// The Debt/Equity (D/E) Ratio is calculated by dividing a company’s total liabilities by its shareholder equity. The ratio is used to evaluate a company's financial leverage.
	DebtToEquity *float32 `json:"debt_to_equity,omitempty"`
	DegreeOfFinancialLeverage *float32 `json:"degree_of_financial_leverage,omitempty"`
	DegreeOfOperatingLeverage *float32 `json:"degree_of_operating_leverage,omitempty"`
	// Cash dividends declared on the company's primary issue of common stock as a percent of funds from operations, on a per-share basis
	DividendToFfo *float32 `json:"dividend_to_ffo,omitempty"`
	EbitdaMargin *float32 `json:"ebitda_margin,omitempty"`
	// Effective interest rate on debt is the usage rate that a borrower actually pays on a debt. It is calculated as the positive value of interest expense divided by its average total debt.
	EffectiveInterestRate *float32 `json:"effective_interest_rate,omitempty"`
	// Equity to Asset ratio is calculated as shareholder's tangible equity divided by its total asset.
	EquityToAsset *float32 `json:"equity_to_asset,omitempty"`
	// FCF Margin is calculated as Free Cash Flow divided by total Revenue.
	FcfMargin *float32 `json:"fcf_margin,omitempty"`
	LiabilitiesToAssets *float32 `json:"liabilities_to_assets,omitempty"`
	// Net margin is calculated as Net Income divided by its Revenue.
	NetMargin *float32 `json:"net_margin,omitempty"`
	// The dividend payout ratio is the ratio of the total amount of dividends paid out to shareholders relative to the net income of the company.
	Payout *float32 `json:"payout,omitempty"`
	// Return on tangible assets is calculated as Net Income divided by its average total tangible assets. Total tangible assets equals to Total Assets minus Intangible Assets.
	ReturnOnTangibleAsset *float32 `json:"return_on_tangible_asset,omitempty"`
	// Return on tangible equity is calculated as Net Income attributable to Common Stockholders divided by its average total shareholder tangible equity.
	ReturnOnTangibleEquity *float32 `json:"return_on_tangible_equity,omitempty"`
	// Return on assets is calculated as Net Income divided by its Total Assets.
	Roa *float32 `json:"roa,omitempty"`
	Roe *float32 `json:"roe,omitempty"`
	// The return on equity adjusted to book values
	RoeAdj *float32 `json:"roe_adj,omitempty"`
	// Return on Invested Capital (ROIC) is calculated as follows:    Return on Invested Capital (ROIC) = (EBIT - Adjusted Taxes) / (Book Value of Debt + Book Value of Equity - Cash)
	Roic *float32 `json:"roic,omitempty"`
	// 1-Year Return on Invested Incremental Capital (1-Year ROIIC %) measures the change in earnings as a percentage of change in investment over 1-year.
	Roiic1y *float32 `json:"roiic_1y,omitempty"`
	// Return on Retained Earnings (RORE) is an indicator of a company's growth potential, it shows how much a company earns by reinvesting its retained earnings.
	Rore5y *float32 `json:"rore_5y,omitempty"`
	// Asset Turnover measures how quickly a company turns over its asset through sales. It is defined as     <b>Asset Turnover = Sales / {Total Assets}</b>    Companies with low profit margins tend to have high asset turnover, while those with high profit margins have low asset turnover. Companies in the retail industry tend to have a very high turnover ratio.
	Turnover *float32 `json:"turnover,omitempty"`
	// The weighted average cost of capital (WACC) is the rate that a company is expected to pay on average to all its security holders to finance its assets.
	Wacc *float32 `json:"wacc,omitempty"`
}

// NewValuationsIREITNODIRECTRatios instantiates a new ValuationsIREITNODIRECTRatios object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValuationsIREITNODIRECTRatios() *ValuationsIREITNODIRECTRatios {
	this := ValuationsIREITNODIRECTRatios{}
	return &this
}

// NewValuationsIREITNODIRECTRatiosWithDefaults instantiates a new ValuationsIREITNODIRECTRatios object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValuationsIREITNODIRECTRatiosWithDefaults() *ValuationsIREITNODIRECTRatios {
	this := ValuationsIREITNODIRECTRatios{}
	return &this
}

// GetCapexToOperatingCashFlow returns the CapexToOperatingCashFlow field value if set, zero value otherwise.
func (o *ValuationsIREITNODIRECTRatios) GetCapexToOperatingCashFlow() float32 {
	if o == nil || IsNil(o.CapexToOperatingCashFlow) {
		var ret float32
		return ret
	}
	return *o.CapexToOperatingCashFlow
}

// GetCapexToOperatingCashFlowOk returns a tuple with the CapexToOperatingCashFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsIREITNODIRECTRatios) GetCapexToOperatingCashFlowOk() (*float32, bool) {
	if o == nil || IsNil(o.CapexToOperatingCashFlow) {
		return nil, false
	}
	return o.CapexToOperatingCashFlow, true
}

// HasCapexToOperatingCashFlow returns a boolean if a field has been set.
func (o *ValuationsIREITNODIRECTRatios) HasCapexToOperatingCashFlow() bool {
	if o != nil && !IsNil(o.CapexToOperatingCashFlow) {
		return true
	}

	return false
}

// SetCapexToOperatingCashFlow gets a reference to the given float32 and assigns it to the CapexToOperatingCashFlow field.
func (o *ValuationsIREITNODIRECTRatios) SetCapexToOperatingCashFlow(v float32) {
	o.CapexToOperatingCashFlow = &v
}

// GetCapexToRevenue returns the CapexToRevenue field value if set, zero value otherwise.
func (o *ValuationsIREITNODIRECTRatios) GetCapexToRevenue() float32 {
	if o == nil || IsNil(o.CapexToRevenue) {
		var ret float32
		return ret
	}
	return *o.CapexToRevenue
}

// GetCapexToRevenueOk returns a tuple with the CapexToRevenue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsIREITNODIRECTRatios) GetCapexToRevenueOk() (*float32, bool) {
	if o == nil || IsNil(o.CapexToRevenue) {
		return nil, false
	}
	return o.CapexToRevenue, true
}

// HasCapexToRevenue returns a boolean if a field has been set.
func (o *ValuationsIREITNODIRECTRatios) HasCapexToRevenue() bool {
	if o != nil && !IsNil(o.CapexToRevenue) {
		return true
	}

	return false
}

// SetCapexToRevenue gets a reference to the given float32 and assigns it to the CapexToRevenue field.
func (o *ValuationsIREITNODIRECTRatios) SetCapexToRevenue(v float32) {
	o.CapexToRevenue = &v
}

// GetDebtToAsset returns the DebtToAsset field value if set, zero value otherwise.
func (o *ValuationsIREITNODIRECTRatios) GetDebtToAsset() float32 {
	if o == nil || IsNil(o.DebtToAsset) {
		var ret float32
		return ret
	}
	return *o.DebtToAsset
}

// GetDebtToAssetOk returns a tuple with the DebtToAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsIREITNODIRECTRatios) GetDebtToAssetOk() (*float32, bool) {
	if o == nil || IsNil(o.DebtToAsset) {
		return nil, false
	}
	return o.DebtToAsset, true
}

// HasDebtToAsset returns a boolean if a field has been set.
func (o *ValuationsIREITNODIRECTRatios) HasDebtToAsset() bool {
	if o != nil && !IsNil(o.DebtToAsset) {
		return true
	}

	return false
}

// SetDebtToAsset gets a reference to the given float32 and assigns it to the DebtToAsset field.
func (o *ValuationsIREITNODIRECTRatios) SetDebtToAsset(v float32) {
	o.DebtToAsset = &v
}

// GetDebtToEquity returns the DebtToEquity field value if set, zero value otherwise.
func (o *ValuationsIREITNODIRECTRatios) GetDebtToEquity() float32 {
	if o == nil || IsNil(o.DebtToEquity) {
		var ret float32
		return ret
	}
	return *o.DebtToEquity
}

// GetDebtToEquityOk returns a tuple with the DebtToEquity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsIREITNODIRECTRatios) GetDebtToEquityOk() (*float32, bool) {
	if o == nil || IsNil(o.DebtToEquity) {
		return nil, false
	}
	return o.DebtToEquity, true
}

// HasDebtToEquity returns a boolean if a field has been set.
func (o *ValuationsIREITNODIRECTRatios) HasDebtToEquity() bool {
	if o != nil && !IsNil(o.DebtToEquity) {
		return true
	}

	return false
}

// SetDebtToEquity gets a reference to the given float32 and assigns it to the DebtToEquity field.
func (o *ValuationsIREITNODIRECTRatios) SetDebtToEquity(v float32) {
	o.DebtToEquity = &v
}

// GetDegreeOfFinancialLeverage returns the DegreeOfFinancialLeverage field value if set, zero value otherwise.
func (o *ValuationsIREITNODIRECTRatios) GetDegreeOfFinancialLeverage() float32 {
	if o == nil || IsNil(o.DegreeOfFinancialLeverage) {
		var ret float32
		return ret
	}
	return *o.DegreeOfFinancialLeverage
}

// GetDegreeOfFinancialLeverageOk returns a tuple with the DegreeOfFinancialLeverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsIREITNODIRECTRatios) GetDegreeOfFinancialLeverageOk() (*float32, bool) {
	if o == nil || IsNil(o.DegreeOfFinancialLeverage) {
		return nil, false
	}
	return o.DegreeOfFinancialLeverage, true
}

// HasDegreeOfFinancialLeverage returns a boolean if a field has been set.
func (o *ValuationsIREITNODIRECTRatios) HasDegreeOfFinancialLeverage() bool {
	if o != nil && !IsNil(o.DegreeOfFinancialLeverage) {
		return true
	}

	return false
}

// SetDegreeOfFinancialLeverage gets a reference to the given float32 and assigns it to the DegreeOfFinancialLeverage field.
func (o *ValuationsIREITNODIRECTRatios) SetDegreeOfFinancialLeverage(v float32) {
	o.DegreeOfFinancialLeverage = &v
}

// GetDegreeOfOperatingLeverage returns the DegreeOfOperatingLeverage field value if set, zero value otherwise.
func (o *ValuationsIREITNODIRECTRatios) GetDegreeOfOperatingLeverage() float32 {
	if o == nil || IsNil(o.DegreeOfOperatingLeverage) {
		var ret float32
		return ret
	}
	return *o.DegreeOfOperatingLeverage
}

// GetDegreeOfOperatingLeverageOk returns a tuple with the DegreeOfOperatingLeverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsIREITNODIRECTRatios) GetDegreeOfOperatingLeverageOk() (*float32, bool) {
	if o == nil || IsNil(o.DegreeOfOperatingLeverage) {
		return nil, false
	}
	return o.DegreeOfOperatingLeverage, true
}

// HasDegreeOfOperatingLeverage returns a boolean if a field has been set.
func (o *ValuationsIREITNODIRECTRatios) HasDegreeOfOperatingLeverage() bool {
	if o != nil && !IsNil(o.DegreeOfOperatingLeverage) {
		return true
	}

	return false
}

// SetDegreeOfOperatingLeverage gets a reference to the given float32 and assigns it to the DegreeOfOperatingLeverage field.
func (o *ValuationsIREITNODIRECTRatios) SetDegreeOfOperatingLeverage(v float32) {
	o.DegreeOfOperatingLeverage = &v
}

// GetDividendToFfo returns the DividendToFfo field value if set, zero value otherwise.
func (o *ValuationsIREITNODIRECTRatios) GetDividendToFfo() float32 {
	if o == nil || IsNil(o.DividendToFfo) {
		var ret float32
		return ret
	}
	return *o.DividendToFfo
}

// GetDividendToFfoOk returns a tuple with the DividendToFfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsIREITNODIRECTRatios) GetDividendToFfoOk() (*float32, bool) {
	if o == nil || IsNil(o.DividendToFfo) {
		return nil, false
	}
	return o.DividendToFfo, true
}

// HasDividendToFfo returns a boolean if a field has been set.
func (o *ValuationsIREITNODIRECTRatios) HasDividendToFfo() bool {
	if o != nil && !IsNil(o.DividendToFfo) {
		return true
	}

	return false
}

// SetDividendToFfo gets a reference to the given float32 and assigns it to the DividendToFfo field.
func (o *ValuationsIREITNODIRECTRatios) SetDividendToFfo(v float32) {
	o.DividendToFfo = &v
}

// GetEbitdaMargin returns the EbitdaMargin field value if set, zero value otherwise.
func (o *ValuationsIREITNODIRECTRatios) GetEbitdaMargin() float32 {
	if o == nil || IsNil(o.EbitdaMargin) {
		var ret float32
		return ret
	}
	return *o.EbitdaMargin
}

// GetEbitdaMarginOk returns a tuple with the EbitdaMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsIREITNODIRECTRatios) GetEbitdaMarginOk() (*float32, bool) {
	if o == nil || IsNil(o.EbitdaMargin) {
		return nil, false
	}
	return o.EbitdaMargin, true
}

// HasEbitdaMargin returns a boolean if a field has been set.
func (o *ValuationsIREITNODIRECTRatios) HasEbitdaMargin() bool {
	if o != nil && !IsNil(o.EbitdaMargin) {
		return true
	}

	return false
}

// SetEbitdaMargin gets a reference to the given float32 and assigns it to the EbitdaMargin field.
func (o *ValuationsIREITNODIRECTRatios) SetEbitdaMargin(v float32) {
	o.EbitdaMargin = &v
}

// GetEffectiveInterestRate returns the EffectiveInterestRate field value if set, zero value otherwise.
func (o *ValuationsIREITNODIRECTRatios) GetEffectiveInterestRate() float32 {
	if o == nil || IsNil(o.EffectiveInterestRate) {
		var ret float32
		return ret
	}
	return *o.EffectiveInterestRate
}

// GetEffectiveInterestRateOk returns a tuple with the EffectiveInterestRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsIREITNODIRECTRatios) GetEffectiveInterestRateOk() (*float32, bool) {
	if o == nil || IsNil(o.EffectiveInterestRate) {
		return nil, false
	}
	return o.EffectiveInterestRate, true
}

// HasEffectiveInterestRate returns a boolean if a field has been set.
func (o *ValuationsIREITNODIRECTRatios) HasEffectiveInterestRate() bool {
	if o != nil && !IsNil(o.EffectiveInterestRate) {
		return true
	}

	return false
}

// SetEffectiveInterestRate gets a reference to the given float32 and assigns it to the EffectiveInterestRate field.
func (o *ValuationsIREITNODIRECTRatios) SetEffectiveInterestRate(v float32) {
	o.EffectiveInterestRate = &v
}

// GetEquityToAsset returns the EquityToAsset field value if set, zero value otherwise.
func (o *ValuationsIREITNODIRECTRatios) GetEquityToAsset() float32 {
	if o == nil || IsNil(o.EquityToAsset) {
		var ret float32
		return ret
	}
	return *o.EquityToAsset
}

// GetEquityToAssetOk returns a tuple with the EquityToAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsIREITNODIRECTRatios) GetEquityToAssetOk() (*float32, bool) {
	if o == nil || IsNil(o.EquityToAsset) {
		return nil, false
	}
	return o.EquityToAsset, true
}

// HasEquityToAsset returns a boolean if a field has been set.
func (o *ValuationsIREITNODIRECTRatios) HasEquityToAsset() bool {
	if o != nil && !IsNil(o.EquityToAsset) {
		return true
	}

	return false
}

// SetEquityToAsset gets a reference to the given float32 and assigns it to the EquityToAsset field.
func (o *ValuationsIREITNODIRECTRatios) SetEquityToAsset(v float32) {
	o.EquityToAsset = &v
}

// GetFcfMargin returns the FcfMargin field value if set, zero value otherwise.
func (o *ValuationsIREITNODIRECTRatios) GetFcfMargin() float32 {
	if o == nil || IsNil(o.FcfMargin) {
		var ret float32
		return ret
	}
	return *o.FcfMargin
}

// GetFcfMarginOk returns a tuple with the FcfMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsIREITNODIRECTRatios) GetFcfMarginOk() (*float32, bool) {
	if o == nil || IsNil(o.FcfMargin) {
		return nil, false
	}
	return o.FcfMargin, true
}

// HasFcfMargin returns a boolean if a field has been set.
func (o *ValuationsIREITNODIRECTRatios) HasFcfMargin() bool {
	if o != nil && !IsNil(o.FcfMargin) {
		return true
	}

	return false
}

// SetFcfMargin gets a reference to the given float32 and assigns it to the FcfMargin field.
func (o *ValuationsIREITNODIRECTRatios) SetFcfMargin(v float32) {
	o.FcfMargin = &v
}

// GetLiabilitiesToAssets returns the LiabilitiesToAssets field value if set, zero value otherwise.
func (o *ValuationsIREITNODIRECTRatios) GetLiabilitiesToAssets() float32 {
	if o == nil || IsNil(o.LiabilitiesToAssets) {
		var ret float32
		return ret
	}
	return *o.LiabilitiesToAssets
}

// GetLiabilitiesToAssetsOk returns a tuple with the LiabilitiesToAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsIREITNODIRECTRatios) GetLiabilitiesToAssetsOk() (*float32, bool) {
	if o == nil || IsNil(o.LiabilitiesToAssets) {
		return nil, false
	}
	return o.LiabilitiesToAssets, true
}

// HasLiabilitiesToAssets returns a boolean if a field has been set.
func (o *ValuationsIREITNODIRECTRatios) HasLiabilitiesToAssets() bool {
	if o != nil && !IsNil(o.LiabilitiesToAssets) {
		return true
	}

	return false
}

// SetLiabilitiesToAssets gets a reference to the given float32 and assigns it to the LiabilitiesToAssets field.
func (o *ValuationsIREITNODIRECTRatios) SetLiabilitiesToAssets(v float32) {
	o.LiabilitiesToAssets = &v
}

// GetNetMargin returns the NetMargin field value if set, zero value otherwise.
func (o *ValuationsIREITNODIRECTRatios) GetNetMargin() float32 {
	if o == nil || IsNil(o.NetMargin) {
		var ret float32
		return ret
	}
	return *o.NetMargin
}

// GetNetMarginOk returns a tuple with the NetMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsIREITNODIRECTRatios) GetNetMarginOk() (*float32, bool) {
	if o == nil || IsNil(o.NetMargin) {
		return nil, false
	}
	return o.NetMargin, true
}

// HasNetMargin returns a boolean if a field has been set.
func (o *ValuationsIREITNODIRECTRatios) HasNetMargin() bool {
	if o != nil && !IsNil(o.NetMargin) {
		return true
	}

	return false
}

// SetNetMargin gets a reference to the given float32 and assigns it to the NetMargin field.
func (o *ValuationsIREITNODIRECTRatios) SetNetMargin(v float32) {
	o.NetMargin = &v
}

// GetPayout returns the Payout field value if set, zero value otherwise.
func (o *ValuationsIREITNODIRECTRatios) GetPayout() float32 {
	if o == nil || IsNil(o.Payout) {
		var ret float32
		return ret
	}
	return *o.Payout
}

// GetPayoutOk returns a tuple with the Payout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsIREITNODIRECTRatios) GetPayoutOk() (*float32, bool) {
	if o == nil || IsNil(o.Payout) {
		return nil, false
	}
	return o.Payout, true
}

// HasPayout returns a boolean if a field has been set.
func (o *ValuationsIREITNODIRECTRatios) HasPayout() bool {
	if o != nil && !IsNil(o.Payout) {
		return true
	}

	return false
}

// SetPayout gets a reference to the given float32 and assigns it to the Payout field.
func (o *ValuationsIREITNODIRECTRatios) SetPayout(v float32) {
	o.Payout = &v
}

// GetReturnOnTangibleAsset returns the ReturnOnTangibleAsset field value if set, zero value otherwise.
func (o *ValuationsIREITNODIRECTRatios) GetReturnOnTangibleAsset() float32 {
	if o == nil || IsNil(o.ReturnOnTangibleAsset) {
		var ret float32
		return ret
	}
	return *o.ReturnOnTangibleAsset
}

// GetReturnOnTangibleAssetOk returns a tuple with the ReturnOnTangibleAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsIREITNODIRECTRatios) GetReturnOnTangibleAssetOk() (*float32, bool) {
	if o == nil || IsNil(o.ReturnOnTangibleAsset) {
		return nil, false
	}
	return o.ReturnOnTangibleAsset, true
}

// HasReturnOnTangibleAsset returns a boolean if a field has been set.
func (o *ValuationsIREITNODIRECTRatios) HasReturnOnTangibleAsset() bool {
	if o != nil && !IsNil(o.ReturnOnTangibleAsset) {
		return true
	}

	return false
}

// SetReturnOnTangibleAsset gets a reference to the given float32 and assigns it to the ReturnOnTangibleAsset field.
func (o *ValuationsIREITNODIRECTRatios) SetReturnOnTangibleAsset(v float32) {
	o.ReturnOnTangibleAsset = &v
}

// GetReturnOnTangibleEquity returns the ReturnOnTangibleEquity field value if set, zero value otherwise.
func (o *ValuationsIREITNODIRECTRatios) GetReturnOnTangibleEquity() float32 {
	if o == nil || IsNil(o.ReturnOnTangibleEquity) {
		var ret float32
		return ret
	}
	return *o.ReturnOnTangibleEquity
}

// GetReturnOnTangibleEquityOk returns a tuple with the ReturnOnTangibleEquity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsIREITNODIRECTRatios) GetReturnOnTangibleEquityOk() (*float32, bool) {
	if o == nil || IsNil(o.ReturnOnTangibleEquity) {
		return nil, false
	}
	return o.ReturnOnTangibleEquity, true
}

// HasReturnOnTangibleEquity returns a boolean if a field has been set.
func (o *ValuationsIREITNODIRECTRatios) HasReturnOnTangibleEquity() bool {
	if o != nil && !IsNil(o.ReturnOnTangibleEquity) {
		return true
	}

	return false
}

// SetReturnOnTangibleEquity gets a reference to the given float32 and assigns it to the ReturnOnTangibleEquity field.
func (o *ValuationsIREITNODIRECTRatios) SetReturnOnTangibleEquity(v float32) {
	o.ReturnOnTangibleEquity = &v
}

// GetRoa returns the Roa field value if set, zero value otherwise.
func (o *ValuationsIREITNODIRECTRatios) GetRoa() float32 {
	if o == nil || IsNil(o.Roa) {
		var ret float32
		return ret
	}
	return *o.Roa
}

// GetRoaOk returns a tuple with the Roa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsIREITNODIRECTRatios) GetRoaOk() (*float32, bool) {
	if o == nil || IsNil(o.Roa) {
		return nil, false
	}
	return o.Roa, true
}

// HasRoa returns a boolean if a field has been set.
func (o *ValuationsIREITNODIRECTRatios) HasRoa() bool {
	if o != nil && !IsNil(o.Roa) {
		return true
	}

	return false
}

// SetRoa gets a reference to the given float32 and assigns it to the Roa field.
func (o *ValuationsIREITNODIRECTRatios) SetRoa(v float32) {
	o.Roa = &v
}

// GetRoe returns the Roe field value if set, zero value otherwise.
func (o *ValuationsIREITNODIRECTRatios) GetRoe() float32 {
	if o == nil || IsNil(o.Roe) {
		var ret float32
		return ret
	}
	return *o.Roe
}

// GetRoeOk returns a tuple with the Roe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsIREITNODIRECTRatios) GetRoeOk() (*float32, bool) {
	if o == nil || IsNil(o.Roe) {
		return nil, false
	}
	return o.Roe, true
}

// HasRoe returns a boolean if a field has been set.
func (o *ValuationsIREITNODIRECTRatios) HasRoe() bool {
	if o != nil && !IsNil(o.Roe) {
		return true
	}

	return false
}

// SetRoe gets a reference to the given float32 and assigns it to the Roe field.
func (o *ValuationsIREITNODIRECTRatios) SetRoe(v float32) {
	o.Roe = &v
}

// GetRoeAdj returns the RoeAdj field value if set, zero value otherwise.
func (o *ValuationsIREITNODIRECTRatios) GetRoeAdj() float32 {
	if o == nil || IsNil(o.RoeAdj) {
		var ret float32
		return ret
	}
	return *o.RoeAdj
}

// GetRoeAdjOk returns a tuple with the RoeAdj field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsIREITNODIRECTRatios) GetRoeAdjOk() (*float32, bool) {
	if o == nil || IsNil(o.RoeAdj) {
		return nil, false
	}
	return o.RoeAdj, true
}

// HasRoeAdj returns a boolean if a field has been set.
func (o *ValuationsIREITNODIRECTRatios) HasRoeAdj() bool {
	if o != nil && !IsNil(o.RoeAdj) {
		return true
	}

	return false
}

// SetRoeAdj gets a reference to the given float32 and assigns it to the RoeAdj field.
func (o *ValuationsIREITNODIRECTRatios) SetRoeAdj(v float32) {
	o.RoeAdj = &v
}

// GetRoic returns the Roic field value if set, zero value otherwise.
func (o *ValuationsIREITNODIRECTRatios) GetRoic() float32 {
	if o == nil || IsNil(o.Roic) {
		var ret float32
		return ret
	}
	return *o.Roic
}

// GetRoicOk returns a tuple with the Roic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsIREITNODIRECTRatios) GetRoicOk() (*float32, bool) {
	if o == nil || IsNil(o.Roic) {
		return nil, false
	}
	return o.Roic, true
}

// HasRoic returns a boolean if a field has been set.
func (o *ValuationsIREITNODIRECTRatios) HasRoic() bool {
	if o != nil && !IsNil(o.Roic) {
		return true
	}

	return false
}

// SetRoic gets a reference to the given float32 and assigns it to the Roic field.
func (o *ValuationsIREITNODIRECTRatios) SetRoic(v float32) {
	o.Roic = &v
}

// GetRoiic1y returns the Roiic1y field value if set, zero value otherwise.
func (o *ValuationsIREITNODIRECTRatios) GetRoiic1y() float32 {
	if o == nil || IsNil(o.Roiic1y) {
		var ret float32
		return ret
	}
	return *o.Roiic1y
}

// GetRoiic1yOk returns a tuple with the Roiic1y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsIREITNODIRECTRatios) GetRoiic1yOk() (*float32, bool) {
	if o == nil || IsNil(o.Roiic1y) {
		return nil, false
	}
	return o.Roiic1y, true
}

// HasRoiic1y returns a boolean if a field has been set.
func (o *ValuationsIREITNODIRECTRatios) HasRoiic1y() bool {
	if o != nil && !IsNil(o.Roiic1y) {
		return true
	}

	return false
}

// SetRoiic1y gets a reference to the given float32 and assigns it to the Roiic1y field.
func (o *ValuationsIREITNODIRECTRatios) SetRoiic1y(v float32) {
	o.Roiic1y = &v
}

// GetRore5y returns the Rore5y field value if set, zero value otherwise.
func (o *ValuationsIREITNODIRECTRatios) GetRore5y() float32 {
	if o == nil || IsNil(o.Rore5y) {
		var ret float32
		return ret
	}
	return *o.Rore5y
}

// GetRore5yOk returns a tuple with the Rore5y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsIREITNODIRECTRatios) GetRore5yOk() (*float32, bool) {
	if o == nil || IsNil(o.Rore5y) {
		return nil, false
	}
	return o.Rore5y, true
}

// HasRore5y returns a boolean if a field has been set.
func (o *ValuationsIREITNODIRECTRatios) HasRore5y() bool {
	if o != nil && !IsNil(o.Rore5y) {
		return true
	}

	return false
}

// SetRore5y gets a reference to the given float32 and assigns it to the Rore5y field.
func (o *ValuationsIREITNODIRECTRatios) SetRore5y(v float32) {
	o.Rore5y = &v
}

// GetTurnover returns the Turnover field value if set, zero value otherwise.
func (o *ValuationsIREITNODIRECTRatios) GetTurnover() float32 {
	if o == nil || IsNil(o.Turnover) {
		var ret float32
		return ret
	}
	return *o.Turnover
}

// GetTurnoverOk returns a tuple with the Turnover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsIREITNODIRECTRatios) GetTurnoverOk() (*float32, bool) {
	if o == nil || IsNil(o.Turnover) {
		return nil, false
	}
	return o.Turnover, true
}

// HasTurnover returns a boolean if a field has been set.
func (o *ValuationsIREITNODIRECTRatios) HasTurnover() bool {
	if o != nil && !IsNil(o.Turnover) {
		return true
	}

	return false
}

// SetTurnover gets a reference to the given float32 and assigns it to the Turnover field.
func (o *ValuationsIREITNODIRECTRatios) SetTurnover(v float32) {
	o.Turnover = &v
}

// GetWacc returns the Wacc field value if set, zero value otherwise.
func (o *ValuationsIREITNODIRECTRatios) GetWacc() float32 {
	if o == nil || IsNil(o.Wacc) {
		var ret float32
		return ret
	}
	return *o.Wacc
}

// GetWaccOk returns a tuple with the Wacc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsIREITNODIRECTRatios) GetWaccOk() (*float32, bool) {
	if o == nil || IsNil(o.Wacc) {
		return nil, false
	}
	return o.Wacc, true
}

// HasWacc returns a boolean if a field has been set.
func (o *ValuationsIREITNODIRECTRatios) HasWacc() bool {
	if o != nil && !IsNil(o.Wacc) {
		return true
	}

	return false
}

// SetWacc gets a reference to the given float32 and assigns it to the Wacc field.
func (o *ValuationsIREITNODIRECTRatios) SetWacc(v float32) {
	o.Wacc = &v
}

func (o ValuationsIREITNODIRECTRatios) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValuationsIREITNODIRECTRatios) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CapexToOperatingCashFlow) {
		toSerialize["capex_to_operating_cash_flow"] = o.CapexToOperatingCashFlow
	}
	if !IsNil(o.CapexToRevenue) {
		toSerialize["capex_to_revenue"] = o.CapexToRevenue
	}
	if !IsNil(o.DebtToAsset) {
		toSerialize["debt_to_asset"] = o.DebtToAsset
	}
	if !IsNil(o.DebtToEquity) {
		toSerialize["debt_to_equity"] = o.DebtToEquity
	}
	if !IsNil(o.DegreeOfFinancialLeverage) {
		toSerialize["degree_of_financial_leverage"] = o.DegreeOfFinancialLeverage
	}
	if !IsNil(o.DegreeOfOperatingLeverage) {
		toSerialize["degree_of_operating_leverage"] = o.DegreeOfOperatingLeverage
	}
	if !IsNil(o.DividendToFfo) {
		toSerialize["dividend_to_ffo"] = o.DividendToFfo
	}
	if !IsNil(o.EbitdaMargin) {
		toSerialize["ebitda_margin"] = o.EbitdaMargin
	}
	if !IsNil(o.EffectiveInterestRate) {
		toSerialize["effective_interest_rate"] = o.EffectiveInterestRate
	}
	if !IsNil(o.EquityToAsset) {
		toSerialize["equity_to_asset"] = o.EquityToAsset
	}
	if !IsNil(o.FcfMargin) {
		toSerialize["fcf_margin"] = o.FcfMargin
	}
	if !IsNil(o.LiabilitiesToAssets) {
		toSerialize["liabilities_to_assets"] = o.LiabilitiesToAssets
	}
	if !IsNil(o.NetMargin) {
		toSerialize["net_margin"] = o.NetMargin
	}
	if !IsNil(o.Payout) {
		toSerialize["payout"] = o.Payout
	}
	if !IsNil(o.ReturnOnTangibleAsset) {
		toSerialize["return_on_tangible_asset"] = o.ReturnOnTangibleAsset
	}
	if !IsNil(o.ReturnOnTangibleEquity) {
		toSerialize["return_on_tangible_equity"] = o.ReturnOnTangibleEquity
	}
	if !IsNil(o.Roa) {
		toSerialize["roa"] = o.Roa
	}
	if !IsNil(o.Roe) {
		toSerialize["roe"] = o.Roe
	}
	if !IsNil(o.RoeAdj) {
		toSerialize["roe_adj"] = o.RoeAdj
	}
	if !IsNil(o.Roic) {
		toSerialize["roic"] = o.Roic
	}
	if !IsNil(o.Roiic1y) {
		toSerialize["roiic_1y"] = o.Roiic1y
	}
	if !IsNil(o.Rore5y) {
		toSerialize["rore_5y"] = o.Rore5y
	}
	if !IsNil(o.Turnover) {
		toSerialize["turnover"] = o.Turnover
	}
	if !IsNil(o.Wacc) {
		toSerialize["wacc"] = o.Wacc
	}
	return toSerialize, nil
}

type NullableValuationsIREITNODIRECTRatios struct {
	value *ValuationsIREITNODIRECTRatios
	isSet bool
}

func (v NullableValuationsIREITNODIRECTRatios) Get() *ValuationsIREITNODIRECTRatios {
	return v.value
}

func (v *NullableValuationsIREITNODIRECTRatios) Set(val *ValuationsIREITNODIRECTRatios) {
	v.value = val
	v.isSet = true
}

func (v NullableValuationsIREITNODIRECTRatios) IsSet() bool {
	return v.isSet
}

func (v *NullableValuationsIREITNODIRECTRatios) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValuationsIREITNODIRECTRatios(val *ValuationsIREITNODIRECTRatios) *NullableValuationsIREITNODIRECTRatios {
	return &NullableValuationsIREITNODIRECTRatios{value: val, isSet: true}
}

func (v NullableValuationsIREITNODIRECTRatios) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValuationsIREITNODIRECTRatios) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


