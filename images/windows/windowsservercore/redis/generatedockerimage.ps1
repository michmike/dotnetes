if ((test-path .\downloads) -eq $false)
{
	new-item .\downloads -type directory
}
Invoke-WebRequest -Method Get -Uri https://github.com/MSOpenTech/redis/releases/download/win-2.8.2400/Redis-x64-2.8.2400.zip -OutFile .\downloads\redis.zip 
docker pull microsoft/windowsservercore:latest
docker build -t redis:windowsservercore-10.0.14393.321 .
