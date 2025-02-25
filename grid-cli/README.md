# tfcmd

<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-53%25-brightgreen.svg?longCache=true&style=flat)</a>

Threefold CLI to manage deployments on Threefold Grid.

## Usage

First [download](#download) tfcmd binaries.

Login using your [mnemonics](https://threefoldtech.github.io/info_grid/dashboard/portal/dashboard_portal_polkadot_create_account.html) and specify which grid network (mainnet/testnet) to deploy on by running:

```bash
tfcmd login
```

For examples and description of tfcmd commands check out:

- [vm](docs/vm.md)
- [gateway-fqdn](docs/gateway-fqdn.md)
- [gateway-name](docs/gateway-name.md)
- [kubernetes](docs/kubernetes.md)
- [ZDB](docs/zdb.md)

## Download

- Download the binaries from [releases](https://github.com/threefoldtech/tfgrid-sdk-go/releases)
- Extract the downloaded files
- Move the binary to any of `$PATH` directories, for example:

```bash
mv tfcmd /usr/local/bin
```

## Configuration

tf-grid saves user configuration in `.tfgridconfig` under default configuration directory for your system see: [UserConfigDir()](https://pkg.go.dev/os#UserConfigDir)

## Build

Clone the repo and run the following command inside the repo directory:

```bash
make build
```
