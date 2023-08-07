defmodule CaesarCipherEncryptor do
  @alph_start 96
  @alph_end 122

  def run(str, key) do
    key = rem(key, 26)

    str
    |> String.to_charlist()
    |> Enum.reduce('', &[build_new_char(&1, key) | &2])
    |> Enum.reverse()
    |> :unicode.characters_to_binary()
  end

  defp build_new_char(char, key) do
    char = char + key
    offset = char - @alph_end

    if offset > 0, do: @alph_start + offset, else: char
  end
end

case System.argv() do
  ["--test"] ->
    ExUnit.start()

    defmodule RunLengthEncodingTest do
      use ExUnit.Case

      test "CaesarCipherEncryptor.run/2" do
        tests = [{"xyz", 2, "zab"}, {"abc", 57, "fgh"}]

        Enum.each(tests, fn {str, key, expected} ->
          assert CaesarCipherEncryptor.run(str, key) == expected
        end)
      end
    end

  _ ->
    IO.puts(:stderr, "\nplease specify --test flag")
end
