package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maxzycon/SIB-Golang-Final-Project-4/internal/config"
	"github.com/maxzycon/SIB-Golang-Final-Project-4/internal/domain/global/service"
	"github.com/maxzycon/SIB-Golang-Final-Project-4/pkg/constant/role"
	"github.com/maxzycon/SIB-Golang-Final-Project-4/pkg/middleware"
)

const (
	CategorysBasePath   = "categories"
	CategorysBaseIdPath = "categories/:id"

	TasksBasePath                 = "tasks"
	TasksBaseIdPath               = "tasks/:id"
	TasksBaseUpdateStatusIdPath   = "tasks/update-status/:id"
	TasksBaseUpdateCategoryIdPath = "tasks/update-category/:id"
)

type GlobalControllerParams struct {
	V1            fiber.Router
	Conf          *config.Config
	GlobalService service.GlobalService
	Middleware    middleware.GlobalMiddleware
}
type GlobalController struct {
	v1            fiber.Router
	conf          *config.Config
	globalService service.GlobalService
	middleware    middleware.GlobalMiddleware
}

func New(params *GlobalControllerParams) *GlobalController {
	return &GlobalController{
		v1:            params.V1,
		conf:          params.Conf,
		globalService: params.GlobalService,
		middleware:    params.Middleware,
	}
}

func (pc *GlobalController) Init() {
	// ---- Categorys API
	pc.v1.Get(CategorysBasePath, pc.middleware.Protected([]uint{role.ROLE_MEMBER, role.ROLE_ADMIN}), pc.handlerGetAllCategory)
	pc.v1.Post(CategorysBasePath, pc.middleware.Protected([]uint{role.ROLE_ADMIN}), pc.handlerCreateCategory)
	pc.v1.Patch(CategorysBaseIdPath, pc.middleware.Protected([]uint{role.ROLE_ADMIN}), pc.handlerUpdateCategory)
	pc.v1.Delete(CategorysBaseIdPath, pc.middleware.Protected([]uint{role.ROLE_ADMIN}), pc.handlerDeleteCategory)

	// ---- Task API
	pc.v1.Get(TasksBasePath, pc.middleware.Protected([]uint{role.ROLE_MEMBER, role.ROLE_ADMIN}), pc.handlerGetAllTask)
	pc.v1.Post(TasksBasePath, pc.middleware.Protected([]uint{role.ROLE_MEMBER, role.ROLE_ADMIN}), pc.handlerCreateTask)
	pc.v1.Put(TasksBaseIdPath, pc.middleware.Protected([]uint{role.ROLE_MEMBER, role.ROLE_ADMIN}), pc.handlerUpdateTask)
	pc.v1.Patch(TasksBaseUpdateStatusIdPath, pc.middleware.Protected([]uint{role.ROLE_MEMBER, role.ROLE_ADMIN}), pc.handlerUpdateStatusTask)
	pc.v1.Patch(TasksBaseUpdateCategoryIdPath, pc.middleware.Protected([]uint{role.ROLE_MEMBER, role.ROLE_ADMIN}), pc.handlerUpdateCategoryTask)
	pc.v1.Delete(TasksBaseIdPath, pc.middleware.Protected([]uint{role.ROLE_MEMBER, role.ROLE_ADMIN}), pc.handlerDeleteTask)
}
