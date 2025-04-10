# GuruTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CikReporting** | Pointer to **string** | Reported CIK of insider | [optional] 
**Company** | Pointer to **string** | The name of the company as identified on its SEC filings. | [optional] 
**Cost** | Pointer to **float32** |  | [optional] 
**Date** | Pointer to **string** | Insiders tab: the date of the insider transaction | [optional] 
**Exchange** | Pointer to **string** | The company&#39;s stock exchange. Example: NAS for Apple. | [optional] 
**FinalShare** | Pointer to **float32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Position** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**SplitFactor** | Pointer to **float32** |  | [optional] 
**Stockid** | Pointer to **string** | A unique identifier for the stock | [optional] 
**Symbol** | Pointer to **string** | The company&#39;s stock ticker symbol | [optional] 
**TransShare** | Pointer to **float32** |  | [optional] 
**Type** | Pointer to **string** | A code that determines if a security represents common stock or preferred stock | [optional] 

## Methods

### NewGuruTransaction

`func NewGuruTransaction() *GuruTransaction`

NewGuruTransaction instantiates a new GuruTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuruTransactionWithDefaults

`func NewGuruTransactionWithDefaults() *GuruTransaction`

NewGuruTransactionWithDefaults instantiates a new GuruTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCikReporting

`func (o *GuruTransaction) GetCikReporting() string`

GetCikReporting returns the CikReporting field if non-nil, zero value otherwise.

### GetCikReportingOk

`func (o *GuruTransaction) GetCikReportingOk() (*string, bool)`

GetCikReportingOk returns a tuple with the CikReporting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCikReporting

`func (o *GuruTransaction) SetCikReporting(v string)`

SetCikReporting sets CikReporting field to given value.

### HasCikReporting

`func (o *GuruTransaction) HasCikReporting() bool`

HasCikReporting returns a boolean if a field has been set.

### GetCompany

`func (o *GuruTransaction) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *GuruTransaction) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *GuruTransaction) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *GuruTransaction) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCost

`func (o *GuruTransaction) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *GuruTransaction) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *GuruTransaction) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *GuruTransaction) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetDate

`func (o *GuruTransaction) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *GuruTransaction) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *GuruTransaction) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *GuruTransaction) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetExchange

`func (o *GuruTransaction) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *GuruTransaction) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *GuruTransaction) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *GuruTransaction) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetFinalShare

`func (o *GuruTransaction) GetFinalShare() float32`

GetFinalShare returns the FinalShare field if non-nil, zero value otherwise.

### GetFinalShareOk

`func (o *GuruTransaction) GetFinalShareOk() (*float32, bool)`

GetFinalShareOk returns a tuple with the FinalShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalShare

`func (o *GuruTransaction) SetFinalShare(v float32)`

SetFinalShare sets FinalShare field to given value.

### HasFinalShare

`func (o *GuruTransaction) HasFinalShare() bool`

HasFinalShare returns a boolean if a field has been set.

### GetName

`func (o *GuruTransaction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GuruTransaction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GuruTransaction) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GuruTransaction) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPosition

`func (o *GuruTransaction) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *GuruTransaction) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *GuruTransaction) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *GuruTransaction) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetPrice

`func (o *GuruTransaction) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *GuruTransaction) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *GuruTransaction) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *GuruTransaction) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetSplitFactor

`func (o *GuruTransaction) GetSplitFactor() float32`

GetSplitFactor returns the SplitFactor field if non-nil, zero value otherwise.

### GetSplitFactorOk

`func (o *GuruTransaction) GetSplitFactorOk() (*float32, bool)`

GetSplitFactorOk returns a tuple with the SplitFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitFactor

`func (o *GuruTransaction) SetSplitFactor(v float32)`

SetSplitFactor sets SplitFactor field to given value.

### HasSplitFactor

`func (o *GuruTransaction) HasSplitFactor() bool`

HasSplitFactor returns a boolean if a field has been set.

### GetStockid

`func (o *GuruTransaction) GetStockid() string`

GetStockid returns the Stockid field if non-nil, zero value otherwise.

### GetStockidOk

`func (o *GuruTransaction) GetStockidOk() (*string, bool)`

GetStockidOk returns a tuple with the Stockid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockid

`func (o *GuruTransaction) SetStockid(v string)`

SetStockid sets Stockid field to given value.

### HasStockid

`func (o *GuruTransaction) HasStockid() bool`

HasStockid returns a boolean if a field has been set.

### GetSymbol

`func (o *GuruTransaction) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GuruTransaction) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GuruTransaction) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *GuruTransaction) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTransShare

`func (o *GuruTransaction) GetTransShare() float32`

GetTransShare returns the TransShare field if non-nil, zero value otherwise.

### GetTransShareOk

`func (o *GuruTransaction) GetTransShareOk() (*float32, bool)`

GetTransShareOk returns a tuple with the TransShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransShare

`func (o *GuruTransaction) SetTransShare(v float32)`

SetTransShare sets TransShare field to given value.

### HasTransShare

`func (o *GuruTransaction) HasTransShare() bool`

HasTransShare returns a boolean if a field has been set.

### GetType

`func (o *GuruTransaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GuruTransaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GuruTransaction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GuruTransaction) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


