# Example use-case:
```shell
1) cd ./vault
2) sudo docker compose up -d
3) podman compose up -d

# export by yourself for security reasons
export VAULT_TOKEN='dev-root-token'

nano ~/.zshrc
alias secretctl="path/to/secretctl"

nano ~/secretctl/config.yaml
# add your backends and it's access vars
```

## Config example
```yaml
apiVersion: v1
kind: Config
preferences: {}

current-context: whole-company-creds

contexts:
  - name: vault-company-name
    backend:
      - v1
  - name: whole-company-creds
    backend:
      - v1
      - g1

backends:
  - name: v1
    type: vault
    address: http://127.0.0.1:8200
    auth:
      env: VAULT_TOKEN

  - name: g1
    type: gitlab
    address: https://gitlab.company.com
    auth:
      env: TOKEN_FOR_EXAMPLE
      exec:
        command: op
        args: ["read", "op://vault/prod-token/credential"]
```

---

## Building to try
```shell
cd /secretctl
go build .
go run ./main.go
```