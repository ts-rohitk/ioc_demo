package mapping

type APIResponse[T any] struct {
	QueryStatus string `json:"query_status"`
	Data        []T    `json:"data"`
}

type RawData struct {
	Id               string   `json:"id"`
	Ioc              string   `json:"ioc"`
	ThreatType       string   `json:"threat_type"`
	ThreatTypeDesc   string   `json:"threat_type_desc"`
	IocType          string   `json:"ioc_type"`
	IocTypeDesc      string   `json:"ioc_type_desc"`
	Malware          string   `json:"malware"`
	MalwarePrintable string   `json:"malware_printable"`
	MalwareAlias     string   `json:"malware_alias"`
	MalwareMalpedia  string   `json:"malware_malpedia"`
	ConfidenceLevel  int      `json:"confidence_level"`
	IsCompromised    bool     `json:"is_compromised"`
	FirstSeen        string   `json:"first_seen"` // into milliseconds
	LastSeen         *string  `json:"last_seen"`  // into milliseconds
	Reference        string   `json:"reference"`
	Reporter         string   `json:"reporter"`
	Tags             []string `json:"tags"`
}
