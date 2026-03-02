package plant

import (
	"errors"
	"mill/miner"
)

func Filter(class string) (miner.MinerInfo, error) {
	for k := range MinerMap {
		if MinerMap[k].Info().Class == class {
			structInfo := MinerMap[k].Info()
			return *structInfo, nil
		}

	}
	return miner.MinerInfo{}, errors.New("shit mate, no such thing has been found")
}

func MinerInfo(id int) *miner.MinerInfo {
	defer mtx.Unlock()

	mtx.Lock()

	return MinerMap[id].Info()
}
