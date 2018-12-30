// this problem has not done yet!!!

#include <stdio.h>
#include <stdlib.h>

struct Node{
  float a;
  struct Node* next;
};

struct Stack {
  struct Node* top;
};

struct Stack stack = {NULL};
float t = 1.5;

float abso(float a, float b){
  if (a-b >= 0){
    return a-b;
  }
  else{
    return b-a;
  }
}

void Pop(){
  struct Node* temp = stack.top;
  stack.top = temp->next;
  free(temp);
}


void checkStack(){
  if (abso(stack.top->a, stack.top->next->a) <= t){
    printf("(%.1f, %.1f)\n",stack.top->next->a, stack.top->a);
  }
  else{
   Pop();
  }
}




void Push(float A[], int index){
  struct Node* node = (struct Node*)malloc(sizeof(struct Node));
  node->a = A[index];
  node->next = stack.top;
  stack.top = node;
}

void findSimilar(float A[], int index, int lengthA){
  if(index == lengthA-1){
    return;
  }
  
  Push(A, index);
  Push(A, index+1);
  checkStack();
  int i;
  for(i = index +2; i < lengthA; i++ ){
    Push(A, i);
    checkStack();
  }
  Pop();
  findSimilar(A, index+1, lengthA);
}



int main(){
  float A[6] = {3.5, 8.0, 4.6, 4.3, 7.5, 8.3};
  findSimilar(A, 0, sizeof(A)/sizeof(A[0]));
}



///////////////////////////// COMPLEXITY //////////////////////
// it is the same like problem 1, at the kth recursion, there are n - k iteration
// so the time complexity is O(n^2)
// space complexity is O(1) ,

