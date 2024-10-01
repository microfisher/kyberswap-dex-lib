package vaultT1

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/KyberNetwork/ethrpc"
	"github.com/KyberNetwork/logger"

	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/entity"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/valueobject"
)

type PoolsListUpdater struct {
	config       Config
	ethrpcClient *ethrpc.Client
}

func NewPoolsListUpdater(config *Config, ethrpcClient *ethrpc.Client) *PoolsListUpdater {
	return &PoolsListUpdater{
		config:       *config,
		ethrpcClient: ethrpcClient,
	}
}

func (u *PoolsListUpdater) GetNewPools(ctx context.Context, metadataBytes []byte) ([]entity.Pool, []byte, error) {
	logger.WithFields(logger.Fields{
		"dexType": DexType,
	}).Infof("Start updating pools list ...")
	defer func() {
		logger.WithFields(logger.Fields{
			"dexType": DexType,
		}).Infof("Finish updating pools list.")
	}()

	extraBytes, err := json.Marshal(&StaticExtra{
		VaultLiquidationResolver: u.config.VaultLiquidationResolver,
	})
	if err != nil {
		return nil, nil, err
	}

	paths, err := u.getSwapPaths(ctx)
	if err != nil {
		return nil, nil, err
	}

	pools := make([]entity.Pool, 0)

	for _, swapPath := range paths {
		pool := entity.Pool{
			Address:  swapPath.Protocol.String(),
			Exchange: string(valueobject.ExchangeFluidVaultT1),
			Type:     DexType,
			Reserves: entity.PoolReserves{"0", "0"},
			Tokens: []*entity.PoolToken{
				{
					Address:   strings.ToLower(swapPath.TokenIn.String()),
					Weight:    1,
					Swappable: true,
				},
				{
					Address:   strings.ToLower(swapPath.TokenOut.String()),
					Weight:    1,
					Swappable: true,
				},
			},
			StaticExtra: string(extraBytes),
		}

		pools = append(pools, pool)
	}

	return pools, metadataBytes, nil
}

func (u *PoolsListUpdater) getSwapPaths(ctx context.Context) ([]SwapPath, error) {
	var paths []SwapPath

	req := u.ethrpcClient.R().SetContext(ctx)

	req.AddCall(&ethrpc.Call{
		ABI:    vaultLiquidationResolverABI,
		Target: u.config.VaultLiquidationResolver,
		Method: VLRMethodGetAllSwapPaths,
	}, []interface{}{&paths})

	if _, err := req.Aggregate(); err != nil {
		logger.WithFields(logger.Fields{
			"dexType": DexType,
		}).Error("aggregate request failed")
		return nil, err
	}

	return paths, nil
}