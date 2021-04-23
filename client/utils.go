package client

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

func StructQuery(v interface{}) url.Values {
	types := reflect.TypeOf(v)
	values := reflect.ValueOf(v)
	query := url.Values{}
	for i := 0; i < types.NumField(); i++ {
		if strings.ToLower(types.Field(i).Tag.Get("query")) == "true" {
			if types.Field(i).Tag.Get("default") != "" {
				query.Set(types.Field(i).Tag.Get("json"), types.Field(i).Tag.Get("default"))
			}

			switch values.Field(i).Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				if values.Field(i).Int() != 0 {
					query.Set(types.Field(i).Tag.Get("json"), strconv.Itoa(int(values.Field(i).Int())))
				}
			case reflect.Bool:
				if values.Field(i).Bool() != false {
					query.Set(types.Field(i).Tag.Get("json"), strconv.FormatBool(values.Field(i).Bool()))
				} else {
					query.Set(types.Field(i).Tag.Get("json"), strconv.FormatBool(false))
				}
			case reflect.String:
				if values.Field(i).String() != "" {
					query.Set(types.Field(i).Tag.Get("json"), values.Field(i).String())
				}
			case reflect.Interface:
				fmt.Println(values.Field(i))
				if values.Field(i).String() != "" {
					query.Set(types.Field(i).Tag.Get("json"), values.Field(i).String())
				}
			case reflect.Float64, reflect.Float32:
				if values.Field(i).Float() != 0 {
					query.Set(types.Field(i).Tag.Get("json"), strconv.FormatFloat(values.Field(i).Float(), 'g', -1, 64))
				}
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				if values.Field(i).Uint() != 0 {
					query.Set(types.Field(i).Tag.Get("json"), strconv.FormatUint(values.Field(i).Uint(), 32))
				}
			}
		}
	}
	return query
}

func StructHeader(v interface{}) map[string][]string {
	types := reflect.TypeOf(v)
	values := reflect.ValueOf(v)
	header := make(map[string][]string)
	for i := 0; i < types.NumField(); i++ {
		if strings.ToLower(types.Field(i).Tag.Get("header")) == "true" {
			if values.Field(i).String() != "" {
				header[types.Field(i).Tag.Get("json")] = []string{values.Field(i).String()}
			} else {
				if types.Field(i).Tag.Get("default") != "" {
					header[types.Field(i).Tag.Get("json")] = []string{types.Field(i).Tag.Get("default")}
				}
			}
		}
	}
	return header
}
