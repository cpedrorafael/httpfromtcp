package request

import (
	"errors"
	"io"
	"strings"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	b, err := io.ReadAll(reader)

	if err != nil {
		return nil, err
	}

	str := string(b)

	requestLine, lineErr := parseRequestLine(str)
	if lineErr != nil {
		return nil, lineErr
	}

	request := &Request{
		RequestLine: *requestLine,
	}
	return request, nil
}

func parseRequestLine(requestStr string) (*RequestLine, error) {
	lines := strings.Split(requestStr, "\r\n")
	if len(lines) == 0 {
		return nil, errors.New("invalid line")
	}
	requestLine := strings.Split(lines[0], " ")
	if len(requestLine) != 3 {
		return nil, errors.New("invalid request line")
	}

	method := strings.ToUpper(requestLine[0])
	requestTarget := requestLine[1]
	rawHttp := requestLine[2]
	httpVersion := strings.Replace(rawHttp, "HTTP/", "", 1)

	result := &RequestLine{
		HttpVersion:   httpVersion,
		RequestTarget: requestTarget,
		Method:        method,
	}
	return result, nil
}
