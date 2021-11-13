CREATE TABLE SpaceZoo (
name VARCHAR(30), 
category VARCHAR(30), 
description VARCHAR(200), 
source VARCHAR(200)
);

INSERT INTO SpaceZoo
VALUES
    ('Art therapy', 'expression', 'Draw your emotions or express your feelings in a painting', 'https://www.healthline.com/health/mental-health/anxiety-drawing#:~:text=Art%20therapy%20refers%20to%20any,have%20trouble%20putting%20into%20words.'),
    ('Therapeutic journaling', 'expression', 'Writing your way through your thoughts is a great way to reflect on your emotions and experiences', 'https://www.va.gov/WHOLEHEALTHLIBRARY/tools/therapeutic-journaling.asp'),
    ('Sand tray therapy', 'relaxation', 'Raking the sand of a mini zen garden gives you a way to express yourself and reduce tension.', 'https://www.healthyplace.com/blogs/anxiety-schmanxiety/2015/07/hey-kids-and-adults-play-with-sand-to-reduce-anxiety'),
    ('Perfect day', 'visualization', 'What''s your idea of a perfect day? Reflecting an imagining on it can help you reduce stress.', 'https://projectlifemastery.com/design-your-perfect-day/'),
    ('Reflect on your good traits', 'reflection', 'What are you good at? What do other people see in you? Reflect on your good traits to improve your self-image', 'https://www.healthyplace.com/blogs/buildingselfesteem/2015/10/identifying-your-good-qualities-when-you-believe-youre-worthless'),
    ('Reflect on your gratitute', 'reflection', 'What are you thankful for? How can you express it?', 'https://resources.wellcertified.com/articles/practicing-gratitude-5-tips-for-positive-reflection/'),
    ('Progressive muscle relaxation', 'relaxation', 'Relax your muscles one group at a time, releasing your stress.', 'https://www.uofmhealth.org/health-library/uz2225'),
    ('Meditation', 'meditation', 'Meditation is one of the best ways to take a break a relax.', 'https://www.mindful.org/how-to-meditate/'),
    ('Roll breathing', 'breathing', 'Lay on your back with one hand on your belly and the other on your chest. Feel your breath move through you like a wave as you breathe in and out.', 'https://myhealth.alberta.ca/Health/Pages/conditions.aspx?hwid=ug4360'),
    ('Square breathing', 'breathing', 'Inhale through your nose for 4 seconds, hold for 4 seconds, exhale through your mouth for 4 seconds, hold for 4 seconds, and repeat.', 'https://blog.zencare.co/square-breathing/'),
    ('Resonant breathing', 'breathing', 'Lie down, eyes shut, and breathe in slowly through your nose. Exhale slowly and repeat.', 'https://www.rosalbacourtney.com/resonance-frequency-breathing/#:~:text=Resonance%20frequency%20breathing%20is%20a,such%20as%20the%20circulatory%20system.'),
    ('Visiualization meditation', 'meditation', 'Picturing positive images and relaxing your body is a good way to meditate if you struggle with traditional meditation.', 'https://www.wellandgood.com/visualization-meditation/'),
    ('Happy Place visualization', 'visualization', 'Imagine a peaceful place, like a beach or the woods. Imagine with all your senses: how it sounds, how it feels, how it smells.', 'https://www.mondaycampaigns.org/destress-monday/find-happy-place'),
    ('Music therapy', 'expression', 'Singing a song or playing an instrument is a great way to express your emotions and release stress.', 'https://www.healthline.com/health/sound-healing'),
    ('Relaxing soundscapes', 'relaxation', 'A relaxing soundscape, whether it''s the beach, traffic, or birds chirping in the woods, can help put your mind at ease.', 'https://www.drjohnlapuma.com/naturetherapy/what-is-sound-or-music-therapy-what-are-soundscapes/'),
    ('Self-reflection', 'reflection', 'Reflecting on what your priorities are and how you are working towards them can help you identify sources of stress and work towards reducing them.', 'https://www.skillsyouneed.com/ps/reflective-practice.html')
;