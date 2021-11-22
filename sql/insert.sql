DO
$$
    DECLARE
        userId1 CONSTANT uuid := uuid_generate_v4();
        userId2 CONSTANT uuid := uuid_generate_v4();
        userId3 CONSTANT uuid := uuid_generate_v4();
        userId4 CONSTANT uuid := uuid_generate_v4();
    BEGIN

        INSERT
        INTO user_qa
        VALUES (userId1, 'Juliette Woodward'),
               (userId2, 'Nasir Worthington'),
               (userId3, 'Gurveer Rahman'),
               (userId4, 'Lacey Mackie');

        INSERT
        INTO question
        VALUES (DEFAULT, 'Which non-exercise technique has been shown to boost workout results?', userId1),
               (DEFAULT, 'Which building material gets its name from Arabic for "brick"?', userId1),
               (DEFAULT, 'What color was Mario''s shirt in his very first video game appearance?', userId1),
               (DEFAULT, 'What was the better known name of outlaw Harry Longabaugh?', userId1),
               (DEFAULT, 'What US state was named for the future James II of England?', userId2),
               (DEFAULT, 'What was the name of the scandal that forced President Nixon to resign?', userId2),
               (DEFAULT, 'What is the correct name for a baby seal, baby shark, or baby bat?', userId2),
               (DEFAULT, 'Which band released the tracks "Monkey Wrench" and "Learn to Fly"?', userId3),
               (DEFAULT, 'Marilyn Monroe made her first movie appearance in 1947 in which film?', userId3),
               (DEFAULT, 'How many Oscar nominations did the 1995 movie "Apollo 13" receive?', userId4);
    END
$$