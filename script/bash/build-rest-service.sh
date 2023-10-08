cd ../../
cd ./cmd/dashboard-rest-server
go build -o dashboard-server

cd ../../
mv ./cmd/dashboard-rest-server/dashboard-server ./build/dashboard-server

cp .env ./build/dashboard-server/.env
