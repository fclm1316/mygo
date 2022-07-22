package tool

import (
	"strconv"
)

type PageOffset struct {
	Page   int `json:"size"`
	Offset int `json:"size"`
}

func NewPageOffset() *PageOffset {

	return &PageOffset{
		Page:   15,
		Offset: 0,
	}
}

func (ps *PageOffset) Convert(page string, offset string) {
	if page == "none" {
		return
	}
	rpage, err := strconv.Atoi(page)
	if err != nil {
		ps.Page = 0
	} else {
		ps.Page = rpage
	}
	if offset == "none" {
		return
	}
	roffset, err := strconv.Atoi(offset)
	if err != nil {
		ps.Page = 0
	} else {
		ps.Page = roffset
	}

}
