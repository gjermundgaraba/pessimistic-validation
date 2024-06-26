package app

import (
	_ "cosmossdk.io/api/cosmos/tx/config/v1" // import for side-effects
	"cosmossdk.io/depinject"
	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	_ "cosmossdk.io/x/circuit" // import for side-effects
	circuitkeeper "cosmossdk.io/x/circuit/keeper"
	_ "cosmossdk.io/x/evidence" // import for side-effects
	evidencekeeper "cosmossdk.io/x/evidence/keeper"
	feegrantkeeper "cosmossdk.io/x/feegrant/keeper"
	_ "cosmossdk.io/x/feegrant/module" // import for side-effects
	nftkeeper "cosmossdk.io/x/nft/keeper"
	_ "cosmossdk.io/x/nft/module" // import for side-effects
	_ "cosmossdk.io/x/upgrade"    // import for side-effects
	upgradekeeper "cosmossdk.io/x/upgrade/keeper"
	"encoding/hex"
	"encoding/json"
	"fmt"
	abci "github.com/cometbft/cometbft/abci/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	_ "github.com/cosmos/cosmos-sdk/x/auth" // import for side-effects
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authsims "github.com/cosmos/cosmos-sdk/x/auth/simulation"
	_ "github.com/cosmos/cosmos-sdk/x/auth/tx/config" // import for side-effects
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	_ "github.com/cosmos/cosmos-sdk/x/auth/vesting" // import for side-effects
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	_ "github.com/cosmos/cosmos-sdk/x/authz/module" // import for side-effects
	_ "github.com/cosmos/cosmos-sdk/x/bank"         // import for side-effects
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	_ "github.com/cosmos/cosmos-sdk/x/consensus" // import for side-effects
	consensuskeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	_ "github.com/cosmos/cosmos-sdk/x/crisis" // import for side-effects
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	_ "github.com/cosmos/cosmos-sdk/x/distribution" // import for side-effects
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	groupkeeper "github.com/cosmos/cosmos-sdk/x/group/keeper"
	_ "github.com/cosmos/cosmos-sdk/x/group/module" // import for side-effects
	_ "github.com/cosmos/cosmos-sdk/x/mint"         // import for side-effects
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	_ "github.com/cosmos/cosmos-sdk/x/params" // import for side-effects
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	_ "github.com/cosmos/cosmos-sdk/x/slashing" // import for side-effects
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	_ "github.com/cosmos/cosmos-sdk/x/staking" // import for side-effects
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	_ "github.com/cosmos/ibc-go/modules/capability" // import for side-effects
	capabilitykeeper "github.com/cosmos/ibc-go/modules/capability/keeper"
	_ "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts" // import for side-effects
	icacontrollerkeeper "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/controller/keeper"
	icahostkeeper "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/host/keeper"
	_ "github.com/cosmos/ibc-go/v8/modules/apps/29-fee" // import for side-effects
	ibcfeekeeper "github.com/cosmos/ibc-go/v8/modules/apps/29-fee/keeper"
	ibctransferkeeper "github.com/cosmos/ibc-go/v8/modules/apps/transfer/keeper"
	connectiontypes "github.com/cosmos/ibc-go/v8/modules/core/03-connection/types"
	ibckeeper "github.com/cosmos/ibc-go/v8/modules/core/keeper"
	"gopkg.in/yaml.v2"
	"hub/x/pessimist/lightclient"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	pessimistmodulekeeper "hub/x/pessimist/keeper"
	pessimisttypes "hub/x/pessimist/types"
	// this line is used by starport scaffolding # stargate/app/moduleImport

	"hub/docs"
)

const (
	AccountAddressPrefix = "hub"
	Name                 = "hub"
)

var (
	// DefaultNodeHome default home directories for the application daemon
	DefaultNodeHome string
)

var (
	_ runtime.AppI            = (*App)(nil)
	_ servertypes.Application = (*App)(nil)
)

