# Proyecto Toolkit

### Configuración para debug
``` .vscode/launch.json
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Launch Package",
      "type": "go",
      "request": "launch",
      "mode": "auto",
      "program": "${workspaceFolder}/cmd/toolkit",
      "env": {},
      "args": [],
      "cwd": "${workspaceFolder}",
    },
    {
      "name": "Attach to Process",
      "type": "go",
      "request": "attach",
      "processId": "${command:pickProcess}"
    }
  ]
}
```
- Se necesita instalar `go install github.com/go-delve/delve/cmd/dlv@latest`
- Verificar con `dlv version`
- Si se necesitan argumentos se pone en:
``` .vscode/launch.json
"args": [
        "-d",
        "criteria"
      ],
```

### Implementaciones
- Conexión a ElasticSearch
- Conexión a Vault para el manejo de secretos
- Tests