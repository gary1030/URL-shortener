package service

import (
	"URL-shortener/src/model"
	"URL-shortener/src/persistence"
	"time"
)

func AddUrl(Original_url string, expireAt time.Time) (*model.Url, error) {
	url := &model.Url{
        Original_url: Original_url,
        Expired_date: expireAt,
	}
	err := persistence.DB.Model(&model.Url{}).Create(&url).Error
	if err != nil {
        return nil, err
    } else {
		return url, nil
	}
}

func GetUrl(url_id int64) (*model.Url, error) {
	url := &model.Url{}
	err := persistence.DB.Select("original_url", "expired_date").Where("id = ?", url_id).First(&url).Error
	if err != nil {
		return nil, err
    } else {
		return url, nil
	}	
}