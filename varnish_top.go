package main

type VarnishTop struct {
	Binary string
	Args   []string
}

type Gauge struct {
	Name string
	Description string
	Value float64
}

type VarnishTopMetrics struct {
	Gauges []Gauge
}

func (vt *VarnishTop) Parse(output string) (VarnishTopMetrics, error) {
	metrics := VarnishTopMetrics{}
	gauges, err := vt.parseGaugeMetrics(output)
	if err != nil {
		log.Println(err)
	}
	metrics.Gauges = []Gauge{
		{
			Name: "responses_not_ok",
			Description: "Total responses not OK",
			Value: 10,
		},
	}
	return metrics, nil
}

func (vt *VarnishTop) parseGaugeMetrics(report string) ([]Gauge, error) {
	var (
		responsesNotOk = "responses_not_ok"
	)

	metrics := map[string]float64{
		responsesNotOk: 0,
	}

	var re *regexp.Regexp = regexp.MustCompile(`^.*RespReason.*$`)
	var lines []string
	scanner := bufio.NewScanner(strings.NewReader(report))
	for scanner.Scan() {
		if lines = re.FindStringSubmatch(scanner.Text()); len(lines) > 0 {
			_ = lines[1]
		}
	}
	if err := scanner.Err(); err != nil {
		return []Gauge{}, err
	}
	var gauges []Gauge
	return gauges, nil
}
