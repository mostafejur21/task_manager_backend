ALTER TABLE
    tasks
ADD
    COLUMN description varchar(255) NOT NULL;

ALTER TABLE
    tasks
ADD
    COLUMN status text;

ALTER TABLE
    tasks
ADD
  COLUMN updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW();

