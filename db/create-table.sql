CREATE EXTENSION pgcrypto;  

CREATE TABLE passengers(
  uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  survived boolean,
  pclass int,
  name varchar(100),
  sex varchar(10),
  age float,
  ssa int,
  pca int,
  fare float
);

COPY passengers(survived,pclass,name,sex,age,ssa,pca,fare)
FROM '/tmp/titanic.csv'
WITH (
  FORMAT CSV,
  HEADER true,
  NULL ''
);