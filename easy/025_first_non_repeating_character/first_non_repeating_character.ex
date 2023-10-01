defmodule FirstNonRepeatingCharacter do
  @moduledoc """
  """

  def run(string) do
    string
    |> String.graphemes()
    |> Enum.with_index()
    |> build_keyword_list_with_idx()
    |> find_non_repeating_char()
  end

  defp build_keyword_list_with_idx(list_of_chars_with_idx) do
    Enum.reduce(list_of_chars_with_idx, [], fn
      {key, idx}, acc ->
        Keyword.update(acc, String.to_atom(key), {1, idx}, fn {value, _idx} ->
          {value + 1, idx}
        end)
    end)
  end

  defp find_non_repeating_char(keyword_list) do
    Enum.reduce_while(keyword_list, -1, fn
      {_, {1, idx}}, _acc -> {:halt, idx}
      _, acc -> {:cont, acc}
    end)
  end
end

case System.argv() do
  ["--test"] ->
    ExUnit.start()

    defmodule FirstNonRepeatingCharacterTest do
      use ExUnit.Case

      test "run/1" do
        assert FirstNonRepeatingCharacter.run("abcdcaf") == 1
        assert FirstNonRepeatingCharacter.run("faadabcbbebdf") == 6
        assert FirstNonRepeatingCharacter.run("a") == 0
        assert FirstNonRepeatingCharacter.run("ab") == 0
        assert FirstNonRepeatingCharacter.run("abc") == 0
        assert FirstNonRepeatingCharacter.run("abac") == 1
        assert FirstNonRepeatingCharacter.run("ababac") == 5
        assert FirstNonRepeatingCharacter.run("lmnopqldsafmnopqsa") == 7
        assert FirstNonRepeatingCharacter.run("") == -1
        assert FirstNonRepeatingCharacter.run("ggyllaylacrhdzedddjsc") == 10

        assert FirstNonRepeatingCharacter.run(
                 "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxy"
               ) == 25

        assert FirstNonRepeatingCharacter.run(
                 "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"
               ) == -1

        assert FirstNonRepeatingCharacter.run(
                 "aaaaaaaaaaaaaaaaaaaabbbbbbbbbbcccccccccccdddddddddddeeeeeeeeffghgh"
               ) == -1

        assert FirstNonRepeatingCharacter.run("aabbccddeeff") == -1
      end
    end

  _ ->
    IO.puts(:stderr, "\nplease specify --test")
end
