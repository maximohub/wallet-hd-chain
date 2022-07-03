package walletdispatcher

import (
	"context"
	"github.com/SavourDao/savour-core/rpc/common"
	"github.com/SavourDao/savour-core/wallet"
	"runtime/debug"
	"strings"

	"github.com/SavourDao/savour-core/config"
	wallet2 "github.com/SavourDao/savour-core/rpc/wallet"
	"github.com/SavourDao/savour-core/wallet/bitcoin"
	"github.com/SavourDao/savour-core/wallet/ethereum"
	"github.com/ethereum/go-ethereum/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CommonRequest interface {
	GetChain() string
}

type CommonReply = wallet2.SupportCoinsResponse

type ChainType = string

type WalletDispatcher struct {
	registry map[ChainType]wallet.WalletAdaptor
}

func (d *WalletDispatcher) GetBalance(ctx context.Context, request *wallet2.BalanceRequest) (*wallet2.BalanceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (d *WalletDispatcher) GetTxByAddress(ctx context.Context, request *wallet2.TxAddressRequest) (*wallet2.TxAddressResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (d *WalletDispatcher) GetTxByHash(ctx context.Context, request *wallet2.TxHashRequest) (*wallet2.TxHashResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (d *WalletDispatcher) GetAccount(ctx context.Context, request *wallet2.AccountRequest) (*wallet2.AccountResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (d *WalletDispatcher) GetUtxo(ctx context.Context, request *wallet2.UtxoRequest) (*wallet2.UtxoResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (d *WalletDispatcher) GetMinRent(ctx context.Context, request *wallet2.MinRentRequest) (*wallet2.MinRentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func New(conf *config.Config) (*WalletDispatcher, error) {
	dispatcher := WalletDispatcher{
		registry: make(map[ChainType]wallet.WalletAdaptor),
	}
	walletAdaptorFactoryMap := map[string]func(conf *config.Config) (wallet.WalletAdaptor, error){
		bitcoin.ChainName:  bitcoin.NewChainAdaptor,
		ethereum.ChainName: ethereum.NewChainAdaptor,
	}
	supportedChains := []string{bitcoin.ChainName, ethereum.ChainName}
	for _, c := range conf.Chains {
		if factory, ok := walletAdaptorFactoryMap[c]; ok {
			adaptor, err := factory(conf)
			if err != nil {
				log.Crit("failed to setup chain", "chain", c, "error", err)
			}
			dispatcher.registry[c] = adaptor
		} else {
			log.Error("unsupported chain", "chain", c, "supportedChains", supportedChains)
		}
	}
	return &dispatcher, nil
}

func NewLocal(network config.NetWorkType) *WalletDispatcher {
	dispatcher := WalletDispatcher{
		registry: make(map[ChainType]wallet.WalletAdaptor),
	}

	walletAdaptorFactoryMap := map[string]func(network config.NetWorkType) wallet.WalletAdaptor{
		bitcoin.ChainName:  bitcoin.NewLocalChainAdaptor,
		ethereum.ChainName: ethereum.NewLocalWalletAdaptor,
	}
	supportedChains := []string{bitcoin.ChainName, ethereum.ChainName}

	for _, c := range supportedChains {
		if factory, ok := walletAdaptorFactoryMap[c]; ok {
			dispatcher.registry[c] = factory(network)
		}
	}
	return &dispatcher
}

func (d *WalletDispatcher) Interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		if e := recover(); e != nil {
			log.Error("panic error", "msg", e)
			log.Debug(string(debug.Stack()))
			err = status.Errorf(codes.Internal, "Panic err: %v", e)
		}
	}()

	pos := strings.LastIndex(info.FullMethod, "/")
	method := info.FullMethod[pos+1:]

	chain := req.(CommonRequest).GetChain()
	log.Info(method, "chain", chain, "req", req)

	resp, err = handler(ctx, req)
	log.Debug("Finish handling", "resp", resp, "err", err)
	return
}

func (d *WalletDispatcher) preHandler(req interface{}) (resp *CommonReply) {
	chain := req.(CommonRequest).GetChain()
	if _, ok := d.registry[chain]; !ok {
		return &CommonReply{
			Error:        &common.Error{Code: 2000},
			SupportCoins: nil,
		}
	}
	return nil
}

func (d *WalletDispatcher) GetNonce(ctx context.Context, request *wallet2.NonceRequest) (*wallet2.NonceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (d *WalletDispatcher) GetSupportCoins(ctx context.Context, request *wallet2.SupportCoinsRequest) (*wallet2.SupportCoinsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (d *WalletDispatcher) GetGasPrice(ctx context.Context, request *wallet2.GasPriceRequest) (*wallet2.GasPriceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (d *WalletDispatcher) SendTx(ctx context.Context, request *wallet2.SendTxRequest) (*wallet2.SendTxResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (d *WalletDispatcher) mustEmbedUnimplementedWalletServiceServer() {
	//TODO implement me
	panic("implement me")
}
