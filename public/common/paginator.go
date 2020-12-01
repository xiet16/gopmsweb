package common

import "math"

type Paging struct {
	Page      int64 `json:"page" form:"page"`
	PageSize  int64 `json:"pagesize" form:"pagesize"`
	Total     int64 `json:"total" form:"total"`
	PageCount int64 `json:"pagecount" form:"pagecount"`
	StartNums int64 `json:"startnums" form:"startnums"`
}

func (p *Paging) GetPages() {
	if p.Page < 1 {
		p.Page = 1
	}
	if p.PageSize < 1 {
		p.PageSize = 10
	}
	page_count := math.Ceil(float64(p.Total) / float64(p.PageSize))
	p.StartNums = p.PageSize * (p.Page - 1)
	p.PageCount = int64(page_count)
}
