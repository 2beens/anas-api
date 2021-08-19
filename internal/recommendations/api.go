package recommendations

type Api interface {
	Get(userId, rId int) (*DBRecommendation, error)
	GetAll(userId int) ([]*DBRecommendation, error)
	Add(userId int, r *DBRecommendation) error
	Remove(userId, rId int) error
}
