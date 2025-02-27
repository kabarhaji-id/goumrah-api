package api

type Response struct {
	Data  any `json:"data"`
	Error any `json:"error"`
	Meta  any `json:"meta,omitempty"`
}

func ResponseData(data any, meta ...any) Response {
	var realMeta any = meta
	if metaLength := len(meta); metaLength == 1 {
		realMeta = meta[0]
	}

	return Response{
		Data:  data,
		Error: nil,
		Meta:  realMeta,
	}
}

func ResponseError(err any, meta ...any) Response {
	var realMeta any = meta
	if metaLength := len(meta); metaLength == 1 {
		realMeta = meta[0]
	}

	if errError, ok := err.(error); ok {
		err = errError.Error()
	}

	return Response{
		Data:  nil,
		Error: err,
		Meta:  realMeta,
	}
}
