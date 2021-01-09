package services

var(
	ItemService itemsServiceInterface = &itemsService{}
)

type itemsService struct{

}

type itemsServiceInterface interface{
	GetItem() 
	SaveIten()
}

func  (s *itemsService)GetItem()  {
	
}

func (s *itemsService)SaveIten(){

}