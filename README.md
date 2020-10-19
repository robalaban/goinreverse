### 🚘 Reverse-proxy build in GO 🚘

**[WIP]** Project is still in the development phase.

`goinreverse` is a simple and configurable reverse-proxy built with go.

#### Configuration

Configure the proxy easily to suit your needs through the `config.yaml` file:

```yaml
timeout: 1 # in seconds
servers:
  - name: andromeda
    url: http://localhost:8081/healthCheck
  - name: milkyway
    url: http://localhost:8082/healthCheck
```

##### Dependencies

1. `gopkg.in/yaml.v2` - used for parsing configuration file.
