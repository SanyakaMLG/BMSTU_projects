#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct Date {
    int Year, Month, Day;
};

int key(int base, struct Date date) {
    if(base == 31) return date.Day - 1;
    if(base == 12) return date.Month - 1;
    if(base == 61) return date.Year - 1970;
}

void dsort(struct Date *dates, int n, int base) {
    int count[base];
    for(int i = 0; i < base; ++i) {
        count[i] = 0;
    }
    for(int i = 0; i < n; ++i) {
        count[key(base, dates[i])]++;
    }
    for(int i = 1; i < base; ++i) {
        count[i] += count[i - 1];
    }
    struct Date sorted[n];
    for(int i = n - 1; i >= 0; --i) {
        int j = count[key(base, dates[i])] - 1;
        count[key(base, dates[i])] = j;
        sorted[j] = dates[i];
    }
    for(int i = 0; i < n; ++i) {
        dates[i] = sorted[i];
    }
}

void datesort(struct Date *dates, int n) {
    dsort(dates, n, 31);
    dsort(dates, n, 12);
    dsort(dates, n, 61);
}

int main(int argc, char ** argv) {
    int n;
    scanf("%d", &n);
    struct Date dates[n];
    for(int i = 0; i < n; ++i) {
        scanf("%d%d%d", &dates[i].Year, &dates[i].Month, &dates[i].Day);
    }
    datesort(dates, n);
    for(int i = 0; i < n; ++i) {
        printf("%04d %02d %02d\n", dates[i].Year, dates[i].Month, dates[i].Day);
    }
    return 0;
}