cd ../../
cd ./cmd/dashboard-tui-client
go build -o dashboard-tui

cd ../../
mv ./cmd/dashboard-tui-client/dashboard-tui ./build/dashboard-tui

cp .env ./build/dashboard-tui/.env
