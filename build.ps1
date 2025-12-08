# 配置输出目录 (相对于 server 目录，或者使用绝对路径)
$OutputName = "pmail"
$ReleaseDir = "../release"

# 1. 编译前端
Write-Host "Start building frontend..." -ForegroundColor Green
Set-Location fe
npm install
npm run build
if ($LASTEXITCODE -ne 0) { 
    Write-Host "Frontend build failed!" -ForegroundColor Red
    exit $LASTEXITCODE 
}
Set-Location ..

# 2. 编译后端
Write-Host "Start building backend..." -ForegroundColor Green
Set-Location server

# 确保输出目录存在
$AbsReleaseDir = Join-Path (Get-Location) $ReleaseDir
if (!(Test-Path $AbsReleaseDir)) {
    New-Item -ItemType Directory -Path $AbsReleaseDir | Out-Null
}

# Windows 下通常添加 .exe 后缀
$OutputFile = Join-Path $AbsReleaseDir "$OutputName.exe"

go build -o $OutputFile main.go

if ($LASTEXITCODE -ne 0) {
    Write-Host "Backend build failed!" -ForegroundColor Red
    exit $LASTEXITCODE
}

Write-Host "Build success! Output file: $OutputFile" -ForegroundColor Green
Set-Location ..
