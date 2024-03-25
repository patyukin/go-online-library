package converter

import (
	"github.com/patyukin/go-online-library/internal/handler/reqdto"
	"github.com/patyukin/go-online-library/internal/usecase/model"
)

func ToDirectoriesModelFromReqDTO(dto []reqdto.Directory) []model.Directory {
	var directories []model.Directory
	for _, directory := range dto {
		directories = append(directories, model.Directory{
			Name: directory.Name,
		})
	}

	return directories
}
