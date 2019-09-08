
#include <stdio.h>
#include <Windows.h>

typedef int(*showui)();
int main(){
    HMODULE h;
    FARPROC proc;
    int ret;
    printf("main start\n");
    h = LoadLibraryA("dll/test.dll");
    proc = GetProcAddress(h, "Showui");
    ret = (showui)(proc)();
    printf("call dll  showui end %d\n",ret);
    FreeLibrary(h);
    printf("main end\n");
    return 0;
}