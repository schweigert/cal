#include "stdio.h"
#include "stdlib.h"

int** create(int n){
  int** r = malloc(sizeof(int*)*n);
  int i;
  for(i = 0; i < n; i++){
    r[i] = malloc(sizeof(int)*n);
  }
  return r;
}

int** mull(int** m1, int** m2, int n){
  int i, j, k;
  int** r = create(n);
  for(i = 0; i < n; i++){
      for(j = 0; j < n; j++){
        r[i][j] = 0;
        for(k = 0; k < n; k++){
          r[i][j] += m1[i][k]*m2[k][j];
        }
      }
  }
  return r;
}

void printM(int** m, int n){
  int i, j;
  for(i = 0; i < n; i++){
    for(j = 0; j < n; j++){
      printf("%d  ", m[i][j]);
    }
    printf("\n");
  }
  printf("\n");
}



int main(void) {
  int** a = create(2);
  int** b = create(2);

  a[0][0] = -1;
  a[0][1] = 3;
  a[1][0] = 4;
  a[1][1] = 2;

  b[0][0] = 1;
  b[0][1] = 2;
  b[1][0] = 3;
  b[1][1] = 4;

  printM(a,2);
  printM(b,2);

  int** r = mull(a,b,2);
  printM(r,2);
}
