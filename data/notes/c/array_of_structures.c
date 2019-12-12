// Programming in C, 4th ed. - Ch. 8

#include <stdio.h>

// struct is a user defined type
struct time {
    int hour;
    int minutes;
    int seconds;
};

int main(void) {
    // array of structures
    struct time testTimes[5] = { {11,59,59},{12,0,0},{1,29,59} };

    int i;

    for ( i = 0; i < 5; ++i) {
        printf("%.2i:%.2i:%.2i\n",
         testTimes[i].hour, testTimes[i].minutes, testTimes[i].seconds);
    }

    return 0;
}
