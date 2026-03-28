/*
Disease Spread

In 1978 the British Medical Journal reported on an outbreak of influenza
at a British boarding school. There were 1000 students. The outbreak
began with one infected student.

We want to study the spread of the disease through the population of
this school. The total population may be divided into three:
the infected (i), those who have recovered (r), and
those who are still susceptible (s) to get the disease.

We will study the disease on a period of tm days. One model of
propagation uses 3 differential equations:

(1) s'(t) = -b * s(t) * i(t)
(2) i'(t) =  b * s(t) * i(t) - a * i(t)
(3) r'(t) =  a * i(t)

Transformed into finite differences:

(I)    S[k+1] = S[k] - dt * b * S[k] * I[k]
(II)   I[k+1] = I[k] + dt * (b * S[k] * I[k] - a * I[k])
(III)  R[k+1] = R[k] + dt * I[k] * a

The function epidemic will return the maximum number of infected
as an integer (truncate to integer the result of max(I)).

Example:
  tm = 14; n = 336; s0 = 996; i0 = 2; b = 0.00206; a = 0.41
  epidemic(tm, n, s0, i0, b, a) --> 483
*/
package kata

func Epidemic(tm, n, s0, i0, b, a float64) int {
}