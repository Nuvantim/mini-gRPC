CREATE SCHEMA IF NOT EXISTS "public";

CREATE  TABLE "public".categories ( 
	id                   integer  NOT NULL GENERATED  BY DEFAULT AS IDENTITY ,
	name                 varchar(100)  NOT NULL  ,
	CONSTRAINT pk_categories PRIMARY KEY ( id )
 );

CREATE  TABLE "public".places ( 
	id                   integer  NOT NULL GENERATED  BY DEFAULT AS IDENTITY ,
	name                 varchar(100)  NOT NULL  ,
	location             varchar(100)  NOT NULL  ,
	description          text  NOT NULL  ,
	entry_fee            bigint  NOT NULL  ,
	CONSTRAINT pk_places PRIMARY KEY ( id )
 );

CREATE  TABLE "public".review ( 
	id                   integer  NOT NULL GENERATED  BY DEFAULT AS IDENTITY ,
	place_id             integer  NOT NULL  ,
	rating               integer DEFAULT 1 NOT NULL  ,
	"comment"            text  NOT NULL  ,
	created_at           timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL  ,
	CONSTRAINT pk_review PRIMARY KEY ( id )
 );

CREATE  TABLE "public".place_has_categories ( 
	id                   integer  NOT NULL GENERATED  BY DEFAULT AS IDENTITY ,
	place_id             integer  NOT NULL  ,
	categories_id        integer  NOT NULL  ,
	CONSTRAINT pk_place_has_categories PRIMARY KEY ( id )
 );

ALTER TABLE "public".place_has_categories ADD CONSTRAINT fk_place_has_categories_categoryid FOREIGN KEY ( categories_id ) REFERENCES "public".categories( id );

ALTER TABLE "public".place_has_categories ADD CONSTRAINT fk_place_has_categories_placesid FOREIGN KEY ( place_id ) REFERENCES "public".places( id );

ALTER TABLE "public".review ADD CONSTRAINT fk_review_places FOREIGN KEY ( place_id ) REFERENCES "public".places( id );

COMMENT ON COLUMN "public".review."comment" IS 'text';
