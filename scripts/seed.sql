INSERT INTO class_sessions(id,course,classroom,starts_at) VALUES ('demo-print','版画入门','Art-203',now()+interval '2 hours') ON CONFLICT DO NOTHING;
