#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char *fibstr(int n) {
    int fib1 = 1, fib2 = 1, fib = 1;
    for(int i = 2; i < n; ++i) {
        fib = fib1 + fib2;
        fib1 = fib2;
        fib2 = fib;
    }
    char *s1 = malloc(fib > 1 ? fib + 1 : 2);
    char *s2 = malloc(fib > 1 ? fib + 1 : 2);
    char *res = malloc(fib > 1 ? fib + 1 : 2);
    strcpy(s1, "a");
    strcpy(s2, "b");
    if(n == 1) {
        strcpy(res, "a");
        free(s1);
        free(s2);
        return res;
    } else {
        if(n == 2) {
            strcpy(res, "b");
            free(s1);
            free(s2);
            return res;
        } else {
            strcpy(res, "");
            for(int i = 2; i < n; ++i) {
                res = strcat(s1, s2);
                s1 = s2;
                s2 = res;
            }
        }
    }
    free(s1);
    return res;
}

int main(int argc, char ** argv) {
    int n;
    scanf("%d", &n);
    char *c;
    c = fibstr(n);
    printf("%s", c);
    free(c);
    return 0;
}