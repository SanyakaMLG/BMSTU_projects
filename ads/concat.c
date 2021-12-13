#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char *concat(char **s, int n) {
    size_t length = 0;
    for(int i = 0; i < n; ++i) length += strlen(s[i]);
    char *res;
    res = malloc(length + 1);
    strcpy(res, "");
    for(int i = 0; i < n; ++i) {
        strncat(res, s[i], strlen(s[i]) + 1);
    }
    return res;
}

int main(int argc, char ** argv) {
    int n;
    scanf("%d", &n);
    char **s;
    s = malloc(n);
    char *str;
    str = malloc(1000);
    for(int i = 0; i < n; ++i) {
        scanf("%s", str);
        s[i] = malloc(strlen(str) + 1);
        strcpy(s[i], str);
    }
    printf("%s", concat(s, n));
    free(s);
    return 0;
}
