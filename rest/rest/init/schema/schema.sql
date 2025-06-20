CREATE SCHEMA IF NOT EXISTS "public";

CREATE  TABLE "public".category ( 
	id                   integer  NOT NULL  ,
	name                 varchar(100)  NOT NULL  ,
	created_at           timestamptz DEFAULT CURRENT_TIMESTAMP   ,
	CONSTRAINT pk_category PRIMARY KEY ( id )
 );

CREATE  TABLE "public".product ( 
	id                   integer  NOT NULL  ,
	name                 varchar(100)  NOT NULL  ,
	description          text  NOT NULL  ,
	category_id          integer  NOT NULL  ,
	price                integer DEFAULT 0 NOT NULL  ,
	created_at           timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL  ,
	CONSTRAINT pk_product PRIMARY KEY ( id ),
	CONSTRAINT fk_product_category FOREIGN KEY ( category_id ) REFERENCES "public".category( id )   
 );
