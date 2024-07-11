@echo off
g++ -shared -o libunhook_ntdll.dll unhook_ntdll.cpp -std=c++11 -static-libgcc -static-libstdc++ -lpsapi
go build -o main.exe main.go
./main.exe
pause
