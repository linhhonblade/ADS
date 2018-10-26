#include <stdio.h>
#include <stdlib.h>
#include <time.h>

//define a node contains the number
struct Node {
  int num;
  struct Node* next;
};

//a pointer to top of the stack
struct Node* top = NULL;

//put a number to stack
void Push(){
  int x = rand()%10;
  if((top==NULL)||((top!=NULL)&&(x!=top->num))){
    struct Node* temp = (struct Node*)malloc(sizeof(struct Node));
    temp->num = x;
    temp->next = top;
    top = temp;
    return;
  }
  else{
    Push();
  }
}

//get a number out of the stack
void Pop(){
  struct Node*temp = top;
  top = top->next;
  free(temp);
  return;
}

//check if the player win or not
int isWin(){
  int x;
  printf("Enter an integer from 0 to 9: ");
  scanf("%d",&x );
  return(((top->num > x) && (top->next->num < x))||((top->num < x)&&(top->next->num > x)));
}

//print out the numbers that had been chosen (just for checking)
void Print(){
  printf("%d %d\n", top->num, top->next->num);
  return;
}

int main(){
  srand(time(NULL));
  Push();
  Push();
  if(isWin()){
    Print();
    printf("YOU WIN!\n");
    return 0;
  }
  else{
    Pop();
    Push();
    if(isWin()){
      Print();
      printf("YOU WIN!\n");
      return 0;
    }
    else{
      Print();
      Pop();
      Pop();
      printf("YOU FAILED!\n");
      return 1;
    }
  }
}
