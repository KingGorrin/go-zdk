package embedded

import (
	"math/big"

	"github.com/ignition-pillar/go-zdk/client"
	"github.com/ignition-pillar/go-zdk/utils/template"
	"github.com/zenon-network/go-zenon/chain/nom"
	"github.com/zenon-network/go-zenon/common"
	"github.com/zenon-network/go-zenon/common/types"
	"github.com/zenon-network/go-zenon/vm/embedded/definition"
)

type PtlcApi struct {
	c client.Client
}

func NewPtlcApi(c client.Client) PtlcApi {
	return PtlcApi{c}
}

func (s PtlcApi) GetById(id types.Hash) (*definition.PtlcInfo, error) {
	var result definition.PtlcInfo
	err := s.c.Call(&result, "embedded.ptlc.getById", id)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Contract methods

func (s PtlcApi) Create(
	zts types.ZenonTokenStandard,
	amount *big.Int,
	expirationTime int64,
	pointType uint8,
	pointLock []byte) (*nom.AccountBlock, error) {
	data, err := definition.ABIPtlc.PackMethod(
		definition.CreatePtlcMethodName,
		expirationTime,
		pointType,
		pointLock,
	)
	if err != nil {
		return nil, err
	}
	return template.CallContract(
		s.c.ProtocolVersion(),
		s.c.ChainIdentifier(),
		types.PtlcContract,
		zts,
		amount,
		data,
	), nil
}

func (s PtlcApi) Reclaim(id types.Hash) (*nom.AccountBlock, error) {
	data, err := definition.ABIPtlc.PackMethod(
		definition.ReclaimPtlcMethodName,
		id,
	)
	if err != nil {
		return nil, err
	}
	return template.CallContract(
		s.c.ProtocolVersion(),
		s.c.ChainIdentifier(),
		types.PtlcContract,
		types.ZnnTokenStandard,
		common.Big0,
		data,
	), nil
}

func (s PtlcApi) Unlock(id types.Hash) (*nom.AccountBlock, error) {
	data, err := definition.ABIPtlc.PackMethod(
		definition.UnlockPtlcMethodName,
		id,
	)
	if err != nil {
		return nil, err
	}
	return template.CallContract(
		s.c.ProtocolVersion(),
		s.c.ChainIdentifier(),
		types.PtlcContract,
		types.ZnnTokenStandard,
		common.Big0,
		data,
	), nil
}

func (s PtlcApi) ProxyUnlock(
	id types.Hash,
	address types.Address,
	signature []uint8) (*nom.AccountBlock, error) {
	data, err := definition.ABIPtlc.PackMethod(
		definition.ProxyUnlockPtlcMethodName,
		id,
		address,
		signature,
	)
	if err != nil {
		return nil, err
	}
	return template.CallContract(
		s.c.ProtocolVersion(),
		s.c.ChainIdentifier(),
		types.PtlcContract,
		types.ZnnTokenStandard,
		common.Big0,
		data,
	), nil
}
