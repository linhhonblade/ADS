#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <math.h>

int NumberOfElement = 0;
int HeightOfTree = 0;
int NumberOfLeaves = 0;

struct bstNode {
  int data;
  int index;
  struct bstNode* left;
  struct bstNode* right;
};

struct bstNode* newbstNode(int data, int index){
  struct bstNode* node = (struct bstNode*)malloc(sizeof(struct bstNode));
  node->data = data;
  node->index = index;
  node->left = NULL;
  node->right = NULL;
  return node;
}

int Search(struct bstNode* root, int data){
  if(root == NULL) return -1;
  else if(root->data == data) return root->index;
  else if(data < root->data) return Search(root->left, data);
  else return Search(root->right, data);
}

void preOrder(struct bstNode* root){
  if(root == NULL) return;
  printf("%d ",root->data);
  preOrder(root->left);
  preOrder(root->right);
}

int Height(int num_element){
  return ceil(log(num_element)*M_LOG2E);
}

int getMid(int left, int right){
  return ceil(left + (right-left)/2);
}

struct bstNode* MakeTreeFromArr(int A[], int left, int right){
  int mid = getMid(left, right);
  struct bstNode* node = newbstNode(A[mid], mid);
  if(mid == left){
    node->left = newbstNode(A[left], left);
    node->right = newbstNode(A[right], right);
  }
  else{
    node->left = MakeTreeFromArr(A, left, mid);
    node->right = MakeTreeFromArr(A, mid+1, right);
  }
  return node;
}
void showAllLeaves(int A[], int n){
  int i = 0;
  while(i < n){
    printf("%d ",A[i]);
    i++;
  }
  printf("\n");
  printf("NumberOfLeaves = %d\nNumberOfElement = %d\nHeight = %d\n",NumberOfLeaves, NumberOfElement,HeightOfTree );
}
int* initDat(){
  int i = 0;
  int* A = (int*)malloc(NumberOfLeaves*sizeof(int));
  while(i < NumberOfLeaves){
    if(i < NumberOfElement){
      printf("A[%d] = ", i);
      scanf("%d",&A[i]);
    }
    else{
      A[i] = A[NumberOfElement-1];
    }
    i++;
  }
  return A;
}
void AddToArr(int A[], int data, int index){
  int i;
  if(index == NumberOfElement){
    for(i = NumberOfElement; i < NumberOfLeaves; i++){
      A[i] = data;
    }
    NumberOfElement++;
    return;
  }
  for(i = NumberOfElement-1; i > index; i--){
    A[i] = A[i-1];
  }
  A[index] = data;
  NumberOfElement++;
  return;
}
int* AddToDat(int A[], int new){
  int i;
  //if(NumberOfElement < NumberOfLeaves){NumberOfElement++;}
    // add new element to tree does not requires increment in height
    // NumberOfElement increase by 1
    // NumberOfLeaves does not change
  if(NumberOfElement == NumberOfLeaves){
    // add new element to tree requires increment in height of the tree
    // create a new array to store data of new tree
    // NumberOfElement increase by 1
    // NumberOfLeaves is doubled
    A = (int*)realloc(A, 2*NumberOfLeaves*sizeof(int)); // resize array
    for(i = NumberOfElement; i < 2*NumberOfLeaves; i++){
      A[i] = A[i-1];
    }
    NumberOfLeaves*=2;
    HeightOfTree++;
  }
  i = 0;
  if(A[0] > new){
    AddToArr(A, new, 0);
  }
  else if(A[NumberOfElement] < new){
    AddToArr(A, new, NumberOfElement);
  }
  else{
    i = 1;
    while(i < NumberOfElement){
      if(A[i] > new && new > A[i-1]){
        AddToArr(A, new, i);
        break;
      }
      i++;
    }
  }
  return A;
}

bool isEmpty(struct bstNode* root){
  if(root == NULL) return true;
  else return false;
}

void RemoveFromDat(int A[], int data, struct bstNode* root){
  int i;
  if(data == A[NumberOfElement-1]){
    for(i = NumberOfElement-1;i < NumberOfLeaves; i++){
      A[i] = A[i-1];
    }
    NumberOfElement--;
  }
  else if(Search(root, data) == -1){
    printf("Item not found\n");
    return;
  }
  else{
    for(i = Search(root, data); i < NumberOfElement; i++){
      A[i] = A[i+1];
    }
    NumberOfElement--;
  }
  if(NumberOfLeaves/2 == NumberOfElement){
    A = (int*)realloc(A, (NumberOfLeaves/2)*sizeof(int));
    NumberOfLeaves/=2;
    HeightOfTree--;
  }
  return;
}


int main(){
  // init an empty tree
  struct bstNode* root = NULL;
  // get data before building tree
  printf("How many elements in your data? \n ");
  scanf("%d",&NumberOfElement);
  HeightOfTree = Height(NumberOfElement);
  NumberOfLeaves = pow(2, HeightOfTree);
  // strore data in an array before building tree
  int* A = initDat();

  // show all the leaves of the tree
  printf("All leaves:\n" );
  showAllLeaves(A, NumberOfLeaves);

  // build tree from data stored in array
  root = MakeTreeFromArr(A, 0, NumberOfLeaves-1);

  // show tree in preOrder
  printf("Tree in PreOrder:\n");
  preOrder(root);
  printf("\n");

  //////// ADD NEW VALUE TO TREE AND CHECK /////////
  // get value to add
  int added_value;
  printf("Enter the value you want to add to tree: ");
  scanf("%d",&added_value );
  // add value to the array
  A = AddToDat(A, added_value);
  // rebuild the tree from array after adding new value
  root = MakeTreeFromArr(A, 0, NumberOfLeaves-1);
  // show the leaves of new tree (after adding new value)
  printf("All the leaves after adding %d:\n",added_value);
  showAllLeaves(A, NumberOfLeaves);
  // show new tree in preOrder
  printf("Tree after adding value %d in PreOrder:\n", added_value);
  preOrder(root);
  printf("\n");
  //////////////////////////////////////////////////
  //////// REMOVE A VALUE FROM TREE AND CHECK //////
  int value_remove;
  printf("Enter the value you want to remove: ");
  scanf("%d",&value_remove );
  // remove value from array
  RemoveFromDat(A, value_remove, root);
  if(Search(root, value_remove) != -1){
    // rebuild the tree after removing a value
    root = MakeTreeFromArr(A, 0, NumberOfLeaves-1);
  }
  // show all the leaves of new tree (after removing a value)
  printf("All the leaves after removing %d:\n",value_remove );
  showAllLeaves(A, NumberOfLeaves);
  // show new tree in PreOrder
  printf("Tree after removing value %d in PreOrder:\n",value_remove );
  preOrder(root);
  printf("\n");
  ///////////////////////////////////////////////////
  return 0;
}
