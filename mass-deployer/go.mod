module github.com/threefoldtech/tfgrid-sdk-go/mass-deployer

go 1.21

require (
	github.com/cosmos/go-bip39 v1.0.0
	github.com/rs/zerolog v1.31.0
	github.com/sethvargo/go-retry v0.2.4
	github.com/spf13/cobra v1.8.0
	github.com/stretchr/testify v1.8.4
	github.com/threefoldtech/tfgrid-sdk-go/grid-client v0.13.4
	github.com/threefoldtech/tfgrid-sdk-go/grid-proxy v0.13.4
	github.com/threefoldtech/zos v0.5.6-0.20231029102142-baf8d9ee77fe
	gopkg.in/validator.v2 v2.0.1
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/ChainSafe/go-schnorrkel v1.1.0 // indirect
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/cenkalti/backoff/v3 v3.2.2 // indirect
	github.com/centrifuge/go-substrate-rpc-client/v4 v4.0.12 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/deckarep/golang-set v1.8.0 // indirect
	github.com/decred/base58 v1.0.5 // indirect
	github.com/decred/dcrd/crypto/blake256 v1.0.1 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.2.0 // indirect
	github.com/ethereum/go-ethereum v1.11.6 // indirect
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/gomodule/redigo v2.0.0+incompatible // indirect
	github.com/google/uuid v1.5.0 // indirect
	github.com/gorilla/schema v1.2.1 // indirect
	github.com/gorilla/websocket v1.5.1 // indirect
	github.com/gtank/merlin v0.1.1 // indirect
	github.com/gtank/ristretto255 v0.1.2 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/holiman/uint256 v1.2.3 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jbenet/go-base58 v0.0.0-20150317085156-6237cf65f3a6 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mimoo/StrobeGo v0.0.0-20220103164710-9a04d6ca976b // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pierrec/xxHash v0.1.5 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/rs/cors v1.10.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/threefoldtech/tfchain/clients/tfchain-client-go v0.0.0-20240101155545-8dfb94126345 // indirect
	github.com/threefoldtech/tfgrid-sdk-go/rmb-sdk-go v0.13.4 // indirect
	github.com/vedhavyas/go-subkey v1.0.3 // indirect
	golang.org/x/crypto v0.17.0 // indirect
	golang.org/x/exp v0.0.0-20231006140011-7918f672742d // indirect
	golang.org/x/net v0.19.0 // indirect
	golang.org/x/sync v0.5.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.zx2c4.com/wireguard/wgctrl v0.0.0-20200609130330-bd2cb7843e1b // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/natefinch/npipe.v2 v2.0.0-20160621034901-c1b8fa8bdcce // indirect
)

replace github.com/threefoldtech/tfgrid-sdk-go/grid-client => ../grid-client

replace github.com/threefoldtech/tfgrid-sdk-go/grid-proxy => ../grid-proxy

replace github.com/threefoldtech/tfgrid-sdk-go/rmb-sdk-go => ../rmb-sdk-go
