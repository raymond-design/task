package controllers

// HTTP params
const (
	idParamName  = "id"
	idsParamName = "ids"
	idParam      = `{` + idParamName + `}:^[0-9]+$`
	idsParam     = `{` + idsParamName + `}:[0-9]+(,[0-9]+)*`
)

type RemindersService interface {
	creator
	//editor
	//fetcher
	//deleter
}
