#include <stdio.h>
#include <math.h>
int checkdigit(char A[9]){
  int sum = 0;
  for(int i = 0; i<9; i++)
    sum += (i+1)*((int)A[i]-48);      // convert char type to integer type
  return fmod(sum, 11);
}

int main(){
  char A[9];
  printf("Enter first 9 digits of your ISBN code:\n");
  scanf("%s",A);
  printf("The check digit is %d\n",checkdigit(A) );
}
