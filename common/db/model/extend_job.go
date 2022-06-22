package model

// import (
// 	"encoding/json"
// 	"database/sql/driver"
// )
//
//
// type ConfType	interface{}
//
// func (c ConfType) Value() (driver.Value, error) {
// 	b, err := json.Marshal(c)
// 	return string(b), err
// }
//
// func (c *ConfType) Scan(src any) error {
// 	return json.Unmarshal(([]byte)(src.(string)), c)
// }
