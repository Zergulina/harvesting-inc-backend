package dto

type CreateCustomerRequestDto struct {
	Ogrn          string  `json:"ogrn"`
	Name          string  `json:"name"`
	Logo          *string `json:"logo"`
	LogoExtension *string `json:"logo_extension"`
}

type CustomerDto struct {
	Id            uint64  `json:"id"`
	Ogrn          string  `json:"ogrn"`
	Name          string  `json:"name"`
	Logo          *string `json:"logo"`
	LogoExtension *string `json:"logo_extension"`
}

type UpdateCustomerRequestDto struct {
	Ogrn          string  `json:"ogrn"`
	Name          string  `json:"name"`
	Logo          *string `json:"logo"`
	LogoExtension *string `json:"logo_extension"`
}

type PatchCustomerRequestDto struct {
	Ogrn string `json:"ogrn"`
	Name string `json:"name"`
}
