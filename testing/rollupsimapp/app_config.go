package rollupsimapp

import (
	"fmt"

	sequencerv1 "github.com/decentrio/rollkit-sdk/api/rollkitsdk/sequencer/module"
	sequencertypes "github.com/decentrio/rollkit-sdk/x/sequencer/types"

	_ "cosmossdk.io/x/circuit"                                // import for side-effects
	_ "cosmossdk.io/x/feegrant/module"                        // import for side-effects
	_ "cosmossdk.io/x/upgrade"                                // import for side-effects
	_ "github.com/cosmos/cosmos-sdk/x/auth/tx/config"         // import for side-effects
	_ "github.com/cosmos/cosmos-sdk/x/auth/vesting"           // import for side-effects
	_ "github.com/cosmos/cosmos-sdk/x/authz/module"           // import for side-effects
	_ "github.com/cosmos/cosmos-sdk/x/bank"                   // import for side-effects
	_ "github.com/cosmos/cosmos-sdk/x/consensus"              // import for side-effects
	_ "github.com/cosmos/cosmos-sdk/x/crisis"                 // import for side-effects
	_ "github.com/cosmos/cosmos-sdk/x/distribution"           // import for side-effects
	_ "github.com/cosmos/cosmos-sdk/x/params"                 // import for side-effects
	_ "github.com/cosmos/cosmos-sdk/x/staking"                // import for side-effects
	_ "github.com/cosmos/interchain-attestation/configmodule" // import for side effects
	_ "github.com/decentrio/rollkit-sdk/x/sequencer"          // import for side-effects
	_ "github.com/decentrio/rollkit-sdk/x/staking"            // import for side-effects

	runtimev1alpha1 "cosmossdk.io/api/cosmos/app/runtime/v1alpha1"
	appv1alpha1 "cosmossdk.io/api/cosmos/app/v1alpha1"
	authmodulev1 "cosmossdk.io/api/cosmos/auth/module/v1"
	authzmodulev1 "cosmossdk.io/api/cosmos/authz/module/v1"
	bankmodulev1 "cosmossdk.io/api/cosmos/bank/module/v1"
	circuitmodulev1 "cosmossdk.io/api/cosmos/circuit/module/v1"
	consensusmodulev1 "cosmossdk.io/api/cosmos/consensus/module/v1"
	crisismodulev1 "cosmossdk.io/api/cosmos/crisis/module/v1"
	distrmodulev1 "cosmossdk.io/api/cosmos/distribution/module/v1"
	feegrantmodulev1 "cosmossdk.io/api/cosmos/feegrant/module/v1"
	genutilmodulev1 "cosmossdk.io/api/cosmos/genutil/module/v1"
	govmodulev1 "cosmossdk.io/api/cosmos/gov/module/v1"
	paramsmodulev1 "cosmossdk.io/api/cosmos/params/module/v1"
	stakingmodulev1 "cosmossdk.io/api/cosmos/staking/module/v1"
	txconfigv1 "cosmossdk.io/api/cosmos/tx/config/v1"
	upgrademodulev1 "cosmossdk.io/api/cosmos/upgrade/module/v1"
	vestingmodulev1 "cosmossdk.io/api/cosmos/vesting/module/v1"
	"cosmossdk.io/core/appconfig"
	"cosmossdk.io/depinject"
	circuittypes "cosmossdk.io/x/circuit/types"
	"cosmossdk.io/x/feegrant"
	upgradetypes "cosmossdk.io/x/upgrade/types"

	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/module"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	consensustypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"
	ibcfeetypes "github.com/cosmos/ibc-go/v9/modules/apps/29-fee/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v9/modules/apps/transfer/types"
	ibcexported "github.com/cosmos/ibc-go/v9/modules/core/exported"

	attestationconfigmodulev1 "github.com/cosmos/interchain-attestation/configmodule/api/configmodule/module/v1"
	attestationconfigtypes "github.com/cosmos/interchain-attestation/configmodule/types"
)

const AccountAddressPrefix = "rollup"

