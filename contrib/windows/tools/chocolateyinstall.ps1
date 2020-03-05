$ErrorActionPreference = 'Stop'
$toolsDir   = "$(Split-Path -parent $MyInvocation.MyCommand.Definition)"
$myRelease  = '2.0.0'
$url        = "https://github.com/jtyr/gbt/releases/download/v$myRelease/gbt-$myRelease-windows-386.zip"
$url64      = "https://github.com/jtyr/gbt/releases/download/v$myRelease/gbt-$myRelease-windows-amd64.zip"

$packageArgs = @{
  packageName   = $env:ChocolateyPackageName
  unzipLocation = $toolsDir
  url           = $url
  url64bit      = $url64

  softwareName  = 'Go Bullet Train (GBT)'

  checksumType  = 'sha256'
  checksum      = '92d0ea961153e82fc63441f4bb9096c0512ac4d142cf81920c526c299b2da2c9'
  checksum64    = 'cb777db5299da8187f1e94cc5389c980e92f56c1030d6c850724410b87b983d9'
}

Install-ChocolateyZipPackage @packageArgs
