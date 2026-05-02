# use-case
```shell
1) sudo docker compose up -d
2) podman compose up -d

# export VAULT_DEV_ROOT_TOKEN_ID='dev-root-token'
export VAULT_TOKEN='dev-root-token'
export VAULT_ADDR='http://127.0.0.1:8200'

# or: vault login dev-root-token
vault kv list /secret

nano ~/secretctl/config.yaml
# add your backends and it's access vars

cd /secretctl
go build .
go run ./main.go

```