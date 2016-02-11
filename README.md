# editdistance

## Introduction

This is an experimental Go package that aims to make Kilian Evang's
[golang-levenshtein](https://github.com/texttheater/golang-levenshtein)
package generic in both:

* The edit operations, making it possible to use different distance
  measures (e.g. shortest edit script, Levenshtein distance,
  Levenshtein-Damerau distance).
* The sequence (element) type (e.g. characters, tokens, nucleotides).

Since Go does not support parametric polymorphism, one of the goals is
to evaluate how much genericity is possible without too much performance
loss.