// App extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type App struct {
	*runtime.App
	legacyAmino       *codec.LegacyAmino
	appCodec          codec.Codec
	txConfig          client.TxConfig
	interfaceRegistry codectypes.InterfaceRegistry

	// keepers
	AccountKeeper         authkeeper.AccountKeeper
	BankKeeper            bankkeeper.Keeper
	StakingKeeper         *stakingkeeper.Keeper
	DistrKeeper           distrkeeper.Keeper
	ConsensusParamsKeeper consensuskeeper.Keeper

	SlashingKeeper       slashingkeeper.Keeper
	MintKeeper           mintkeeper.Keeper
	GovKeeper            *govkeeper.Keeper
	CrisisKeeper         *crisiskeeper.Keeper
	UpgradeKeeper        *upgradekeeper.Keeper
	ParamsKeeper         paramskeeper.Keeper
	AuthzKeeper          authzkeeper.Keeper
	EvidenceKeeper       evidencekeeper.Keeper
	FeeGrantKeeper       feegrantkeeper.Keeper
	GroupKeeper          groupkeeper.Keeper
	NFTKeeper            nftkeeper.Keeper
	CircuitBreakerKeeper circuitkeeper.Keeper

	// IBC
	IBCKeeper           *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	CapabilityKeeper    *capabilitykeeper.Keeper
	IBCFeeKeeper        ibcfeekeeper.Keeper
	ICAControllerKeeper icacontrollerkeeper.Keeper
	ICAHostKeeper       icahostkeeper.Keeper
	TransferKeeper      ibctransferkeeper.Keeper

	// Scoped IBC
	ScopedIBCKeeper           capabilitykeeper.ScopedKeeper
	ScopedIBCTransferKeeper   capabilitykeeper.ScopedKeeper
	ScopedICAControllerKeeper capabilitykeeper.ScopedKeeper
	ScopedICAHostKeeper       capabilitykeeper.ScopedKeeper

	// TODO: Remove all this silly pointer stuff after IBC dep inj exists...
	PessimistKeeper              pessimistmodulekeeper.Keeper
	PessimisticLightClientModule lightclient.LightClientModule
	// this line is used by starport scaffolding # stargate/app/keeperDeclaration

	// simulation manager
	sm *module.SimulationManager
}

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, "."+Name)
}

// getGovProposalHandlers return the chain proposal handlers.
func getGovProposalHandlers() []govclient.ProposalHandler {
	var govProposalHandlers []govclient.ProposalHandler
	// this line is used by starport scaffolding # stargate/app/govProposalHandlers

	govProposalHandlers = append(govProposalHandlers,
		paramsclient.ProposalHandler,
		// this line is used by starport scaffolding # stargate/app/govProposalHandler
	)

	return govProposalHandlers
}

// AppConfig returns the default app config.
func AppConfig() depinject.Config {
	return depinject.Configs(
		appConfig,
		// Loads the app config from a YAML file.
		// appconfig.LoadYAML(AppConfigYAML),
		depinject.Supply(
			// supply custom module basics
			map[string]module.AppModuleBasic{
				genutiltypes.ModuleName: genutil.NewAppModuleBasic(genutiltypes.DefaultMessageValidator),
				govtypes.ModuleName:     gov.NewAppModuleBasic(getGovProposalHandlers()),
				// this line is used by starport scaffolding # stargate/appConfig/moduleBasic
			},
		),
	)
}

