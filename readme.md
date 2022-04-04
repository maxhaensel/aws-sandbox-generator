# Codecoverage

[![codecov](https://codecov.io/gh/maxhaensel/aws-sandbox-generator/branch/main/graph/badge.svg?token=15JNFW7CGS)](https://codecov.io/gh/maxhaensel/aws-sandbox-generator)

## Local Starten

nur die Test starten oder einen Mock-Server starten

für Test muss du "nur"
`go test ./...` oder wenn du was spezielles Testen willst `go test -v -run TestDeallocateSandbox ./resolver -count=1` 
Ich starte die Test/Debug immer über VS-Code

Falls du einen "Server" starten willst musst du im Main.go die funcktion local einkommentieren und in func main =>
```
func main() {
    local()
    // lambda.Start(Handler)
}
```
local aufrufen und das Lambda dings raus nehmen.

`go run main.go`

Grundsätzlich bevor es los geht muss das schema.graphql in go-code kompiliert werden `go generate ./schema` das erzeugt dann ein .go file von wo go dann das schema lesen kann

## Frontend

Jedes Mal nachdem man eine GraphQL Query oder Mutation geschrieben hat, sollte man `yarn generate` ausführen. Dieses Skript erzeugt automatisch TypeScript-Types basierend auf dem Schema auf dem Server und den entwickelten GraphQL Documents.

**Achtung:** Damit das Script funktioniert, muss zuvor der GraphQL-Server lokal gestartet werden und unter Port 8080 laufen.