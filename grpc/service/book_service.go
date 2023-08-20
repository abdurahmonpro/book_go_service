package service

import (
	"book/config"
	"book/genproto/book_service"
	"book/grpc/client"
	"book/models"
	"book/pkg/logger"
	"book/storage"

	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BookService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	book_service.UnimplementedBookServiceServer
}

func NewBookService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *BookService {
	return &BookService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *BookService) Create(ctx context.Context, req *book_service.CreateBook) (*book_service.BookResponse, error) {
	i.log.Info("---CreateBook------>", logger.Any("req", req))

	bookpk, err := i.strg.Book().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateBook->Book->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	respons, err := i.strg.Book().GetByPKey(ctx, bookpk)
	if err != nil {
		i.log.Error("!!!GetBookByID->Book->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	response := &book_service.BookResponse{
		Data: []*book_service.BookData{
			{
				Book:   respons,
				Status: respons.Status,
			},
		},
		IsOk:    true,
		Message: "ok",
	}

	return response, nil
}

func (i *BookService) GetByID(ctx context.Context, req *book_service.BookPK) (resp *book_service.Book, err error) {

	i.log.Info("---GetBookByID------>", logger.Any("req", req))

	resp, err = i.strg.Book().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetBookByID->Book->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *BookService) GetBookByTitle(ctx context.Context, req *book_service.BookByTitle) (*book_service.BookResponse, error) {
	i.log.Info("---GetBookByTitle------>", logger.Any("req", req))

	book, err := i.strg.Book().GetBookByTitle(ctx, req)
	if err != nil {
		i.log.Error("!!!GetBookByTitle->Book->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	bookData := &book_service.BookData{
		Book:   book,
		Status: book.Status,
	}

	response := &book_service.BookResponse{
		Data:    []*book_service.BookData{bookData},
		IsOk:    true,
		Message: "ok",
	}

	return response, nil
}

func (i *BookService) GetList(ctx context.Context, req *book_service.BookListRequest) (*book_service.BookResponse, error) {
	i.log.Info("---GetBooks------>", logger.Any("req", req))

	resp, err := i.strg.Book().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetBooks->Book->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	bookDataList := make([]*book_service.BookData, 0, len(resp.Books))
	for _, book := range resp.Books {
		bookData := &book_service.BookData{
			Book:   book,
			Status: book.Status,
		}
		bookDataList = append(bookDataList, bookData)
	}

	response := &book_service.BookResponse{
		Data:    bookDataList,
		IsOk:    true,
		Message: "ok",
	}

	return response, nil
}

func (i *BookService) Update(ctx context.Context, req *book_service.UpdateBook) (resp *book_service.Book, err error) {

	i.log.Info("---UpdateBook------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Book().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateBook--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Book().GetByPKey(ctx, &book_service.BookPK{Id: req.Id})
	if err != nil {
		i.log.Error("!!!GetBook->Book->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *BookService) UpdatePatch(ctx context.Context, req *book_service.UpdatePatchBook) (resp *book_service.BookResponse, err error) {

	i.log.Info("---UpdatePatchBook------>", logger.Any("req", req))

	updatePatchModel := models.UpdatePatchRequest{
		Id:     req.GetId(),
		Status: req.GetStatus(),
	}

	rowsAffected, err := i.strg.Book().UpdatePatch(ctx, &updatePatchModel)

	if err != nil {
		i.log.Error("!!!UpdatePatchBook--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	respons, err := i.strg.Book().GetByPKey(ctx, &book_service.BookPK{Id: req.Id})
	if err != nil {
		i.log.Error("!!!GetBookByID->Book->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp = &book_service.BookResponse{
		Data: []*book_service.BookData{
			{
				Book:   respons,
				Status: respons.Status,
			},
		},
		IsOk:    true,
		Message: "ok",
	}

	return
}

func (i *BookService) Delete(ctx context.Context, req *book_service.BookPK) (resp *book_service.BookResponse, err error) {

	i.log.Info("---DeleteBook------>", logger.Any("req", req))

	err = i.strg.Book().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteBook->Book->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	response := &book_service.BookResponse{
		Data:    []*book_service.BookData{},
		IsOk:    true,
		Message: "ok",
	}

	return response, nil
}
