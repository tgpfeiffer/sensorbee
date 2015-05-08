package tuple

import (
	"time"
)

type Timestamp time.Time

func (t Timestamp) Type() TypeID {
	return TypeTimestamp
}

func (t Timestamp) AsBool() (bool, error) {
	return false, castError(t.Type(), TypeBool)
}

func (t Timestamp) AsInt() (int64, error) {
	return 0, castError(t.Type(), TypeInt)
}

func (t Timestamp) AsFloat() (float64, error) {
	return 0, castError(t.Type(), TypeFloat)
}

func (t Timestamp) AsString() (string, error) {
	return "", castError(t.Type(), TypeString)
}

func (t Timestamp) AsBlob() ([]byte, error) {
	return nil, castError(t.Type(), TypeBlob)
}

func (t Timestamp) AsTimestamp() (time.Time, error) {
	return time.Time(t), nil
}

func (t Timestamp) AsArray() (Array, error) {
	return nil, castError(t.Type(), TypeArray)
}

func (t Timestamp) AsMap() (Map, error) {
	return nil, castError(t.Type(), TypeMap)
}

func (t Timestamp) clone() Value {
	return Timestamp(t)
}