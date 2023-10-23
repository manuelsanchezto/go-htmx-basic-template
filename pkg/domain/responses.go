package domain

import "htmx.try/m/v2/pkg/domain/dto"

type Response struct {
	Base           BaseResponse
	SectionsToEdit []string
}

type BaseResponse struct {
	Business_line_data dto.BusinessLineData
}

func NewResponse(sections []string, businessline dto.BusinessLineData) Response {
	return Response{
		Base:           BaseResponse{Business_line_data: businessline},
		SectionsToEdit: sections,
	}
}
