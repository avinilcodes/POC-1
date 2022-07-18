package task

import (
	"context"
	"taskmanager/app"
	"taskmanager/db"
	storemocks "taskmanager/db/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type ServiceTestSuite struct {
	suite.Suite
	store   *storemocks.Storer
	logger  *zap.SugaredLogger
	service Service
}

func init() {
	app.InitLogger()
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

func (suite *ServiceTestSuite) SetupTest() {
	suite.logger = app.GetLogger()
	suite.store = &storemocks.Storer{}
	suite.service = NewService(suite.store, suite.logger)
}

func (suite *ServiceTestSuite) TearDownTest() {
	suite.store.AssertExpectations(suite.T())
}

func (suite *ServiceTestSuite) TestAddTask() {
	t := suite.T()
	var task db.Task
	task.Description = "Do test before demo"
	task.TaskStatusCode = "not_scoped"
	ctx := context.Background()

	suite.store.On("CreateTask", ctx, task).Return(nil)

	gotErr := suite.service.addTask(ctx, task)
	assert.NoError(t, gotErr)
}
