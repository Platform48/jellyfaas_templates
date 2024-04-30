# f_pound_to_dollar
Sample FaaS function to convert a pound to a dollar.

``zip -x "*.git*" "*.idea*" -r f_go_pound_to_dollar.zip f_go_pound_to_dollar``


go mod tidy

(mac/nix)
FUNCTION_TARGET=ConvertGBPToUSD LOCAL_ONLY=true go run cmd/main.go

win32 (using cmd, not powershell)
set "FUNCTION_TARGET=ConvertGBPToUSD" && set "LOCAL_ONLY=true" && go run cmd/main.go

