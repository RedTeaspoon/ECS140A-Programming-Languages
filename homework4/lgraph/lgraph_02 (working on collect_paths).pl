% /* Graph 1 */
% edge(g1, 0, a, 0).
% edge(g1, 0, b, 1).
% edge(g1, 1, c, 2).
% edge(g1, 1, f, 3).
% edge(g1, 3, k, 4).
% edge(g1, 3, j, 7).
% edge(g1, 4, m, 7).
% edge(g1, 7, l, 6).

% /* Graph 2 */
% edge(g2, 0, a, 0).
% edge(g2, 0, b, 1).
% edge(g2, 1, c, 2).
% edge(g2, 1, f, 3).
% edge(g2, 2, d, 2).
% edge(g2, 2, e, 0).
% edge(g2, 3, g, 4).
% edge(g2, 3, j, 5).
% edge(g2, 4, h, 1).
% edge(g2, 5, l, 6).

% % Start is the same as Target, G1 doesn't contain node
% find_sequence(Graph1, Graph2, Target, Target, 0, Sequence) :-
%     edge(Graph1, Target, _, _) == false; edge(Graph1, _, _, Target) == false,
%     edge(Graph2, Target, _, _) == true; edge(Graph2, _, _, Target) == true,
%     fail.

% Default fail
find_sequence(_, _, _, _, _, _) :- fail.

% Start is the same as Target, G1 contains node & G2 doesn't
% find_sequence(Graph1, Graph2, Target, Target, 0, Sequence) :-
%     edge(Graph1, Target, _, _) == true; edge(Graph1, _, _, Target) == true,
%     edge(Graph2, Target, _, _) == false, edge(Graph2, _, _, Target) == false,
%     Sequence = [[]|Sequence].

find_sequence(Graph1, Graph2, Start, Target, K, Sequence) :-
    K >= 0,
    collect_paths(Graph1, Start, Target, K, [], Sequence).

% Find all paths of length K and append to Paths
collect_paths(Graph1, Start, Target, K, List, Paths) :-
    K >= 0,
    find_paths(Graph1, Start, Target, K, Path),
    collect_paths(Graph1, Start, Target, K, [Path|List], Paths).

% Found all paths, return List
collect_paths(_, _, _, _, List, List).

% No paths found, return fail
collect_paths(_, _, _, _, [], _) :- fail.

find_paths(Graph1, Current, Target, K, [Label|Path]) :-
    K > 0,
    edge(Graph1, Current, Label, Next) == true,  % For every edge from Current
    K1 is K - 1,
    find_paths(Graph1, Next, Target, K1, Path).  % Recursively call find_paths

% If K == 0 & Current == Target, add Path to end of PathList
find_paths(Graph1, Target, Target, 0, []) :-
    edge(Graph1, Target, _, _) == true; edge(Graph1, _, _, Target) == true.  % Make sure that Target is in Graph1
    