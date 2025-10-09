package tempupdater

import "errors"

var (
	ErrInvalidCmpOperator = errors.New("invalid compare operator")
	ErrTempOutOfRange     = errors.New("-1")
)

type TempUpdater struct {
	minTemp uint
	maxTemp uint
}

func NewTempUpdater() *TempUpdater {
	const (
		lowerBorder = 15
		upperBorder = 30
	)

	return &TempUpdater{
		minTemp: lowerBorder,
		maxTemp: upperBorder,
	}
}

func (tempUpd *TempUpdater) Update(cmpOperator string, temp uint) error {
	switch cmpOperator {
	case ">=":
		if temp > tempUpd.maxTemp {
			tempUpd.minTemp = temp
			return ErrTempOutOfRange
		} else if temp > tempUpd.minTemp {
			tempUpd.minTemp = temp
		}
	case "<=":
		if temp < tempUpd.minTemp {
			tempUpd.maxTemp = temp
			return ErrTempOutOfRange
		} else if temp < tempUpd.maxTemp {
			tempUpd.maxTemp = temp
		}
	default:
		return ErrInvalidCmpOperator
	}

	return nil
}

func (tempUpd *TempUpdater) GetCurrentTemp() uint {
	return tempUpd.minTemp
}
