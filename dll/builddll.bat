:: go build -buildmode=c-shared -o test.dll dll.go
:: go 1.13  1.12.1 all checked
windres -i ui/ui.rc -O coff -o ui.syso
go build  -buildmode=c-shared -o test.dll
exit