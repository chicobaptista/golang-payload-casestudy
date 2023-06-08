package changeemployee

var CHANGE_NAME = "Name"
var CHANGE_ADDRESS = "Address"
var CHANGE_CLASSIFICATION_HOURLY = "Hourly"
var CHANGE_CLASSIFICATION_SALARIED = "Salaried"
var CHANGE_CLASSIFICATION_COMISSIONED = "Commissioned"
var CHANGE_PAYMENT_HOLD = "Hold"
var CHANGE_PAYMENT_DIRECT = "Direct"
var CHANGE_PAYMENT_MAIL = "Mail"
var CHANGE_AFFILIATION_MEMBER = "Member"
var CHANGE_AFFILIATION_NO_MEMBER = "NoMember"

type ChangeEmployee struct{}

func (tx ChangeEmployee) Execute() (success bool, err error) {
	panic("Not implemented")
}
