package requestDto

type GetByNip struct {
	Nip string `json:"nip" binding:"required"`
}
