# sme-scaffold

gin clean scaffold

- cobra
- viper
- wire

## wire

cd /wire

wire -> wire_gen.go

## start

core.yaml.sample -> core.yaml

- go run . version
- go run . start

## start ENV

- CORE_SETTINGS_APPLICATION_PORT=11911 go run . start

## figlet

figlet -f digital sme-scaffold

## sonar

    sonar-scanner -Dsonar.projectKey=test-golang -Dsonar.sources=. -Dsonar.host.url=http://127.0.0.1:9001 -Dsonar.login=13f5314fa5f29152a2de0d8031fbd18b57407820

## todo
