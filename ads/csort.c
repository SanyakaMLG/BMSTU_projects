#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int strcomp(char *a, char *b) {
    if(strlen(a) > strlen(b)) {
        return 1;
    } else {
        if(strlen(a) == strlen(b)) {
            for(int i = 0; i < strlen(a); ++i) {
                if(a[i] == b[i]) continue;
                if(a[i] > b[i]) return 1;
                else return 0;
            }
        } else {
            return 0;
        }
    }
    return 0;
}

void csort(char *src, char *dest) {
    char **s;
    strcpy(dest, "");
    s = (char**)malloc(1000 * sizeof(char*));
    int i = 0;
    char *pch = strtok(src, " ");
    while(pch != NULL) {
        s[i] = (char*)malloc(1000);
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
            if(strcomp(s[i], s[j])) {
                ++k;
            } else {
                if(!strcmp(s[i], s[j]) && j < i) {
                    ++k;
                }
            }
        }
        sorted[k] = (char*)malloc(1000);
        strcpy(sorted[k], s[i]);
        k = 0;
    }
    for(i = 0; i < n; ++i) {
        strncat(dest, sorted[i], strlen(sorted[i]) + 1);
        strncat(dest, " ", 2);
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
    return 0;
}