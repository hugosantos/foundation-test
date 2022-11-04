providers: {
	"namespacelabs.dev/foundation/library/storage/s3:Bucket": {
		initializedWith: imageFrom: binary: "namespacelabs.dev/foundation/library/oss/localstack/prepare"

		resources: {
			server: {
				class: "namespacelabs.dev/foundation/library/runtime:Server"
				intent: package_name: "namespacelabs.dev/foundation/library/oss/localstack/server"
			}
		}
	}
}