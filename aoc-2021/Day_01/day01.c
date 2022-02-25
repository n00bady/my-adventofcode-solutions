#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main (int argc, char *argv[]) {
    FILE *fp=fopen(argv[1], "r");
    int lines_count = 0, increase_count = 0, temp = -1;
    char *line = NULL;
    size_t len = 0, nread;

    if (argc != 2) {
        perror("First arguments need to be the input text file!\nNo more no less!\n");
        return 1;
    }
    else if (fp == NULL) {
        perror("Unable to open file!");
        return 1;
    }

    while (nread = getline(&line, &len, fp) != -1) {
        if (temp != -1) {
            if (temp < atoi(line)) {
                increase_count++;
            }
        }
        temp = atoi(line);
        lines_count++;
    }
    // Reset
    rewind(fp); 
    *line = NULL;
    len = 0;
    // Define array with lines_count
    int array[lines_count];
    int i = 0, sum_count = 0;
    // Do the deed 
    // Probably there is a better way to do it but I am not that smart...
    while (nread = getline(&line, &len, fp) != -1) {
        array[i] = atoi(line);
        i++;
    }
    int prevsum = array[0] + array[1] + array[2], sum = 0;
    for (int j=4;j<=lines_count;j++) {
        sum = array[j-2] + array[j-1] + array[j];
        if (prevsum < sum) {
            sum_count++;
        }
        prevsum = sum;
    }
    
    printf("\n--------------------\n");
    printf("Total lines: %d\n", lines_count);
    printf("Total increases: %d\n", increase_count);
    printf("Total increases for pt.2: %d\n", sum_count);
    
    free(line);
    fclose(fp);
    return 0;
}

