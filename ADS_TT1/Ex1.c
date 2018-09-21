#include <stdio.h>
int sum(int*a, int*b){
  return *a + *b;
}

int main(){
  int a, b;
  printf("Exercise#1:\nFind sum off two integer a and b\n");
  printf("a = ");
  scanf("%d",&a);
  printf("b = ");
  scanf("%d",&b);
  printf("%d + %d = %d\n",a, b, sum(&a, &b));
  return 0;
}
