package currenciesprocessing

type Currency struct {
	NumCode  int    `json:"num_code"  xml:"NumCode"`
	CharCode string `json:"char_code" xml:"CharCode"`
	Value    string `json:"value"     xml:"Value"`
}

type Currencies struct {
	Data []Currency `xml:"Valute"`
}
