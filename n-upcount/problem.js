/*
 * Define the n-upcount of an array to be the number of times the partial sum
 * goes from less than n to greater than or equal to n during the calculation
 * of the sum of the elements of the array. Assume the initial value of the
 * sum is 0.
 *
 * eg.1 if n=6, the 6-upcount of the array [6,3,1] is 1.
 *
 * index    array[index]    partial sum    upcount    comment
 * 0            6            6                1        0 to 6
 * 1            3            9
 * 2            1            10
 *
 * eg.2 the 20-upcount of the array [1,12,-1] is 0.
 *
 * index    array[index]    partial sum    upcount    comment
 * 0            1            1
 * 1            12           13
 * 2            -1           12
 */

function upcount(arr, n) {
}