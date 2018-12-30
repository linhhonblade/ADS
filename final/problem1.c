#include <stdio.h>
#include <math.h>

float t = 1.5;
int k = 2;

float abso(float a, float b){
  if (a-b >= 0){
    return a-b;
  }
  else{
    return b-a;
  }
}


void findSimilar(float A[], int index, int length){
  //base case
  if((length - index) == 2){
    if(abso(A[index],A[index+1]) <= t){
      printf("(%.1f, %.1f)\n",A[index], A[index+1] );
      return;
    }
    else{
      return;
    }
  }
  else{
    int i ;
    int count = 0;
    for(i = index +1; i<length; i++){
      if(abso(A[i],A[index]) <= t){
        printf("(%.1f, %.1f)\n",A[index], A[i]);
        count ++;
      }
    }
    if(count==k){
      printf("%.1f is perfectly similar value with k = %d\n",A[index], k );
    }
    findSimilar(A, index +1, length);
  }
}

int main(){
  float A[6] = {3.5, 8.0, 4.6, 4.3, 7.5, 8.3};
  int lengthA = sizeof(A)/sizeof(A[0]);
  printf("length = %d\n" ,lengthA);
  findSimilar(A, 0, lengthA);
}

///////////////////////////// COMPLEXITY ////////////////////////////
// for finding the similar values of the first element of the array
// it took n-1 iteration
// for the second element, it took n - 2
// ...
// the time complexity is (n-1) + (n-2) + (n-3) + ... ~ O(^2)
// space complexity is O(1)
