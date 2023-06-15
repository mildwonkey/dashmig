gen:
	protoc -I schemas --go_opt=paths=source_relative --go_out=internal/kinds \
		--validate_out="lang=go,paths=source_relative:internal/kinds"  \
		./schemas/dashboard/dashboard1.0.proto

	protoc -I schemas --go_opt=paths=source_relative --go_out=internal/legacydash \
		--validate_out="lang=go,paths=source_relative:internal/kinds"  \
		./schemas/dashboard/dashboard0.1.proto