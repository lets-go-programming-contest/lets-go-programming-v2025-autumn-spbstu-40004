package tempupdater

import "errors"

var ErrInvalidCmpOperator = errors.New("invalid compare operator")

type TempUpdater struct {
	minTemp int
	maxTemp int
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

func (tempUpd *TempUpdater) Update(cmpOperator string, temp int) error {
	switch cmpOperator {
	case ">=":
		if temp > tempUpd.maxTemp {
			tempUpd.minTemp = -1
			tempUpd.maxTemp = -1
		} else if temp > tempUpd.minTemp {
			tempUpd.minTemp = temp
		}
	case "<=":
		if temp < tempUpd.minTemp {
			tempUpd.minTemp = -1
			tempUpd.maxTemp = -1
		} else if temp < tempUpd.maxTemp {
			tempUpd.maxTemp = temp
		}
	default:
		return ErrInvalidCmpOperator
	}

	return nil
}

func (tempUpd *TempUpdater) GetCurrentTemp() int {
	return tempUpd.minTemp
}
