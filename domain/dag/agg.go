package flow

import "gorm.io/gorm"

type Agg struct {
	db *gorm.DB
}

func NewAgg(db *gorm.DB) *Agg  {
	agg:= &Agg{}
	if db != nil {
		agg.db = db
	}
	return agg
}
func (a *Agg)()  {

}