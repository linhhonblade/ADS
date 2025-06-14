def lexi_sort(arr):
    mem = [[False, False] for _ in range(len(arr))]
    mem[0][0] = True
    mem[0][1] = True
    trace = [[None, None] for _ in range(len(arr))]
    for i in range(1, len(arr)):
        # Update current trace
        pre0 = arr[i-1]
        pre1 = arr[i-1][::-1]
        curr0 = arr[i]
        curr1 = arr[i][::-1]
        if mem[i-1][0]:
            if curr0 >= pre0:
                mem[i][0] = True # curr0 possible
                trace[i][0] = 0 # curr0 possible with pre0
            if curr1 >= pre0:
                mem[i][1] = True # curr1 possible
                trace[i][1] = 0 # curr1 possible with pre0
        if mem[i-1][1]:
            if curr0 >= pre1:
                mem[i][0] = True # curr0 possible
                if trace[i][0] is None:
                    trace[i][0] = 1 # curr0 possible with pre1
            if curr1 >= pre1:
                mem[i][1] = True # curr1 possible
                if trace[i][1] is None:
                    trace[i][1] = 1 # curr1 possible with pre1
        # Reconstruct the result
    result = [0] * len(arr)
    curr = 0 if mem[i][0] else 1

    for i in range(len(arr)-1, -1, -1):
        result[i] = curr
        curr = trace[i][curr]
    print(result)
    return "".join(map(str, result))


def test_lexi_sort():
    with open('input.txt', 'r') as f:
        lines = f.readlines()[1:]
        input_data = []
        while lines:
            n = int(lines.pop(0).strip())
            arr = []
            while n > 0:
                arr.append(lines.pop(0).strip())
                n -= 1
            input_data.append(arr)
        print(input_data)
    with open('output.txt', 'r') as f:
        lines = f.readlines()
        output_data = [line.strip() for line in lines]
        print(output_data)
    i = 0
    j = 0
    while i < len(input_data) and j < len(output_data):
        print(i, j)
        print(input_data[i], output_data[j])
        assert lexi_sort(input_data[i]) == output_data[j]
        i += 1
        j += 1