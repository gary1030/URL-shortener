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

func GetOriginalUrl(url_id int64) (string, error) {
	url := &model.Url{}
	err := persistence.DB.Select("original_url").Where("id = ?", url_id).First(&url).Error
	if err != nil {
        return "", err
    } else {
		return url.Original_url, nil
	}	
}