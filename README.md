# tinet2inet-henge
convert [tinet](https://github.com/slankdev/tinet) and [tinet-go](https://github.com/ak1ra24/tinet-go) config yaml to [inet-henge](https://github.com/codeout/inet-henge) json and Start web server and display inet-henge.

## Usage
`docker build -t tn2ih .`

- exampledata

`docker run --name tm2ih -p 8080:8080 --rm -it tn2ih tn2ih exampledata/spec.yaml`

- your custom tinet config file

`docker run --name tm2ih -v `pwd`/custom.yml:/work/custom.yml -p 8080:8080 --rm -it tn2ih tn2ih custom.yml`

Please access `http://localhost:8080/` .
