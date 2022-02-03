package utils

import (
	"lambda/aws-sandbox/graph-ql-api/models"
	"log"
	"strconv"
)

func FindAvailableSandbox(items []models.SandboxItem) (*models.SandboxItem, error) {

	for _, item := range items {
		b, err := strconv.ParseBool(item.Available)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		if b {
			return &item, nil
		}
	}
	return nil, nil
}
