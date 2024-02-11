package product

type DetailResponse struct {
	Id          int64  `json:"id"`
	NamaProduct string `json:"nama_product"`
	Deskripsi   string `json:"deskripsi"`
}