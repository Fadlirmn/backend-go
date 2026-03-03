package service

import(
	"backend-api-belajar/repository"
	"backend-api-belajar/model"
)

type ProductService struct{
	repo repository.ProductRepository
}
func NewProductService(r repository.ProductRepository) *ProductService  {
	return &ProductService{repo:r}
}
func (s *ProductService)GetProduct()[]model.Product  {
	return s.repo.FindAllProduct()
}
func (s *ProductService)CreateProduct(product model.Product)  {
	s.repo.SaveProduct(product)
}
func (s *ProductService)UpdateProduct(id int,product model.Product)error  {
	return s.repo.UpdateProduct(id, product)
}
func (s *ProductService)DeleteProduct(id int)error {
	return s.repo.DeleteProduct(id)
}