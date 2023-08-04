defmodule RunLengthEncoding do
  @target 9

  def run(str), do: do_run(str, 1, "", "")

  # base case
  defp do_run("", c, prvChar, output), do: "#{output}#{c}#{prvChar}"

  defp do_run(<<char::binary-size(1)>> <> tail, c, "", output),
    do: do_run(tail, c, "#{char}", output)

  defp do_run(<<char::binary-size(1)>> <> tail, c, prvChar, output)
       when c == @target or char != prvChar,
       do: do_run(tail, 1, "#{char}", "#{output}#{c}#{prvChar}")

  defp do_run(<<_>> <> tail, c, prvChar, output),
    do: do_run(tail, c + 1, prvChar, output)
end

case System.argv() do
  ["--test"] ->
    ExUnit.start()

    defmodule RunLengthEncodingTest do
      use ExUnit.Case

      test "RunLengthEncoding.run/2" do
        tests = [
          {"AAAAAAAAAAAAABBCCCCDD", "9A4A2B4C2D"},
          {"************^^^^^^^$$$$$$%%%%%%%!!!!!!AAAAAAAAAAAAAAAAAAAA", "9*3*7^6$7%6!9A9A2A"},
          {"........______=========AAAA   AAABBBB   BBB", "8.6_9=4A3 3A4B3 3B"}
        ]

        Enum.each(tests, fn {input, expected} ->
          assert RunLengthEncoding.run(input) == expected
        end)
      end
    end

  _ ->
    IO.puts(:stderr, "\nplease specify --test flag")
end
