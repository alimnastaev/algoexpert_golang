defmodule ClassPhotos do
  @moduledoc """
  """

  def run(red_shirt_heights, blue_shirt_heights) do
    red = Enum.sort(red_shirt_heights, &(&1 >= &2))
    blue = Enum.sort(blue_shirt_heights, &(&1 >= &2))

    {red, blue}
    |> find_back_row()
    |> build_rows
  end

  defp find_back_row({[red | r_rest], [blue | b_rest]}) when red < blue,
    do: {:blue, b_rest, r_rest, true}

  defp find_back_row({[_ | r_rest], [_ | b_rest]}),
    do: {:red, r_rest, b_rest, true}

  defp build_rows({_, [], [], result}), do: result

  defp build_rows({:blue, [blue | b_rest], [red | r_rest], true}) when blue > red do
    build_rows({:blue, b_rest, r_rest, true})
  end

  defp build_rows({:blue, _, _, _}), do: build_rows({:blue, [], [], false})

  defp build_rows({:red, [blue | b_rest], [red | r_rest], true}) when blue < red do
    build_rows({:blue, b_rest, r_rest, true})
  end

  defp build_rows({:red, _, _, _}), do: build_rows({:red, [], [], false})
end

case System.argv() do
  ["--test"] ->
    ExUnit.start()

    defmodule ClassPhotosTest do
      use ExUnit.Case

      test "run/2" do
        assert ClassPhotos.run([5, 8, 1, 3, 4], [6, 9, 2, 4, 5]) == true
        assert ClassPhotos.run([5, 8, 2, 3, 4], [6, 9, 2, 4, 1]) == false
      end
    end

  _ ->
    IO.puts(:stderr, "\nplease specify --test")
end
