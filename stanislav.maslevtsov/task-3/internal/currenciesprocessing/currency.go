package currenciesprocessing

type Currency struct {
	NumCode  int    `xml:"NumCode"  json:"num_code"`
	CharCode string `xml:"CharCode" json:"char_code"`
	Value    string `xml:"Value"    json:"value"`
}

type Currencies struct {
	Data []Currency `xml:"Valute"`
}
