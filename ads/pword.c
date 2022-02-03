#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void pword(char *s, char *t) {
    int length_s = strlen(s);
    int length_t = strlen(t);
    char *str = (char*)malloc(length_s + length_t + 2);
    strcpy(str, "");
    strncat(str, s, length_s + 1);
    strncat(str, "#", 2);
    strncat(str, t, length_t + 1);
    int n = strlen(str);
    int pi[n];
    int k;
    pi[0] = k = 0;
    for(int i = 1; i < n; ++i) {
        while(k > 0 && str[k] != str[i]) {
            k = pi[k - 1];
        }
        if(str[k] == str[i]) k++;
        pi[i] = k;
    }
    for(int i = length_s + 1; i < n; ++i) {
        //printf("%d", pi[i]);
        //if(pi[i] == length_s) printf("%d ", i - 2 * length_s);
        if(!pi[i]) {
            printf("no");
            free(str);
            return;
        }
    }
    printf("yes");
    free(str);
}

int main(int argc, char ** argv) {
    char *str = argv[1];
    char *str2 = argv[2];
    pword(str, str2);
    return 0;
}