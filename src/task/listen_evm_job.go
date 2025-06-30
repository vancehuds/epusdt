package task

import (
	"sync"

	"github.com/assimon/luuu/model"
	"github.com/assimon/luuu/model/data"
	"github.com/assimon/luuu/model/service"
	"github.com/assimon/luuu/util/log"
)

type ListenEvmJob struct {
}

var gListenEvmJobLock sync.Mutex

func (r ListenEvmJob) Run() {
	gListenEvmJobLock.Lock()
	defer gListenEvmJobLock.Unlock()

	var wg sync.WaitGroup

	polygon := func() {
		walletAddress, err := data.GetAvailablePolygonWallet()
		if err != nil {
			log.Sugar.Error(err)
			return
		}
		if len(walletAddress) <= 0 {
			return
		}
		for _, address := range walletAddress {
			wg.Add(1)
			go service.EtherscanCallBack(model.ChainNamePolygonPOS, address.Token, &wg)
		}
	}
	polygon()

	bsc := func() {
		walletAddress, err := data.GetAvailableBSCWallet()
		if err != nil {
			log.Sugar.Error(err)
			return
		}
		if len(walletAddress) <= 0 {
			return
		}
		for _, address := range walletAddress {
			wg.Add(1)
			go service.EtherscanCallBack(model.ChainNameBSC, address.Token, &wg)
		}
	}
	bsc()

	wg.Wait()
}
