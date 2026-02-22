function greet() {
  var t = !![];
  var zero = +[];
  var one = +t;
  var two = one + one;
  var three = two + one;
  var four = two + two;
  var seven = three + four;
  var eight = four + four;
  var nine = eight + one;
  var ten = nine + one;

  var fc = ([] + []).constructor.fromCharCode;

  return fc(
    eight * nine,            // 72  H
    ten * ten + one,         // 101 e
    ten * ten + eight,       // 108 l
    ten * ten + eight,       // 108 l
    ten * ten + ten + one,   // 111 o
    four * eight,            // 32  (space)
    eight * ten + seven,     // 87  W
    ten * ten + ten + one,   // 111 o
    ten * ten + ten + four,  // 114 r
    ten * ten + eight,       // 108 l
    ten * ten,               // 100 d
    four * eight + one       // 33  !
  );
}