const cons = x => list => c => n => c(x)(list(c)(n));

const snoc = x => list => c => n => list(c)(c(x)(n));

const map = f => list => c => n => list(x => c(f(x)))(n);

const filter = p => list => c => n => list(x => rest => p(x)(c(x)(rest))(rest))(n);