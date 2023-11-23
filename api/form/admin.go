package form

type AddCompanyForm struct {
	Name           string `form:"name" json:"name" binding:"required"`
	Address        string `form:"address" json:"address" binding:"required"`
	Introduction   string `form:"introduction" json:"introduction" binding:"required"`
	OfficialMobile string `form:"officialMobile" json:"officialMobile" binding:"required,mobile"`
	OfficialSite   string `form:"officialSite" json:"officialSite" binding:"required,url"`
	CompanyType    string `form:"companyType" json:"companyType" binding:"required"`
	Picture        string `form:"picture" json:"picture" binding:"required,url"`
}

type UpdateCompanyForm struct {
	Id             int    `form:"cid" json:"cid" binding:"required"`
	Name           string `form:"name" json:"name" binding:"required"`
	Address        string `form:"address" json:"address" binding:"required"`
	Introduction   string `form:"introduction" json:"introduction" binding:"required"`
	OfficialMobile string `form:"officialMobile" json:"officialMobile" binding:"required,mobile"`
	OfficialSite   string `form:"officialSite" json:"officialSite" binding:"required,url"`
	CompanyType    string `form:"companyType" json:"companyType" binding:"required"`
	Picture        string `form:"picture" json:"picture" binding:"required,url"`
}

type AddRoomForm struct {
	Cid      int      `form:"cid" json:"cid" binding:"required"`
	Name     string   `form:"name" json:"name" binding:"required"`
	Capacity int      `form:"capacity" json:"capacity" binding:"required"`
	Location string   `form:"location" json:"location" binding:"required"`
	Facility string   `form:"facility" json:"facility" binding:"required"`
	Photo    []string `form:"photo" json:"photo" binding:"required,dive"`
}

type UpdateRoomForm struct {
	Id       int      `form:"rid" json:"rid" binding:"required"`
	Name     string   `form:"name" json:"name" binding:"required"`
	Capacity int      `form:"capacity" json:"capacity" binding:"required"`
	Location string   `form:"location" json:"location" binding:"required"`
	Facility string   `form:"facility" json:"facility" binding:"required"`
	Photo    []string `form:"photo" json:"photo" binding:"required,dive"`
}
