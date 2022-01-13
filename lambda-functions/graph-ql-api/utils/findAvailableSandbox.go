package utils

import (
	"fmt"
	"lambda/aws-sandbox/graph-ql-api/models"
	"strconv"
)

func FindAvailableSandbox(items []models.SandboxItem) (*models.SandboxItem, error) {

	for _, item := range items {
		b, err := strconv.ParseBool(item.Available)
		if err != nil {
			return nil, err
		}
		if b {
			return &item, nil
		}
	}
	return nil, fmt.Errorf("")

}
