/*
  Given the following function definition, we can construct 2-dimensional matrices
  that are infinite in both dimensions.

  type Matrix = [[Bool]]

  generate :: Int -> Int -> Matrix
  generate n m =
    let falses = repeat False
        oneTrue = replicate m False ++ [ True ] ++ falses
    in replicate n falses ++ [oneTrue] ++ repeat falses

  This generates a matrix where every element is False except the one under indices n, m.
  Write a function that takes such a matrix and returns the pair of indices where the
  matrix is True.

  For JavaScript: The kata is implemented via ES6 proxies in the preloaded function generate:

  function generate(y,x) {
    const colHandler = row => ({ get: (_,col) => Number.isFinite(Number(col)) && Number(col) >= 0 ?
                                                   row === y && Number(col) === x :
                                                   undefined ,
                                 set: () => false });
    const rowHandler =         { get: (_,row) => Number.isFinite(Number(row)) && Number(row) >= 0 ?
                                                   new Proxy( {}, colHandler(Number(row)) ) :
                                                   undefined ,
                                 set: () => false };
    return new Proxy( {}, rowHandler );
  }

  You can use mat[i][j] normally to access "matrix" elements, which will be all false
  except when i = y and j = x. However, proxies behave entirely differently from arrays,
  so regular array methods will not work at all.
*/

function find(matrix) {
}