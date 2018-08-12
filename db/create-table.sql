CREATE TABLE passengers
(
  survived boolean,
  pclass int,
  name varchar(100),
  sex varchar(10),
  age float,
  ssa int,
  pca int,
  fare float
);

COPY passengers
FROM '/tmp/titanic.csv'
WITH (
  FORMAT CSV,
  HEADER true,
  NULL ''
);