# Test Images for Promotion Simulation

Available container images for testing the XSDLC promotion pipeline. Each image exposes an OpenAPI spec at `/openapi.json` on port 8000.

## Images

| Image | Tag | DDT Result | Description |
|-------|-----|------------|-------------|
| `ghcr.io/novelcore/perfect-api` | `latest` | All principles pass | Reference API — complete OpenAPI spec, proper error handling, versioned, test-ready |
| `ghcr.io/novelcore/bad-docs-api` | `latest` | Fails P002 (Documentation), P003 (Errors) | API with missing descriptions, undocumented error codes, incomplete schemas |
| `ghcr.io/novelcore/broken-api` | `latest` | Fails P001 (Compliance), P004 (Schema) | API with invalid OpenAPI spec, missing schemas, broken references |

## Simulating Deployments

To test the promotion pipeline, update the Kustomize overlay for a specific environment in the gitops repo:

### 1. Deploy perfect-api (should pass all gates)

```bash
cd /tmp && git clone https://github.com/novelcore/perfect-api-gitops.git && cd perfect-api-gitops

# Update dev overlay to use perfect-api
cat > dry/overlays/dev/kustomization.yaml <<'EOF'
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
resources:
  - ../../base
images:
  - name: ghcr.io/novelcore/perfect-api
    newTag: latest
EOF

git add . && git commit -m "deploy perfect-api:latest to dev" && git push origin main
```

### 2. Deploy bad-docs-api (should FAIL staging gate in strict mode)

```bash
cat > dry/overlays/staging/kustomization.yaml <<'EOF'
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
resources:
  - ../../base
patches:
  - target:
      kind: Deployment
      name: perfect-api
    patch: |
      - op: replace
        path: /spec/template/spec/containers/0/image
        value: ghcr.io/novelcore/bad-docs-api:latest
EOF

git add . && git commit -m "test: deploy bad-docs-api to staging" && git push origin main
```

### 3. Deploy broken-api (should FAIL any gate)

```bash
cat > dry/overlays/dev/kustomization.yaml <<'EOF'
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
resources:
  - ../../base
patches:
  - target:
      kind: Deployment
      name: perfect-api
    patch: |
      - op: replace
        path: /spec/template/spec/containers/0/image
        value: ghcr.io/novelcore/broken-api:latest
EOF

git add . && git commit -m "test: deploy broken-api to dev" && git push origin main
```

### 4. Rollback to perfect-api

```bash
cat > dry/overlays/dev/kustomization.yaml <<'EOF'
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
resources:
  - ../../base
EOF

git add . && git commit -m "rollback: restore default overlay" && git push origin main
```

## Monitoring

```bash
# Watch workflows
kubectl get workflows -n driveby -w

# Watch promoter PRs
gh pr list --repo novelcore/perfect-api-gitops --state all

# Watch ArgoCD apps
kubectl get applications -n argocd | grep perfect-api

# Watch commit statuses
kubectl get commitstatuses -n driveby

# Check promotion strategy
kubectl get changetransferpolicies -n driveby
```

## Expected Behavior

| Scenario | Dev | Staging Gate | Prod Gate |
|----------|-----|-------------|-----------|
| perfect-api | Deploys | Pass (validate-only + functional-test) | Pass (validate-only + load-test) |
| bad-docs-api | Deploys | Fail (P002/P003 fail strict validation) | N/A (blocked at staging) |
| broken-api | Deploys | Fail (P001/P004 fail any mode) | N/A (blocked at staging) |
