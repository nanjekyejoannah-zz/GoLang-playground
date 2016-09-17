package db  
//note name of package should be the same as name of the folder
type item struct {
	Price float64
}


func loadItem(id int) *item{
	return &item{
		Price : 9001
	}
}