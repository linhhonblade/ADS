#include <stdio.h>
#include <stdlib.h>

//////////////////////// STRUCT /////////////////////////////
struct Edge {
  char startVertice;
  char endVertice;
  float weight;
};

struct Graph {
  int verticesNum;
  int edgesNum;
  char* V;
  struct Edge* E;
};

///////////////////////////// GLOBAL VARIABLES /////////////////////////

//////////////////////////// FUNCTION /////////////////////////////////
struct Edge initEdge(char s, char d, float val){
  struct Edge edge;
  edge.startVertice = s;
  edge.endVertice = d;
  edge.weight = val;
  return edge;
}

struct Graph initGraph(){
  struct Graph myGraph;
  // input vertices
  printf("How many vertices: ");
  scanf("%d",&myGraph.verticesNum);
  myGraph.V = (char*)malloc(myGraph.verticesNum*sizeof(char));
  for(int i = 0;  i < myGraph.verticesNum; i++){
    printf("name of vertice %d: ", i);
    scanf("\n%c",&(myGraph.V[i]));
  }
  // input edges
  char startVertice, endVertice;
  float weight = 0;
  printf("Now, get edges data\nHow many edges? ");
  scanf("%d",&myGraph.edgesNum);
  myGraph.E = (struct Edge*)malloc(myGraph.edgesNum*sizeof(struct Edge));
  for(int i = 0; i < myGraph.edgesNum; i++){
    printf("startVertice: ");
    scanf("\n%c", &startVertice);
    printf("endVertice: ");
    scanf("\n%c",&endVertice );
    printf("weight: ");
    scanf("%f",&weight);
    myGraph.E[i] = initEdge(startVertice, endVertice, weight);
  }
  return myGraph;
}

void showVertices(struct Graph graph){
  for(int i = 0; i < graph.verticesNum; i++){
    printf("%c\t", graph.V[i]);
  }
  printf("\n");
}

void showEdges(struct Graph graph){
  for(int i = 0; i < graph.edgesNum; i++){
    printf("%c <-> %c : %.2f\n",graph.E[i].startVertice, graph.E[i].endVertice, graph.E[i].weight);
  }
}

int asIndex(char c){
  return (int)(c - 'a');
}
// Array based
float* getAdjArr(struct Graph graph){
  float* AdjArr = (float*)malloc(graph.verticesNum*graph.verticesNum*sizeof(float));
  for(int i = 0; i < graph.edgesNum; i++){
    int r = asIndex(graph.E[i].startVertice);
    int c = asIndex(graph.E[i].endVertice);
    AdjArr[r*graph.verticesNum+c] = graph.E[i].weight;
    AdjArr[c*graph.verticesNum+r] = graph.E[i].weight; // undirect graph
  }
  return AdjArr;
}

void showAdjArr(struct Graph graph, float* AdjArr){
  for(int i = 0; i < graph.verticesNum; i++){
    for(int j = 0; j < graph.verticesNum; j++){
      printf("%.1f\t",AdjArr[i*graph.verticesNum+j]);
    }
    printf("\n");
  }
  printf("\n");
}





int main(){
  struct Graph myGraph = initGraph();
  showVertices(myGraph);
  showEdges(myGraph);
  float* Arr = getAdjArr(myGraph);
  showAdjArr(myGraph, Arr);
  return 0;
}
