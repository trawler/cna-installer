# cna-installer

```bash
cna-installer is a binary that installs, sets-up and configures a
Kubernetes cluster with the CNA stack applications.

Usage:
  cna-installer [command]

Available Commands:
  backend     Manage the remote backend
  cluster     Create a Cluster
  help        Help about any command

Flags:
      --config string   config file (default is $HOME/.cna-installer.yaml)
  -h, --help            help for cna-installer

Use "cna-installer [command] --help" for more information about a command.
```

## Getting Started

### Building the Installer

```bash
git clone git@github.com:trawler/cna-installer.git
cd cna-installer
make build
```

The binary is built inside the cloned repository under the `build/bin` directory.
Output logs are saved under the `build/logs` directory.

### Azure CLI

You will need to install the [Azure Cli](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli?view=azure-cli-latest).

### Az Login

To authenticate the cli agent, simply run `az login` and log in using the
portal's username and password. Once logged in, note the id field of the output
from the `az login` command. This is a simple way to retrieve
the Subscription ID for the Azure account.

```bash
az login

Note, we have launched a browser for you to login. For old experience with device code, use "az login --use-device-code"
You have logged in. Now let us find all the subscriptions to which you have access...
[
  {
    "cloudName": "AzureCloud",
    "id": "SUBSCRIPTION_ID", # ARM_SUBSCRIPTION_ID
    "isDefault": true,
    ...
    "tenantId": "xxxxxx-xxxxx-xxxxx-xxxxx",
    "user": {
      "name": "MY_USER_NAME",
      "type": "user",
    }
  }
]
```

#### Adding an Azure Service Principal

Next, add a new role assignment for the Installer to use:

```bash
az ad sp create-for-rbac -n "my_az_sp" --role contributor

{
  "appId": "APP_ID", # ARM_CLIENT_ID
  "displayName": "my_az_sp",
  "name": "http://my_az_sp",
  "password": "SOME_PASSWORD", # ARM_CLIENT_SECRET
  "tenant": "MY_TENANT_ID" # ARM_TENANT_ID
}
```

#### Set-Up Your Environment Variables

Set the following environment variables, per the mapping below:

```bash
# id field in az login output
export ARM_SUBSCRIPTION_ID=SUBSCRIPTION_ID

# appID field in az ad output
export ARM_CLIENT_ID=APP_ID

# password field in az ad output
export ARM_CLIENT_SECRET=SOME_PASSWORD

# tenant field in az ad output
export ARM_TENANT_ID=MY_TENANT_ID
```

## Creating A Cluster

### Config File

```bash
cp cna-installer.example.yaml $HOME/.cna-installer.yaml
```

Edit the file to set-up your cluster settings.

### Create a Remote Backend

Go to the directory, where the binary was created:

```bash
cd build/bin
```

To do that, simply run:

```bash
./cna-installer backend init
```

### Create Your Cluster

```bash
./cna-installer cluster create
```
