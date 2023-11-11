# scripts/build.ps1

# 保存当前工作目录
$currentDir = Get-Location

# 检查Go是否安装
if (-not (Get-Command "go" -ErrorAction SilentlyContinue)) {
    Write-Error "Go is not installed. Please install Go to proceed."
    exit 1
}

# 定义项目根目录相对路径
$rootDir = ".."

# 进入项目根目录
Set-Location -Path $rootDir

# 清理之前的构建
Write-Host "Cleaning up previous builds..."
go clean

# 构建项目 - 输出Windows可执行文件
Write-Host "Building project for Windows..."
go build -o ./bin/app.exe ./cmd/main.go

# 运行测试 (可选)
# Write-Host "Running tests..."
# go test ./...

# 返回原始工作目录
Set-Location -Path $currentDir

Write-Host "Build completed successfully."
