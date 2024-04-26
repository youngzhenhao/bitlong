cd
$start_time = Get-Date
$currentPath = Get-Location
$currentFolderName = Split-Path -Path $currentPath -Leaf
$specificString = "wallet"
if ($currentFolderName -eq $specificString) {
    Write-Host "gomobile is in progress, please wait..."
    Set-Location api
    gomobile bind -target android
    Set-Location ..
    $end_time = Get-Date
    $time_taken = $end_time - $start_time
    Write-Host "Time cost: $($time_taken.TotalSeconds) seconds."
} else {
	Write-Output "Wrong current directory, please run script in wallet."
    pause
}
