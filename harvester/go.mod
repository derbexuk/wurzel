module github.com/derbexuk/wurzel/harvester

go 1.15

require (
	github.com/derbexuk/wurzel/combiner v0.0.0
	github.com/basgys/goxml2json v1.1.0
	github.com/beevik/etree v1.1.0
	github.com/bitly/go-simplejson v0.5.0 // indirect
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869 // indirect
	github.com/onsi/gomega v1.16.0
	github.com/teris-io/shortid v0.0.0-20201117134242-e59966efd125
	go.mongodb.org/mongo-driver v1.7.2
	golang.org/x/net v0.0.0-20210917221730-978cfadd31cf
	golang.org/x/oauth2 v0.0.0-20210819190943-2bc19b11175f
	gopkg.in/redis.v5 v5.2.9
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/derbexuk/wurzel/harvester => ../harvester

replace github.com/derbexuk/wurzel/combiner => ../combiner
