package utils

import (
	"lambda/aws-sandbox/graph-ql-api/models"
	"strconv"
)

func FindAvailableSandbox(items []models.SandboxItem) (*models.SandboxItem, bool, error) {
	result := []models.SandboxItem{}

	for _, item := range items {
		b, err := strconv.ParseBool(item.Available)
		if err != nil {
			return &models.SandboxItem{}, false, err
		}
		if b {
			result = append(result, item)
		}
	}

	if len(result) > 0 {
		return &result[0], true, nil
	} else {
		return nil, false, nil
	}
}
