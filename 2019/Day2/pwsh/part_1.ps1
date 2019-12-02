[int[]]$puzzleInput = (Get-Content -Path "./input.txt") -split ','
$opCode = [System.Collections.Generic.List[int]]::new()
$opCode.AddRange($puzzleInput)

<#
To do this, before running the program,
    replace position 1 with the value 12
and...
    replace position 2 with the value 2.
What value is left at position 0 after the program halts?
#>
$opCode[1] = 12
$opCode[2] = 2

:outer for ($i = 0; $i -lt $opCode.Count; $i = $i + 4) {
    $currentSet = $opCode.GetRange($i,4)

    switch ($currentSet[0]) {
        1 {
            $opCode[$currentSet[3]] = $opCode[$currentSet[1]] + $opCode[$currentSet[2]]
            break
        }
        2 {
            $opCode[$currentSet[3]] = $opCode[$currentSet[1]] * $opCode[$currentSet[2]]
            break
        }
        99 {
            Write-Warning -Message 'OpCode 99 hit, exiting program.'
            break outer
        }
        Default {Write-Error -Message 'Invalid computer operation.  Get to escape pod!' -ErrorAction Stop }
    }
}
Write-Output -InputObject $opCode[0]
# Write-Output -InputObject ($opCode -Join ',')
