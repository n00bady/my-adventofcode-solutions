#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct BingoCard {
    int board[5][5];
} card;

int getCharCount (FILE *fp);
void nums2array (FILE *fp, int nums[]);
void markBoards (card c[], int num);
int check4win (card c);

int main (int argc, char *argv[]) {
    FILE *fp = fopen(argv[1], "r");  // First argument is the random numbers
    FILE *fp2 = fopen(argv[2], "r"); // and second argument are the bingo cards

    if (fp == NULL) {
        perror("Unable to open file!");
        exit(EXIT_FAILURE);
    }

    int wc = getCharCount(fp);
    int nums[wc];
    nums2array(fp, nums);
    printf("---- The numbers: \n");
    for (int i=0; i<wc; i++) {
        printf("%d ", nums[i]);
    }
    printf("\n--------------------\n");

    char *line = NULL;
    size_t len = 0;
    char *token;
    card c[100]; //100 total boards on actual input, too lazy to make a check in here

    // adding boards as int in an array of structs of a 2D int array
    // can't figure out how to ignore empty lines so in the input so
    // you need to remove the empty lines.
    for (int i=0; i<3; i++) {
        for (int row=0; row<5; row++){
            getline(&line, &len, fp2);
            token = strtok(line, " ");
            for (int col=0; col<5; col++){
                c[i].board[row][col] = atoi(token);
                token = strtok(NULL, " ");
            }
        }
    }
    free(line);

    for (int k=0; k<3; k++) {
        for (int r=0; r<5; r++) {
            for (int l=0; l<5; l++) {
                printf("%d |", c[k].board[r][l]);
            }
            printf("\n");
        }
    }

    for (int nn=0; nn<wc; nn++) {
        printf(">Marking board for %d number.\n", nums[nn]);
        markBoards(c, nums[nn]);
    }
    
    for (int k=0; k<3; k++) {
        for (int r=0; r<5; r++) {
            for (int l=0; l<5; l++) {
                printf("%d |", c[k].board[r][l]);
            }
            printf("\n");
        }
    }

    fclose(fp);
    fclose(fp2);
    return 0;
}

// mark the board by changing the value to -1
void markBoards (card c[], int num) {
    for (int i=0; i<3; i++) {
        for (int row=0; row<5; row++) {
            for (int col=0; col<5; col++) {
                printf("value: %d ", c[i].board[row][col]);
                printf("number: %d ", num);
                printf("\n-------\n");
                if (num == c[i].board[row][col]) {
                    c[i].board[row][col] = -1;
                    printf("MARKED!\n\n");
                }
            }
            printf("\n");
        }
        printf("End of card #%d checkingwins...\n\n\n", i);
        if (check4win(c[i]) == 0) {
            printf("The board with #%d WINS!!!\n", i);
            exit(EXIT_SUCCESS);
        }
    }
}

int check4win (card c) {
    int row_count =0 , col_count = 0;

    for (int row=0; row<5; row++) {
        row_count = 0;
        col_count = 0;
        for (int col=0; col<5; col++) {
            if (c.board[row][col] == -1) {
                row_count++;
                printf("+1 row\n");
            }
            if (c.board[col][row] == -1) {
                col_count++;
                printf("+1 col\n");
            }
        }
    }
    printf("\nr_c: %d\tc_c: %d\n\n", row_count, col_count);
    if (row_count == 5 || col_count == 5) {
        return 0;
    } else {
        return 1;
    }
}

void nums2array(FILE *fp,  int nums[]) {
    char *line = NULL;
    size_t len =0;
    char *token;
    int n = 0;

    getline(&line, &len, fp);
    token = strtok(line, ",");
    while (token != NULL) {
        nums[n] = atoi(token);
        token = strtok(NULL, ",");
        n++;
    }
    rewind(fp);
}

int getCharCount (FILE *fp) {
    char *line = NULL;
    size_t len = 0;
    char *token;
    int count = 0;

    getline(&line, &len, fp);
    token = strtok(line, ",");
    while (token != NULL) {
        count++;
        token = strtok(NULL, ",");
    }

    rewind(fp);
    return count;
}
