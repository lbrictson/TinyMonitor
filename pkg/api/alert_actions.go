package api

import (
	"context"
	"fmt"
	"github.com/lbrictson/TinyMonitor/pkg/db"
	"github.com/sirupsen/logrus"
)

func fireAlertDown(monitor *db.BaseMonitor, message string, dbConn *db.DatabaseConnection, logger *logrus.Logger) {
	ctx := context.TODO()
	for _, alertChannel := range monitor.AlertChannels {
		alertChannelModel, err := dbConn.GetAlertChannel(ctx, alertChannel)
		if err != nil {
			logger.WithError(err).Errorf("Failed to get alert channel %v to send down alert for monitor %v", alertChannel, monitor.Name)
			continue
		}
		err = convertDBAlertChannelToAlertChannelModel(alertChannelModel).SendDown(monitor.Name, message, dbConn)
		if err != nil {
			logger.WithError(err).Errorf("Failed to send down alert for monitor %v to alert channel %v", monitor.Name, alertChannel)
			continue
		}
	}
	return
}

func fireAlertUp(monitor *db.BaseMonitor, dbConn *db.DatabaseConnection, logger *logrus.Logger) {
	ctx := context.TODO()
	for _, alertChannel := range monitor.AlertChannels {
		alertChannelModel, err := dbConn.GetAlertChannel(ctx, alertChannel)
		if err != nil {
			logger.WithError(err).Errorf("Failed to get alert channel %v to send up alert for monitor %v", alertChannel, monitor.Name)
			continue
		}
		err = convertDBAlertChannelToAlertChannelModel(alertChannelModel).SendUp(monitor.Name, fmt.Sprintf("Monitor %v is back up", monitor.Name), dbConn)
		if err != nil {
			logger.WithError(err).Errorf("Failed to send up alert for monitor %v to alert channel %v", monitor.Name, alertChannel)
			continue
		}
	}
}
