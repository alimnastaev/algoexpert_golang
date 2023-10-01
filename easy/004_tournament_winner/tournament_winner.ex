defmodule TournamentWinner do
  @moduledoc """
  There's an algorithms tournament taking place in which teams of programmers
  compete against each other to solve algorithmic problems as fast as possible.
  Teams compete in a round robin, where each team faces off against all other
  teams. Only two teams compete against each other at a time, and for each
  competition, one team is designated the home team, while the other team is the
  away team. In each competition there's always one winner and one loser; there
  are no ties. A team receives 3 points if it wins and 0 points if it loses. The
  winner of the tournament is the team that receives the most amount of points.

  Given an array of pairs representing the teams that have competed against each
  other and an array containing the results of each competition, write a
  function that returns the winner of the tournament.
  The input arrays are named `competitions` and `results`, respectively.
  The `competitions` array has elements in the form of `[homeTeam, awayTeam]`,
  where each team is a string of at most 30 characters representing the name of the team.
  The `results` array contains information about the winner of each corresponding competition
  in the `competitions` array. Specifically, `results[i]` denotes the winner of `competitions[i]`,
  where a `1` in the `results` array means that the home team in the corresponding
  competition won and a `0` means that the away team won.

  ## SAMPLE INPUT:
    competitions = [["HTML", "C#"], ["C#", "Python"], ["Python", "HTML"]]
    results =  = [0, 0, 1]

  ## SAMPLE OUTPUT: "Python"

    # C# beats HTML, Python Beats C#, and Python Beats HTML.
    # HTML - 0 points
    # C# -  3 points
    # Python -  6 points
  """

  def run(competitions, results) do
    competitions
    |> Enum.zip(results)
    |> map_score()
    |> find_winner()
  end

  defp map_score(list_of_tuples) do
    Enum.reduce(list_of_tuples, %{}, fn {[home_team, away_team], winner}, acc ->
      winner_team = if winner == 0, do: away_team, else: home_team

      Map.update(acc, winner_team, 3, &(&1 + 3))
    end)
  end

  defp find_winner(map) do
    Enum.reduce(map, {"", 0}, fn
      {team, score}, {_, highest_score} when score > highest_score -> {team, score}
      _, acc -> acc
    end)
    |> then(fn {winner, _} -> winner end)
  end
end

case System.argv() do
  ["--test"] ->
    ExUnit.start()

    defmodule TournamentWinnerTest do
      use ExUnit.Case

      test "run/1" do
        assert TournamentWinner.run(
                 [["HTML", "C#"], ["C#", "Python"], ["Python", "HTML"]],
                 [0, 0, 1]
               ) == "Python"
      end
    end

  _ ->
    IO.puts(:stderr, "\nplease specify --test")
end
