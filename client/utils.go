package client

import (
	"net/url"
	"reflect"
	"strings"
)

func StructQuery(v interface{}) url.Values {
	types := reflect.TypeOf(v)
	values := reflect.ValueOf(v)
	query := url.Values{}
	for i := 0; i < types.NumField(); i++ {
		if strings.ToLower(types.Field(i).Tag.Get("not")) != "true" {
			if values.Field(i).String() != "" {
				query.Set(types.Field(i).Tag.Get("json"), values.Field(i).String())
			}
		}
	}
	return query
}
