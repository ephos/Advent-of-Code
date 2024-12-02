$lists = Get-Content input.txt

$splitLists = $lists.Split("`n")

$leftNumbersList = [System.Collections.Generic.List[Int]]::new()
$rightNumbersList = [System.Collections.Generic.List[Int]]::new()

foreach ($line in $splitLists) {
    $item = $line.Split('  ').Trim()
    $leftNumbersList.Add($item[0]) | Out-Null
    $rightNumbersList.Add($item[1]) | Out-Null
}

if ($leftNumbersList.Count -ne $rightNumbersList.Count) {throw "Lists aren't equal..."}

$leftNumbersList = $leftNumbersList | Sort-Object
$rightNumbersList = $rightNumbersList | Sort-Object
$final = [System.Collections.ArrayList]::new()

for ($i = 0; $i -lt $leftNumbersList.count; $i++) {

   $leftNum = $leftNumbersList[$i]
   $rightNum = $rightNumbersList[$i]

   $counter = 0
   foreach ($n in $rightNumbersList) {
     #"check if $leftNum -eq $n"
     if ($leftNum -eq $n) {
         $counter++
     }
   }
   $simScore = $leftNum * $counter

   $ret = [PSCustomObject]@{
      Index = $i
      LeftNum = $leftNum
      RightNum = $rightNum
      Sum = $leftNum + $rightNum
      SimilarityScore = $simScore
      Diff = [Math]::Abs($leftNum - $rightNum)
   }

   $final.Add($ret) | out-null
}
$part1 = $final.Diff | Measure-Object -Sum
$part2 = $final.SimilarityScore | Measure-Object -Sum

$part1.Sum
$part2.Sum

