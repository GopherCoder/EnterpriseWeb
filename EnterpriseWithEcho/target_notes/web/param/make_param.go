package param

type ReturnParam struct {
	Return string `json:"return" validate:"eq=return_all|eq=return_list|eq=count"`
	Search string `json:"search"`
}

type ValidWithParam interface {
	Valid() error
}
