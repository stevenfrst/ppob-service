package delivery

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"ppob-service/delivery"
	"ppob-service/delivery/product/request"
	"ppob-service/delivery/product/response"
	"ppob-service/helpers/errorHelper"
	"ppob-service/usecase/product"
	"strconv"
)

type ProductDelivery struct {
	usecase product.IProductUsecase
}

func NewProductDelivery(uc product.IProductUsecase) *ProductDelivery {
	return &ProductDelivery{
		usecase: uc,
	}
}

func (p *ProductDelivery) GetTagihanPLN(c echo.Context) error {
	resp, err := p.usecase.GetTagihanPLN()
	if err != nil {
		return delivery.ErrorResponse(c, http.StatusInternalServerError, "", err)
	}
	return delivery.SuccessResponse(c, response.FromDomain(resp))
}

func (p *ProductDelivery) GetProduct(c echo.Context) error {
	idParam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return delivery.ErrorResponse(c, http.StatusInternalServerError, "Wrong Params", err)
	}
	resp, err := p.usecase.GetProduct(idParam)
	if errors.As(err, &errorHelper.ErrRecordNotFound) {
		return delivery.ErrorResponse(c, http.StatusNoContent, "", err)
	} else if err != nil {
		return delivery.ErrorResponse(c, http.StatusInternalServerError, "", err)
	}
	return delivery.SuccessResponse(c, response.FromDomainList(resp))
}

func (p *ProductDelivery) EditProduct(c echo.Context) error {
	var productReq request.EditProduct
	err := c.Bind(&productReq)
	if err != nil {
		return delivery.ErrorResponse(c, http.StatusInternalServerError, "Failed to Bind Data", err)
	}
	err = c.Validate(&productReq)
	if err != nil {
		return delivery.ErrorResponse(c, http.StatusBadRequest, "Failed, Wrong Input", err)
	}
	err = p.usecase.EditProduct(productReq.ToDomain())
	if err != nil {
		return delivery.ErrorResponse(c, http.StatusInternalServerError, "Failed to Edit Product", err)
	}
	return delivery.SuccessResponse(c, "success")
}

func (p *ProductDelivery) DeleteProduct(c echo.Context) error {
	idParam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return delivery.ErrorResponse(c, http.StatusBadRequest, "Wrong Params", err)
	}
	err = p.usecase.Delete(idParam)
	if err != nil {
		return delivery.ErrorResponse(c, http.StatusInternalServerError, "Internal error", err)
	}

	return delivery.SuccessResponse(c, "success")
}

func (p *ProductDelivery) CreateProduct(c echo.Context) error {
	var deliveryModel request.CreateProduct
	if err := c.Bind(&deliveryModel); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(&deliveryModel); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err := p.usecase.Create(deliveryModel.ToDomain())
	if errors.As(err, &errorHelper.DuplicateData) {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return delivery.SuccessResponse(c, "success")
}

func (p *ProductDelivery) GetAllProducts(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page == 0 {
		page = 1
	}
	pageSize, _ := strconv.Atoi(c.QueryParam("page_size"))
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	offset := (page - 1) * pageSize

	resp, err := p.usecase.GetAll(offset, pageSize)
	if resp[0].ID == 0 {
		return delivery.ErrorResponse(c, 204, "Nil Data", nil)
	} else if err != nil {
		return delivery.ErrorResponse(c, http.StatusInternalServerError, "error", err)
	}

	return delivery.SuccessResponse(c, response.FromDomainList(resp))
}

func (p *ProductDelivery) GetCategory(c echo.Context) error {
	resp := p.usecase.GetAllCategory()
	if resp[0].ID == 0 {
		return delivery.ErrorResponse(c, http.StatusNoContent, "nil", nil)
	}
	return delivery.SuccessResponse(c, response.FromDomainCategoryList(resp))
}

func (p *ProductDelivery) GetSubCategory(c echo.Context) error {
	resp := p.usecase.GetAllSubCategory()
	if resp[0].ID == 0 {
		return delivery.ErrorResponse(c, http.StatusNoContent, "nil", nil)
	}
	return delivery.SuccessResponse(c, response.FromDomainSubCategoryList(resp))
}

func (p *ProductDelivery) EditSubCategory(c echo.Context) error {
	var deliveryModel request.SubCategory
	err := c.Bind(&deliveryModel)
	if err != nil {
		return delivery.ErrorResponse(c, http.StatusBadRequest, "Failed to Bind Data", err)
	}
	err = p.usecase.EditSubCategory(deliveryModel.ToDomainSubCategory())
	if err != nil {
		return delivery.ErrorResponse(c, http.StatusInternalServerError, "error", err)
	}
	return delivery.SuccessResponse(c, "success")
}

func (p *ProductDelivery) CreateCategory(c echo.Context) error {
	var deliveryModel request.Category
	err := c.Bind(&deliveryModel)
	if err != nil {
		return delivery.ErrorResponse(c, http.StatusBadRequest, "Failed to Bind Data", err)
	}
	err = c.Validate(&deliveryModel)
	if err != nil {
		return delivery.ErrorResponse(c, http.StatusBadRequest, "Required Data", err)
	}
	err = p.usecase.CreateCategory(deliveryModel.ToDomainCategory())
	if errors.As(err, &errorHelper.DuplicateData) {
		return delivery.ErrorResponse(c, http.StatusBadRequest, "not found", err)
	} else if err != nil {
		return delivery.ErrorResponse(c, http.StatusInternalServerError, "failed to create category", err)
	}
	return delivery.SuccessResponse(c, "success")
}

func (p *ProductDelivery) DeleteCategory(c echo.Context) error {
	idParam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return delivery.ErrorResponse(c, http.StatusBadRequest, "Wrong Params", err)
	}
	err = p.usecase.DeleteCategory(idParam)
	if err != nil {
		return delivery.ErrorResponse(c, http.StatusInternalServerError, "Internal error", err)
	}

	return delivery.SuccessResponse(c, "success")
}

func (p *ProductDelivery) DeleteSubCategory(c echo.Context) error {
	idParam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return delivery.ErrorResponse(c, http.StatusBadRequest, "Wrong Params", err)
	}
	err = p.usecase.DeleteSubCategory(idParam)
	if err != nil {
		return delivery.ErrorResponse(c, http.StatusInternalServerError, "Internal error", err)
	}

	return delivery.SuccessResponse(c, "success")
}

func (p *ProductDelivery) CreateSubCategory(c echo.Context) error {
	var deliveryModel = request.SubCategory{}
	name := c.FormValue("name")
	tax, err := strconv.Atoi(c.FormValue("tax"))
	if err != nil {
		return delivery.ErrorResponse(c, http.StatusBadRequest, "error input", err)
	}
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	tempFile, err := ioutil.TempFile("temp", "*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()
	log.Println(tempFile.Name())
	fileBytes, err := ioutil.ReadAll(src)
	if err != nil {
		fmt.Println(err)
	}
	tempFile.Write(fileBytes)

	defer func() {
		os.Remove(tempFile.Name())
	}()

	deliveryModel.Name = name
	deliveryModel.Tax = tax
	deliveryModel.ImageURL = tempFile.Name()

	err = p.usecase.CreateSubCategory(deliveryModel.ToDomainSubCategory(), fmt.Sprintf("%v.png", name))
	if errors.As(err, &errorHelper.DuplicateData) {
		return delivery.ErrorResponse(c, http.StatusBadRequest, "", err)
	} else if err != nil {
		return delivery.ErrorResponse(c, http.StatusInternalServerError, "", err)
	}
	return delivery.SuccessResponse(c, "success")
}
