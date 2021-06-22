package model

import ()

type Review struct {
	Id        int64  `json:"id"`
	ProductId int64  `json:"productId"`
	Comment   string `json:"comment"`
	Rating    int    `json:"rating"`
}
