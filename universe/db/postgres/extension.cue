import (
	"namespacelabs.dev/foundation/std/fn"
	"namespacelabs.dev/foundation/std/fn:inputs"
	"namespacelabs.dev/foundation/std/monitoring/tracing"
)

$providerProto: inputs.#Proto & {
	source: "provider.proto"
}

extension: fn.#Extension & {
	instantiate: {
		openTelemetry: tracing.#Exports.TracerProvider
	}

	provides: {
		Database: {
			input: $providerProto.types.DatabaseArgs

			availableIn: {
				go: {
					package: "namespacelabs.dev/foundation/universe/db/postgres"
					type:    "*DB"
				}
			}
		}

		DatabaseReference: {
			input: $providerProto.types.DatabaseReferenceArgs

			availableIn: {
				go: {
					package: "namespacelabs.dev/foundation/universe/db/postgres"
					type:    "string"
				}
			}
		}

		Factory: {
			input: $providerProto.types.FactoryArgs

			availableIn: {
				go: {
					package: "namespacelabs.dev/foundation/universe/db/postgres"
					type:    "Factory"
				}
			}
		}
	}
}
