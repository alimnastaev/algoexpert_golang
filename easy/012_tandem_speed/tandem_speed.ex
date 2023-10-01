defmodule TandemSpeed do
  @moduledoc """
  """

  def run(red_shirt_speeds, blue_shirt_speeds, fastest) do
    order = if fastest == true, do: &(&1 >= &2), else: &(&1 <= &2)

    red = Enum.sort(red_shirt_speeds, order)
    blue = Enum.sort(blue_shirt_speeds)

    Enum.zip_reduce([red, blue], 0, fn
      [l, r], acc -> acc + max(l, r)
    end)
  end
end

case System.argv() do
  ["--test"] ->
    ExUnit.start()

    defmodule TandemSpeedTest do
      use ExUnit.Case

      test "run/2" do
        assert TandemSpeed.run([5, 5, 3, 9, 2], [3, 6, 7, 2, 1], true) == 32
        assert TandemSpeed.run([5, 5, 3, 9, 2], [3, 6, 7, 2, 1], false) == 25
      end
    end

  _ ->
    IO.puts(:stderr, "\nplease specify --test")
end
