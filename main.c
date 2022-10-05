#include <stdio.h>
#include <sys/types.h>
#include <unistd.h>

int main()
{
    printf("Starting ...");
    while (1) {
        fork();
    }
    return 0;
}