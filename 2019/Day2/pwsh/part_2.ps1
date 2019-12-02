function Invoke-IntCodeComputer {
    [CmdletBinding()]
    param (
        # Noun
        [Parameter(Mandatory=$true)]
        [ValidateRange(1,100)]
        [int]
        $Noun,
        # Verb
        [Parameter(Mandatory=$true)]
        [ValidateRange(1,100)]
        [int]
        $Verb
    )
    process {
        [int[]]$puzzleInput = (Get-Content -Path "./input.txt") -split ','
        $opCode = [System.Collections.Generic.List[int]]::new()
        $opCode.AddRange($puzzleInput)

        $opCode[1] = $Noun
        $opCode[2] = $Verb

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
                    # Write-Warning -Message 'OpCode 99 hit, exiting program.'
                    break outer
                }
                Default {Write-Error -Message 'Invalid computer operation.  Get to escape pod!' -ErrorAction Stop }
            }
        }
        Write-Output -InputObject $opCode[0]
    }
}

$nouns = 1..100
$verbs = 1..100

:outer for ($n = 1; $n -lt $nouns.count; $n++) {
    for ($v = 1; $v -lt $verbs.Count; $v++) {
        $result = Invoke-IntCodeComputer -Noun $n -Verb $v
        if ($result -eq 19690720) {
            break outer
        }
    }
}
Write-Output -InputObject ("Noun: {0},Verb: {1}, ResultingNumber: {2}" -f $n, $v, $result)