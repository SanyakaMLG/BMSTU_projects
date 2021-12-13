#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int wcount(char *s) {
    int k = 0, res = 0;
    for(int i = 0; i < strlen(s); ++i) {
        if(s[i] != ' ') {
            ++k;
        } else {
            if(k) {
                res++;
                k = 0;
            }
        }
    }
    if(k) res++;
    return res;
}

int main(int argc, char ** argv) {
    char *str = malloc(1000);
    gets(str);
    printf("%d", wcount(str));
    free(str);
    return 0;
}
