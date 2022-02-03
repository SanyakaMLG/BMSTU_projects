#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void csort(char *src, char *dest) {
    char **s;
    strcpy(dest, "");
    s = (char**)malloc(1001 * sizeof(char*));
    int i = 0;
    char *pch = strtok(src, " ");
    while(pch != NULL) {
        s[i] = (char*)malloc(1001);
        strcpy(s[i], pch);
        pch = strtok(NULL, " ");
        ++i;
    }
    free(pch);
    int n = i, k = 0;
    char **sorted;
    sorted = (char**)malloc(n * sizeof(char*));
    for(i = 0; i < n; ++i) {
        for(int j = 0; j < n; ++j) {
            int len1 = strlen(s[i]);
            int len2 = strlen(s[j]);
            if(len1 > len2) {
                ++k;
            } else {
                if(len1 == len2 && j < i) {
                    ++k;
                }
            }
        }
        sorted[k] = (char*)malloc(1001);
        strcpy(sorted[k], s[i]);
        k = 0;
    }
    for(i = 0; i < n; ++i) {
        if(i != n - 1) {
            strncat(dest, sorted[i], strlen(sorted[i]) + 1);
            strncat(dest, " ", 2);
        } else {
            strncat(dest, sorted[i], strlen(sorted[i]) + 1);
        }
    }
    for(i = 0; i < n; ++i) {
        free(s[i]);
    }
    for(i = 0; i < n; ++i) {
        free(sorted[i]);
    }
    free(s);
    free(sorted);
}

int main(int argc, char ** argv) {
    char *str = (char*)malloc(1500);
    char *res = (char*)malloc(1005);
    gets(str);
    csort(str, res);
    printf("%s", res);
    free(str);
    free(res);
    return 0;
}