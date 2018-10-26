#include <stdio.h>

double multi(int a, int b){
  int i = 0;
  // case condition
  if(b==0)
    return 0;
  if(b < 0){
    for(i = 0; i < -b; i++)
      return multi(a, b+1) - a;
  }
  for(i = 0; i < b; i++){
    return a+multi(a, b-1);
  }
}

int main(){
  int a, b;
  printf("Enter the first factor: ");
  scanf("%d",&a );
  printf("Enter the second factor: ");
  scanf("%d",&b );
  printf("%f\n",multi(a, b) );
  return 0;
}
