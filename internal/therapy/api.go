package therapy

type Api interface {
	GetDay(userId int, id int) (*Day, error)
}
