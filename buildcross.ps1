function main {
	$latestTag = git describe --tags
	go install github.com/mitchellh/gox@latest
	gox -output="./releases/{{.Dir}}_{{.OS}}_{{.Arch}}_$($latestTag)" 
    Set-Location $PSScriptRoot
}

main