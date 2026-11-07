package tasks_transport_http

import (
	"fmt"
	"net/http"

	"github.com/popochoo/todo-list-golang/internal/core/domain"
	core_logger "github.com/popochoo/todo-list-golang/internal/core/logger"
	core_http_request "github.com/popochoo/todo-list-golang/internal/core/transport/http/request"
	core_http_response "github.com/popochoo/todo-list-golang/internal/core/transport/http/response"
	core_http_types "github.com/popochoo/todo-list-golang/internal/core/transport/http/types"
)

type PatchTaskRequest struct {
	Title       core_http_types.Nullable[string] `json:"title" swaggertype:"string" example:"Погулять с собакой"`
	Description core_http_types.Nullable[string] `json:"description" swaggertype:"string" example:"null"`
	Completed   core_http_types.Nullable[bool]   `json:"completed" swaggertype:"boolean"`
}

func (r *PatchTaskRequest) Validate() error {
	if r.Title.Set {
		if r.Title.Value == nil {
			return fmt.Errorf("'Title' can't be NULL")
		}

		titleLength := len([]rune(*r.Title.Value))
		if titleLength < 1 || titleLength > 100 {
			return fmt.Errorf("'Title' must be between 1 and 100 symbols")
		}
	}

	if r.Description.Set {
		if r.Description.Value != nil {
			descriptionLength := len([]rune(*r.Description.Value))
			if descriptionLength < 1 || descriptionLength > 1000 {
				return fmt.Errorf("'Description' must be between 1 and 1000 symbols")
			}
		}
	}

	if r.Completed.Set {
		if r.Completed.Value == nil {
			return fmt.Errorf("'Completed' can't be NULL")
		}
	}

	return nil
}

type PatchTaskResponse TaskDTOResponse

// PatchTask godoc
// @Summary Изменение задачи
// @Description Изменение информации об уже существующей в системе задачи
// @Description ### Логика обновление полей (Three-state logic):
// @Description 1. **Поле не передано**: `description` игнорируется, значение в БД не меняется
// @Description 2. **Явно передано значение**: `description`: `Утром в 06:30 выйти на прогулку с собакой` - устанавливает новый номер телефона в БД
// @Description 3. **Передан null**: `description`: `null` - очищает поле в БД (set to NULL)
// @Description Ограничение: `title` и `completed` не может быть выставлен как null
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "ID изменяемой задачи"
// @Param request body PatchTaskRequest true "PatchTask тело запроса"
// @Success 200 {object} PatchTaskResponse "Успешно измененная задача"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 404 {object} core_http_response.ErrorResponse "User not found"
// @Failure 409 {object} core_http_response.ErrorResponse "Conflict"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /tasks/{id} [patch]
func (h *TasksHTTPHandler) PatchTask(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	taskID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get taskID path value",
		)

		return
	}

	var request PatchTaskRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate HTTP request",
		)

		return
	}

	taskPatch := taskPatchFromRequest(request)

	taskDomain, err := h.tasksService.PatchTask(ctx, taskID, taskPatch)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to patch task",
		)

		return
	}

	response := PatchTaskResponse(taskDTOFromDomain(taskDomain))

	responseHandler.JSONResponse(response, http.StatusOK)
}

func taskPatchFromRequest(request PatchTaskRequest) domain.TaskPatch {
	return domain.TaskPatch{
		request.Title.ToDomain(),
		request.Description.ToDomain(),
		request.Completed.ToDomain(),
	}
}
