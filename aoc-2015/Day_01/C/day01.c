#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(int argc, char *argv[]) {
    FILE *fp = fopen(argv[1], "r");
    int ch;
    int floor = 0;
    int pos = 1;

    if (argc != 2) {
        perror("First argument need to be the input text file!\n No more no less!\n");
        return 1;
    }
    if (fp == NULL) {
        perror("Unable to open file!");
        return 1;
    }
    while ((ch = fgetc(fp)) != EOF) {
        if (ch == '(') {
            floor++;
        }
        else if (ch == ')') {
            floor--;
        }
        // This is a lazy way to do part 2 :)
        if (floor == -1) {
            printf("%d\n", pos);
        }
        pos++;
    }
    printf("%i\n", floor);

    return 0;
}