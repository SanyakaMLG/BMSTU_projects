#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int strdiff(char *a, char *b) {
    if(!strcmp(a, b)) return -1;
    char symbol1, symbol2;
    int res = 0;
    size_t length;
    if(strlen(a) < strlen(b)) length = strlen(a);
    else length = strlen(b);
    for(int i = 0; i < length + 1; ++i) {
        if(a[i] == b[i]) res += 8;
        else {
            symbol1 = a[i];
            symbol2 = b[i];
            break;
        }
    }
    for(int i = 0; i < 7; ++i) {
        if((symbol1 >> i) % 2 == (symbol2 >> i) % 2) res++;
        else break;
    }
    return res;
}

int main(int argc, char ** argv) {
    char *a, *b;
    a = malloc(30);
    b = malloc(30);
    scanf("%s%s", a, b);
    printf("%d", strdiff(a, b));
    free(a);
    free(b);
    return 0;
}
