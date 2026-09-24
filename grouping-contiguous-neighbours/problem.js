/*
 * Grouping Contiguous Neighbours
 * 
 * Write a function that groups contiguous indices based on a merge condition.
 * The function will iterate through indices from firstIndex to lastIndex and use
 * handler functions to determine whether consecutive indices should be merged into
 * a group or treated as single values.
 * 
 * Parameters:
 * - firstIndex: starting index (integer >= 0)
 * - lastIndex: ending index (integer >= firstIndex)
 * - shouldMerge: function(index) returns true/false to merge index with index+1
 * - groupValues: function(firstIndex, lastIndex) called for each group
 * - singleValue: function(index) called for each ungrouped index
 */

function grouper(firstIndex, lastIndex, shouldMerge, groupValues, singleValue) {
  // Your solution here
}