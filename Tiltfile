# Build Docker image
docker_build('micro-writer',
             context='.',
             dockerfile='./Dockerfile',
)

# ConfigKUSTOMIZE_DIR="./resources/dev/"# LoadRestrictor option is passed as migrations need to be accessed outside of base directory
# See: https://github.com/kubernetes-sigs/kustomize/issues/865
manifests = local("kustomize build --load-restrictor LoadRestrictionsNone {dir}".format(dir=KUSTOMIZE_DIR), quiet=True)

# Apply Kubernetes manifests
# Allow duplcates is marked true for when you import multiple go-micro Tiltfiles
#  into a single Tiltfile it will mark the clusterrole as duplicate.
k8s_yaml(manifests, allow_duplicates=True)
