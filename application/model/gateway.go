/*
 * @Author: Wen Jiajun
 * @Date: 2022-05-08 20:13:00
 * @LastEditors: Wen Jiajun
 * @LastEditTime: 2022-07-02 12:51:33
 * @FilePath: \application\model\gateway.go
 * @Description:
 */
package model

import (
	"app/config"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	cfg "github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

var Contract *gateway.Contract

func initFabric() error {
	// Connect to Fabric
	err := os.Setenv("DISCOVERY_AS_LOCALHOST", "true")
	if err != nil {
		log.Fatalf("Error setting DISCOVERY_AS_LOCALHOST environemnt variable: %v", err)
	}

	wallet, err := createNewWallet(config.Conf.WalletPath)
	if err != nil {
		log.Println(err)
		return err
	}

	err = createNewUser(wallet, config.Conf.User)
	if err != nil {
		log.Println(err)
		return err
	}

	gw := initGateway(wallet, config.Conf.User, config.Conf.CcpPath)
	defer gw.Close()

	network, err := gw.GetNetwork(config.Conf.ChannelName)
	if err != nil {
		log.Println(err)
		return err
	}

	Contract = network.GetContract(config.Conf.ChaincodeName)

	return nil
}

func createNewWallet(wpath string) (*gateway.Wallet, error) {
	_, err := os.Stat(wpath)
	if err == nil {
		log.Printf("Wallet: %v exsits, it has been removed", wpath)
		os.RemoveAll(wpath)
	}

	wallet, err := gateway.NewFileSystemWallet(wpath)
	if err != nil {
		log.Fatalf("Failed to create wallet: %v", err)
	}

	return wallet, nil
}

func createNewUser(wallet *gateway.Wallet, user string) error {
	if wallet.Exists(user) {
		log.Panicf("User identidy %v already in the wallet", user)
	}

	credPath := config.Conf.CredPath

	certPath := filepath.Join(credPath, "signcerts", "User1@org1.example.com-cert.pem")
	// read the certificate pem
	cert, err := ioutil.ReadFile(filepath.Clean(certPath))
	if err != nil {
		return err
	}

	keyDir := filepath.Join(credPath, "keystore")
	// there's a single file in this dir containing the private key
	files, err := ioutil.ReadDir(keyDir)
	if err != nil {
		return err
	}
	if len(files) != 1 {
		return fmt.Errorf("keystore folder should have contain one file\n")
	}
	keyPath := filepath.Join(keyDir, files[0].Name())
	key, err := ioutil.ReadFile(filepath.Clean(keyPath))
	if err != nil {
		return err
	}

	identity := gateway.NewX509Identity("Org1MSP", string(cert), string(key))
	return wallet.Put(user, identity)
}

func initGateway(wallet *gateway.Wallet, user string, ccp ...string) *gateway.Gateway {
	gw, err := gateway.Connect(
		gateway.WithConfig(cfg.FromFile(filepath.Clean(config.Conf.CcpPath))),
		gateway.WithIdentity(wallet, user),
	)
	if err != nil {
		log.Fatalf("Unable to instantiate a gateway: %v", err)
	}
	return gw
}

func initContract(gw *gateway.Gateway) error {
	if len(config.Conf.ChannelName) == 0 {
		panic("Nil channel ID")
	}
	network, err := gw.GetNetwork(config.Conf.ChannelName)
	if err != nil {
		return err
	}

	Contract = network.GetContract(config.Conf.ChaincodeName)

	return nil
}

func InitFabric() error {
	err := config.LoadConfig()
	if err != nil {
		log.Println(err)
		return err
	}
	err = initFabric()
	if err != nil {
		log.Println(err)
		return err
	}

	GetReqID()
	GetResID()
	GetTableID()
	return nil
}
