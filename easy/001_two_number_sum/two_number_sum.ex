defmodule TwoNumberSum do
  @moduledoc """
  """

  def run(array, target_sum) do
    Enum.reduce_while(array, {[], MapSet.new()}, fn int, {result, m} ->
      corresponding_number = target_sum - int

      if MapSet.member?(m, corresponding_number) do
        {:halt, {[int, corresponding_number | result], m}}
      else
        {:cont, {result, MapSet.put(m, int)}}
      end
    end)
    |> then(fn {result, _} -> result end)
  end
end

case System.argv() do
  ["--test"] ->
    ExUnit.start()

    defmodule TwoNumberSumTest do
      use ExUnit.Case

      test "run/1" do
        assert TwoNumberSum.run([3, 5, -4, 8, 11, 1, -1, 6], 10) == [-1, 11]
      end
    end

  _ ->
    IO.puts(:stderr, "\nplease specify --test")
end
