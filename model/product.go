package model

type Product struct{
	ID int `json:"id"`
	NamaProduct string `json:"nameProduct"`
	Tipe string `json:"type"`
	Stok int `json:"stock"`
} 