// New returns a reference to an initialized App.
func New(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	appOpts servertypes.AppOptions,
	baseAppOptions ...func(*baseapp.BaseApp),
) (*App, error) {
	var (
		app        = &App{}
		appBuilder *runtime.AppBuilder

		// merge the AppConfig and other configuration in one config
		appConfig = depinject.Configs(
			AppConfig(),
			depinject.Supply(
				// Supply the application options
				appOpts,
				// Supply with IBC keeper getter for the IBC modules with App Wiring.
				// The IBC Keeper cannot be passed because it has not been initiated yet.
				// Passing the getter, the app IBC Keeper will always be accessible.
				// This needs to be removed after IBC supports App Wiring.
				app.GetIBCKeeper,
				app.GetCapabilityScopedKeeper,
				// Supply the logger
				logger,

				// ADVANCED CONFIGURATION
				//
				// AUTH
				//
				// For providing a custom function required in auth to generate custom account types
				// add it below. By default the auth module uses simulation.RandomGenesisAccounts.
				//
				// authtypes.RandomGenesisAccountsFn(simulation.RandomGenesisAccounts),
				//
				// For providing a custom a base account type add it below.
				// By default the auth module uses authtypes.ProtoBaseAccount().
				//
				// func() sdk.AccountI { return authtypes.ProtoBaseAccount() },
				//
				// For providing a different address codec, add it below.
				// By default the auth module uses a Bech32 address codec,
				// with the prefix defined in the auth module configuration.
				//
				// func() address.Codec { return <- custom address codec type -> }

				//
				// STAKING
				//
				// For provinding a different validator and consensus address codec, add it below.
				// By default the staking module uses the bech32 prefix provided in the auth config,
				// and appends "valoper" and "valcons" for validator and consensus addresses respectively.
				// When providing a custom address codec in auth, custom address codecs must be provided here as well.
				//
				// func() runtime.ValidatorAddressCodec { return <- custom validator address codec type -> }
				// func() runtime.ConsensusAddressCodec { return <- custom consensus address codec type -> }

				//
				// MINT
				//

				// For providing a custom inflation function for x/mint add here your
				// custom function that implements the minttypes.InflationCalculationFn
				// interface.
			),
		)
	)

	if err := depinject.Inject(appConfig,
		&appBuilder,
		&app.appCodec,
		&app.legacyAmino,
		&app.txConfig,
		&app.interfaceRegistry,
		&app.AccountKeeper,
		&app.BankKeeper,
		&app.StakingKeeper,
		&app.DistrKeeper,
		&app.ConsensusParamsKeeper,
		&app.SlashingKeeper,
		&app.MintKeeper,
		&app.GovKeeper,
		&app.CrisisKeeper,
		&app.UpgradeKeeper,
		&app.ParamsKeeper,
		&app.AuthzKeeper,
		&app.EvidenceKeeper,
		&app.FeeGrantKeeper,
		&app.NFTKeeper,
		&app.GroupKeeper,
		&app.CircuitBreakerKeeper,
		&app.PessimistKeeper,
		&app.PessimisticLightClientModule,
		// this line is used by starport scaffolding # stargate/app/keeperDefinition
	); err != nil {
		panic(err)
	}

	// Below we could construct and set an application specific mempool and
	// ABCI 1.0 PrepareProposal and ProcessProposal handlers. These defaults are
	// already set in the SDK's BaseApp, this shows an example of how to override
	// them.
	//
	// Example:
	//
	// app.App = appBuilder.Build(...)
	// nonceMempool := mempool.NewSenderNonceMempool()
	// abciPropHandler := NewDefaultProposalHandler(nonceMempool, app.App.BaseApp)
	//
	// app.App.BaseApp.SetMempool(nonceMempool)
	// app.App.BaseApp.SetPrepareProposal(abciPropHandler.PrepareProposalHandler())
	// app.App.BaseApp.SetProcessProposal(abciPropHandler.ProcessProposalHandler())
	//
	// Alternatively, you can construct BaseApp options, append those to
	// baseAppOptions and pass them to the appBuilder.
	//
	// Example:
	//
	// prepareOpt = func(app *baseapp.BaseApp) {
	// 	abciPropHandler := baseapp.NewDefaultProposalHandler(nonceMempool, app)
	// 	app.SetPrepareProposal(abciPropHandler.PrepareProposalHandler())
	// }
	// baseAppOptions = append(baseAppOptions, prepareOpt)
	//
	// create and set vote extension handler
	// voteExtOp := func(bApp *baseapp.BaseApp) {
	// 	voteExtHandler := NewVoteExtensionHandler()
	// 	voteExtHandler.SetHandlers(bApp)
	// }

	voteExtOp := func(bApp *baseapp.BaseApp) {
		bApp.SetExtendVoteHandler(func(ctx sdk.Context, vote *abci.RequestExtendVote) (*abci.ResponseExtendVote, error) {
			// TODO: Move all this into something async and just pick the latest value from memory to not slow things down
			ctx.Logger().Info("vote extension handler", "height", ctx.BlockHeight())

			validationObjectives := app.PessimistKeeper.GetAllValidationObjectives(ctx)

			validationVotes := make([]pessimisttypes.ValidationVote, 0)

			pessimistConfigPath := os.Getenv("PESSIMIST_CONFIG_PATH")
			if pessimistConfigPath == "" {
				return &abci.ResponseExtendVote{
					VoteExtension: []byte(""), // empty vote extension
				}, nil
			}
			file, err := os.ReadFile(pessimistConfigPath)
			if err != nil {
				ctx.Logger().Error("failed to read pessimist config", "error", err)
				return nil, err
			}
			var pessimistConfig PessimisticValidationConfig
			if err := yaml.Unmarshal(file, &pessimistConfig); err != nil {
				ctx.Logger().Error("failed to unmarshal pessimist config", "error", err)
				return nil, err
			}
			ctx.Logger().Info("pessimist config", "config", pessimistConfig)

			for _, objective := range validationObjectives {
				if !objective.Activated {
					continue
				}
				rpcAddr := pessimistConfig.ChainsToValidate[objective.ClientIdToValidate].RPC
				if rpcAddr == "" {
					ctx.Logger().Warn("rpc address not found (this might be OK)", "client", objective.ClientIdToValidate)
					continue
				}

				statusResp, err := http.Get(fmt.Sprintf("%s/status", rpcAddr))
				if err != nil {
					ctx.Logger().Error("failed to get status from rpc", "rpc", rpcAddr, "error", err)
					continue
				}

				if statusResp.StatusCode != http.StatusOK {
					ctx.Logger().Error("rpc status not ok", "rpc", rpcAddr, "status", statusResp.Status)
					continue
				}

				statusBody, err := io.ReadAll(statusResp.Body)
				if err != nil {
					ctx.Logger().Error("failed to read body", "error", err)
					continue
				}

				var status Status
				if err := json.Unmarshal(statusBody, &status); err != nil {
					ctx.Logger().Error("failed to unmarshal body", "error", err, "body", string(statusBody), "endpoint", fmt.Sprintf("%s/status", rpcAddr))
					continue
				}

				headerResp, err := http.Get(fmt.Sprintf("%s/header?height=%s", rpcAddr, status.Result.SyncInfo.LatestBlockHeight))
				if err != nil {
					ctx.Logger().Error("failed to get header from rpc", "rpc", rpcAddr, "error", err, "height", status.Result.SyncInfo.LatestBlockHeight)
					continue
				}

				if headerResp.StatusCode != http.StatusOK {
					ctx.Logger().Error("rpc status not ok", "rpc", rpcAddr, "status", headerResp.Status)
					continue
				}

				body, err := io.ReadAll(headerResp.Body)
				if err != nil {
					ctx.Logger().Error("failed to read body", "error", err)
					continue
				}

				var header CometHeader
				if err := json.Unmarshal(body, &header); err != nil {
					ctx.Logger().Error("failed to unmarshal body", "error", err, "body", string(body), "endpoint", fmt.Sprintf("%s/header", rpcAddr))
					continue
				}
				heightInt, err := strconv.Atoi(header.Result.Header.Height)
				if err != nil {
					ctx.Logger().Error("failed to parse height", "error", err, "body", string(body), "endpoint", fmt.Sprintf("%s/header", rpcAddr))
					continue
				}

				hashBytes, err := hex.DecodeString(header.Result.Header.AppHash)
				if err != nil {
					ctx.Logger().Error("failed to decode hash", "error", err)
					continue
				}
				nextValHashBytes, err := hex.DecodeString(header.Result.Header.NextValidatorsHash)
				if err != nil {
					ctx.Logger().Error("failed to decode next validators hash", "error", err)
					continue
				}
				validationVotes = append(validationVotes, pessimisttypes.ValidationVote{
					ClientIdToValidate: objective.ClientIdToValidate,
					ClientIdToUpdate:   objective.ClientIdToNotify,
					Height:             int64(heightInt),
					Timestamp:          header.Result.Header.Time,
					MerkleRoot:         pessimisttypes.MerkleRoot{
						Hash: hashBytes,
					},
					NextValidatorHash:  nextValHashBytes,
				})

				// Check for stuff to relay back
				paths, _ := app.GetIBCKeeper().ConnectionKeeper.GetClientConnectionPaths(ctx, objective.ClientIdToValidate)
				for _, path := range paths {
					ctx.Logger().Info("path", "path", path)
					connection, found := app.GetIBCKeeper().ConnectionKeeper.GetConnection(ctx, path)
					if !found {
						ctx.Logger().Error("connection not found", "path", path)
						continue
					}
					if connection.State == connectiontypes.INIT {
						ctx.Logger().Info("connection in init state, will try some relaying here", "path", path)
						continue
					}
				}
			}

			voteExtension := pessimisttypes.VoteExtension{
				ValidationVotes: validationVotes,
			}
			marshalledVoteExtensions, err := app.appCodec.Marshal(&voteExtension)
			if err != nil {
				ctx.Logger().Error("failed to marshal vote extensions", "error", err)
				return nil, err
			}

			return &abci.ResponseExtendVote{
				VoteExtension: marshalledVoteExtensions,
			}, nil
		})
		bApp.SetVerifyVoteExtensionHandler(func(ctx sdk.Context, vote *abci.RequestVerifyVoteExtension) (*abci.ResponseVerifyVoteExtension, error) {
			consAddr := sdk.ConsAddress(vote.ValidatorAddress)
			ctx.Logger().Info("verify vote extension handler", "height", ctx.BlockHeight(), "vote", vote, "vote extension", string(vote.VoteExtension), "validator hex", hex.EncodeToString(vote.ValidatorAddress), "consAddr", consAddr.String())

			var voteExt pessimisttypes.VoteExtension
			if err := app.appCodec.Unmarshal(vote.VoteExtension, &voteExt); err != nil {
				ctx.Logger().Error("failed to unmarshal vote extension", "error", err)
				return &abci.ResponseVerifyVoteExtension{
					Status: abci.ResponseVerifyVoteExtension_REJECT,
				}, nil
			}

			for _, validationVote := range voteExt.ValidationVotes {
				if validationVote.Height <= 0 {
					ctx.Logger().Error("invalid height", "height", validationVote.Height)
					return &abci.ResponseVerifyVoteExtension{
						Status: abci.ResponseVerifyVoteExtension_REJECT,
					}, nil
				}

				if validationVote.ClientIdToValidate == "" {
					ctx.Logger().Error("invalid client id to validate", "client", validationVote.ClientIdToValidate)
					return &abci.ResponseVerifyVoteExtension{
						Status: abci.ResponseVerifyVoteExtension_REJECT,
					}, nil
				}

				if validationVote.ClientIdToUpdate == "" {
					ctx.Logger().Error("invalid client id to update", "client", validationVote.ClientIdToUpdate)
					return &abci.ResponseVerifyVoteExtension{
						Status: abci.ResponseVerifyVoteExtension_REJECT,
					}, nil
				}

				if validationVote.Timestamp.IsZero() {
					ctx.Logger().Error("invalid timestamp", "timestamp", validationVote.Timestamp)
					return &abci.ResponseVerifyVoteExtension{
						Status: abci.ResponseVerifyVoteExtension_REJECT,
					}, nil
				}
			}

			return &abci.ResponseVerifyVoteExtension{
				Status: abci.ResponseVerifyVoteExtension_ACCEPT,
			}, nil
		})
		bApp.SetPrepareProposal(func(ctx sdk.Context, proposal *abci.RequestPrepareProposal) (*abci.ResponsePrepareProposal, error) {
			ctx.Logger().Info("prepare proposal handler", "height", ctx.BlockHeight(), "extensionvotes", len(proposal.LocalLastCommit.Votes), "enabled height", ctx.ConsensusParams().Abci.VoteExtensionsEnableHeight)

			if ctx.ConsensusParams().Abci.VoteExtensionsEnableHeight > 0 && proposal.Height > ctx.ConsensusParams().Abci.VoteExtensionsEnableHeight {
				ctx.Logger().Info("prepare proposal: vote extension enabled", "height", ctx.BlockHeight())

				clientIdsToUpdate := make(map[string]bool)
				committeeProposal := pessimisttypes.CommitteeProposal{
					Commitments: []pessimisttypes.Commitment{},
				}
				for _, vote := range proposal.LocalLastCommit.Votes {
					ctx.Logger().Info("vote extension received", "vote", vote.String())


					var voteExt pessimisttypes.VoteExtension
					if err := app.appCodec.Unmarshal(vote.VoteExtension, &voteExt); err != nil {
						ctx.Logger().Error("failed to unmarshal vote extension", "error", err)
						return nil, err
					}
					for _, validationVote := range voteExt.ValidationVotes {
						committeeProposal.Commitments = append(committeeProposal.Commitments, pessimisttypes.Commitment{
							ValidatorAddress: vote.Validator.Address,
							CanonicalVoteExtension: pessimisttypes.CanonicalVoteExtension{
								Extension: vote.VoteExtension,
								Height:    proposal.Height - 1,
								Round:     int64(proposal.LocalLastCommit.Round),
								ChainId:   "hub", // TODO get chain id
							},
							ExtensionSignature:     vote.ExtensionSignature,
							Timestamp: validationVote.Timestamp,
						})

						clientIdsToUpdate[validationVote.ClientIdToUpdate] = true
					}
				}

				if len(committeeProposal.Commitments) > 0 {
					specialTx := pessimisttypes.CommitteeProposalSpecialTx{
						CommitteeProposal: committeeProposal,
						ClientIdsToSendTo: []string{},
					}

					for clientId, _ := range clientIdsToUpdate {
						specialTx.ClientIdsToSendTo = append(specialTx.ClientIdsToSendTo, clientId)
					}

					specialTxBz, err := app.appCodec.Marshal(&specialTx)
					if err != nil {
						ctx.Logger().Error("failed to marshal special tx", "error", err)
						return nil, err
					}

					return &abci.ResponsePrepareProposal{
						Txs: append([][]byte{specialTxBz}, proposal.Txs...),
					}, nil
				}

			}

			return &abci.ResponsePrepareProposal{
				Txs: proposal.Txs,
			}, nil
		})
		bApp.SetProcessProposal(func(ctx sdk.Context, proposal *abci.RequestProcessProposal) (*abci.ResponseProcessProposal, error) {
			ctx.Logger().Info("process proposal handler", "height", ctx.BlockHeight())

			if ctx.ConsensusParams().Abci.VoteExtensionsEnableHeight > 0 && proposal.Height > ctx.ConsensusParams().Abci.VoteExtensionsEnableHeight {
				ctx.Logger().Info("process proposal: vote extension enabled", "height", ctx.BlockHeight())
			}

			return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_ACCEPT}, nil
		})
		bApp.SetPreBlocker(func(ctx sdk.Context, req *abci.RequestFinalizeBlock) (*sdk.ResponsePreBlock, error) {
			ctx.Logger().Info("pre-blocker", "height", ctx.BlockHeight())
			res := &sdk.ResponsePreBlock{}
			if len(req.Txs) == 0 {
				return res, nil
			}

			if ctx.ConsensusParams().Abci.VoteExtensionsEnableHeight > 0 && req.Height > ctx.ConsensusParams().Abci.VoteExtensionsEnableHeight {
				specialTxBz := req.Txs[0]
				var specialTx pessimisttypes.CommitteeProposalSpecialTx
				if err := app.appCodec.Unmarshal(specialTxBz, &specialTx); err != nil {
					ctx.Logger().Info("failed to unmarshal special tx", "error", err)
					return res, nil // Dont want this to fail everything and this might be an OK scenario?
				}

				for _, clientId := range specialTx.ClientIdsToSendTo {
					if err := app.GetIBCKeeper().ClientKeeper.UpdateClient(ctx, clientId, &specialTx.CommitteeProposal); err != nil {
						ctx.Logger().Error("failed to update client", "error", err, "client", clientId)
						return res, nil // Don't want the chain to halt because of this
					}
				}
			}

			return res, nil
		})
	}

	baseAppOptions = append(baseAppOptions, voteExtOp)

	app.App = appBuilder.Build(db, traceStore, baseAppOptions...)

	// Register legacy modules
	if err := app.registerIBCModules(appOpts); err != nil {
		return nil, err
	}

	// register streaming services
	if err := app.RegisterStreamingServices(appOpts, app.kvStoreKeys()); err != nil {
		return nil, err
	}

	/****  Module Options ****/

	app.ModuleManager.RegisterInvariants(app.CrisisKeeper)

	// create the simulation manager and define the order of the modules for deterministic simulations
	//
	// NOTE: this is not required apps that don't use the simulator for fuzz testing transactions
	overrideModules := map[string]module.AppModuleSimulation{
		authtypes.ModuleName: auth.NewAppModule(app.appCodec, app.AccountKeeper, authsims.RandomGenesisAccounts, app.GetSubspace(authtypes.ModuleName)),
	}
	app.sm = module.NewSimulationManagerFromAppModules(app.ModuleManager.Modules, overrideModules)
	app.sm.RegisterStoreDecoders()

	// A custom InitChainer can be set if extra pre-init-genesis logic is required.
	// By default, when using app wiring enabled module, this is not required.
	// For instance, the upgrade module will set automatically the module version map in its init genesis thanks to app wiring.
	// However, when registering a module manually (i.e. that does not support app wiring), the module version map
	// must be set manually as follow. The upgrade module will de-duplicate the module version map.
	//
	// app.SetInitChainer(func(ctx sdk.Context, req *abci.RequestInitChain) (*abci.ResponseInitChain, error) {
	// 	app.UpgradeKeeper.SetModuleVersionMap(ctx, app.ModuleManager.GetVersionMap())
	// 	return app.App.InitChainer(ctx, req)
	// })

	if err := app.Load(loadLatest); err != nil {
		return nil, err
	}

	return app, nil
}

