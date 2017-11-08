package models

import (
	"testing"
	"errors"
	"fmt"
	"microservice-wallet-server/models"
)

func TestGetTbExchangeWalletConfigById(t *testing.T) {
	tbExchangeWalletConfig, err := models.GetTbExchangeWalletConfigById(1)
	if err != nil{
		errors.New("query config error")
	}
	fmt.Println(tbExchangeWalletConfig)


}
