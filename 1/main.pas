program BubbleSortApp;
{$mode objfpc}{$H+}

uses
  sysutils;

type
  TIntArray = array of integer;

var
  Numbers: TIntArray;

// 4.0: Dodanie parametrów (minVal, maxVal, count)
procedure GenerateNumbers(var arr: TIntArray; minVal, maxVal, count: integer);
var i: integer;
begin
  SetLength(arr, count);
  Randomize;
  for i := 0 to count - 1 do
    arr[i] := minVal + Random(maxVal - minVal + 1);
end;

procedure BubbleSort(var arr: TIntArray);
var i, j, temp, n: integer;
begin
  n := Length(arr);
  for i := 0 to n - 2 do
    for j := 0 to n - 2 - i do
      if arr[j] > arr[j + 1] then
      begin
        temp := arr[j]; arr[j] := arr[j + 1]; arr[j + 1] := temp;
      end;
end;

begin
  GenerateNumbers(Numbers, 0, 100, 50);
  BubbleSort(Numbers);
  Writeln('Tablica z 50 elementami zostala wygenerowana i posortowana.');
end.