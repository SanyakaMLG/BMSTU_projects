#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void prefixes(char *str) {
    int n = strlen(str);
    int pi[n];
    int t;
    pi[0] = t = 0;
    for(int i = 1; i < n; ++i) {
        while(t > 0 && str[t] != str[i]) {
            t = pi[t - 1];
        }
        if(str[t] == str[i]) t++;
        pi[i] = t;
        if(pi[i] != 0 && (i + 1) % (i + 1 - pi[i]) == 0) printf("%d %d\n", i + 1, (i + 1)/(i + 1 - pi[i]));
    }
}

int main(int argc, char ** argv) {
    char *str;
    str = argv[1];
    prefixes(str);
    return 0;
}