package indecoder

type CurrencyCollection struct {
	Items []CurrencyItem `xml:"Valute"`
}

func (cc CurrencyCollection) Len() int {
	return len(cc.Items)
}

func (cc CurrencyCollection) Swap(i, j int) {
	cc.Items[i], cc.Items[j] = cc.Items[j], cc.Items[i]
}

func (cc CurrencyCollection) Less(i, j int) bool {
	return cc.Items[i].ConvertedAmount > cc.Items[j].ConvertedAmount
}
