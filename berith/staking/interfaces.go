package staking

import (
	"math/big"

	"bitbucket.org/ibizsoftware/berith-chain/common"
)

//StakingList list of staked accounts
type StakingList interface {
	GetInfoWithIndex(idx int) (StakingInfo, error)
	GetInfo(address common.Address) (StakingInfo, error)
	SetInfo(address common.Address, x interface{}) error
	Delete(address common.Address) error
	Encode() ([]byte, error)
	Decode(rlpData []byte) (StakingList, error)
	Len() int
	Finalize()
	Print()
}

type StakingInfo interface {
	Address() common.Address
	Value() *big.Int
}

type DataBase interface {
	GetStakingList(key string) (StakingList, error)
	Commit(key string, stakingList StakingList) error
	NewStakingList() StakingList
	Close()
}
