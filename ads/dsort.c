#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void dsort(char *str) {
    int *count = (int*)malloc(26 * sizeof(int));
    for(int i = 0; i < 26; ++i) {
        count[i] = 0;
    }
    int length = strlen(str);
    for(int i = 0; i < length; ++i) {
        count[(int)(str[i] - 'a')]++;
    }
    for(int i = 1; i < 26; ++i) {
        count[i] = count[i] + count[i - 1];
    }
    char *sorted = (char*)malloc(length + 1);
    for(int i = length - 1; i >= 0; --i) {
        int j = count[(int)(str[i] - 'a')] - 1;
        count[(int)(str[i] - 'a')] = j;
        sorted[j] = str[i];
    }
    for(int i = 0; i < length; ++i) {
        str[i] = sorted[i];
    }
    free(sorted);
    free(count);
}

int main(int argc, char ** argv) {
    char *str = (char*)malloc(1000001);
    scanf("%s", str);
    dsort(str);
    printf("%s", str);
    free(str);
    return 0;
}