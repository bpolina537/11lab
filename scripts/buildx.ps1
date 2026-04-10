Write-Host "=== Setup docker buildx ==="
docker buildx create --name mybuilder --use
docker buildx inspect --bootstrap

Write-Host "=== Build image for linux/amd64 ==="
docker buildx build --platform linux/amd64 -t go-app:amd64 --load ./go-app

Write-Host "=== Build image for linux/arm64 ==="
docker buildx build --platform linux/arm64 -t go-app:arm64 --load ./go-app

Write-Host "=== Image sizes ==="
docker images --format "table {{.Repository}}\t{{.Tag}}\t{{.Size}}" | Select-String go-app

Write-Host "=== Done ==="