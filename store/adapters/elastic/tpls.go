package elastic

import (
	ET "github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"strconv"
)

const (
	NormedLCaseKeyword = "normed_lcase_kw"
	NormedUCaseKeyword = "normed_ucase_kw"
)

var (
	FieldIntMap     = map[string]interface{}{"type": "long"}
	FieldBoolMap    = map[string]interface{}{"type": "boolean"}
	FieldDateMap    = map[string]interface{}{"type": "date"}
	FieldFloat64Map = map[string]interface{}{"type": "double"}
	FieldTextMap    = map[string]interface{}{"type": "text", "index":"false", "norms":"false"}

	FieldPriceMap = map[string]interface{}{
		"type"			: "scaled_float",
		"scaling_factor": 100,
	}

	FieldRatingMap = map[string]interface{}{
		"type"			: "scaled_float",
		"scaling_factor": 10,
	}

	FieldSimpleKeywordLC = FieldSets{
		Type		:   "keyword",
		Index		:  true,
		Norms		:  false,
		Normalizer	: NormedLCaseKeyword,
	}

	FieldSimpleKeyword = FieldSets{
		Type	: "keyword",
		Index	: true,
		Norms	: false,
	}

	FieldNonIndexKeyword = FieldSets{
		Type	: "keyword",
		Index	: false,
		Norms	: false,
	}

	// TypePrices nested object type for prices
	TypePrices = map[string]interface{}{
		"properties": map[string]ET.Property{
			"type"		:  FieldSimpleKeyword,
			"value"		: FieldPriceMap,
			"currency"	:  FieldSimpleKeyword,
			"context"	:   FieldSimpleKeyword,
		},
	}

	// TypeShop nested object type for shop
	TypeShop = map[string]interface{}{
		"properties": map[string]ET.Property{
			"id"		:   FieldSimpleKeyword,
			"name"		: FieldSimpleKeyword,
			"link"		: FieldSimpleKeyword,
			"reviews"	:  FieldIntMap,
			"rating"	:   FieldRatingMap,
		},
	}

	// TypeImages common nested type for images
	TypeImages =  map[string]interface{}{
		"properties": map[string]ET.Property{
			"url"			: FieldSimpleKeyword,
		},
	}

	TypeMapStringString = map[string]interface{}{
		"properties": map[string]string{},
	}

)

type AnalysisNormalizer struct {
	Type		string		`json:"type,omitempty"`
	Filter		[]string	`json:"filter,omitempty"`
}

type FieldSets struct {
	Index		bool		`json:"index"`
	Norms		bool		`json:"norms"`
	Type		string		`json:"type,omitempty"`
	Analyzer	string		`json:"analyzer,omitempty"`
	Normalizer	string		`json:"normalizer,omitempty"`
}

var DefaultIndexSettings = &ET.IndexSettings{
	Index	: &ET.IndexSettings{
		NumberOfShards		: strconv.Itoa(20),
		NumberOfReplicas	: strconv.Itoa(1),
		RefreshInterval		: "1s",
	},
	Analysis	: &ET.IndexSettingsAnalysis{
		Normalizer	: map[string]ET.Normalizer{
			NormedLCaseKeyword: AnalysisNormalizer{
				Type	: "custom",
				Filter	: []string{"lowercase"},
			},
			//NormedUCaseKeyword: AnalysisNormalizer{
			//	Type	: "custom",
			//	Filter	: []string{"uppercase"},
			//},
		},
	},
}
