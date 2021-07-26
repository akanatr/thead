#include <stdio.h>

int main() {

    int num;
    int factorial=1;
    int count = 0;

    printf("input number:");
    scanf("%d", &num);

    while (count < num)
    {
        count++;
        factorial = factorial * count;
    }
    
    printf("%d\n", factorial);

    return 0;
}