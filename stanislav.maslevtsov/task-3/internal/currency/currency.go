package currency

type Currency struct {
	NumCode  int     `json:"num_code"  xml:"NumCode"`
	CharCode string  `json:"char_code" xml:"CharCode"`
	Value    float32 `json:"value"     xml:"Value"`
}

type Currencies []*Currency
