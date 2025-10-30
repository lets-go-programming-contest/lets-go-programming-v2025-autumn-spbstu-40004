package codingProcessor

type PathHolder struct {
	InPath  string `yaml:"input-file"`
	OutPath string `yaml:"output-file"`
}

type Currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

type xmlCurrency struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}
