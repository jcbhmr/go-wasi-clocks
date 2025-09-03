//go:generate rm -rf ./.out/bindings/ ./monotonic-clock/ ./wall-clock/ ./timezone/
//go:generate go tool wit-bindgen-go generate --out ./.out/bindings/ ./wit/
//go:generate mv ./.out/bindings/wasi/clocks/monotonic-clock/ ./monotonic-clock/
//go:generate mv ./.out/bindings/wasi/clocks/wall-clock/ ./wall-clock/
//go:generate mv ./.out/bindings/wasi/clocks/timezone/ ./timezone/
//go:generate rm -rf ./monotonic-clock/empty.s ./wall-clock/empty.s ./timezone/empty.s
//go:generate go tool jet -g "*.go" -e "DO NOT EDIT.\n" "DO NOT EDIT.\n\n//go:build wasip2\n" -e "\\.out/bindings/wasi/clocks/" "" -e "github\\.com/jcbhmr/go-wasi-clocks/\\.out/bindings/wasi/io/" "github.com/jcbhmr/go-wasi-io/" ./monotonic-clock/ ./wall-clock/ ./timezone/

package clocks_test
