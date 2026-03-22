program BubbleSortApp;
{$mode objfpc}{$H+}

uses
  sysutils;

type
  TIntArray = array of integer;

var
  Numbers: TIntArray;

procedure GenerateNumbers(var arr: TIntArray);
var i: integer;
begin
  SetLength(arr, 50);
  Randomize;
  for i := 0 to 49 do
    arr[i] := Random(101);
end;

// 3.5: Procedura do sortowania liczb
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
  GenerateNumbers(Numbers);
  BubbleSort(Numbers); 
  Writeln('Tablica z 50 elementami zostala wygenerowana i posortowana.');
end.