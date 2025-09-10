#!/usr/bin/env -S just --justfile
set dotenv-load

gen:
	go tool templ generate && just tailwind &

dev:
	go tool templ generate --watch --proxy="https://localhost:8080" --cmd="go run ."

tailwind:
    tailwindcss -i ./assets/css/input.css -o ./assets/css/output.css --watch

clean-templ:
	#!/usr/bin/env bash
	set -euo pipefail
	find . -type f -name '*_templ.go' -print -exec rm -f {} +

list-templ:
	#!/usr/bin/env bash
	set -euo pipefail
	find . -type f -name '*_templ.go' -print | sed 's|^\./||' | sort
