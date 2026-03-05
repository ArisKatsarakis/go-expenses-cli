CREATE table  income (
        income_id  BIGINT NOT null auto_increment,
        name text not null, 
        money BIGINT NOT NULL,
        PRIMARY KEY(income_id) 
);

ALTER TABLE income ADD COLUMN timestamp datetime not null;
