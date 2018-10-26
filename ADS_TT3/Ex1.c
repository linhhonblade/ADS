#include <stdio.h>
#include <stdlib.h>

// define a Node contains a letter
struct Node {
  char c;
  struct Node*next;
};


struct Stack {
  int size;
  struct Node* top;
};

// initialize a stack
struct Stack stack = {0, NULL};

// Push a letter to stack
void Push(struct Node* letter){
  letter->next = stack.top;
  stack.top = letter;
  stack.size++;
  return;
}

// read name from screen and store it to stack
void storeName(){
  char c;
  printf("Enter a name: ");
  while((c = getchar()) != '\n'){
    struct Node* temp = (struct Node*)malloc(sizeof(struct Node));
    temp->c = c;
    Push(temp);
  }
  return;
}

// Pop a letter out from stack
void Pop(){
  struct Node* temp = stack.top;
  stack.top = stack.top->next;
  free(temp);
  stack.size--;
  return;
}

// Print out the name in reverse order
void printRev(){
  while(stack.top!=NULL){
    printf("%c",stack.top->c);
    Pop();
  }
  printf("\n");
  return;
}


int main(){
  storeName();
  printRev();
  return 0;
}
