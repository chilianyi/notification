// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package models

import "time"

type Notification struct {
	NotificationId string    `gorm:"column:notification_id"`
	ContentType    string    `gorm:"column:content_type"`
	SentType       string    `gorm:"column:sent_type"`
	AddrsStr       string    `gorm:"column:addrs_str"`
	Title          string    `gorm:"column:title"`
	Content        string    `gorm:"column:content"`
	ShortContent   string    `gorm:"column:short_content"`
	ExporedDays    int64     `gorm:"column:expired_days"`
	Owner          string    `gorm:"column:owner"`
	Status         string    `gorm:"column:status"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
}
