package usecase_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	mock "github.com/yudai2929/connpass-keyword-bot-v1/mocks/repository"
	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/domain/entity"
	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/usecase"
	"go.uber.org/mock/gomock"
)

func TestSend_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	eventRepo := mock.NewMockEventRepository(ctrl)
	messageRepo := mock.NewMockMessageRepository(ctrl)
	notifiedEventRepo := mock.NewMockNotifiedEventRepository(ctrl)

	u := usecase.NewNotificationUsecase(eventRepo, messageRepo, notifiedEventRepo)

	expectedEvents := []entity.Event{
		{
			EventID:  1,
			Title:    "test1",
			EventURL: "https://test1.com",
		},
	}

	eventRepo.EXPECT().GetByKeyword([]string{"名古屋", "愛知"}).Return(expectedEvents, nil)

	notifiedEventRepo.EXPECT().FindByEventIDs([]int{1}).Return([]int{}, nil)

	messageRepo.EXPECT().Send([]entity.Message{
		{
			Text: "test1\nhttps://test1.com",
		},
	}).Return(nil)

	notifiedEventRepo.EXPECT().Save([]int{1}).Return(nil)

	err := u.Send()
	assert.NoError(t, err)
}

func TestSend_ErrorOnGetByKeyword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	eventRepo := mock.NewMockEventRepository(ctrl)
	messageRepo := mock.NewMockMessageRepository(ctrl)
	notifiedEventRepo := mock.NewMockNotifiedEventRepository(ctrl)

	u := usecase.NewNotificationUsecase(eventRepo, messageRepo, notifiedEventRepo)

	eventRepo.EXPECT().GetByKeyword([]string{"名古屋", "愛知"}).Return(nil, errors.New("error"))

	err := u.Send()
	assert.Error(t, err)
}

func TestSend_ErrorOnFindByEventIDs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	eventRepo := mock.NewMockEventRepository(ctrl)
	messageRepo := mock.NewMockMessageRepository(ctrl)
	notifiedEventRepo := mock.NewMockNotifiedEventRepository(ctrl)

	u := usecase.NewNotificationUsecase(eventRepo, messageRepo, notifiedEventRepo)

	expectedEvents := []entity.Event{
		{
			EventID:  1,
			Title:    "test1",
			EventURL: "https://test1.com",
		},
	}

	eventRepo.EXPECT().GetByKeyword([]string{"名古屋", "愛知"}).Return(expectedEvents, nil)

	notifiedEventRepo.EXPECT().FindByEventIDs([]int{1}).Return(nil, errors.New("error"))

	err := u.Send()
	assert.Error(t, err)
}

func TestSend_ErrorOnSend(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	eventRepo := mock.NewMockEventRepository(ctrl)
	messageRepo := mock.NewMockMessageRepository(ctrl)
	notifiedEventRepo := mock.NewMockNotifiedEventRepository(ctrl)

	u := usecase.NewNotificationUsecase(eventRepo, messageRepo, notifiedEventRepo)

	expectedEvents := []entity.Event{
		{
			EventID:  1,
			Title:    "test1",
			EventURL: "https://test1.com",
		},
	}

	eventRepo.EXPECT().GetByKeyword([]string{"名古屋", "愛知"}).Return(expectedEvents, nil)

	notifiedEventRepo.EXPECT().FindByEventIDs([]int{1}).Return([]int{}, nil)

	messageRepo.EXPECT().Send([]entity.Message{
		{
			Text: "test1\nhttps://test1.com",
		},
	}).Return(errors.New("error"))

	err := u.Send()
	assert.Error(t, err)
}

func TestSend_ErrorOnSave(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	eventRepo := mock.NewMockEventRepository(ctrl)
	messageRepo := mock.NewMockMessageRepository(ctrl)
	notifiedEventRepo := mock.NewMockNotifiedEventRepository(ctrl)

	u := usecase.NewNotificationUsecase(eventRepo, messageRepo, notifiedEventRepo)

	expectedEvents := []entity.Event{
		{
			EventID:  1,
			Title:    "test1",
			EventURL: "https://test1.com",
		},
	}

	eventRepo.EXPECT().GetByKeyword([]string{"名古屋", "愛知"}).Return(expectedEvents, nil)

	notifiedEventRepo.EXPECT().FindByEventIDs([]int{1}).Return([]int{}, nil)

	messageRepo.EXPECT().Send([]entity.Message{
		{
			Text: "test1\nhttps://test1.com",
		},
	}).Return(nil)

	notifiedEventRepo.EXPECT().Save([]int{1}).Return(errors.New("error"))

	err := u.Send()
	assert.Error(t, err)
}
