package repositories

type promotionRepository struct {
}

func NewPromotionRepository() PromotionRepository {
	return promotionRepository{}
}

func (promotionRepository) GetPromotion() (Promotion, error) {
	panic("")
}
