#!/bin/bash

# Helm deployment script for Jelastic Golang Hello application
set -e

RELEASE_NAME="jelastic-golang-hello"
NAMESPACE="jelastic-golang-hello"
CHART_PATH="./helm/jelastic-golang-hello"
ENVIRONMENT=${1:-"production"}
IMAGE_TAG=${2:-"latest"}
REGISTRY=${3:-"ghcr.io/ruk-com-cloud/jelastic-golang-hello"}

echo "🚀 Deploying Jelastic Golang Hello with Helm"
echo "Release: $RELEASE_NAME"
echo "Namespace: $NAMESPACE"
echo "Environment: $ENVIRONMENT"
echo "Image: $REGISTRY:$IMAGE_TAG"

# Check if helm is available
if ! command -v helm &> /dev/null; then
    echo "❌ Helm not found. Please install Helm first."
    exit 1
fi

# Check if kubectl is available
if ! command -v kubectl &> /dev/null; then
    echo "❌ kubectl not found. Please install kubectl first."
    exit 1
fi

# Check if cluster is accessible
if ! kubectl cluster-info &> /dev/null; then
    echo "❌ Cannot connect to Kubernetes cluster. Please check your kubeconfig."
    exit 1
fi

echo "✅ Kubernetes cluster connection verified"

# Create namespace if it doesn't exist
echo "📦 Creating namespace if needed..."
kubectl create namespace $NAMESPACE --dry-run=client -o yaml | kubectl apply -f -

# Prepare Helm values based on environment
VALUES_FILE=""
case $ENVIRONMENT in
    "development")
        VALUES_FILE="--set replicaCount=1 --set config.environment=development --set config.testMessage='Development with Helm!' --set resources.limits.cpu=100m --set resources.limits.memory=64Mi"
        ;;
    "staging")
        VALUES_FILE="--set replicaCount=2 --set config.environment=staging --set config.testMessage='Staging with Helm!' --set ingress.enabled=true"
        ;;
    "production")
        VALUES_FILE="--set replicaCount=3 --set config.environment=production --set config.testMessage='Production with Helm!' --set ingress.enabled=true --set autoscaling.enabled=true"
        ;;
    *)
        echo "⚠️  Unknown environment: $ENVIRONMENT. Using default values."
        ;;
esac

# Deploy with Helm
echo "🚢 Deploying with Helm..."
helm upgrade --install $RELEASE_NAME $CHART_PATH \
    --namespace $NAMESPACE \
    --set image.repository=$REGISTRY \
    --set image.tag=$IMAGE_TAG \
    $VALUES_FILE \
    --wait --timeout=5m

# Show deployment status
echo "📊 Deployment Status:"
kubectl get all -n $NAMESPACE

# Show Helm status
echo "📈 Helm Release Status:"
helm status $RELEASE_NAME -n $NAMESPACE

# Show access information
echo ""
echo "🎉 Deployment completed successfully!"
echo ""
echo "📝 Access methods:"

# Check if NodePort is enabled
if kubectl get svc -n $NAMESPACE | grep -q NodePort; then
    NODE_PORT=$(kubectl get svc -n $NAMESPACE -o jsonpath='{.items[?(@.spec.type=="NodePort")].spec.ports[0].nodePort}')
    echo "1. NodePort: http://<node-ip>:$NODE_PORT/"
fi

echo "2. Port forward: kubectl port-forward -n $NAMESPACE svc/$RELEASE_NAME 8080:80"
echo "3. Direct pod access: kubectl exec -n $NAMESPACE -it deployment/$RELEASE_NAME -- sh"

# Check if ingress is enabled
if kubectl get ingress -n $NAMESPACE &> /dev/null; then
    echo "4. Ingress: Check 'kubectl get ingress -n $NAMESPACE' for external access"
fi

echo ""
echo "📋 Useful Helm commands:"
echo "• Upgrade: helm upgrade $RELEASE_NAME $CHART_PATH -n $NAMESPACE"
echo "• Rollback: helm rollback $RELEASE_NAME -n $NAMESPACE"
echo "• Uninstall: helm uninstall $RELEASE_NAME -n $NAMESPACE"
echo "• History: helm history $RELEASE_NAME -n $NAMESPACE"
echo ""
echo "📋 Testing commands:"
echo "• Health check: kubectl port-forward -n $NAMESPACE svc/$RELEASE_NAME 8080:80 && curl http://localhost:8080/api/health"
echo "• View logs: kubectl logs -n $NAMESPACE -l app.kubernetes.io/name=$RELEASE_NAME"