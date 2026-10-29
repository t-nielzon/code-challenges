const cons = element => list => c => n => c(element)(list(c)(n));

const snoc = element => list => c => n => list(c)(c(element)(n));

const map = transform => list => c => n => list(x => c(transform(x)))(n);

const filter = predicate => list => c => n => 
  list(x => r => predicate(x)(c(x)(r))(r))(n);