// LegacyAmino returns App's amino codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *App) LegacyAmino() *codec.LegacyAmino {
	return app.legacyAmino
}

// AppCodec returns App's app codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *App) AppCodec() codec.Codec {
	return app.appCodec
}

// GetKey returns the KVStoreKey for the provided store key.
func (app *App) GetKey(storeKey string) *storetypes.KVStoreKey {
	kvStoreKey, ok := app.UnsafeFindStoreKey(storeKey).(*storetypes.KVStoreKey)
	if !ok {
		return nil
	}
	return kvStoreKey
}

// GetMemKey returns the MemoryStoreKey for the provided store key.
func (app *App) GetMemKey(storeKey string) *storetypes.MemoryStoreKey {
	key, ok := app.UnsafeFindStoreKey(storeKey).(*storetypes.MemoryStoreKey)
	if !ok {
		return nil
	}

	return key
}

// kvStoreKeys returns all the kv store keys registered inside App.
func (app *App) kvStoreKeys() map[string]*storetypes.KVStoreKey {
	keys := make(map[string]*storetypes.KVStoreKey)
	for _, k := range app.GetStoreKeys() {
		if kv, ok := k.(*storetypes.KVStoreKey); ok {
			keys[kv.Name()] = kv
		}
	}

	return keys
}

