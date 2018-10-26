#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#define CAPACITY 20

// define a Node contains the name of a passenger
struct Node {
  char name[CAPACITY];
  struct Node*next;
};

// define a quque
struct Queue {
  struct Node *front, *back;
  int size;
};

// initialize a queue
struct Queue queue = {NULL, NULL, 0};

// put a new passenger to queue
void enQueue(char name[]){
  struct Node* temp = (struct Node*)malloc(sizeof(struct Node));
  strcpy(temp->name, name);
  temp->next = NULL;
  if(queue.size == 0){
    queue.back = temp;
    queue.front = temp;
    queue.size++;
    return;
  }
  queue.back->next = temp;
  queue.back = temp;
  queue.size++;
  return;
}

// get a passenger out (get into the train)
void deQueue(){
  if(queue.size == 0){
    printf("There isn't any passenger waitting\n");
    return;
  }
  struct Node*temp = queue.front;
  queue.front = queue.front->next;
  free(temp);
  queue.size--;
  return;
}

// show the queue (just for checking)
void printQueue(){
  struct Node* temp = queue.front;
  if(temp==NULL){
    printf("There isn't passenger in queue\n");
  }
  while(temp != NULL){
    printf("%s\n",temp->name);
    temp = temp->next;
  }
  printf("\n");
  return;
}

int main(){
  enQueue("misa");printQueue();
  enQueue("mai dien");printQueue();
  enQueue("hoa ngo");printQueue();
  deQueue();printQueue();
  deQueue();printQueue();
  deQueue();printQueue();
  deQueue();printQueue();
  return 0;
}
