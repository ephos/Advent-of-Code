$input = Get-Content -Path "./input.txt"

$totalFuel = 0

foreach ($module in $input) {
    $module = [Math]::Floor(($module / 3) - 2)
    $totalFuel = $totalFuel + $module
}

Write-Output -InputObject $totalFuel