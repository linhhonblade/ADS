#include <stdio.h>
#include <time.h>

int calc_burnt(int candles, int makeNew) {
    int burnt = candles;    // number of candles has burnt
    int leftovers = candles;    // number of leftovers in next turn
    int to_burn = 0;    // number of candle will form by leftovers in next turn
    while (leftovers >= makeNew) {
        to_burn = leftovers / makeNew;
        leftovers = to_burn + leftovers % makeNew;
        burnt += to_burn;
    }
    return burnt;
}

int main() {
    int result;
    clock_t begin = clock();
    result = calc_burnt(200, 4);
    clock_t end = clock();
    double time_spent = (double)(end - begin)/CLOCKS_PER_SEC;
    printf("return %d in %f secs\n", result, time_spent);
    return 0;
}