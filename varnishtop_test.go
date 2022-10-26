package main

import (
	"testing"

	"exporter/test/assert"
)

var report = `
 19588.00 VCL_return deliver
 10640.00 VCL_call HIT
  9810.00 VCL_use boot
  9789.00 ReqURL /
  9789.00 RespStatus 200
  9789.00 RespReason OK
  9789.00 VCL_call HASH
  9789.00 VCL_call RECV
  9789.00 VCL_return hash
  9789.00 ReqProtocol HTTP/1.1
  9789.00 RespProtocol HTTP/1.1
  9789.00 VCL_call DELIVER
  9789.00 VCL_return lookup
  9789.00 Filters  gunzip
  9789.00 ReqHeader Accept: */*
  9789.00 Begin sess 0 HTTP/1
  9789.00 ReqHeader Host: 10.88.84.160
  9789.00 RespHeader ETag: W/"6357d5f7-264"
  9789.00 ReqHeader User-Agent: curl/7.85.0
  9789.00 RespHeader Accept-Ranges: bytes
  9789.00 RespHeader Vary: Accept-Encoding
  9789.00 ReqHeader X-Forwarded-For: 10.88.84.1
  9789.00 RespHeader Content-Encoding: gzip
  9789.00 RespHeader Connection: keep-alive
  9789.00 RespHeader Content-Type: text/html
   102.00 RespStatus 404
   101.00 RespStatus 400
   100.00 RespReason Not Found
   100.00 ReqURL /login
   100.00 RespReason Not Found
   100.00 RespHeader Content-Length: 162
   100.00 RespHeader Date: Tue, 25 Oct 2022 21:34:00 GMT
    98.00 RespStatus 500
    99.00 ReqAcct 85 0 85 251 0 251
    39.00 VCL_call BACKEND_FETCH
    39.00 BereqHeader Host: 10.88.84.160
    39.00 VCL_call BACKEND_RESPONSE
    23.00 BerespHeader Server: nginx/1.18.0 (Ubuntu)
    22.00 BereqURL /
    22.00 BerespStatus 200
    22.00 BerespReason OK
    22.00 Length 384
    22.00 BereqMethod HEAD
    17.00 VCL_call MISS
    11.00 ReqURL /api/v1/someaction
    10.00 ReqMethod POST
    11.00 ReqMethod GET
    10.00 ReqURL /login.php?email=me@gmail.com
    10.00 VCL_call PASS
     9.00 ReqURL /?s=some%20%search
`

func TestParseRespStatus(t *testing.T) {
	varnishTop := NewVarnishTop("", "")
	report := varnishTop.ParseRespStatus("98.00 RespStatus 500", &RespStatusMetrics{})
	assert.Equal(t, report.Good, 0.0)
	assert.Equal(t, report.Bad, 98.0)

	varnishTop.ParseRespStatus("100.00 RespStatus 200", report)
	assert.Equal(t, report.Good, 100.0)
	assert.Equal(t, report.Bad, 98.0)

	varnishTop.ParseRespStatus("100.00 RespStatus 307", report)
	assert.Equal(t, report.Good, 100.0)
	assert.Equal(t, report.Bad, 98.0)
}

func TestParseReqMethod(t *testing.T) {
	varnishTop := NewVarnishTop("", "")
	report := varnishTop.ParseReqMethod("98.00 ReqMethod POST", &ReqMethodMetrics{})
	assert.Equal(t, report.Post, 98.0)
}

func TestParseVclCall(t *testing.T) {
	varnishTop := NewVarnishTop("", "")
	report := varnishTop.ParseVclCall("10640.00 VCL_call HIT", &VclCallMetrics{})
	assert.Equal(t, report.Hit, 10640.0)
}
