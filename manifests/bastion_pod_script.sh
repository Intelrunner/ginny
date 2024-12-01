#!/bin/bash

set -e

# Configuration variables
NAMESPACE="kube-system"
SECRET_NAME="bastion-ssh-keys"
POD_NAME="bastion-pod"
SSH_KEY_PATH="./bastion-ssh-key"
MANIFEST_FILE="./bastion-pod.yaml"

# Step 1: Generate SSH key pair
echo "Generating SSH key pair..."
ssh-keygen -t rsa -b 2048 -f $SSH_KEY_PATH -N "" -C "bastion-key" >/dev/null

# Step 2: Create a Kubernetes secret for the SSH key
echo "Creating Kubernetes secret..."
kubectl create secret generic $SECRET_NAME \
  --from-file=id_rsa=$SSH_KEY_PATH \
  --from-file=id_rsa.pub=${SSH_KEY_PATH}.pub \
  -n $NAMESPACE --dry-run=client -o yaml | kubectl apply -f -

# Step 3: Create the bastion pod manifest
echo "Creating bastion pod manifest..."
cat <<EOF > $MANIFEST_FILE
apiVersion: v1
kind: Pod
metadata:
  name: $POD_NAME
  namespace: $NAMESPACE
  labels:
    app: bastion
spec:
  containers:
  - name: bastion
    image: debian:bullseye-slim
    command: ["/bin/sh", "-c", "while true; do sleep 3600; done"]
    resources:
      requests:
        memory: "64Mi"
        cpu: "250m"
      limits:
        memory: "128Mi"
        cpu: "500m"
    stdin: true
    tty: true
    volumeMounts:
    - name: ssh-keys
      mountPath: /root/.ssh
      readOnly: true
  volumes:
  - name: ssh-keys
    secret:
      secretName: $SECRET_NAME
  securityContext:
    runAsUser: 1000
    runAsGroup: 1000
    fsGroup: 1000
  restartPolicy: Always
EOF

# Step 4: Apply the bastion pod manifest
echo "Deploying bastion pod..."
kubectl apply -f $MANIFEST_FILE

# Step 5 (Optional): Create a NetworkPolicy for restricted access
read -p "Do you want to restrict access to the bastion pod using a NetworkPolicy? (y/n): " RESTRICT_ACCESS
if [[ "$RESTRICT_ACCESS" == "y" ]]; then
  echo "Creating NetworkPolicy..."
  cat <<EOF | kubectl apply -f -
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: bastion-access
  namespace: $NAMESPACE
spec:
  podSelector:
    matchLabels:
      app: bastion
  policyTypes:
  - Ingress
  - Egress
  ingress:
  - from:
    - ipBlock:
        cidr: $(curl -s https://ipinfo.io/ip)/32 # Your current IP
EOF
else
  echo "Skipping NetworkPolicy creation."
fi

echo "Bastion pod setup complete!"
echo "To access the pod, use: kubectl exec -it $POD_NAME -n $NAMESPACE -- /bin/bash"
