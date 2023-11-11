# scripts/run.ps1

# 检查Go是否安装
if (-not (Get-Command "go" -ErrorAction SilentlyContinue)) {
    Write-Error "Go is not installed. Please install Go to proceed."
    exit 1
}

# 运行项目
Write-Host "Running application..."
go run ../cmd/main.go

Write-Host "Application is running."
