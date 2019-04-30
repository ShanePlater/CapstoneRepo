package models

import (
	"errors"
	"reflect"
	"server/orm"
	"server/types"
	"server/utils"
	"sort"
)

// GetNews return News records between startDate and endDate, where length is number of News to be returned.
func (c *Context) GetNews(data *types.GetNewsJSON) interface{} {
	// Set given time period.
	res := &copy{reserveString: []string{data.StartDate, data.EndDate}}

	// Obtain records in time period
	c.GetNewsTable().Range(res.rangeNews)

	// Sort records in order latest to oldest.
	sort.Sort(byDate(res.intf))

	// Check number of announcement message required.
	if data.Length == 0 {
		return res.intf
	}

	// Return slice as specific number of News records is required.
	return res.intf[:data.Length]
}

// rangeNews pass records in time period to copy.
func (a *copy) rangeNews(key, value interface{}) bool {
	// Check if in time period.
	if !inTimePeriod(a.reserveString[0], a.reserveString[1], reflect.ValueOf(value).FieldByName("DateTimePosted").String()) {
		return true
	}

	// Pass value to interface in copy.
	a.intf = append(a.intf, types.News{
		ID:     utils.Atoi(reflect.ValueOf(value).FieldByName("ArticleID").String()),
		Author: reflect.ValueOf(value).FieldByName("PostedBy").String(),
		Date:   reflect.ValueOf(value).FieldByName("DateTimePosted").String(),
		Title:  reflect.ValueOf(value).FieldByName("ArticleHeadline").String(),
		Text:   reflect.ValueOf(value).FieldByName("ArticleText").String(),
	})

	// Return true for keep looping through map. Otherwise, loop exit.
	return true
}

// CreateOrUpdateNews update a News record. If ID is not defined, it will create a new News record.
func (c *Context) CreateOrUpdateNews(data *types.News) error {
	r := orm.News{
		ArticleID:       utils.Itoa(data.ID),
		PostedBy:        data.Author,
		DateTimePosted:  data.Date,
		ArticleHeadline: data.Title,
		ArticleText:     data.Text,
	}

	if err := orm.CreateOrUpdateNews(&r, c.db); err != nil {
		return err
	}

	// Check if cache is enabled.
	if c.config.IsCache() {
		c.news.Store(r.ArticleID, r)
	}

	// Update data.
	*data = types.News{
		ID:     utils.Atoi(r.ArticleID),
		Author: r.PostedBy,
		Date:   r.DateTimePosted,
		Title:  r.ArticleHeadline,
		Text:   r.ArticleText,
	}

	return nil
}

// RemoveNews remove a given News record.
func (c *Context) RemoveNews(id *int) error {
	if err := orm.RemoveNews(id, c.db); err != nil {
		return err
	}

	// Check if cache is enabled.
	if !c.config.IsCache() {
		return nil
	}

	c.news.Delete(id)

	if _, ok := c.news.Load(id); ok {
		return errors.New("fail to remove news")
	}

	return nil
}
