
[![pipeline status](https://gitlab.test.igdcs.com/finops/utils/basics/reformat/badges/master/pipeline.svg)](https://gitlab.test.igdcs.com/finops/utils/basics/reformat/commits/master)
[![coverage report](https://gitlab.test.igdcs.com/finops/utils/basics/reformat/badges/master/coverage.svg)](https://gitlab.test.igdcs.com/finops/utils/basics/reformat/commits/master)

## Reformat 

Package Reformat is a generic Golang package that exposes functionality to convert an arbitrary map[string]interface{} into a native Go structure.

* Decode url.Values to a golang struct using tag names. The default setting is "json". Other tag names can be used if you setup a custom config.
* Decode map[string]interface{} to a golang struct, regardless of key names.
* Keep track of unused data. (Fields not mapped to the struct)
* Custom interface{} can be registered in the config to handle custom conversions. See *test.go for details.

## Inspired Projects

[fatih/structs](https://github.com/fatih/structs)  
[mitchellh/mapstructure](https://github.com/mitchellh/mapstructure)


It supports the following types;

	IntToFloat         float32
	IntToUint          uint   
	IntToBool          bool
	IntToString        string
	UintToInt          int
	UintToFloat        float32
	UintToBool         bool
	UintToString       string
	BoolToInt          int
	BoolToUint         uint
	BoolToFloat        float32
	BoolToString       string
	FloatToInt         int
	FloatToUint        uint
	FloatToBool        bool
	FloatToString      string
	SliceUint8ToString string
	StringToSliceUint8 []byte
	ArrayUint8ToString string
	StringToInt        int
	StringToUint       uint
	StringToBool       bool
	StringToFloat      float32
	StringToStrSlice   []string
	StringToIntSlice   []int
	StringToStrArray   [1]string
	StringToIntArray   [1]int
	SliceToMap         map[string]interface{}
	MapToSlice         []interface{}
	ArrayToMap         map[string]interface{}
	MapToArray         [1]interface{}
	Maybe some other stuff who knows


We also receive gorale.types (goracle.number) from the oracle driver package that will be supported if the "weak conversion" is switched on.
This allows for type casting and ignores explicit type conversions.


Example how to use on input url.Values() // This is a map[string][] string
Please see *_test.go for more functionality 
Most common function;

```go

package main

import (
	"fmt"
	"net/url"
	"gitlab.test.igdcs.com/finops/utils/basics/reformat.git"
)

// settlementInputStruct
type settlementInputStruct struct {
	fileID                  string `json:"file_id"`
	RecordID                string `json:"record_id,omitempty"`
	MatchStatus             string `json:"match_status,omitempty"`
	PaymentReference        string `json:"payment_reference,omitempty"`
	Moea                    string `json:"moea,omitempty"`
	AcquirerReferenceNumber string `json:"acquirer_reference_number,omitempty"`
	FromDate                string `json:"from_date,omitempty"`
	ToDate                  string `json:"to_date,omitempty"`
	Offset                  string `json:"offset"`
	Limit                   string `json:"limit"`
}

func main() {
	inputUrl := parseForm()
	// The values in the struct is not a slice, we receive the input as a slice map[string][]string we need to flatten the map before we transform it to the struct.
	input := reformat.FlattenURLValues(inputUrl)
	var inputParms settlementInputStruct
	// Setup reformat decoder. Must be a pointer to the struct, so we can change the values..
	err := reformat.Decode(input, &inputParms)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%#v\n", inputParms)
}

// this simulates the results of http.Request's ParseForm() function it returns a map[string][]string
func parseForm() url.Values {
	return url.Values{
		"file_id":                   []string{"1234"},
		"record_id":                 []string{"3"},
		"match_status":              []string{"M"},
	}
}
```

