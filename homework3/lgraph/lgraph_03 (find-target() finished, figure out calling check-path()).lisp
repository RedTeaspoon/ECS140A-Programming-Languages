;; You may define helper functions here

;; check if sequence contains rune label, returns index of edge containing label
defun edges-contains-label (edges, label) int{
	for index, edge := range edges{
		if edge.label == label{
			return index
		}
	}
	return -1
}

;; graph g1, target, allowed depth ;; current-edge, nodes visited so far, depth so far
;; make sure to pass in current-edge & target as a list
(defun find-target (g1 current-edge target k visited depth)
  (cond
      ;; at k & reached target
      ((and (equal k depth) (equal target current-edge))
         (list visited)
      )
      ;; at k & not reached target
      ((and (equal k depth) (not (equal target (cdr current-edge))))
         nil
      )
      ;; otherwise for each edge in edges adjacent to current-edge's node apply find-target -> sum up results in list
      (t
         (apply 'append (mapcar 
               (lambda (edge)
                  (find-target g1 (cdr edge) target k (append visited (list (car edge))) (+ depth 1)))
               (g1 (car current-edge)) 
            )
         )
      ) 

   )
)

(defun check-path (g2 current-path target)

)

(defun find-sequence (g1 g2 start target k)
    ;; TODO: Incomplete function
    ;; The next line should not be in your solution.
    (list 'incomplete)

    (let ((g1StartEdges (g1 start)) (g1EndExists (g1 target)) (g2StartEdges (g2 start)) g2EndExists (g2 target))
        ;; start == target, g1Start exists, but g2Start doesn't exist
        (cond ((and (equal start target) (not (null car g1StartEdges)) (null g2StartExists)) 
            (cons nil t))
          ;; k >= 1 && g1StartExists && g1EndExists --> return (cons S nil)
          ((and (zerop k) (not (null g1StartEdges)) (not (null g1EndExists)))
            (let ((paths (find-target g1 (list start) (list target) k nil 0)))
               (cond ((null paths) nil)
                  ((check-path) nil)
               )
            )
          )
          (t nil)
            
          ;; s == t && g1StartExists && ! g2StartExists

         )
    )
    
	

	;; check if start & end node are the same + path is at least 0
	;;or nodes in g1 don't exist in g2
	if s == t && g1StartExists && ! g2StartExists {
		return visited, true;
	// check that both start and end nodes exist & path length is long enough
	} else if k >= 1 && g1StartExists && g1EndExists{
)