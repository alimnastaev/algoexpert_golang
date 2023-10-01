defmodule NonConstructibleChange do
  @moduledoc """
  """

  def run(coins) do
    coins
    |> Enum.sort()
    |> Enum.reduce(0, fn
      c, result when c > result + 1 -> result + 1
      c, result -> result + c
    end)
  end
end

case System.argv() do
  ["--test"] ->
    ExUnit.start()

    defmodule TournamentWinnerTest do
      use ExUnit.Case

      test "run/1" do
        assert NonConstructibleChange.run([1, 2, 5]) == 4
        assert NonConstructibleChange.run([5, 7, 1, 1, 2, 3, 22]) == 20
      end
    end

  _ ->
    IO.puts(:stderr, "\nplease specify --test")
end
