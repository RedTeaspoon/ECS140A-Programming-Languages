
% Default fail
find_sequence(_, _, _, _, _, _) :- fail.

% K not given, find K
find_sequence(Graph1, Graph2, Start, Target, K, Sequence) :-
    var(K),
    not(var(Sequence)) -> length(Sequence, K),
    % print("Find1"),
    find_paths(Graph1, Start, Target, K, Sequence),
    not(find_paths(Graph2, Start, Target, K, Sequence)).

% K given, looking for Sequence
find_sequence(Graph1, Graph2, Start, Target, K, Sequence) :-
    not(var(K)) -> K >= 0,
    % print("Find2"),
    find_paths(Graph1, Start, Target, K, Sequence),
    % print(Sequence),
    not(find_paths(Graph2, Start, Target, K, Sequence)).

% Assume K is given
find_paths(Graph1, Current, Target, K, [Label|Path]) :-
    not(var(K)) -> K > 0,
    edge(Graph1, Current, Label, Next),  % For every edge from Current
    K2 is K - 1,
    find_paths(Graph1, Next, Target, K2, Path).  % Recursively call find_paths

% find_paths(Graph1, Current, Target, K, [_|Path]) :-
%     K > 0,
%     edge(Graph1, Current, _, Next),  % For every edge from Current
%     K2 is K - 1,
%     find_paths(Graph1, Next, Target, K2, Path).  % Recursively call find_paths

% Reached end of path and Current == Target
find_paths(Graph1, Target, Target, 0, []) :-
    edge(Graph1, Target, _, _); edge(Graph1, _, _, Target). 