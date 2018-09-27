
#include <stdio.h>
#include <stdlib.h>

#define CAPACITY 10

struct Car {
  int num_pass;
  char name[CAPACITY];
  struct Car* next;
};        // define a car

struct Car* head;        // initialize a head
int size;

void AddCar(char name[], int num_pass){
  struct Car* temp1 = (struct Car*)malloc(sizeof(struct Car));
  int i = 0;
  while(name[i] != '\0'){
      temp1->name[i] = name[i];
      i++;
  }//read and save the name of the car
  temp1->num_pass = num_pass;
  if (head == NULL){
    temp1->next = head;
    head = temp1;
    size++;
    return;
  }
  temp1->next = head;
  head = temp1;
  size++;
  return;
}

void displayAll(){
  struct Car* temp = head;
  while(temp != NULL){
    int i =0;
    while(temp->name[i] != '\0'){
      printf("%c", temp->name[i]);
      i++;
    }
    printf(": %d\n",temp->num_pass );
    temp = temp->next;
  }
  printf("The train contains %d cars.\n", size);
}

void rmEmpty(){
  struct Car* temp1 = head;
  struct Car* temp2 = NULL;
  while(temp1 != NULL){
    if(temp1->num_pass==0){
      temp2->next = temp1->next;
      free(temp1);
      size--;
    }
    temp2 = temp1;
    temp1 = temp1->next;
  }
}

int main(){
  head = NULL;
  AddCar("car1", 3);
  AddCar("car2", 2);
  AddCar("misa", 0);
  AddCar("car3", 6);
  AddCar("car4", 5);
  AddCar("caremp", 0);
  AddCar("misakute", 2);
  displayAll();
  rmEmpty();
  displayAll();
}
