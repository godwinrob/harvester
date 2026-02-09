INSERT INTO users (user_id, name, email, roles, password_hash, guild, enabled, date_created, date_updated) VALUES
    ('5cf37266-3473-4006-984f-9325122678b7', 'Luke Skywalker', 'admin@example.com', '{ADMIN}', '$2a$10$1ggfMVZV6Js0ybvJufLRUOWHS5f6KneuP0XwwHpJ8L8ipdry9f2/a', NULL, true, '2019-03-24 00:00:00', '2019-03-24 00:00:00'),
    ('45b5fbd3-755f-4379-8f07-a58d4a30fa2f', 'Darth Vader', 'user@example.com', '{USER}', '$2a$10$9/XASPKBbJKVfCAZKDH.UuhsuALDr5vVm6VrYA9VFR8rccK86C1hW', NULL, true, '2019-03-24 00:00:00', '2019-03-24 00:00:00')
ON CONFLICT DO NOTHING;

INSERT INTO galaxies (galaxy_id, galaxy_name, owner_user_id) VALUES
    ('681672b7-95a8-4871-8832-e5774799c0e3', 'Finalizer', '5cf37266-3473-4006-984f-9325122678b7'),
    ('b4629864-500c-4f06-8e4a-d31cea7bcfae', 'Bria', '45b5fbd3-755f-4379-8f07-a58d4a30fa2f')
ON CONFLICT DO NOTHING;

INSERT INTO resources (
    resource_id,
    resource_name,
    galaxy_id,
    added_user_id,
    resource_type,
    cr,
    cd,
    dr,
    fl,
    "hr",
    ma,
    pe,
    oq,
    sr,
    ut,
    er)
VALUES (
    '47318be2-049e-4029-9c59-a04e1bae0b01',
    'akiium',
    '681672b7-95a8-4871-8832-e5774799c0e3',
    '5cf37266-3473-4006-984f-9325122678b7',
    'petrochem_fuel_liquid_type1',
    0,
    0,
    3,
    0,
    0,
    0,
    991,
    23,
    0,
    0,
    0),
    ('9e453180-621f-49a0-90c6-bf7a5e2a419b',
    'havai',
    '681672b7-95a8-4871-8832-e5774799c0e3',
    '5cf37266-3473-4006-984f-9325122678b7',
    'iron_kammris',
    0,
    747,
    461,
    935,
    0,
    986,
    546,
    0,
    40,
    980,
    913)
ON CONFLICT DO NOTHING;