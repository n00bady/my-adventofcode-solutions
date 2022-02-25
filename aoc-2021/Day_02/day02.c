#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int part2 (char *file);

int main (int argc, char *argv[]) {
    FILE *fp=fopen(argv[1], "r");
    int pos_h = 0, pos_v = 0;

    if (fp == NULL) {
        perror("Unable to open file!");
        return 1;
    }

    char *line = NULL;
    size_t len = 0, nread;
    while (nread = getline(&line, &len, fp) != -1) {
        char *p = strtok(line, " ");
        char *direction = p; 
        int unit;
        while (p != NULL) {
            unit = atoi(p);
            p = strtok(NULL, " ");
        }
        if (strcmp("forward", direction) == 0) {
            pos_h = pos_h + unit;
        } else if (strcmp("down", direction) == 0) {
            pos_v = pos_v + unit;
        } else if (strcmp("up", direction) == 0) {
            pos_v = pos_v - unit;
        } else {
            perror("\nThere is something wrong I can feel it.\n");
        }
    }
    printf("--- Part 1 ---\n");
    printf("Horizontal pos: %d\nVertical pos: %d\n", pos_h, pos_v);
    printf("Horizontal pos * Vertical pos: %d\n", pos_h*pos_v);
    int ans = part2(argv[1]);
    printf("--- Part 2 ---\n");
    printf("Horizontal pos * Vertical pos: %d\n", ans);
    
    free(line);
    fclose(fp);
    return 0;
}

int part2 (char *file) {
    FILE *fp=fopen(file, "r");
    int pos_h = 0, pos_v = 0, aim = 0;
    char *line = NULL;
    size_t len = 0, nread;

    while (nread = getline(&line, &len, fp) != -1) {
        char *p = strtok(line, " ");
        char *direction = p;
        int unit;
        while (p != NULL) {
            unit = atoi(p);
            p = strtok(NULL, " ");
        }
        if (strcmp("forward", direction) == 0) {
            pos_h = pos_h + unit;
            pos_v = pos_v + unit * aim;
        } else if (strcmp("down", direction) == 0) {
            aim = aim + unit;
        } else if (strcmp("up", direction) == 0) {
            aim = aim - unit;
        } else {
            perror("\nSomething borked!\n");
        }
    }
    
    return pos_h*pos_v;
}
