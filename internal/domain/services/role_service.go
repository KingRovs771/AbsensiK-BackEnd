package services

import (
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/repository"
)

type RoleService struct {
	RoleRepository *repository.RoleRepository
}

func NewRoleService(roleRepository *repository.RoleRepository) *RoleService {
	return &RoleService{RoleRepository: roleRepository}
}

func (s *RoleService) GetAllService() map[string]interface{} {
	roles, err := s.RoleRepository.GetAllRole()
	if err != nil {
		return map[string]interface{}{
			"Status":   "Error",
			"Messsage": "Error",
			"Error":    err.Error(),
		}
	}

	if len(roles) == 0 {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Status Not Found",
		}
	}

	return map[string]interface{}{
		"Status":  "Success",
		"Messafe": "Berhasil Mendapatkan Data Role",
		"Roles":   roles,
	}
}

func (s *RoleService) CreateService(role *models.Ak_Roles) map[string]interface{} {

	if role.RoleId == "" || role.NameRole == "" || role.Description == "" {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Silakan Lengkapi Data Yang Harus di isi",
		}
	}

	if err := s.RoleRepository.CreateRole(role); err != nil {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Gagal Membuat Role Baru",
		}
	}
	return map[string]interface{}{
		"Status":  "Success",
		"Message": "Berhasil Membuat Role Baru",
		"Role":    role,
	}

}

func (s *RoleService) GetRolesById(RoleId string) map[string]interface{} {
	roles, err := s.RoleRepository.GetRoleById(RoleId)
	if err != nil {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Pengguna Tidak Ditemukan",
			"Error":   err.Error(),
		}
	}

	return map[string]interface{}{
		"Status":  "Success",
		"Message": "Berhasil Mendapatkan Data Role",
		"Role":    roles,
	}
}

func (s *RoleService) UpdateRole(roles *models.Ak_Roles) map[string]interface{} {
	if err := s.RoleRepository.UpdateRole(roles); err != nil {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Gagal Memperbarui Role",
			"Error":   err.Error(),
		}
	}

	return map[string]interface{}{
		"Status":  "Success",
		"Message": "Berhasil Memperbarui Roles",
		"Roles":   roles,
	}
}

func (s *RoleService) DeleteRole(RoleId string) map[string]interface{} {
	if err := s.RoleRepository.DeleteRole(RoleId); err != nil {
		return map[string]interface{}{
			"Status ": "Error",
			"Message": "Role Tidak Dapat di Hapus",
			"Error":   err.Error(),
		}
	}

	return map[string]interface{}{
		"Status":  "Success",
		"Message": "Berhasil Menghapus Role",
		"Role":    RoleId,
	}

}
