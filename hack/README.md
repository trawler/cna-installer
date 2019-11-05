# Hacks

## Creating a Cluster from scratch using docker

Few notes how to use the `cna-installer` from scratch using Docker.
This can be used for testing / trying the `cna-installer` without
installing any programs / tools to local environment (everything can be tested
only inside the Docker).

Run Docker container:

```bash
docker run --rm -it -e USER="$USER" -e ARM_CLIENT_ID="$ARM_CLIENT_ID" -e ARM_CLIENT_SECRET="$ARM_CLIENT_SECRET" -e ARM_SUBSCRIPTION_ID="$ARM_SUBSCRIPTION_ID" -e ARM_TENANT_ID="$ARM_TENANT_ID" -e SSH_AUTH_SOCK=$SSH_AUTH_SOCK -v $HOME/.ssh:/root/.ssh:ro -v $SSH_AUTH_SOCK:$SSH_AUTH_SOCK ubuntu
```

Execute following command in Docker container...

Install terraform + git and other handy tools:

```bash
apt update
apt install -y curl git golang unzip
```

Clone the application repository and build the binary:

```bash
git clone https://github.com/trawler/cna-installer.git
cd cna-installer
make build
```

Generate SSH keys if not exists:

```bash
test -f $HOME/.ssh/id_rsa || ( install -m 0700 -d $HOME/.ssh && ssh-keygen -b 2048 -t rsa -f $HOME/.ssh/id_rsa -q -N "" )
```

Set the Azure environment variables (described above) if needed:

```bash
export ARM_SUBSCRIPTION_ID=SUBSCRIPTION_ID
export ARM_CLIENT_ID=APP_ID
export ARM_CLIENT_SECRET=SOME_PASSWORD
export ARM_TENANT_ID=MY_TENANT_ID
```

Copy the cna-installer configuration file to the proper location:

```bash
cp cna-installer.example.yaml ~/.cna-installer.yaml
```

Initiate backend:

```bash
cd build/bin
./cna-installer backend init
```

Create cluster:

```bash
./cna-installer cluster create
```
