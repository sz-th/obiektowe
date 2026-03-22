program BubbleSortApp;
{$mode objfpc}{$H+}

uses
  sysutils;

type
  TIntArray = array of integer;

var
  Numbers: TIntArray;

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

// 4.5: TESTY JEDNOSTKOWE
procedure AssertTrue(condition: boolean; testName: string);
begin
  if condition then Writeln('[PASS] ', testName)
  else Writeln('[FAIL] ', testName);
end;

procedure RunTests;
var testArr: TIntArray; i: integer; boundsOk: boolean;
begin
  Writeln('--- Uruchamianie testow ---');
  
  // Test 1
  GenerateNumbers(testArr, 10, 20, 5);
  boundsOk := (Length(testArr) = 5);
  for i := 0 to Length(testArr) - 1 do
    if (testArr[i] < 10) or (testArr[i] > 20) then boundsOk := false;
  AssertTrue(boundsOk, 'Test 1: GenerateNumbers poprawny zakres (10-20) i dlugosc (5)');

  // Test 2
  SetLength(testArr, 3); testArr[0]:=3; testArr[1]:=1; testArr[2]:=2;
  BubbleSort(testArr);
  AssertTrue((testArr[0]=1) and (testArr[1]=2) and (testArr[2]=3), 'Test 2: Sortowanie nieuporzadkowanej tablicy');

  // Test 3
  SetLength(testArr, 3); testArr[0]:=1; testArr[1]:=2; testArr[2]:=3;
  BubbleSort(testArr);
  AssertTrue((testArr[0]=1) and (testArr[1]=2) and (testArr[2]=3), 'Test 3: Sortowanie posortowanej tablicy');

  // Test 4
  SetLength(testArr, 3); testArr[0]:=5; testArr[1]:=4; testArr[2]:=0;
  BubbleSort(testArr);
  AssertTrue((testArr[0]=0) and (testArr[1]=4) and (testArr[2]=5), 'Test 4: Sortowanie odwrotnie posortowanej tablicy');

  // Test 5
  SetLength(testArr, 3); testArr[0]:=7; testArr[1]:=7; testArr[2]:=7;
  BubbleSort(testArr);
  AssertTrue((testArr[0]=7) and (testArr[1]=7) and (testArr[2]=7), 'Test 5: Sortowanie tablicy z identycznymi elementami');
  
  Writeln('---------------------------');
end;

begin
  RunTests;
  
  GenerateNumbers(Numbers, 0, 100, 50);
  BubbleSort(Numbers);
  Writeln('Tablica z 50 elementami zostala wygenerowana i posortowana.');
end.