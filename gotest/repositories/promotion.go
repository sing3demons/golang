package repositories

type PromotionRepository interface {
	GetPromotion() (Promotion, error)
}

type Promotion struct {
	ID              uint
	PurchaseMin     int
	DiscountPercent int
}
