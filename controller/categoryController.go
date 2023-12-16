package controller

import (
	"01-quickstart/model"
	"01-quickstart/repository"
	"01-quickstart/response"
	"01-quickstart/vo"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	Repository repository.CategoryRepository
}

func NewCategoryController() ICategoryController {
	repository := repository.NewCategoryRepository()
	repository.DB.AutoMigrate(model.Category{})
	return CategoryController{Repository: repository}
}

func (c CategoryController) Create(ctx *gin.Context) {
	var categoryVo vo.CreateCategoryRequest
	//ctx.Bind(&category)
	// 使用categoryVo可根据categoryVo的标签进行数据验证
	if err := ctx.ShouldBind(&categoryVo); err != nil {
		response.Fail(ctx, nil, "数据验证错误，分类名称必填")
		return
	}

	category, err := c.Repository.Create(categoryVo.Name)

	if err != nil {
		response.Fail(ctx, nil, "创建失败！")
	}

	response.Success(ctx, gin.H{"category": category}, "创建成功！")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	if err := c.Repository.Delete(categoryId); err != nil {
		response.Fail(ctx, nil, "删除失败！")
		return
	}
	response.Success(ctx, nil, "删除成功！")
}

func (c CategoryController) Update(ctx *gin.Context) {
	var categoryVo vo.CreateCategoryRequest
	// 修改后的category
	if err := ctx.ShouldBind(&categoryVo); err != nil {
		response.Fail(ctx, nil, "数据验证错误，分类名称必填")
		return
	}

	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	category, err := c.Repository.Query(categoryId)
	if err != nil {
		response.Fail(ctx, nil, "分类不存在！")
		return
	}

	newCategory, err := c.Repository.Update(*category, categoryVo.Name)
	if err != nil {
		response.Fail(ctx, nil, "未找到对应记录")
		return
	}

	response.Success(ctx, gin.H{"category": newCategory}, "修改成功！")
}

func (c CategoryController) Query(ctx *gin.Context) {
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	category, err := c.Repository.Query(categoryId)
	if err != nil {
		response.Fail(ctx, nil, "未找到对应记录")
		return
	}

	response.Success(ctx, gin.H{"category": category}, "查询成功！")
}
