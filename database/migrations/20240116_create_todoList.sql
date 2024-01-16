
CREATE TABLE IF NOT EXISTS todoList (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  description TEXT,
  completed BOOLEAN DEFAULT FALSE,
  created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
-- 刪除觸發器（如果存在）
DROP TRIGGER IF EXISTS update_todoList_modtime ON todoList;
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = NOW();
   RETURN NEW;   
END;
$$ language 'plpgsql';

CREATE TRIGGER update_todoList_modtime
  BEFORE UPDATE ON todoList
  FOR EACH ROW
  EXECUTE FUNCTION update_updated_at_column();
