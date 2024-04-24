package qdrant

import (
	"github.com/qdrant/go-client/grpc"
)

func NewMatchKeyword(field, keyword string) *grpc.Condition {
	return &grpc.Condition{
		ConditionOneOf: &grpc.Condition_Field{
			Field: &grpc.FieldCondition{
				Key: field,
				Match: &grpc.Match{
					MatchValue: &grpc.Match_Keyword{Keyword: keyword},
				},
			},
		},
	}
}

func NewMatchText(field, text string) *grpc.Condition {
	return &grpc.Condition{
		ConditionOneOf: &grpc.Condition_Field{
			Field: &grpc.FieldCondition{
				Key: field,
				Match: &grpc.Match{
					MatchValue: &grpc.Match_Text{Text: text},
				},
			},
		},
	}
}

func NewMatchBool(field string, value bool) *grpc.Condition {
	return &grpc.Condition{
		ConditionOneOf: &grpc.Condition_Field{
			Field: &grpc.FieldCondition{
				Key: field,
				Match: &grpc.Match{
					MatchValue: &grpc.Match_Boolean{Boolean: value},
				},
			},
		},
	}
}

func NewMatchInt(field string, value int64) *grpc.Condition {
	return &grpc.Condition{
		ConditionOneOf: &grpc.Condition_Field{
			Field: &grpc.FieldCondition{
				Key: field,
				Match: &grpc.Match{
					MatchValue: &grpc.Match_Integer{Integer: value},
				},
			},
		},
	}
}

func NewMatchKeywords(field string, keywords ...string) *grpc.Condition {
	return &grpc.Condition{
		ConditionOneOf: &grpc.Condition_Field{
			Field: &grpc.FieldCondition{
				Key: field,
				Match: &grpc.Match{
					MatchValue: &grpc.Match_Keywords{Keywords: &grpc.RepeatedStrings{
						Strings: keywords,
					}},
				},
			},
		},
	}
}

func NewMatchInts(field string, values ...int64) *grpc.Condition {
	return &grpc.Condition{
		ConditionOneOf: &grpc.Condition_Field{
			Field: &grpc.FieldCondition{
				Key: field,
				Match: &grpc.Match{
					MatchValue: &grpc.Match_Integers{Integers: &grpc.RepeatedIntegers{
						Integers: values,
					}},
				},
			},
		},
	}
}

func NewMatchExceptKeywords(field string, keywords ...string) *grpc.Condition {
	return &grpc.Condition{
		ConditionOneOf: &grpc.Condition_Field{
			Field: &grpc.FieldCondition{
				Key: field,
				Match: &grpc.Match{
					MatchValue: &grpc.Match_ExceptKeywords{ExceptKeywords: &grpc.RepeatedStrings{
						Strings: keywords,
					}},
				},
			},
		},
	}
}

func NewMatchExceptInts(field string, values ...int64) *grpc.Condition {
	return &grpc.Condition{
		ConditionOneOf: &grpc.Condition_Field{
			Field: &grpc.FieldCondition{
				Key: field,
				Match: &grpc.Match{
					MatchValue: &grpc.Match_ExceptIntegers{ExceptIntegers: &grpc.RepeatedIntegers{
						Integers: values,
					}},
				},
			},
		},
	}
}

func NewIsNull(field string) *grpc.Condition {
	return &grpc.Condition{
		ConditionOneOf: &grpc.Condition_IsNull{
			IsNull: &grpc.IsNullCondition{
				Key: field,
			},
		},
	}
}

func NewIsEmpty(field string) *grpc.Condition {
	return &grpc.Condition{
		ConditionOneOf: &grpc.Condition_IsEmpty{
			IsEmpty: &grpc.IsEmptyCondition{
				Key: field,
			},
		},
	}
}

func NewHasId(ids ...*grpc.PointId) *grpc.Condition {
	return &grpc.Condition{
		ConditionOneOf: &grpc.Condition_HasId{
			HasId: &grpc.HasIdCondition{
				HasId: ids,
			},
		},
	}
}

func NewNestedCondtion(field string, conditon *grpc.Condition) *grpc.Condition {

	filter := &grpc.Filter{
		Must: []*grpc.Condition{conditon},
	}

	return &grpc.Condition{
		ConditionOneOf: &grpc.Condition_Nested{
			Nested: &grpc.NestedCondition{
				Key:    field,
				Filter: filter,
			},
		},
	}
}

func NewNestedFilter(field string, filter *grpc.Filter) *grpc.Condition {
	return &grpc.Condition{
		ConditionOneOf: &grpc.Condition_Nested{
			Nested: &grpc.NestedCondition{
				Key:    field,
				Filter: filter,
			},
		},
	}
}

func NewFilterAsCondition(filter *grpc.Filter) *grpc.Condition {
	return &grpc.Condition{
		ConditionOneOf: &grpc.Condition_Filter{
			Filter: filter,
		},
	}
}

func NewRange(field string, rangeVal *grpc.Range) *grpc.Condition {
	return &grpc.Condition{
		ConditionOneOf: &grpc.Condition_Field{
			Field: &grpc.FieldCondition{
				Key:   field,
				Range: rangeVal,
			},
		},
	}
}

func NewGeoRadius(field string, lat, long float64, radius float32) *grpc.Condition {
	return &grpc.Condition{
		ConditionOneOf: &grpc.Condition_Field{
			Field: &grpc.FieldCondition{
				Key: field,
				GeoRadius: &grpc.GeoRadius{
					Radius: radius,
					Center: &grpc.GeoPoint{
						Lat: lat,
						Lon: long,
					},
				},
			},
		},
	}
}

func NewGeoBoundingBox(field string, topLeftLat, topLeftLon, bottomRightLat, bottomRightLon float64) *grpc.Condition {
	return &grpc.Condition{
		ConditionOneOf: &grpc.Condition_Field{
			Field: &grpc.FieldCondition{
				Key: field,
				GeoBoundingBox: &grpc.GeoBoundingBox{
					TopLeft: &grpc.GeoPoint{
						Lat: topLeftLat,
						Lon: topLeftLon,
					},
					BottomRight: &grpc.GeoPoint{
						Lat: bottomRightLat,
						Lon: bottomRightLon,
					},
				},
			},
		},
	}
}

func NewGeoPolygon(field string, exterior *grpc.GeoLineString, interior ...*grpc.GeoLineString) *grpc.Condition {

	geoPolygon := &grpc.GeoPolygon{
		Exterior:  exterior,
		Interiors: interior,
	}

	return &grpc.Condition{
		ConditionOneOf: &grpc.Condition_Field{
			Field: &grpc.FieldCondition{
				Key:        field,
				GeoPolygon: geoPolygon,
			},
		},
	}
}

func NewValuesCount(field string, valuesCount *grpc.ValuesCount) *grpc.Condition {
	return &grpc.Condition{
		ConditionOneOf: &grpc.Condition_Field{
			Field: &grpc.FieldCondition{
				Key:         field,
				ValuesCount: valuesCount,
			},
		},
	}
}

func NewDatetimeRange(field string, dateTimeRange *grpc.DatetimeRange) *grpc.Condition {
	return &grpc.Condition{
		ConditionOneOf: &grpc.Condition_Field{
			Field: &grpc.FieldCondition{
				Key:           field,
				DatetimeRange: dateTimeRange,
			},
		},
	}
}
