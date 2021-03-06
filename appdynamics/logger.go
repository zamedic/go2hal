package appdynamics

import (
	"context"
	"github.com/go-kit/kit/log"
	"time"
)

type loggingService struct {
	logger log.Logger
	Service
}

func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (s *loggingService) sendAppdynamicsAlert(ctx context.Context, chatId uint32, message AppdynamicsMessage) (err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "send appdynamics alert",
			"message", message,
			"chat", chatId,
			"error", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	return s.Service.sendAppdynamicsAlert(ctx, chatId, message)
}
func (s *loggingService) addAppdynamicsEndpoint(chat uint32, endpoint string) (err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "add appdynamics endpoint",
			"endpoint", endpoint,
			"error", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	return s.Service.addAppdynamicsEndpoint(chat, endpoint)
}

func (s *loggingService) executeCommandFromAppd(ctx context.Context, chatId uint32, commandName, applicationID, nodeID string) (err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "execute app dynamics command",
			"command_name", commandName,
			"applicatiom_id", applicationID,
			"node id", nodeID,
			"chat", chatId,
			"error", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	return s.Service.executeCommandFromAppd(ctx, chatId, commandName, applicationID, nodeID)
}