var (
	// module account permissions
	moduleAccPerms = []*authmodulev1.ModuleAccountPermission{
		{Account: authtypes.FeeCollectorName},
		{Account: distrtypes.ModuleName},
		{Account: stakingtypes.BondedPoolName, Permissions: []string{authtypes.Burner, stakingtypes.ModuleName}},
		{Account: stakingtypes.NotBondedPoolName, Permissions: []string{authtypes.Burner, stakingtypes.ModuleName}},
		{Account: govtypes.ModuleName, Permissions: []string{authtypes.Burner}},
		{Account: ibctransfertypes.ModuleName, Permissions: []string{authtypes.Minter, authtypes.Burner}},
		{Account: ibcfeetypes.ModuleName},
	}

	// blocked account addresses
	blockAccAddrs = []string{
		authtypes.FeeCollectorName,
		distrtypes.ModuleName,
		stakingtypes.BondedPoolName,
		stakingtypes.NotBondedPoolName,
		// We allow the following module accounts to receive funds:
		// govtypes.ModuleName
	}

	// application configuration (used by depinject)
	AppConfig = depinject.Configs(appconfig.Compose(&appv1alpha1.Config{
		Modules: []*appv1alpha1.ModuleConfig{
			{
				Name: runtime.ModuleName,
				Config: appconfig.WrapAny(&runtimev1alpha1.Module{
					AppName: "RollupSimApp",
					// NOTE: upgrade module is required to be prioritized
					PreBlockers: []string{
						upgradetypes.ModuleName,
					},
					// NOTE: staking module is required if HistoricalEntries param > 0
					BeginBlockers: []string{
						// cosmos sdk modules
						distrtypes.ModuleName,
						stakingtypes.ModuleName,
						authz.ModuleName,
						genutiltypes.ModuleName,
						// ibc modules
						capabilitytypes.ModuleName,
						ibcexported.ModuleName,
						ibctransfertypes.ModuleName,
						ibcfeetypes.ModuleName,
					},
					EndBlockers: []string{
						// cosmos sdk modules
						crisistypes.ModuleName,
						govtypes.ModuleName,
						sequencertypes.ModuleName,
						stakingtypes.ModuleName,
						feegrant.ModuleName,
						genutiltypes.ModuleName,
						// ibc modules
						ibcexported.ModuleName,
						ibctransfertypes.ModuleName,
						capabilitytypes.ModuleName,
						ibcfeetypes.ModuleName,
					},
					OverrideStoreKeys: []*runtimev1alpha1.StoreKeyConfig{
						{
							ModuleName: authtypes.ModuleName,
							KvStoreKey: "acc",
						},
					},
					// NOTE: The genutils module must occur after staking so that pools are
					// properly initialized with tokens from genesis accounts.
					// NOTE: The genutils module must also occur after auth so that it can access the params from auth.
					InitGenesis: []string{
						capabilitytypes.ModuleName,
						authtypes.ModuleName,
						banktypes.ModuleName,
						distrtypes.ModuleName,
						sequencertypes.ModuleName,
						stakingtypes.ModuleName,
						govtypes.ModuleName,
						crisistypes.ModuleName,
						ibcexported.ModuleName,
						genutiltypes.ModuleName,
						authz.ModuleName,
						ibctransfertypes.ModuleName,
						ibcfeetypes.ModuleName,
						feegrant.ModuleName,
						paramstypes.ModuleName,
						upgradetypes.ModuleName,
						vestingtypes.ModuleName,
						consensustypes.ModuleName,
						circuittypes.ModuleName,
						attestationconfigtypes.ModuleName,
					},
					// When ExportGenesis is not specified, the export genesis module order
					// is equal to the init genesis order
					// ExportGenesis: []string{},
					// Uncomment if you want to set a custom migration order here.
					// OrderMigrations: []string{},
				}),
			},
			{
				Name: authtypes.ModuleName,
				Config: appconfig.WrapAny(&authmodulev1.Module{
					Bech32Prefix:             AccountAddressPrefix,
					ModuleAccountPermissions: moduleAccPerms,
					// By default modules authority is the governance module. This is configurable with the following:
					// Authority: "group", // A custom module authority can be set using a module name
					// Authority: "cosmos1cwwv22j5ca08ggdv9c2uky355k908694z577tv", // or a specific address
				}),
			},
			{
				Name:   vestingtypes.ModuleName,
				Config: appconfig.WrapAny(&vestingmodulev1.Module{}),
			},
			{
				Name: banktypes.ModuleName,
				Config: appconfig.WrapAny(&bankmodulev1.Module{
					BlockedModuleAccountsOverride: blockAccAddrs,
				}),
			},
			{
				Name: stakingtypes.ModuleName,
				Config: appconfig.WrapAny(&stakingmodulev1.Module{
					// NOTE: specifying a prefix is only necessary when using bech32 addresses
					// If not specfied, the auth Bech32Prefix appended with "valoper" and "valcons" is used by default
					Bech32PrefixValidator: fmt.Sprintf("%svaloper", AccountAddressPrefix),
					Bech32PrefixConsensus: fmt.Sprintf("%svalcons", AccountAddressPrefix),
				}),
			},
			{
				Name:   paramstypes.ModuleName,
				Config: appconfig.WrapAny(&paramsmodulev1.Module{}),
			},
			{
				Name:   "tx",
				Config: appconfig.WrapAny(&txconfigv1.Config{}),
			},
			{
				Name:   genutiltypes.ModuleName,
				Config: appconfig.WrapAny(&genutilmodulev1.Module{}),
			},
			{
				Name:   authz.ModuleName,
				Config: appconfig.WrapAny(&authzmodulev1.Module{}),
			},
			{
				Name:   upgradetypes.ModuleName,
				Config: appconfig.WrapAny(&upgrademodulev1.Module{}),
			},
			{
				Name:   distrtypes.ModuleName,
				Config: appconfig.WrapAny(&distrmodulev1.Module{}),
			},
			{
				Name:   feegrant.ModuleName,
				Config: appconfig.WrapAny(&feegrantmodulev1.Module{}),
			},
			{
				Name:   govtypes.ModuleName,
				Config: appconfig.WrapAny(&govmodulev1.Module{}),
			},
			{
				Name:   crisistypes.ModuleName,
				Config: appconfig.WrapAny(&crisismodulev1.Module{}),
			},
			{
				Name:   consensustypes.ModuleName,
				Config: appconfig.WrapAny(&consensusmodulev1.Module{}),
			},
			{
				Name:   circuittypes.ModuleName,
				Config: appconfig.WrapAny(&circuitmodulev1.Module{}),
			},
			{
				Name:   sequencertypes.ModuleName,
				Config: appconfig.WrapAny(&sequencerv1.Module{}),
			},
			{
				Name:   attestationconfigtypes.ModuleName,
				Config: appconfig.WrapAny(&attestationconfigmodulev1.Module{}),
			},
		},
	}),
		depinject.Supply(
			// supply custom module basics
			map[string]module.AppModuleBasic{
				genutiltypes.ModuleName: genutil.NewAppModuleBasic(genutiltypes.DefaultMessageValidator),
				govtypes.ModuleName: gov.NewAppModuleBasic(
					[]govclient.ProposalHandler{
						paramsclient.ProposalHandler,
					},
				),
			},
		))
)
