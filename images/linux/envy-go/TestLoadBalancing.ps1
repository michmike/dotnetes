param (
    [Parameter(Mandatory=$false)][string]$url="http://envygo--v1.apprenda.jon/env",
    [Parameter(Mandatory=$false)][int]$numRequests=3
)

function PrintMap($psmap)
{
    write-host ""
    write-host "--------> Printing Map <--------"
    $keys = $psmap.GetEnumerator() | Sort-Object Value
    foreach ($key in $keys.Name)
    {
        Write-Host "$key`t" $psmap[$key]
    }
}

$UniqueContent = @{}

do {
    for ($i = 1; $i -le $numRequests ; $i++)
    {
        $webresponse = Invoke-WebRequest $url -DisableKeepAlive
        $contentjson = $webresponse.Content
        $content = ConvertFrom-Json $contentjson
        write-host "Received Data: " $content.IPNETv4 $content.HOSTNAME
        $UniqueContent[$content.IPNETv4] += 1
    }

    PrintMap $UniqueContent
    Start-Sleep -Seconds 2
}
while ($true)