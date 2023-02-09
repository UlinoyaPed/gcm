function main {
	$latestTag = git describe --tags
	gox -output="./releases/{{.Dir}}_{{.OS}}_{{.Arch}}_$($latestTag)" 
    Set-Location $PSScriptRoot
}

main