// GetSubspace returns a param subspace for a given module name.
func (app *App) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := app.ParamsKeeper.GetSubspace(moduleName)
	return subspace
}

// GetIBCKeeper returns the IBC keeper.
func (app *App) GetIBCKeeper() *ibckeeper.Keeper {
	return app.IBCKeeper
}

// GetCapabilityScopedKeeper returns the capability scoped keeper.
func (app *App) GetCapabilityScopedKeeper(moduleName string) capabilitykeeper.ScopedKeeper {
	return app.CapabilityKeeper.ScopeToModule(moduleName)
}

// SimulationManager implements the SimulationApp interface.
func (app *App) SimulationManager() *module.SimulationManager {
	return app.sm
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *App) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	app.App.RegisterAPIRoutes(apiSvr, apiConfig)
	// register swagger API in app.go so that other applications can override easily
	if err := server.RegisterSwaggerAPI(apiSvr.ClientCtx, apiSvr.Router, apiConfig.Swagger); err != nil {
		panic(err)
	}

	// register app's OpenAPI routes.
	docs.RegisterOpenAPIService(Name, apiSvr.Router)
}

// GetMaccPerms returns a copy of the module account permissions
//
// NOTE: This is solely to be used for testing purposes.
func GetMaccPerms() map[string][]string {
	dup := make(map[string][]string)
	for _, perms := range moduleAccPerms {
		dup[perms.Account] = perms.Permissions
	}
	return dup
}

// BlockedAddresses returns all the app's blocked account addresses.
func BlockedAddresses() map[string]bool {
	result := make(map[string]bool)
	if len(blockAccAddrs) > 0 {
		for _, addr := range blockAccAddrs {
			result[addr] = true
		}
	} else {
		for addr := range GetMaccPerms() {
			result[addr] = true
		}
	}
	return result
}
