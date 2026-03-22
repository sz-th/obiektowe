program BubbleSortApp;
{$mode objfpc}{$H+}

uses
  sysutils;

type
  TIntArray = array of integer;

var
  Numbers: TIntArray;

// 3.0: Procedura na sztywno generująca 50 liczb od 0 do 100
procedure GenerateNumbers(var arr: TIntArray);
var i: integer;
begin
  SetLength(arr, 50);
  Randomize;
  for i := 0 to 49 do
    arr[i] := Random(101);
end;

begin
  GenerateNumbers(Numbers);
  Writeln('Tablica z 50 elementami zostala wygenerowana.');
end.