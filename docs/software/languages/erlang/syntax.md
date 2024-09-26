# Syntax

`Erlang` files use the `.erl` extension.

```erlang
%% filename = /tmp/tut.erl

-module(tut).

%% The `'/NUMBER` means that the method takes in NUMBER parameters
-import(string, [len/1, concat/2, chr/2, substr/3, str/2, to_lower/1, to_upper/1]).

%% Define functions to export
-export([hello_world/0, add/2]).

%% Function definitions
hello_world() ->
    io:fwrite("Hello World\n").

add(A,B) -> 
    hello_world(),
    A + B.
```

To run the file `/tmp/tutorial.erl`, we can use the `erl` CLI:

```erlang
cd("/tmp/").

%% Name of the file to compile
c(tut).

%% Execute method `hello_world` from `tut` module
tut:hello_world().
Hello World

tut:add(5,4).
Hello World
9

%% Print module information (method signatures, etc)
tut:module_info().
```

### Variables

```erlang
print_one() ->
    Num = 1,
    Num.
```

An `Atom` is a variable which is equal to it's value.

```erlang
atom_stuff() -> 
    'An Atom'.
```

### Math

```erlang
do_math(A,B) -> 
    A + B,
    A - B,
    A * B,
    A div B,
    A rem B,
    math:exp(1),
    math:log(1),
    math:pow(10,2),
    math:sqrt(100),
    random:uniform(10)
    % sin, cos, etc

compare(A,B) -> 
    A =:= B, % returns true/false
    A == B,
    A /= B,
    A > B,
    A <= B,
    Age = 18,
    (Age >= 5) or (Age =< 18).
                  
```

### Conditionals

IF/ELSE:

```erlang
preschool() -> 
    'Go to preschool'.

kindergarten() ->
    'Go to kindergarten'.

grade_school() ->
    'Go to grade school'.

what_grade(X) ->
    if X < 5 -> preschool()
    ; X == 5 -> kindergarten()
    ; X > 5 -> grade_shool()
    end.
```

CASE:

```erlang
say_hello(language) -> 
    case language of
        french -> 'Bonjour';
        german -> 'Guten Tag';
        spanish -> 'Buenos dias';
    end.
```

### Strings

```erlang
string_stuff() -> 
    Str1 = "Random string",
    Str2 = "Another string",
    io:fwrite("String: ~p ~p\n", [Str1, Str2]),

    Str3 = io_lib:format("It's a ~s and ~s\n", [Str1, Str2]),
    io:fwrite(Str3),
    
    len(Str3),
    Str4 = concat(Str1, Str2),
    CharIndex = chr(Str4, $n),
    
    Str5 = substring(Str4, 8, 6),
    
    StrIndex = str(Str4, Str2),
    
    to_upper(Str1),
    to_lower(Str1)
```

### Tuples

```erlang
tuple_stuff() -> 
    My_Data = {42, 175, 6.25},
    
    {A,B,C} = My_Data,
    C,
    
    {D, _, _} = My_Data,
    D,
    
    My_Data_2 = {height, 6.25},
    {height, Ht} = My_Data2,
    Ht,
                
```

### Lists

```erlang
list_stuff() -> 
    List1 = [1,2,3],
    List2 = [4,5,6],
    % Adding/substracting lists
    List3 = List1 ++ List2,
    List4 = List3 -- List1,
    
    % Retrieving head and tail element from list
    hd(List4),
    tl(List4),
    [Head|Tail] = List5,
    
    % Add value to list
    List5 = [3|List4]
```

Comprehensions

```erlang
lc_stuff() -> 
    List1 = [1,2,3],
    
    % Multiply all values from List1 by 2 and store in List2
    List2 = [2*N || N <- List1],
    
    % Even values
    List3 = [1,2,3,4],
    evens = [N || N <- List3, N rem 2 == 0],
    
    % Search a list of tuples
    City_Weather = [{pittsburgh, 50}, {'new york', 53}, {charlotte, 68}, {miami, 78}],
    Great_Temp = [{City, Temp} || {City, Temp} <- City_Weather, Temp >= 65]
```

### Type Conversions

```erlang
type_stuff() ->
    is_atom(name),
    is_float(3.14),
    is_integer(10),
    is_boolean(false),
    is_list([1,2,3]),
    is_tuple({height, 6.24}),
    % Can use the following methods to perform type conversions
    % atom_to_binary(), tuple_to_list(), etc.
```
