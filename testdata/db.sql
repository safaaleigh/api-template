DROP TABLE IF EXISTS todo;

CREATE TABLE todo (
  id uuid NOT NULL default uuid_generate_v4(),
  create_date timestamp with time zone NOT NULL,
  update_date timestamp with time zone NOT NULL,
  name text NOT NULL,
  completed boolean NOT NULL DEFAULT false,
  due timestamp with time zone
);

CREATE TRIGGER todo_insert BEFORE INSERT ON todo FOR EACH ROW EXECUTE PROCEDURE on_insert_column();
CREATE TRIGGER todo_update BEFORE UPDATE ON todo FOR EACH ROW EXECUTE PROCEDURE on_update_column();

INSERT INTO todo (name) VALUES ('Grocery Shopping');
