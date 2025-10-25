package indecoder

type BankData struct {
	ValCurs []bankValute `json:"valute" xml:"Valute"`
}

func (data BankData) Len() int {
	return len(data.ValCurs)
}

func (data BankData) Swap(index1, index2 int) {
	data.ValCurs[index1], data.ValCurs[index2] = data.ValCurs[index2], data.ValCurs[index1]
}

func (data BankData) Less(index1, index2 int) bool {
	return data.ValCurs[index1].FloatValue > data.ValCurs[index2].FloatValue
}
