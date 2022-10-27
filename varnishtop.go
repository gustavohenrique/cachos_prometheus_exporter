package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type VarnishTopMetrics struct {
	RespStatus *RespStatusMetrics
	ReqMethod  *ReqMethodMetrics
	VclCall    *VclCallMetrics
}

type VarnishTop struct {
	Binary string
	Args   []string
}

func NewVarnishTop(binary, arg string) VarnishTop {
	return VarnishTop{
		Binary: binary,
		Args:   []string{arg},
	}
}

func (vt *VarnishTop) Parse(report string) (VarnishTopMetrics, error) {
	respStatusMetrics := &RespStatusMetrics{}
	reqMethodMetrics := &ReqMethodMetrics{}
	vclCallMetrics := &VclCallMetrics{}
	scanner := bufio.NewScanner(strings.NewReader(report))
	for scanner.Scan() {
		line := scanner.Text()
		respStatusMetrics = vt.ParseRespStatus(line, respStatusMetrics)
		reqMethodMetrics = vt.ParseReqMethod(line, reqMethodMetrics)
		vclCallMetrics = vt.ParseVclCall(line, vclCallMetrics)
	}
	if err := scanner.Err(); err != nil {
		return VarnishTopMetrics{}, err
	}
	metrics := VarnishTopMetrics{
		RespStatus: respStatusMetrics,
		ReqMethod:  reqMethodMetrics,
		VclCall:    vclCallMetrics,
	}
	return metrics, nil
}

func (vt *VarnishTop) ParseRespStatus(line string, report *RespStatusMetrics) *RespStatusMetrics {
	num, text, err := split(line, "RespStatus")
	if err != nil {
		return report
	}
	quantity := toFloat(num)
	statusCode := toInt(text)
	if statusCode >= 400 {
		report.Bad = quantity
	}
	if statusCode >= 200 && statusCode < 300 {
		report.Good = quantity
	}
	return report
}

func (vt *VarnishTop) ParseReqMethod(line string, report *ReqMethodMetrics) *ReqMethodMetrics {
	num, text, err := split(line, "ReqMethod")
	if err != nil {
		return report
	}
	quantity := toFloat(num)
	method := toUpper(trim(text))
	switch method {
	case http.MethodGet:
		report.Get = quantity
	case http.MethodPost:
		report.Post = quantity
	case http.MethodPut:
		report.Put = quantity
	case http.MethodDelete:
		report.Delete = quantity
	case http.MethodHead:
		report.Head = quantity
	}
	return report
}

func (vt *VarnishTop) ParseVclCall(line string, report *VclCallMetrics) *VclCallMetrics {
	num, text, err := split(line, "VCL_call")
	if err != nil {
		return report
	}
	quantity := toFloat(num)
	call := toUpper(trim(text))
	switch call {
	case "HIT":
		report.Hit = quantity
	case "MISS":
		report.Miss = quantity
	case "PASS":
		report.Pass = quantity
	case "HASH":
		report.Hash = quantity
	case "RECV":
		report.Recv = quantity
	case "DELIVER":
		report.Deliver = quantity
	case "SYNTH":
		report.Synth = quantity
	case "PIPE":
		report.Pipe = quantity
	}
	return report
}

func toInt(s string) int {
	val, err := strconv.Atoi(trim(s))
	if err != nil {
		log.Println("cannot convert string", s, "to int")
		return 0
	}
	return val
}

func toFloat(s string) float64 {
	val, err := strconv.ParseFloat(trim(s), 64)
	if err != nil {
		log.Println("cannot convert string", s, "to float64")
		return 0
	}
	return val
}

func toUpper(s string) string {
	return strings.ToUpper(s)
}

func trim(s string) string {
	return strings.TrimSpace(s)
}

func split(line, tag string) (string, string, error) {
	splitted := strings.Split(line, tag)
	if len(splitted) != 2 {
		return "", "", fmt.Errorf("line doesnt contains the tag %s", tag)
	}
	return splitted[0], splitted[1], nil
}
