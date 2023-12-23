package controller

import (
	"gin-web/common"
	"gin-web/model"
	"gin-web/response"
	"gin-web/vo"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"strconv"
)

type IPostController interface {
	RestController
	PageList(ctx *gin.Context)
}

type PostController struct {
	DB *gorm.DB
}

func NewPostController() IPostController {
	db := common.GetDB()
	db.AutoMigrate(&model.Post{})
	return PostController{DB: db}
}

func (p PostController) Create(ctx *gin.Context) {
	var postVo vo.CreatePostRequest
	//ctx.ShouldBind(&postVo)

	if err := ctx.ShouldBind(&postVo); err != nil {
		log.Print(err.Error())
		response.Fail(ctx, nil, "数据验证错误！")
		return
	}

	// 获取登录用户
	user, _ := ctx.Get("user")

	var post = model.Post{
		UserId:     user.(model.User).ID,
		CategoryId: postVo.CategoryId,
		Title:      postVo.Title,
		HeadImg:    postVo.HeadImg,
		Content:    postVo.Content,
	}
	if err := p.DB.Create(&post).Error; err != nil {
		response.Fail(ctx, nil, "创建post失败！")
		return
	}
	response.Success(ctx, gin.H{"post": post}, "创建成功！")
}

func (p PostController) Delete(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p PostController) Update(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p PostController) Query(ctx *gin.Context) {
	// 获取path中的id
	postId := ctx.Params.ByName("id")

	var post model.Post
	// Preload
	if p.DB.Preload("Category").Where("id = ?", postId).First(&post).RecordNotFound() {
		response.Fail(ctx, nil, "查询失败！")
		return
	}

	response.Success(ctx, gin.H{"post": post}, "成功")
}

func (p PostController) PageList(ctx *gin.Context) {
	// 获取分页参数 注意获取的格式
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	// 分页
	var posts []model.Post
	p.DB.Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&posts)

	// 记录的总条数
	var total int
	p.DB.Model(model.Post{}).Count(&total)

	// 返回数据
	response.Success(ctx, gin.H{"data": posts, "total": total}, "成功")

}
