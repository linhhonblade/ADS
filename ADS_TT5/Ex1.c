#include <stdio.h>
#include <stdlib.h>
#include <time.h>

//function to generate aa array with random value
int* makeArr(int size, int low, int high){
  int i;
  int* A = (int*)malloc(size*sizeof(int));
  for(i=0; i<size; i++){
    A[i] = low + (rand()%((high-low)+1));
  }
  return A;
}

//function to print all values in an array
void Print(int* A, int n){
  int i = 0;
  while(i < n){
    printf("%d ",A[i]);
    i++;
  }
  printf("\n");
}

//function to swap two value with given index in an array
void swap(int* a, int* b){
  *a = *a + *b;
  *b = *a - *b;
  *a = *a - *b;
}

//function for reversible bubble sort
void revBubble(int* A, int left, int right){
  //base case
  if(left == right)
    return;
  int i, j, k;
  //sort forward
  for(i = left; i<right; i++){
    if(A[i] > A[i+1]){
      swap(&A[i], &A[i+1]);
    }
  }
  //sort backward
  if(left == right-1) return;
  for(i = right - 1; i > left; i--){
    if(A[i] < A[i+1]){
      swap(&A[i], &A[i+1]);
    }
  }
  //recur the sort
  revBubble(A, left+1, right-1);
}

int main(){
  srand(time(NULL));
  int size = 10;
  int low = 0;
  int high = 100;
  int* A = makeArr(size, low, high);
  printf("Initial Array:\n");
  Print(A, size);
  clock_t begin = clock();
  revBubble(A, 0, size-1);
  clock_t end = clock();
  printf("Sorted Array:\n");
  Print(A, size);
  printf("Time takec: %lfs\n",(double)(end-begin)/CLOCKS_PER_SEC);
  return 0;
}

/*
##################### COMPLEXITY ###########################
- In the initial bubble sort, each recursion return one element
sorted. In this reversible bubble sort, each recursion gives
two elements sorted (one the largest and one the smalles).
However, each recursion has two loops (one is forward sorting
and the other is backward sorting) instead one loop. therefore,
the number of recursion needed reduces by half. As a result,
the time complexity remains the same which is O(n^2).
- Space complexity is O(1) because all changes are made on the
initial array.
*/
