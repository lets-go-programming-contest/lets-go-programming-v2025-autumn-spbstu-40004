package indecoder

type CurrencyCollection struct {
	Items []CurrencyItem `xml:"Valute"`
}

// Len возвращает количество элементов
func (cc CurrencyCollection) Len() int {
	return len(cc.Items)
}

// Swap меняет местами элементы i и j
func (cc CurrencyCollection) Swap(i, j int) {
	cc.Items[i], cc.Items[j] = cc.Items[j], cc.Items[i]
}

// Less определяет порядок сортировки (по убыванию значения)
func (cc CurrencyCollection) Less(i, j int) bool {
	return cc.Items[i].ConvertedAmount > cc.Items[j].ConvertedAmount
}
