DO
$$
    DECLARE
        userId1   CONSTANT uuid := uuid_generate_v4();
        userId2   CONSTANT uuid := uuid_generate_v4();
        userId3   CONSTANT uuid := uuid_generate_v4();
        userId4   CONSTANT uuid := uuid_generate_v4();
        question1 CONSTANT uuid := uuid_generate_v4();
        question2 CONSTANT uuid := uuid_generate_v4();
        question3 CONSTANT uuid := uuid_generate_v4();
        question4 CONSTANT uuid := uuid_generate_v4();
    BEGIN

        INSERT
        INTO user_qa
        VALUES (userId1, 'Juliette Woodward'),
               (userId2, 'Nasir Worthington'),
               (userId3, 'Gurveer Rahman'),
               (userId4, 'Lacey Mackie');

        INSERT
        INTO question
        VALUES (question1, 'Which non-exercise technique has been shown to boost workout results?', userId1),
               (question2, 'Which building material gets its name from Arabic for "brick"?', userId1),
               (question3, 'What color was Mario''s shirt in his very first video game appearance?', userId1),
               (question4, 'What was the better known name of outlaw Harry Longabaugh?', userId1),
               (DEFAULT, 'What US state was named for the future James II of England?', userId2),
               (DEFAULT, 'What was the name of the scandal that forced President Nixon to resign?', userId2),
               (DEFAULT, 'What is the correct name for a baby seal, baby shark, or baby bat?', userId2),
               (DEFAULT, 'Which band released the tracks "Monkey Wrench" and "Learn to Fly"?', userId3),
               (DEFAULT, 'Marilyn Monroe made her first movie appearance in 1947 in which film?', userId3),
               (DEFAULT, 'How many Oscar nominations did the 1995 movie "Apollo 13" receive?', userId4);

        INSERT
        INTO answer
        VALUES (DEFAULT, 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam tortor.', userId1, question1),
               (DEFAULT, 'Vestibulum accumsan tristique commodo. Nunc maximus neque vel ornare mollis.', userId1, question2),
               (DEFAULT, 'Curabitur feugiat turpis eu consequat sodales. Integer vel lacus mi.', userId3, question3),
               (DEFAULT, 'Suspendisse potenti. Maecenas vitae mattis sem. Ut faucibus egestas nulla.', userId3, question4);
    END
$$