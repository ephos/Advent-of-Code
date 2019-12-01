$input = Get-Content -Path "./input.txt"

$totalFuel = 0

foreach ($module in $input) {

    $mass = $module
    while ($mass -ge 0) {
        $mass = [Math]::Floor(($mass / 3) - 2)
        if ($mass -le 0) { continue }
        $totalFuel = $totalFuel + $mass
    }
}
Write-Output -InputObject $totalFuel