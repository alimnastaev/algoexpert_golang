defmodule GenerateDocument do
  def generate_document(_characters, "" = _document), do: true

  def generate_document(characters, document) do
    lookup_table =
      characters
      |> String.graphemes()
      |> Enum.frequencies()

    document
    |> String.graphemes()
    |> Enum.reduce_while({false, lookup_table}, fn str, {_, lt} ->
      case Map.get(lt, str) do
        value when value == 0 or value == nil ->
          {:halt, {false, lt}}

        value ->
          {:cont, {true, Map.put(lt, str, value - 1)}}
      end
    end)
    |> then(fn {r, _} -> r end)
  end
end

case System.argv() do
  ["--test"] ->
    ExUnit.start()

    defmodule GenerateDocumentTest do
      use ExUnit.Case

      test "GenerateDocument.generate_document/2" do
        tests = [
          {"Bste!hetsi ogEAxpelrt x ", "AlgoExpert is the Best!", true},
          {"A", "a", false},
          {"a", "a", true},
          {"a hsgalhsa sanbjksbdkjba kjx", "", true},
          {" ", "hello", false},
          {"     ", "     ", true},
          {"aheaollabbhb", "hello", true},
          {"aheaolabbhb", "hello", false},
          {"estssa", "testing", false},
          {"Bste!hetsi ogEAxpert", "AlgoExpert is the Best!", false},
          {"helloworld ", "hello wOrld", false},
          {"helloworldO", "hello wOrld", false},
          {"helloworldO ", "hello wOrld", true},
          {"&*&you^a%^&8766 _=-09     docanCMakemthisdocument", "Can you make this document &", true},
          {"abcabcabcacbcdaabc", "bacaccadac", true}
        ]

        Enum.each(tests, fn {characters, document, expected} ->
          assert GenerateDocument.generate_document(characters, document) == expected
        end)
      end
    end

  _ ->
    IO.puts(:stderr, "\nplease specify --test flag")
end
