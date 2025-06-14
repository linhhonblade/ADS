def toy_race(params, V):
  max_speed = max(V)
  me_idx = params[0] - 1
  if V[me_idx] == max_speed:
    return 0
  min_race_time = params[1] / max_speed
  # Z = (V*min_race_time - X) / (V-1)
  z = params[1] - V[me_idx] * (min_race_time-1)
  if z < params[2]:
     return int(z + 1)
  else:
     return -1

def test_toy_race():
    with open('input.txt', 'r') as f:
        lines = f.readlines()[1:]
        input_data = [list(map(int, line.split())) for line in lines]
        print(input_data)
    with open('output.txt', 'r') as f:
        lines = f.readlines()
        output_data = [list(map(int, line.split())) for line in lines]
        print(output_data)
    i = 0
    j = 0
    while i < len(input_data) and j < len(output_data):
      print(i, j)
      assert toy_race(input_data[i], input_data[i+1]) == output_data[j][0]
      i += 2
      j += 1