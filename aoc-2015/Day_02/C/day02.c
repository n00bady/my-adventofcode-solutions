#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// Whata mess...
int main (int argc, char *argv[]) {
    FILE *fp = fopen(argv[1], "r");
    char line[12];
    int l, w, h;
    int lwh[3];
    int paper = 0;
    int ribbon =0;
    char *d = 0; // dimensions of the cuboid from the string
    //rectangular cuboid surface 2(lw + wh + hl)

    if (argc != 2) {
        perror("First argument need to be the input text file!\n No more no less!\n");
        return 1;
    }
    else if (fp == NULL) {
        perror("Unable to open file!");
        return 1;
    }

    while (fgets(line, sizeof(line), fp)) {
        d = strdup(line);
        // There sould be and easier way to do that ?
        int d_len = strlen(d);
        char *dptr = strtok(d, "x");
        int i = 0;
        while (dptr != NULL) {
            lwh[i] = atoi(dptr);
            dptr = strtok(NULL, "x");
            i++;
        }

        // I could just use the array only but I prefer that way because I can see easier the formula.
        l = lwh[0];
        w = lwh[1];
        h = lwh[2];
        int sides[3];
        sides[0] = l*w;
        sides[1] = w*h;
        sides[2] = h*l;
        int small = sides[0];
        int j=1;
        while (j<=2) {
            if (small > sides[j]) {
                small = sides[j];
            }
            j++;
        }
        paper = paper + 2*l*w + 2*w*h + 2*h*l + small;
        // Ribbon
        for (int s=1;s<=2;s++) {
            if (lwh[s-1] > lwh[s]) {
                int temp = lwh[s-1];
                lwh[s-1] = lwh[s];
                lwh[s] = temp;
            }
        }
        ribbon = ribbon + 2*lwh[0] + 2*lwh[1] + l*w*h;
        printf("&%d&\n", ribbon);
    }
    printf("\nTotal warping: %d sqfeet\n", paper);
    printf("Total ribbon: %d feet\n", ribbon);
    fclose(fp);

    return 0;
}
