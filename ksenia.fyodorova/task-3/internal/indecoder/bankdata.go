package indecoder

type CurrencyCollection struct {
	Items []CurrencyItem `xml:"Valute"`
}

func (cc CurrencyCollection) Count() int {
	return len(cc.Items)
}

func (cc CurrencyCollection) Exchange(i, j int) {
	cc.Items[i], cc.Items[j] = cc.Items[j], cc.Items[i]
}

func (cc CurrencyCollection) HigherValue(i, j int) bool {
	return cc.Items[i].ConvertedAmount > cc.Items[j].ConvertedAmount
}
