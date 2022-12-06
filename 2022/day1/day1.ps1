$puzzle = Get-Content -Path ./input.txt -Raw

$elfcals = $puzzle.split("`n`n").trim("")

$elfList = [System.Collections.Generic.List[PSCustomObject]]::new()
foreach ($e in $elfcals) {
    $cals = $e.Split("`n") | Measure-Object -Sum
    $elf = [PSCustomObject]@{
        "Elf" = $elfcals.IndexOf($e)
        "Calories" = $cals.Sum
    }
    $elfList.Add($elf)
}

$elves = $elfList | Sort-Object -Descending -Property Calories

"Part 1: {0}" -f $elves[0].Calories
"Part 2: {0}" -f ($elves | Select-Object -First 3 | Measure-Object -Property Calories -Sum).Sum
