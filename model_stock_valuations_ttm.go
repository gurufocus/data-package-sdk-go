/*
Gurufocus Data Package API

API for accessing Gurufocus data packages, please go to [https://www.gurufocus.com/user/me?tab=account&subtab=api-token](https://www.gurufocus.com/user/me?tab=account&subtab=api-token) to view or generate authorization keys.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"gopkg.in/validator.v2"
)

// StockValuationsTtm - struct for StockValuationsTtm
type StockValuationsTtm struct {
	ValuationsINOREITNODIRECT *ValuationsINOREITNODIRECT
	ValuationsIREITNODIRECT *ValuationsIREITNODIRECT
	ValuationsNNOREITDIRECT *ValuationsNNOREITDIRECT
	ValuationsNNOREITNODIRECT *ValuationsNNOREITNODIRECT
	ValuationsNREITDIRECT *ValuationsNREITDIRECT
	ValuationsNREITNODIRECT *ValuationsNREITNODIRECT
}

// ValuationsINOREITNODIRECTAsStockValuationsTtm is a convenience function that returns ValuationsINOREITNODIRECT wrapped in StockValuationsTtm
func ValuationsINOREITNODIRECTAsStockValuationsTtm(v *ValuationsINOREITNODIRECT) StockValuationsTtm {
	return StockValuationsTtm{
		ValuationsINOREITNODIRECT: v,
	}
}

// ValuationsIREITNODIRECTAsStockValuationsTtm is a convenience function that returns ValuationsIREITNODIRECT wrapped in StockValuationsTtm
func ValuationsIREITNODIRECTAsStockValuationsTtm(v *ValuationsIREITNODIRECT) StockValuationsTtm {
	return StockValuationsTtm{
		ValuationsIREITNODIRECT: v,
	}
}

// ValuationsNNOREITDIRECTAsStockValuationsTtm is a convenience function that returns ValuationsNNOREITDIRECT wrapped in StockValuationsTtm
func ValuationsNNOREITDIRECTAsStockValuationsTtm(v *ValuationsNNOREITDIRECT) StockValuationsTtm {
	return StockValuationsTtm{
		ValuationsNNOREITDIRECT: v,
	}
}

// ValuationsNNOREITNODIRECTAsStockValuationsTtm is a convenience function that returns ValuationsNNOREITNODIRECT wrapped in StockValuationsTtm
func ValuationsNNOREITNODIRECTAsStockValuationsTtm(v *ValuationsNNOREITNODIRECT) StockValuationsTtm {
	return StockValuationsTtm{
		ValuationsNNOREITNODIRECT: v,
	}
}

// ValuationsNREITDIRECTAsStockValuationsTtm is a convenience function that returns ValuationsNREITDIRECT wrapped in StockValuationsTtm
func ValuationsNREITDIRECTAsStockValuationsTtm(v *ValuationsNREITDIRECT) StockValuationsTtm {
	return StockValuationsTtm{
		ValuationsNREITDIRECT: v,
	}
}

// ValuationsNREITNODIRECTAsStockValuationsTtm is a convenience function that returns ValuationsNREITNODIRECT wrapped in StockValuationsTtm
func ValuationsNREITNODIRECTAsStockValuationsTtm(v *ValuationsNREITNODIRECT) StockValuationsTtm {
	return StockValuationsTtm{
		ValuationsNREITNODIRECT: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *StockValuationsTtm) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ValuationsINOREITNODIRECT
	err = newStrictDecoder(data).Decode(&dst.ValuationsINOREITNODIRECT)
	if err == nil {
		jsonValuationsINOREITNODIRECT, _ := json.Marshal(dst.ValuationsINOREITNODIRECT)
		if string(jsonValuationsINOREITNODIRECT) == "{}" { // empty struct
			dst.ValuationsINOREITNODIRECT = nil
		} else {
			if err = validator.Validate(dst.ValuationsINOREITNODIRECT); err != nil {
				dst.ValuationsINOREITNODIRECT = nil
			} else {
				match++
			}
		}
	} else {
		dst.ValuationsINOREITNODIRECT = nil
	}

	// try to unmarshal data into ValuationsIREITNODIRECT
	err = newStrictDecoder(data).Decode(&dst.ValuationsIREITNODIRECT)
	if err == nil {
		jsonValuationsIREITNODIRECT, _ := json.Marshal(dst.ValuationsIREITNODIRECT)
		if string(jsonValuationsIREITNODIRECT) == "{}" { // empty struct
			dst.ValuationsIREITNODIRECT = nil
		} else {
			if err = validator.Validate(dst.ValuationsIREITNODIRECT); err != nil {
				dst.ValuationsIREITNODIRECT = nil
			} else {
				match++
			}
		}
	} else {
		dst.ValuationsIREITNODIRECT = nil
	}

	// try to unmarshal data into ValuationsNNOREITDIRECT
	err = newStrictDecoder(data).Decode(&dst.ValuationsNNOREITDIRECT)
	if err == nil {
		jsonValuationsNNOREITDIRECT, _ := json.Marshal(dst.ValuationsNNOREITDIRECT)
		if string(jsonValuationsNNOREITDIRECT) == "{}" { // empty struct
			dst.ValuationsNNOREITDIRECT = nil
		} else {
			if err = validator.Validate(dst.ValuationsNNOREITDIRECT); err != nil {
				dst.ValuationsNNOREITDIRECT = nil
			} else {
				match++
			}
		}
	} else {
		dst.ValuationsNNOREITDIRECT = nil
	}

	// try to unmarshal data into ValuationsNNOREITNODIRECT
	err = newStrictDecoder(data).Decode(&dst.ValuationsNNOREITNODIRECT)
	if err == nil {
		jsonValuationsNNOREITNODIRECT, _ := json.Marshal(dst.ValuationsNNOREITNODIRECT)
		if string(jsonValuationsNNOREITNODIRECT) == "{}" { // empty struct
			dst.ValuationsNNOREITNODIRECT = nil
		} else {
			if err = validator.Validate(dst.ValuationsNNOREITNODIRECT); err != nil {
				dst.ValuationsNNOREITNODIRECT = nil
			} else {
				match++
			}
		}
	} else {
		dst.ValuationsNNOREITNODIRECT = nil
	}

	// try to unmarshal data into ValuationsNREITDIRECT
	err = newStrictDecoder(data).Decode(&dst.ValuationsNREITDIRECT)
	if err == nil {
		jsonValuationsNREITDIRECT, _ := json.Marshal(dst.ValuationsNREITDIRECT)
		if string(jsonValuationsNREITDIRECT) == "{}" { // empty struct
			dst.ValuationsNREITDIRECT = nil
		} else {
			if err = validator.Validate(dst.ValuationsNREITDIRECT); err != nil {
				dst.ValuationsNREITDIRECT = nil
			} else {
				match++
			}
		}
	} else {
		dst.ValuationsNREITDIRECT = nil
	}

	// try to unmarshal data into ValuationsNREITNODIRECT
	err = newStrictDecoder(data).Decode(&dst.ValuationsNREITNODIRECT)
	if err == nil {
		jsonValuationsNREITNODIRECT, _ := json.Marshal(dst.ValuationsNREITNODIRECT)
		if string(jsonValuationsNREITNODIRECT) == "{}" { // empty struct
			dst.ValuationsNREITNODIRECT = nil
		} else {
			if err = validator.Validate(dst.ValuationsNREITNODIRECT); err != nil {
				dst.ValuationsNREITNODIRECT = nil
			} else {
				match++
			}
		}
	} else {
		dst.ValuationsNREITNODIRECT = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ValuationsINOREITNODIRECT = nil
		dst.ValuationsIREITNODIRECT = nil
		dst.ValuationsNNOREITDIRECT = nil
		dst.ValuationsNNOREITNODIRECT = nil
		dst.ValuationsNREITDIRECT = nil
		dst.ValuationsNREITNODIRECT = nil

		return fmt.Errorf("data matches more than one schema in oneOf(StockValuationsTtm)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(StockValuationsTtm)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src StockValuationsTtm) MarshalJSON() ([]byte, error) {
	if src.ValuationsINOREITNODIRECT != nil {
		return json.Marshal(&src.ValuationsINOREITNODIRECT)
	}

	if src.ValuationsIREITNODIRECT != nil {
		return json.Marshal(&src.ValuationsIREITNODIRECT)
	}

	if src.ValuationsNNOREITDIRECT != nil {
		return json.Marshal(&src.ValuationsNNOREITDIRECT)
	}

	if src.ValuationsNNOREITNODIRECT != nil {
		return json.Marshal(&src.ValuationsNNOREITNODIRECT)
	}

	if src.ValuationsNREITDIRECT != nil {
		return json.Marshal(&src.ValuationsNREITDIRECT)
	}

	if src.ValuationsNREITNODIRECT != nil {
		return json.Marshal(&src.ValuationsNREITNODIRECT)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *StockValuationsTtm) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ValuationsINOREITNODIRECT != nil {
		return obj.ValuationsINOREITNODIRECT
	}

	if obj.ValuationsIREITNODIRECT != nil {
		return obj.ValuationsIREITNODIRECT
	}

	if obj.ValuationsNNOREITDIRECT != nil {
		return obj.ValuationsNNOREITDIRECT
	}

	if obj.ValuationsNNOREITNODIRECT != nil {
		return obj.ValuationsNNOREITNODIRECT
	}

	if obj.ValuationsNREITDIRECT != nil {
		return obj.ValuationsNREITDIRECT
	}

	if obj.ValuationsNREITNODIRECT != nil {
		return obj.ValuationsNREITNODIRECT
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj StockValuationsTtm) GetActualInstanceValue() (interface{}) {
	if obj.ValuationsINOREITNODIRECT != nil {
		return *obj.ValuationsINOREITNODIRECT
	}

	if obj.ValuationsIREITNODIRECT != nil {
		return *obj.ValuationsIREITNODIRECT
	}

	if obj.ValuationsNNOREITDIRECT != nil {
		return *obj.ValuationsNNOREITDIRECT
	}

	if obj.ValuationsNNOREITNODIRECT != nil {
		return *obj.ValuationsNNOREITNODIRECT
	}

	if obj.ValuationsNREITDIRECT != nil {
		return *obj.ValuationsNREITDIRECT
	}

	if obj.ValuationsNREITNODIRECT != nil {
		return *obj.ValuationsNREITNODIRECT
	}

	// all schemas are nil
	return nil
}

type NullableStockValuationsTtm struct {
	value *StockValuationsTtm
	isSet bool
}

func (v NullableStockValuationsTtm) Get() *StockValuationsTtm {
	return v.value
}

func (v *NullableStockValuationsTtm) Set(val *StockValuationsTtm) {
	v.value = val
	v.isSet = true
}

func (v NullableStockValuationsTtm) IsSet() bool {
	return v.isSet
}

func (v *NullableStockValuationsTtm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStockValuationsTtm(val *StockValuationsTtm) *NullableStockValuationsTtm {
	return &NullableStockValuationsTtm{value: val, isSet: true}
}

func (v NullableStockValuationsTtm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStockValuationsTtm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


