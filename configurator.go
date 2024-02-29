func DefaultConfig() network.Config {
	encCfg := MakeEncodingConfig()

	genState := ModuleBasics.DefaultGenesis(encCfg.Codec)
	genState["pylons"] = CustomGenesisHelper(encCfg.Codec)}

	// add default SendEnabled
	bankModuleName := banktypes.ModuleName
	bankModule, ok := gs[bankModuleName]

func CustomGenesisHelper(cdc codec.Codec) json.RawMessage {
	return cdc.MustMarshalJSON(types.NetworkTestGenesis())
}
