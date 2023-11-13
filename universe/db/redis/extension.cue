import (
	"encoding/json"
	"namespacelabs.dev/foundation/std/fn"
	"namespacelabs.dev/foundation/std/fn:inputs"
	"namespacelabs.dev/foundation/std/core"
	"namespacelabs.dev/foundation/std/monitoring/tracing"
)

$providerProto: inputs.#Proto & {
	source: "provider.proto"
}

extension: fn.#Extension & {
	instantiate: {
		readinessCheck: core.#Exports.ReadinessCheck
		openTelemetry:  tracing.#Exports.TracerProvider
	}

	provides: {
		Redis: {
			input: $providerProto.types.RedisArgs

			availableIn: {
				go: {
					package: "github.com/go-redis/redis/v8"
					type:    "*Client"
				}
			}
		}
	}
}

$redisServer: inputs.#Server & {
	packageName: "namespacelabs.dev/foundation/library/oss/redis/server"
}

configure: fn.#Configure & {
	stack: {
		append: [$redisServer]
	}

	if $redisServer.$addressMap.redis != _|_ {
		startup: {
			args: {
				redis_endpoint: $redisServer.$addressMap.redis
			}
			env: {
				REDIS_ROOT_PASSWORD: fromSecret: "namespacelabs.dev/foundation/library/oss/redis/server:password"
			}
		}
	}
}
