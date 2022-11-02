package handler

import (
	"clean-architecture/internal/user/mock"
	"clean-architecture/internal/user/models"
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetListUser(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	apiLogger := logger.NewApiLogger(nil)
	mockNewsUC := mock.NewMockUseCase(ctrl)
	newsHandlers := NewNewsHandlers(nil, mockNewsUC, apiLogger)

	handlerFunc := newsHandlers.Create()

	userID := uuid.New()

	news := &models.News{
		AuthorID: userID,
		Title:    "TestNewsHandlers_Create title",
		Content:  "TestNewsHandlers_Create title content some text content",
	}

	buf, err := converter.AnyToBytesBuffer(news)
	require.NoError(t, err)
	require.NotNil(t, buf)
	require.Nil(t, err)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/news/create", strings.NewReader(buf.String()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	u := &models.User{
		UserID: userID,
	}
	ctxWithValue := context.WithValue(context.Background(), utils.UserCtxKey{}, u)
	req = req.WithContext(ctxWithValue)
	e := echo.New()
	ctx := e.NewContext(req, res)
	ctxWithReqID := utils.GetRequestCtx(ctx)
	span, ctxWithTrace := opentracing.StartSpanFromContext(ctxWithReqID, "newsHandlers.Create")
	defer span.Finish()

	mockNews := &models.News{
		AuthorID: userID,
		Title:    "TestNewsHandlers_Create title",
		Content:  "TestNewsHandlers_Create title content asdasdsadsadadsad",
	}

	mockNewsUC.EXPECT().Create(ctxWithTrace, gomock.Any()).Return(mockNews, nil)

	err = handlerFunc(ctx)
	require.NoError(t, err)
}
