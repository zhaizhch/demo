cd backend && go mod tidy && cd ..
cd frontend && npm install && npm run build && cd ..

# 启动后端
cd backend && go run main.go