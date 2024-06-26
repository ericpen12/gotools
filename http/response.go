package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Parser interface {
	Obj() any
	ErrMsg() string
}

type EasyResponse interface {
	Buffer() *bytes.Buffer
	Bind(data interface{}) error
	UnescapeString() EasyResponse
	FormatJson() string
	StatusCode() int
}

func Read(resp *http.Response, format any) (EasyResponse, error) {
	buf := bytes.NewBuffer(make([]byte, 0, 1024))
	_, err := buf.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}
	newBuffer, err := getBuffer(buf, format)
	if err != nil {
		return nil, err
	}
	return &result{
		buf:        newBuffer,
		statusCode: resp.StatusCode,
	}, nil
}

func getBuffer(r *bytes.Buffer, format any) (*bytes.Buffer, error) {
	if format == nil {
		return r, nil
	}
	err := json.Unmarshal(r.Bytes(), format)
	if err != nil {
		return nil, err
	}
	p, ok := format.(Parser)
	if !ok {
		b, err := json.Marshal(format)
		return bytes.NewBuffer(b), err
	}

	if p.ErrMsg() != "" {
		return nil, fmt.Errorf(p.ErrMsg())
	}

	b, err := json.Marshal(p.Obj())
	return bytes.NewBuffer(b), err
}

type result struct {
	buf        *bytes.Buffer
	statusCode int
}

func (r *result) Buffer() *bytes.Buffer {
	return r.buf
}

func (r *result) String() string {
	return r.buf.String()
}

func (r *result) StatusCode() int {
	return r.statusCode
}

func (r *result) Bind(data interface{}) error {
	return json.Unmarshal(r.buf.Bytes(), data)
}

func (r *result) UnescapeString() EasyResponse {
	m := make(map[string]interface{})
	_ = json.Unmarshal([]byte(UnescapeString(r.buf.Bytes())), &m)
	b, _ := json.Marshal(m)
	r.buf = bytes.NewBuffer(b)
	return r
}

func (r *result) FormatJson() string {
	if !json.Valid(r.buf.Bytes()) {
		return r.buf.String()
	}
	var data any
	json.Unmarshal(r.buf.Bytes(), &data)
	b, _ := json.MarshalIndent(data, "", "\t")
	return string(b)
}

func UnescapeString(input []byte) string {
	var s string
	_ = json.Unmarshal(input, &s)
	return s
}
