package service

import (
	"Blog/blog/dao"
	"Blog/blog/model"
)

type AdminService struct {}

var adminModel model.AdminModel

func (as AdminService) SelectByNameAndPassword(ad dao.Admin) []map[string]interface{} {
	return adminModel.SelectByNameAndPassword(ad)
}

func (as AdminService) SelectById(ad dao.Admin) []map[string]interface{} {
	return adminModel.SelectById(ad)
}

func (as AdminService) Add(ad dao.Admin) int64 {
	return adminModel.Add(ad)
}

func (as AdminService) ChangeStatus(ad dao.Admin) int64 {
	return adminModel.ChangeStatus(ad)
}

func (as AdminService) Delete(ad dao.Admin) int64 {
	return adminModel.Delete(ad)
}