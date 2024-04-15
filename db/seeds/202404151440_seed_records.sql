INSERT INTO records (name, marks, createdAt)
SELECT name, marks::integer[], createdAt FROM (
    VALUES
    ('John', '{80, 85, 90}'::integer[], NOW()),
    ('Alice', '{75, 82, 88}'::integer[], NOW()),
    ('Bob', '{90, 85, 92}'::integer[], NOW()),
    ('Emily', '{88, 90, 87}'::integer[], NOW()),
    ('Michael', '{82, 84, 86}'::integer[], NOW()),
    ('Sophia', '{91, 92, 95}'::integer[], NOW()),
    ('William', '{89, 86, 83}'::integer[], NOW()),
    ('Emma', '{93, 94, 90}'::integer[], NOW()),
    ('James', '{87, 85, 91}'::integer[], NOW()),
    ('Olivia', '{94, 95, 96}'::integer[], NOW())
) AS s(name, marks, createdAt)
WHERE NOT EXISTS (
    SELECT 1 FROM records r WHERE r.name = s.name
);
