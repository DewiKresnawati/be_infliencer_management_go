package config

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

var SnapClient snap.Client

func InitMidtrans() {
	SnapClient.New("MIDTRANS_SERVER_KEY", midtrans.Sandbox)
}
