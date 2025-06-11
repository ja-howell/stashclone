DROP TABLE IF EXISTS stashitems;
CREATE TABLE stashitems (
  id         INT AUTO_INCREMENT NOT NULL,
  item_name      VARCHAR(128) NOT NULL,
  item_type     VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO stashitems
  (item_name, item_type)
VALUES
  ('Blue', 'Fabric'),
  ('Square Ruler', 'Ruler'),
  ('Purple Scissors', 'Scissor'),
  ('Pattern Star', 'Pattern');