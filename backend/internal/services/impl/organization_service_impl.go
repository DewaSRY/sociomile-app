package impl

import (
	"DewaSRY/sociomile-app/internal/services"
	"DewaSRY/sociomile-app/pkg/dtos/filtersdto"
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	"DewaSRY/sociomile-app/pkg/dtos/responsedto"
	"DewaSRY/sociomile-app/pkg/lib/jwt"
	"DewaSRY/sociomile-app/pkg/models"
	"errors"

	"gorm.io/gorm"
)

type organizationServiceImpl struct {
	db *gorm.DB
}

// CreateStaff implements services.OrganizationService.
func (t *organizationServiceImpl) CreateStaff(req requestdto.RegisterRequest, user *jwt.Claims) error {
	var existingUser models.UserModel
	if err := t.db.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		return errors.New("user with this email already exists")
	}

	var ownerRole models.UserRoleModel
	if err := t.db.Where("name = ?", models.RoleOrganizationSales).
		First(&ownerRole).Error; err != nil {
		return errors.New("organization sales role not found")
	}

	userOwner := models.UserModel{
		Email:          req.Email,
		Name:           req.Name,
		Password:       req.Password,
		RoleID:         ownerRole.ID,
		OrganizationID: user.OrganizationId,
	}

	if err := t.db.Create(&userOwner).Error; err != nil {
		return errors.New("failed to create user")
	}

	return nil

}

// GetStaffList implements services.OrganizationService.
func (t *organizationServiceImpl) GetStaffList(filter filtersdto.FiltersDto, user *jwt.Claims) (*responsedto.OrganizationStaffPagination, error) {
	var staffList []models.UserModel
	var total int64
	offset := (*filter.Page - 1) * *filter.Limit
	if err := t.db.Model(&models.UserModel{}).
		Where("organization_id = ?", user.OrganizationId).
		Count(&total).Error; err != nil {
		return nil, errors.New("failed to count user staff")
	}

 if err:= t.db.Model(&models.UserModel{}).
		Where("organization_id = ?", user.OrganizationId).
		Preload("Role").
		Offset(offset).Find(&staffList).Error; err!= nil{
			return  nil, errors.New("failed to populate user")
		}
	staffListResponse := make([]responsedto.OrganizationStaffRecord, 0)

	for _, uStaff := range staffList{
		staffListResponse = append(staffListResponse, *t.mappedToOrganizationStaffRecord(&uStaff) )
	}

return  &responsedto.OrganizationStaffPagination{
	Data: staffListResponse,
	Metadata: responsedto.PaginateMetaData{
		Total: int(total),
		Limit: *filter.Limit,
		Page: *filter.Page,
	},
}, nil
}

func (t *organizationServiceImpl) mappedToOrganizationStaffRecord(user *models.UserModel) *responsedto.OrganizationStaffRecord{
	return &responsedto.OrganizationStaffRecord{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		RoleName: user.Role.Name,
	}
}




func NewOrganizationService(db *gorm.DB) services.OrganizationService {
	return &organizationServiceImpl{db: db}
}
