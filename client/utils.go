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
					if types.Field(i).Tag.Get("default") != "true" {
						query.Set(types.Field(i).Tag.Get("json"), "false")
					}
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

func StructBody(v interface{}) map[string]interface{} {
	resp := make(map[string]interface{})
	types := reflect.TypeOf(v)
	values := reflect.ValueOf(v)

	for i := 0; i < types.NumField(); i++ {
		if strings.ToLower(types.Field(i).Tag.Get("body")) == "true" {
			if types.Field(i).Tag.Get("default") != "" {
				resp[types.Field(i).Tag.Get("json")] = types.Field(i).Tag.Get("default")
			}
			switch values.Field(i).Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				if values.Field(i).Int() != 0 {
					resp[types.Field(i).Tag.Get("json")] = values.Field(i).Int()
				}
			case reflect.Bool:
				if values.Field(i).Bool() != false {
					resp[types.Field(i).Tag.Get("json")] = values.Field(i).Bool()
				} else {
					if types.Field(i).Tag.Get("default") != "true" {
						resp[types.Field(i).Tag.Get("json")] = false
					}
				}
			case reflect.String:
				if values.Field(i).String() != "" {
					resp[types.Field(i).Tag.Get("json")] = values.Field(i).String()
				}
			case reflect.Interface:
				resp[types.Field(i).Tag.Get("json")] = StructBody(values.Field(i).Interface())
			case reflect.Float64, reflect.Float32:
				if values.Field(i).Float() != 0 {
					resp[types.Field(i).Tag.Get("json")] = values.Field(i).Float()
				}
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				if values.Field(i).Uint() != 0 {
					resp[types.Field(i).Tag.Get("json")] = values.Field(i).Uint()
				}
			case reflect.Struct:
				resp[types.Field(i).Tag.Get("json")] = StructBody(values.Field(i).Interface())
			default:
				resp[types.Field(i).Tag.Get("json")] = values.Field(i).String()
			}

		}
	}
	return resp
}

func StructHeader(v interface{}) []map[string]string {
	types := reflect.TypeOf(v)
	values := reflect.ValueOf(v)
	header := make([]map[string]string, 0)
	for i := 0; i < types.NumField(); i++ {
		if strings.ToLower(types.Field(i).Tag.Get("header")) == "true" {
			if values.Field(i).String() != "" {
				header = append(header, map[string]string{types.Field(i).Tag.Get("json"): values.Field(i).String()})
			} else {
				if types.Field(i).Tag.Get("default") != "" {
					header = append(header, map[string]string{types.Field(i).Tag.Get("json"): types.Field(i).Tag.Get("default")})

				}
			}
		}
	}
	return header
}
