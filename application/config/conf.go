/*
 * @Author: Wen Jiajun
 * @Date: 2022-05-06 11:59:46
 * @LastEditors: Wen Jiajun
 * @LastEditTime: 2022-07-02 12:54:27
 * @FilePath: \application\config\conf.go
 * @Description:
 */
package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/ini.v1"
)

/*
 ccpPath       = "./organizations/peerOrganizations/org1.example.com/connection-org1.yaml"
 user          = "appUser"
 walletPath    = "wallet"
 channelName   = "mychannel"
 chaincodeName = "stateGrid"
*/
type MyConfig struct {
	WalletPath    string `ini:"wallet_path"`
	CcpPath       string `ini:"ccp_path"`
	User          string `ini:"user"`
	CredPath      string `ini:"cred_path"`
	ChannelName   string `ini:"channel_name"`
	ChaincodeName string `ini:"chaincode"`
}

var Conf MyConfig

func LoadConfig() error {
	pwd, _ := os.Getwd()
	fileConfig, err := ini.Load(pwd + "/config/config.ini")
	if err != nil {
		log.Println(err)
		return err
	}

	err = fileConfig.MapTo(&Conf)
	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Println(Conf)
	return nil
}
