#include <stdio.h>
#include <math.h>

struct Point {
  float x;
  float y;
};             // define struct Point

char checkRect(struct Point A, struct Point B){
  if (A.x == B.x || A.y == B.y)
    return 'F';
  return 'T';
}             // define function to check the rectangle

void inputPoint(struct Point *A){
  printf("Point:\nx = " );
  scanf("%f",&A->x );
  printf("y = " );
  scanf("%f",&A->y );
}         // define function to input value for two points


int checkPoint(struct Point *C, struct Point *A, struct Point *B){
  if (fabs(C->x-B->x)<fabs(A->x-B->x) && fabs(C->y-B->y)<fabs(A->y-B->y))
    return 1;
  return 0;
}           // define function to check whether a point falls into the rectangle


int main(){
  struct Point A ;
  struct Point B ;
  struct Point C ;
  inputPoint(&A);
  inputPoint(&B);

  while (checkRect(A, B)=='F'){
    printf("These points cannot construct a rectangle\nTry another set of points.\n");
    inputPoint(&A);
    inputPoint(&B);
  }
   printf("These points can construct a rectangle\nThe area of this rectangle is %6.2f.\nPlease input another point.\n", fabs((A.x - B.x)*(A.y-B.y)));
   inputPoint(&C);
   printf("%d\n",checkPoint(&C, &A, &B));
  return 0;
}
