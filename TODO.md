
# Main features
- fuzzy search
  (To search through many backends at a time?) No because of same backends in different ctxs possabilities. Dont list em at last...
- correct working graph with beauty
- COPY from one backend (GitLab -> Vault)
- ctx with interactive kubectx-like mode (also dont forget for backends like kubens)

# Building Platform:
-  Parse backend type (Vault, GitLab, AWS, k8s) the smart way

# Errors
- Hadnle Ctx errors
- ??? always more

# Cli/ui
- Convinient way to pront connections to backens
- Smart way to list actual data for backrnds in ctx? For example user wants to list/graph/etc from only one backend not from the whole ctx

# Code
- refactor vault explicit executes with implicit vault "github.com/hashicorp/vault/api" (HOT)