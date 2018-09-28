package toolBox

import "secretBox/models"

func GetRes() models.Res {
	var res models.Res
	res.Code = -1
	res.Info = "error"
	return res
}