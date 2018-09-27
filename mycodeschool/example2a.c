#include <stdio.h>
#include <stdlib.h>

struct Node{
  int data;
  struct Node* next;
};

struct Node* head;

void Insert(int x, int index){
  struct Node* temp1 = (struct Node*)malloc(sizeof(struct Node));
  temp1->data = x;
  temp1->next = NULL;
  if(index ==1){
    temp1->next = head;
    head = temp1;
    return;
  }
  struct Node* temp2 = head;
  for (int i =0; i<index-2; i++){
  temp2 = temp2->next;
  }
  temp1->next = temp2->next;
  temp2->next = temp1;
}

void Print(){
  struct Node *temp = head;
  while(temp != NULL){
    printf("%d ",temp->data );
    temp = temp->next;
  }
  printf("\n");
}

int main(){
  head = NULL;
  Insert(4,1);
  Insert(3,2);
  Insert(5,3);
  Insert(1,2);
  Print();
}
