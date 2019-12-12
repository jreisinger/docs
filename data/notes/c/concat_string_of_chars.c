// Programming in C, 4th ed. - Ch. 9

#include <stdio.h>

// Concatenate two character arrays
int concat( char result[],
            const char str1[], int n1,
            const char str2[], int n2 )
{
    int i, j;

    // copy str1 to result
    for (i = 0; i < n1; ++i)
        result[i] = str1[i];

    // copy str2 to result
    for (j = 0; j < n2; ++j)
        result[n1 + j] = str2[j];
}

int main(void)
{
    const char s1[] = {'a', 'b', 'c'};
    int n1 = 3;
    const char s2[] = {'d', 'e'};
    int n2 = 2;
    int n3 = n1 + n2;
    char s3[ n3 ];
    int i;

    concat(s3, s1, n1, s2, n2);

    for (i = 0; i < n3; ++i)
        printf("%c", s3[i]);

    printf("\n");

    return 0;
}
