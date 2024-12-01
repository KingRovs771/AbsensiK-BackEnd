package repository

type TipePotonganRepository struct{
	DB *gorm.DB
}

func NewTipePotonganRepository(db *gorm.DB) *TipePotonganRepository{
	return &TipePotonganRepository{ DB: db}
}


func (r *TipePotonganRepository) GetAllTipePotongan()([]models.Ak_TipePotongan, error){
	var tipePotongan []models.Ak_TipePotongan
	if err := r.DB.Find(&tipePotongan).Error;err !=nil{
		return nil, err
	}
	return tipePotongan, nil
}

func (r *TipePotonganRepository) CreateTipePotongan(tipePotongan *models.Ak_TipePotongan) error{
	if err := r.DB.Create(tipePotongan).Error; err !=nil{
		return err 
	}
	return nil
}

func (r *TipePotonganRepository) GetTipePotonganById(TipePotonganId int64)(*models.Ak_TipePotongan, error){
	var tipePotongan models.Ak_TipePotongan
	if err:= r.DB.First(&tipePotongan, TipePotonganId).Error; err !=nil{
		return nil, err
	}

	return &TipePotongan, nil
}

func (r *TipePotonganRepository) UpdateTipePotongan(tipePotongan *models.Ak_TipePotongan) error{
	return r.DB.Save(tipePotongan).Error
}

func (r *TipePotongan) DeleteTipePotongan(TipePotonganId int64) error{
	return r.DB.Delete(&models.Ak_TipePotongan{}, TipePotonganId).error
}