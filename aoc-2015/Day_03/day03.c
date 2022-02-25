#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// coordinates struct
struct coord {
  int x;
  int y;
};

int exists_in(struct coord cur, struct coord visited[], int i);
void init_santas(struct coord *cur, int santas);
int find_visited(FILE **fp, int santas);

int main (int argc, char *argv[]) {

    FILE *fp; 
    fp = fopen(argv[1], "r");

    if (argc != 2) {
        perror("First argument need to be the input text file!\n No more no less!\n");
        return 1;
    }
    else if (fp == NULL) {
        perror("Unable to open file!");
        return 1;
    }

    int i = find_visited(&fp, 1);
    printf("Part1>> Count of locations visited at least once: %d\n", i);
    rewind(fp);
    int i2 = find_visited(&fp, 2);
    printf("Part2>> Count of locations visited by all santas at least once: %d\n", i2);

    fclose(fp);
    return 0;
}

// TODO: fix initialization for multiple santas
int find_visited (FILE **fp, int santas) {
  struct coord visited[10000];
  struct coord cur[santas];
  int visited_count = 1;
  char ch;
  visited[0].x = 0;
  visited[0].y = 0;
  int robo_santas = 0;
  init_santas(&cur[santas], santas);

  while ((ch = fgetc(*fp)) != EOF) {
    if (ch == '^') {
      cur[robo_santas].y += 1;
    }
    else if (ch == 'v') {
      cur[robo_santas].y -= 1;
    }
    else if (ch == '>') {
      cur[robo_santas].x += 1;
    }
    else if (ch == '<') {
      cur[robo_santas].x -= 1;
    }
    if (exists_in(cur[robo_santas], visited, visited_count) == 1) {
      visited[visited_count].x = cur[robo_santas].x;
      visited[visited_count].y = cur[robo_santas].y;
      visited_count++;
    }
    robo_santas++;
    robo_santas %= santas;
  }

  return visited_count;
}

void init_santas(struct coord *cur, int santas) {
  for (int i=0;i<santas;i++) {
    cur[i].x = 0;
    cur[i].y = 0;
  }
}

// checks if current coordinates exist in the visited coordinate array
int exists_in (struct coord cur , struct coord visited[], int i) {
  for (int j=0;j<i;j++) {
    if (cur.x == visited[j].x && cur.y == visited[j].y) {
      return 0;
    }
  }
  return 1;
}
