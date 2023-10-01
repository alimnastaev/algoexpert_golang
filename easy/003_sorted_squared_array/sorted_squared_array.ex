defmodule SortedSquaredArray do
  def sorted_squared_array(list) do
    [map, length] = build_map_with_indexed(list)

    build_result(map, 1, length - 1, [])
  end

  defp build_map_with_indexed(list) do
    Enum.reduce(list, [%{}, 1], fn key, [map, idx] ->
      [Map.put(map, idx, key), idx + 1]
    end)
  end

  defp build_result(map, _, _, acc) when map == %{}, do: acc

  defp build_result(map, left_idx, length, acc) do
    if abs(map[left_idx]) > abs(map[length]) do
      squared_number = map[left_idx] * map[left_idx]

      map
      |> Map.delete(left_idx)
      |> build_result(left_idx + 1, length, [squared_number | acc])
    else
      squared_number = map[length] * map[length]

      map
      |> Map.delete(length)
      |> build_result(left_idx, length - 1, [squared_number | acc])
    end
  end
end

case System.argv() do
  ["--test"] ->
    ExUnit.start()

    defmodule SortedSquaredArrayTest do
      use ExUnit.Case

      test "sorted_squared_array/1" do
        assert SortedSquaredArray.sorted_squared_array([-7, -5, -4, 3, 6, 8, 9]) ==
                 [9, 16, 25, 36, 49, 64, 81]

        assert SortedSquaredArray.sorted_squared_array([-4, -1]) == [1, 16]
        assert SortedSquaredArray.sorted_squared_array([2]) == [4]
        assert SortedSquaredArray.sorted_squared_array([-2]) == [4]
      end
    end

  _ ->
    IO.puts(:stderr, "\nplease specify --test")
end
