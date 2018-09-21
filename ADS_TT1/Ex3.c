#include <stdio.h>
#include <math.h>
struct Point {
  float x;
  float y;
};

float EuDistance(struct Point A,struct Point B){
  return sqrt((A.x - B.x)*(A.x - B.x)+(A.y-B.y)*(A.y-B.y));
}

int main(){
  struct Point A;
  struct Point B;
  printf("Point A:\nx = ");
  scanf("%f",&A.x);
  printf("y = ");
  scanf("%f", &A.y);
  printf("Point B:\nx = " );
  scanf("%f",&B.x);
  printf("y = " );
  scanf("%f",&B.y);
  printf("The Euclid distance between two points is %6.2f\n",EuDistance(A,B));
  return 0;
}
