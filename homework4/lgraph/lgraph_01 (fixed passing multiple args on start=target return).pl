

% Start is the same as Target, G1 doesn't contain node
find_sequence(Graph1, Graph2, Target, Target, 0, Sequence) :-
    edge(Graph1, Target, _, _) == false; edge(Graph1, _, _, Target) == false,
    edge(Graph2, Target, _, _) == true; edge(Graph2, _, _, Target) == true,
    fail.

% Start is the same as Target, G2 contains node
find_sequence(Graph1, Graph2, Target, Target, 0, Sequence) :-
    edge(Graph1, Target, _, _) == false; edge(Graph1, _, _, Target) == false,
    fail.

% Start is the same as Target, G1 contains node & G2 doesn't
find_sequence(Graph1, Graph2, Target, Target, 0, Sequence) :-
    edge(Graph1, Target, _, _) == true; edge(Graph1, _, _, Target) == true,
    edge(Graph2, Target, _, _) == false, edge(Graph2, _, _, Target) == false,
    % length(X, 1), X=[[]],
    % print(Sequence),
    % length(Sequence,0),
    Sequence = [].

find_sequence(Graph1, Graph2, Start, Target, K, Sequence) :-
    %% TODO: remove fail and add body/other cases for this predicate
    Start \= Target,
    print("test2"),
    fail.