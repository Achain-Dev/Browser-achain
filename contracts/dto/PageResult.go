package dto

import "errors"

type PageResult struct {

	TotalRecords	int

	CurrentPage		int

	TotalPage		int

	PageSize		int

	DataList		[]interface{}
}

func (p PageResult) New(currentPage int, pageSize int, list []interface{}, totalRecords int) {
	p.CurrentPage = currentPage
	p.PageSize = pageSize
	p.DataList = list
	p.TotalRecords = totalRecords
	p.TotalPage = (p.TotalRecords - 1) / p.PageSize + 1
}

func (p PageResult) NewPageResult(list []interface{}, totalRecords int) {
	p.DataList = list
	p.TotalRecords = totalRecords
}

func (p PageResult) SetPageResult(totalRecords int, dataList []interface{}) {
	if totalRecords < 0 {
		panic(errors.New("the total number of records cannot be less than 0"))
		return
	}
	p.TotalRecords = totalRecords
	p.DataList = dataList
}

func (p PageResult) SetTotalRecords(totalRecords int) {
	p.TotalRecords = totalRecords
	p.TotalPage = (p.TotalRecords - 1) / p.PageSize + 1
}
