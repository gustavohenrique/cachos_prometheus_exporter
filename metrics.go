package main

type MetricsCounter interface {
	Counters() []Counter
}

type Counter struct {
	Name        string
	Description string
	Value       float64
}

type RespStatusMetrics struct {
	Good float64
	Bad  float64
}

func (m *RespStatusMetrics) Counters() []Counter {
	return []Counter{
		{
			Name:        "resp_status_good",
			Description: "Total of good responses status code",
			Value:       m.Good,
		},
		{
			Name:        "resp_status_bad",
			Description: "Total of bad responses status code",
			Value:       m.Bad,
		},
	}
}

type ReqMethodMetrics struct {
	Post   float64
	Get    float64
	Put    float64
	Delete float64
	Head   float64
}

func (m *ReqMethodMetrics) Counters() []Counter {
	return []Counter{
		{
			Name:        "req_method_post",
			Description: "Total of POST requests",
			Value:       m.Post,
		},
		{
			Name:        "req_method_get",
			Description: "Total of GET requests",
			Value:       m.Get,
		},
		{
			Name:        "req_method_put",
			Description: "Total of PUT requests",
			Value:       m.Put,
		},
		{
			Name:        "req_method_delete",
			Description: "Total of DELETE requests",
			Value:       m.Delete,
		},
		{
			Name:        "req_method_head",
			Description: "Total of HEAD requests",
			Value:       m.Head,
		},
	}
}

type VclCallMetrics struct {
	Hit     float64
	Miss    float64
	Pass    float64
	Recv    float64
	Hash    float64
	Deliver float64
	Pipe    float64
	Synth   float64
}

func (m *VclCallMetrics) Counters() []Counter {
	return []Counter{
		{
			Name:        "vcl_call_hit",
			Description: "Total varnish HIT",
			Value:       m.Hit,
		},
		{
			Name:        "vcl_call_miss",
			Description: "Total varnish MISS",
			Value:       m.Miss,
		},
		{
			Name:        "vcl_call_pass",
			Description: "Total varnish PASS",
			Value:       m.Pass,
		},
		{
			Name:        "vcl_call_recv",
			Description: "Total varnish RECV",
			Value:       m.Recv,
		},
		{
			Name:        "vcl_call_hash",
			Description: "Total varnish HASH",
			Value:       m.Hash,
		},
		{
			Name:        "vcl_call_deliver",
			Description: "Total varnish DELIVER",
			Value:       m.Deliver,
		},
		{
			Name:        "vcl_call_pipe",
			Description: "Total varnish PIPE",
			Value:       m.Pipe,
		},
		{
			Name:        "vcl_call_synth",
			Description: "Total varnish SYNTH",
			Value:       m.Synth,
		},
	}
}
