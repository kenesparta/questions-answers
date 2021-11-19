CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


/* ---------------------------------------------------------------------------------------------------------------------
 *
 *
**/
CREATE TABLE IF NOT EXISTS user_qa
(
    id   uuid        NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
    name varchar(50) NOT NULL
);
COMMENT ON COLUMN user_qa.name IS 'User name';


/* ---------------------------------------------------------------------------------------------------------------------
 *
 *
**/
CREATE TABLE IF NOT EXISTS question
(
    id      uuid NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
    content text NOT NULL,

    user_id uuid NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user_qa (id)
);
COMMENT ON COLUMN question.content IS 'Content of the question';
CREATE INDEX question_user_id ON question (user_id);


/* ---------------------------------------------------------------------------------------------------------------------
 *
 *
**/
CREATE TABLE IF NOT EXISTS answer
(
    id          uuid NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
    content     text NOT NULL,

    user_id     uuid NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user_qa (id),

    question_id uuid NOT NULL,
    FOREIGN KEY (question_id) REFERENCES question (id)
);
CREATE INDEX answer_user_id ON answer (user_id);
CREATE INDEX answer_question_id ON answer (question_id);