#include <stdio.h>
#include <stdlib.h>

int sumarr(int A[], int size){
  int sum = 0;
  for (int i= 0; i < size; i++){
    sum += A[i];
  }
  return sum;
}

int main(){
  int n;
  printf("How many element do you want to include in your array?  ");
  scanf("%d",&n );
  int *A = malloc(n*sizeof(int));
  for (int i =0; i<n; i++){
    printf("A[%d] = ", i);
    scanf("%d",&A[i]);
  }
  printf("sum of array is %d\n",sumarr(A,n));
  return 0;
}
