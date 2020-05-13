package service

import (
	"learnIris/model"
	"time"
	"xorm.io/xorm"
)

type StatisService interface {
	// 查询某一天的用户增长数量
	GetUserDailyCount(date string) int64
	GetOrderDailyCount(date string) int64
	GetAdminDailyCount(date string) int64
}

type statisService struct {
	engine *xorm.Engine
}

func NewStatisService(db *xorm.Engine) StatisService {
	return &statisService{
		engine: db,
	}
}

// 查询某一日用户的增长数量
func (ss *statisService) GetUserDailyCount(date string) int64 {
	if date == "NaN-NaN-NaN" {
		date = time.Now().Format("2006-01-02")
	}

	startDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return 0
	}
	endDate := startDate.AddDate(0, 0, 1)
	result, err := ss.engine.Where("register_time between ? and ? and del_falg = 0 ", startDate.Format("2006-01-02"), endDate.Format("2006-01-02")).Count(model.User{})
	if err != nil {
		return 0
	}
	return result
}

func (ss *statisService) GetOrderDailyCount(date string) int64 {
	if date == "NaN-NaN-NaN" {
		date = time.Now().Format("2006-01-02")
	}

	startDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return 0
	}
	endDate := startDate.AddDate(0, 0, 1)
	result, err := ss.engine.Where("register_time between ? and ? and del_falg = 0 ", startDate.Format("2006-01-02"), endDate.Format("2006-01-02")).Count(model.UserOrder{})
	if err != nil {
		return 0
	}
	return result
}

func (ss *statisService) GetAdminDailyCount(date string) int64 {
	if date == "NaN-NaN-NaN" {
		date = time.Now().Format("2006-01-02")
	}

	startDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return 0
	}
	endDate := startDate.AddDate(0, 0, 1)
	result, err := ss.engine.Where("create_time between ? and ? ", startDate.Format("2006-01-02"), endDate.Format("2006-01-02")).Count(model.Admin{})
	if err != nil {
		return 0
	}
	return result
}
