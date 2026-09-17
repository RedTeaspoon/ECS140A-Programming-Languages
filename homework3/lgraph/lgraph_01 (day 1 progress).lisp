;; You may define helper functions here

(defun find-sequence (g1 g2 start target k)
    ;; TODO: Incomplete function
    ;; The next line should not be in your solution.
    (list 'incomplete)

    (let ((g1StartEdges g1(start)) (g1EndExists g1(target)) (g2StartEdges g2(start)) g2EndExists g2(target))
        (cond ((and (equal start target)
                (not (equal (car g1StartEdges) nil))
                (equal (car g2StartExists) nil)) 
                t)
            ((and ()
                ()))
            (t (nil))
            
        s == t && g1StartExists && ! g2StartExists

        )
    )
    
	

	;; check if start & end node are the same + path is at least 0
	;;or nodes in g1 don't exist in g2
	if s == t && g1StartExists && ! g2StartExists {
		return visited, true;
	// check that both start and end nodes exist & path length is long enough
	} else if k >= 1 && g1StartExists && g1EndExists{
)

(defun find-target (g1 target current-edge visited visited-arr)
)

(defun check-path (g2 target current-edge visited-arr)
)

;; check if sequence contains rune label, returns index of edge containing label
defun edges-contains-label (edges, label) int{
	for index, edge := range edges{
		if edge.label == label{
			return index
		}
	}
	return -1
}