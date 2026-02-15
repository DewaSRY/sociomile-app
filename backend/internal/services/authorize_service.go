package services

type AuthorizeService interface{
	IsUserAuthorize(roleId uint, allowedRole [] string ) error
}